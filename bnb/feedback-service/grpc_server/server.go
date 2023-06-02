package grpcserver

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/rruzicic/globetrotter/bnb/feedback-service/models"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/pb"

	//"github.com/rruzicic/globetrotter/bnb/feedback-service/repos"
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
