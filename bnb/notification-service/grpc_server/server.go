package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/rruzicic/globetrotter/bnb/notification-service/model"
	"github.com/rruzicic/globetrotter/bnb/notification-service/pb"
	"github.com/rruzicic/globetrotter/bnb/notification-service/repos"
	"github.com/rruzicic/globetrotter/bnb/notification-service/socket"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationServiceServer struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *NotificationServiceServer) ReservationCreated(ctx context.Context, res *pb.ReservationNotification) (*emptypb.Empty, error){

	notification := model.Notification{
		UserId: res.UserId,
		AccommodationId: &res.AccommodationId,
		AccommodationName: &res.AccommodationName,
	}

	notif, err := repos.CreateReservationNotification(notification)
	if err != nil {
		log.Panic("Notification creation failed")
	}
	socket.SendNotification(*notif)

	return &emptypb.Empty{}, nil
}
func (s *NotificationServiceServer) ReservationCanceled(ctx context.Context, res *pb.ReservationNotification) (*emptypb.Empty, error){

	notification := model.Notification{
		UserId: res.UserId,
		AccommodationId: &res.AccommodationId,
		AccommodationName: &res.AccommodationName,
	}

	notif, err := repos.CreateCancellationNotification(notification)
	if err != nil {
		log.Panic("Notification creation failed")
	}
	socket.SendNotification(*notif)

	return &emptypb.Empty{}, nil
}

func (s *NotificationServiceServer) HostRated(ctx context.Context, rating *pb.HostRatingNotification) (*emptypb.Empty, error) {
	log.Println("Notification server hit")
	ratingValue := int(rating.Rating)
	notification := model.Notification{
		UserId: rating.RatedId,
		RaterId: &rating.RaterId,
		Rating:  &ratingValue,
	}

	notif, err := repos.CreateRatingNotification(notification);
	if err != nil {
		log.Println("Error server.go notification service")
	}
	socket.SendNotification(*notif)

	return &emptypb.Empty{}, nil
}

func (s *NotificationServiceServer) AccommodationRated(ctx context.Context, rating *pb.AccommodationRatingNotification) (*emptypb.Empty, error) {
	ratingValue := int(rating.Rating)
	notification := model.Notification{
		UserId: rating.OwnerId,
		AccommodationId: &rating.RatedId,
		RaterId: &rating.RaterId,
		Rating:  &ratingValue,
	}

	notif, err := repos.CreateAccommodationRatingNotification(notification);
	if err != nil {
		log.Println("Error server.go notification service")
	}
	socket.SendNotification(*notif)

	return &emptypb.Empty{}, nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panic("Notification service failed to listen. Error: ", err)
	}

	server := grpc.NewServer()
	pb.RegisterNotificationServiceServer(server, &NotificationServiceServer{})

	log.Println("Notification gRPC server listening..")
	if err := server.Serve(listen); err != nil {
		log.Panic("Could not start Notification gRPC Server. Error: ", err)
	}
}
