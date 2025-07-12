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

type FormGenerator struct {
	templates map[int]*template.Template
}

func NewFormGenerator() (*FormGenerator, error) {
	fg := &FormGenerator{
		templates: make(map[int]*template.Template),
	}

	// Pre-load all templates
	for i := 1; i <= 6; i++ {
		tmplFile, err := templateFS.ReadFile(fmt.Sprintf("templates/page%d.html", i))
		if err != nil {
			return nil, fmt.Errorf("failed to load template %d: %w", i, err)
		}

		tmpl, err := template.New(fmt.Sprintf("page%d", i)).Parse(string(tmplFile))
		if err != nil {
			return nil, fmt.Errorf("failed to parse template %d: %w", i, err)
		}
		fg.templates[i] = tmpl
	}

	return fg, nil
}

func (fg *FormGenerator) GeneratePages(formData models.FormData) (map[int][]byte, error) {
	pages := make(map[int][]byte)

	for pageNumber, tmpl := range fg.templates {
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, formData); err != nil {
			return nil, fmt.Errorf("failed to generate page %d: %w", pageNumber, err)
		}
		pages[pageNumber] = buf.Bytes()
	}

	return pages, nil
}
