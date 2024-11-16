package main

import (
	"log"

	"github.com/apoorv-x12/social-media-app/internal/env"
	"github.com/apoorv-x12/social-media-app/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	config := &config{
		address: env.GetStr("addr",":8080"),
	}
    
    storage, err :=store.NewStorage(nil)
	if err!=nil {
		log.Fatal(err)
	}

	app := &application{
		config: *config,
		store: *storage,
	}

	router := app.route()

	log.Fatal(app.serve(router))
}