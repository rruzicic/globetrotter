package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	Id         *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId string        `json:"userId" bson:"user_id"`
	Flight
}
