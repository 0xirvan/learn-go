package main

import (
	"fmt"
	"net/http"
)

// Simple web server
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test")
	})

	fmt.Println("Server is running at http://localhost:3001")
	http.ListenAndServe(":3001", nil)
}
