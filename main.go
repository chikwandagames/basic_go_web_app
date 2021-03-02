package main

import (
	"errors"
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
	// %d, for an int, %s for a string, %f for a float
	_, _ = fmt.Fprintf(w, fmt.Sprintf("About Page: 2 + 2 = %d", sum))
}

// Divide is ...
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(0.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Error: cannot divid by 0 \n")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("Divide Page: 0.0 / 0.0 = %f", f))
}

func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("Cannot devide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func addValues(x, y int) int {
	return x + y
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
