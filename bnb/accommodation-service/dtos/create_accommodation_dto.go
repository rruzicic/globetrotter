package dtos

import "github.com/rruzicic/globetrotter/bnb/accommodation-service/models"

type CreateAccommodationDTO struct {
	Name                  string                `json:"name"`
	Location              models.Address        `json:"location"`
	AvailableCommodations []models.Commodations `json:"availableCommodations"`
	Photos                []string              `json:"photos"`
	Guests                int                   `json:"guests"`
	User                  string                `json:"user"`
	AutoApprove           bool                  `json:"autoApprove"`
}
