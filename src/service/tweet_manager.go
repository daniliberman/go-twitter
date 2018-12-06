package service

import (
	"github.com/daniliberman/twitter/src/domain"
	"time"
	"fmt"
)

// Initialization of tweets slice
var tweets []*domain.Tweet

func InitializeServiceTweet() {
	tweets = make([]*domain.Tweet, 0) 
}

func AddTweet(tweet *domain.Tweet) error{
	tweets = append(tweets, tweet)
	return nil
//	return fmt.Errorf("adding tweet faild")
}

func PublishTweet(tweet *domain.Tweet) error{
	if tweet.User == "" {
		return fmt.Errorf("user is required")	
	}

	if tweet.Text == "" {
		return fmt.Errorf("text is required")	
	}

	length := len(tweet.Text)
	if length > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}

	date := time.Now()
	tweet.Date = &date

	tweets = append(tweets, tweet)

	return nil
}

func GetTweets() []*domain.Tweet {
	return tweets;
}

