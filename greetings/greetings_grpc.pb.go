// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: greetings/greetings.proto

package greetings

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
	GreeterServiceGrpc_Hello_FullMethodName     = "/greetings.GreeterServiceGrpc/Hello"
	GreeterServiceGrpc_HelloWold_FullMethodName = "/greetings.GreeterServiceGrpc/HelloWold"
)

// GreeterServiceGrpcClient is the client API for GreeterServiceGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceGrpcClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloWold(ctx context.Context, in *HelloWoldRequest, opts ...grpc.CallOption) (GreeterServiceGrpc_HelloWoldClient, error)
}

type greeterServiceGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceGrpcClient(cc grpc.ClientConnInterface) GreeterServiceGrpcClient {
	return &greeterServiceGrpcClient{cc}
}

func (c *greeterServiceGrpcClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, GreeterServiceGrpc_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceGrpcClient) HelloWold(ctx context.Context, in *HelloWoldRequest, opts ...grpc.CallOption) (GreeterServiceGrpc_HelloWoldClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterServiceGrpc_ServiceDesc.Streams[0], GreeterServiceGrpc_HelloWold_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceGrpcHelloWoldClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreeterServiceGrpc_HelloWoldClient interface {
	Recv() (*HelloWoldResponse, error)
	grpc.ClientStream
}

type greeterServiceGrpcHelloWoldClient struct {
	grpc.ClientStream
}

func (x *greeterServiceGrpcHelloWoldClient) Recv() (*HelloWoldResponse, error) {
	m := new(HelloWoldResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServiceGrpcServer is the server API for GreeterServiceGrpc service.
// All implementations must embed UnimplementedGreeterServiceGrpcServer
// for forward compatibility
type GreeterServiceGrpcServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloWold(*HelloWoldRequest, GreeterServiceGrpc_HelloWoldServer) error
	mustEmbedUnimplementedGreeterServiceGrpcServer()
}

// UnimplementedGreeterServiceGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceGrpcServer struct {
}

func (UnimplementedGreeterServiceGrpcServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreeterServiceGrpcServer) HelloWold(*HelloWoldRequest, GreeterServiceGrpc_HelloWoldServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWold not implemented")
}
func (UnimplementedGreeterServiceGrpcServer) mustEmbedUnimplementedGreeterServiceGrpcServer() {}

// UnsafeGreeterServiceGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceGrpcServer will
// result in compilation errors.
type UnsafeGreeterServiceGrpcServer interface {
	mustEmbedUnimplementedGreeterServiceGrpcServer()
}

func RegisterGreeterServiceGrpcServer(s grpc.ServiceRegistrar, srv GreeterServiceGrpcServer) {
	s.RegisterService(&GreeterServiceGrpc_ServiceDesc, srv)
}

func _GreeterServiceGrpc_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceGrpcServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreeterServiceGrpc_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceGrpcServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterServiceGrpc_HelloWold_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloWoldRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServiceGrpcServer).HelloWold(m, &greeterServiceGrpcHelloWoldServer{stream})
}

type GreeterServiceGrpc_HelloWoldServer interface {
	Send(*HelloWoldResponse) error
	grpc.ServerStream
}

type greeterServiceGrpcHelloWoldServer struct {
	grpc.ServerStream
}

func (x *greeterServiceGrpcHelloWoldServer) Send(m *HelloWoldResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GreeterServiceGrpc_ServiceDesc is the grpc.ServiceDesc for GreeterServiceGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterServiceGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greetings.GreeterServiceGrpc",
	HandlerType: (*GreeterServiceGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreeterServiceGrpc_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloWold",
			Handler:       _GreeterServiceGrpc_HelloWold_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greetings/greetings.proto",
}