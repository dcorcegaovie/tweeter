package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type PrintableTweet interface {
	PrintableTweet() string
	DecorateTweet()
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
}

func (tweet *Tweet) GetUser() string {
	return tweet.User
}
func (tweet *Tweet) GetText() string {
	return tweet.Text
}
func (tweet *Tweet) GetDate() *time.Time {
	return tweet.Date
}
func (tweet *Tweet) GetId() int {
	return tweet.Id
}

type Tweet struct {
	Id   int        `json: "id"`
	User string     `json: "user"`
	Text string     `json: "text"`
	Date *time.Time `json: "date"`
}

type TextTweet struct {
	Tweet
}

type ImageTweet struct {
	Tweet
	Url string `json: "url" binding: "required"`
}

type QuoteTweet struct {
	Tweet
	QuotedTweet PrintableTweet `json: "quotedTweet" binding:"required"`
}

func (tweet *Tweet) DecorateTweet() {
	var time time.Time = time.Now()
	tweet.Id = rand.Int()
	tweet.Date = &time
}

func NewSimpleTweet(user string, text string) Tweet {
	var time time.Time = time.Now()
	return Tweet{rand.Int(), user, text, &time}
}

func (tweet *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func NewTextTweet(user string, text string) *TextTweet {
	return &TextTweet{NewSimpleTweet(user, text)}
}

func (tweet *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("%s", tweet.Tweet.PrintableTweet())
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
	return &ImageTweet{NewSimpleTweet(user, text), url}
}

func (tweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("%s %s", tweet.Tweet.PrintableTweet(), tweet.Url)
}

func NewQuoteTweet(user string, text string, tweet PrintableTweet) *QuoteTweet {
	return &QuoteTweet{NewSimpleTweet(user, text), tweet}
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf(`%s "%s"`, tweet.Tweet.PrintableTweet(), tweet.QuotedTweet.PrintableTweet())
}

func (tweet *Tweet) String() string {
	return tweet.PrintableTweet()
}
