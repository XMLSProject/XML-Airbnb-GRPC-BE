// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: reservation/reservation-service.proto

package reservation

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
	GreetFromReservation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Reserve(ctx context.Context, in *RequestForReserve, opts ...grpc.CallOption) (*ResponseForReserve, error)
	DeleteReservation(ctx context.Context, in *RequestDeleteReservation, opts ...grpc.CallOption) (*ResponseDeleteReservation, error)
	AcceptReservation(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	CheckReservations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CheckReservationsByDates(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	GetAllReservations(ctx context.Context, in *AllReservationsRequest, opts ...grpc.CallOption) (*AllReservationsResponse, error)
	CheckForGuests(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) GreetFromReservation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ReservationService/GreetFromReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Reserve(ctx context.Context, in *RequestForReserve, opts ...grpc.CallOption) (*ResponseForReserve, error) {
	out := new(ResponseForReserve)
	err := c.cc.Invoke(ctx, "/ReservationService/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteReservation(ctx context.Context, in *RequestDeleteReservation, opts ...grpc.CallOption) (*ResponseDeleteReservation, error) {
	out := new(ResponseDeleteReservation)
	err := c.cc.Invoke(ctx, "/ReservationService/DeleteReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) AcceptReservation(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/ReservationService/AcceptReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckReservations(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ReservationService/CheckReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckReservationsByDates(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/ReservationService/CheckReservationsByDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllReservations(ctx context.Context, in *AllReservationsRequest, opts ...grpc.CallOption) (*AllReservationsResponse, error) {
	out := new(AllReservationsResponse)
	err := c.cc.Invoke(ctx, "/ReservationService/GetAllReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckForGuests(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ReservationService/CheckForGuests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	GreetFromReservation(context.Context, *Request) (*Response, error)
	Reserve(context.Context, *RequestForReserve) (*ResponseForReserve, error)
	DeleteReservation(context.Context, *RequestDeleteReservation) (*ResponseDeleteReservation, error)
	AcceptReservation(context.Context, *DeleteRequest) (*DeleteResponse, error)
	CheckReservations(context.Context, *Request) (*Response, error)
	CheckReservationsByDates(context.Context, *CheckRequest) (*CheckResponse, error)
	GetAllReservations(context.Context, *AllReservationsRequest) (*AllReservationsResponse, error)
	CheckForGuests(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) GreetFromReservation(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetFromReservation not implemented")
}
func (UnimplementedReservationServiceServer) Reserve(context.Context, *RequestForReserve) (*ResponseForReserve, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}
func (UnimplementedReservationServiceServer) DeleteReservation(context.Context, *RequestDeleteReservation) (*ResponseDeleteReservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (UnimplementedReservationServiceServer) AcceptReservation(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptReservation not implemented")
}
func (UnimplementedReservationServiceServer) CheckReservations(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckReservations not implemented")
}
func (UnimplementedReservationServiceServer) CheckReservationsByDates(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckReservationsByDates not implemented")
}
func (UnimplementedReservationServiceServer) GetAllReservations(context.Context, *AllReservationsRequest) (*AllReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservations not implemented")
}
func (UnimplementedReservationServiceServer) CheckForGuests(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckForGuests not implemented")
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

func _ReservationService_GreetFromReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GreetFromReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/GreetFromReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GreetFromReservation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForReserve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Reserve(ctx, req.(*RequestForReserve))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteReservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/DeleteReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, req.(*RequestDeleteReservation))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_AcceptReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).AcceptReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/AcceptReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).AcceptReservation(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/CheckReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckReservations(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckReservationsByDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckReservationsByDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/CheckReservationsByDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckReservationsByDates(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/GetAllReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllReservations(ctx, req.(*AllReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckForGuests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckForGuests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReservationService/CheckForGuests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckForGuests(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetFromReservation",
			Handler:    _ReservationService_GreetFromReservation_Handler,
		},
		{
			MethodName: "Reserve",
			Handler:    _ReservationService_Reserve_Handler,
		},
		{
			MethodName: "DeleteReservation",
			Handler:    _ReservationService_DeleteReservation_Handler,
		},
		{
			MethodName: "AcceptReservation",
			Handler:    _ReservationService_AcceptReservation_Handler,
		},
		{
			MethodName: "CheckReservations",
			Handler:    _ReservationService_CheckReservations_Handler,
		},
		{
			MethodName: "CheckReservationsByDates",
			Handler:    _ReservationService_CheckReservationsByDates_Handler,
		},
		{
			MethodName: "GetAllReservations",
			Handler:    _ReservationService_GetAllReservations_Handler,
		},
		{
			MethodName: "CheckForGuests",
			Handler:    _ReservationService_CheckForGuests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation/reservation-service.proto",
}
