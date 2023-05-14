package grpcclient

import (
	"context"
	"io"
	"log"

	"github.com/rruzicic/globetrotter/bnb/account-service/pb"
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

func GetActiveReservationsByUser(id string) ([](*pb.Reservation), error) {
	conn, _ := connectToReservationService()
	client := pb.NewReservationServiceClient(conn)

	var reservations [](*pb.Reservation)
	stream, err := client.GetActiveReservationsByUser(context.Background(), &pb.RequestUserId{Id: id})
	if err != nil {
		//log.Println("Could not get stream for reservations by user id, error: ", err.Error())
		return []*pb.Reservation{}, nil
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
		reservations = append(reservations, reservation)
	}

	defer conn.Close()

	return reservations, nil
}

func GetFutureActiveReservationsByHost(id string) ([](*pb.Reservation), error) {
	conn, _ := connectToReservationService()
	client := pb.NewReservationServiceClient(conn)

	var reservations [](*pb.Reservation)
	stream, err := client.GetFutureActiveReservationsByHost(context.Background(), &pb.RequestUserId{Id: id})
	if err != nil {
		//log.Println("Could not get stream for reservations by user id, error: ", err.Error())
		return []*pb.Reservation{}, err
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
		reservations = append(reservations, reservation)
	}

	defer conn.Close()

	return reservations, nil
}
