package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/rruzicic/globetrotter/bnb/reservation-service/models"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/pb"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/repos"
	"github.com/rruzicic/globetrotter/bnb/reservation-service/services"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReservationServiceServer struct {
	pb.UnimplementedReservationServiceServer
}

func buildGRPCReservation(reservation models.Reservation) pb.Reservation {
	return pb.Reservation{
		AccommodationId: reservation.AccommodationId.Hex(),
		UserId:          reservation.UserId.Hex(),
		StartDate:       timestamppb.New(reservation.DateInterval.Start),
		EndDate:         timestamppb.New(reservation.DateInterval.End),
		NumOfGuests:     int32(reservation.NumOfGuests),
		IsApproved:      reservation.IsApproved,
		TotalPrice:      reservation.TotalPrice,
	}
}

func (s *ReservationServiceServer) GetReservationById(ctx context.Context, req *pb.RequestReservationById) (*pb.Reservation, error) {
	reservation, err := repos.GetReservationById(req.GetId())
	if err != nil {
		log.Panic("Could not get reservation with id: ", req.GetId())
		return nil, err
	}

	grpc_reservation := buildGRPCReservation(*reservation)

	return &grpc_reservation, nil
}

func (s *ReservationServiceServer) GetReservationsByAccommodationId(req *pb.RequestReservationsByAccommodationId, stream pb.ReservationService_GetReservationsByAccommodationIdServer) error {
	reservations, err := repos.GetReservationsByAccommodationId(req.GetId())
	if err != nil {
		log.Panic("Could not get reservations with accommodation id: ", req.GetId())
		return err
	}

	for _, reservation := range reservations {
		grpc_reservation := buildGRPCReservation(reservation)
		if err := stream.Send(&grpc_reservation); err != nil {
			return err
		}
	}

	return nil
}

func (s *ReservationServiceServer) GetActiveReservationsByUser(req *pb.RequestUserId, stream pb.ReservationService_GetActiveReservationsByUserServer) error {
	reservations, err := services.GetActiveReservationsByUser(req.GetId())
	if err != nil {
		log.Println("Could not get reservations for user id: ", req.GetId())
		return err
	}

	for _, reservation := range reservations {
		grpc_reservation := buildGRPCReservation(reservation)
		if err := stream.Send(&grpc_reservation); err != nil {
			return err
		}
	}

	return nil
}

func (s *ReservationServiceServer) GetFinishedReservationsByUser(req *pb.RequestUserId, stream pb.ReservationService_GetFinishedReservationsByUserServer) error {
	reservations, err := services.GetFinishedReservationsByUser(req.GetId())
	if err != nil {
		log.Println("Could not get finished reservations for user id: ", req.GetId())
		return err
	}

	for _, reservation := range reservations {
		grpc_reservation := buildGRPCReservation(reservation)
		if err := stream.Send(&grpc_reservation); err != nil {
			return err
		}
	}

	return nil
}

func (s *ReservationServiceServer) GetFutureActiveReservationsByHost(req *pb.RequestUserId, stream pb.ReservationService_GetFutureActiveReservationsByHostServer) error {
	reservations, err := services.GetFutureActiveReservationsByHost(req.GetId())
	if err != nil {
		log.Panic("Could not get future active reservations by host with id: ", req.GetId())
		return err
	}
	//log.Println("IM HEREEEEEEEEE", reservations)
	for _, reservation := range reservations {
		grpc_reservation := buildGRPCReservation(reservation)
		if err := stream.Send(&grpc_reservation); err != nil {
			return err
		}
	}

	return nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panic("Reservation service failed to listen. Error: ", err)
	}

	server := grpc.NewServer()
	pb.RegisterReservationServiceServer(server, &ReservationServiceServer{})

	log.Println("Reservation gRPC server listening..")
	if err := server.Serve(listen); err != nil {
		log.Panic("Could not start Reservation gRPC Server. Error: ", err)
	}
}
