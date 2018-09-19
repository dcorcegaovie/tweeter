package service

import (
	"github.com/tweeter/src/domain"
)

type MemoryTweetWriter struct {
	Tweets []domain.PrintableTweet
}

type ChannelTweetWriter struct {
	Memory *MemoryTweetWriter
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	tweetWriter := MemoryTweetWriter{make([]domain.PrintableTweet, 0)}
	return &tweetWriter
}

func NewChannelTweetWriter(writer *MemoryTweetWriter) *ChannelTweetWriter {
	channelTweetWriter := ChannelTweetWriter{writer}
	return &channelTweetWriter
}

func (writer *ChannelTweetWriter) WriteTweet(tweetToWrite chan domain.PrintableTweet, quit chan bool) {
	for {
		tweet, ok := <-tweetToWrite

		if !ok {
			quit <- false
			break
		} else {
			writer.Memory.Tweets = append(writer.Memory.Tweets, tweet)
		}
	}
}
