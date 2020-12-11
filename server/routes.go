package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/omarsagoo/legitWiki/handlers"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

// Routes renders all the routes for the server
func (s Server) Routes() {
	h := handlers.Handler{}
	h.TwitterClient = s.t

	// Logging and recovery
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	// render
	s.e.Renderer = renderer

	// public files
	s.e.Static("/", "public")

	s.e.GET("/", handlers.IndexHandler)
	s.e.GET("/top", handlers.TopArticlesHandler)
	s.e.GET("/tweet", h.GetTweet)
	s.e.GET("/searched", handlers.SearchArticlesHandler)
	s.e.GET("/tweet/:id", h.GetTweet)
}
