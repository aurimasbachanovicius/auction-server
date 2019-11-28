package app

import (
	"github.com/pkg/errors"

	"github.com/3auris/auction-server/internal/user"
)

// Auth user and create new user session
func (app App) Auth(email string, pass string) (*user.Session, error) {
	storedUser := app.store.User.GetByEmail(email)
	if storedUser == nil {
		return nil, errors.New("could not find user by email")
	}

	password := app.store.UserPassword.GetByEmail(email)
	if !password.IsMatch(pass) {
		return nil, errors.New("could not match password")
	}

	session := user.NewSession();
	app.store.UserSession.Add(session)

	return &session, nil
}

// NewUser creates user and password
func (app App) NewUser(email string, password string) error {
	storedUser := app.store.User.GetByEmail(email)
	if storedUser != nil {
		return errors.New("user with %s email already exist")
	}

	u := user.User{Email: email}
	if err := u.Validate(); err != nil {
		return errors.Wrap(err, "could not validate user")
	}

	app.store.User.Create(u)
	app.store.UserPassword.AddOrChangeToEmail(email, password)

	return nil
}
