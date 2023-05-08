package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Model           `bson:",inline"`
	AccommodationId *primitive.ObjectID `json:"accommodationId" bson:"accommodation_id"`
	UserId          *primitive.ObjectID `json:"userId" bson:"user_id"`
	StartDate       time.Time           `json:"startDate" bson:"start_date"`
	EndDate         time.Time           `json:"endtDate" bson:"end_date"`
	NumOfGuests     int                 `json:"numOfGuests" bson:"num_of_guests"`
	IsApproved      bool                `json:"isApproved" bson:"is_approved"`
}
