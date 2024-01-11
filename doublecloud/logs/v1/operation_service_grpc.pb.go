// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: doublecloud/logs/v1/operation_service.proto

package logs

import (
	context "context"
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OperationService_Get_FullMethodName = "/doublecloud.logs.v1.OperationService/Get"
)

// OperationServiceClient is the client API for OperationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationServiceClient interface {
	// Returns the specified operation.
	Get(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type operationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationServiceClient(cc grpc.ClientConnInterface) OperationServiceClient {
	return &operationServiceClient{cc}
}

func (c *operationServiceClient) Get(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, OperationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServiceServer is the server API for OperationService service.
// All implementations must embed UnimplementedOperationServiceServer
// for forward compatibility
type OperationServiceServer interface {
	// Returns the specified operation.
	Get(context.Context, *GetOperationRequest) (*v1.Operation, error)
	mustEmbedUnimplementedOperationServiceServer()
}

// UnimplementedOperationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationServiceServer struct {
}

func (UnimplementedOperationServiceServer) Get(context.Context, *GetOperationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOperationServiceServer) mustEmbedUnimplementedOperationServiceServer() {}

// UnsafeOperationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationServiceServer will
// result in compilation errors.
type UnsafeOperationServiceServer interface {
	mustEmbedUnimplementedOperationServiceServer()
}

func RegisterOperationServiceServer(s grpc.ServiceRegistrar, srv OperationServiceServer) {
	s.RegisterService(&OperationService_ServiceDesc, srv)
}

func _OperationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Get(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationService_ServiceDesc is the grpc.ServiceDesc for OperationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.logs.v1.OperationService",
	HandlerType: (*OperationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OperationService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/logs/v1/operation_service.proto",
}