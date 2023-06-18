package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/pb"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/repos"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type RecommendationServiceDBEventsServer struct {
	pb.UnimplementedRecommendationServiceDBEventsServer
}

func buildLocalAccommodation(accommodation *pb.GraphAccommodation) models.Accommodation {
	return models.Accommodation{
		Name:     accommodation.Name,
		Location: accommodation.Location,
		Price:    accommodation.Price,
		MongoId:  accommodation.MongoId,
	}
}

func buildLocalUser(user *pb.GraphUser) models.User {
	return models.User{
		Name:    user.Name,
		MongoId: user.MongoId,
	}
}

func buildLocalReservation(reservation *pb.GraphReservation) models.Reservation {
	return models.Reservation{
		MongoId:              reservation.MongoId,
		UserMongoId:          reservation.UserMongoId,
		AccommodationMongoId: reservation.AccommodationMongoId,
		ReservationEnd:       reservation.ReservationEnd.AsTime(),
	}
}

func buildLocalReview(review *pb.GraphReview) models.Review {
	return models.Review{
		Value:                int(review.Value),
		MongoId:              review.MongoId,
		UserMongoId:          review.UserMongoId,
		AccommodationMongoId: review.AccommodationMongoId,
	}
}

func (s *RecommendationServiceDBEventsServer) CreateAccommodation(ctx context.Context, req *pb.GraphAccommodation) (*pb.GraphEmpty, error) {
	if err := repos.CreateAccommodationNode(buildLocalAccommodation(req)); err != nil {
		log.Print("Could not create accommodation node. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) CreateUser(ctx context.Context, req *pb.GraphUser) (*pb.GraphEmpty, error) {
	if err := repos.CreateUserNode(buildLocalUser(req)); err != nil {
		log.Print("Could not create user node. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) CreateReservation(ctx context.Context, req *pb.GraphReservation) (*pb.GraphEmpty, error) {
	if err := repos.CreateReservationRelationship(buildLocalReservation(req)); err != nil {
		log.Print("Could not create reservation relationship. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) CreateReview(ctx context.Context, req *pb.GraphReview) (*pb.GraphEmpty, error) {
	if err := repos.CreateReviewRelationship(buildLocalReview(req)); err != nil {
		log.Print("Could not create review relationship. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) DeleteAccommodation(ctx context.Context, req *pb.GraphAccommodation) (*pb.GraphEmpty, error) {
	if err := repos.DeleteAccommodationNode(buildLocalAccommodation(req)); err != nil {
		log.Print("Could not delete accommodation node. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) DeleteUser(ctx context.Context, req *pb.GraphUser) (*pb.GraphEmpty, error) {
	if err := repos.DeleteUserNode(buildLocalUser(req)); err != nil {
		log.Print("Could not delete user node. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) DeleteReservation(ctx context.Context, req *pb.GraphReservation) (*pb.GraphEmpty, error) {
	if err := repos.DeleteReservationRelationship(buildLocalReservation(req)); err != nil {
		log.Print("Could not delete reservation relationship. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) DeleteReview(ctx context.Context, req *pb.GraphReview) (*pb.GraphEmpty, error) {
	if err := repos.DeleteReviewRelationship(buildLocalReview(req)); err != nil {
		log.Print("Could not delete review relationship. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) UpdateAccommodation(ctx context.Context, req *pb.GraphAccommodation) (*pb.GraphEmpty, error) {
	if err := repos.UpdateAccommodationNode(buildLocalAccommodation(req)); err != nil {
		log.Print("Could not update accommodation node. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) UpdateUser(ctx context.Context, req *pb.GraphUser) (*pb.GraphEmpty, error) {
	if err := repos.UpdateUserNode(buildLocalUser(req)); err != nil {
		log.Print("Could not update user node. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) UpdateReservation(ctx context.Context, req *pb.GraphReservation) (*pb.GraphEmpty, error) {
	if err := repos.UpdateReservationRelationship(buildLocalReservation(req)); err != nil {
		log.Print("Could not update reservation relationship. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func (s *RecommendationServiceDBEventsServer) UpdateReview(ctx context.Context, req *pb.GraphReview) (*pb.GraphEmpty, error) {
	if err := repos.UpdateReviewRelationship(buildLocalReview(req)); err != nil {
		log.Print("Could not update review relationship. Error: ", err.Error())
		return nil, err
	}

	return &pb.GraphEmpty{}, nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Recommendation service failed to listen. Error: ", err.Error())
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()), grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()))
	pb.RegisterRecommendationServiceDBEventsServer(server, &RecommendationServiceDBEventsServer{})

	log.Println("Recommendation server gRPC server listening..")
	if err := server.Serve(listen); err != nil {
		log.Println("Could not start Recommendation gRPC Server. Error: ", err.Error())
	}
}
