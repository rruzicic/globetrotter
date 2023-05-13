package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var acommodationsCollection *mongo.Collection

var client *mongo.Client

func Connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))
	if err != nil {
		log.Panic("Could not connect to MongoDB")
	}
	acommodationsCollection = client.Database("bnb-accommodations").Collection("accommodations")

	// TODO: add uniqueness to needed fields
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panic(err)
	}
}
