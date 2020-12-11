package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/omarsagoo/legitWiki/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.New()
	s.Start(os.Getenv("PORT"))
}
