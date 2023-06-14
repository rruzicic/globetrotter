package repos

import (
	"context"
	"log"
	"time"

	grpcclient "github.com/rruzicic/globetrotter/bnb/feedback-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func CreateHostReview(hostReview models.HostReview) (*models.HostReview, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	hostReview.Id = &obj_id
	hostReview.CreatedOn = int(time.Now().Unix())
	hostReview.ModifiedOn = int(time.Now().Unix())

	_, err := hostReviewCollection.InsertOne(context.TODO(), hostReview)

	if err != nil {
		log.Print("Could not create a review! err: ", err.Error())
		return nil, err
	}
	return &hostReview, nil
}

func GetHostReviewById(id string) (*models.HostReview, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	hostReview := models.HostReview{}

	if err != nil {
		log.Print("Could not get id from hex string: ", id)
	}

	filter := bson.M{"_id": bson.M{"$eq": objectId}}
	if err := hostReviewCollection.FindOne(context.TODO(), filter).Decode(&hostReview); err != nil {
		log.Print("could not find a host review with id: ", id)
		return nil, err
	}

	return &hostReview, nil
}

func GetHostReviewsByUserId(id string) ([]models.HostReview, error) {
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get user id from string: ", id)
		return nil, err
	}

	hostReviews := []models.HostReview{}
	filter := bson.M{"user_id": bson.M{"$eq": userId}}
	cursor, err := hostReviewCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Print("Could not get host review")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var hostReview models.HostReview
		err := cursor.Decode(&hostReview)

		if err != nil {
			log.Print("Could not unmarshall host review on cursor")
			return nil, err
		}

		hostReviews = append(hostReviews, hostReview)
	}

	return hostReviews, nil
}

func GetHostReviewsByHostId(id string) ([]models.HostReview, error) {
	hostId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get host id from string: ", id)
		return nil, err
	}

	hostReviews := []models.HostReview{}
	filter := bson.M{"host_id": bson.M{"$eq": hostId}}
	cursor, err := hostReviewCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Print("Could not get host review")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var hostReview models.HostReview
		err := cursor.Decode(&hostReview)

		if err != nil {
			log.Print("Could not unmarshall host review on cursor")
			return nil, err
		}

		hostReviews = append(hostReviews, hostReview)
	}

	return hostReviews, nil
}

func DeleteHostReview(id string) error {
	hostReviewId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get id from hex string: ", id)
	}

	filter := bson.M{"_id": bson.M{"$eq": hostReviewId}}
	if _, err := hostReviewCollection.DeleteOne(context.TODO(), filter); err != nil {
		log.Print("Could not delete host review with hex id", id)
		return err
	}

	return nil
}

func UpdateHostReview(hostReview models.HostReview) error {
	hostReview.ModifiedOn = int(time.Now().Unix())

	objId, err := primitive.ObjectIDFromHex(hostReview.Id.Hex())
	if err != nil {
		log.Print("Could not convert host review hex to id")
		return err
	}

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.M{"$set": bson.M{"rating": hostReview.Rating}}

	if _, err := hostReviewCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Print("Could not update host review")
		return err
	}

	return nil
}

//============================================================================================
//============================================================================================

func CreateAccommodationReview(accommodationReview models.AccommodationReview) (*models.AccommodationReview, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	accommodationReview.Id = &obj_id
	accommodationReview.CreatedOn = int(time.Now().Unix())
	accommodationReview.ModifiedOn = int(time.Now().Unix())

	_, err := accommodationReviewCollection.InsertOne(context.TODO(), accommodationReview)

	if err != nil {
		log.Print("Could not create a review! err: ", err.Error())
		return nil, err
	}

	if err := grpcclient.CreateReview(accommodationReview); err != nil {
		return nil, err
	}

	return &accommodationReview, nil
}

func GetAllAccommodationReviews() ([]models.AccommodationReview, error) {
	reviews := []models.AccommodationReview{}
	cursor, err := accommodationReviewCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Println("Could not get all accommodation reviews")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var review models.AccommodationReview
		if err := cursor.Decode(&review); err != nil {
			log.Println("Could not decode review at cursor. Error: ", err.Error())
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}

func GetAccommodationReviewById(id string) (*models.AccommodationReview, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	accommodationReview := models.AccommodationReview{}

	if err != nil {
		log.Print("Could not get id from hex string: ", id)
	}

	filter := bson.M{"_id": bson.M{"$eq": objectId}}
	if err := accommodationReviewCollection.FindOne(context.TODO(), filter).Decode(&accommodationReview); err != nil {
		log.Print("could not find a accommodation review with id: ", id)
		return nil, err
	}

	return &accommodationReview, nil
}

func GetAccommodationReviewsByUserId(id string) ([]models.AccommodationReview, error) {
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get user id from string: ", id)
		return nil, err
	}

	accommodationReviews := []models.AccommodationReview{}
	filter := bson.M{"user_id": bson.M{"$eq": userId}}
	cursor, err := accommodationReviewCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Print("Could not get accommodation review")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var accommodationReview models.AccommodationReview
		err := cursor.Decode(&accommodationReview)

		if err != nil {
			log.Print("Could not unmarshall accommodation review on cursor")
			return nil, err
		}

		accommodationReviews = append(accommodationReviews, accommodationReview)
	}

	return accommodationReviews, nil
}

func GetAccommodationReviewsByAccommodationId(id string) ([]models.AccommodationReview, error) {
	accommodationId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get accommodation id from string: ", id)
		return nil, err
	}

	accommodationReviews := []models.AccommodationReview{}
	filter := bson.M{"accommodation_id": bson.M{"$eq": accommodationId}}
	cursor, err := accommodationReviewCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Print("Could not get accommodation review")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var accommodationReview models.AccommodationReview
		err := cursor.Decode(&accommodationReview)

		if err != nil {
			log.Print("Could not unmarshall accommodation review on cursor")
			return nil, err
		}

		accommodationReviews = append(accommodationReviews, accommodationReview)
	}

	return accommodationReviews, nil
}

func DeleteAccommodationReview(id string) error {
	accommodationReviewId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Print("Could not get id from hex string: ", id)
	}

	filter := bson.M{"_id": bson.M{"$eq": accommodationReviewId}}
	if _, err := accommodationReviewCollection.DeleteOne(context.TODO(), filter); err != nil {
		log.Print("Could not delete accommodation review with hex id", id)
		return err
	}

	accommodationReview, err := GetAccommodationReviewById(id)
	if err := grpcclient.DeleteReview(*accommodationReview); err != nil {
		return err
	}

	return nil
}

func UpdateAccommodationReview(accommodationReview models.AccommodationReview) error {
	accommodationReview.ModifiedOn = int(time.Now().Unix())

	objId, err := primitive.ObjectIDFromHex(accommodationReview.Id.Hex())
	if err != nil {
		log.Print("Could not convert accommodation review hex to id")
		return err
	}

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.M{"$set": bson.M{"rating": accommodationReview.Rating}}

	if _, err := accommodationReviewCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Print("Could not update accommodation review")
		return err
	}

	if err := grpcclient.UpdateReview(accommodationReview); err != nil {
		return err
	}

	return nil
}
