package grpcclient

import (
	"context"
	"io"
	"log"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectToReservationService() (*grpc.ClientConn, error) {
	// var opts []grpc.DialOption
	conn, err := grpc.Dial("reservation-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to reservation service")
		return nil, err
	}

	return conn, nil
}

func GetReservationById(id string) (*pb.Reservation, error) {
	conn, _ := connectToReservationService()
	client := pb.NewReservationServiceClient(conn)

	reservation, err := client.GetReservationById(context.Background(), &pb.RequestReservationById{Id: id})
	if err != nil {
		log.Panic("Could not get reservation by id from reservation grpc service")
		return nil, err
	}

	defer conn.Close()

	return reservation, nil
}

func GetReservationsByAccommodationId(id string) ([](*pb.Reservation), error) {
	conn, _ := connectToReservationService()
	client := pb.NewReservationServiceClient(conn)

	var reservations [](*pb.Reservation)
	stream, err := client.GetReservationsByAccommodationId(context.Background(), &pb.RequestReservationsByAccommodationId{Id: id})
	if err != nil {
		log.Panic("Could not get stream for reservations by accommodation id")
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
