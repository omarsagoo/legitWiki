package hoaxy

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
)

// GetTweetsFromArticle gets a tweet from an article
func GetTweetsFromArticle(client *twitter.Client, articleID string) string {
	var tweets Tweets
	url := "https://api-hoaxy.p.rapidapi.com/tweets?ids=%5B" + articleID + "%5D"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("HOAXY_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &tweets)

	num := tweets.Tweets[0].TweetID

	return num
}
