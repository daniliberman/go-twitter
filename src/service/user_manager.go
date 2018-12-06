package service


import (
	//"fmt"
	"github.com/daniliberman/twitter/src/domain"
)

// Initialization of users slice
var users []domain.User

func AddUser(user *domain.User) error{
	users = append(users, *user)
	

	return nil
//	return fmt.Errorf("adding user faild")
}