package rest

import (
	"github.com/daniliberman/twitter/src/domain"
	"github.com/daniliberman/twitter/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GinServer struct {
	TweetManager *service.TweetManager
}

func NewGinServer(tweetManager *service.TweetManager) *GinServer{
	var ginServer GinServer
	ginServer.TweetManager = tweetManager

	return &ginServer
}

func (ginServer *GinServer)StartGinServer() {
	router := gin.Default()
	

	// router.GET("/unGet/:parametro", funcionQueHaceGet)
	// router.POST("/unPost", funcionQueHacePost)

	/*
	showTweets             Shows all tweets							GET
	showUsers              Shows all users							GET
	checkLogin             checks if user is logged in				GET
	showTweetById          Shows the tweet identified by an id		GET
	showTweetsForUser      Shows all tweets for a specific user		GET
	tweetsWithQuote        finds all tweets with specific quote		GET
	publishTweet           Publishes a tweet						POST
	addUser                adds new twitter user					POST
	login                  allows user to login						POST
	logout                 allows user to logout					POST
	*/
	
	router.GET("/showTweets", ginServer.ginShowTweets)
	router.GET("/showUsers", ginServer.ginShowUsers)
	router.GET("/checkLogin/:nick", ginServer.ginCheckLogin)
	router.GET("/showTweetById/:id", ginServer.showTweetById)
	router.GET("/showTweetsForUser/:nick", ginServer.showTweetsForUser)
	router.GET("/tweetsWithQuote/:quote", ginServer.ginTweetsWithQuote)

	router.POST("/publishTweet/text", ginServer.ginPublishTextTweet)
	router.POST("/publishTweet/image", ginServer.ginPublishImageTweet)
	router.POST("/publishTweet/quote", ginServer.ginPublishQuoteTweet)
	router.POST("/addUser", ginServer.ginAddUser)
	router.POST("/login", ginServer.ginLogin)
	router.POST("/logout", ginServer.ginLogout)

    go router.Run()

}

func (ginServer *GinServer)ginShowTweets(c * gin.Context) {
	c.JSON(http.StatusOK, ginServer.TweetManager.GetTweets())
}

func (ginServer *GinServer)ginShowUsers(c * gin.Context) {
	c.JSON(http.StatusOK, ginServer.TweetManager.GetUsers())
}

func (ginServer *GinServer)ginCheckLogin(c * gin.Context) {
	user := ginServer.TweetManager.GetUserWithNick(c.Param("nick"))
	c.JSON(http.StatusOK, ginServer.TweetManager.IsUserLoggedIn(user))
}

func (ginServer *GinServer)showTweetById(c * gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, ginServer.TweetManager.GetTweetById(id))
}

func (ginServer *GinServer)showTweetsForUser(c * gin.Context) {
	nick := c.Param("nick")
	user := ginServer.TweetManager.GetUserWithNick(nick)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with nick: " + nick + " does not exist\n"})
		return
	}

	tweetsForUser, err := ginServer.TweetManager.GetTweetsByUser(user)

	if err == nil {
		c.JSON(http.StatusOK, tweetsForUser)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ginServer *GinServer)ginTweetsWithQuote(c * gin.Context) {
	searchResult := make(chan domain.Tweet)
	ginServer.TweetManager.SearchTweetsContaining(c.Param("quote"), searchResult)
	tweetsList := make([]domain.Tweet, 0) 
	func() {
		for toPrint := range searchResult {
			if toPrint != nil {
				tweetsList = append(tweetsList, toPrint)
			}			
		}
	}()
	c.JSON(http.StatusOK, tweetsList)
}

func (ginServer *GinServer)ginPublishTextTweet(c *gin.Context) {
	var tweet domain.TextTweet
	err := c.ShouldBindJSON(&tweet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user := ginServer.TweetManager.GetUserWithNick(tweet.GetUser().Nick)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with nick: " + tweet.GetUser().Nick + " does not exist\n"})
		return
	}

	tweetToPublish := domain.NewTextTweet(user, tweet.GetText())

	_, err = ginServer.TweetManager.PublishTweet(tweetToPublish)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ginServer *GinServer)ginPublishImageTweet(c *gin.Context) {
	var imageTweet domain.ImageTweet
	err := c.ShouldBindJSON(&imageTweet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user := ginServer.TweetManager.GetUserWithNick(imageTweet.GetUser().Nick)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with nick: " + imageTweet.GetUser().Nick + " does not exist\n"})
		return
	}

	tweetToPublish := domain.NewImageTweet(user, imageTweet.GetText(), imageTweet.GetUrl())

	_, err = ginServer.TweetManager.PublishTweet(tweetToPublish)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

type GinQuoteTweet struct {
	domain.TextTweet
	QuotedTweetId int
}

func (ginServer *GinServer)ginPublishQuoteTweet(c *gin.Context) {
	var ginQuoteTweet GinQuoteTweet
	err := c.ShouldBindJSON(&ginQuoteTweet)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	quotedTweet := ginServer.TweetManager.GetTweetById(ginQuoteTweet.QuotedTweetId)
	if quotedTweet == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is not a tweet with that id\n"})
		return
	}

	user := ginServer.TweetManager.GetUserWithNick(ginQuoteTweet.GetUser().Nick)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with nick: " + ginQuoteTweet.GetUser().Nick + " does not exist\n"})
		return
	}

	tweetToPublish := domain.NewQuoteTweet(user, ginQuoteTweet.GetText(), quotedTweet)

	_, err = ginServer.TweetManager.PublishTweet(tweetToPublish)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ginServer *GinServer)ginAddUser(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ginServer.TweetManager.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ginServer *GinServer)ginLogin(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ginServer.TweetManager.Login(user.Nick, user.Pass)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ginServer *GinServer)ginLogout(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ginServer.TweetManager.Logout(user.Nick)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}