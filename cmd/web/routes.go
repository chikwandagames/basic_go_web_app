package main

import (
	"net/http"

	"github.com/chikwandagames/basic_go_web_app/pkg/config"
	"github.com/chikwandagames/basic_go_web_app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Here we user the bmizerany pat external package for routing

func routes(app *config.AppConfig) http.Handler {
	// Create a multiplexer, which is an http handler
	mux := chi.NewRouter()
	// Middlewere
	// Middleware lets you process a request as it comes into the application
	// and performes some action on it

	// Help application, Gracefully absorb panics and prints the stack trace
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
