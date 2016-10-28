package handle

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"net/url"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
)

func GoogleOAuth(clientID, redirect string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//next := r.FormValue("next")
		//tok := r.Context().Value(request.Token).(*rocket.Token)
		//url := authUrl("https://raketa.rocks")
		http.Redirect(w, r, authUrl(clientID, redirect), 302)
	}
}

func GoogleOAuthCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request) {
		code := r.FormValue("code")
		if code == "" {
			http.Error(w, `The "code" parameter is missing`, http.StatusInternalServerError)
			return
		}
		res, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=" + url.QueryEscape(code))
		if err != nil {
			http.Error(w, `The service is currently unavailable`, http.StatusServiceUnavailable)
			return
		}
		spew.Dump(ioutil.ReadAll(res.Body))
		res.Body.Close()
	}
}

var e = html.EscapeString
var scopes = e(strings.Join([]string{"https://www.googleapis.com/auth/userinfo.email"}, "|"))

func authUrl(clientID, redirect string) string {
	return fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?response_type=code&scope=%s&client_id=%s&redirect_uri=%s&prompt=consent", scopes, e(clientID), e(redirect))
}
