package account

type User struct {
	email string
}

func (u *User) setEmail(email string) {
	u.email = email
}

func NewUser(email string) User {
	user := User{}
	user.setEmail(email)
	
	return user
}
