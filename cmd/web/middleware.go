package main

import (
	"fmt"
	"net/http"
)

// WriteToConsole is ...
func WriteToConsole(next http.Handler) http.Handler {
	// The anolymous func is cast to a http.Handler func
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}
