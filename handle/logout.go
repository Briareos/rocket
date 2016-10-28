package handle

import (
	"net/http"
	"github.com/Briareos/rocket/request"
)

func LogoutAndRedirect(home string) http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request) {
		request.GetToken(r).SetUser(nil)
		http.Redirect(w, r, home, 302)
	}
}
