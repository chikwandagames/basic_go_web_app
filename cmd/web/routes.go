package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/chikwandagames/basic_go_web_app/pkg/config"
	"github.com/chikwandagames/basic_go_web_app/pkg/handlers"
)

// Here we user the bmizerany pat external package for routing

func routes(app *config.AppConfig) http.Handler {
	// Create a multiplexer, which is an http handler
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux

}
