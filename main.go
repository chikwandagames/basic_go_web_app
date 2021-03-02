package main

import (
	"fmt"
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
	fmt.Fprintf(w, "Home Page")
}

// About is a handler function
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	// %d, for an int, %s for a string
	_, _ = fmt.Fprintf(w, fmt.Sprintf("About Page: 2 + 2 = %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %v \n", port)

	// Start a webserver
	// _ mean, if theres an error, forget about it
	_ = http.ListenAndServe(port, nil)

}
