package domain

import (
	"time"
)


type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet{
	var tweet Tweet
	tweet.User = user
	tweet.Text = text

	return &tweet
}
