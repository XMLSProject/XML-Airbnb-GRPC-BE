// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: greeter/greeter-service.proto

package greeter

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
	GreeterService_Greet_FullMethodName     = "/GreeterService/Greet"
	GreeterService_GreetTest_FullMethodName = "/GreeterService/GreetTest"
)

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceClient interface {
	Greet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GreetTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Greet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, GreeterService_Greet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) GreetTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, GreeterService_GreetTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServiceServer is the server API for GreeterService service.
// All implementations must embed UnimplementedGreeterServiceServer
// for forward compatibility
type GreeterServiceServer interface {
	Greet(context.Context, *Request) (*Response, error)
	GreetTest(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGreeterServiceServer()
}

// UnimplementedGreeterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (UnimplementedGreeterServiceServer) Greet(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreeterServiceServer) GreetTest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetTest not implemented")
}
func (UnimplementedGreeterServiceServer) mustEmbedUnimplementedGreeterServiceServer() {}

// UnsafeGreeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceServer will
// result in compilation errors.
type UnsafeGreeterServiceServer interface {
	mustEmbedUnimplementedGreeterServiceServer()
}

func RegisterGreeterServiceServer(s grpc.ServiceRegistrar, srv GreeterServiceServer) {
	s.RegisterService(&GreeterService_ServiceDesc, srv)
}

func _GreeterService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterService_Greet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Greet(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_GreetTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).GreetTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterService_GreetTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).GreetTest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GreeterService_ServiceDesc is the grpc.ServiceDesc for GreeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreeterService_Greet_Handler,
		},
		{
			MethodName: "GreetTest",
			Handler:    _GreeterService_GreetTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter/greeter-service.proto",
}

const (
	ReservationService_GreetFromReservation_FullMethodName = "/ReservationService/GreetFromReservation"
	ReservationService_Reserve_FullMethodName              = "/ReservationService/Reserve"
	ReservationService_DeleteReservation_FullMethodName    = "/ReservationService/DeleteReservation"
	ReservationService_AcceptReservation_FullMethodName    = "/ReservationService/AcceptReservation"
)

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
	err := c.cc.Invoke(ctx, ReservationService_GreetFromReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Reserve(ctx context.Context, in *RequestForReserve, opts ...grpc.CallOption) (*ResponseForReserve, error) {
	out := new(ResponseForReserve)
	err := c.cc.Invoke(ctx, ReservationService_Reserve_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteReservation(ctx context.Context, in *RequestDeleteReservation, opts ...grpc.CallOption) (*ResponseDeleteReservation, error) {
	out := new(ResponseDeleteReservation)
	err := c.cc.Invoke(ctx, ReservationService_DeleteReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) AcceptReservation(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ReservationService_AcceptReservation_FullMethodName, in, out, opts...)
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
		FullMethod: ReservationService_GreetFromReservation_FullMethodName,
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
		FullMethod: ReservationService_Reserve_FullMethodName,
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
		FullMethod: ReservationService_DeleteReservation_FullMethodName,
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
		FullMethod: ReservationService_AcceptReservation_FullMethodName,
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
	Metadata: "greeter/greeter-service.proto",
}

const (
	LoginService_GreetFromLogin_FullMethodName     = "/LoginService/GreetFromLogin"
	LoginService_GreetFromLoginTest_FullMethodName = "/LoginService/GreetFromLoginTest"
	LoginService_CreateUser_FullMethodName         = "/LoginService/CreateUser"
	LoginService_Login_FullMethodName              = "/LoginService/Login"
	LoginService_UpdateUser_FullMethodName         = "/LoginService/UpdateUser"
	LoginService_DeleteUser_FullMethodName         = "/LoginService/DeleteUser"
)

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	GreetFromLogin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GreetFromLoginTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UpdateUser(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) GreetFromLogin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, LoginService_GreetFromLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) GreetFromLoginTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, LoginService_GreetFromLoginTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, LoginService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, LoginService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) UpdateUser(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, LoginService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) DeleteUser(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, LoginService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	GreetFromLogin(context.Context, *Request) (*Response, error)
	GreetFromLoginTest(context.Context, *Request) (*Response, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UpdateUser(context.Context, *UpdateRequest) (*UpdateResponse, error)
	DeleteUser(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) GreetFromLogin(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetFromLogin not implemented")
}
func (UnimplementedLoginServiceServer) GreetFromLoginTest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetFromLoginTest not implemented")
}
func (UnimplementedLoginServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedLoginServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServiceServer) UpdateUser(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedLoginServiceServer) DeleteUser(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_GreetFromLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).GreetFromLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_GreetFromLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).GreetFromLogin(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_GreetFromLoginTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).GreetFromLoginTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_GreetFromLoginTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).GreetFromLoginTest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).UpdateUser(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).DeleteUser(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetFromLogin",
			Handler:    _LoginService_GreetFromLogin_Handler,
		},
		{
			MethodName: "GreetFromLoginTest",
			Handler:    _LoginService_GreetFromLoginTest_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _LoginService_CreateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _LoginService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _LoginService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter/greeter-service.proto",
}

const (
	AccommodationService_GreetFromAccommodation_FullMethodName = "/AccommodationService/GreetFromAccommodation"
	AccommodationService_CreateAccommodation_FullMethodName    = "/AccommodationService/CreateAccommodation"
	AccommodationService_EditAccommodation_FullMethodName      = "/AccommodationService/EditAccommodation"
	AccommodationService_SearchAccommodation_FullMethodName    = "/AccommodationService/SearchAccommodation"
	AccommodationService_GetAllAccommodations_FullMethodName   = "/AccommodationService/GetAllAccommodations"
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	GreetFromAccommodation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error)
	EditAccommodation(ctx context.Context, in *EditAccoRequest, opts ...grpc.CallOption) (*EditAccoResponse, error)
	SearchAccommodation(ctx context.Context, in *SearchAccoRequest, opts ...grpc.CallOption) (*SearchAccoResponse, error)
	GetAllAccommodations(ctx context.Context, in *AllAccommodationsRequest, opts ...grpc.CallOption) (*AllAccommodationsResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) GreetFromAccommodation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, AccommodationService_GreetFromAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error) {
	out := new(CreateAccommodationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_CreateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) EditAccommodation(ctx context.Context, in *EditAccoRequest, opts ...grpc.CallOption) (*EditAccoResponse, error) {
	out := new(EditAccoResponse)
	err := c.cc.Invoke(ctx, AccommodationService_EditAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) SearchAccommodation(ctx context.Context, in *SearchAccoRequest, opts ...grpc.CallOption) (*SearchAccoResponse, error) {
	out := new(SearchAccoResponse)
	err := c.cc.Invoke(ctx, AccommodationService_SearchAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllAccommodations(ctx context.Context, in *AllAccommodationsRequest, opts ...grpc.CallOption) (*AllAccommodationsResponse, error) {
	out := new(AllAccommodationsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllAccommodations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	GreetFromAccommodation(context.Context, *Request) (*Response, error)
	CreateAccommodation(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error)
	EditAccommodation(context.Context, *EditAccoRequest) (*EditAccoResponse, error)
	SearchAccommodation(context.Context, *SearchAccoRequest) (*SearchAccoResponse, error)
	GetAllAccommodations(context.Context, *AllAccommodationsRequest) (*AllAccommodationsResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) GreetFromAccommodation(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetFromAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAccommodation(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) EditAccommodation(context.Context, *EditAccoRequest) (*EditAccoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) SearchAccommodation(context.Context, *SearchAccoRequest) (*SearchAccoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllAccommodations(context.Context, *AllAccommodationsRequest) (*AllAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccommodations not implemented")
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

func _AccommodationService_GreetFromAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GreetFromAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GreetFromAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GreetFromAccommodation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, req.(*CreateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_EditAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAccoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).EditAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_EditAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).EditAccommodation(ctx, req.(*EditAccoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_SearchAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAccoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).SearchAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_SearchAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).SearchAccommodation(ctx, req.(*SearchAccoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAllAccommodations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllAccommodations(ctx, req.(*AllAccommodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetFromAccommodation",
			Handler:    _AccommodationService_GreetFromAccommodation_Handler,
		},
		{
			MethodName: "CreateAccommodation",
			Handler:    _AccommodationService_CreateAccommodation_Handler,
		},
		{
			MethodName: "EditAccommodation",
			Handler:    _AccommodationService_EditAccommodation_Handler,
		},
		{
			MethodName: "SearchAccommodation",
			Handler:    _AccommodationService_SearchAccommodation_Handler,
		},
		{
			MethodName: "GetAllAccommodations",
			Handler:    _AccommodationService_GetAllAccommodations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter/greeter-service.proto",
}
