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

func connectToUserService() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("account-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Panic("Could not connect to account service. Error: ", err.Error())
		return nil, err
	}

	return conn, nil
}

func buildLocalUser(user *pb.User) models.User {
	return models.User{
		Name:    user.FirstName + " " + user.LastName,
		MongoId: user.Id,
	}
}

func GetAllUsers() ([]models.User, error) {
	conn, _ := connectToUserService()
	client := pb.NewUserServiceClient(conn)

	users := []models.User{}
	stream, err := client.GetUsers(context.Background(), &pb.UserRequestPagination{})
	if err != nil {
		log.Println("Could not get stream of users, Error: ", err.Error())
		return nil, err
	}

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error in reading from user stream. Error: ", err.Error())
			return nil, err
		}
		users = append(users, buildLocalUser(user.User))
	}

	defer conn.Close()

	return users, nil
}
