package grpcclient

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/timestamp"
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

func ReservationCreated(reservation models.Reservation, accommodationName string, hostId string) (*pb.UserResponse, error) {
	conn, _ := connectToNotificationService()
	client := pb.NewNotificationServiceClient(conn)

	_, err := client.ReservationCreated(context.Background(),
	&pb.ReservationNotification{
		AccommodationId: reservation.AccommodationId.Hex(),
		UserId: hostId,
		StartDate: &timestamp.Timestamp{Seconds: reservation.DateInterval.Start.Unix()},
		EndDate: &timestamp.Timestamp{Seconds: reservation.DateInterval.End.Unix()},
		NumOfGuests: int32(reservation.NumOfGuests),
		IsApproved: reservation.IsApproved,
		AccommodationName: accommodationName,
	})
	if err != nil {
		log.Panic("Could not notify about reservation. Error: ", err)
		return nil, err
	}
	
	defer conn.Close()

	return nil, nil
}
func ReservationCanceled(reservation models.Reservation, accommodationName string, hostId string) (*pb.UserResponse, error) {
	conn, _ := connectToNotificationService()
	client := pb.NewNotificationServiceClient(conn)

	_, err := client.ReservationCanceled(context.Background(),
	&pb.ReservationNotification{
		AccommodationId: reservation.AccommodationId.Hex(),
		UserId: hostId,
		StartDate: &timestamp.Timestamp{Seconds: reservation.DateInterval.Start.Unix()},
		EndDate: &timestamp.Timestamp{Seconds: reservation.DateInterval.End.Unix()},
		NumOfGuests: int32(reservation.NumOfGuests),
		IsApproved: reservation.IsApproved,
		AccommodationName: accommodationName,
	})
	if err != nil {
		log.Panic("Could not notify about reservation cancellation. Error: ", err)
		return nil, err
	}
	
	defer conn.Close()

	return nil, nil
}