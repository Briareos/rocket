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

	Groups   []*Group  `json:"groups"`
	Statuses []*Status `json:"statuses"`
}

type UserService interface {
	Get(int) (*User, error)
}