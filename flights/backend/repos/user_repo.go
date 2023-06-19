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
	update := bson.M{"$set": bson.M{
		"first_name":         user.FirstName,
		"last_name":          user.LastName,
		"email":              user.EMail,
		"password":           user.Password,
		"role":               user.Role,
		"api_key.key":        user.ApiKey.Key,
		"api_key.expiration": user.ApiKey.Expiration,
		"address.country":    user.Address.Country,
		"address.street":     user.Address.Street,
		"address.street_num": user.Address.StreetNum,
		"address.zip":        user.Address.ZIPCode,
	}}

	if _, err := usersCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Print("Error: ", err.Error())
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
