package dtos

import "time"

type ReservationDTO struct {
	ReservationStartDate time.Time `json:"reservationStartDate"`
	ReservationEndDate   time.Time `json:"reservationEndDate"`
	ArrivalDestination   string    `json:"arrivalDestination"`   // place where you want to land when coming to the reservation place
	DepartureDestination string    `json:"departureDestination"` // place where you want to land coming from the reservation place
	People               int       `json:"people"`               // how many people are staying at the reservation to know minimum fligh spaces needed
}
