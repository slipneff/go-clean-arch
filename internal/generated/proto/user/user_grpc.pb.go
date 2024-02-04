// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/user/user.proto

package user

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
	UserService_CreateUser_FullMethodName    = "/user.UserService/CreateUser"
	UserService_GetUserById_FullMethodName   = "/user.UserService/GetUserById"
	UserService_GetUserByName_FullMethodName = "/user.UserService/GetUserByName"
	UserService_GetAllUser_FullMethodName    = "/user.UserService/GetAllUser"
	UserService_UpdateUser_FullMethodName    = "/user.UserService/UpdateUser"
	UserService_DeleteUser_FullMethodName    = "/user.UserService/DeleteUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserResponse, error) {
	out := new(GetAllUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*UserResponse, error)
	GetUserByName(context.Context, *GetUserByNameRequest) (*UserResponse, error)
	GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUserByName(context.Context, *GetUserByNameRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedUserServiceServer) GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByName(ctx, req.(*GetUserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUser(ctx, req.(*GetAllUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _UserService_GetUserByName_Handler,
		},
		{
			MethodName: "GetAllUser",
			Handler:    _UserService_GetAllUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/user.proto",
}
