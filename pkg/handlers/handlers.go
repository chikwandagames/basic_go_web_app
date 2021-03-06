package handlers

import (
	"net/http"

	"github.com/chikwandagames/basic_go_web_app/pkg/render"
)

// Home page, Handler function
// In order for a function to respond to requests from a web browser
// it has to handle 2 params, ResponseWriter and Request
// Home is ...
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateThree(w, "home.page.html")

}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateThree(w, "about.page.html")

}
