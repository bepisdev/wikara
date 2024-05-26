package api

import (
	"net/http"
	"github.com/joshburnsxyz/go-wiki/pkg/templates"
	"github.com/joshburnsxyz/go-wiki/pkg/page"
)

// viewHandler displays the view page.
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	templates.RenderTemplate(w, "view", p)
}

// frontPageHandler displays the front page.
func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := page.LoadPage("FrontPage")
	if err != nil {
		http.Redirect(w, r, "/edit/FrontPage", http.StatusFound)
		return
	}
	templates.RenderTemplate(w, "view", p)
}

// editHandler displays the edit page.
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	templates.RenderTemplate(w, "edit", p)
}

// saveHandler saves the edited page.
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &page.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
