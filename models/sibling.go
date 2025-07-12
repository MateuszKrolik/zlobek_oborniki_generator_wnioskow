package models

type Siblings struct {
	Count int    `json:"Count"`
	Ages  string `json:"Ages"` // comma separated list of int's
}
