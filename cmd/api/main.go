package main

import (
	"log"
    "github.com/joho/godotenv"
	"github.com/apoorv-x12/social-media-app/internal/env"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	config := &config{
		address: env.GetStr("addr",":8080"),
	}

	app := &application{
		config: *config,
	}

	router := app.route()

	log.Fatal(app.serve(router))
}