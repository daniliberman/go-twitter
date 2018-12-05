package service

import (
	"github.com/daniliberman/twitter/src/domain"
	"time"
)

var tweet domain.Tweet

func PublishTweet (tweet2 *domain.Tweet) {
	tweet.User = tweet2.User
	tweet.Text = tweet2.Text
	date := time.Now()
	tweet.Date = &date
}

func GetTweet () domain.Tweet {
	return tweet;
}
