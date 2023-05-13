package repos

import (
	"context"
	"log"
	"time"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAccommodation(accommodation models.Accommodation) error {
	accommodation.CreatedOn = int(time.Now().Unix())
	accommodation.ModifiedOn = int(time.Now().Unix())

	_, err := accommodationsCollection.InsertOne(context.TODO(), accommodation)
	if err != nil {
		log.Panic("Could not save Accommodation because: ", err.Error())
		return err
	}
	return nil
}

func GetAllAccommodations() ([]models.Accommodation, error) {
	var accommodations []models.Accommodation
	cursor, err := accommodationsCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Panic("Could not get all accommodations")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var accommodation models.Accommodation
		if err := cursor.Decode(&accommodation); err != nil {
			log.Panic("Could not decode accommodation from cursor")
			return nil, err
		}

		accommodations = append(accommodations, accommodation)
	}

	return accommodations, nil
}

func UpdateAccommodation(accommodation models.Accommodation) error {
	accommodation.ModifiedOn = int(time.Now().Unix())

	objID, err := primitive.ObjectIDFromHex(accommodation.Id.Hex())
	if err != nil {
		log.Panic("Could not convert accommodation hex to id")
		return err
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": bson.M{
		"reservations":           accommodation.Reservations,
		"available_commodations": accommodation.AvailableCommodations,
		"photos":                 accommodation.Photos,
		"availability":           accommodation.Availability,
		"unit_price":             accommodation.UnitPrice,
		"price_for_person":       accommodation.PriceForPerson,
		"auto_approve":           accommodation.AutoApprove,
	},
	}

	if _, err := accommodationsCollection.UpdateByID(context.TODO(), filter, update); err != nil {
		log.Panic("Could not update accommodation")
		return err
	}

	return nil
}

func GetAccommodationById(id string) (*models.Accommodation, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panic("Could not convert accommodation hex to id")
		return nil, err
	}

	var accommodation models.Accommodation
	filter := bson.M{"_id": bson.M{"$eq": objId}}
	if err := accommodationsCollection.FindOne(context.TODO(), filter).Decode(&accommodation); err != nil {
		return nil, err
	}

	return &accommodation, nil
}

func GetAccommodationsByHostId(id string) ([]models.Accommodation, error) {
	host_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panic("Could not get user id from hex: ", id)
		return nil, err
	}

	var accommodations []models.Accommodation
	filter := bson.M{"user": bson.M{"$eq": host_id}}
	cursor, err := accommodationsCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Panic("Could not get cursor for accommodations. Error: ", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var accommodation models.Accommodation
		err := cursor.Decode(&accommodation)

		if err != nil {
			log.Panic("Could not decode accommodation from cursor. Error: ", err)
			return nil, err
		}

		accommodations = append(accommodations, accommodation)
	}

	return accommodations, nil
}
