package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"

	"kindergarden_recruitment_app_pdf_gen/models"
)

func main() {
	// Load seed data from JSON
	data, err := os.ReadFile("seedData/formData.json")
	if err != nil {
		log.Fatalf("Error reading seed file: %v", err)
	}

	var formData models.FormData
	err = json.Unmarshal(data, &formData)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Parse the template
	tmpl, err := template.ParseFiles("templates/html/page1.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Create output file
	outputFile, err := os.Create("generated_form.html")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Execute the template with our data
	err = tmpl.Execute(outputFile, formData)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	fmt.Println("Form successfully generated as generated_form.html")
}
