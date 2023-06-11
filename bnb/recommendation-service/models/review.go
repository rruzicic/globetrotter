package models

type Review struct {
	Value                int    `json:"value"`
	MongoId              string `json:"mongoId"`
	UserMongoId          string `json:"userMongoId"`
	AccommodationMongoId string `json:"accommodationMongoId"`
}
