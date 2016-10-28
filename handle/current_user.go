package handle

import (
	"encoding/json"
	"github.com/Briareos/rocket"
	"github.com/Briareos/rocket/request"
	"net/http"
)

func CurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		tok := r.Context().Value(request.Token).(*rocket.Token)
		if tok.User == nil {
			enc.Encode(map[string]interface{}{
				"ok":    false,
				"error": "logged_out",
			})
		} else {
			enc.Encode(map[string]interface{}{
				"ok":   true,
				"user": tok.User,
			})
		}
	}
}
