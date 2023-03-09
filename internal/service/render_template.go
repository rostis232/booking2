package service

import (
	"html/template"
	"net/http"
	"log"
)

func RenderTamplate (w http.ResponseWriter, tmpl string) {
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