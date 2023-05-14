package dtos

import "github.com/rruzicic/globetrotter/bnb/accommodation-service/models"

type UpdateAvailabilityDTO struct {
	AccommodationId string              `json:"accommodationId"`
	NewInterval     models.TimeInterval `json:"newInterval"`
}
