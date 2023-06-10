package dtos

import "time"

type SearchFlightsDTO struct {
	ArrivalDateTime   time.Time `json:"arrivalDateTime"`
	Destination       string    `json:"destination"`
	DepartureDateTime time.Time `json:"departureDateTime"`
	Departure         string    `json:"departure"`
	PassengerNumber   int       `json:"passengerNumber"`
}
