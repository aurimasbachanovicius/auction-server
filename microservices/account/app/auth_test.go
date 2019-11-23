package app

import (
	"fmt"
	"testing"

	"github.com/3auris/auction-server/internal/user"
	"github.com/3auris/auction-server/store"
)

var sessionMock = &user.Session{}
var addedSession = false

type userStorageMock struct {
	user.Storage
	user.SessionStorage
}

func (s userStorageMock) Add(session user.Session) { addedSession = true }

func (s userStorageMock) GetByEmail(email string) *user.User {
	if email != "admin@admin.com" {
		return nil
	}
	return &user.User{}
}

type userPasswordStorageMock struct{ user.PasswordStorage }

func (u userPasswordStorageMock) GetByEmail(email string) user.HashedPassword {
	return user.NewHashedPassword("admin")
}

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
			addedSession = false
			userStorageMock := userStorageMock{}

			app := App{
				store: &store.Storage{
					User:         userStorageMock,
					UserSession:  userStorageMock,
					UserPassword: userPasswordStorageMock{},
				},
			}

			err, session := app.Auth(tt.email, tt.pass)

			if tt.want1 != "" {
				if err == nil || err.Error() != tt.want1 {
					t.Errorf("Auth() want error %v, got %v", tt.want1, err)
				}
			}

			if err != nil && session != tt.want2 {
				t.Errorf("Auth() want session %v, got %v", tt.want2, session)
			}

			if tt.addedSession != addedSession {
				t.Errorf("Auth() want add to session %t, got %t", tt.addedSession, addedSession)
			}

		})
	}
}
