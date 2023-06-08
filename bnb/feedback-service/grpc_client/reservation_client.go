package grpcclient

import (
	"context"
	"io"
	"log"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/pb"
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

func GetFinishedReservationsByUser(userId string) ([](*pb.Reservation), error) {
	conn, _ := connectToReservationService()
	client := pb.NewReservationServiceClient(conn)

	var reservations [](*pb.Reservation)
	stream, err := client.GetFinishedReservationsByUser(context.Background(), &pb.RequestUserId{Id: userId})
	if err != nil {
		log.Panic("Could not get stream for finished reservations by user")
		return nil, err
	}

	for {
		reservation, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Error in reading from reservations stream. Error: ", err)
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	defer conn.Close()

	return reservations, nil
}
