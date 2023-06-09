// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: api/user/service/v1/user.proto

package v1

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
	UserService_GetUser_FullMethodName           = "/api.user.service.v1.UserService/GetUser"
	UserService_GetUserByUsername_FullMethodName = "/api.user.service.v1.UserService/GetUserByUsername"
	UserService_Save_FullMethodName              = "/api.user.service.v1.UserService/Save"
	UserService_CreateUser_FullMethodName        = "/api.user.service.v1.UserService/CreateUser"
	UserService_ListUser_FullMethodName          = "/api.user.service.v1.UserService/ListUser"
	UserService_VerifyPassword_FullMethodName    = "/api.user.service.v1.UserService/VerifyPassword"
	UserService_AddFollower_FullMethodName       = "/api.user.service.v1.UserService/AddFollower"
	UserService_ConfirmFriendship_FullMethodName = "/api.user.service.v1.UserService/ConfirmFriendship"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserReply, error)
	GetUserByUsername(ctx context.Context, in *GetUserByUsernameReq, opts ...grpc.CallOption) (*GetUserByUsernameReply, error)
	Save(ctx context.Context, in *SaveUserReq, opts ...grpc.CallOption) (*SaveUserReply, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserReply, error)
	ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserReply, error)
	VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...grpc.CallOption) (*VerifyPasswordReply, error)
	AddFollower(ctx context.Context, in *AddFollowerReq, opts ...grpc.CallOption) (*AddFollowerReply, error)
	ConfirmFriendship(ctx context.Context, in *ConfirmFriendshipReq, opts ...grpc.CallOption) (*ConfirmFriendshipReply, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *GetUserByUsernameReq, opts ...grpc.CallOption) (*GetUserByUsernameReply, error) {
	out := new(GetUserByUsernameReply)
	err := c.cc.Invoke(ctx, UserService_GetUserByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Save(ctx context.Context, in *SaveUserReq, opts ...grpc.CallOption) (*SaveUserReply, error) {
	out := new(SaveUserReply)
	err := c.cc.Invoke(ctx, UserService_Save_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUser(ctx context.Context, in *ListUserReq, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, UserService_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...grpc.CallOption) (*VerifyPasswordReply, error) {
	out := new(VerifyPasswordReply)
	err := c.cc.Invoke(ctx, UserService_VerifyPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddFollower(ctx context.Context, in *AddFollowerReq, opts ...grpc.CallOption) (*AddFollowerReply, error) {
	out := new(AddFollowerReply)
	err := c.cc.Invoke(ctx, UserService_AddFollower_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmFriendship(ctx context.Context, in *ConfirmFriendshipReq, opts ...grpc.CallOption) (*ConfirmFriendshipReply, error) {
	out := new(ConfirmFriendshipReply)
	err := c.cc.Invoke(ctx, UserService_ConfirmFriendship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUser(context.Context, *GetUserReq) (*GetUserReply, error)
	GetUserByUsername(context.Context, *GetUserByUsernameReq) (*GetUserByUsernameReply, error)
	Save(context.Context, *SaveUserReq) (*SaveUserReply, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserReply, error)
	ListUser(context.Context, *ListUserReq) (*ListUserReply, error)
	VerifyPassword(context.Context, *VerifyPasswordReq) (*VerifyPasswordReply, error)
	AddFollower(context.Context, *AddFollowerReq) (*AddFollowerReply, error)
	ConfirmFriendship(context.Context, *ConfirmFriendshipReq) (*ConfirmFriendshipReply, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserReq) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByUsername(context.Context, *GetUserByUsernameReq) (*GetUserByUsernameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUserServiceServer) Save(context.Context, *SaveUserReq) (*SaveUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ListUser(context.Context, *ListUserReq) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServiceServer) VerifyPassword(context.Context, *VerifyPasswordReq) (*VerifyPasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPassword not implemented")
}
func (UnimplementedUserServiceServer) AddFollower(context.Context, *AddFollowerReq) (*AddFollowerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFollower not implemented")
}
func (UnimplementedUserServiceServer) ConfirmFriendship(context.Context, *ConfirmFriendshipReq) (*ConfirmFriendshipReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmFriendship not implemented")
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

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByUsername(ctx, req.(*GetUserByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Save(ctx, req.(*SaveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
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
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUser(ctx, req.(*ListUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_VerifyPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyPassword(ctx, req.(*VerifyPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFollowerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddFollower_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddFollower(ctx, req.(*AddFollowerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmFriendship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmFriendshipReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmFriendship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ConfirmFriendship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmFriendship(ctx, req.(*ConfirmFriendshipReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.service.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UserService_GetUserByUsername_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _UserService_Save_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _UserService_ListUser_Handler,
		},
		{
			MethodName: "VerifyPassword",
			Handler:    _UserService_VerifyPassword_Handler,
		},
		{
			MethodName: "AddFollower",
			Handler:    _UserService_AddFollower_Handler,
		},
		{
			MethodName: "ConfirmFriendship",
			Handler:    _UserService_ConfirmFriendship_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/service/v1/user.proto",
}
