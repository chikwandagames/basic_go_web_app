package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chikwandagames/basic_go_web_app/pkg/config"
	"github.com/chikwandagames/basic_go_web_app/pkg/handlers"
	"github.com/chikwandagames/basic_go_web_app/pkg/render"
)

const portNumber = ":8080"

// if A function or variable begins with a CAPITAL, this means its accessible outside
// that package

func main() {

	var app config.AppConfig
	// Load Templates cache, when the app starts
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	// Store template cache in the app
	app.TemplateCache = tc
	// We use this to load Templates from disc or cache, depending
	// on weather we in development of production mode
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	// Pass ther repo to the handlers
	handlers.NewHandlers(repo)

	// Pass the template cache to render package
	// giving the render package access to app (template)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	// Start a webserver
	// _ mean, if theres an error, forget about it
	_ = http.ListenAndServe(portNumber, nil)

}
