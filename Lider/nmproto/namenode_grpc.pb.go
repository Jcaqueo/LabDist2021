// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package nmproto

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

// StartServerClient is the client API for StartServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StartServerClient interface {
	NameNodeStorePlayersMoves(ctx context.Context, in *Playersmoves, opts ...grpc.CallOption) (*Status, error)
	DataNodeStoreMove(ctx context.Context, in *Playermove, opts ...grpc.CallOption) (*Status, error)
}

type startServerClient struct {
	cc grpc.ClientConnInterface
}

func NewStartServerClient(cc grpc.ClientConnInterface) StartServerClient {
	return &startServerClient{cc}
}

func (c *startServerClient) NameNodeStorePlayersMoves(ctx context.Context, in *Playersmoves, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/namenode.StartServer/nameNodeStorePlayersMoves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startServerClient) DataNodeStoreMove(ctx context.Context, in *Playermove, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/namenode.StartServer/dataNodeStoreMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartServerServer is the server API for StartServer service.
// All implementations must embed UnimplementedStartServerServer
// for forward compatibility
type StartServerServer interface {
	NameNodeStorePlayersMoves(context.Context, *Playersmoves) (*Status, error)
	DataNodeStoreMove(context.Context, *Playermove) (*Status, error)
	mustEmbedUnimplementedStartServerServer()
}

// UnimplementedStartServerServer must be embedded to have forward compatible implementations.
type UnimplementedStartServerServer struct {
}

func (UnimplementedStartServerServer) NameNodeStorePlayersMoves(context.Context, *Playersmoves) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NameNodeStorePlayersMoves not implemented")
}
func (UnimplementedStartServerServer) DataNodeStoreMove(context.Context, *Playermove) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataNodeStoreMove not implemented")
}
func (UnimplementedStartServerServer) mustEmbedUnimplementedStartServerServer() {}

// UnsafeStartServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StartServerServer will
// result in compilation errors.
type UnsafeStartServerServer interface {
	mustEmbedUnimplementedStartServerServer()
}

func RegisterStartServerServer(s grpc.ServiceRegistrar, srv StartServerServer) {
	s.RegisterService(&StartServer_ServiceDesc, srv)
}

func _StartServer_NameNodeStorePlayersMoves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playersmoves)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).NameNodeStorePlayersMoves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namenode.StartServer/nameNodeStorePlayersMoves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).NameNodeStorePlayersMoves(ctx, req.(*Playersmoves))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartServer_DataNodeStoreMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playermove)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).DataNodeStoreMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namenode.StartServer/dataNodeStoreMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).DataNodeStoreMove(ctx, req.(*Playermove))
	}
	return interceptor(ctx, in, info, handler)
}

// StartServer_ServiceDesc is the grpc.ServiceDesc for StartServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StartServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "namenode.StartServer",
	HandlerType: (*StartServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "nameNodeStorePlayersMoves",
			Handler:    _StartServer_NameNodeStorePlayersMoves_Handler,
		},
		{
			MethodName: "dataNodeStoreMove",
			Handler:    _StartServer_DataNodeStoreMove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namenode.proto",
}
