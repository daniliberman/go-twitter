package service_test

import (
	"testing"
	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"

)

func TestRegisterNewUser(t *testing.T) {
	//newUser parameters: name, mail, nick, pass
	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	
	service.AddUser(user)
}

func TestNewUserIsSaved(t *testing.T) {

	//initialization: newUser parameters: name, mail, nick, pass
	service.InitializeServiceUser()
	name := "dani"
	mail := "dani@mail.com"
	nick := "dliberman"
	pass := "password1"
	user := domain.NewUser(name, mail, nick, pass)

	//Operation
	service.AddUser(user)

	//Validation
	user = service.GetUsers()[0]
	if user.Name != name || user.Mail != mail || 
		user.Nick != nick || user.Pass != pass {
			t.Errorf("Expected user is name: %s, mail: %s, nick: %s, pass: %s \nbut is name: %s, mail: %s, nick: %s, pass: %s", 
				name, mail, nick, pass, user.Name, user.Mail, user.Nick, user.Pass)
		return	
	}
}
