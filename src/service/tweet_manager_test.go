package service_test

import (
	"testing"
	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"

)

func TestPublishedTweetIsSaved(t *testing.T) {

	//initialization
	service.InitializeServiceTweet()
	service.InitializeServiceUser()
	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	service.AddUser(user)

	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweets()[0]
	
	if !service.CompareUsers(publishedTweet.User, user) ||
		publishedTweet.Text != text {
			t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", 
				user.Nick, text, publishedTweet.User.Nick, publishedTweet.Text)
		return	
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	//Initialization
	service.InitializeServiceTweet()
	service.InitializeServiceUser()
	var tweet *domain.Tweet
	
	var user *domain.User
	text := "this is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error if user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	//Initialization
	service.InitializeServiceTweet()
	service.InitializeServiceUser()
	var tweet *domain.Tweet

	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	service.AddUser(user)

	var text string

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error if text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	//Initialization
	service.InitializeServiceTweet()
	service.InitializeServiceUser()
	var tweet *domain.Tweet
	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	service.AddUser(user)
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error if text exceeds 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	service.InitializeServiceTweet()
	service.InitializeServiceUser()

	var tweet, seccondTweet *domain.Tweet
	user1 := domain.NewUser("user1", "user2@mail.com", "user1Nick", "password1")
	service.AddUser(user1)
	user2 := domain.NewUser("user2", "user2@mail.com", "user2Nick", "password2")
	service.AddUser(user2)

	tweet = domain.NewTweet(user1, "this is the first tweet")
	seccondTweet = domain.NewTweet(user2, "this is the seccond tweet")

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(seccondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("expected size is 2 but was %d", len(publishedTweets))
		return
	}
	
	
	firstPublishedTweet := publishedTweets[0]
	seccondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, 0, tweet.User, tweet.Text) {
		return
	}
	if !isValidTweet(t, seccondPublishedTweet, 1, seccondTweet.User, seccondTweet.Text) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user *domain.User, text string) bool {
	if !service.CompareUsers(tweet.User, user) {
		t.Errorf("no match in users: tweet.User: %s, user: %s", tweet.User.Nick, user.Nick)
		return false
	}
	if tweet.Text != text {
		t.Error("no match in texts")
		return false
	}
	if tweet.Id != id {
		t.Error("no match in ids")
		return false
	}
	return true
}

func TestCanRetrieveTweetById(t * testing.T) {
	// Initialization
	service.InitializeServiceTweet()
	service.InitializeServiceUser()

	var tweet *domain.Tweet
	var id int

	user := domain.NewUser("user", "user@mail.com", "userNick", "password")
	service.AddUser(user)

	text := "this is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validaiton
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCantPublishIfUserDoesNotExist(t * testing.T) {
	//initialization
	service.InitializeServiceTweet()
	service.InitializeServiceUser()
	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")

	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user does not exist" {
		t.Error("Expected error if user does not exist")
	}
	return	
}