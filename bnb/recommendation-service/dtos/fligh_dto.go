package dtos

import "time"

type Flight struct {
	FlightId          string    `json:"flightId"`
	DepartureDateTime time.Time `json:"departureDateTime"`
	ArrivalDateTime   time.Time `json:"arrivalDateTime"`
	Departure         string    `json:"departure" `
	Destination       string    `json:"destination"`
	Price             float32   `json:"price"`
	Seats             int       `json:"seats"`
	Duration          int       `json:"duration"`
}
