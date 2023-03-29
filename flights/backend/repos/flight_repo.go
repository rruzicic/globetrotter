package repos

import (
	"context"
	"log"
	"time"

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

	startOfDay := time.Date(searchParams.ArrivalDateTime.Year(), searchParams.ArrivalDateTime.Month(), searchParams.ArrivalDateTime.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)
	startOfDay1 := time.Date(searchParams.DepartureDateTime.Year(), searchParams.DepartureDateTime.Month(), searchParams.DepartureDateTime.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay1 := startOfDay1.Add(24 * time.Hour)

    filter := bson.M{
        "$and": []bson.M{
            {"destination": bson.M{"$regex": searchParams.Destination, "$options": "i"}},
            {"departure": bson.M{"$regex": searchParams.Departure, "$options": "i"}},
        },
    }

    if !searchParams.ArrivalDateTime.IsZero() || !searchParams.DepartureDateTime.IsZero() {
        andClauses := make([]bson.M, 0)

        if !searchParams.ArrivalDateTime.IsZero() {
            andClauses = append(andClauses, bson.M{
                "arrival_date_time": bson.M{
                    "$gte": startOfDay,
                    "$lte": endOfDay,
                },
            })
        }

        if !searchParams.DepartureDateTime.IsZero() {
            andClauses = append(andClauses, bson.M{
                "departure_date_time": bson.M{
                    "$gte": startOfDay1,
                    "$lte": endOfDay1,
                },
            })
        }

        filter["$and"] = andClauses
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
