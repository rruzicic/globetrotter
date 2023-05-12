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

	_, err := acommodationsCollection.InsertOne(context.TODO(), accommodation)
	if err != nil {
		log.Panic("Could not save Accommodation because: ", err.Error())
		return err
	}
	return nil
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

	if _, err := acommodationsCollection.UpdateByID(context.TODO(), filter, update); err != nil {
		log.Panic("Could not update accommodation")
		return err
	}

	return nil
}
