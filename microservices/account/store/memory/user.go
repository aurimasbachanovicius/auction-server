package memory

import (
	"errors"

	"github.com/3auris/auction-server/internal/user"
)

type users map[string]user.User
type sessions map[user.Token]user.Session

type UserMemoryStorage struct {
	users    users
	sessions sessions
}

func NewUserMemoryStorage() *UserMemoryStorage {
	return &UserMemoryStorage{
		users: users{
			//"admin": user.User{
			//	Email:   "admin@admin.com",
			//	Name:    "AdminName",
			//	Surname: "AdminSurname",
			//	Address: user.Address{
			//		Address1: "Address line 1",
			//		Address2: "Address line 2",
			//	},
			//},
		},
		sessions: sessions{},
	}
}

func (u UserMemoryStorage) GetByEmail(email string) (*user.User, error) {
	usr, ok := u.users[email]
	if !ok {
		return nil, errors.New("could not find user in memory storage")
	}

	return &usr, nil
}

func (u UserMemoryStorage) Exists(email string) (bool, error) {
	_, ok := u.users[email]

	return ok, nil
}

func (u *UserMemoryStorage) Create(user user.User) error {
	u.users[user.Email] = user
	return nil
}

func (u UserMemoryStorage) Get(token user.Token) (*user.Session, error) {
	session, ok := u.sessions[token]
	if !ok {
		return nil, errors.New("could not find session in memory storage")
	}

	return &session, nil
}

func (u *UserMemoryStorage) Add(session user.Session) error {
	u.sessions[session.GetToken()] = session
	return nil
}
