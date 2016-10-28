package response

import (
	"github.com/Briareos/rocket"
	"time"
)

type GroupDays struct {
	Group rocket.Group
	Days  []Day
}

type Day struct {
	Date               time.Time
	AvailableBodyCount int           `json:"availableBodyCount"`
	ActiveRules        []rocket.Rule `json:"activeRules"`
}
