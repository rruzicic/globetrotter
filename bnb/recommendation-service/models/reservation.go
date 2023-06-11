package models

type Reservation struct {
	MongoId              string `json:"mongoId"`
	UserMongoId          string `json:"userMongoId"`
	AccommodationMongoId string `json:"accommodationMongoId"`
}
