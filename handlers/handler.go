package handlers

import "github.com/dghubble/go-twitter/twitter"

// Handler holds the information needed for the handlers
type Handler struct {
	TwitterClient *twitter.Client
}