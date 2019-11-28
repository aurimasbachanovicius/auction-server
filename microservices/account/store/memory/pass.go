package memory

import "github.com/3auris/auction-server/internal/user"

type passwords map[string]user.HashedPassword

type UserPasswordStorage struct {
	passwords passwords
}

func NewUserPasswordStorage() *UserPasswordStorage {
	return &UserPasswordStorage{
		passwords: passwords{"admin@admin.com": user.NewHashedPassword("admin")},
	}
}

func (ups UserPasswordStorage) GetByEmail(email string) user.HashedPassword {
	return ups.passwords[email]
}

func (ups *UserPasswordStorage) AddOrChangeToEmail(email string, password string) {
	ups.passwords[email] = user.NewHashedPassword(password)
}
