package main

import (
	"net/http"
)

// Home page, Handler function
// In order for a function to respond to requests from a web browser
// it has to handle 2 params, ResponseWriter and Request
// Home is ...
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is a handler function
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
