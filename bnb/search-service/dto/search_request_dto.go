package dto

import (
	"time"
)

type SearchRequestDTO struct {
	Location  string    `json:"location"`
	Guests    int       `json:"guests"` // number of guests that can be accommodated
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
