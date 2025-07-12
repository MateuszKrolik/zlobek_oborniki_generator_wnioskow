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
	// page 3
	IsSingleParentFamily bool   `json:"IsSingleParentFamily"`
	ParentDisability     bool   `json:"ParentDisability"`
	PreviouslyRejected   bool   `json:"PreviouslyRejected"`
	DifficultSituation   bool   `json:"DifficultSituation"`
	SituationDescription string `json:"SituationDescription"`
	LargeFamily          bool   `json:"LargeFamily"`
	Vaccinated           bool   `json:"Vaccinated"`
	VaccinationExemption bool   `json:"VaccinationExemption"`
	// page 5
	Points        int `json:"Points"`
	PendingPoints int `json:"PendingPoints"`
}

type Siblings struct {
	Count int    `json:"Count"`
	Ages  string `json:"Ages"` // comma separated list of ints
	// page 3
	SiblingInNursery  bool `json:"SiblingInNursery"`
	SiblingDisability bool `json:"SiblingDisability"`
}
