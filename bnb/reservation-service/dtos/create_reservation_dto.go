package dtos

import "github.com/rruzicic/globetrotter/bnb/reservation-service/models"

type CreateReservationDTO struct {
	AccommodationId string              `json:"accommodationId"`
	UserId          string              `json:"userId"`
	DateInterval    models.TimeInterval `json:"dateInterval"`
	NumOfGuests     int                 `json:"numOfGuests"`
	TotalPrice      float32             `json:"totalPrice"`
}
