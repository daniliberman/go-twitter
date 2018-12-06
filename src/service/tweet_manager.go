package service

import (
	"github.com/daniliberman/twitter/src/domain"
	"time"
	"fmt"
)

// Initialization of tweets slice
var tweets []*domain.Tweet
var nextId int
var tweetsByUsers map[*domain.User][]*domain.Tweet

func InitializeServiceTweet() {
	tweets = make([]*domain.Tweet, 0) 
	tweetsByUsers = make(map[*domain.User][]*domain.Tweet) 
	nextId = 0
}

func AddTweet(tweet *domain.Tweet) error{
	tweets = append(tweets, tweet)
	return nil
//	return fmt.Errorf("adding tweet faild")
}

func PublishTweet(tweet *domain.Tweet) (int,error) {
	if tweet.User == nil {
		return -1, fmt.Errorf("user is required")	
	}

	if GetUserWithNick(tweet.User.Nick) == nil {
		return -1, fmt.Errorf("user does not exist")
	}

	if !IsUserLoggedIn(tweet.User){
		return -1, fmt.Errorf("user is not logged in")
	}

	if tweet.Text == "" {
		return -1, fmt.Errorf("text is required")	
	}

	length := len(tweet.Text)
	if length > 140 {
		return -1, fmt.Errorf("text exceeds 140 characters")
	}

	date := time.Now()
	tweet.Date = &date
	tweet.Id = nextId
	nextId = nextId+1

	tweets = append(tweets, tweet)
	tweetsByUsers[tweet.User] = append(tweetsByUsers[tweet.User], tweet)

	return tweet.Id, nil
}

func GetTweets() []*domain.Tweet {
	return tweets;
}

func GetTweetById(id int) *domain.Tweet {
	var tweet *domain.Tweet
	for i := 0; i < len(GetTweets()); i++ {
		tweet = GetTweets()[i];
		if(tweet.Id == id){
			return tweet
		}
	}
	return nil
}

func GetTweetsByUser(user *domain.User) ([]*domain.Tweet, error) {

	userFound := GetUserWithNick(user.Nick)
	if userFound == nil {
		return nil, fmt.Errorf("User does not exist")	
	}

	tweetsForUser := tweetsByUsers[user]

	if tweetsForUser == nil {
		return nil, fmt.Errorf("User had not tweet yet")
	}

	return tweetsByUsers[user], nil
}
