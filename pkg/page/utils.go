package page

import (
	"github.com/joshburnsxyz/wikara/pkg/utils"
	"github.com/spf13/viper"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getContentDir() string {
	p := fmt.Sprintf("%s/%s", utils.GetExecPath(), viper.GetString("ContentDir"))
	return p
}

// ensureDir checks if a directory exists, and creates it if it does not.
func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}

// SanitizeHTML sanitizes HTML content to prevent XSS attacks
func sanitizeHTML(htmlContent []byte) []byte{
    policy := bluemonday.UGCPolicy()
	return []byte(policy.Sanitize(string(htmlContent)))
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

func formatTitle(title string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	sub := re.FindAllString(title, -1)
	return strings.Join(sub, " ")
}
