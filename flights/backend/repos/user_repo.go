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

func FindAllUsers() []models.User {
	result := []models.User{}
	cursor, err := usersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Panic("could not find users err: ", err.Error())
		return []models.User{}
	}
	cursor.All(context.TODO(), &result)
	return result
}
func CreateUser(user models.User) bool {
	user.CreatedOn = int(time.Now().Unix())
	user.ModifiedOn = int(time.Now().Unix())

	if _, err := FindUserByEmail(user.EMail); err == nil {
		return false
	}
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

func FindUserByEmail(mail string) (*models.User, error) {
	user := models.User{}
	filter := bson.M{"email": bson.M{"$eq": mail}}

	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByAPIKey(api_key string) (*models.User, error) {
	var user models.User

	filter := bson.M{"api_key": bson.M{"value": bson.M{"$eq": api_key}}}

	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		log.Panic("Could not find user by api key: ", api_key)
		return nil, err
	}

	return &user, nil
}

func AddUserAPIKey(user models.User, api_key models.API_Key) bool {
	user.APIKey = api_key
	return UpdateUser(user)
}

func DeleteUserAPIKey(api_key string) bool {
	user, err := FindUserByAPIKey(api_key)

	if err != nil {
		return false
	}

	user.APIKey = models.API_Key{}

	return UpdateUser(*user)
}
