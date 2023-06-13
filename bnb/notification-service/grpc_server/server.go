package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/rruzicic/globetrotter/bnb/notification-service/pb"
	"github.com/rruzicic/globetrotter/bnb/notification-service/socket"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationServiceServer struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *NotificationServiceServer) ReservationCreated(ctx context.Context, res *pb.Reservation) (*emptypb.Empty, error){
	socket.SendNotification("This is the title", "This is the message", "user@email.com")

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
