package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"kindergarden_recruitment_app_pdf_gen/models"
	"kindergarden_recruitment_app_pdf_gen/services"
)

func main() {
	// Load seed data from JSON
	data, err := os.ReadFile("seedData/formData.json")
	if err != nil {
		log.Fatalf("Error reading seed file: %v", err)
	}

	var formData models.FormData
	if err := json.Unmarshal(data, &formData); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	formGenerator := services.FormGenerator{}

	// Generate pages
	for pageNumber := 1; pageNumber <= 5; pageNumber++ {
		if err := formGenerator.GeneratePage(pageNumber, formData); err != nil {
			log.Fatalf("Error generating page %d: %v", pageNumber, err)
		}
	}

	fmt.Println("Form generation completed successfully")
}
