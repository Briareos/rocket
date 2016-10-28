package rocket

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type Token struct {
	session *sessions.Session
	user    *User
	r       *http.Request
	w       http.ResponseWriter
}

const userID = "userID"

func NewToken(s *sessions.Session, userSvc UserService, r *http.Request, w http.ResponseWriter) *Token {
	t := &Token{
		session: s,
		r:       r,
		w:       w,
	}
	if id, ok := s.Values[userID].(int); ok {
		if user, err := userSvc.Get(id); err == nil {
			t.user = user
		}
	}
	return t
}

func (t *Token) SetUser(user *User) {
	if user == nil {
		delete(t.session.Values, userID)
	} else {
		t.session.Values[userID] = user.ID
	}
	t.session.Save(t.r, t.w)
	t.user = user
}

func (t *Token) User() *User {
	return t.user
}
