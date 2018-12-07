package service_test

import (
	"testing"
	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"
	"reflect"

)

// TEXT TWEET
func TestPublishedTextTweetIsSaved(t *testing.T) {

	//initialization
	tweetManager := service.NewTweetManager()

	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)

	text := "This is my first tweet"
	tweet := domain.NewTextTweet(user, text)

	//Operation
	tweetManager.PublishTweet(tweet)

	//Validation
	publishedTweet := tweetManager.GetTweets()[0]
	
	if !service.CompareUsers(publishedTweet.GetUser(), user) ||
		publishedTweet.GetText() != text {
			t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", 
				user.Nick, text, publishedTweet.GetUser().Nick, publishedTweet.GetText())
		return	
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

// IMAGE TWEET
func TestPublishedImageTweetIsSaved(t *testing.T) {

	//initialization
	tweetManager := service.NewTweetManager()

	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)

	text := "This is my first image tweet"
	url := "https://www.google.com.ar/search?q=perro&rlz=1C5CHFA_enAR825AR826&source=lnms&tbm=isch&sa=X&ved=0ahUKEwjUndeX2I3fAhXJkpAKHepeDyAQ_AUIDigB&biw=1440&bih=718#imgrc=U0Xq_HR3t7p7dM:"
	tweet := domain.NewImageTweet(user, text, url)

	//Operation
	tweetManager.PublishTweet(tweet)

	//Validation
	publishedTweet := tweetManager.GetTweets()[0].(*domain.ImageTweet)
	
	if !service.CompareUsers(publishedTweet.GetUser(), user) ||
		publishedTweet.GetText() != text || publishedTweet.GetUrl() != url {
			t.Errorf("Expected tweet is %s: %s, url: %s\nbut is %s: %s, url: %s", 
				user.Nick, text, url, publishedTweet.GetUser().Nick, publishedTweet.GetText(), publishedTweet.GetUrl())
		return	
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

// QUOTE TWEET
func TestPublishedQuoteTweetIsSaved(t *testing.T) {

	//initialization
	tweetManager := service.NewTweetManager()

	user1 := domain.NewUser("user1", "user1@mail.com", "user1", "password")
	tweetManager.AddUser(user1)
	tweetManager.Login(user1.Nick, user1.Pass)

	text1 := "This is user1's tweet"
	tweet1 := domain.NewTextTweet(user1, text1)

	user2 := domain.NewUser("user2", "user2@mail.com", "user2", "password")
	tweetManager.AddUser(user2)
	tweetManager.Login(user2.Nick, user2.Pass)

	text2 := "This is user2's tweet"
	tweet2 := domain.NewQuoteTweet(user2, text2, tweet1)

	//Operation
	tweetManager.PublishTweet(tweet2)

	//Validation
	publishedTweet := tweetManager.GetTweets()[0].(*domain.QuoteTweet)
	
	if !service.CompareUsers(publishedTweet.GetUser(), user2) ||
		publishedTweet.GetText() != text2 {
			t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", 
				user2.Nick, text2, publishedTweet.GetUser().Nick, publishedTweet.GetText())
		return	
	}
	if !reflect.DeepEqual(publishedTweet.Tweet, tweet1) {
		t.Errorf("Expected QuoteTweet's tweet is %s: %s \nbut is %s: %s", 
			user1.Nick, text1, publishedTweet.GetTweet().GetUser().Nick, publishedTweet.GetTweet().GetText())
		return	
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet
	
	var user *domain.User
	text := "this is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error if user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet

	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)

	var text string

	tweet = domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error if text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet
	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	tweet = domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error if text exceeds 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()


	var tweet, seccondTweet *domain.TextTweet
	user1 := domain.NewUser("user1", "user2@mail.com", "user1Nick", "password1")
	tweetManager.AddUser(user1)
	user2 := domain.NewUser("user2", "user2@mail.com", "user2Nick", "password2")
	tweetManager.AddUser(user2)
	tweetManager.Login(user1.Nick, user1.Pass)
	tweetManager.Login(user2.Nick, user2.Pass)

	tweet = domain.NewTextTweet(user1, "this is the first tweet")
	seccondTweet = domain.NewTextTweet(user2, "this is the seccond tweet")

	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(seccondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("expected size is 2 but was %d", len(publishedTweets))
		return
	}
	
	
	firstPublishedTweet := publishedTweets[0]
	seccondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, 0, tweet.GetUser(), tweet.GetText()) {
		return
	}
	if !isValidTweet(t, seccondPublishedTweet, 1, seccondTweet.GetUser(), seccondTweet.GetText()) {
		return
	}

}

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user *domain.User, text string) bool {
	if !service.CompareUsers(tweet.GetUser(), user) {
		t.Errorf("no match in users: tweet.User: %s, user: %s", tweet.GetUser().Nick, user.Nick)
		return false
	}
	if tweet.GetText() != text {
		t.Error("no match in texts")
		return false
	}
	if tweet.GetId() != id {
		t.Errorf("no match in ids: tweet.GetId() = %d, id = %d", tweet.GetId(), id)
		return false
	}
	return true
}

func TestCanRetrieveTweetById(t * testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int

	user := domain.NewUser("user", "user@mail.com", "userNick", "password")
	tweetManager.AddUser(user)
	tweetManager.Login(user.Nick, user.Pass)

	text := "this is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validaiton
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCantPublishIfUserDoesNotExist(t * testing.T) {
	//initialization
	tweetManager := service.NewTweetManager()

	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")

	text := "This is my first tweet"
	tweet := domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user does not exist" {
		t.Error("Expected error if user does not exist")
	}
	return	
}


func TestCantPublishIfUserDoesIsNotLoggedIn(t * testing.T) {
	//initialization
	tweetManager := service.NewTweetManager()

	user := domain.NewUser("dani", "dani@mail.com", "danidani", "password")
	tweetManager.AddUser(user)

	text := "This is my first tweet"
	tweet := domain.NewTextTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is not logged in" {
		t.Errorf("Expected error if user is not logged in and got %s", err)
	}
	return	
}

func TestGetTweetsByUser(t * testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	user1 := domain.NewUser("user1", "user1@mail.com", "user1Nick", "password")
	tweetManager.AddUser(user1)
	user2 := domain.NewUser("user2", "user2@mail.com", "user2Nick", "password")
	tweetManager.AddUser(user2)
	textsToTweetByUser1 := []string{"primer tweet de user1", "segundo tweet de user1", "tercer tweet de user1"}

	// Operation
	for _, textToTweet := range textsToTweetByUser1 {
		domain.NewTextTweet(user1, textToTweet)
	}
	text4 := "primer tweet de user2"
	domain.NewTextTweet(user2, text4)

	tweetsFromUser1,_ := tweetManager.GetTweetsByUser(user1)

	for i, tweet := range tweetsFromUser1 {
		if(tweet.GetText() != textsToTweetByUser1[i]) {
			t.Errorf("no match in texts: expected %s but was %s", textsToTweetByUser1[i], tweet.GetText())
		}
	}
}

func TestGetTweetsByUserReturnsErrorIfUserDoesNotExist(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	user1 := domain.NewUser("user1", "user1@mail.com", "user1Nick", "password")

	// Operation
	_, err := tweetManager.GetTweetsByUser(user1)

	// Validation
	if err != nil && err.Error() != "User does not exist" {
		t.Error("Expected error if user does not exist")
	}
}

func TestGetTweetsByUserReturnsErrorIfUserHasNoTweets(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	user1 := domain.NewUser("user1", "user1@mail.com", "user1Nick", "password")
	tweetManager.AddUser(user1)

	// Operation
	_, err := tweetManager.GetTweetsByUser(user1)

	// Validation
	if err != nil && err.Error() != "User had not tweet yet" {
		t.Error("Expected error if user had not tweet yet")
	}
}

func TestCanGetAPrintableTweet(t *testing.T) {
	// Initialization
	user1 := domain.NewUser("user1", "user1@mail.com", "user1Nick", "password")
	tweet := domain.NewTextTweet(user1, "this is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@user1Nick: this is my tweet"
	if text != expectedText {
		t.Errorf("the expected text is %s but was %s", expectedText, text)
	}
}

// func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {
// 	// Initialization
// 	var tweetWriter service.TweetWriter
// 	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
// 	tweetManager := service.NewTweetManager(tweetWriter)

// 	var tweet domain.Tweet // Fill the tweet with data
	
// 	// Operation
// 	id, _ := tweetManager.PublishTweet(tweet)

// 	// Validation
// 	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
// 	savedTweet := memoryWriter.GetLastSavedTweet()

// 	if savedTweet == nil {
// 		//TODO
// 	}
// 	if savedTweet.GetIt() != id {
// 		//TODO
// 	}
// }