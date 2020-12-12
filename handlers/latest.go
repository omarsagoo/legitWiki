package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/omarsagoo/legitWiki/hoaxy"
)

// LatestArticlesHandler handles rendering all the Latest articles
func LatestArticlesHandler(c echo.Context) error {
	var articles hoaxy.Articles

	dur := c.Param("dur")
	if dur < "2" {
		dur = "2"
	}

	hoaxy.GetLatestArticles(&articles, dur)

	return c.Render(http.StatusOK, "latest.html", articles)
}
