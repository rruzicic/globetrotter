package repos

import (
	"context"
	"log"
	"time"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type CustomError struct{}

func (m *CustomError) Error() string {
	return "Cannot cancel a reservation less than one day in advance error"
}

func CreateReservation(reservation models.Reservation) error {
	reservation.CreatedOn = int(time.Now().Unix())
	reservation.ModifiedOn = int(time.Now().Unix())

	//TODO:Provera da li je taj smestaj vec rezervisan u to vreme. grpc metoda od Accomodation servisa da vrati sve rezervacije za taj smestaj ovde.

	_, err := reservationCollection.InsertOne(context.TODO(), reservation)

	if err != nil {
		log.Print("Could not create reservation! err: ", err.Error())
		return err
	}
	return nil
}

func GetReservationById(id string) (*models.Reservation, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	reservation := models.Reservation{}

	if err != nil {
		log.Print("Could not get id from hex string: ", id)
	}

	filter := bson.M{"_id": bson.M{"$eq": objectId}}
	if err := reservationCollection.FindOne(context.TODO(), filter).Decode(&reservation); err != nil {
		log.Print("could not find reservation with id: ", id)
		return nil, err
	}

	return &reservation, nil
}

func GetReservationsByUserId(id string) ([]models.Reservation, error) {
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get user id from string: ", id)
		return nil, err
	}

	reservations := []models.Reservation{}
	filter := bson.M{"user_id": bson.M{"$eq": userId}}
	cursor, err := reservationCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Print("Could not get reservation")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var reservation models.Reservation
		err := cursor.Decode(&reservation)

		if err != nil {
			log.Print("Could not unmarshall reservation on cursor")
			return nil, err
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func GetReservationsByAccommodationId(id string) ([]models.Reservation, error) {
	acc_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panic("Could not get accommodation id from id: ", id)
		return nil, err
	}

	var reservations []models.Reservation
	filter := bson.M{"accommodation_id": bson.M{"$eq": acc_id}}
	cursor, err := reservationCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Panic("Could not get reservations from accommodation id")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var reservation models.Reservation
		err := cursor.Decode(&reservation)

		if err != nil {
			log.Panic("Could not unmarshall reservation on cursor")
			return nil, err
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func DeleteReservation(id string) error {
	reservation_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get id from hex string: ", id)
	}

	currentTime := time.Now()
	reservation, _ := GetReservationById(id)

	if currentTime.After(reservation.DateInterval.Start.Add(time.Hour * 24 * -1)) {
		log.Print("You cannot cancel a reservation less than one day in advance")
		return &CustomError{}
	}

	filter := bson.M{"_id": bson.M{"$eq": reservation_id}}
	if _, err := reservationCollection.DeleteOne(context.TODO(), filter); err != nil {
		log.Print("Could not delete flight with hex id", id)
		return err
	}

	//TODO: Povecati brojac otkazanih rezervacija u useru. grpc metoda prema Ratku

	return nil
}
