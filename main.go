package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MateuszKrolik/zlobek_oborniki_generator_wnioskow/extensions"
	"github.com/MateuszKrolik/zlobek_oborniki_generator_wnioskow/models"
	"github.com/MateuszKrolik/zlobek_oborniki_generator_wnioskow/services"
)

func saveAsSinglePdf(pages extensions.Pages) {
	pdfBytes, err := pages.ToPdf("http://localhost:3001")
	if err != nil {
		log.Fatalf("Error converting to PDF: %v", err)
	}

	if err := os.WriteFile("genOutput/form.pdf", pdfBytes, 0644); err != nil {
		log.Fatalf("Error saving PDF: %v", err)
	}

	fmt.Println("PDF generation completed successfully")
}

func main() {
	data, err := os.ReadFile("seedData/formData.json")
	if err != nil {
		log.Fatalf("Error reading seed file: %v", err)
	}

	var formData models.FormData
	if err := json.Unmarshal(data, &formData); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	formGenerator, err := services.NewFormGenerator()
	if err != nil {
		log.Fatalf("Error instantiating generator: %v", err)
	}
	pages, err := formGenerator.GeneratePages(formData)
	if err != nil {
		log.Fatalf("Error generating pages: %v", err)
	}

	// save html pages to output dir
	for pageNum, content := range pages {
		fmt.Printf("Page %d size: %d bytes\n", pageNum, len(content))
		if err := os.WriteFile(fmt.Sprintf("genOutput/page%d.html", pageNum), content, 0644); err != nil {
			log.Printf("Warning: failed to save page %d: %v", pageNum, err)
		}
	}

	saveAsSinglePdf(pages)

	fmt.Println("Form generation completed successfully")
}
