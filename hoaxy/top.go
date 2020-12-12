package hoaxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// GetTopArticles will return all the top articles
func GetTopArticles(articles *Articles) {
	url := "https://api-hoaxy.p.rapidapi.com/top-articles?most_recent=true&exclude_tags=%5B%5D"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("HOAXY_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &articles)

}

// GetSearchArticles will return all the top articles
func GetSearchArticles(articles *Articles, query string) {
	url := "https://api-hoaxy.p.rapidapi.com/articles?query=" + query + "&sort_by=relevant&use_lucene_syntax=true"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("HOAXY_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &articles)
}
