package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/account-service/models"
	"github.com/rruzicic/globetrotter/bnb/account-service/pb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectToRecommendationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("recommendation-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Panicf("Could not connect to recommendation service")
		return nil, err
	}

	return conn, nil
}

func buildGraphUser(user models.User) *pb.GraphUser {
	return &pb.GraphUser{
		Name:    user.FirstName + " " + user.LastName,
		MongoId: user.Id.Hex(),
	}
}

func CreateUser(user models.User) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.CreateUser(context.Background(), buildGraphUser(user))
	if err != nil {
		log.Panic("Could not create user node. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func DeleteUser(user models.User) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.DeleteUser(context.Background(), buildGraphUser(user))
	if err != nil {
		log.Panic("Could not delete user node. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func UpdateUser(user models.User) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.UpdateUser(context.Background(), buildGraphUser(user))
	if err != nil {
		log.Panic("Could not update user node. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}
