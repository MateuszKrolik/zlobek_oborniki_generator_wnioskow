package models

type Child struct {
	// page 1
	FullName        string   `json:"FullName"`
	BirthDate       string   `json:"BirthDate"`
	PESEL           string   `json:"PESEL"`
	Address         string   `json:"Address"`
	Diet            string   `json:"Diet"`
	DevelopmentInfo string   `json:"DevelopmentInfo"`
	Siblings        Siblings `json:"Siblings"`
	// page 2
	HasDisability         bool `json:"HasDisability"`
	NeedsSpecialEducation bool `json:"NeedsSpecialEducation"`
}

type Siblings struct {
	Count int    `json:"Count"`
	Ages  string `json:"Ages"` // comma separated list of ints
}
