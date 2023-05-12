package models

type Price struct {
	Amount   float32      `json:"amount" bson:"amount"`
	Duration TimeInterval `json:"duration" bson:"duration"`
}
