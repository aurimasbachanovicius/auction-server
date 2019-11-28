package app

import (
	"fmt"
	"testing"

	"github.com/3auris/auction-server/internal/user"
	"github.com/3auris/auction-server/store"
)

var (
	sessionMock = &user.Session{}
	adminPass   = user.NewHashedPassword("admin")
)

type userStorageMock struct {
	user.Storage
	user.SessionStorage

	onCall func()
}

func (s userStorageMock) Add(session user.Session) { s.onCall() }
func (s userStorageMock) Create(user user.User)    { s.onCall() }
func (s userStorageMock) GetByEmail(email string) *user.User {
	if email != "admin@admin.com" {
		return nil
	}
	return &user.User{}
}

type userPasswordStorageMock struct {
	user.PasswordStorage
	onCall func()
}

func (u userPasswordStorageMock) GetByEmail(email string) user.HashedPassword      { return adminPass }
func (u userPasswordStorageMock) AddOrChangeToEmail(email string, password string) { u.onCall() }

func TestApp_Auth(t *testing.T) {
	tests := []struct {
		email string
		pass  string

		want1 string
		want2 *user.Session

		addedSession bool
	}{
		{email: "test@test.com", pass: "test", want1: "could not find user by email", want2: nil},
		{email: "admin@admin.com", pass: "test", want1: "could not match password", want2: nil},
		{email: "admin@admin.com", pass: "admin", want1: "", want2: sessionMock, addedSession: true},
		{email: "admin@user.com", pass: "admin", want1: "could not find user by email", want2: nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("given email:%s and pass:%s", tt.email, tt.pass), func(t *testing.T) {
			usrStg := userStorageMock{}

			var addSession bool
			usrSessionStg := userStorageMock{onCall: func() {
				addSession = true
			}}

			passStg := userPasswordStorageMock{}

			app := App{
				store: &store.Storage{
					User:         &usrStg,
					UserSession:  &usrSessionStg,
					UserPassword: &passStg,
				},
			}

			session, err := app.Auth(tt.email, tt.pass)

			if tt.want1 != "" {
				if err == nil || err.Error() != tt.want1 {
					t.Errorf("Auth() want error %v, got %v", tt.want1, err)
				}
			}

			if err != nil && session != tt.want2 {
				t.Errorf("Auth() want session %v, got %v", tt.want2, session)
			}

			if tt.addedSession != addSession {
				t.Errorf("Auth() want add to session %t, got %t", tt.addedSession, addSession)
			}

		})
	}
}

func TestApp_NewUser(t *testing.T) {
	tests := []struct {
		email string
		pass  string

		err bool

		addedUser         bool
		addedUserPassword bool
	}{
		{email: "test@test.com", pass: "test", addedUser: true, addedUserPassword: true},
		{email: "admin@admin.com", pass: "test", err: true},
		{email: "admin@admin.com", pass: "admin", err: true},
		{email: "admin@user.com", pass: "admin", addedUser: true, addedUserPassword: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("given email:%s and pass:%s", tt.email, tt.pass), func(t *testing.T) {
			var addUser, addPass bool

			passStg := userPasswordStorageMock{onCall: func() {
				addPass = true
			}}
			userStg := userStorageMock{onCall: func() {
				addUser = true
			}}

			app := App{
				store: &store.Storage{
					User:         &userStg,
					UserSession:  &userStg,
					UserPassword: &passStg,
				},
			}

			err := app.NewUser(tt.email, tt.pass)
			if err != nil && !tt.err {
				t.Errorf("NewUser() do not want error, got: %v", err)
			}

			if tt.addedUser != addUser {
				t.Errorf("NewUser() should add user to store, want:%t , got: %t", tt.addedUser, addUser)
			}

			if tt.addedUserPassword != addPass {
				t.Errorf("NewUser() should add user's password to store, want:%t , got: %t", tt.addedUser, addPass, )
			}
		})
	}
}
