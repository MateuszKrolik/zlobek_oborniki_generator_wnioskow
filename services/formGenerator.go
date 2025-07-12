package services

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"kindergarden_recruitment_app_pdf_gen/models"
)

type FormGenerator struct{}

func (FormGenerator) GeneratePage(pageNumber int, formData models.FormData) error {
	// Ensure output directory exists
	if err := os.MkdirAll("genOutput", 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Create template and output paths
	templatePath := fmt.Sprintf("templates/page%d.html", pageNumber)
	outputPath := filepath.Join("genOutput", fmt.Sprintf("page%d.html", pageNumber))

	// Parse template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error parsing page%d template: %v", pageNumber, err)
	}

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating page%d output file: %v", pageNumber, err)
	}
	defer outputFile.Close()

	// Execute template
	if err := tmpl.Execute(outputFile, formData); err != nil {
		return fmt.Errorf("error executing page%d template: %v", pageNumber, err)
	}

	fmt.Printf("Successfully generated %s\n", outputPath)
	return nil
}
