package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/pb"
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

func ReservationCreated(res models.Reservation) (*pb.UserResponse, error) {
	conn, _ := connectToNotificationService()
	client := pb.NewNotificationServiceClient(conn)

	_, err := client.ReservationCreated(context.Background(), &pb.ReservationNotification{})
	if err != nil {
		log.Panic("Could not notify about reservation. Error: ", err)
		return nil, err
	}
	
	defer conn.Close()

	return nil, nil
}