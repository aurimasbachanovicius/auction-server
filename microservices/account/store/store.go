package store

import (
	"github.com/3auris/auction-server/internal/user"
	"github.com/3auris/auction-server/store/memory"
)

type Storage struct {
	User         user.Storage
	UserSession  user.SessionStorage
	UserPassword user.PasswordStorage
}

func NewMemoryStorage() *Storage {
	userStorage := memory.NewUserMemoryStorage()

	return &Storage{
		User:         userStorage,
		UserSession:  userStorage,
		UserPassword: memory.UserPasswordStorage{},
	}
}
