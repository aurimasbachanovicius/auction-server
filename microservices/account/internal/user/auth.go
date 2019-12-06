package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Authenticator interface {
	Register(email string, pass string) (error, User)
	Login(email string, pass string) (error, User)
}

type HashedPassword []byte

func NewHashedPassword(pass string) HashedPassword {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)

	return hashed
}

func (hashedPassword HashedPassword) IsMatch(pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass)); err != nil {
		return false
	}

	return true
}

type PasswordStorage interface {
	GetByEmail(email string) (HashedPassword, error)
	AddOrChangeToEmail(email, password string) error
}
