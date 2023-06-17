package model

//type possible values: RESERVATION, CANCELLATION, RATING, A_RATING, HOST_STATUS, RESPONSE
type Notification struct {
	Model  				`bson:",inline"`
	UserId 				string `json:"userId" bson:"user_id"` //for whom the notification is
	Type              	string  `json:"type" bson:"type"`
	AccommodationId  	*string `json:"accommodationId,omitempty" bson:"accommodation_id,omitempty"`
	AccommodationName 	*string `json:"accommodationName,omitempty" bson:"accommodation_name,omitempty"`
	Rating          	*int    `json:"rating,omitempty" bson:"rating,omitempty"`
	RaterId				*string `json:"raterId,omitempty" bson:"rater_id,omitempty"`
}