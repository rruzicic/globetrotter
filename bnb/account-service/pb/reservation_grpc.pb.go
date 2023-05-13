// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: reservation.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ReservationService_GetActiveReservationsByUser_FullMethodName       = "/pb.ReservationService/GetActiveReservationsByUser"
	ReservationService_GetFutureActiveReservationsByHost_FullMethodName = "/pb.ReservationService/GetFutureActiveReservationsByHost"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	GetActiveReservationsByUser(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (ReservationService_GetActiveReservationsByUserClient, error)
	GetFutureActiveReservationsByHost(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (ReservationService_GetFutureActiveReservationsByHostClient, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) GetActiveReservationsByUser(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (ReservationService_GetActiveReservationsByUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReservationService_ServiceDesc.Streams[0], ReservationService_GetActiveReservationsByUser_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &reservationServiceGetActiveReservationsByUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReservationService_GetActiveReservationsByUserClient interface {
	Recv() (*Reservation, error)
	grpc.ClientStream
}

type reservationServiceGetActiveReservationsByUserClient struct {
	grpc.ClientStream
}

func (x *reservationServiceGetActiveReservationsByUserClient) Recv() (*Reservation, error) {
	m := new(Reservation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reservationServiceClient) GetFutureActiveReservationsByHost(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (ReservationService_GetFutureActiveReservationsByHostClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReservationService_ServiceDesc.Streams[1], ReservationService_GetFutureActiveReservationsByHost_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &reservationServiceGetFutureActiveReservationsByHostClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReservationService_GetFutureActiveReservationsByHostClient interface {
	Recv() (*Reservation, error)
	grpc.ClientStream
}

type reservationServiceGetFutureActiveReservationsByHostClient struct {
	grpc.ClientStream
}

func (x *reservationServiceGetFutureActiveReservationsByHostClient) Recv() (*Reservation, error) {
	m := new(Reservation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	GetActiveReservationsByUser(*RequestUserId, ReservationService_GetActiveReservationsByUserServer) error
	GetFutureActiveReservationsByHost(*RequestUserId, ReservationService_GetFutureActiveReservationsByHostServer) error
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) GetActiveReservationsByUser(*RequestUserId, ReservationService_GetActiveReservationsByUserServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActiveReservationsByUser not implemented")
}
func (UnimplementedReservationServiceServer) GetFutureActiveReservationsByHost(*RequestUserId, ReservationService_GetFutureActiveReservationsByHostServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFutureActiveReservationsByHost not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_GetActiveReservationsByUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReservationServiceServer).GetActiveReservationsByUser(m, &reservationServiceGetActiveReservationsByUserServer{stream})
}

type ReservationService_GetActiveReservationsByUserServer interface {
	Send(*Reservation) error
	grpc.ServerStream
}

type reservationServiceGetActiveReservationsByUserServer struct {
	grpc.ServerStream
}

func (x *reservationServiceGetActiveReservationsByUserServer) Send(m *Reservation) error {
	return x.ServerStream.SendMsg(m)
}

func _ReservationService_GetFutureActiveReservationsByHost_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestUserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReservationServiceServer).GetFutureActiveReservationsByHost(m, &reservationServiceGetFutureActiveReservationsByHostServer{stream})
}

type ReservationService_GetFutureActiveReservationsByHostServer interface {
	Send(*Reservation) error
	grpc.ServerStream
}

type reservationServiceGetFutureActiveReservationsByHostServer struct {
	grpc.ServerStream
}

func (x *reservationServiceGetFutureActiveReservationsByHostServer) Send(m *Reservation) error {
	return x.ServerStream.SendMsg(m)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActiveReservationsByUser",
			Handler:       _ReservationService_GetActiveReservationsByUser_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFutureActiveReservationsByHost",
			Handler:       _ReservationService_GetFutureActiveReservationsByHost_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reservation.proto",
}