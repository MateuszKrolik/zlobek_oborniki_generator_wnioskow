package models

type Parent struct {
	FullName  string `json:"FullName"`
	PESEL     string `json:"PESEL"`
	BirthDate string `json:"BirthDate"`
	Address   string `json:"Address"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	Workplace string `json:"Workplace"`
}
