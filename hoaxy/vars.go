package hoaxy

import "time"

// Articles holds a list of Articles from the API
type Articles struct {
	Articles []Article `json:"articles"`
	Count    int       `json:"num_of_entries"`
}

// Article holds an Article from the API
type Article struct {
	URL    string    `json:"canonical_url"`
	Date   time.Time `json:"date_captured"`
	Tweets int       `json:"number_of_tweets"`
	Type   string    `json:"site_type"`
	Title  string    `json:"title"`
}