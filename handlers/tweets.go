package handlers

import (
	"github.com/labstack/echo"
	"github.com/omarsagoo/legitWiki/hoaxy"
)

// GetTweet does
func (h Handler) GetTweet(c echo.Context) error {
	ID := c.Param("id")
	tweet := hoaxy.GetTweetsFromArticle(h.TwitterClient, ID)
	return c.JSON(200, tweet)
}
