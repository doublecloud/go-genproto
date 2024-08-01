// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: doublecloud/organizationmanager/saml/v1/federation_service.proto

package saml

import (
	context "context"
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FederationService_Get_FullMethodName    = "/doublecloud.organizationmanager.saml.v1.FederationService/Get"
	FederationService_Create_FullMethodName = "/doublecloud.organizationmanager.saml.v1.FederationService/Create"
	FederationService_Update_FullMethodName = "/doublecloud.organizationmanager.saml.v1.FederationService/Update"
	FederationService_Delete_FullMethodName = "/doublecloud.organizationmanager.saml.v1.FederationService/Delete"
)

// FederationServiceClient is the client API for FederationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing federations.
type FederationServiceClient interface {
	// Returns the specified federation.
	//
	// To get the list of available federations, make a [List] request.
	Get(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*Federation, error)
	// Creates a federation in the specified organization.
	Create(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Updates the specified federation.
	Update(ctx context.Context, in *UpdateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Deletes the specified federation.
	Delete(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type federationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFederationServiceClient(cc grpc.ClientConnInterface) FederationServiceClient {
	return &federationServiceClient{cc}
}

func (c *federationServiceClient) Get(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*Federation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Federation)
	err := c.cc.Invoke(ctx, FederationService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Create(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederationService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Update(ctx context.Context, in *UpdateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederationService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Delete(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederationService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederationServiceServer is the server API for FederationService service.
// All implementations must embed UnimplementedFederationServiceServer
// for forward compatibility.
//
// A set of methods for managing federations.
type FederationServiceServer interface {
	// Returns the specified federation.
	//
	// To get the list of available federations, make a [List] request.
	Get(context.Context, *GetFederationRequest) (*Federation, error)
	// Creates a federation in the specified organization.
	Create(context.Context, *CreateFederationRequest) (*v1.Operation, error)
	// Updates the specified federation.
	Update(context.Context, *UpdateFederationRequest) (*v1.Operation, error)
	// Deletes the specified federation.
	Delete(context.Context, *DeleteFederationRequest) (*v1.Operation, error)
	mustEmbedUnimplementedFederationServiceServer()
}

// UnimplementedFederationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFederationServiceServer struct{}

func (UnimplementedFederationServiceServer) Get(context.Context, *GetFederationRequest) (*Federation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFederationServiceServer) Create(context.Context, *CreateFederationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFederationServiceServer) Update(context.Context, *UpdateFederationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFederationServiceServer) Delete(context.Context, *DeleteFederationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFederationServiceServer) mustEmbedUnimplementedFederationServiceServer() {}
func (UnimplementedFederationServiceServer) testEmbeddedByValue()                           {}

// UnsafeFederationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FederationServiceServer will
// result in compilation errors.
type UnsafeFederationServiceServer interface {
	mustEmbedUnimplementedFederationServiceServer()
}

func RegisterFederationServiceServer(s grpc.ServiceRegistrar, srv FederationServiceServer) {
	// If the following call pancis, it indicates UnimplementedFederationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FederationService_ServiceDesc, srv)
}

func _FederationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Get(ctx, req.(*GetFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Create(ctx, req.(*CreateFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Update(ctx, req.(*UpdateFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Delete(ctx, req.(*DeleteFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FederationService_ServiceDesc is the grpc.ServiceDesc for FederationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FederationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.organizationmanager.saml.v1.FederationService",
	HandlerType: (*FederationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FederationService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FederationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FederationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FederationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/organizationmanager/saml/v1/federation_service.proto",
}
