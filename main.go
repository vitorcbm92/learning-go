package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Format the print to be a response type writer (api response)
	fmt.Fprintf(w, "Hello World!")
}

func handlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! 2!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/2", handlerTwo)
	fmt.Println("Server is running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
