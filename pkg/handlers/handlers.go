package handlers

import (
	"net/http"

	"github.com/chikwandagames/basic_go_web_app/pkg/config"
	"github.com/chikwandagames/basic_go_web_app/pkg/render"
)

// Here we use a repository pattern, to help us swap components
// in and out off the application with minimal changes to the code.
// We want the handlers to have access to config

// Repo is the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page, Handler function
// In order for a function to respond to requests from a web browser
// it has to handle 2 params, ResponseWriter and Request
// Home is ...
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")

}

// About is the handler for the about page
// Adding the receiver, gives About handler access to everything inside
// the repository
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")

}
