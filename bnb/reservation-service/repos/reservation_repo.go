package repos

import (
	"context"
	"log"
	"time"

	grpcclient "github.com/rruzicic/globetrotter/bnb/reservation-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type CustomError struct{}

func (m *CustomError) Error() string {
	return "Cannot cancel a reservation less than one day in advance error"
}

func CreateReservation(reservation models.Reservation) (*models.Reservation, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	reservation.Id = &obj_id
	reservation.CreatedOn = int(time.Now().Unix())
	reservation.ModifiedOn = int(time.Now().Unix())

	inserted_id, err := reservationCollection.InsertOne(context.TODO(), reservation)

	if err != nil {
		log.Print("Could not create reservation! err: ", err.Error())
		return nil, err
	}
	id := inserted_id.InsertedID.(primitive.ObjectID)
	reservation.Id = &id

	if err := grpcclient.CreateReservation(reservation); err != nil {
		return nil, err
	}

	return &reservation, nil
}

func GetAllReservations() ([]models.Reservation, error) {
	reservations := []models.Reservation{}

	cursor, err := reservationCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Could not get all reservations")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var reservation models.Reservation
		if err := cursor.Decode(&reservation); err != nil {
			log.Println("Could not decode reservation from cursos")
			return nil, err
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
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

func GetFinishedReservationsByUser(id string) ([]models.Reservation, error) {
	allReservations, err := GetActiveReservationsByUser(id)
	if err != nil {
		log.Print("Could not reservations for user id: ", id)
		return nil, err
	}

	finishedReservations := []models.Reservation{}
	currentTime := time.Now()

	for _, reservation := range allReservations {
		if currentTime.After(reservation.DateInterval.End) {
			finishedReservations = append(finishedReservations, reservation)
		}
	}

	return finishedReservations, nil
}

func GetActiveReservationsByUser(id string) ([]models.Reservation, error) {
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get user id from string: ", id)
		return nil, err
	}

	reservations := []models.Reservation{}
	filter := bson.M{"user_id": bson.M{"$eq": userId}, "is_approved": bson.M{"$eq": true}}
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

	res, err := GetReservationById(id)
	if err != nil {
		return err
	}

	if err := grpcclient.DeleteReservation(*res); err != nil {
		return err
	}

	return nil
}

func UpdateReservation(reservation models.Reservation) error {
	reservation.ModifiedOn = int(time.Now().Unix())

	objID, err := primitive.ObjectIDFromHex(reservation.Id.Hex())
	if err != nil {
		log.Panic("Could not convert reservation hex to id")
		return err
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": bson.M{
		"date_interval": reservation.DateInterval,
		"num_of_guests": reservation.NumOfGuests,
		"is_approved":   reservation.IsApproved,
	},
	}

	if _, err := reservationCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Panic("Could not update reservation")
		return err
	}

	if err := grpcclient.UpdateReservation(reservation); err != nil {
		return err
	}

	return nil
}
