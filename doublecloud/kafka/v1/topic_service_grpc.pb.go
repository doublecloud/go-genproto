// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: doublecloud/kafka/v1/topic_service.proto

package kafka

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
	TopicService_Get_FullMethodName    = "/doublecloud.kafka.v1.TopicService/Get"
	TopicService_List_FullMethodName   = "/doublecloud.kafka.v1.TopicService/List"
	TopicService_Create_FullMethodName = "/doublecloud.kafka.v1.TopicService/Create"
	TopicService_Update_FullMethodName = "/doublecloud.kafka.v1.TopicService/Update"
	TopicService_Delete_FullMethodName = "/doublecloud.kafka.v1.TopicService/Delete"
)

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicServiceClient interface {
	// Returns the specified Apache Kafka Topic resource.
	//
	// To get the list of available Apache Kafka Topic resources, make a [List]
	// request.
	Get(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	// Retrieves the list of Apache Kafka Topic resources in the specified cluster.
	List(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	// Creates a new Apache Kafka topic in the specified cluster.
	Create(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Updates the specified Apache Kafka topic.
	Update(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Deletes the specified Apache Kafka topic.
	Delete(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type topicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicServiceClient(cc grpc.ClientConnInterface) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) Get(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, TopicService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) List(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, TopicService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) Create(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, TopicService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) Update(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, TopicService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) Delete(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, TopicService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
// All implementations must embed UnimplementedTopicServiceServer
// for forward compatibility
type TopicServiceServer interface {
	// Returns the specified Apache Kafka Topic resource.
	//
	// To get the list of available Apache Kafka Topic resources, make a [List]
	// request.
	Get(context.Context, *GetTopicRequest) (*Topic, error)
	// Retrieves the list of Apache Kafka Topic resources in the specified cluster.
	List(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	// Creates a new Apache Kafka topic in the specified cluster.
	Create(context.Context, *CreateTopicRequest) (*v1.Operation, error)
	// Updates the specified Apache Kafka topic.
	Update(context.Context, *UpdateTopicRequest) (*v1.Operation, error)
	// Deletes the specified Apache Kafka topic.
	Delete(context.Context, *DeleteTopicRequest) (*v1.Operation, error)
	mustEmbedUnimplementedTopicServiceServer()
}

// UnimplementedTopicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (UnimplementedTopicServiceServer) Get(context.Context, *GetTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTopicServiceServer) List(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTopicServiceServer) Create(context.Context, *CreateTopicRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTopicServiceServer) Update(context.Context, *UpdateTopicRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTopicServiceServer) Delete(context.Context, *DeleteTopicRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTopicServiceServer) mustEmbedUnimplementedTopicServiceServer() {}

// UnsafeTopicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServiceServer will
// result in compilation errors.
type UnsafeTopicServiceServer interface {
	mustEmbedUnimplementedTopicServiceServer()
}

func RegisterTopicServiceServer(s grpc.ServiceRegistrar, srv TopicServiceServer) {
	s.RegisterService(&TopicService_ServiceDesc, srv)
}

func _TopicService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).Get(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).List(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).Create(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).Update(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopicService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).Delete(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TopicService_ServiceDesc is the grpc.ServiceDesc for TopicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.kafka.v1.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TopicService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TopicService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TopicService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TopicService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TopicService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/kafka/v1/topic_service.proto",
}
