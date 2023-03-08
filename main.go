package main

import (
	"log"
	"net/http"
	"github.com/rostis232/booking2/handlers"
)

const portNumber = ":8080"




func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Printf("Starting app on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
