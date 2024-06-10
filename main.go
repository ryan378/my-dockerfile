package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone accessed me")
	fmt.Fprintf(w, "hahahaha!")
	fmt.Fprintf(w, "<html><body><h1>Here is a wise word: Stay curious!</h1></body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
