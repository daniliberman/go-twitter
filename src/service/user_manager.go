package service


import (
	"fmt"
	"github.com/daniliberman/twitter/src/domain"
	"reflect"
)


func (tweetManager *TweetManager)AddUser(user *domain.User) error{
	tweetManager.Users = append(tweetManager.Users, user)
	
	return nil
//	return fmt.Errorf("adding user faild")
}

func (tweetManager *TweetManager)GetUsers() []*domain.User {
	return tweetManager.Users
}

func CompareUsers(user1 *domain.User, user2 *domain.User) bool{
	return reflect.DeepEqual(user1, user2)
}

// If user does not exist returns nil
func (tweetManager *TweetManager)GetUserWithNick(nick string) *domain.User {
	for _, user := range tweetManager.Users {
		if user.Nick == nick {
			return user
		}
	}
	return nil
}

func (tweetManager *TweetManager)Login(nick string, pass string) error{
	user := tweetManager.GetUserWithNick(nick)
	if user == nil || user.Pass != pass {
		return fmt.Errorf("Wrong nick or password")
	}
	tweetManager.LoggedInUsers = append(tweetManager.LoggedInUsers, user)
	return nil
}

func (tweetManager *TweetManager)GetLoggedInUsers() []*domain.User {
	return tweetManager.LoggedInUsers
}

func (tweetManager *TweetManager)IsUserLoggedIn(user *domain.User) bool {
	for _,currentUser := range tweetManager.GetLoggedInUsers() {
		if(CompareUsers(user, currentUser)) {
			return true
		}
	}
	return false
}
