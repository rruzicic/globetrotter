package connections

import (
	"context"
	"log"
	"time"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func MakeUserServiceConnection() (pb.UserServiceClient, context.Context) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c, ctx
}
