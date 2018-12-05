package domain

type User struct {
	name string
	mail string
	nick string
	pass string
}

func NewUser(name string, mail string, nick string, pass string) *User{
	var user User
	user.name = name
	user.mail = mail
	user.nick = nick
	user.pass = pass

	return &user
}