package response

import "github.com/Briareos/rocket"

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
