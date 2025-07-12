package models

type Child struct {
	FullName        string `json:"FullName"`
	BirthDate       string `json:"BirthDate"`
	PESEL           string `json:"PESEL"`
	Address         string `json:"Address"`
	Diet            string `json:"Diet"`
	DevelopmentInfo string `json:"DevelopmentInfo"`
}
