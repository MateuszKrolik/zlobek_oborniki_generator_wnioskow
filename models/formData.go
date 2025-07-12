package models

type FormData struct {
	Child    Child    `json:"Child"`
	Mother   Parent   `json:"Mother"`
	Father   Parent   `json:"Father"`
	Siblings Siblings `json:"Siblings"`
}
