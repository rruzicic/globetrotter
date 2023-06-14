package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/rruzicic/globetrotter/bnb/accommodation-service/models"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/pb"
	"github.com/rruzicic/globetrotter/bnb/accommodation-service/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func buildGRPCHostId(hostId string) pb.HostAnswer {
	return pb.HostAnswer{
		HostId: hostId,
	}
}

func (s *AccommodationServiceServer) GetAccommodationById(ctx context.Context, req *pb.RequestAccommodationById) (*pb.Accommodation, error) {
	accommodation, err := repos.GetAccommodationById(req.GetId())
	if err != nil {
		log.Println("Could not get accommodation with id", req.GetId(), ", error: ", err.Error())
		return nil, err
	}

	grpc_accommodation := buildGRPCAccommodation(*accommodation)

	return &grpc_accommodation, nil
}

func (s *AccommodationServiceServer) GetAccommodationByHostId(req *pb.RequestAccommodationByHostId, stream pb.AccommodationService_GetAccommodationByHostIdServer) error {
	accommodations, err := repos.GetAccommodationsByHostId(req.GetId())
	if err != nil {
		log.Println("Could not get accommodations for host id: ", req.GetId(), ", error: ", err.Error())
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

func (s *AccommodationServiceServer) GetPastHostsByAccommodations(req *pb.RequestGetPastHostsByAccommodations, stream pb.AccommodationService_GetPastHostsByAccommodationsServer) error {
	pastHosts := []string{}
	for _, id := range req.AccommodationId {
		accommodation, err := repos.GetAccommodationById(id)
		if err != nil {
			log.Println("Could not get accommodations for host id: ", id, ", error: ", err.Error())
			return err
		}
		pastHosts = append(pastHosts, accommodation.User.Hex())
	}

	for _, pastHost := range pastHosts {
		grpc_hostId := buildGRPCHostId(pastHost)
		if err := stream.Send(&grpc_hostId); err != nil {
			return err
		}
	}

	return nil
}

func (s *AccommodationServiceServer) TestConnection(ctx context.Context, req *pb.TestMessage) (*pb.TestMessage, error) {
	log.Print("Hello from accommodation service, message is: ", req.GetMsg())
	return req, nil
}

func (s *AccommodationServiceServer) AddReservationToAccommodation(ctx context.Context, req *pb.AddReservationToAccommodationRequest) (*pb.BoolAnswer, error) {
	accommodation, err := repos.GetAccommodationById(req.GetAccommodationId())
	if err != nil {
		log.Panic("Could not get accommodation by id: ", req.GetAccommodationId())
		return &pb.BoolAnswer{Answer: false}, err
	}

	primitive_res_id, err := primitive.ObjectIDFromHex(req.GetReservationId())
	if err != nil {
		log.Panic("Could not get id from reservation id. Error: ", err.Error())
		return &pb.BoolAnswer{Answer: false}, err
	}

	accommodation.Reservations = append(accommodation.Reservations, &primitive_res_id)
	if err := repos.UpdateAccommodation(*accommodation); err != nil {
		log.Panic("Could not add reservation to accommodation. Error: ", err.Error())
		return &pb.BoolAnswer{Answer: false}, err
	}

	return &pb.BoolAnswer{Answer: true}, nil
}

func (s *AccommodationServiceServer) RemoveReservationFromAccommodation(ctx context.Context, req *pb.AddReservationToAccommodationRequest) (*pb.BoolAnswer, error) {
	accommodation, err := repos.GetAccommodationById(req.GetAccommodationId())
	if err != nil {
		log.Panic("Could not get accommodation by id: ", req.GetAccommodationId())
		return &pb.BoolAnswer{Answer: false}, err
	}

	primitive_res_id, err := primitive.ObjectIDFromHex(req.GetReservationId())
	if err != nil {
		log.Panic("Could not get id from reservation id. Error: ", err.Error())
		return &pb.BoolAnswer{Answer: false}, err
	}

	i := 0
	for index, res_id := range accommodation.Reservations {
		if primitive_res_id == *res_id {
			i = index
			break
		}
	}

	accommodation.Reservations = append(accommodation.Reservations[:i], accommodation.Reservations[i+1:]...)
	if err := repos.UpdateAccommodation(*accommodation); err != nil {
		log.Panic("Could not add reservation to accommodation. Error: ", err.Error())
		return &pb.BoolAnswer{Answer: false}, err
	}

	return &pb.BoolAnswer{Answer: true}, nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Accommodation service failed to listen. Error: ", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterAccommodationServiceServer(server, &AccommodationServiceServer{})

	log.Println("Accommodation gRPC server listening..")
	if err := server.Serve(listen); err != nil {
		log.Println("Could not start Accommodation gRPC Server. Error: ", err.Error())
	}
}
