package dtos

import "time"

type ReservationDTO struct {
	ReservationStartDate             time.Time `json:"reservationStartDate"`
	ReservationEndDate               time.Time `json:"reservationEndDate"`
	DepartureLocationToReservation   string    `json:"departureLocationToReservation"`   // location from which you want to take off when going to your accommodation
	ArrivalLocationAtReservation     string    `json:"arrivalLocationAtReservation"`     // location at which you want to land when going to your accommodation
	DepartureLocationFromReservation string    `json:"departureLocationFromReservation"` // location from which you want to take off when going home
	ArrivalLocationAtHome            string    `json:"arrivalLocationAtHome"`            // location at which you want to land when going home
	People                           int       `json:"people"`                           // how many people are staying at the reservation to know minimum fligh spaces needed
}
