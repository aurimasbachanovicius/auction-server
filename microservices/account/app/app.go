package app

import (
	"github.com/3auris/auction-server/store"
)

// App main application where the business logic is done
type App struct {
	store *store.Storage
}

// NewApp creates new application with dependencies
func NewApp() App {
	return App{
		store: store.NewMemoryStorage(),
	}
}
