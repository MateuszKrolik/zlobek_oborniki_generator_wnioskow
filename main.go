package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"kindergarden_recruitment_app_pdf_gen/models"
	"kindergarden_recruitment_app_pdf_gen/services"
)

func htmlToPDF(pages map[int][]byte) ([]byte, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	var mergedHTML bytes.Buffer
	for i := 1; i <= len(pages); i++ {
		mergedHTML.Write(pages[i])
	}

	// gotenberg requires at least 1 file named "index.html"
	part, err := writer.CreateFormFile("files", "index.html")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := part.Write(mergedHTML.Bytes()); err != nil {
		return nil, fmt.Errorf("failed to write HTML content: %w", err)
	}

	writer.Close()

	resp, err := http.Post(
		"http://localhost:3001/forms/chromium/convert/html",
		writer.FormDataContentType(),
		&body,
	)
	if err != nil {
		return nil, fmt.Errorf("Gotenberg request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Gotenberg error (%d): %s", resp.StatusCode, string(errorBody))
	}

	return io.ReadAll(resp.Body)
}

func saveAsSinglePdf(pages map[int][]byte) {
	pdfBytes, err := htmlToPDF(pages)
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

	formGenerator := services.FormGenerator{}
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
