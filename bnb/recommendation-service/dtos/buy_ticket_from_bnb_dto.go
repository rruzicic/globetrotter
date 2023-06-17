package dtos

type BuyTicketFromBnBDTO struct {
	ApiKey               string `json:"apiKey"`
	FlightId             string `json:"flightId"`
	NumOfTicketsOptional []int  `json:"numOfTicketsOptional"`
}
