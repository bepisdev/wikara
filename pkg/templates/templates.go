package templates

import (
	"log"
	"net/http"

	layouts "github.com/joshburnsxyz/go-view-layouts"
	"github.com/joshburnsxyz/wikara/pkg/page"
)

var (
	templates map[string]string
)

func Init() {
	templates = map[string]string{
		"view": "tmpl/view.html",
		"edit": "tmpl/edit.html",
	}
	err := layouts.Init(templates, "tmpl/base.html")
	if err != nil {
		log.Fatalf("Failed to process templates: %v", err)
	}
}

func RenderTemplate(w http.ResponseWriter, tmplname string, p *page.Page) {
	layouts.RenderTemplate(w, tmplname, "base", p)
}
