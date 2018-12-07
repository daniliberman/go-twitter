package service

import (
	"github.com/daniliberman/twitter/src/domain"
	"time"
	"fmt"
)

type TweetManager struct {
	Tweets []domain.Tweet
	NextId int
	TweetsByUsers map[*domain.User][]domain.Tweet
	Users []*domain.User
	LoggedInUsers []*domain.User
}

func NewTweetManager() *TweetManager{
	var tweetManager TweetManager
	tweetManager.Tweets = make([]domain.Tweet, 0) 
	tweetManager.TweetsByUsers = make(map[*domain.User][]domain.Tweet) 
	tweetManager.NextId = 0
	tweetManager.Users = make([]*domain.User, 0) 

	return &tweetManager
}

func (tweetManager *TweetManager)AddTweet(tweet domain.Tweet) error{
	tweetManager.Tweets = append(tweetManager.Tweets, tweet)
	return nil
//	return fmt.Errorf("adding tweet faild")
}

func (tweetManager *TweetManager)PublishTweet(tweet domain.Tweet) (int,error) {
	if tweet.GetUser() == nil {
		return -1, fmt.Errorf("user is required")	
	}

	if tweetManager.GetUserWithNick(tweet.GetUser().Nick) == nil {
		return -1, fmt.Errorf("user does not exist")
	}

	if !tweetManager.IsUserLoggedIn(tweet.GetUser()){
		return -1, fmt.Errorf("user is not logged in")
	}

	if tweet.GetText() == "" {
		return -1, fmt.Errorf("text is required")	
	}

	length := len(tweet.GetText())
	if length > 140 {
		return -1, fmt.Errorf("text exceeds 140 characters")
	}

	date := time.Now()
	tweet.SetDate(&date)
	tweet.SetId(tweetManager.NextId)
	tweetManager.NextId = tweetManager.NextId+1

	tweetManager.Tweets = append(tweetManager.Tweets, tweet)
	tweetManager.TweetsByUsers[tweet.GetUser()] = append(tweetManager.TweetsByUsers[tweet.GetUser()], tweet)

	return tweet.GetId(), nil
}

func (tweetManager *TweetManager)GetTweets() []domain.Tweet {
	return tweetManager.Tweets;
}

func (tweetManager *TweetManager)GetTweetById(id int) domain.Tweet {
	for _,tweet := range tweetManager.GetTweets() {
		if(tweet.GetId() == id){
			return tweet
		}
	}
	return nil
}

func (tweetManager *TweetManager)GetTweetsByUser(user *domain.User) ([]domain.Tweet, error) {

	userFound := tweetManager.GetUserWithNick(user.Nick)
	if userFound == nil {
		return nil, fmt.Errorf("User does not exist")	
	}

	tweetsForUser := tweetManager.TweetsByUsers[user]

	if tweetsForUser == nil {
		return nil, fmt.Errorf("User had not tweet yet")
	}

	return tweetManager.TweetsByUsers[user], nil
}
