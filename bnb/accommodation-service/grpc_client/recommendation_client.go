package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/models"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectToRecommendationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("recommendation-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panicf("Could not connect to recommendation service")
		return nil, err
	}

	return conn, nil
}

func buildGraphAccommodation(accommodation models.Accommodation) *pb.GraphAccommodation {
	return &pb.GraphAccommodation{
		Name:     accommodation.Name,
		Location: accommodation.Location.Country + ", " + accommodation.Location.Street + " " + accommodation.Location.StreetNum,
		Price:    accommodation.UnitPrice.Amount,
		MongoId:  accommodation.Id.Hex(),
	}
}

func CreateAccommodation(accommodation models.Accommodation) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.CreateAccommodation(context.Background(), buildGraphAccommodation(accommodation))
	if err != nil {
		log.Panic("Could not create accommodation node. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func DeleteAccommodation(accommodation models.Accommodation) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.DeleteAccommodation(context.Background(), buildGraphAccommodation(accommodation))
	if err != nil {
		log.Panic("Could not delete accommodation node. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func UpdateAccommodation(accommodation models.Accommodation) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.UpdateAccommodation(context.Background(), buildGraphAccommodation(accommodation))
	if err != nil {
		log.Panic("Could not update accommodation node. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}
