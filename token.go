package rocket

import "net/http"

type Token struct {
	req  *http.Request
	user *User
}

func NewToken(r *http.Request) *Token {
	return new(Token)
}

func (t *Token) SetUser(user *User) {
	t.user = user
}

func (t *Token) User() *User {
	return t.user
}
