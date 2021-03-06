package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/chikwandagames/basic_go_web_app/pkg/config"
)

// A FuncMap is a map of function that can be used in a template
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplateThree is ...
func RenderTemplateThree(w http.ResponseWriter, tmpl string) {
	// Get the template cache from the app config
	tc := app.TemplateCache

	// Retrieve the template out of the map
	// Use comma ok idiom to check if the variable exists
	t, ok := tc[tmpl]
	if !ok {
		// exit
		log.Fatal("could not get template from cache")
	}

	// We need to read the template from disk
	// We need to put that parsed template in memory into some bytes
	// Buffer to hold bytes
	buf := new(bytes.Buffer)
	// Execute the template and store in buf variable
	_ = t.Execute(buf, nil)
	// Write to ResponseWriter
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}

}

// CreateTemplateCache creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// myCache is template cache to hold all templates, and is created
	// at the start of the application, key string, value template pointer
	myCache := map[string]*template.Template{}

	// Add files with .page.html suffix to pages var
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// Because we don't want the full path, we remove the ./templates/
		name := filepath.Base(page)
		// fmt.Println("Page is currently", page)
		// fmt.Println("name is currently", name)
		// Create a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// If we find a match, i.e. a file with the suffix, ./templates/*.layout.html
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		// Now add the template to the cache
		myCache[name] = ts

	}
	return myCache, nil
}
