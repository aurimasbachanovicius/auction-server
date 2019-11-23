package user

type Address struct {
	Address1 string
	Address2 string
}

type User struct {
	Avatar  string
	Email   string
	Name    string
	Surname string
	Address Address
}

type Storage interface {
	GetByEmail(email string) *User
}
