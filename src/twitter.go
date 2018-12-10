package main

import (
	"github.com/abiosoft/ishell"
	"github.com/daniliberman/twitter/src/service"
	"github.com/daniliberman/twitter/src/domain"
	"github.com/daniliberman/twitter/src/rest"
	"strconv"
)

func main() {
	
	
	tweetManager := service.NewTweetManager()
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager.TweetWriter = tweetWriter

	ginServer := rest.NewGinServer(tweetManager)
	ginServer.StartGinServer()


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
			user := tweetManager.GetUserWithNick(nick)

			if user == nil {
				c.Printf("User with nick: %s does not exist\n", nick)
				return
			}

			c.Print("Write your tweet's kind:\n TextTweet, ImageTweet or QuoteTweet?\n")
			tweetsKind := c.ReadLine()

			c.Print("Write your tweet's text: ")
			text := c.ReadLine()
			var tweet domain.Tweet
			switch tweetsKind {
				case "ImageTweet", "image tweet", "image":
					c.Print("Enter your image's url: ")
					url := c.ReadLine()

					tweet = domain.NewImageTweet(user, text, url)
				case "QuoteTweet", "quote tweet", "quote":
					tweets := tweetManager.GetTweets()
					
					c.Print("These are the current tweets:")
					for i := 0 ; i < len(tweets); i++ {
						c.Println("[" + strconv.Itoa(tweets[i].GetId()) + "]" + tweets[i].String())
					}
					c.Print("Which one would you like to quote? Enter it's id: ")
					otherTweetId, _ := strconv.Atoi(c.ReadLine())
					quotedTweet := tweetManager.GetTweetById(otherTweetId)
					if quotedTweet == nil {
						c.Printf("Unable to tweet: '%d' is an invalid id!\n", otherTweetId)
					}
					tweet = domain.NewQuoteTweet(user, text, quotedTweet)
				case "TextTweet", "text tweet", "text":
					tweet = domain.NewTextTweet(user, text)
				default:
					c.Printf("Unable to tweet: '%s' is not a kind of tweet!\n", tweetsKind)
					return
			}

			_, error := tweetManager.PublishTweet(tweet)

			if error == nil {
				c.Printf("Tweet %s sent\n", tweet.String())
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

			tweets := tweetManager.GetTweets()

			for i := 0 ; i < len(tweets); i++ {
				c.Println("[" + strconv.Itoa(tweets[i].GetId()) + "]" + tweets[i].String())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsForUser",
		Help: "Shows all tweets for a specific user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user's nick: ")

			nick := c.ReadLine()
			user := tweetManager.GetUserWithNick(nick)

			if user == nil {
				c.Printf("User with nick: %s does not exist\n", nick)
				return
			}

			tweetsForUser, err := tweetManager.GetTweetsByUser(user)

			if err == nil {
				c.Printf("Tweets for user %s:\n", nick)
				for _,tweet := range(tweetsForUser) {
					c.Printf("%s\n", tweet.String())
				}
			} else {
				c.Printf("%s\n", err)
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

			tweet := tweetManager.GetTweetById(id)

			c.Println(tweet.String())

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
			error := tweetManager.AddUser(user)

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

			users := tweetManager.GetUsers()

			for i := 0 ; i < len(users); i++ {
				c.Println(users[i].Nick)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "allows user to login",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the nick: ")
			nick := c.ReadLine()

			c.Print("Write the password: ")
			pass := c.ReadLine()

			err := tweetManager.Login(nick, pass)

			if err == nil {
				c.Printf("Logged in successful\n")
			} else {
				c.Printf("%s\n", err)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "allows user to logout",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the nick: ")
			nick := c.ReadLine()

			err := tweetManager.Logout(nick)

			if err == nil {
				c.Printf("Logged out successful\n")
			} else {
				c.Printf("%s\n", err)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "checkLogin",
		Help: "checks if user is logged in",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the nick: ")
			nick := c.ReadLine()
			user := tweetManager.GetUserWithNick(nick)
			isLoggedIn := tweetManager.IsUserLoggedIn(user)

			if isLoggedIn {
				c.Printf("User with nick %s is logged in\n", nick)
			} else {
				c.Printf("User with nick %s is not logged in\n", nick)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "tweetsWithQuote",
		Help: "finds all tweets with specific quote",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the quote: ")
			quote := c.ReadLine()

			searchResult := make(chan domain.Tweet)
			tweetManager.SearchTweetsContaining(quote, searchResult)
			go func() {
				for toPrint := range searchResult {
					if toPrint != nil {
						c.Printf("%s\n", toPrint.String())
					}			
				}
			}()
			return
		},
	})

	shell.Run()

}