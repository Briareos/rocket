package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Briareos/rocket"
	"github.com/Briareos/rocket/container"
	"strconv"
	"time"
	"github.com/Briareos/rocket/response"
	"github.com/Briareos/rocket/sql"
)

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

		responseData := response.ProfileResponse{
			User: response.ProfileUser{
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

func makeGroupDays(groupService rocket.GroupService) http.Handler {
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

		var groupDays response.GroupDays

		// First day of month in given year
		date := time.Date(yearNumber, time.Month(monthNumber), 1, 0, 0, 0, 0, time.UTC)
		daysInMonth := time.Date(yearNumber, time.Month(monthNumber)+1, 0, 0, 0, 0, 0, time.UTC).Day()
		availableBodyCounts, err := groupService.GetAvailableBodyCounts(group, date)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(date)
		for i := 1; i <= daysInMonth; i++ {
			if _, ok := availableBodyCounts[date]; !ok {
				availableBodyCounts[date] = 0
			}
			// Move to next day
			date = date.AddDate(0, 0, 1)
		}

		for date, count := range availableBodyCounts {
			groupDays.Days = append(groupDays.Days, response.Day{
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

func main() {
	c := container.MustLoadFromPath(filepath.Join("..", "..", "config.yml"))
	c.MustWarmUp()

	userService := sql.NewUserService(c.DB())
	groupService := sql.NewGroupService(c.DB())

	c.HTTPHandler().Handle("/api/v1/profile", makeProfileFunc(userService, groupService))
	c.HTTPHandler().Handle("/api/v1/groupDays", makeGroupDays(groupService))

	c.HTTPServer().ListenAndServe()
}
