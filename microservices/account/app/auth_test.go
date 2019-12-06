package app

import (
	"errors"
	"testing"

	"github.com/3auris/auction-server/internal/user"
	"github.com/3auris/auction-server/store"
)

type userSessionStorageMock struct {
	addError bool
}

func (s userSessionStorageMock) Get(token user.Token) (*user.Session, error) {
	panic("implement me")
}

func (s userSessionStorageMock) Add(_ user.Session) error {
	if s.addError {
		return errors.New("")
	}

	return nil
}

type userStorageMock struct {
	user.Storage

	createReturnErr bool
}

func (s userStorageMock) Create(user user.User) error {
	if s.createReturnErr {
		return errors.New("")
	}

	return nil
}

func (s userStorageMock) Get(token user.Token) (*user.Session, error) {
	return &user.Session{}, nil
}

func (s userStorageMock) Exists(email string) (bool, error) {
	if email == "existinguser@mail.com" || email == "secondexistinguser@mail.com" {
		return true, nil
	}

	if email == "notexistinguser@mail.com" || email == "not_existing_invalid_mail" {
		return false, nil
	}

	return false, errors.New("")
}

func (s userStorageMock) GetByEmail(email string) (*user.User, error) {
	if email == "admin@admin.com" || email == "admin@onlystorage.com" {
		return &user.User{}, nil
	}
	return nil, errors.New("")
}

type userPasswordStorageMock struct {
	addOrChangeToEmailErr bool
}

func (u userPasswordStorageMock) GetByEmail(email string) (user.HashedPassword, error) {
	if email == "secondexistinguser@mail.com" {
		return user.NewHashedPassword("admin"), nil
	}
	return nil, errors.New("")
}

func (u userPasswordStorageMock) AddOrChangeToEmail(_, _ string) error {
	if u.addOrChangeToEmailErr {
		return errors.New("")
	}

	return nil
}

func TestApp_Auth(t *testing.T) {
	tests := map[string]struct {
		email string
		pass  string

		wantSession bool

		userSessionAddError bool

		wantErr bool
	}{
		"search in store do not work": {
			email:   "admin@admin.com",
			pass:    "admin",
			wantErr: true,
		},
		"could not find user": {
			email:   "notexistinguser@mail.com",
			pass:    "admin",
			wantErr: true,
		},
		"could not find password": {
			email:   "existinguser@mail.com",
			pass:    "admin",
			wantErr: true,
		},
		"password do not match": {
			email:   "secondexistinguser@mail.com",
			pass:    "adminwrongpassword",
			wantErr: true,
		},
		"could not add user session": {
			email:               "secondexistinguser@mail.com",
			pass:                "admin",
			userSessionAddError: true,
			wantErr:             true,
		},
		"everything is fine": {
			email:       "secondexistinguser@mail.com",
			pass:        "admin",
			wantSession: true,
			wantErr:     false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			app := App{
				store: &store.Storage{
					User:         userStorageMock{},
					UserSession:  userSessionStorageMock{addError: tt.userSessionAddError},
					UserPassword: userPasswordStorageMock{},
				},
			}

			session, err := app.Auth(tt.email, tt.pass)

			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (session != nil) != tt.wantSession {
				t.Errorf("Auth() session = %v, wantSession %v", session, tt.wantSession)
			}
		})
	}
}

func TestApp_NewUser(t *testing.T) {
	tests := map[string]struct {
		email string
		pass  string

		createErr      bool
		addOrChangeErr bool

		wantErr bool
	}{
		"could not fetch from store": {
			email:   "not_fetchable_email",
			wantErr: true,
		},
		"user already exist": {
			email:   "existinguser@mail.com",
			wantErr: true,
		},
		"invalid user should return error": {
			email:   "not_existing_invalid_mail",
			wantErr: true,
		},
		"could not create user": {
			email:     "notexistinguser@mail.com",
			wantErr:   true,
			createErr: true,
		},
		"could not change or add email": {
			email:          "notexistinguser@mail.com",
			wantErr:        true,
			addOrChangeErr: true,
		},
		"everything is fine": {
			email: "notexistinguser@mail.com",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			app := App{
				store: &store.Storage{
					User:         userStorageMock{createReturnErr: tt.createErr},
					UserSession:  userSessionStorageMock{},
					UserPassword: userPasswordStorageMock{addOrChangeToEmailErr: tt.addOrChangeErr},
				},
			}

			err := app.NewUser(tt.email, tt.pass)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
