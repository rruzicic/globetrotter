package models

import (
	"time"
)

type Flight struct {
	Model             `bson:",inline"`
	DepartureDateTime time.Time `json:"departureDateTime" bson:"departure_date_time"`
	ArrivalDateTime   time.Time `json:"arrivalDateTime" bson:"arrival_date_time"`
	Departure         string    `json:"departure" bson:"departure" `
	Destination       string    `json:"destination" bson:"destination"`
	Price             float32   `json:"price" bson:"price"`
	Seats             int       `json:"seats" bson:"seats"`
	Duration          int       `json:"duration" bson:"duration"`
}
