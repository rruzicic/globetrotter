package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Flight struct {
	// Model
	Id                bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DepartureDateTime time.Time     `json:"departureDateTime" bson:"departure_date_time"`
	Departure         string        `json:"departure" bson:"departure" `
	Destination       string        `json:"destination" bson:"destination"`
	Price             float32       `json:"price" bson:"price"`
	Seats             int           `json:"seats" bson:"seats"`
}
