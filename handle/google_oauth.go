package handle

import (
	"fmt"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"encoding/json"
	"net/url"
)

func GoogleOAuth(clientID, redirect string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//next := r.FormValue("next")
		//tok := r.Context().Value(request.Token).(*rocket.Token)
		//url := authUrl("https://raketa.rocks")
		http.Redirect(w, r, authURL(clientID, redirect), 302)
	}
}

type userInfo struct {
	email string `json:"email"`
}

func GoogleOAuthCallback(clientID, clientSecret, redirect string) http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request) {
		code := r.FormValue("code")
		if code == "" {
			http.Error(w, `The "code" parameter is missing`, http.StatusInternalServerError)
			return
		}
		res, err := http.Get(tokenURL(code, clientID, clientSecret, redirect))
		if err != nil {
			http.Error(w, `The service is currently unavailable`, http.StatusServiceUnavailable)
			return
		}
		body, err := ioutil.ReadAll(res.Body)
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
		spew.Dump(body, info)

		// todo: find out why we get 401
		// if user exists - log them in
		// if user doesn't exist - create them
	}
}

var e = url.QueryEscape
var scopes = e("https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")

func authURL(clientID, redirect string) string {
	return fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?response_type=code&scope=%s&client_id=%s&redirect_uri=%s&prompt=consent", scopes, e(clientID), e(redirect))
}

func tokenURL(code, clientID, clientSecret, redirect string) string {
	a := fmt.Sprintf("https://accounts.google.com/o/oauth2/token?code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code", e(code), e(clientID), e(clientSecret), e(redirect))
	println(a)
	return a
}
