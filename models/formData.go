package models

import "time"

type FormData struct {
	Child               Child     `json:"Child"`
	Mother              Parent    `json:"Mother"`
	Father              Parent    `json:"Father"`
	CurrentDate         time.Time `json:"CurrentDate"         example:"2025-01-02T15:04:05Z"`
	CommitteeMemberName string    `json:"CommitteeMemberName"`
}
