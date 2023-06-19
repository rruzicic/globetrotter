package repos

import (
	"context"
	"log"
	"time"

	"github.com/rruzicic/globetrotter/bnb/notification-service/model"
	"github.com/rruzicic/globetrotter/bnb/notification-service/pb"
	"github.com/rruzicic/globetrotter/bnb/notification-service/socket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func GetNotificationsByUserId(id string) ([]model.Notification, error) {

	notifications := []model.Notification{}
	filter := bson.M{"user_id": bson.M{"$eq": id}}
	cursor, err := notificationCollection.Find(context.TODO(), filter)

	if err != nil {
		log.Print("Could not get notifications")
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var notification model.Notification
		err := cursor.Decode(&notification)

		if err != nil {
			log.Print("Could not unmarshall notification on cursor")
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func CreateReservationNotification(notification model.Notification) (*model.Notification, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	notification.Id = &obj_id
	notification.CreatedOn = int(time.Now().Unix())
	notification.ModifiedOn = int(time.Now().Unix())
	notification.Type = "RESERVATION"

	_, err := notificationCollection.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Print("Could not create notification! err: ", err.Error())
		return nil, err
	}
	return &notification, nil
}

func CreateCancellationNotification(notification model.Notification) (*model.Notification, error) {
	log.Println("Notification repo hit")
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	notification.Id = &obj_id
	notification.CreatedOn = int(time.Now().Unix())
	notification.ModifiedOn = int(time.Now().Unix())
	notification.Type = "CANCELLATION"

	_, err := notificationCollection.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Print("Could not create notification! err: ", err.Error())
		return nil, err
	}
	return &notification, nil
}

func CreateRatingNotification(notification model.Notification) (*model.Notification, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	notification.Id = &obj_id
	notification.CreatedOn = int(time.Now().Unix())
	notification.ModifiedOn = int(time.Now().Unix())
	notification.Type = "RATING"

	_, err := notificationCollection.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Print("Could not create notification! err: ", err.Error())
		return nil, err
	}
	return &notification, nil
}

func CreateAccommodationRatingNotification(notification model.Notification) (*model.Notification, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	notification.Id = &obj_id
	notification.CreatedOn = int(time.Now().Unix())
	notification.ModifiedOn = int(time.Now().Unix())
	notification.Type = "A_RATING"

	_, err := notificationCollection.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Print("Could not create notification! err: ", err.Error())
		return nil, err
	}
	return &notification, nil
}

func CreateHostStatusNotification(notification model.Notification) (*model.Notification, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	notification.Id = &obj_id
	notification.CreatedOn = int(time.Now().Unix())
	notification.ModifiedOn = int(time.Now().Unix())
	notification.Type = "HOST_STATUS"

	_, err := notificationCollection.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Print("Could not create notification! err: ", err.Error())
		return nil, err
	}
	return &notification, nil
}

func CreateReservationResponseNotification(notification model.Notification) (*model.Notification, error) {
	obj_id := primitive.NewObjectIDFromTimestamp(time.Now())
	notification.Id = &obj_id
	notification.CreatedOn = int(time.Now().Unix())
	notification.ModifiedOn = int(time.Now().Unix())
	notification.Type = "RESPONSE"

	_, err := notificationCollection.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Print("Could not create notification! err: ", err.Error())
		return nil, err
	}
	return &notification, nil
}

func ReservationCanceled(event *pb.ReservationEvent) error {
	log.Println("notification server hit")
	notification := model.Notification{
		UserId:            event.UserId,
		AccommodationId:   &event.AccommodationId,
		AccommodationName: &event.AccommodationName,
	}

	notif, err := CreateCancellationNotification(notification)
	if err != nil {
		log.Panic("Notification creation failed")
		return err
	}
	socket.SendNotification(*notif)

	return nil
}
