package main

import (
	"net/http"
)

// viewHandler displays the view page.
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// frontPageHandler displays the front page.
func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("FrontPage")
	if err != nil {
		http.Redirect(w, r, "/edit/FrontPage", http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// editHandler displays the edit page.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// saveHandler saves the edited page.
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
