// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: doublecloud/visualization/v1/workbook_service.proto

package visualization

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
	WorkbookService_Get_FullMethodName                 = "/doublecloud.visualization.v1.WorkbookService/Get"
	WorkbookService_Create_FullMethodName              = "/doublecloud.visualization.v1.WorkbookService/Create"
	WorkbookService_Update_FullMethodName              = "/doublecloud.visualization.v1.WorkbookService/Update"
	WorkbookService_Delete_FullMethodName              = "/doublecloud.visualization.v1.WorkbookService/Delete"
	WorkbookService_GetConnection_FullMethodName       = "/doublecloud.visualization.v1.WorkbookService/GetConnection"
	WorkbookService_CreateConnection_FullMethodName    = "/doublecloud.visualization.v1.WorkbookService/CreateConnection"
	WorkbookService_UpdateConnection_FullMethodName    = "/doublecloud.visualization.v1.WorkbookService/UpdateConnection"
	WorkbookService_DeleteConnection_FullMethodName    = "/doublecloud.visualization.v1.WorkbookService/DeleteConnection"
	WorkbookService_AdviseDatasetFields_FullMethodName = "/doublecloud.visualization.v1.WorkbookService/AdviseDatasetFields"
)

// WorkbookServiceClient is the client API for WorkbookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkbookServiceClient interface {
	// Workbook operations
	Get(ctx context.Context, in *GetWorkbookRequest, opts ...grpc.CallOption) (*GetWorkbookResponse, error)
	Create(ctx context.Context, in *CreateWorkbookRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateWorkbookRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteWorkbookRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Connections operations
	GetConnection(ctx context.Context, in *GetWorkbookConnectionRequest, opts ...grpc.CallOption) (*GetWorkbookConnectionResponse, error)
	CreateConnection(ctx context.Context, in *CreateWorkbookConnectionRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	UpdateConnection(ctx context.Context, in *UpdateWorkbookConnectionRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	DeleteConnection(ctx context.Context, in *DeleteWorkbookConnectionRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// List all direct fields of dataset with configured data source
	AdviseDatasetFields(ctx context.Context, in *AdviseDatasetFieldsRequest, opts ...grpc.CallOption) (*AdviseDatasetFieldsResponse, error)
}

type workbookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkbookServiceClient(cc grpc.ClientConnInterface) WorkbookServiceClient {
	return &workbookServiceClient{cc}
}

func (c *workbookServiceClient) Get(ctx context.Context, in *GetWorkbookRequest, opts ...grpc.CallOption) (*GetWorkbookResponse, error) {
	out := new(GetWorkbookResponse)
	err := c.cc.Invoke(ctx, WorkbookService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) Create(ctx context.Context, in *CreateWorkbookRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, WorkbookService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) Update(ctx context.Context, in *UpdateWorkbookRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, WorkbookService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) Delete(ctx context.Context, in *DeleteWorkbookRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, WorkbookService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) GetConnection(ctx context.Context, in *GetWorkbookConnectionRequest, opts ...grpc.CallOption) (*GetWorkbookConnectionResponse, error) {
	out := new(GetWorkbookConnectionResponse)
	err := c.cc.Invoke(ctx, WorkbookService_GetConnection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) CreateConnection(ctx context.Context, in *CreateWorkbookConnectionRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, WorkbookService_CreateConnection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) UpdateConnection(ctx context.Context, in *UpdateWorkbookConnectionRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, WorkbookService_UpdateConnection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) DeleteConnection(ctx context.Context, in *DeleteWorkbookConnectionRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, WorkbookService_DeleteConnection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workbookServiceClient) AdviseDatasetFields(ctx context.Context, in *AdviseDatasetFieldsRequest, opts ...grpc.CallOption) (*AdviseDatasetFieldsResponse, error) {
	out := new(AdviseDatasetFieldsResponse)
	err := c.cc.Invoke(ctx, WorkbookService_AdviseDatasetFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkbookServiceServer is the server API for WorkbookService service.
// All implementations must embed UnimplementedWorkbookServiceServer
// for forward compatibility
type WorkbookServiceServer interface {
	// Workbook operations
	Get(context.Context, *GetWorkbookRequest) (*GetWorkbookResponse, error)
	Create(context.Context, *CreateWorkbookRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateWorkbookRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteWorkbookRequest) (*v1.Operation, error)
	// Connections operations
	GetConnection(context.Context, *GetWorkbookConnectionRequest) (*GetWorkbookConnectionResponse, error)
	CreateConnection(context.Context, *CreateWorkbookConnectionRequest) (*v1.Operation, error)
	UpdateConnection(context.Context, *UpdateWorkbookConnectionRequest) (*v1.Operation, error)
	DeleteConnection(context.Context, *DeleteWorkbookConnectionRequest) (*v1.Operation, error)
	// List all direct fields of dataset with configured data source
	AdviseDatasetFields(context.Context, *AdviseDatasetFieldsRequest) (*AdviseDatasetFieldsResponse, error)
	mustEmbedUnimplementedWorkbookServiceServer()
}

// UnimplementedWorkbookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkbookServiceServer struct {
}

func (UnimplementedWorkbookServiceServer) Get(context.Context, *GetWorkbookRequest) (*GetWorkbookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWorkbookServiceServer) Create(context.Context, *CreateWorkbookRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWorkbookServiceServer) Update(context.Context, *UpdateWorkbookRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWorkbookServiceServer) Delete(context.Context, *DeleteWorkbookRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWorkbookServiceServer) GetConnection(context.Context, *GetWorkbookConnectionRequest) (*GetWorkbookConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (UnimplementedWorkbookServiceServer) CreateConnection(context.Context, *CreateWorkbookConnectionRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (UnimplementedWorkbookServiceServer) UpdateConnection(context.Context, *UpdateWorkbookConnectionRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnection not implemented")
}
func (UnimplementedWorkbookServiceServer) DeleteConnection(context.Context, *DeleteWorkbookConnectionRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (UnimplementedWorkbookServiceServer) AdviseDatasetFields(context.Context, *AdviseDatasetFieldsRequest) (*AdviseDatasetFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdviseDatasetFields not implemented")
}
func (UnimplementedWorkbookServiceServer) mustEmbedUnimplementedWorkbookServiceServer() {}

// UnsafeWorkbookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkbookServiceServer will
// result in compilation errors.
type UnsafeWorkbookServiceServer interface {
	mustEmbedUnimplementedWorkbookServiceServer()
}

func RegisterWorkbookServiceServer(s grpc.ServiceRegistrar, srv WorkbookServiceServer) {
	s.RegisterService(&WorkbookService_ServiceDesc, srv)
}

func _WorkbookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkbookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).Get(ctx, req.(*GetWorkbookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkbookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).Create(ctx, req.(*CreateWorkbookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkbookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).Update(ctx, req.(*UpdateWorkbookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkbookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).Delete(ctx, req.(*DeleteWorkbookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkbookConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_GetConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).GetConnection(ctx, req.(*GetWorkbookConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkbookConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_CreateConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).CreateConnection(ctx, req.(*CreateWorkbookConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_UpdateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkbookConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).UpdateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_UpdateConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).UpdateConnection(ctx, req.(*UpdateWorkbookConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkbookConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_DeleteConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).DeleteConnection(ctx, req.(*DeleteWorkbookConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkbookService_AdviseDatasetFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdviseDatasetFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkbookServiceServer).AdviseDatasetFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkbookService_AdviseDatasetFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkbookServiceServer).AdviseDatasetFields(ctx, req.(*AdviseDatasetFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkbookService_ServiceDesc is the grpc.ServiceDesc for WorkbookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkbookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.visualization.v1.WorkbookService",
	HandlerType: (*WorkbookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _WorkbookService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _WorkbookService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WorkbookService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WorkbookService_Delete_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _WorkbookService_GetConnection_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _WorkbookService_CreateConnection_Handler,
		},
		{
			MethodName: "UpdateConnection",
			Handler:    _WorkbookService_UpdateConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _WorkbookService_DeleteConnection_Handler,
		},
		{
			MethodName: "AdviseDatasetFields",
			Handler:    _WorkbookService_AdviseDatasetFields_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/visualization/v1/workbook_service.proto",
}
