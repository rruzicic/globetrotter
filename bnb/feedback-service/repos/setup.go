package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var hostReviewCollection *mongo.Collection
var accommodationReviewCollection *mongo.Collection

var client *mongo.Client

func Connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))
	if err != nil {
		log.Panic("Could not connect to MongoDB")
	}

	hostReviewCollection = client.Database("bnb-reviews").Collection("host_reviews")
	accommodationReviewCollection = client.Database("bnb-reviews").Collection("accommodation_reviews")
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panic(err)
	}
}
