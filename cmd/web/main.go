package main

import (
	"fmt"
	"net/http"

	"github.com/chikwandagames/basic_go_web_app/pkg/handlers"
)

const port = ":8080"

// if A function or variable begins with a CAPITAL, this means its accessible outside
// that package

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %v \n", port)

	// Start a webserver
	// _ mean, if theres an error, forget about it
	_ = http.ListenAndServe(port, nil)

}
