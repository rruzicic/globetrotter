package model

type Notification struct {
	Model  `bson:",inline"`
	UserId string `json:"userId" bson:"user_id"`
	//enum: RESERVATION, CANCELLATION, RATING, A_RATING, HOST_STATUS, RESPONSE
	Type            string  `json:"type" bson:"type"`
	AccommodationId *string `json:"accommodationId,omitempty" bson:"accommodation_id,omitempty"`
	// Rating          *int    `json:"rating,omitempty" bson:"rating,omitempty"`
}