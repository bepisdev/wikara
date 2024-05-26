package page

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/joshburnsxyz/wikara/pkg/utils"
	"github.com/spf13/viper"
)

// Page struct holds the title and body of a wiki page.
type Page struct {
	Title string
	Body  []byte
	HtmlContent template.HTML
}

const fileExtension = ".txt"

func getContentDir() string {
	p := fmt.Sprintf("%s/%s", utils.GetExecPath(), viper.GetString("ContentDir"))
	return p
}

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
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body, HtmlContent: html}, nil
}

// ensureDir checks if a directory exists, and creates it if it does not.
func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
