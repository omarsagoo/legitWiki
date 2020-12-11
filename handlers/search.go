package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/omarsagoo/legitWiki/hoaxy"
)

// SearchArticlesHandler handles rendering all the queried articles
func SearchArticlesHandler(c echo.Context) error {
	var articles hoaxy.Articles
	query := c.FormValue("search")
	hoaxy.GetSearchArticles(&articles, query)

	return c.Render(http.StatusOK, "search.html", map[string]interface{}{
		"articles": articles,
		"query":    query,
	})
}
