package repos

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_URI = "mongodb://mongo"

var UsersCollection *mongo.Collection

var client *mongo.Client

func Setup() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))

	if err != nil {
		panic(err)
	}

	UsersCollection = client.Database("flights").Collection("users")
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
