package app

import (
	"github.com/3auris/auction-server/store"
)

type App struct {
	store *store.Storage
}

func NewApp() App {
	return App{
		store: store.NewMemoryStorage(),
	}
}
