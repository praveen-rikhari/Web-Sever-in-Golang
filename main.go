package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Creating Web Server in Golang")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
