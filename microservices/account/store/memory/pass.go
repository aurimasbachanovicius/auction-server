package memory

import (
	"errors"

	"github.com/3auris/auction-server/internal/user"
)

type passwords map[string]user.HashedPassword

// UserPasswordStorage storage for keeping the users' passwords
type UserPasswordStorage struct {
	passwords passwords
}

// NewUserPasswordStorage creates new user password storage
func NewUserPasswordStorage() *UserPasswordStorage {
	return &UserPasswordStorage{
		passwords: passwords{"admin@admin.com": user.NewHashedPassword("admin")},
	}
}

// GetByEmail gives the users hashed password from storage
func (ups UserPasswordStorage) GetByEmail(email string) (user.HashedPassword, error) {
	pass, ok := ups.passwords[email]
	if !ok {
		return nil, errors.New("could not find user's password")
	}

	return pass, nil
}

// AddOrChangeToEmail changes or add new password for given user
func (ups *UserPasswordStorage) AddOrChangeToEmail(email, password string) error {
	ups.passwords[email] = user.NewHashedPassword(password)
	return nil
}
