package rocket

import (
	"github.com/gorilla/sessions"
)

type Token struct {
	session *sessions.Session
	user    *User
}

func NewToken(s *sessions.Session) *Token {
	return &Token{
		session: s,
	}
}

func (t *Token) SetUser(user *User) {
	if user == nil {
		delete(t.session.Values, "userID")
	} else {
		t.session.Values["userID"] = user.ID
	}
	t.user = user
}

func (t *Token) User() *User {
	return t.user
}
