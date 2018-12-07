package domain

import (
	"time"
)

type Tweet interface {
	GetUser() *User
	GetText() string
	GetId() int
	GetDate() *time.Time
	SetDate(date *time.Time)
	SetId(id int)
	String() string
}

type TextTweet struct {
	User *User
	Text string
	Date *time.Time
	Id int
}

func NewTextTweet(user *User, text string) *TextTweet{
	var tweet TextTweet
	tweet.User = user
	tweet.Text = text

	return &tweet
}

func (tweet *TextTweet)String() string{
	str := "@"
	str += tweet.GetUser().Nick
	str += ": "
	str += tweet.GetText()

	return str
}

func (tweet *TextTweet)GetUser() *User{
	return tweet.User
}

func (tweet *TextTweet)GetText() string{
	return tweet.Text	
}

func (tweet *TextTweet)GetId() int{
	return tweet.Id
}

func (tweet *TextTweet)GetDate() *time.Time{
	return tweet.Date
}

func (tweet *TextTweet)SetDate(date *time.Time) {
	tweet.Date = date
}
func (tweet *TextTweet)SetId(id int){
	tweet.Id = id
}

type ImageTweet struct {
	TextTweet
	Url string
}

func NewImageTweet(user *User, text string, url string) *ImageTweet{
	imageTweet := ImageTweet{TextTweet:TextTweet{User: user, Text:text,}, Url: url,}

	return &imageTweet
}
func (tweet *ImageTweet)GetUrl()string{
	return tweet.Url
}
func (tweet *ImageTweet)String() string{
	str := "@"
	str += tweet.GetUser().Nick
	str += ": "
	str += tweet.GetText()
	str += ", image's url: "
	str += tweet.GetUrl()

	return str
}

type QuoteTweet struct {
	TextTweet
	Tweet Tweet
}

func NewQuoteTweet(user *User, text string, tweet Tweet) *QuoteTweet{
	quoteTweet := QuoteTweet{TextTweet:TextTweet{User: user, Text:text,}, Tweet: tweet,}

	return &quoteTweet
}
func (tweet *QuoteTweet)GetTweet()Tweet{
	return tweet.Tweet
}

func (tweet *QuoteTweet)String() string{
	str := "@"
	str += tweet.GetUser().Nick
	str += ": "
	str += tweet.GetText()
	str += ", quoted tweet: \n -->"
	str += tweet.GetTweet().String()

	return str
}