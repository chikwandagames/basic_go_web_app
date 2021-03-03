package render

import (
	"fmt"
	"html/template"
	"net/http"
)

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
