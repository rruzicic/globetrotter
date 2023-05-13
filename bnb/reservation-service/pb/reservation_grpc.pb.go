// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
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

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	GetReservationById(ctx context.Context, in *RequestReservationById, opts ...grpc.CallOption) (*Reservation, error)
	GetReservationsByAccommodationId(ctx context.Context, in *RequestReservationsByAccommodationId, opts ...grpc.CallOption) (ReservationService_GetReservationsByAccommodationIdClient, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) GetReservationById(ctx context.Context, in *RequestReservationById, opts ...grpc.CallOption) (*Reservation, error) {
	out := new(Reservation)
	err := c.cc.Invoke(ctx, "/pb.ReservationService/GetReservationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservationsByAccommodationId(ctx context.Context, in *RequestReservationsByAccommodationId, opts ...grpc.CallOption) (ReservationService_GetReservationsByAccommodationIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReservationService_ServiceDesc.Streams[0], "/pb.ReservationService/GetReservationsByAccommodationId", opts...)
	if err != nil {
		return nil, err
	}
	x := &reservationServiceGetReservationsByAccommodationIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReservationService_GetReservationsByAccommodationIdClient interface {
	Recv() (*Reservation, error)
	grpc.ClientStream
}

type reservationServiceGetReservationsByAccommodationIdClient struct {
	grpc.ClientStream
}

func (x *reservationServiceGetReservationsByAccommodationIdClient) Recv() (*Reservation, error) {
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
	GetReservationById(context.Context, *RequestReservationById) (*Reservation, error)
	GetReservationsByAccommodationId(*RequestReservationsByAccommodationId, ReservationService_GetReservationsByAccommodationIdServer) error
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) GetReservationById(context.Context, *RequestReservationById) (*Reservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationById not implemented")
}
func (UnimplementedReservationServiceServer) GetReservationsByAccommodationId(*RequestReservationsByAccommodationId, ReservationService_GetReservationsByAccommodationIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetReservationsByAccommodationId not implemented")
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

func _ReservationService_GetReservationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReservationById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReservationService/GetReservationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservationById(ctx, req.(*RequestReservationById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservationsByAccommodationId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestReservationsByAccommodationId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReservationServiceServer).GetReservationsByAccommodationId(m, &reservationServiceGetReservationsByAccommodationIdServer{stream})
}

type ReservationService_GetReservationsByAccommodationIdServer interface {
	Send(*Reservation) error
	grpc.ServerStream
}

type reservationServiceGetReservationsByAccommodationIdServer struct {
	grpc.ServerStream
}

func (x *reservationServiceGetReservationsByAccommodationIdServer) Send(m *Reservation) error {
	return x.ServerStream.SendMsg(m)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReservationById",
			Handler:    _ReservationService_GetReservationById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetReservationsByAccommodationId",
			Handler:       _ReservationService_GetReservationsByAccommodationId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reservation.proto",
}
