package main

import "log"

func main() {
	app := &application{
		config: config{
			addr: ":8001",
		},
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
