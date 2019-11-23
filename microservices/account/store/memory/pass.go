package memory

import "github.com/3auris/auction-server/internal/user"

type UserPasswordStorage struct{}

func (UserPasswordStorage) GetByEmail(email string) user.HashedPassword {
	return user.NewHashedPassword("admin")
}

func (UserPasswordStorage) AddOrChangeToEmail(email string, password string) user.HashedPassword {
	return user.NewHashedPassword(password)
}
