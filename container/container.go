package container

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"reflect"
	"github.com/Briareos/rocket/handle"
)

type Config struct {
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBHost     string `yaml:"db_host"`
	DBName     string `yaml:"db_name"`
	HTTPAddr   string `yaml:"http_addr"`
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
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", c.conf.DBUser, c.conf.DBPassword, c.conf.DBHost, c.conf.DBName))
		if err != nil {
			panic(errors.New(`container: failed to connect to mysql server`))
		}
		return db
	}).(*sql.DB)
}

func (c *Container) HTTPHandler() *http.ServeMux {
	return c.once.Do("HTTPHandler", func() interface{} {
		http.HandleFunc("/", handle.Index())
		return http.DefaultServeMux
	}).(*http.ServeMux)
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
