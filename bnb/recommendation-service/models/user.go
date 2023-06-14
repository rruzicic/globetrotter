package models

type User struct {
	Name    string `json:"name"`
	MongoId string `json:"mongoId"`
}
