// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: doublecloud/transfer/v1/endpoint_service.proto

package transfer

import (
	context "context"
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EndpointService_Get_FullMethodName    = "/doublecloud.transfer.v1.EndpointService/Get"
	EndpointService_List_FullMethodName   = "/doublecloud.transfer.v1.EndpointService/List"
	EndpointService_Create_FullMethodName = "/doublecloud.transfer.v1.EndpointService/Create"
	EndpointService_Update_FullMethodName = "/doublecloud.transfer.v1.EndpointService/Update"
	EndpointService_Delete_FullMethodName = "/doublecloud.transfer.v1.EndpointService/Delete"
)

// EndpointServiceClient is the client API for EndpointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointServiceClient interface {
	Get(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error)
	List(ctx context.Context, in *ListEndpointsRequest, opts ...grpc.CallOption) (*ListEndpointsResponse, error)
	Create(ctx context.Context, in *CreateEndpointRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateEndpointRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type endpointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointServiceClient(cc grpc.ClientConnInterface) EndpointServiceClient {
	return &endpointServiceClient{cc}
}

func (c *endpointServiceClient) Get(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*Endpoint, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Endpoint)
	err := c.cc.Invoke(ctx, EndpointService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointServiceClient) List(ctx context.Context, in *ListEndpointsRequest, opts ...grpc.CallOption) (*ListEndpointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEndpointsResponse)
	err := c.cc.Invoke(ctx, EndpointService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointServiceClient) Create(ctx context.Context, in *CreateEndpointRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, EndpointService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointServiceClient) Update(ctx context.Context, in *UpdateEndpointRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, EndpointService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointServiceClient) Delete(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, EndpointService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointServiceServer is the server API for EndpointService service.
// All implementations must embed UnimplementedEndpointServiceServer
// for forward compatibility
type EndpointServiceServer interface {
	Get(context.Context, *GetEndpointRequest) (*Endpoint, error)
	List(context.Context, *ListEndpointsRequest) (*ListEndpointsResponse, error)
	Create(context.Context, *CreateEndpointRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateEndpointRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteEndpointRequest) (*v1.Operation, error)
	mustEmbedUnimplementedEndpointServiceServer()
}

// UnimplementedEndpointServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointServiceServer struct {
}

func (UnimplementedEndpointServiceServer) Get(context.Context, *GetEndpointRequest) (*Endpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEndpointServiceServer) List(context.Context, *ListEndpointsRequest) (*ListEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEndpointServiceServer) Create(context.Context, *CreateEndpointRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEndpointServiceServer) Update(context.Context, *UpdateEndpointRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEndpointServiceServer) Delete(context.Context, *DeleteEndpointRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEndpointServiceServer) mustEmbedUnimplementedEndpointServiceServer() {}

// UnsafeEndpointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointServiceServer will
// result in compilation errors.
type UnsafeEndpointServiceServer interface {
	mustEmbedUnimplementedEndpointServiceServer()
}

func RegisterEndpointServiceServer(s grpc.ServiceRegistrar, srv EndpointServiceServer) {
	s.RegisterService(&EndpointService_ServiceDesc, srv)
}

func _EndpointService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndpointService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServiceServer).Get(ctx, req.(*GetEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndpointService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServiceServer).List(ctx, req.(*ListEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndpointService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServiceServer).Create(ctx, req.(*CreateEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndpointService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServiceServer).Update(ctx, req.(*UpdateEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndpointService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServiceServer).Delete(ctx, req.(*DeleteEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EndpointService_ServiceDesc is the grpc.ServiceDesc for EndpointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndpointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.transfer.v1.EndpointService",
	HandlerType: (*EndpointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _EndpointService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _EndpointService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EndpointService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EndpointService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EndpointService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/transfer/v1/endpoint_service.proto",
}
