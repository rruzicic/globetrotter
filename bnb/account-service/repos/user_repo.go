package repos

import (
	"context"
	"log"
	"time"

	grpcclient "github.com/rruzicic/globetrotter/bnb/account-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/account-service/models"
	"github.com/rruzicic/globetrotter/bnb/account-service/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(user models.User) (*models.User, error) {
	user.CreatedOn = int(time.Now().Unix())
	user.ModifiedOn = int(time.Now().Unix())
	user.WantedNotifications = []string{"RESERVATION", "CANCELLATION", "RATING", "A_RATING", "HOST_STATUS", "RESPONSE"}

	inserted_id, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Panic("could not save document to database! err: ", err.Error())
		return &models.User{}, err
	}
	id := inserted_id.InsertedID.(primitive.ObjectID)
	user.Id = &id

	if err := grpcclient.CreateUser(user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() []models.User {
	result := []models.User{}
	cursor, err := usersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Panic("could not find users err: ", err.Error())
		return []models.User{}
	}
	cursor.All(context.TODO(), &result)
	return result
}

func GetUserByEmail(mail string) (*models.User, error) {
	user := models.User{}
	filter := bson.M{"email": bson.M{"$eq": mail}}

	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(id primitive.ObjectID) (*models.User, error) {
	user := models.User{}

	filter := bson.M{"_id": bson.M{"$eq": id}}
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user models.User) (*models.User, error) {
	user.ModifiedOn = int(time.Now().Unix())

	objID, err := primitive.ObjectIDFromHex(user.Id.Hex())
	if err != nil {
		log.Println("could not convert hex to id")
		return &models.User{}, err
	}
	edited := bson.M{
		"first_name":                   user.FirstName,
		"last_name":                    user.LastName,
		"email":                        user.EMail,
		"password":                     user.Password,
		"address":                      user.Address,
		"super_host":                   user.SuperHost,
		"reservation_counter":          user.ReservationCounter,
		"canceled_reservation_counter": user.CanceledReservationsCounter,
		"total_reservation_duration":   user.TotalReservationDuration,
		"api_key":                      user.ApiKey,
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": edited}
	var updatedUser models.User
	err = usersCollection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedUser)
	if err != nil {
		return &models.User{}, err
	}

	if err := grpcclient.UpdateUser(user); err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func DeleteUser(id primitive.ObjectID) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	if _, err := usersCollection.DeleteOne(context.TODO(), filter); err != nil {
		log.Print("Could not delete user with hex id", id)
		return err
	}

	user, _ := GetUserById(id)
	if err := grpcclient.DeleteUser(*user); err != nil {
		return err
	}

	return nil
}

func AvgRatingChanged(hostId string, avgRating float32) error {
	objectId, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return err
	}
	user, err := GetUserById(objectId)
	if err != nil {
		return err
	}
	log.Println("USAO U AVGRATINGCHANGED: ", avgRating)
	user.Rating = avgRating
	err = CheckSuperHostStatus(user)
	if err != nil {
		return err
	}

	return nil
}

func HandleNewReservationEvent(reservationEvent *pb.ReservationEvent) error {
	objectId, err := primitive.ObjectIDFromHex(reservationEvent.HostId)
	if err != nil {
		return err
	}
	host, err := GetUserById(objectId)
	if err != nil {
		return err
	}

	host.ReservationCounter = host.ReservationCounter + 1
	durationOfReservation := reservationEvent.EndDate.AsTime().Sub(reservationEvent.StartDate.AsTime())
	host.TotalReservationDuration = host.TotalReservationDuration + durationOfReservation

	err = CheckSuperHostStatus(host)
	if err != nil {
		return err
	}

	return nil
}

func HandleCanceledReservationEvent(reservationEvent *pb.ReservationEvent) error {
	objectId, err := primitive.ObjectIDFromHex(reservationEvent.HostId)
	if err != nil {
		return err
	}
	host, err := GetUserById(objectId)
	if err != nil {
		return err
	}

	host.CanceledReservationsCounter = host.CanceledReservationsCounter + 1

	err = CheckSuperHostStatus(host)
	if err != nil {
		return err
	}

	return nil
}

func HandleRollbackCancleReservation(reservationEvent *pb.ReservationEvent) error {
	objectId, err := primitive.ObjectIDFromHex(reservationEvent.HostId)
	if err != nil {
		return err
	}
	host, err := GetUserById(objectId)
	if err != nil {
		return err
	}

	host.CanceledReservationsCounter = host.CanceledReservationsCounter - 1

	err = CheckSuperHostStatus(host)
	if err != nil {
		return err
	}

	return nil
}

func CheckSuperHostStatus(user *models.User) error {
	hostStatusStart := user.SuperHost
	condition1 := false
	condition2 := false
	condition3 := false
	condition4 := false
	if user.Rating > 4.7 {
		condition1 = true
	}
	if float32(user.CanceledReservationsCounter)/float32(user.ReservationCounter) <= 0.05 {
		condition2 = true
	}
	if user.ReservationCounter >= 5 {
		condition3 = true
	}
	if user.TotalReservationDuration.Hours() >= (24 * 50) {
		condition4 = true
	}
	//TODO add other conditions

	if condition1 && condition2 && condition3 && condition4 {
		user.SuperHost = true
	} else {
		user.SuperHost = false
	}
	hostStatusEnd := user.SuperHost
	if(hostStatusStart != hostStatusEnd) {
		grpcclient.HostStatusChanged(user.Id.Hex())
	}

	_, err := UpdateUser(*user)
	if err != nil {
		return err
	}

	return nil
}

func AddAPIKeyToUser(email string, key string) (*models.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		log.Print("Could not get user by email. Error: ", err.Error())
		return nil, err
	}

	user.ApiKey = key

	return UpdateUser(*user)
}
