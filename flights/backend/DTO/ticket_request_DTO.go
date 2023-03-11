package DTO

type TicketRequest struct {
	FlightId             string `json:"flightId" bson:"flightId"`
	UserId               string `json:"userId" bson:"userId"`
	NumOfTicketsOptional []int  `json:"numOfTicketsOptional" bson:"numOfTicketsOptional"`
}
