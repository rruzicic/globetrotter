package models

type Accommodation struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Price    float32 `json:"price"`
	MongoId  string  `json:"mongoId"`
}
