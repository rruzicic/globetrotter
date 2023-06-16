package grpcserver

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/rruzicic/globetrotter/bnb/account-service/jwt"
	"github.com/rruzicic/globetrotter/bnb/account-service/models"
	"github.com/rruzicic/globetrotter/bnb/account-service/pb"
	"github.com/rruzicic/globetrotter/bnb/account-service/repos"
	"github.com/rruzicic/globetrotter/bnb/account-service/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func userToUserResponse(user *models.User) *pb.UserResponse {
	return &pb.UserResponse{
		User: &pb.User{
			Id:                   user.Id.Hex(),
			CreatedOn:            timestamppb.New(time.Unix(int64(user.CreatedOn), 0)),
			ModifiedOn:           timestamppb.New(time.Unix(int64(user.ModifiedOn), 0)),
			DeletedOn:            timestamppb.New(time.Unix(int64(user.DeletedOn), 0)),
			FirstName:            user.FirstName,
			LastName:             user.LastName,
			Email:                user.EMail,
			Password:             user.Password,
			Role:                 user.Role,
			Country:              user.Country,
			Street:               user.Street,
			StreetNum:            user.StreetNum,
			ZipCode:              int32(user.ZIPCode),
			SuperHost:            user.SuperHost,
			Rating:               user.Rating,
			RatingNum:            int32(user.RatingNum),
			CancellationsCounter: int32(user.CancellationsCounter),
		},
	}
}

func (s *server) GetUserById(ctx context.Context, in *pb.UserRequestId) (*pb.UserResponse, error) {
	objId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return &pb.UserResponse{}, err
	}
	user, err := services.GetById(objId)
	if err != nil {
		return &pb.UserResponse{}, err
	}
	return userToUserResponse(user), nil
}

func (s *server) GetUserByEmail(ctx context.Context, in *pb.UserRequestEmail) (*pb.UserResponse, error) {
	user, err := services.GetByEmail(in.Email)
	if err != nil {
		return &pb.UserResponse{}, err
	}
	return userToUserResponse(user), nil
}

func (s *server) GetUsers(pagination *pb.UserRequestPagination, stream pb.UserService_GetUsersServer) error {
	users := repos.GetAllUsers()

	for _, user := range users {
		grpc_user := userToUserResponse(&user)
		if err := stream.Send(grpc_user); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.UserRequestId) (*pb.UserResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *server) IncrementCancellationsCounter(ctx context.Context, in *pb.UserRequestId) (*pb.UserResponse, error) {
	objId, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return &pb.UserResponse{}, err
	}
	user, err := services.IncrementCancellationsCounter(objId)
	if err != nil {
		return &pb.UserResponse{}, err
	}
	return userToUserResponse(user), nil
}

func (s *server) VerifyToken(ctx context.Context, in *pb.TokenRequest) (*pb.BooleanReturn, error) {
	return &pb.BooleanReturn{Boolean: jwt.IsTokenValid(in.Token, in.Role)}, nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("gRPC failed to listen")
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})
	log.Printf("gRPC server listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}
