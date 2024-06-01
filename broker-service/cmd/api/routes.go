package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		MaxAge:         300,
	}))
	mux.Use(middleware.Logger)
	mux.Use(middleware.Heartbeat("/health"))

	mux.Post("/", brokerHandler)
	return mux
}
