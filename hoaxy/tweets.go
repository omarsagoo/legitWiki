package hoaxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// GetTweets will return tweets
func GetTweets(articles *Articles) {
	url := "https://api-hoaxy.p.rapidapi.com/tweets?ids=%5B29317%2C%2068363%5D"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("HOAXY_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &articles)
}
