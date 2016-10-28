package handle

import (
	"encoding/json"
	"github.com/Briareos/rocket"
	"net/http"
)

type RuleActionType string

const (
	Mute   RuleActionType = "mute"
	UnMute RuleActionType = "unmute"
)

type RuleActionRequest struct {
	RuleID int            `json:"ruleID"`
	Type   RuleActionType `json:"type"`
}

func RuleAction(userService rocket.UserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)

		var requestBody RuleActionRequest

		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// TODO: Losmi spoji sa sesijom
		//user, err := userService.Get(1)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//
		//if requestBody.Type == Mute {
		//	err := userService.MuteRule(user, requestBody.RuleID)
		//	if err != nil {
		//		http.Error(w, err.Error(), http.StatusInternalServerError)
		//		return
		//	}
		//}
		//
		//if requestBody.Type == UnMute {
		//	err := userService.UnMuteRule(user, requestBody.RuleID)
		//	if err != nil {
		//		http.Error(w, err.Error(), http.StatusInternalServerError)
		//		return
		//	}
		//}
	})
}
