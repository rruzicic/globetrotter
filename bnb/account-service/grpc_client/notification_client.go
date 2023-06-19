package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/account-service/pb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HostStatusChanged(id string) (*pb.UserResponse, error) {
	conn, _ := connectToNotificationService()
	client := pb.NewNotificationServiceClient(conn)

	_, err := client.HostStatusChanged(context.Background(),
	&pb.HostStatusNotification{
		UserId:            id,
	})
	if err != nil {
		log.Panic("Could not notify about host status. Error: ", err)
		return nil, err
	}

	defer conn.Close()

	return nil, nil
}

func connectToNotificationService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("notification-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Panicf("Could not connect to notification service")
		return nil, err
	}

	return conn, nil
}