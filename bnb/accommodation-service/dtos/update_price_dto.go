package dtos

import "github.com/rruzicic/globetrotter/bnb/accommodation-service/models"

type UpdatePriceDTO struct {
	AccommodationId string              `json:"accommodationId"`
	NewPrice        float32             `json:"newPrice"`
	NewInterval     models.TimeInterval `json:"newInterval"`
}
