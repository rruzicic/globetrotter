package repos

import (
	"context"
	"log"
	"time"

	"github.com/rruzicic/globetrotter/flights/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func FindUserByObjectId(id string) {

}

func FindUserByEmail(email string) (models.User, error) {
	result := models.User{}
	filter := bson.M{"email": bson.M{"$eq": email}}
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func FindAllUsers() []models.User {
	result := []models.User{}
	//filter := bson.D{{Name: "deleted_on", Value: bson.D{{Name: "$eq", Value: "0"}}}}
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic("could not find users err: ", err.Error())
		return []models.User{}
	}
	cursor.All(context.TODO(), &result)
	return result
}
func CreateUser(user models.User) bool {
	if _, err := usersCollection.InsertOne(context.TODO(), user); err != nil {
		log.Panic("could not save document to database! err: ", err.Error())
		return false
	}
	return true
}

func UpdateUser(user models.User) bool {
	user.ModifiedOn = int(time.Now().Unix())

	objID, err := primitive.ObjectIDFromHex(user.Id.Hex())
	if err != nil {
		log.Panic("could not convert hex to id")
		return false
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.D{{Name: "$set", Value: bson.D{}}}

	if _, err := usersCollection.UpdateByID(context.TODO(), filter, update); err != nil {
		return false
	}
	return true
}

func DeleteUser(id bson.ObjectId) {

}
