package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const DATABASE_URI = "mongodb://mongo"

var usersCollection *mongo.Collection

var client *mongo.Client

func Setup() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo"))

	if err != nil {
		log.Panic(err)
	}

	usersCollection = client.Database("flights").Collection("users")

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
	if _, err := usersCollection.Indexes().CreateOne(context.TODO(), index); err != nil {
		log.Println("couldn't set email field to unique")
	}
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panic(err)
	}
}
