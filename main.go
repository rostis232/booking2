package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func renderTamplate (w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Println("error parsing template")
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error executing template")
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTamplate(w, "home.page.tmpl")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTamplate(w, "about.page.tmpl")
}


func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	log.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
