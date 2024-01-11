// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: doublecloud/logs/v1/log_export_service.proto

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
	LogExportService_Get_FullMethodName    = "/doublecloud.logs.v1.LogExportService/Get"
	LogExportService_List_FullMethodName   = "/doublecloud.logs.v1.LogExportService/List"
	LogExportService_Create_FullMethodName = "/doublecloud.logs.v1.LogExportService/Create"
	LogExportService_Update_FullMethodName = "/doublecloud.logs.v1.LogExportService/Update"
	LogExportService_Delete_FullMethodName = "/doublecloud.logs.v1.LogExportService/Delete"
)

// LogExportServiceClient is the client API for LogExportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogExportServiceClient interface {
	Get(ctx context.Context, in *GetExportRequest, opts ...grpc.CallOption) (*LogsExport, error)
	List(ctx context.Context, in *ListExportRequest, opts ...grpc.CallOption) (*ListExportResponse, error)
	Create(ctx context.Context, in *CreateExportRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateExportRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteExportRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type logExportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogExportServiceClient(cc grpc.ClientConnInterface) LogExportServiceClient {
	return &logExportServiceClient{cc}
}

func (c *logExportServiceClient) Get(ctx context.Context, in *GetExportRequest, opts ...grpc.CallOption) (*LogsExport, error) {
	out := new(LogsExport)
	err := c.cc.Invoke(ctx, LogExportService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logExportServiceClient) List(ctx context.Context, in *ListExportRequest, opts ...grpc.CallOption) (*ListExportResponse, error) {
	out := new(ListExportResponse)
	err := c.cc.Invoke(ctx, LogExportService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logExportServiceClient) Create(ctx context.Context, in *CreateExportRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, LogExportService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logExportServiceClient) Update(ctx context.Context, in *UpdateExportRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, LogExportService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logExportServiceClient) Delete(ctx context.Context, in *DeleteExportRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, LogExportService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogExportServiceServer is the server API for LogExportService service.
// All implementations must embed UnimplementedLogExportServiceServer
// for forward compatibility
type LogExportServiceServer interface {
	Get(context.Context, *GetExportRequest) (*LogsExport, error)
	List(context.Context, *ListExportRequest) (*ListExportResponse, error)
	Create(context.Context, *CreateExportRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateExportRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteExportRequest) (*v1.Operation, error)
	mustEmbedUnimplementedLogExportServiceServer()
}

// UnimplementedLogExportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogExportServiceServer struct {
}

func (UnimplementedLogExportServiceServer) Get(context.Context, *GetExportRequest) (*LogsExport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLogExportServiceServer) List(context.Context, *ListExportRequest) (*ListExportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLogExportServiceServer) Create(context.Context, *CreateExportRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLogExportServiceServer) Update(context.Context, *UpdateExportRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLogExportServiceServer) Delete(context.Context, *DeleteExportRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLogExportServiceServer) mustEmbedUnimplementedLogExportServiceServer() {}

// UnsafeLogExportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogExportServiceServer will
// result in compilation errors.
type UnsafeLogExportServiceServer interface {
	mustEmbedUnimplementedLogExportServiceServer()
}

func RegisterLogExportServiceServer(s grpc.ServiceRegistrar, srv LogExportServiceServer) {
	s.RegisterService(&LogExportService_ServiceDesc, srv)
}

func _LogExportService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogExportServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogExportService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogExportServiceServer).Get(ctx, req.(*GetExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogExportService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogExportServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogExportService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogExportServiceServer).List(ctx, req.(*ListExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogExportService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogExportServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogExportService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogExportServiceServer).Create(ctx, req.(*CreateExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogExportService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogExportServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogExportService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogExportServiceServer).Update(ctx, req.(*UpdateExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogExportService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogExportServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogExportService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogExportServiceServer).Delete(ctx, req.(*DeleteExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogExportService_ServiceDesc is the grpc.ServiceDesc for LogExportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogExportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.logs.v1.LogExportService",
	HandlerType: (*LogExportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _LogExportService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LogExportService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LogExportService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LogExportService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LogExportService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/logs/v1/log_export_service.proto",
}