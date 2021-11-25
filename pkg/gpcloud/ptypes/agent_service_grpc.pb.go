// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ptypes

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

// NetworkAgentServiceClient is the client API for NetworkAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkAgentServiceClient interface {
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error)
	QueueJob(ctx context.Context, in *QueueJobRequest, opts ...grpc.CallOption) (*QueueJobResponse, error)
}

type networkAgentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkAgentServiceClient(cc grpc.ClientConnInterface) NetworkAgentServiceClient {
	return &networkAgentServiceClient{cc}
}

func (c *networkAgentServiceClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error) {
	out := new(GetJobResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkAgentService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkAgentServiceClient) QueueJob(ctx context.Context, in *QueueJobRequest, opts ...grpc.CallOption) (*QueueJobResponse, error) {
	out := new(QueueJobResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkAgentService/QueueJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkAgentServiceServer is the server API for NetworkAgentService service.
// All implementations must embed UnimplementedNetworkAgentServiceServer
// for forward compatibility
type NetworkAgentServiceServer interface {
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)
	QueueJob(context.Context, *QueueJobRequest) (*QueueJobResponse, error)
	mustEmbedUnimplementedNetworkAgentServiceServer()
}

// UnimplementedNetworkAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkAgentServiceServer struct {
}

func (UnimplementedNetworkAgentServiceServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedNetworkAgentServiceServer) QueueJob(context.Context, *QueueJobRequest) (*QueueJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueJob not implemented")
}
func (UnimplementedNetworkAgentServiceServer) mustEmbedUnimplementedNetworkAgentServiceServer() {}

// UnsafeNetworkAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkAgentServiceServer will
// result in compilation errors.
type UnsafeNetworkAgentServiceServer interface {
	mustEmbedUnimplementedNetworkAgentServiceServer()
}

func RegisterNetworkAgentServiceServer(s grpc.ServiceRegistrar, srv NetworkAgentServiceServer) {
	s.RegisterService(&NetworkAgentService_ServiceDesc, srv)
}

func _NetworkAgentService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAgentServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkAgentService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAgentServiceServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkAgentService_QueueJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkAgentServiceServer).QueueJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkAgentService/QueueJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkAgentServiceServer).QueueJob(ctx, req.(*QueueJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkAgentService_ServiceDesc is the grpc.ServiceDesc for NetworkAgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkAgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.NetworkAgentService",
	HandlerType: (*NetworkAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _NetworkAgentService_GetJob_Handler,
		},
		{
			MethodName: "QueueJob",
			Handler:    _NetworkAgentService_QueueJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ptypes/agent_service.proto",
}
