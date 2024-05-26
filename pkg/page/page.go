package page

import (
	"os"
	"html/template"
	"path/filepath"
	"log"
)

// Page struct holds the title and body of a wiki page.
type Page struct {
	Title string
	FTitle string
	Body  []byte
	HtmlContent template.HTML
}

const fileExtension = ".txt"

// Save method writes the Page's content to a text file.
func (p *Page) Save() error {
	dataDir := getContentDir()
	if err := ensureDir(dataDir); err != nil {
		return err
	}

	filename := filepath.Join(dataDir, p.Title+fileExtension)
	return os.WriteFile(filename, p.Body, 0600)
}

// LoadPage loads a Page from a text file.
func LoadPage(title string) (*Page, error) {
	dataDir := getContentDir()
	filename := filepath.Join(dataDir, title+fileExtension)
	body, err := os.ReadFile(filename)
	html := template.HTML(mdToHTML(body))
	ftitle := formatTitle(title)
	if err != nil {
		return nil, err
	}
	p := &Page{Title: title, FTitle: ftitle, Body: body, HtmlContent: html}
	log.Println(p)
	return p, nil
}

