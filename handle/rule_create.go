package handle

import (
	"encoding/json"
	"github.com/Briareos/rocket"
	"net/http"
)

type RuleCreateRequest struct {
	GroupID int         `json:"groupID"`
	Rule    rocket.Rule `json:"rule"`
}

func RuleCreate(groupService rocket.GroupService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)

		var requestBody RuleCreateRequest

		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		group, err := groupService.Get(requestBody.GroupID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = groupService.AddRule(group, &requestBody.Rule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
