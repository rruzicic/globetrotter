package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	Model                 `bson:",inline"`
	Reservations          []*primitive.ObjectID `json:"reservations" bson:"reservations"`
	Name                  string                `json:"name" bson:"name"`
	Location              Address               `json:"location" bson:"location"`
	AvailableCommodations []Commodations        `json:"availableCommodations" bson:"available_commodations"`
	Photos                []string              `json:"photos" bson:"photos"` // b64 strings
	Guests                int                   `json:"guests" bson:"guests"`
	AvailabilityStartDate time.Time             `json:"availabilityStartDate" bson:"availability_start_date"`
	AvailabilityEndDate   time.Time             `json:"availabilityEndDate" bson:"availability_end_date"`
	UnitPrice             float32               `json:"unitPrice" bson:"unit_price"`            // price of 1 person per night or room per night
	PriceForPerson        bool                  `json:"priceForPerson" bson:"price_for_person"` // flag wether the unit price is for person/night or room/night
	User                  *primitive.ObjectID   `json:"user" bson:"user"`
	AutoApprove           bool                  `json:"autoApprove" bson:"auto_approve"` // wether to automatically approve reservations
}
