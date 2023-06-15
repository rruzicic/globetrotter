package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationReview struct {
	Model           `bson:",inline"`
	Rating          int                 `json:"rating" bson:"rating"`
	UserId          *primitive.ObjectID `json:"userId" bson:"user_id"`                   //reviewer
	AccommodationId *primitive.ObjectID `json:"accommodationId" bson:"accommodation_id"` //reviewed
}
