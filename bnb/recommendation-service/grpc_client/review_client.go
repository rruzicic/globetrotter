package grpcclient

import (
	"context"
	"io"
	"log"

	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectToReviewService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("feedback-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panic("Could not connect to feedback service. Error: ", err.Error())
		return nil, err
	}

	return conn, nil
}

func buildLocalReview(review *pb.AccommodationReview) models.Review {
	return models.Review{
		Value:                int(review.Rating),
		MongoId:              review.Id,
		UserMongoId:          review.UserId,
		AccommodationMongoId: review.AccommodationId,
	}
}

func GetAllReviews() ([]models.Review, error) {
	conn, _ := connectToReviewService()
	client := pb.NewFeedbackServiceClient(conn)

	reviews := []models.Review{}
	stream, err := client.GetAllAccommodationReviews(context.Background(), &pb.EmptyReviewMsg{})
	if err != nil {
		log.Println("Could not get stream of reviews, Error: ", err.Error())
		return nil, err
	}

	for {
		review, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error in reading from review stream. Error: ", err.Error())
			return nil, err
		}

		reviews = append(reviews, buildLocalReview(review))
	}

	return reviews, nil
}
