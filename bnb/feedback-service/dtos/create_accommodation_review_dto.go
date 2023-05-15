package dtos

type CreateAccommodationReviewDTO struct {
	Rating          int    `json:"rating"`
	UserId          string `json:"userId"`          //reviewer
	AccommodationId string `json:"accommodationId"` //reviewed
}
