package container

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Briareos/rocket"
	"github.com/Briareos/rocket/handle"
	"github.com/Briareos/rocket/request"
	oursql "github.com/Briareos/rocket/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"
)

type Config struct {
	DBUser            string `yaml:"db_user"`
	DBPassword        string `yaml:"db_password"`
	DBHost            string `yaml:"db_host"`
	DBName            string `yaml:"db_name"`
	HTTPAddr          string `yaml:"http_addr"`
	GoogleOAuthID     string `yaml:"google_oauth_id"`
	GoogleOAuthSecret string `yaml:"google_oauth_secret"`
	Secret            string `yaml:"secret"`
}

func LoadConfig(path string) (*Config, error) {
	conf := new(Config)
	if err := configor.Load(conf, path); err != nil {
		return nil, errors.Wrapf(err, `container: failed to load config`)
	}
	return conf, nil
}

func NewFromConfig(conf *Config) *Container {
	return &Container{
		conf: conf,
		once: new(once),
	}
}

func LoadFromPath(confPath string) (*Container, error) {
	conf, err := LoadConfig(confPath)
	if err != nil {
		return nil, errors.Wrapf(err, `container: failed to create container from config file %s`, confPath)
	}
	return NewFromConfig(conf), nil
}

func MustLoadFromPath(confPath string) *Container {
	conf, err := LoadFromPath(confPath)
	if err != nil {
		panic(err)
	}
	return conf
}

type Container struct {
	conf *Config
	once *once
}

func (c *Container) Params() *Config {
	return c.conf
}

// WarmUp runs each getter once to create the services in memory.
// It also captures panic and wraps it in an error.
// Multiple calls to WarmUp are OK, but only makes sense if some
// network was unreachable and it caused a panic previously.
func (c *Container) WarmUp() (warmUpErr error) {
	defer func() {
		switch rec := recover().(type) {
		case nil:
		// Do nothing by default.
		case error:
			warmUpErr = errors.Wrap(rec, `container: warm-up panicked with error`)
		default:
			warmUpErr = errors.Errorf(`container: warm-up panicked with value %+v`, rec)
		}
	}()
	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)
	arg := []reflect.Value{}
	for i := 0; i < val.NumMethod(); i++ {
		if name := typ.Method(i).Name; name == "WarmUp" || name == "MustWarmUp" {
			continue
		} else {
			log.Printf("container: warming up service %s\n", name)
			val.Method(i).Call(arg)
		}
	}
	return
}

func (c *Container) MustWarmUp() {
	if err := c.WarmUp(); err != nil {
		panic(err)
	}
}

func (c *Container) DB() *sql.DB {
	return c.once.Do("DB", func() interface{} {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", c.conf.DBUser, c.conf.DBPassword, c.conf.DBHost, c.conf.DBName))
		if err != nil {
			panic(errors.New(`container: failed to connect to mysql server`))
		}
		return db
	}).(*sql.DB)
}

func (c *Container) UserService() rocket.UserService {
	return c.once.Do("UserService", func() interface{} {
		return oursql.NewUserService(c.DB())
	}).(rocket.UserService)
}

func (c *Container) GroupService() rocket.GroupService {
	return c.once.Do("GroupService", func() interface{} {
		return oursql.NewGroupService(c.DB())
	}).(rocket.GroupService)
}

func (c *Container) RuleService() rocket.RuleService {
	return c.once.Do("RuleService", func() interface{} {
		return oursql.NewRuleService(c.DB())
	}).(rocket.RuleService)
}

func (c *Container) HomeURL() string {
	return c.once.Do("HomeURL", func() interface{} {
		host, port, err := net.SplitHostPort(c.conf.HTTPAddr)
		if err != nil {
			panic(errors.Wrap(err, `container: failed to parse HTTP address`))
		}
		if host == "" {
			host = "localhost"
		}
		proto := "http://"
		if port == "http" {
			port = ""
		} else if port == "https" {
			proto = "https://"
			port = ""
		} else if port != "" {
			port = ":" + port
		}
		return strings.TrimRight(proto+host+port, "/")
	}).(string)
}

func (c *Container) HTTPHandler() *http.ServeMux {
	return c.once.Do("HTTPHandler", func() interface{} {
		redirectURI := c.HomeURL() + "/oauth/google/callback"
		http.HandleFunc("/oauth/google/callback", c.makeHandle(handle.GoogleOAuthCallback(c.conf.GoogleOAuthID, c.conf.GoogleOAuthSecret, redirectURI)))
		http.HandleFunc("/oauth/google", c.makeHandle(handle.GoogleOAuth(c.conf.GoogleOAuthID, redirectURI)))
		http.HandleFunc("/api/current-user", c.makeHandle(handle.CurrentUser()))
		http.HandleFunc("/api/profile", c.makeHandle(handle.Profile(c.UserService(), c.GroupService())))

		http.HandleFunc("/api/groupDays", c.makeHandle(handle.GroupDays(c.GroupService())))
		http.HandleFunc("/api/groupCreate", c.makeHandle(handle.GroupCreate(c.GroupService())))
		http.HandleFunc("/api/groupAction", c.makeHandle(handle.GroupAction(c.UserService(), c.GroupService())))

		http.HandleFunc("/api/ruleCreate", c.makeHandle(handle.RuleCreate(c.GroupService())))
		http.HandleFunc("/api/ruleAction", c.makeHandle(handle.RuleAction(c.RuleService(), c.UserService())))

		http.HandleFunc("/", c.makeHandle(handle.Index()))
		return http.DefaultServeMux
	}).(*http.ServeMux)
}

func (c *Container) makeHandle(h http.HandlerFunc) http.HandlerFunc {
	return c.enableCORS(c.injectToken(h))
}

func (c *Container) injectToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := c.Session().Get(r, "session")
		if err != nil {
			http.Error(w, "Invalid session provided", 500)
			return
		}
		tok := rocket.NewToken(session)
		//tok.SetUser(user)
		r = r.WithContext(context.WithValue(r.Context(), request.Token, tok))
		h(w, r)
		session.Store()
	}
}

func (c *Container) enableCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}

func (c *Container) HTTPServer() *http.Server {
	return c.once.Do("HTTPServer", func() interface{} {
		srv := &http.Server{
			Addr:    c.conf.HTTPAddr,
			Handler: c.HTTPHandler(),
		}
		return srv
	}).(*http.Server)
}

func (c *Container) Session() *sessions.CookieStore {
	return c.once.Do("Session", func() interface{} {
		return sessions.NewCookieStore([]byte("something-very-secret"))
	}).(*sessions.CookieStore)
}
