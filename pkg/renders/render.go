package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/agniadvani/go-course/pkg/config"
	"github.com/agniadvani/go-course/pkg/models"
)

//Adds user-defined functions to the templates
var function = template.FuncMap{}

//Variable to be refrenced outside the package
var app *config.AppConfig

//Function for Initialising the variable outside the package
func NewTemplate(a *config.AppConfig) {
	app = a
}

//Adds default data to be parsed in every template
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

//Finds the parsed template from the map and executes it. To be used by the handlers.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("No template found.")
	}

	td = AddDefaultData(td)
	err := t.Execute(w, td)
	if err != nil {
		log.Fatalln(err)
	}

}

//Parses the pages, and associated layouts for every page and stores it in a map with the index of the page name
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(function).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
