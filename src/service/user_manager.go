package service


import (
	//"fmt"
	"github.com/daniliberman/twitter/src/domain"
	"reflect"
)

// Initialization of users slice
var users []*domain.User

func AddUser(user *domain.User) error{
	users = append(users, user)
	
	return nil
//	return fmt.Errorf("adding user faild")
}

func InitializeServiceUser() {
	users = make([]*domain.User, 0) 
}

func GetUsers() []*domain.User {
	return users
}

func CompareUsers(user1 *domain.User, user2 *domain.User) bool{
	return reflect.DeepEqual(user1, user2)
}

// If user does not exist returns nil
func GetUserWithNick(nick string) *domain.User {
	for _, user := range users {
		if user.Nick == nick {
			return user
		}
	}
	return nil
}

