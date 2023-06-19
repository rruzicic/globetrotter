package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/feedback-service/models"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectToNotificationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("notification-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panic("Could not connect to notification service. Error: ", err.Error())
		return nil, err
	}

	return conn, nil
}

func HostRated(review models.HostReview) (*pb.UserResponse, error) {
	conn, _ := connectToNotificationService()
	client := pb.NewNotificationServiceClient(conn)

	log.Println("notification client Host Rated")

	_, err := client.HostRated(context.Background(), &pb.HostRatingNotification{RatedId: review.HostId.Hex(), RaterId: review.UserId.Hex(), Rating: int64(review.Rating)})
	if err != nil {
		log.Println("Error: notification client of feedback service")
	}
	defer conn.Close()

	return nil, nil
}
func AccommodationRated(review models.AccommodationReview, hostId string, accommodationName string) (*pb.UserResponse, error) {
	conn, _ := connectToNotificationService()
	client := pb.NewNotificationServiceClient(conn)

	_, err := client.AccommodationRated(context.Background(), &pb.AccommodationRatingNotification{OwnerId: hostId, RatedId: review.AccommodationId.Hex(), RaterId: review.UserId.Hex(), Rating: int64(review.Rating), AccommodationName: accommodationName})
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	return nil, nil
}