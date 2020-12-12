package hoaxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GetLatestArticles will return all the latest articles
func GetLatestArticles(articles *Articles, duration string) {
	url := "https://api-hoaxy.p.rapidapi.com/latest-articles?past_hours=" + duration

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("HOAXY_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &articles)
	fmt.Println(len(articles.Articles))
}
