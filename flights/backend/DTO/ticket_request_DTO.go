package dto

type TicketRequest struct {
	FlightId             string `json:"flightId" bson:"flight_id"`
	UserId               string `json:"userId" bson:"user_id"`
	NumOfTicketsOptional []int  `json:"numOfTicketsOptional" bson:"num_of_tickets_optional"`
}
