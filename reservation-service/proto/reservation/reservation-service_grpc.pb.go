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

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	GreetFromReservation(context.Context, *Request) (*Response, error)
	Reserve(context.Context, *RequestForReserve) (*ResponseForReserve, error)
	DeleteReservation(context.Context, *RequestDeleteReservation) (*ResponseDeleteReservation, error)
	AcceptReservation(context.Context, *DeleteRequest) (*DeleteResponse, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation/reservation-service.proto",
}
