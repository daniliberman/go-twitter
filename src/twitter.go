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

			c.Print("Write your user's nick: ")

			nick := c.ReadLine()
			user := service.GetUserWithNick(nick)

			if user == nil {
				c.Printf("User with nick: %s does not exist\n", nick)
				return
			}

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)
			id, error := service.PublishTweet(tweet)

			if error == nil {
				c.Printf("Tweet %d sent\n", id)
			} else {
				c.Printf("Tweet faild with error: %s\n", error)
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
				c.Println("[" + strconv.Itoa(tweets[i].Id) + "]" + tweets[i].User.Nick + ": " + tweets[i].Text)
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

			c.Println(tweet.User.Nick + ": " + tweet.Text)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "addUser",
		Help: "adds new twitter user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			//name, mail, nick, pass
			c.Print("Write the name: ")
			name := c.ReadLine()

			c.Print("Write the mail: ")
			mail := c.ReadLine()

			c.Print("Write the nick: ")
			nick := c.ReadLine()

			c.Print("Write the password: ")
			pass := c.ReadLine()

			user := domain.NewUser(name, mail, nick, pass)
			error := service.AddUser(user)

			if error == nil {
				c.Printf("User %s was added\n", nick)
			} else {
				c.Print("Adding user faild\n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUsers",
		Help: "Shows all users",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			users := service.GetUsers()

			for i := 0 ; i < len(users); i++ {
				c.Println(users[i].Nick)
			}
			return
		},
	})

	shell.Run()

}