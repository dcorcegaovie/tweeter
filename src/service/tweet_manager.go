package service

import (
	"errors"
	"github.com/tweeter/src/domain"
)

type TweetManager struct {
	Tweets             []domain.PrintableTweet
	TweetsMapByUser    map[string][]domain.PrintableTweet
	ChannelTweetWriter *ChannelTweetWriter
}

func NewTweetManager(channelTweetWriter *ChannelTweetWriter) *TweetManager {
	tweetManager := TweetManager{}
	tweetManager.InitializeService()
	tweetManager.ChannelTweetWriter = channelTweetWriter
	return &tweetManager
}

func (manager *TweetManager) InitializeService() {
	manager.Tweets = make([]domain.PrintableTweet, 0)
	manager.TweetsMapByUser = make(map[string][]domain.PrintableTweet)
}

func (manager *TweetManager) PublishTweet(new_tweet domain.PrintableTweet, quit chan bool) (int, error) {

	TweetsToWrite := make(chan domain.PrintableTweet)

	if len(new_tweet.GetUser()) == 0 {
		return 0, errors.New("user is required")
	}

	if len(new_tweet.GetText()) == 0 {
		return 0, errors.New("text is required")
	}

	if len(new_tweet.GetText()) > 140 {
		return 0, errors.New("text exceeds 140 characters")
	}

	user_tweets, exists := manager.TweetsMapByUser[new_tweet.GetUser()]

	if exists {
		user_tweets = append(user_tweets, new_tweet)
		manager.TweetsMapByUser[new_tweet.GetUser()] = user_tweets
	} else {
		user_tweets := make([]domain.PrintableTweet, 0)
		user_tweets = append(user_tweets, new_tweet)
		manager.TweetsMapByUser[new_tweet.GetUser()] = user_tweets
	}

	manager.Tweets = append(manager.Tweets, new_tweet)

	go manager.ChannelTweetWriter.WriteTweet(TweetsToWrite, quit)
	TweetsToWrite <- new_tweet
	close(TweetsToWrite)

	return new_tweet.GetId(), nil
}

func (manager *TweetManager) GetTweet() domain.PrintableTweet {
	return manager.Tweets[len(manager.Tweets)-1]
}

func (manager *TweetManager) GetTweets() []domain.PrintableTweet {
	return manager.Tweets
}

func (manager *TweetManager) GetTweetById(id int) domain.PrintableTweet {

	for _, tweet := range manager.Tweets {
		if tweet.GetId() == id {
			return tweet
		}
	}

	return nil
}

func (manager *TweetManager) GetTweetsByUser(user string) []domain.PrintableTweet {
	return manager.TweetsMapByUser[user]
}

func (manager *TweetManager) CountTweetsByUser(user string) int {
	return len(manager.TweetsMapByUser[user])
}
