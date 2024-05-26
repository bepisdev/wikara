package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/joshburnsxyz/go-wiki/pkg/page"
)

var (
	templates     map[string]*template.Template
	templatesLock sync.Mutex
)

func Init() {
	templates = make(map[string]*template.Template)
	loadTemplates()
}

func loadTemplates() {
	templates["view.html"] = template.Must(template.ParseFiles(filepath.Join("tmpl", "view.html"), filepath.Join("tmpl", "base.html")))
	templates["edit.html"] = template.Must(template.ParseFiles(filepath.Join("tmpl", "edit.html"), filepath.Join("tmpl", "base.html")))
}

func RenderTemplate(w http.ResponseWriter, tmplname string, p *page.Page) {
	templatesLock.Lock()
	defer templatesLock.Unlock()

	tmpl, ok := templates[tmplname+".html"]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err := tmpl.ExecuteTemplate(w, "base", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
