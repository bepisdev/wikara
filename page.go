package main

import (
	"os"
)

// Page struct holds the title and body of a wiki page.
type Page struct {
	Title string
	Body  []byte
}

// save method writes the Page's content to a text file.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// loadPage loads a Page from a text file.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
