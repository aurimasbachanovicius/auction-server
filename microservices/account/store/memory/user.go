package memory

import (
	"errors"

	"github.com/3auris/auction-server/internal/user"
)

type users map[string]user.User
type sessions map[user.Token]user.Session

// UserMemoryStorage storage for storing users data
type UserMemoryStorage struct {
	users    users
	sessions sessions
}

// NewUserMemoryStorage creates new user memory storage
func NewUserMemoryStorage() *UserMemoryStorage {
	return &UserMemoryStorage{
		users:    users{},
		sessions: sessions{},
	}
}

// GetByEmail gives users data by provided email
func (u UserMemoryStorage) GetByEmail(email string) (*user.User, error) {
	usr, ok := u.users[email]
	if !ok {
		return nil, errors.New("could not find user in memory storage")
	}

	return &usr, nil
}

// Exists checks if given user email exists in storage
func (u UserMemoryStorage) Exists(email string) (bool, error) {
	_, ok := u.users[email]

	return ok, nil
}

// Create creates adds new users to the storage
func (u *UserMemoryStorage) Create(user user.User) error {
	u.users[user.Email] = user
	return nil
}

// Get gets user session from the memory storage by given token which can
// be used for further validations
func (u UserMemoryStorage) Get(token user.Token) (*user.Session, error) {
	session, ok := u.sessions[token]
	if !ok {
		return nil, errors.New("could not find session in memory storage")
	}

	return &session, nil
}

// Add adds new user session to the storage
func (u *UserMemoryStorage) Add(session user.Session) error {
	u.sessions[session.GetToken()] = session
	return nil
}
