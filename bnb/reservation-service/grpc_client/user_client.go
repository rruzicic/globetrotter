package grpcclient

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/pb"
	"google.golang.org/grpc"
)

func connectToUserService() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	conn, err := grpc.Dial("account-service:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect to reservation service")
		return nil, err
	}

	return conn, nil
}

func IncrementCancellationsCounter(id string) (*pb.UserResponse, error) {
	conn, _ := connectToUserService()
	client := pb.NewUserServiceClient(conn)

	user_response, err := client.IncrementCancellationsCounter(context.Background(), &pb.UserRequestId{Id: id})
	if err != nil {
		log.Panic("Could not increment user cancelations from account service. Error: ", err)
		return nil, err
	}

	defer conn.Close()

	return user_response, nil
}
