package services

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFlight(flight models.Flight) error {
	// TODO: add created on

	if _, err := repos.FlightsCollection.InsertOne(context.TODO(), flight); err != nil {
		return err
	}

	return nil
}

func DeleteFlight(flight models.Flight) error {
	// TODO: add deleted on

	// most/all mongodb collection funcitons take a filter. Which represents the query basically
	filter := bson.M{"_id": bson.M{"$eq": flight.Id}}
	if _, err := repos.FlightsCollection.DeleteOne(context.TODO(), filter); err != nil {
		return err
	}

	return nil
}

func GetAllFlights() ([]models.Flight, error) {
	flights := []models.Flight{}
	// second arg represents query/filter by which to search for using bson names
	cursor, err := repos.FlightsCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var flight models.Flight
		err := cursor.Decode(&flight)

		if err != nil {
			return nil, err
		}

		flights = append(flights, flight)
	}

	return flights, nil
}

func GetFlightById(id string) (*models.Flight, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	flight := models.Flight{}

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": bson.M{"$eq": objectId}}
	if err := repos.FlightsCollection.FindOne(context.TODO(), filter).Decode(&flight); err != nil {
		return nil, err
	}

	return &flight, nil

}

func BuyTicket(flightId string, userId string, numOfTicketsOptional ...int) error { //numOfTickets je opcioni
	numOfTickets := 1                  //default vrednost
	if len(numOfTicketsOptional) > 0 { //hendlovanje opcionog parametra
		numOfTickets = numOfTicketsOptional[0]
	}

	flight, err := GetFlightById(flightId)
	if err != nil {
		return err //nije nasao flight sa tim Id-em
	}
	if flight.Seats-numOfTickets < 0 {
		return err //nema dovoljno slobodnih mesta
	}

	flightObjId, _ := primitive.ObjectIDFromHex(flightId)

	//smanji broj slobodnih mesta za taj let
	result, err := repos.FlightsCollection.UpdateOne(context.TODO(), bson.D{{"_id", flightObjId}}, bson.D{{"$set", bson.D{{"seats", flight.Seats - numOfTickets}}}})
	if result.MatchedCount != 0 {
		log.Println("Updated number of seats on the flight")
	} else {
		log.Println("Didn't update number of seats on the flight")
	}
	if err != nil {
		return err //Nije updejtovao broj mesta
	}

	ticket := models.Ticket{}
	ticket.Flight = *flight
	ticket.UserId = userId

	for numOfTickets > 0 {
		if _, err := repos.TicketsCollection.InsertOne(context.TODO(), ticket); err != nil {
			return err
		}
		numOfTickets--
	}

	return nil
}
