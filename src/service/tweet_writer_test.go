package service_test

import (
	"testing"
	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"
	//"github.com/golang/mock/gomock"

)

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager()
	tweetManager.TweetWriter = tweetWriter

	var tweet *domain.TextTweet
	user := domain.NewUser("user", "user@mail.com", "userNick", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)
	text := "this is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	
	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Errorf("Expected tweet on file, but there was no one")
	}
	if savedTweet.GetId() != id {
		t.Errorf("expected same id but got: savedTweet's id: %d, id: %d", savedTweet.GetId(), id)
	}
}

func BenchmarkPublishTweetWithFileTweetWriter(b *testing.B){
	// Initialization
	fileTweetWriter := service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager()
	tweetManager.TweetWriter = fileTweetWriter

	var tweet *domain.TextTweet
	user := domain.NewUser("user", "user@mail.com", "userNick", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)
	text := "this is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet)
	}
}

func BenchmarkPublishTweetWithMemoryTweetWriter(b *testing.B){
	// Initialization
	fileTweetWriter := service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager()
	tweetManager.TweetWriter = fileTweetWriter

	var tweet *domain.TextTweet
	user := domain.NewUser("user", "user@mail.com", "userNick", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)
	text := "this is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet)
	}
}