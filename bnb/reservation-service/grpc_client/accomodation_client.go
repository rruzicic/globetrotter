package grpcclient

import (
	"context"
	"io"
	"log"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/pb"
	"google.golang.org/grpc"
)

func connectToAccomodationService() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	conn, err := grpc.Dial("accomodation-service:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect to reservation service")
		return nil, err
	}

	return conn, nil
}

func GetAccommodationById(id string) (*pb.Accommodation, error) {
	conn, _ := connectToAccomodationService()
	client := pb.NewAccommodationServiceClient(conn)

	accommodation, err := client.GetAccommodationById(context.Background(), &pb.RequestAccommodationById{Id: id})
	if err != nil {
		log.Panic("Could not get accommodation by id from accommodation service. Error: ", err)
		return nil, err
	}

	return accommodation, err
}

func GetAccommodationByHostId(id string) ([](*pb.Accommodation), error) {
	conn, _ := connectToAccomodationService()
	client := pb.NewAccommodationServiceClient(conn)

	var accomodations [](*pb.Accommodation)
	stream, err := client.GetAccommodationByHostId(context.Background(), &pb.RequestAccommodationByHostId{Id: id})
	if err != nil {
		log.Println("Could not get stream for accomodations by host id, error: ", err.Error())
		return nil, err
	}

	for {
		accomodation, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error in reading from accomodations stream. Error: ", err.Error())
			return nil, err
		}
		accomodations = append(accomodations, accomodation)
	}

	defer conn.Close()

	return accomodations, nil
}
