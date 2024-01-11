// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/proto/user_v1/service.proto

package __

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
	Userv1_CreateUser_FullMethodName = "/user_v1.Userv1/CreateUser"
	Userv1_GetUser_FullMethodName    = "/user_v1.Userv1/GetUser"
)

// Userv1Client is the client API for Userv1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Userv1Client interface {
	CreateUser(ctx context.Context, in *PostRequestUser, opts ...grpc.CallOption) (*PostResponseUser, error)
	GetUser(ctx context.Context, in *GetRequestUser, opts ...grpc.CallOption) (*GetResponseUser, error)
}

type userv1Client struct {
	cc grpc.ClientConnInterface
}

func NewUserv1Client(cc grpc.ClientConnInterface) Userv1Client {
	return &userv1Client{cc}
}

func (c *userv1Client) CreateUser(ctx context.Context, in *PostRequestUser, opts ...grpc.CallOption) (*PostResponseUser, error) {
	out := new(PostResponseUser)
	err := c.cc.Invoke(ctx, Userv1_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userv1Client) GetUser(ctx context.Context, in *GetRequestUser, opts ...grpc.CallOption) (*GetResponseUser, error) {
	out := new(GetResponseUser)
	err := c.cc.Invoke(ctx, Userv1_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Userv1Server is the server API for Userv1 service.
// All implementations must embed UnimplementedUserv1Server
// for forward compatibility
type Userv1Server interface {
	CreateUser(context.Context, *PostRequestUser) (*PostResponseUser, error)
	GetUser(context.Context, *GetRequestUser) (*GetResponseUser, error)
	mustEmbedUnimplementedUserv1Server()
}

// UnimplementedUserv1Server must be embedded to have forward compatible implementations.
type UnimplementedUserv1Server struct {
}

func (UnimplementedUserv1Server) CreateUser(context.Context, *PostRequestUser) (*PostResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserv1Server) GetUser(context.Context, *GetRequestUser) (*GetResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserv1Server) mustEmbedUnimplementedUserv1Server() {}

// UnsafeUserv1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Userv1Server will
// result in compilation errors.
type UnsafeUserv1Server interface {
	mustEmbedUnimplementedUserv1Server()
}

func RegisterUserv1Server(s grpc.ServiceRegistrar, srv Userv1Server) {
	s.RegisterService(&Userv1_ServiceDesc, srv)
}

func _Userv1_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequestUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Userv1Server).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userv1_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Userv1Server).CreateUser(ctx, req.(*PostRequestUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userv1_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Userv1Server).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userv1_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Userv1Server).GetUser(ctx, req.(*GetRequestUser))
	}
	return interceptor(ctx, in, info, handler)
}

// Userv1_ServiceDesc is the grpc.ServiceDesc for Userv1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Userv1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.Userv1",
	HandlerType: (*Userv1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Userv1_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Userv1_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/user_v1/service.proto",
}