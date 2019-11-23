package memory

import "github.com/3auris/auction-server/internal/user"

type users map[string]user.User
type sessions map[user.Token]user.Session

type UserMemoryStorage struct {
	users    users
	sessions sessions
}

func NewUserMemoryStorage() *UserMemoryStorage {
	return &UserMemoryStorage{
		users: users{
			"admin": user.User{
				Avatar:  "",
				Email:   "admin@admin.com",
				Name:    "AdminName",
				Surname: "AdminSurname",
				Address: user.Address{
					Address1: "Address line 1",
					Address2: "Address line 2",
				},
			},
		},
		sessions: sessions{},
	}
}

func (u UserMemoryStorage) GetByEmail(email string) *user.User {
	usr := u.users[email]
	return &usr
}

func (u UserMemoryStorage) Get(token user.Token) *user.Session {
	session := u.sessions[token]
	return &session
}

func (u *UserMemoryStorage) Add(session user.Session) {
	u.sessions[session.GetToken()] = session
}
