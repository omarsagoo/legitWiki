package hoaxy

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

// GetTweets will return tweets
func GetTweets(twitterClient *twitter.Client) {
	tweet, _, _ := twitterClient.Statuses.Show(781709708726038529, nil)
	fmt.Println(tweet)
	// url := "https://api-hoaxy.p.rapidapi.com/tweets?ids=%5B29317%2C%2068363%5D"

	// req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("x-rapidapi-key", os.Getenv("HOAXY_KEY"))
	// req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	// res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// json.Unmarshal(body, &tweets)
}
