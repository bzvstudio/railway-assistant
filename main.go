package main

import (
	"log"
	"railway-assistant/env"
)

func main() {
	env.Load()
	app := &application{port: env.GetString("PORT", "8080")}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
