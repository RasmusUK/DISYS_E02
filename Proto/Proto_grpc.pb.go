// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Proto

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

// BenchPressQueueClient is the client API for BenchPressQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BenchPressQueueClient interface {
	SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type benchPressQueueClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchPressQueueClient(cc grpc.ClientConnInterface) BenchPressQueueClient {
	return &benchPressQueueClient{cc}
}

func (c *benchPressQueueClient) SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Proto.BenchPressQueue/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BenchPressQueueServer is the server API for BenchPressQueue service.
// All implementations must embed UnimplementedBenchPressQueueServer
// for forward compatibility
type BenchPressQueueServer interface {
	SendRequest(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedBenchPressQueueServer()
}

// UnimplementedBenchPressQueueServer must be embedded to have forward compatible implementations.
type UnimplementedBenchPressQueueServer struct {
}

func (UnimplementedBenchPressQueueServer) SendRequest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedBenchPressQueueServer) mustEmbedUnimplementedBenchPressQueueServer() {}

// UnsafeBenchPressQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BenchPressQueueServer will
// result in compilation errors.
type UnsafeBenchPressQueueServer interface {
	mustEmbedUnimplementedBenchPressQueueServer()
}

func RegisterBenchPressQueueServer(s grpc.ServiceRegistrar, srv BenchPressQueueServer) {
	s.RegisterService(&BenchPressQueue_ServiceDesc, srv)
}

func _BenchPressQueue_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchPressQueueServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.BenchPressQueue/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchPressQueueServer).SendRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// BenchPressQueue_ServiceDesc is the grpc.ServiceDesc for BenchPressQueue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BenchPressQueue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.BenchPressQueue",
	HandlerType: (*BenchPressQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _BenchPressQueue_SendRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/Proto.proto",
}
