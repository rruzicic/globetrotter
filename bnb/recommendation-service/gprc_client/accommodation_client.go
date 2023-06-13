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

func connectToAccomodationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("accommodation-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panic("Could not connect to reservation service. Error: ", err.Error())
		return nil, err
	}

	return conn, nil
}

func buildLocalAccommodation(accommodation *pb.Accommodation) models.Accommodation {
	return models.Accommodation{
		Name:     accommodation.Name,
		Location: accommodation.Country + ", " + accommodation.Street + " " + accommodation.StreetNum,
		Price:    accommodation.Amount,
		MongoId:  accommodation.Id,
	}
}

func GetAllAccommodations() ([](models.Accommodation), error) {
	conn, _ := connectToAccomodationService()
	client := pb.NewAccommodationServiceClient(conn)

	accommodations := []models.Accommodation{}
	stream, err := client.GetAllAccommodations(context.Background(), &pb.Empty{})
	if err != nil {
		log.Println("Could not get stream of accomodations, Error: ", err.Error())
		return nil, err
	}

	for {
		accommodation, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error in reading from accomodations stream. Error: ", err.Error())
			return nil, err
		}
		accommodations = append(accommodations, buildLocalAccommodation(accommodation))
	}

	defer conn.Close()

	return accommodations, nil
}
