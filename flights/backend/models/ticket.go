package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Ticket struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId string        `json:"userId" bson:"user_id"`
	Flight
}
