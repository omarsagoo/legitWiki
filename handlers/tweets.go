package handlers

import (
	"github.com/labstack/echo"
	"github.com/omarsagoo/legitWiki/hoaxy"
)

// GetTweet does
func (h Handler) GetTweet(c echo.Context) error {
	hoaxy.GetTweets(h.TwitterClient)
	return nil
}
