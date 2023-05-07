package dto

import (
	"time"
)

type SearchDTO struct {
	Location  string    `json:"location"`
	Guests    int       `json:"guests"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
