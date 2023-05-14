package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/models"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/pb"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/repos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccommodationServiceServer struct {
	pb.UnimplementedAccommodationServiceServer
}

func buildGRPCAccommodation(accommodation models.Accommodation) pb.Accommodation {
	var reservation_ids []string
	for _, id := range accommodation.Reservations {
		reservation_ids = append(reservation_ids, id.Hex())
	}

	var commodations []string
	for _, commodation := range accommodation.AvailableCommodations {
		commodations = append(commodations, string(commodation))
	}

	return pb.Accommodation{
		Reservations:          reservation_ids,
		Name:                  accommodation.Name,
		Country:               accommodation.Location.Country,
		Street:                accommodation.Location.Street,
		StreetNum:             accommodation.Location.StreetNum,
		ZipCode:               int32(accommodation.Location.ZIPCode),
		Commodations:          commodations,
		Photos:                accommodation.Photos,
		Guests:                int32(accommodation.Guests),
		AvailabilityStartDate: timestamppb.New(accommodation.Availability.Start),
		AvailabilityEndDate:   timestamppb.New(accommodation.Availability.End),
		Amount:                accommodation.UnitPrice.Amount,
		PriceStartDate:        timestamppb.New(accommodation.UnitPrice.Duration.Start),
		PriceEndDate:          timestamppb.New(accommodation.UnitPrice.Duration.End),
		PriceForPerson:        accommodation.PriceForPerson,
		User:                  accommodation.User.Hex(),
		AutoApprove:           accommodation.AutoApprove,
		Id:                    accommodation.Id.Hex(),
	}
}

func (s *AccommodationServiceServer) GetAccommodationById(ctx context.Context, req *pb.RequestAccommodationById) (*pb.Accommodation, error) {
	accommodation, err := repos.GetAccommodationById(req.GetId())
	if err != nil {
		log.Panic("Could not get accommodation with id", req.GetId())
		return nil, err
	}

	grpc_accommodation := buildGRPCAccommodation(*accommodation)

	return &grpc_accommodation, nil
}

func (s *AccommodationServiceServer) GetAccommodationByHostId(req *pb.RequestAccommodationByHostId, stream pb.AccommodationService_GetAccommodationByHostIdServer) error {
	accommodations, err := repos.GetAccommodationsByHostId(req.GetId())
	if err != nil {
		log.Panic("Could not get accommodations for host id: ", req.GetId())
		return err
	}

	for _, accommodation := range accommodations {
		grpc_accommodation := buildGRPCAccommodation(accommodation)
		if err := stream.Send(&grpc_accommodation); err != nil {
			return err
		}
	}

	return nil
}

func (s *AccommodationServiceServer) TestConnection(ctx context.Context, req *pb.TestMessage) (*pb.TestMessage, error) {
	log.Print("Hello from accommodation service, message is: ", req.GetMsg())
	return req, nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Accommodation service failed to listen. Error: ", err)
	}

	server := grpc.NewServer()
	pb.RegisterAccommodationServiceServer(server, &AccommodationServiceServer{})

	log.Println("Accommodation gRPC server listening..")
	if err := server.Serve(listen); err != nil {
		log.Fatal("Could not start Accommodation gRPC Server. Error: ", err)
	}
}
