package main

import (
	"github.com/abiosoft/ishell"
	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"
	"strconv"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)
			id, error := service.PublishTweet(tweet)

			if error == nil {
				c.Printf("Tweet %d sent\n", id)
			} else {
				c.Print("Tweet faild\n")
			}
			
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for i := 0 ; i < len(tweets); i++ {
				c.Println("[" + strconv.Itoa(tweets[i].Id) + "]" + tweets[i].User + ": " + tweets[i].Text)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows the tweet identified by an id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the id: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweet := service.GetTweetById(id)

			c.Println(tweet.User + ": " + tweet.Text)

			return
		},
	})

	shell.Run()

}