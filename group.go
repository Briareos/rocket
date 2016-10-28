package rocket

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
	Add(*Group) error
	GetAll() ([]*Group, error)

	// AddRule creates a Rule and assigns it to the Group
	AddRule(*Group, *Rule) error
}
