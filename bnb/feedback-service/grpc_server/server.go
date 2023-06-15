package grpcserver

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/rruzicic/globetrotter/bnb/feedback-service/models"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/pb"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/repos"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FeedbackServiceServer struct {
	pb.UnimplementedFeedbackServiceServer
}

func buildGRPCHostReview(hostReview models.HostReview) pb.HostReview {
	return pb.HostReview{
		Id:         hostReview.Id.String(),
		CreatedOn:  timestamppb.New(time.Unix(int64(hostReview.CreatedOn), 0)),
		ModifiedOn: timestamppb.New(time.Unix(int64(hostReview.ModifiedOn), 0)),
		DeletedOn:  timestamppb.New(time.Unix(int64(hostReview.DeletedOn), 0)),
		Rating:     int32(hostReview.Rating),
		UserId:     hostReview.UserId.String(),
		HostId:     hostReview.HostId.String(),
	}
}

func buildGRPCAccommodationReview(accommodationReview models.AccommodationReview) pb.AccommodationReview {
	return pb.AccommodationReview{
		Id:              accommodationReview.Id.String(),
		CreatedOn:       timestamppb.New(time.Unix(int64(accommodationReview.CreatedOn), 0)),
		ModifiedOn:      timestamppb.New(time.Unix(int64(accommodationReview.ModifiedOn), 0)),
		DeletedOn:       timestamppb.New(time.Unix(int64(accommodationReview.DeletedOn), 0)),
		Rating:          int32(accommodationReview.Rating),
		UserId:          accommodationReview.UserId.String(),
		AccommodationId: accommodationReview.AccommodationId.String(),
	}
}

func buildGRPCAvgRating(avgRating float32) pb.AvgRatingResponse {
	return pb.AvgRatingResponse{
		AvgRating: float32(avgRating),
	}
}

func (s *FeedbackServiceServer) GetAllAccommodationReviews(req *pb.EmptyReviewMsg, stream pb.FeedbackService_GetAllAccommodationReviewsServer) error {
	reviews, err := repos.GetAllAccommodationReviews()
	if err != nil {
		log.Println("Could not get all accommodation reviews. Error: ", err.Error())
		return err
	}

	for _, review := range reviews {
		grpc_review := buildGRPCAccommodationReview(review)
		if err := stream.Send(&grpc_review); err != nil {
			return err
		}
	}

	return nil
}

func (s *FeedbackServiceServer) GetHostReviewById(ctx context.Context, req *pb.RequestReviewById) (*pb.HostReview, error) {
	hostReview, err := repos.GetHostReviewById(req.GetId())
	if err != nil {
		log.Println("Could not get host review with id", req.GetId(), ", error: ", err.Error())
		return nil, err
	}

	grpc_hostReview := buildGRPCHostReview(*hostReview)

	return &grpc_hostReview, nil
}

func (s *FeedbackServiceServer) GetHostReviewsByUserId(req *pb.RequestReviewsByUserId, stream pb.FeedbackService_GetHostReviewsByUserIdServer) error {
	hostReviews, err := repos.GetHostReviewsByUserId(req.GetId())
	if err != nil {
		log.Println("Could not get host reviews for user id", req.GetId(), ", error: ", err.Error())
		return err
	}

	for _, review := range hostReviews {
		grpc_hostReview := buildGRPCHostReview(review)
		if err := stream.Send(&grpc_hostReview); err != nil {
			return err
		}
	}

	return nil
}

func (s *FeedbackServiceServer) GetHostReviewsByHostId(req *pb.RequestReviewsByHostId, stream pb.FeedbackService_GetHostReviewsByHostIdServer) error {
	hostReviews, err := repos.GetHostReviewsByHostId(req.GetId())
	if err != nil {
		log.Println("Could not get host reviews for host id", req.GetId(), ", error: ", err.Error())
		return err
	}

	for _, review := range hostReviews {
		grpc_hostReview := buildGRPCHostReview(review)
		if err := stream.Send(&grpc_hostReview); err != nil {
			return err
		}
	}

	return nil
}

func (s *FeedbackServiceServer) CalcAvgRatingForHost(ctx context.Context, req *pb.RequestAvgRating) (*pb.AvgRatingResponse, error) {
	hostReviews, err := repos.GetHostReviewsByHostId(req.GetId())
	if err != nil {
		log.Println("Could not get host reviews for host id", req.GetId(), ", error: ", err.Error())
		return nil, err
	}

	var sumRating float32 = 0
	for _, review := range hostReviews {
		sumRating += float32(review.Rating)
	}
	var avgRating float32 = sumRating / float32(len(hostReviews))

	grpc_avgRating := buildGRPCAvgRating(avgRating)

	return &grpc_avgRating, nil
}

//============================================================================================
//============================================================================================

func (s *FeedbackServiceServer) GetAccommodationReviewById(ctx context.Context, req *pb.RequestReviewById) (*pb.AccommodationReview, error) {
	accommodationReview, err := repos.GetAccommodationReviewById(req.GetId())
	if err != nil {
		log.Println("Could not get accommodation review with id", req.GetId(), ", error: ", err.Error())
		return nil, err
	}

	grpc_accommodationReview := buildGRPCAccommodationReview(*accommodationReview)

	return &grpc_accommodationReview, nil
}

func (s *FeedbackServiceServer) GetaccommodationReviewsByUserId(req *pb.RequestReviewsByUserId, stream pb.FeedbackService_GetAccommodationReviewsByUserIdServer) error {
	accommodationReviews, err := repos.GetAccommodationReviewsByUserId(req.GetId())
	if err != nil {
		log.Println("Could not get accommodation reviews for user id", req.GetId(), ", error: ", err.Error())
		return err
	}

	for _, review := range accommodationReviews {
		grpc_accommodationReview := buildGRPCAccommodationReview(review)
		if err := stream.Send(&grpc_accommodationReview); err != nil {
			return err
		}
	}

	return nil
}

func (s *FeedbackServiceServer) GetAccommodationReviewsByAccommodationId(req *pb.RequestReviewsByAccommodationId, stream pb.FeedbackService_GetAccommodationReviewsByAccommodationIdServer) error {
	accommodationReviews, err := repos.GetAccommodationReviewsByAccommodationId(req.GetId())
	if err != nil {
		log.Println("Could not get accommodation reviews for accommodation id", req.GetId(), ", error: ", err.Error())
		return err
	}

	for _, review := range accommodationReviews {
		grpc_accommodationReview := buildGRPCAccommodationReview(review)
		if err := stream.Send(&grpc_accommodationReview); err != nil {
			return err
		}
	}

	return nil
}

func (s *FeedbackServiceServer) CalcAvgRatingForAccommodation(ctx context.Context, req *pb.RequestAvgRating) (*pb.AvgRatingResponse, error) {
	accommodationReviews, err := repos.GetAccommodationReviewsByAccommodationId(req.GetId())
	if err != nil {
		log.Println("Could not get accommodation reviews for accommodation id", req.GetId(), ", error: ", err.Error())
		return nil, err
	}

	var sumRating float32 = 0
	for _, review := range accommodationReviews {
		sumRating += float32(review.Rating)
	}
	var avgRating float32 = sumRating / float32(len(accommodationReviews))

	grpc_avgRating := buildGRPCAvgRating(avgRating)

	return &grpc_avgRating, nil
}

func (s *FeedbackServiceServer) TestConnection(ctx context.Context, req *pb.TestMessage) (*pb.TestMessage, error) {
	log.Print("Hello from feedback service, message is: ", req.GetMsg())
	return req, nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Feedback service failed to listen. Error: ", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterFeedbackServiceServer(server, &FeedbackServiceServer{})

	log.Println("Feedback gRPC server listening at ", listen.Addr())
	if err := server.Serve(listen); err != nil {
		log.Println("Could not start Feedback gRPC Server. Error: ", err.Error())
	}
}
