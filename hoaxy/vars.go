package hoaxy

import (
	"time"
)

// Articles holds a list of Articles from the API
type Articles struct {
	Articles []Article `json:"articles"`
	Count    int       `json:"num_of_entries"`
}

// Article holds an Article from the API
type Article struct {
	URL         string    `json:"canonical_url"`
	Date        time.Time `json:"date_published"`
	Tweets      int       `json:"number_of_tweets"`
	Type        string    `json:"site_type"`
	Title       string    `json:"title"`
	ID          int       `json:"id"`
	CaptureDate time.Time `json:"date_captured"`
}

// Tweets holds a list of all the tweets and other data
type Tweets struct {
	Tweets []Tweet `json:"tweets"`
}

// Tweet holds information about an individual tweet
type Tweet struct {
	URL       string    `json:"canonical_url"`
	Published time.Time `json:"date_published"`
	Created   time.Time `json:"tweet_created_at"`
	TweetID   string    `json:"tweet_id"`
	Type      string    `json:"site_type"`
	Title     string    `json:"title"`
	Domain    string    `json:"domain"`
	ID        int       `json:"id"`
}
