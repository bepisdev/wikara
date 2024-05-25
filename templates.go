package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

// initTemplates initializes the templates.
func initTemplates() {
	templates = template.Must(template.ParseFiles(
		filepath.Join("tmpl", "edit.html"),
		filepath.Join("tmpl", "view.html"),
	))
}

// renderTemplate renders a template.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
