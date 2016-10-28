package handle

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"encoding/json"
)

func GoogleOAuth(clientID, redirect string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//next := r.FormValue("next")
		//tok := r.Context().Value(request.Token).(*rocket.Token)
		//url := authUrl("https://raketa.rocks")
		http.Redirect(w, r, authUrl(clientID, redirect), 302)
	}
}

type userInfo struct {
	email string `json:"email"`
}

func GoogleOAuthCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request) {
		code := r.FormValue("code")
		if code == "" {
			http.Error(w, `The "code" parameter is missing`, http.StatusInternalServerError)
			return
		}
		spew.Dump(code)
		res, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="+code)
		if err != nil {
			http.Error(w, `The service is currently unavailable`, http.StatusServiceUnavailable)
			return
		}
		body, err :=ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, `The service is currently unavailable`, http.StatusServiceUnavailable)
			return
		}
		res.Body.Close()
		info := new(userInfo)
		if err := json.Unmarshal(body, info); err != nil {
			http.Error(w, `The service is currently unavailable`, http.StatusServiceUnavailable)
			return
		}
		spew.Dump(body)

		// todo: find out why we get 401
		// if user exists - log them in
		// if user doesn't exist - create them
	}
}

var e = html.EscapeString
var scopes = e("https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")

func authUrl(clientID, redirect string) string {
	return fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?response_type=code&scope=%s&client_id=%s&redirect_uri=%s&prompt=consent", scopes, e(clientID), e(redirect))
}
