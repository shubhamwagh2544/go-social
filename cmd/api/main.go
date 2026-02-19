package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shubhamwagh2544/go-social/internal/db"
	"github.com/shubhamwagh2544/go-social/internal/env"
	"github.com/shubhamwagh2544/go-social/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dbConfig := dbConfig{
		addr: env.GetString(
			"DB_ADDR",
			"postgres://postgres:postgres@localhost/go_social?sslmode=disable",
		),
		maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 10),
		maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
	}

	db, err := db.NewDbConn(
		dbConfig.addr,
		dbConfig.maxOpenConns,
		dbConfig.maxIdleConns,
		dbConfig.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("db connection pool established successfully")

	store := store.NewStorage(db)

	config := config{
		addr:     env.GetString("ADDR", ":8001"),
		dbConfig: dbConfig,
	}

	app := &application{
		config: config,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
