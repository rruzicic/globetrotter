package models

import (
	"time"
)

type ReservationStatus struct {
	ReservationCounter          int           `json:"reservationCounter" bson:"reservation_counter"`
	CanceledReservationsCounter int           `json:"canceledReservationCounter" bson:"canceled_reservation_counter"`
	TotalReservationDuration    time.Duration `json:"totalReservationDuration" bson:"total_reservation_duration"`
}
