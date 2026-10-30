package main

import (
	"log"
	"railway-assistant/env"
	"railway-assistant/jobs"
	"railway-assistant/utils"
)

func main() {
	env.Load()

	if utils.IsRailwayStatusEnabled() {
		if utils.IsAtLeastOneProviderEnabled() {
			go jobs.StartRailwayStatusPoller()
		} else {
			log.Println("RAILWAY_STATUS_ENABLED is set but no notification provider is enabled, status polling disabled")
		}
	}

	app := &application{port: env.GetString("PORT", "8080")}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
