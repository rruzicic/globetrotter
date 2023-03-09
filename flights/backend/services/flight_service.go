package services

import (
	"context"
	"fmt"

	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFlight(flight models.Flight) {
	// TODO: add created on

	if _, err := repos.FlightsCollection.InsertOne(context.TODO(), flight); err != nil {
		fmt.Println("Could not save flight document into database")
		fmt.Println(err.Error())
	}
}

func DeleteFlight(flight models.Flight) {
	// TODO: add deleted on

	// most/all mongodb collection funcitons take a filter. Which represents the query basically
	filter := bson.M{"_id": bson.M{"$eq": flight.Id}}
	if _, err := repos.FlightsCollection.DeleteOne(context.TODO(), filter); err != nil {
		fmt.Println("Could not delete flight document from database")
		fmt.Println(err.Error())
	}
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
