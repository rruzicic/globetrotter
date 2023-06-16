package services

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFlight(flight models.Flight) error {
	if err := repos.CreateFlight(flight); err != nil {
		return err
	}

	return nil
}

func DeleteFlight(id string) error {
	if err := repos.DeleteFlight(id); err != nil {
		return err
	}

	return nil
}

func GetAllFlights() ([]models.Flight, error) {
	flights, err := repos.GetAllFlights()

	if err != nil {
		return nil, err
	}

	return flights, nil
}

func GetFlightById(id string) (*models.Flight, error) {
	flight, err := repos.GetFlightById(id)

	if err != nil {
		return nil, err
	}

	return flight, err
}

func BuyTicket(flightId string, userEmail string, numOfTicketsOptional ...int) error { //numOfTicketsOptional is gonna be optional
	numOfTickets := 1                  //default value
	if len(numOfTicketsOptional) > 0 { //handling of default value
		numOfTickets = numOfTicketsOptional[0]
	}

	flight, err := GetFlightById(flightId)
	if err != nil {
		return err //Didn't find a flight with that id
	}
	if flight.Seats-numOfTickets < 0 {
		return err //Not enough seats on the flight
	}

	flightObjId, _ := primitive.ObjectIDFromHex(flightId)

	//reduce number of seats on the flight
	result, err := repos.FlightsCollection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": flightObjId}}, bson.M{"$set": bson.M{"seats": flight.Seats - numOfTickets}})
	if result.MatchedCount != 0 {
		log.Println("Updated number of seats on the flight")
	} else {
		log.Println("Didn't update number of seats on the flight")
	}
	if err != nil {
		return err //Failed to update number of seats
	}

	userId, err := FindUserByEmail(userEmail)

	if err != nil {
		log.Println("Could not find user with id")
	}

	ticket := models.Ticket{}
	ticket.Flight = *flight
	ticket.UserId = userId.Id.String()

	for numOfTickets > 0 {
		if _, err := repos.TicketsCollection.InsertOne(context.TODO(), ticket); err != nil {
			return err
		}
		numOfTickets--
	}

	return nil
}

func GetTicketsByUser(userId string) ([]models.Ticket, error) {
	tickets := []models.Ticket{}
	cursor, err := repos.TicketsCollection.Find(context.TODO(), bson.M{"user_id": bson.M{"$eq": userId}})
	if err != nil {
		return nil, err
	}

	cursor.All(context.TODO(), &tickets)

	return tickets, nil
}

func SearchFlights(searchFLightsDTO dto.SearchFlightsDTO) ([]models.Flight, error) {
	flights, err := repos.GetFlightBySearchParams(searchFLightsDTO)

	if err != nil {
		return nil, err
	}

	return flights, nil
}
