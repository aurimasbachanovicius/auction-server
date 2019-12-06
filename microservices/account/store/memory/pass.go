package memory

import (
	"errors"

	"github.com/3auris/auction-server/internal/user"
)

type passwords map[string]user.HashedPassword

type UserPasswordStorage struct {
	passwords passwords
}

func NewUserPasswordStorage() *UserPasswordStorage {
	return &UserPasswordStorage{
		passwords: passwords{"admin@admin.com": user.NewHashedPassword("admin")},
	}
}

func (ups UserPasswordStorage) GetByEmail(email string) (user.HashedPassword, error) {
	pass, ok := ups.passwords[email]
	if !ok {
		return nil, errors.New("could not find user's password")
	}

	return pass, nil
}

func (ups *UserPasswordStorage) AddOrChangeToEmail(email, password string) error {
	ups.passwords[email] = user.NewHashedPassword(password)
	return nil
}
