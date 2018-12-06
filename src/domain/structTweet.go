package domain

import (
	"time"
)


type Tweet struct {
	User *User
	Text string
	Date *time.Time
	Id int
}

func NewTweet(user *User, text string) *Tweet{
	var tweet Tweet
	tweet.User = user
	tweet.Text = text

	return &tweet
}
