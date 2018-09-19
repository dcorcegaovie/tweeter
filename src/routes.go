package main

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
	"net/http"
)

var TweetManager *service.TweetManager
var Quit chan bool

func InitializeService() {
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)
	TweetManager = service.NewTweetManager(tweetWriter)
	Quit = make(chan bool)
}

func main() {
	r := gin.Default()
	InitializeService()
	r.GET("/tweets", getTweet)
	r.POST("/tweets", createTweet)
	r.Run()
}

func getTweet(c *gin.Context) {
	c.JSON(http.StatusOK, TweetManager.GetTweets())
}

func createTweet(c *gin.Context) {
	var quoteTweet domain.QuoteTweet
	var imageTweet domain.ImageTweet
	var textTweet domain.TextTweet
	var tweet domain.PrintableTweet

	body, _ := c.GetRawData()

	if quoteId, err := jsonparser.GetInt(body, "quote"); err == nil {
		json.Unmarshal(body, &quoteTweet)
		quoteTweet.QuotedTweet = TweetManager.GetTweetById(int(quoteId))
		tweet = &quoteTweet
	} else if _, err := jsonparser.GetString(body, "url"); err == nil {
		json.Unmarshal(body, &imageTweet)
		tweet = &imageTweet
	} else if err := json.Unmarshal(body, &textTweet); err == nil {
		tweet = &textTweet
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	tweet.DecorateTweet()
	id, err := TweetManager.PublishTweet(tweet, Quit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, TweetManager.GetTweetById(id))
	}
}
