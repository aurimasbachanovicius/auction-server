package user

import "crypto/rand"

type Token string

type Session struct {
	token  Token
	expire string
}

func (s Session) GetToken() Token {
	return s.token
}

func (s Session) GetExpire() string {
	return s.expire
}

func NewSession() Session {
	return Session{token: Token(generateToken()), expire: "2020-01-20T14:48"}
}

type SessionStorage interface {
	Get(token Token) *Session
	Add(session Session)
}

func generateToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return string(b)
}
