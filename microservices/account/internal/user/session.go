package user

import (
	"math/rand"
	"time"
)

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
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 64)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
