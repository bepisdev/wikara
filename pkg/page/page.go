package page

import (
	"fmt"
	"os"
	"path/filepath"
)

// Page struct holds the title and body of a wiki page.
type Page struct {
	Title string
	Body  []byte
}

const fileExtension = ".txt"
const dataDir = "data"

// Save method writes the Page's content to a text file.
func (p *Page) Save() error {
	if err := ensureDir(dataDir); err != nil {
		return err
	}
	filename := filepath.Join(dataDir, p.Title+fileExtension)
	return os.WriteFile(filename, p.Body, 0600)
}

// LoadPage loads a Page from a text file.
func LoadPage(title string) (*Page, error) {
	filename := filepath.Join(dataDir, title+fileExtension)
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// ensureDir checks if a directory exists, and creates it if it does not.
func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}
