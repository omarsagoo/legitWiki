package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/omarsagoo/legitWiki/hoaxy"
)

// IndexHandler handles the index route
func IndexHandler(c echo.Context) error {
	var articles hoaxy.Articles

	hoaxy.GetTopArticles(&articles)

	return c.Render(http.StatusOK, "index.html", articles)
}
