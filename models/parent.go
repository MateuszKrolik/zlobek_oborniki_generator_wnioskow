package models

type Parent struct {
	// Page 1 fields
	FullName  string `json:"FullName"`
	PESEL     string `json:"PESEL"`
	BirthDate string `json:"BirthDate"`
	Address   string `json:"Address"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	Workplace string `json:"Workplace"`
	// Page 2
	IsEmployed         bool `json:"IsEmployed"`
	IsSelfEmployed     bool `json:"IsSelfEmployed"`
	IsStudent          bool `json:"IsStudent"`
	FiledTaxInOborniki bool `json:"FiledTaxInOborniki"`
	ResidesInOborniki  bool `json:"ResidesInOborniki"`
	// Page 4
	IsApplicant bool `json:"IsApplicant"`
}
