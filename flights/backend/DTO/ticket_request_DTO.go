package dto

type TicketRequest struct {
	FlightId             string `json:"flightId" bson:"flight_id"`
	UserEmail            string `json:"userEmail" bson:"user_email"`
	NumOfTicketsOptional []int  `json:"numOfTicketsOptional" bson:"num_of_tickets_optional"`
}
