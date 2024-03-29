// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: accommodation.proto

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

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	TestConnection(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error)
	GetAllAccommodations(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AccommodationService_GetAllAccommodationsClient, error)
	GetAccommodationById(ctx context.Context, in *RequestAccommodationById, opts ...grpc.CallOption) (*Accommodation, error)
	GetAccommodationByHostId(ctx context.Context, in *RequestAccommodationByHostId, opts ...grpc.CallOption) (AccommodationService_GetAccommodationByHostIdClient, error)
	AddReservationToAccommodation(ctx context.Context, in *AddReservationToAccommodationRequest, opts ...grpc.CallOption) (*BoolAnswer, error)
	RemoveReservationFromAccommodation(ctx context.Context, in *AddReservationToAccommodationRequest, opts ...grpc.CallOption) (*BoolAnswer, error)
	GetPastHostsByAccommodations(ctx context.Context, in *RequestGetPastHostsByAccommodations, opts ...grpc.CallOption) (AccommodationService_GetPastHostsByAccommodationsClient, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) TestConnection(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error) {
	out := new(TestMessage)
	err := c.cc.Invoke(ctx, "/pb.AccommodationService/TestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllAccommodations(ctx context.Context, in *Empty, opts ...grpc.CallOption) (AccommodationService_GetAllAccommodationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccommodationService_ServiceDesc.Streams[0], "/pb.AccommodationService/GetAllAccommodations", opts...)
	if err != nil {
		return nil, err
	}
	x := &accommodationServiceGetAllAccommodationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccommodationService_GetAllAccommodationsClient interface {
	Recv() (*Accommodation, error)
	grpc.ClientStream
}

type accommodationServiceGetAllAccommodationsClient struct {
	grpc.ClientStream
}

func (x *accommodationServiceGetAllAccommodationsClient) Recv() (*Accommodation, error) {
	m := new(Accommodation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accommodationServiceClient) GetAccommodationById(ctx context.Context, in *RequestAccommodationById, opts ...grpc.CallOption) (*Accommodation, error) {
	out := new(Accommodation)
	err := c.cc.Invoke(ctx, "/pb.AccommodationService/GetAccommodationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccommodationByHostId(ctx context.Context, in *RequestAccommodationByHostId, opts ...grpc.CallOption) (AccommodationService_GetAccommodationByHostIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccommodationService_ServiceDesc.Streams[1], "/pb.AccommodationService/GetAccommodationByHostId", opts...)
	if err != nil {
		return nil, err
	}
	x := &accommodationServiceGetAccommodationByHostIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccommodationService_GetAccommodationByHostIdClient interface {
	Recv() (*Accommodation, error)
	grpc.ClientStream
}

type accommodationServiceGetAccommodationByHostIdClient struct {
	grpc.ClientStream
}

func (x *accommodationServiceGetAccommodationByHostIdClient) Recv() (*Accommodation, error) {
	m := new(Accommodation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accommodationServiceClient) AddReservationToAccommodation(ctx context.Context, in *AddReservationToAccommodationRequest, opts ...grpc.CallOption) (*BoolAnswer, error) {
	out := new(BoolAnswer)
	err := c.cc.Invoke(ctx, "/pb.AccommodationService/AddReservationToAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) RemoveReservationFromAccommodation(ctx context.Context, in *AddReservationToAccommodationRequest, opts ...grpc.CallOption) (*BoolAnswer, error) {
	out := new(BoolAnswer)
	err := c.cc.Invoke(ctx, "/pb.AccommodationService/RemoveReservationFromAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetPastHostsByAccommodations(ctx context.Context, in *RequestGetPastHostsByAccommodations, opts ...grpc.CallOption) (AccommodationService_GetPastHostsByAccommodationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccommodationService_ServiceDesc.Streams[2], "/pb.AccommodationService/GetPastHostsByAccommodations", opts...)
	if err != nil {
		return nil, err
	}
	x := &accommodationServiceGetPastHostsByAccommodationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccommodationService_GetPastHostsByAccommodationsClient interface {
	Recv() (*HostAnswer, error)
	grpc.ClientStream
}

type accommodationServiceGetPastHostsByAccommodationsClient struct {
	grpc.ClientStream
}

func (x *accommodationServiceGetPastHostsByAccommodationsClient) Recv() (*HostAnswer, error) {
	m := new(HostAnswer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	TestConnection(context.Context, *TestMessage) (*TestMessage, error)
	GetAllAccommodations(*Empty, AccommodationService_GetAllAccommodationsServer) error
	GetAccommodationById(context.Context, *RequestAccommodationById) (*Accommodation, error)
	GetAccommodationByHostId(*RequestAccommodationByHostId, AccommodationService_GetAccommodationByHostIdServer) error
	AddReservationToAccommodation(context.Context, *AddReservationToAccommodationRequest) (*BoolAnswer, error)
	RemoveReservationFromAccommodation(context.Context, *AddReservationToAccommodationRequest) (*BoolAnswer, error)
	GetPastHostsByAccommodations(*RequestGetPastHostsByAccommodations, AccommodationService_GetPastHostsByAccommodationsServer) error
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) TestConnection(context.Context, *TestMessage) (*TestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllAccommodations(*Empty, AccommodationService_GetAllAccommodationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAccommodations not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAccommodationById(context.Context, *RequestAccommodationById) (*Accommodation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationById not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAccommodationByHostId(*RequestAccommodationByHostId, AccommodationService_GetAccommodationByHostIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAccommodationByHostId not implemented")
}
func (UnimplementedAccommodationServiceServer) AddReservationToAccommodation(context.Context, *AddReservationToAccommodationRequest) (*BoolAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReservationToAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) RemoveReservationFromAccommodation(context.Context, *AddReservationToAccommodationRequest) (*BoolAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReservationFromAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) GetPastHostsByAccommodations(*RequestGetPastHostsByAccommodations, AccommodationService_GetPastHostsByAccommodationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPastHostsByAccommodations not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccommodationService/TestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).TestConnection(ctx, req.(*TestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllAccommodations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccommodationServiceServer).GetAllAccommodations(m, &accommodationServiceGetAllAccommodationsServer{stream})
}

type AccommodationService_GetAllAccommodationsServer interface {
	Send(*Accommodation) error
	grpc.ServerStream
}

type accommodationServiceGetAllAccommodationsServer struct {
	grpc.ServerStream
}

func (x *accommodationServiceGetAllAccommodationsServer) Send(m *Accommodation) error {
	return x.ServerStream.SendMsg(m)
}

func _AccommodationService_GetAccommodationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccommodationById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccommodationService/GetAccommodationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, req.(*RequestAccommodationById))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccommodationByHostId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestAccommodationByHostId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccommodationServiceServer).GetAccommodationByHostId(m, &accommodationServiceGetAccommodationByHostIdServer{stream})
}

type AccommodationService_GetAccommodationByHostIdServer interface {
	Send(*Accommodation) error
	grpc.ServerStream
}

type accommodationServiceGetAccommodationByHostIdServer struct {
	grpc.ServerStream
}

func (x *accommodationServiceGetAccommodationByHostIdServer) Send(m *Accommodation) error {
	return x.ServerStream.SendMsg(m)
}

func _AccommodationService_AddReservationToAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReservationToAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).AddReservationToAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccommodationService/AddReservationToAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).AddReservationToAccommodation(ctx, req.(*AddReservationToAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_RemoveReservationFromAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReservationToAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).RemoveReservationFromAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AccommodationService/RemoveReservationFromAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).RemoveReservationFromAccommodation(ctx, req.(*AddReservationToAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetPastHostsByAccommodations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestGetPastHostsByAccommodations)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccommodationServiceServer).GetPastHostsByAccommodations(m, &accommodationServiceGetPastHostsByAccommodationsServer{stream})
}

type AccommodationService_GetPastHostsByAccommodationsServer interface {
	Send(*HostAnswer) error
	grpc.ServerStream
}

type accommodationServiceGetPastHostsByAccommodationsServer struct {
	grpc.ServerStream
}

func (x *accommodationServiceGetPastHostsByAccommodationsServer) Send(m *HostAnswer) error {
	return x.ServerStream.SendMsg(m)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestConnection",
			Handler:    _AccommodationService_TestConnection_Handler,
		},
		{
			MethodName: "GetAccommodationById",
			Handler:    _AccommodationService_GetAccommodationById_Handler,
		},
		{
			MethodName: "AddReservationToAccommodation",
			Handler:    _AccommodationService_AddReservationToAccommodation_Handler,
		},
		{
			MethodName: "RemoveReservationFromAccommodation",
			Handler:    _AccommodationService_RemoveReservationFromAccommodation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllAccommodations",
			Handler:       _AccommodationService_GetAllAccommodations_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAccommodationByHostId",
			Handler:       _AccommodationService_GetAccommodationByHostId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPastHostsByAccommodations",
			Handler:       _AccommodationService_GetPastHostsByAccommodations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "accommodation.proto",
}
