package services

import (
	"context"
	"fmt"

	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/repos"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(user models.User) {
	//user.CreatedOn = int(time.Now().Unix())
	//_, err :=
	if _, err := repos.UsersCollection.InsertOne(context.TODO(), user); err != nil {
		fmt.Println("could not save document to database!")
		fmt.Println(err.Error())
	}

}

func GetAllUsers() []models.User {
	result := []models.User{}
	cursor, _ := repos.UsersCollection.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var elem models.User
		err := cursor.Decode(&elem)
		if err != nil {
			return nil
		}
		result = append(result, elem)
	}
	return result
}
