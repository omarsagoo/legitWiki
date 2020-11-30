package server

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/labstack/echo"
	"github.com/omarsagoo/legitWiki/config"
)

// Server is a wrapper around our core
type Server struct {
	e *echo.Echo
	t *twitter.Client
}

//New will create a new instance of the server.
func New() *Server {

	return &Server{
		e: echo.New(), // new echo server to server the api
		t: config.GetTwitterClient(),
	}
}

// Start will start the server instance
func (s *Server) Start(port string) {

	// default port 8080
	if port == "" {
		port = ":8080"
	}

	// register routes
	s.Routes()
	// start server
	s.e.Logger.Fatal(s.e.Start(port))
}

// Stop will stop the server
func (s *Server) Stop() {
	// stop the server
	s.e.Close()
}
