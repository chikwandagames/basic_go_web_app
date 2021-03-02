package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

// if A function or variable begins with a CAPITAL, this means its accessible outside
// that package

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

// Divide is ...
func Divide(w http.ResponseWriter, r *http.Request) {

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// ParseFiles("./templates/" + tmpl), loads the file ./templates folder
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %v \n", port)

	// Start a webserver
	// _ mean, if theres an error, forget about it
	_ = http.ListenAndServe(port, nil)

}
