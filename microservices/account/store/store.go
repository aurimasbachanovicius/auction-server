package store

import (
	"github.com/3auris/auction-server/internal/user"
	"github.com/3auris/auction-server/store/memory"
)

// Storage main structure of storage of application
type Storage struct {
	User         user.Storage
	UserSession  user.SessionStorage
	UserPassword user.PasswordStorage
}

// NewMemoryStorage creates new storage with all dependencies
func NewMemoryStorage() *Storage {
	userStorage := memory.NewUserMemoryStorage()

	return &Storage{
		User:         userStorage,
		UserSession:  userStorage,
		UserPassword: memory.NewUserPasswordStorage(),
	}
}
