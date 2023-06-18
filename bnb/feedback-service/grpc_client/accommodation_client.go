package grpcclient

import (
	"context"
	"io"
	"log"

	"github.com/rruzicic/globetrotter/bnb/feedback-service/pb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Helper function
func contains(list [](*pb.HostAnswer), item *pb.HostAnswer) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}

func GetAccommodationById(id string) (*pb.Accommodation, error) {
	conn, _ := connectToAccommodationService()
	client := pb.NewAccommodationServiceClient(conn)

	accommodation, err := client.GetAccommodationById(context.Background(), &pb.RequestAccommodationById{Id: id})
	if err != nil {
		log.Panic("Could not get accommodation by id from accommodation service. Error: ", err)
		return nil, err
	}

	defer conn.Close()

	return accommodation, err
}

func connectToAccommodationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("accommodation-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
	grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
)

	if err != nil {
		log.Panic("Could not connect to accommodation service. Error: ", err.Error())
		return nil, err
	}

	return conn, nil
}

func GetPastHostsByAccommodations(accommodationIds []string) ([](*pb.HostAnswer), error) {
	conn, _ := connectToAccommodationService()
	client := pb.NewAccommodationServiceClient(conn)

	var pastHosts [](*pb.HostAnswer)
	stream, err := client.GetPastHostsByAccommodations(context.Background(), &pb.RequestGetPastHostsByAccommodations{AccommodationId: accommodationIds})
	if err != nil {
		log.Panic("Could not get stream for past hosts by user")
		return nil, err
	}

	for {
		pastHost, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Error in reading from reservations stream. Error: ", err)
			return nil, err
		}
		if !contains(pastHosts, pastHost) {
			pastHosts = append(pastHosts, pastHost)
		}
	}

	defer conn.Close()

	return pastHosts, nil
}
