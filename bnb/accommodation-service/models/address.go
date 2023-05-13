package models

type Address struct {
	Country   string `json:"country" bson:"country"`
	City      string `json:"city" bson:"city"`
	Street    string `json:"street" bson:"street"`
	StreetNum string `json:"streetNum" bson:"street_num"`
	ZIPCode   int    `json:"zip" bson:"zip"`
}
