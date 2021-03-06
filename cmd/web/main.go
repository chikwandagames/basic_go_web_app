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
	// Load Templates cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	// Store template cache in the app
	app.TemplateCache = tc

	// Pass the template cache to render package
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/base", handlers.Base)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	// Start a webserver
	// _ mean, if theres an error, forget about it
	_ = http.ListenAndServe(portNumber, nil)

}
