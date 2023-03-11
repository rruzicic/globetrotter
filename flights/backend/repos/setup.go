package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_URI = "mongodb://mongo"

var UsersCollection *mongo.Collection
var FlightsCollection *mongo.Collection
var TicketsCollection *mongo.Collection

var client *mongo.Client

func Setup() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))

	if err != nil {
		log.Panic(err)
	}

	UsersCollection = client.Database("flights").Collection("users")
	FlightsCollection = client.Database("flights").Collection("flights")
	TicketsCollection = client.Database("flights").Collection("tickets")
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panic(err)
	}
}
