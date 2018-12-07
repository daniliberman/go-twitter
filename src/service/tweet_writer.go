package service

import (
	"github.com/daniliberman/twitter/src/domain"
	"os"
	//"time"
	//"fmt"
)

type TweetWriter interface {
	WriteTweet(tweet domain.Tweet)
}

type FileTweetWriter struct {
	SavedTweets *os.File
}
//////////////////////////////////////////////////////////////////////////////

type MemoryTweetWriter struct {
	lastTweet domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{}
}

func (m *MemoryTweetWriter) WriteTweet(tweet domain.Tweet) {
	m.lastTweet = tweet
}

func (m *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return m.lastTweet
}

//////////////////////////////////////////////////////////////////////////////

func NewFileTweetWriter() *FileTweetWriter{
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.SavedTweets = file

	return writer
}

func (fileTweetWriter FileTweetWriter) WriteTweet(tweet domain.Tweet) {
	go func() {
		byteSlice := []byte(tweet.String() + "\n")
		fileTweetWriter.SavedTweets.Write(byteSlice)
	}()
}