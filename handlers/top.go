package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/omarsagoo/legitWiki/hoaxy"
)

// TopArticlesHandler handles rendering all the top articles
func TopArticlesHandler(c echo.Context) error {
	var articles hoaxy.Articles

	hoaxy.GetTopArticles(&articles)

	return c.Render(http.StatusOK, "top.html", articles)
}
