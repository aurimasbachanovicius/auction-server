package user

import (
	"errors"
	"regexp"
)

// Address user's address structure
type Address struct {
	Address1 string
	Address2 string
}

// User user's data structure
type User struct {
	Avatar  string
	Email   string
	Name    string
	Surname string
	Address Address
}

// Storage interface which provides user's persistent data manipulation
type Storage interface {
	GetByEmail(email string) (*User, error)
	Exists(email string) (bool, error)
	Create(user User) error
}

// Validate validates is the user's data is correct and not corrupted
func (u User) Validate() error {
	re := regexp.MustCompile(`^(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)])$`)

	if !re.MatchString(u.Email) {
		return errors.New("invalid email")
	}

	return nil
}
