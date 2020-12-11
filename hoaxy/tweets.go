package hoaxy

import (
	"github.com/dghubble/go-twitter/twitter"
)

// GetTweet will return the tweet
func GetTweet(client *twitter.Client, id int64) *twitter.Tweet {
	tweet, _, _ := client.Statuses.Show(id, nil)

	return tweet
}
