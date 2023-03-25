package repos

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/flights/backend/dto"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFlight(flight models.Flight) error {
	if _, err := FlightsCollection.InsertOne(context.TODO(), flight); err != nil {
		log.Panic("Could not create flight: ", flight)
		return err
	}

	return nil
}

func DeleteFlight(flight models.Flight) error {
	filter := bson.M{"_id": bson.M{"$eq": flight.Id}}
	if _, err := FlightsCollection.DeleteOne(context.TODO(), filter); err != nil {
		log.Panic("Could not delete flight ", flight)
		return err
	}

	return nil
}

func GetAllFlights() ([]models.Flight, error) {
	flights := []models.Flight{}
	cursor, err := FlightsCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Panic("Could not get all flights")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var flight models.Flight
		err := cursor.Decode(&flight)

		if err != nil {
			log.Panic("Could not decode unmarshall flight on cursor")
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
		log.Panic("Could not get object id from string: ", id)
		return nil, err
	}

	filter := bson.M{"_id": bson.M{"$eq": objectId}}
	if err := FlightsCollection.FindOne(context.TODO(), filter).Decode(&flight); err != nil {
		log.Panic("Could not find flight with id: ", id)
		return nil, err
	}

	return &flight, nil

}

func GetFlightBySearchParams(searchParams dto.SearchFlightsDTO) ([]models.Flight, error) {

	filter := bson.M{
        "$or": []bson.M{
            bson.M{"arrival_date_time": searchParams.ArrivalDateTime},
            bson.M{"departure_date_time": searchParams.DepartureDateTime},
            bson.M{"destination": bson.M{"$regex": searchParams.Destination, "$options": "i"}},
            bson.M{"departure": bson.M{"$regex": searchParams.Departure, "$options": "i"}},
        },
    }
	
	cursor, err := FlightsCollection.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }

	var flights []models.Flight
    if err := cursor.All(context.Background(), &flights); err != nil {
        return nil, err
    }

	return flights, nil;
}
