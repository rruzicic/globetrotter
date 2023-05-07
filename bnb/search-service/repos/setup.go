package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var accomodationsCollection *mongo.Collection

var client *mongo.Client

func Connect() {
	// TODO: add uniqueness like in user service if needed
	// TODO: add read-only param somehwere

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))
	if err != nil {
		log.Panic("Could not connect to MongoDB")
	}

	accomodationsCollection = client.Database("bnb-accommodations").Collection("accommodations")
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panic(err)
	}
}
