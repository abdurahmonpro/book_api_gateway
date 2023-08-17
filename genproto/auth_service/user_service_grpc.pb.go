// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user_service.proto

package auth_service

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
	UserService_Create_FullMethodName        = "/auth_service.UserService/Create"
	UserService_GetByID_FullMethodName       = "/auth_service.UserService/GetByID"
	UserService_GetUserList_FullMethodName   = "/auth_service.UserService/GetUserList"
	UserService_CheckUser_FullMethodName     = "/auth_service.UserService/CheckUser"
	UserService_GetUserByName_FullMethodName = "/auth_service.UserService/GetUserByName"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetByID(ctx context.Context, in *UserPK, opts ...grpc.CallOption) (*User, error)
	GetUserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	CheckUser(ctx context.Context, in *CheckUserRequest, opts ...grpc.CallOption) (*CheckUserResponse, error)
	GetUserByName(ctx context.Context, in *GetByName, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByID(ctx context.Context, in *UserPK, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckUser(ctx context.Context, in *CheckUserRequest, opts ...grpc.CallOption) (*CheckUserResponse, error) {
	out := new(CheckUserResponse)
	err := c.cc.Invoke(ctx, UserService_CheckUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByName(ctx context.Context, in *GetByName, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Create(context.Context, *CreateUser) (*CreateUserResponse, error)
	GetByID(context.Context, *UserPK) (*User, error)
	GetUserList(context.Context, *UserListRequest) (*UserListResponse, error)
	CheckUser(context.Context, *CheckUserRequest) (*CheckUserResponse, error)
	GetUserByName(context.Context, *GetByName) (*CreateUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Create(context.Context, *CreateUser) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) GetByID(context.Context, *UserPK) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserServiceServer) GetUserList(context.Context, *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServiceServer) CheckUser(context.Context, *CheckUserRequest) (*CheckUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByName(context.Context, *GetByName) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
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

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByID(ctx, req.(*UserPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckUser(ctx, req.(*CheckUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByName)
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
		return srv.(UserServiceServer).GetUserByName(ctx, req.(*GetByName))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _UserService_GetByID_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _UserService_GetUserList_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _UserService_CheckUser_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _UserService_GetUserByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
