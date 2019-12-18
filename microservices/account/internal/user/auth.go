package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Authenticator interface which provides methods for authenticating and registering the user
type Authenticator interface {
	Register(email string, pass string) (error, User)
	Login(email string, pass string) (error, User)
}

// HashedPassword bytes of hashed password
type HashedPassword []byte

// NewHashedPassword bcrypts the given password ant gives hashed password
func NewHashedPassword(pass string) HashedPassword {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)

	return hashed
}

// IsMatch checks is the provided password matches with the hashed password
func (hashedPassword HashedPassword) IsMatch(pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass)); err != nil {
		return false
	}

	return true
}

// PasswordStorage interface for which provides methods for user's password storage manipulation
type PasswordStorage interface {
	GetByEmail(email string) (HashedPassword, error)
	AddOrChangeToEmail(email, password string) error
}
