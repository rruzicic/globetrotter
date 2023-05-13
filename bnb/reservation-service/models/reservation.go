package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Model           `bson:",inline"`
	AccommodationId *primitive.ObjectID `json:"accommodationId" bson:"accommodation_id"`
	UserId          *primitive.ObjectID `json:"userId" bson:"user_id"`
	DateInterval    TimeInterval        `json:"dateInterval" bson:"date_interval"`
	NumOfGuests     int                 `json:"numOfGuests" bson:"num_of_guests"`
	IsApproved      bool                `json:"isApproved" bson:"is_approved"`
	TotalPrice      float32             `json:"totalPrice" bson:"total_price"`
}
