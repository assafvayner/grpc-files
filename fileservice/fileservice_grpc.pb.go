// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fileservice

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

// RetreiverClient is the client API for Retreiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetreiverClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
}

type retreiverClient struct {
	cc grpc.ClientConnInterface
}

func NewRetreiverClient(cc grpc.ClientConnInterface) RetreiverClient {
	return &retreiverClient{cc}
}

func (c *retreiverClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/files.Retreiver/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retreiverClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, "/files.Retreiver/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retreiverClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/files.Retreiver/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetreiverServer is the server API for Retreiver service.
// All implementations must embed UnimplementedRetreiverServer
// for forward compatibility
type RetreiverServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	Push(context.Context, *PushRequest) (*PushResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	mustEmbedUnimplementedRetreiverServer()
}

// UnimplementedRetreiverServer must be embedded to have forward compatible implementations.
type UnimplementedRetreiverServer struct {
}

func (UnimplementedRetreiverServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedRetreiverServer) Push(context.Context, *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedRetreiverServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedRetreiverServer) mustEmbedUnimplementedRetreiverServer() {}

// UnsafeRetreiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetreiverServer will
// result in compilation errors.
type UnsafeRetreiverServer interface {
	mustEmbedUnimplementedRetreiverServer()
}

func RegisterRetreiverServer(s grpc.ServiceRegistrar, srv RetreiverServer) {
	s.RegisterService(&Retreiver_ServiceDesc, srv)
}

func _Retreiver_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetreiverServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/files.Retreiver/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetreiverServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retreiver_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetreiverServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/files.Retreiver/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetreiverServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retreiver_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetreiverServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/files.Retreiver/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetreiverServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Retreiver_ServiceDesc is the grpc.ServiceDesc for Retreiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Retreiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "files.Retreiver",
	HandlerType: (*RetreiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Retreiver_Fetch_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Retreiver_Push_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Retreiver_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fileservice/fileservice.proto",
}
