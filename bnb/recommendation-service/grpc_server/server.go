package grpcserver

import (
	"context"

	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/pb"
)

type RecommendationServiceDBEventsServer struct {
	pb.UnimplementedRecommendationServiceDBEventsServer
}

func buildLocalAccommodation(accommodation *pb.Accommodation) models.Accommodation {
	return models.Accommodation{
		Name:     accommodation.Name,
		Location: accommodation.Location,
		Price:    accommodation.Price,
		MongoId:  accommodation.MongoId,
	}
}

func buildLocalUser(user *pb.User) models.User {
	return models.User{
		Name:    user.Name,
		MongoId: user.MongoId,
	}
}

func buildLocalReservation(reservation *pb.Reservation) models.Reservation {
	return models.Reservation{
		MongoId:              reservation.MongoId,
		UserMongoId:          reservation.UserMongoId,
		AccommodationMongoId: reservation.AccommodationMongoId,
		ReservationEnd:       reservation.ReservationEnd.AsTime(),
	}
}

func buildLocalReview(review *pb.Review) models.Review {
	return models.Review{
		Value:                int(review.Value),
		MongoId:              review.MongoId,
		UserMongoId:          review.UserMongoId,
		AccommodationMongoId: review.AccommodationMongoId,
	}
}

func (s *RecommendationServiceDBEventsServer) GetAllAccommodations(pb.RecommendationServiceDBEvents_GetAllAccommodationsServer) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) GetAllUsers(pb.RecommendationServiceDBEvents_GetAllUsersServer) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) GetAllReservations(pb.RecommendationServiceDBEvents_GetAllReservationsServer) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) GetAllReviews(pb.RecommendationServiceDBEvents_GetAllReviewsServer) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) CreateAccommodation(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) CreateUser(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) CreateReservation(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) CreateReview(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) DeleteAccommodation(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) DeleteUser(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) DeleteReservation(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) DeleteReview(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) UpdateAccommodation(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) UpdateUser(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) UpdateReservation(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
func (s *RecommendationServiceDBEventsServer) UpdateReview(ctx context.Context, req *pb.Accommodation) (*pb.Empty, error) {
}
