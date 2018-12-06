package domain

type User struct {
	Name string
	Mail string
	Nick string
	Pass string
}

func NewUser(name string, mail string, nick string, pass string) *User{
	var user User
	user.Name = name
	user.Mail = mail
	user.Nick = nick
	user.Pass = pass

	return &user
}