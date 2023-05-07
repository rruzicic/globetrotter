package dto

import (
	"time"
)

type SearchResponseDTO struct {
	Location   string    `json:"location"`
	Guests     int       `json:"guests"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	TotalPrice float32   `json:"totalPrice"`
	UnitPrice  float32   `json:"unitPrice"` // price for 1 person per night
}
