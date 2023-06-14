package gprcclient

import (
	"context"
	"io"
	"log"

	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectToReservationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("reservation-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panic("Could not connect to reservation service. Error: ", err.Error())
		return nil, err
	}

	return conn, nil
}

func buildLocalReservation(reservation *pb.Reservation) models.Reservation {
	return models.Reservation{
		MongoId:              reservation.Id,
		UserMongoId:          reservation.UserId,
		AccommodationMongoId: reservation.AccommodationId,
		ReservationEnd:       reservation.EndDate.AsTime(),
	}
}

func GetAllReservations() ([]models.Reservation, error) {
	conn, _ := connectToReservationService()
	client := pb.NewReservationServiceClient(conn)

	reservations := []models.Reservation{}
	stream, err := client.GetAllReservations(context.Background(), &pb.EmptyResMsg{})
	if err != nil {
		log.Println("Could not get stream of accomodations, Error: ", err.Error())
		return nil, err
	}

	for {
		reservation, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error in reading from reservations stream. Error: ", err.Error())
			return nil, err
		}
		reservations = append(reservations, buildLocalReservation(reservation))
	}

	defer conn.Close()

	return reservations, nil
}
