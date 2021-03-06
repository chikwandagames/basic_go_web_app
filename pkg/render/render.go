package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// A FuncMap is a map of function that can be used in a template
var functions = template.FuncMap{}

// RenderTemplate is ...
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// ParseFiles("./templates/" + tmpl), loads the file ./templates folder
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// RendderTemplateTwo is ...
func RenderTemplateTwo(w http.ResponseWriter, files ...string) {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// RenderTemplateThree is ...
func RenderTemplateThree(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		// This will kill n exit the application
		log.Fatal(err)
	}

	// Retrieve the template out of the map
	// Use comma ok idiom to check if the variable exists
	t, ok := tc[tmpl]
	if !ok {
		// exit
		log.Fatal(err)
	}

	// We need to read the template from disk
	// We need to put that parsed template in memory into some bytes
	// Buffer to hold bytes
	buf := new(bytes.Buffer)
	// Execute the template and store in buf variable
	_ = t.Execute(buf, nil)
	// Write to ResponseWriter
	_, err = buf.WriteTo(w)

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
