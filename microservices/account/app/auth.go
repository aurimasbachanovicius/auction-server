package app

import (
	"github.com/pkg/errors"

	"github.com/3auris/auction-server/internal/user"
)

// Auth user and create new user session
func (app App) Auth(email string, pass string) (*user.Session, error) {
	exist, err := app.store.User.Exists(email)
	if err != nil {
		return nil, errors.Wrap(err, "could not check if user exists")
	}

	if !exist {
		return nil, errors.New("could not find user by given email")
	}

	password, err := app.store.UserPassword.GetByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "could not get user's password")
	}

	if !password.IsMatch(pass) {
		return nil, errors.New("could not match password")
	}

	session := user.NewSession()
	if err := app.store.UserSession.Add(session); err != nil {
		return nil, errors.Wrap(err, "could not add user session")
	}

	return &session, nil
}

// NewUser creates user and password
func (app App) NewUser(email string, password string) error {
	exist, err := app.store.User.Exists(email)
	if err != nil {
		return errors.Wrap(err, "could not fetch from store")
	}

	if exist {
		return errors.Errorf("user with %s email already exist", email)
	}

	u := user.User{Email: email}
	if err := u.Validate(); err != nil {
		return errors.Wrap(err, "could not validate user")
	}

	if err := app.store.User.Create(u); err != nil {
		return errors.Wrap(err, "could not create user")
	}

	if err := app.store.UserPassword.AddOrChangeToEmail(email, password); err != nil {
		return errors.Wrap(err, "could not add or change user's password")
	}

	return nil
}
