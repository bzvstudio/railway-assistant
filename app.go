package main

import (
	"log"
	"net/http"
	"railway-assistant/handlers"
	"time"
)

type application struct {
	port string
}

func (app *application) mount() http.Handler {
	r := http.NewServeMux()

	r.Handle("/", http.HandlerFunc(handlers.HomeHandler))
	r.Handle("/robots.txt", http.FileServer(http.Dir(".")))
	r.Handle("/favicon.ico", http.HandlerFunc(handlers.FaviconHandler))
	r.Handle("/health", http.HandlerFunc(handlers.HealthHandler))
	r.Handle("/railway/alerts", http.HandlerFunc(handlers.RailwayAlertsHandler))

	return Chain(r, RequestID, RealIP, Logger, Recoverer)
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         ":" + app.port,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s", app.port)

	return srv.ListenAndServe()
}
