package app

import (
	"github.com/pkg/errors"

	"github.com/3auris/auction-server/internal/user"
)

// Auth user and create new user session
func (app App) Auth(email string, pass string) (error, *user.Session) {
	storedUser := app.store.User.GetByEmail(email)
	if storedUser == nil {
		return errors.New("could not find user by email"), nil
	}

	password := app.store.UserPassword.GetByEmail(email)
	if !password.IsMatch(pass) {
		return errors.New("could not match password"), nil
	}

	session := user.NewSession();
	app.store.UserSession.Add(session)


	return nil, &session
}
