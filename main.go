package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, world!")
	if err != nil {
		log.Fatal(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "About page")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	log.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
