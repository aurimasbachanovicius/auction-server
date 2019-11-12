package user

import "crypto/rand"

type User struct {
	email string
}

type Session struct {
	token  string
	expire string
}

func NewUser(email string) User {
	return User{email: email}
}

func NewSession() Session {
	return Session{token: generateToken(), expire: "2020-01-20T14:48"}
}

func generateToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return string(b)
}
