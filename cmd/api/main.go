package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shubhamwagh2544/go-social/internal/env"
	"github.com/shubhamwagh2544/go-social/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	store := store.NewStorage(nil)

	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8001"),
		},
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
