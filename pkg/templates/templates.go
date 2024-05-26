package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
	"github.com/joshburnsxyz/go-wiki/pkg/page"
)

var templates *template.Template

// renderTemplate renders a template.
func RenderTemplate(w http.ResponseWriter, tmplname string, p *page.Page) {
	tmpl := make(map[string]*template.Template)
	tmpl["view.html"] = template.Must(template.ParseFiles(filepath.Join("tmpl", "view.html"), filepath.Join("tmpl", "base.html")))
	tmpl["edit.html"] = template.Must(template.ParseFiles(filepath.Join("tmpl", "edit.html"), filepath.Join("tmpl", "base.html")))

	tmpl[tmplname+".html"].ExecuteTemplate(w, "base", p)
}
