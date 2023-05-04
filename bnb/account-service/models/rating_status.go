package models

type RatingStatus struct {
	Rating    float32 `json:"rating" bson:"rating"`
	RatingNum int     `json:"ratingNum" bson:"rating_num"`
}
