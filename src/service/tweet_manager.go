package service

import (
	"github.com/daniliberman/twitter/src/domain"
	"time"
	"fmt"
)

var tweet domain.Tweet

func PublishTweet (tweet2 *domain.Tweet) error{
	if tweet2.User == "" {
		return fmt.Errorf("user is required")	
	}

	if tweet2.Text == "" {
		return fmt.Errorf("text is required")	
	}

	length := len(tweet2.Text)
	if length > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}

	tweet.User = tweet2.User
	tweet.Text = tweet2.Text
	date := time.Now()
	tweet.Date = &date
	return nil
}

func GetTweet () domain.Tweet {
	return tweet;
}
