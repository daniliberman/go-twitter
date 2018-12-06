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