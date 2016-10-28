package rocket

import "time"

var DefaultAvailability = AvailabilityMap{
	Available:   true,
	Remote:      true,
	Busy:        false,
	Unavailable: false,
}

type AvailabilityMap struct {
	Available   bool `json:"available"`
	Unavailable bool `json:"unavailable"`
	Busy        bool `json:"busy"`
	Remote      bool `json:"remote"`
}

type Group struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Availability AvailabilityMap `json:"availability"`

	Rules []*Rule `json:"rules"`
}

type GroupService interface {
	Get(int) (*Group, error)

	Add(*Group) error
	GetAll() ([]*Group, error)

	// AddRule creates a Rule and assigns it to the Group
	AddRule(*Group, *Rule) error

	// GetBodyCount returns the number of available people
	GetAvailableBodyCounts(*Group, time.Time) (map[time.Time]int, error)

	// GetTotalBodyCount returns the total number of people in group
	GetTotalBodyCount(*Group) (int, error)
}
