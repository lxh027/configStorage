// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package register

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

// RegisterRaftClient is the client API for RegisterRaft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterRaftClient interface {
	// register a new raft service
	RegisterRaft(ctx context.Context, in *RegisterRaftArgs, opts ...grpc.CallOption) (*RegisterRaftReply, error)
	// get raft config
	GetRaftRegistrations(ctx context.Context, in *GetRaftRegistrationsArgs, opts ...grpc.CallOption) (*GetRaftRegistrationsReply, error)
}

type registerRaftClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterRaftClient(cc grpc.ClientConnInterface) RegisterRaftClient {
	return &registerRaftClient{cc}
}

func (c *registerRaftClient) RegisterRaft(ctx context.Context, in *RegisterRaftArgs, opts ...grpc.CallOption) (*RegisterRaftReply, error) {
	out := new(RegisterRaftReply)
	err := c.cc.Invoke(ctx, "/registerRaft/RegisterRaft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerRaftClient) GetRaftRegistrations(ctx context.Context, in *GetRaftRegistrationsArgs, opts ...grpc.CallOption) (*GetRaftRegistrationsReply, error) {
	out := new(GetRaftRegistrationsReply)
	err := c.cc.Invoke(ctx, "/registerRaft/GetRaftRegistrations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterRaftServer is the server API for RegisterRaft service.
// All implementations must embed UnimplementedRegisterRaftServer
// for forward compatibility
type RegisterRaftServer interface {
	// register a new raft service
	RegisterRaft(context.Context, *RegisterRaftArgs) (*RegisterRaftReply, error)
	// get raft config
	GetRaftRegistrations(context.Context, *GetRaftRegistrationsArgs) (*GetRaftRegistrationsReply, error)
	mustEmbedUnimplementedRegisterRaftServer()
}

// UnimplementedRegisterRaftServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterRaftServer struct {
}

func (UnimplementedRegisterRaftServer) RegisterRaft(context.Context, *RegisterRaftArgs) (*RegisterRaftReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRaft not implemented")
}
func (UnimplementedRegisterRaftServer) GetRaftRegistrations(context.Context, *GetRaftRegistrationsArgs) (*GetRaftRegistrationsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRaftRegistrations not implemented")
}
func (UnimplementedRegisterRaftServer) mustEmbedUnimplementedRegisterRaftServer() {}

// UnsafeRegisterRaftServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterRaftServer will
// result in compilation errors.
type UnsafeRegisterRaftServer interface {
	mustEmbedUnimplementedRegisterRaftServer()
}

func RegisterRegisterRaftServer(s grpc.ServiceRegistrar, srv RegisterRaftServer) {
	s.RegisterService(&RegisterRaft_ServiceDesc, srv)
}

func _RegisterRaft_RegisterRaft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRaftArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterRaftServer).RegisterRaft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerRaft/RegisterRaft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterRaftServer).RegisterRaft(ctx, req.(*RegisterRaftArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegisterRaft_GetRaftRegistrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRaftRegistrationsArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterRaftServer).GetRaftRegistrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerRaft/GetRaftRegistrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterRaftServer).GetRaftRegistrations(ctx, req.(*GetRaftRegistrationsArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterRaft_ServiceDesc is the grpc.ServiceDesc for RegisterRaft service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegisterRaft_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registerRaft",
	HandlerType: (*RegisterRaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterRaft",
			Handler:    _RegisterRaft_RegisterRaft_Handler,
		},
		{
			MethodName: "GetRaftRegistrations",
			Handler:    _RegisterRaft_GetRaftRegistrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/register/register.proto",
}
