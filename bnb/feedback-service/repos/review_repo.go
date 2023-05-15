package repos

import (
	"context"
	"log"
	"time"

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
