package request

import (
	"net/http"
	"github.com/Briareos/rocket"
	"github.com/gorilla/sessions"
)

const (
	Token = iota
	Session
)

func GetToken(r*http.Request) *rocket.Token {
	return r.Context().Value(Token).(*rocket.Token)
}

func GetSession(r*http.Request) *sessions.Session {
	return r.Context().Value(Session).(*sessions.Session)
}
