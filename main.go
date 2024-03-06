package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Creating Web Server in Golang")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Batman 🦇")
	})

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
