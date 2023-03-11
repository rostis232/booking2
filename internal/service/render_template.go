package service

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

const templateFilesExtension string = ".go.html"

func RenderTamplate (w http.ResponseWriter, tmpl string) {

	_, err := RenderTamplateTest()
	if err != nil {
		log.Println(err)
	}
	
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl + templateFilesExtension)
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

func RenderTamplateTest () (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("/templates/*.page"+templateFilesExtension)
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		matches, err := filepath.Glob("./templates/*.layout"+templateFilesExtension)
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout"+templateFilesExtension)
			if err != nil {
				return nil, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}