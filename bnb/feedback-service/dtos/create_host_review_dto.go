package dtos

type CreateHostReviewDTO struct {
	Rating int    `json:"rating"`
	UserId string `json:"userId"` //reviewer
	HostId string `json:"hostId"` //reviewed
}
