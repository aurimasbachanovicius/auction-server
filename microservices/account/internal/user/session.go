package user

import (
	"math/rand"
	"time"
)

// Token token for authenticating and identifying user session
type Token string

// Session user's session
type Session struct {
	token  Token
	expire string
}

// GetToken gives the session's token
func (s Session) GetToken() Token {
	return s.token
}

// GetExpire gives the expiration string of token
func (s Session) GetExpire() string {
	return s.expire
}

// NewSession creates new session with generated token
func NewSession() Session {
	return Session{token: Token(generateToken()), expire: "2020-01-20T14:48"}
}

// SessionStorage interface which provides methods for setting the user's session
type SessionStorage interface {
	Get(token Token) (*Session, error)
	Add(session Session) error
}

func generateToken() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 64)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
