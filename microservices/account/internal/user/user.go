package user

type UUID string

type User struct {
	uuid  UUID
	email string
}

type SessionStorage interface {
	Get(UUID) User
	Set(User)
}
