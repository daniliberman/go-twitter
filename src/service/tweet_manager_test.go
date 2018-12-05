package service_test

import (
	"testing"

	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"

)

func TestPublishedTweetIsSaved(t *testing.T) {
	
	/*
	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
	*/

	//initialization
	var tweet *domain.Tweet
	user := "daniliberman"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
			t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", 
				user, text, publishedTweet.User, publishedTweet.Text)
		return	
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}


}

//func TestRegisterNewUser(t *testing.T)