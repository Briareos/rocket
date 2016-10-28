package handle

import (
	"github.com/Briareos/rocket"
	"time"
	"net/http"
	"strconv"
	"fmt"
	"encoding/json"
)

type GroupDaysResponse struct {
	Group rocket.Group
	Days  []Day
}

type Day struct {
	Date               time.Time
	AvailableBodyCount int           `json:"availableBodyCount"`
	ActiveRules        []rocket.Rule `json:"activeRules"`
}

func GroupDays(groupService rocket.GroupService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if len(id) == 0 {
			http.Error(w, "ID not provided correctly", http.StatusInternalServerError)
			return
		}

		idNumber, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Month not provided correctly", http.StatusInternalServerError)
			return
		}

		group, err := groupService.Get(idNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		month := r.URL.Query().Get("month")
		if len(month) == 0 {
			http.Error(w, "Month not provided correctly", http.StatusInternalServerError)
			return
		}

		monthNumber, err := strconv.Atoi(month)
		if err != nil {
			http.Error(w, "Month not provided correctly", http.StatusInternalServerError)
			return
		}

		year := r.URL.Query().Get("year")
		if len(year) == 0 {
			http.Error(w, "Year not provided correctly", http.StatusInternalServerError)
			return
		}

		yearNumber, err := strconv.Atoi(year)
		if err != nil {
			http.Error(w, "Year not provided correctly", http.StatusInternalServerError)
			return
		}

		var groupDays GroupDaysResponse

		// First day of month in given year
		date := time.Date(yearNumber, time.Month(monthNumber), 1, 0, 0, 0, 0, time.UTC)
		daysInMonth := time.Date(yearNumber, time.Month(monthNumber)+1, 0, 0, 0, 0, 0, time.UTC).Day()
		availableBodyCounts, err := groupService.GetAvailableBodyCounts(group, date)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i := 1; i <= daysInMonth; i++ {
			if _, ok := availableBodyCounts[date]; !ok {
				availableBodyCounts[date] = 0
			}
			// Move to next day
			date = date.AddDate(0, 0, 1)
		}

		for date, count := range availableBodyCounts {
			groupDays.Days = append(groupDays.Days, Day{
				Date:               date,
				AvailableBodyCount: count,
			})
		}

		data, err := json.Marshal(groupDays)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(data))

	})
}
