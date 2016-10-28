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

type ProfileUser struct {
	*rocket.User
	JoinedGroups  []int       `json:"joined_groups"`
	WatchedGroups []int       `json:"watched_groups"`
	MutedRules    []int       `json:"muted_rules"`
	ActiveRules   map[int]int `json:"active_rules"` //Map for each group (key) how many alerts are active (value)
}

type profileResponse struct {
	User   ProfileUser     `json:"user"`
	Groups []*rocket.Group `json:"groups"`
	Users  []*rocket.User  `json:"users"`
}

func makeProfileFunc(userService rocket.UserService, groupService rocket.GroupService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := userService.Get(1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		joinedGroups := []int{}
		watchedGroups := []int{}
		mutedRules := []int{}

		for _, group := range user.JoinedGroups {
			joinedGroups = append(joinedGroups, group.ID)
		}
		for _, group := range user.WatchedGroups {
			watchedGroups = append(watchedGroups, group.ID)
		}
		for _, rule := range user.MutedRules {
			mutedRules = append(mutedRules, rule.ID)
		}

		allUsers, err := userService.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		allGroups, err := groupService.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseData := profileResponse{
			User: ProfileUser{
				User:          user,
				JoinedGroups:  joinedGroups,
				WatchedGroups: watchedGroups,
				MutedRules:    mutedRules,
			},
			Groups: allGroups,
			Users:  allUsers,
		}

		fmt.Printf("%#v\n", responseData)

		data, err := json.Marshal(responseData)
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
	groupService := sql.NewGroupService(c.DB())

	c.HTTPHandler().Handle("/api/v1/profile", makeProfileFunc(userService, groupService))
	c.HTTPHandler().Handle("/api/v1/groupDays", makeGroupDays())

	c.HTTPServer().ListenAndServe()
}
