package dto

//location guest number date range

import (
	"time"
)

type SearchAccomodationDTO struct {
    Location    string `json:"location" binding:"required"`
    DateStart   time.Time `json:"dateStart" binding:"required"`
    DateEnd   time.Time `json:"dateEnd" binding:"required"`
    GuestNumber int `json:"guestNumber" binding:"required"`
}