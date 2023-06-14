package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/feedback-service/models"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/pb"
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

func buildGraphReview(review models.AccommodationReview) *pb.GraphReview {
	return &pb.GraphReview{
		Value:                int32(review.Rating),
		MongoId:              review.Id.Hex(),
		UserMongoId:          review.UserId.Hex(),
		AccommodationMongoId: review.AccommodationId.Hex(),
	}
}

func CreateReview(review models.AccommodationReview) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.CreateReview(context.Background(), buildGraphReview(review))
	if err != nil {
		log.Panic("Could not create review relatioship. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func DeleteReview(review models.AccommodationReview) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.DeleteReview(context.Background(), buildGraphReview(review))
	if err != nil {
		log.Panic("Could not delete review relatioship. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}

func UpdateReview(review models.AccommodationReview) error {
	conn, _ := connectToRecommendationService()
	client := pb.NewRecommendationServiceDBEventsClient(conn)

	_, err := client.UpdateReview(context.Background(), buildGraphReview(review))
	if err != nil {
		log.Panic("Could not update review relatioship. Error: ", err.Error())
		return err
	}

	defer conn.Close()

	return nil
}
