package models

import "time"

type Reservation struct {
	MongoId              string    `json:"mongoId"`
	UserMongoId          string    `json:"userMongoId"`
	AccommodationMongoId string    `json:"accommodationMongoId"`
	ReservationEnd       time.Time `json:"reservationEnd"`
}
