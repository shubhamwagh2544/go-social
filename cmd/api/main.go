package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shubhamwagh2544/go-social/internal/env"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	app := &application{
		config: config{
			addr: env.GetString("ADDR", ":8001"),
		},
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
