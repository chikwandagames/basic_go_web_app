package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world")
		// if we have an error
		if err != nil {
			fmt.Println(err)
		}

		// Sprintln lets you take different data types and return them as a string
		// fmt.Println(fmt.Sprintln("Bytes written: %d", n))
		fmt.Printf("Bytes written: %d \n", n)

	})

	// Start a webserver
	// _ mean, if theres an error, forget about it
	_ = http.ListenAndServe(":8080", nil)

}
