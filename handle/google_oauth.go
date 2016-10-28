package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"io"
	"bytes"
	"github.com/Briareos/rocket"
	"github.com/Briareos/rocket/request"
	"database/sql"
)

func GoogleOAuth(clientID, redirect string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//@todo secure this
		state := ""
		http.Redirect(w, r, authURL(clientID, redirect, state), 302)
	}
}

type tokenInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type userInfo struct {
	LastName      string `json:"family_name"`
	FullName      string `json:"name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	Link          string `json:"link"`
	FirstName     string `json:"given_name"`
	ID            string `json:"id"`
	HostedDomain  string `json:"hd"`
	VerifiedEmail bool `json:"verified_email"`
}

func GoogleOAuthCallback(clientID, clientSecret, redirectURI, returnURI string, userSvc rocket.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := getUserInfo(r.FormValue("code"), clientID, clientSecret, redirectURI)
		if err != nil {
			http.Error(w, `The service is currently unavailable: ` + err.Error(), http.StatusServiceUnavailable)
			return
		}
		user, err := userSvc.GetByGoogleID(info.ID)
		if err == sql.ErrNoRows {
			user = new(rocket.User)
			user.Email = info.Email
			user.FirstName = info.FirstName
			user.LastName = info.LastName
			user.GoogleID = info.ID
			if err = userSvc.Add(user); err != nil {
				http.Error(w, `The service is currently unavailable: ` + err.Error(), http.StatusServiceUnavailable)
				return
			}
		} else if err != nil {
			http.Error(w, `The service is currently unavailable: ` + err.Error(), http.StatusServiceUnavailable)
			return
		}
		token := request.GetToken(r)
		token.SetUser(user)
		http.Redirect(w, r, returnURI, 302)
	}
}

var e = url.QueryEscape
var scopes = e("https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")

func authURL(clientID, redirect, state string) string {
	return fmt.Sprintf("https://accounts.google.com/o/oauth2/auth?response_type=code&scope=%s&client_id=%s&redirect_uri=%s&prompt=consent&access_type=offline&state=%s", scopes, e(clientID), e(redirect), e(state))
}

func tokenParams(code, clientID, clientSecret, redirect string) io.Reader {
	return bytes.NewBufferString(fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code", e(code), e(clientID), e(clientSecret), e(redirect)))
}

func getUserInfo(code, clientID, clientSecret, redirect string) (*userInfo, error) {
	res, err := http.Post("https://accounts.google.com/o/oauth2/token", "application/x-www-form-urlencoded", tokenParams(code, clientID, clientSecret, redirect))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	tokInfo := new(tokenInfo)
	if err := json.Unmarshal(body, tokInfo); err != nil {
		return nil, err
	}
	res, err = http.Get("https://www.googleapis.com/userinfo/v2/me?access_token=" + e(tokInfo.AccessToken))
	if err != nil {
		return nil, err
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	uInfo := new(userInfo)
	if err = json.Unmarshal(body, uInfo); err != nil {
		return nil, err
	}
	return uInfo, nil
}
