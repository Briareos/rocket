package handle

import (
	"github.com/Briareos/rocket"
	"net/http"
	"encoding/json"
)

type GroupActionType string

const (
	Join GroupActionType = "join"
	Leave GroupActionType = "leave"
	Watch GroupActionType = "watch"
	UnWatch GroupActionType = "unwatch"
)

type GroupActionRequest struct {
	GroupID int `json:"groupID"`
	Type GroupActionType `json:"type"`
}

func GroupAction(userService rocket.UserService, groupService rocket.GroupService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)

		var requestBody GroupActionRequest

		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		if requestBody.Type != Join &&
			requestBody.Type != Leave &&
			requestBody.Type != Watch &&
			requestBody.Type != UnWatch {
			http.Error(w, "Missing parameter 'type'.", http.StatusInternalServerError)
			return
		}

		group, err := groupService.Get(requestBody.GroupID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// TODO: Losmi spoji sa sesijom
		user, err := userService.Get(1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if requestBody.Type == Join {
			err = userService.JoinGroup(user, group)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if requestBody.Type == Leave {
			err = userService.LeaveGroup(user, group)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if requestBody.Type == Watch {
			err = userService.WatchGroup(user, group)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if requestBody.Type == UnWatch {
			err = userService.UnWatchGroup(user, group)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	})
}
