package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Briareos/rocket"
	"github.com/Briareos/rocket/container"
	"github.com/Briareos/rocket/sql"
)

type profileResponse struct {
	user struct {
		rocket.User
		joinedGroups  []int       `json:"joined_groups"`
		watchedGroups []int       `json:"watched_groups"`
		mutedRules    []int       `json:"muted_rules"`
		active_rules  map[int]int `json:"active_rules"` //Map for each group (key) how many alerts are active (value)
	} `json:"user"`
	groups []rocket.Group `json:"groups"`
	users  []rocket.User  `json:"users"`
}

func makeProfileFunc(userService rocket.UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := userService.Get(1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	})
}

func makeGroupDays() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func main() {
	c := container.MustLoadFromPath(filepath.Join("..", "..", "config.yml"))
	c.MustWarmUp()

	userService := sql.NewUserService(c.DB())

	c.HTTPHandler().Handle("/api/v1/profile", makeProfileFunc(userService))
	c.HTTPHandler().Handle("/api/v1/groupDays", makeGroupDays())

	c.HTTPServer().ListenAndServe()
}
