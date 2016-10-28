package handle

import (
	"encoding/json"
	"fmt"
	"github.com/Briareos/rocket"
	"net/http"
	"github.com/Briareos/rocket/request"
)

type ProfileUser struct {
	*rocket.User
	JoinedGroups  []int       `json:"joined_groups"`
	WatchedGroups []int       `json:"watched_groups"`
	MutedRules    []int       `json:"muted_rules"`
	ActiveRules   map[int]int `json:"active_rules"` //Map for each group (key) how many alerts are active (value)
}

type ProfileResponse struct {
	User   ProfileUser     `json:"user"`
	Groups []*rocket.Group `json:"groups"`
	Users  []*rocket.User  `json:"users"`
}

func Profile(userService rocket.UserService, groupService rocket.GroupService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := request.GetToken(r).User()
		if user == nil {
			http.Error(w, "you must log in first", http.StatusForbidden)
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

		responseData := ProfileResponse{
			User: ProfileUser{
				User:          user,
				JoinedGroups:  joinedGroups,
				WatchedGroups: watchedGroups,
				MutedRules:    mutedRules,
			},
			Groups: allGroups,
			Users:  allUsers,
		}

		data, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	})
}
