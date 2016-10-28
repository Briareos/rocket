package rocket

import "time"

const (
	StatusAvailable   = "available"
	StatusUnavailable = "unavailable"
	StatusBusy        = "busy"
	StatusRemote      = "remote"
)

type Status struct {
	ID     int       `json:"id"`
	Type   string    `json:"type"`
	Reason string    `json:"reason"`
	Date   time.Time `json:"date"`
}

type User struct {
	ID        int    `json:"id"`
	GoogleID  string `json:"google_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`

	JoinedGroups  []*Group  `json:"-"` // indexed
	WatchedGroups []*Group  `json:"-"` // indexed
	Statuses      []*Status `json:"statuses,omitempty"`
	MutedRules    []*Rule   `json:"-"` // indexed
}

type UserService interface {
	Get(int) (*User, error)
	GetAll() ([]*User, error)
	Add(*User) error

	//JoinGroup(*User, *Group) error
	//LeaveGroup(*User, *Group) error
	//
	//WatchGroup(*User, *Group) error
	//UnWatchGroup(*User, *Group) error
	//
	//MuteRule(*User, *Rule) error
	//UnMuteRule(*User, *Rule) error
}
