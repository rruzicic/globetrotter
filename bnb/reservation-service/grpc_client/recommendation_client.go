package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func connectToRecommendationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("recommendation-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panicf("Could not connect to recommendation service")
		return nil, err
	}

	return conn, nil
}

func buildGraphReservation(reservation models.Reservation) *pb.GraphReservation {
	return &pb.GraphReservation{
		MongoId:              reservation.Id.Hex(),
		UserMongoId:          reservation.UserId.Hex(),
		AccommodationMongoId: reservation.AccommodationId.Hex(),
		ReservationEnd:       timestamppb.New(reservation.DateInterval.End),
	}
}

func CreateReservation(reservation models.Reservation) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.CreateReservation(context.Background(), buildGraphReservation(reservation))
	if err != nil {
		log.Panic("Could not create reservation relatioship. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func DeleteReservation(reservation models.Reservation) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.DeleteReservation(context.Background(), buildGraphReservation(reservation))
	if err != nil {
		log.Panic("Could not delete reservation relatioship. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func UpdateReservation(reservation models.Reservation) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.UpdateReservation(context.Background(), buildGraphReservation(reservation))
	if err != nil {
		log.Panic("Could not update reservation relatioship. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}
