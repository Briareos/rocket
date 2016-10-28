package handle

import (
	"encoding/json"
	"fmt"
	"github.com/Briareos/rocket"
	"net/http"
)

type GroupCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	BusyValue   bool `json:"busyValue"`
	RemoteValue bool `json:"remoteValue"`
}

func GroupCreate(groupService rocket.GroupService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)

		var requestBody GroupCreateRequest

		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		group := &rocket.Group{
			Name:         requestBody.Name,
			Description:  requestBody.Description,
			Availability: rocket.DefaultAvailability,
		}

		group.Availability.Busy = requestBody.BusyValue
		group.Availability.Remote = requestBody.RemoteValue

		err = groupService.Add(group)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(group)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	})
}
