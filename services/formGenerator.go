package services

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/MateuszKrolik/zlobek_oborniki_generator_wnioskow/models"
)

//go:embed templates/*.html
var templateFS embed.FS

type FormGenerator struct{}

func (FormGenerator) GeneratePages(formData models.FormData) (map[int][]byte, error) {
	pages := make(map[int][]byte)

	for pageNumber := 1; pageNumber <= 6; pageNumber++ {
		content, err := generatePage(pageNumber, formData)
		if err != nil {
			return nil, fmt.Errorf("failed to generate page %d: %w", pageNumber, err)
		}
		pages[pageNumber] = content
	}

	return pages, nil
}

func generatePage(pageNumber int, formData models.FormData) ([]byte, error) {
	// Read template from embedded filesystem
	tmplFile, err := templateFS.ReadFile(fmt.Sprintf("templates/page%d.html", pageNumber))
	if err != nil {
		return nil, fmt.Errorf("error reading template: %w", err)
	}

	// Parse the template
	tmpl, err := template.New(fmt.Sprintf("page%d", pageNumber)).Parse(string(tmplFile))
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, formData); err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	return buf.Bytes(), nil
}
