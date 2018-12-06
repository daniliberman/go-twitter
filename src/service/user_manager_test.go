package service_test

import (
	"testing"
	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"

)

func TestNewUserIsSaved(t *testing.T) {

	//initialization: newUser parameters: name, mail, nick, pass
	tweetManager := service.NewTweetManager()
	name := "dani"
	mail := "dani@mail.com"
	nick := "dliberman"
	pass := "password1"
	user := domain.NewUser(name, mail, nick, pass)

	//Operation
	tweetManager.AddUser(user)

	//Validation
	user = tweetManager.GetUsers()[0]
	if user.Name != name || user.Mail != mail || 
		user.Nick != nick || user.Pass != pass {
			t.Errorf("Expected user is name: %s, mail: %s, nick: %s, pass: %s \nbut is name: %s, mail: %s, nick: %s, pass: %s", 
				name, mail, nick, pass, user.Name, user.Mail, user.Nick, user.Pass)
		return	
	}
}

func TestLoginUserOK(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()

	name := "dani"
	mail := "dani@mail.com"
	nick := "dliberman"
	pass := "password1"
	user := domain.NewUser(name, mail, nick, pass)
	tweetManager.AddUser(user)
	
	err := tweetManager.Login(nick, pass)
	if err != nil {
		t.Errorf("TestLoginUserOK faild with error: %s", err)
	}
}

func TestLoginUserFaildWithWrongNick(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()

	name := "dani"
	mail := "dani@mail.com"
	nick := "dliberman"
	pass := "password1"
	user := domain.NewUser(name, mail, nick, pass)
	tweetManager.AddUser(user)
	
	err := tweetManager.Login("otro nick", pass)
	//Validation
	
	if err != nil && err.Error() != "Wrong nick or password" {
		t.Error("Expected error if wrong nick or password")
	}
}

func TestLoginUserFaildWithWrongPass(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()

	name := "dani"
	mail := "dani@mail.com"
	nick := "dliberman"
	pass := "password1"
	user := domain.NewUser(name, mail, nick, pass)
	tweetManager.AddUser(user)
	
	err := tweetManager.Login(nick, "otraPass")
	//Validation
	
	if err != nil && err.Error() != "Wrong nick or password" {
		t.Error("Expected error if wrong nick or password")
	}
}
