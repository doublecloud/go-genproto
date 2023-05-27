// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: doublecloud/kafka/v1/cluster_service.proto

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
	ClusterService_Get_FullMethodName                   = "/doublecloud.kafka.v1.ClusterService/Get"
	ClusterService_List_FullMethodName                  = "/doublecloud.kafka.v1.ClusterService/List"
	ClusterService_Create_FullMethodName                = "/doublecloud.kafka.v1.ClusterService/Create"
	ClusterService_Update_FullMethodName                = "/doublecloud.kafka.v1.ClusterService/Update"
	ClusterService_Delete_FullMethodName                = "/doublecloud.kafka.v1.ClusterService/Delete"
	ClusterService_ResetCredentials_FullMethodName      = "/doublecloud.kafka.v1.ClusterService/ResetCredentials"
	ClusterService_ListHosts_FullMethodName             = "/doublecloud.kafka.v1.ClusterService/ListHosts"
	ClusterService_Start_FullMethodName                 = "/doublecloud.kafka.v1.ClusterService/Start"
	ClusterService_Stop_FullMethodName                  = "/doublecloud.kafka.v1.ClusterService/Stop"
	ClusterService_RescheduleMaintenance_FullMethodName = "/doublecloud.kafka.v1.ClusterService/RescheduleMaintenance"
	ClusterService_ListOperations_FullMethodName        = "/doublecloud.kafka.v1.ClusterService/ListOperations"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	// Returns the specified Apache Kafka® cluster.
	Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Retrieves a list of Apache Kafka® clusters that belong to the specified
	// project.
	List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Creates an Apache Kafka® cluster in the specified project.
	Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Updates the specified Apache Kafka® cluster.
	Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Deletes the specified Apache Kafka® cluster.
	Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Reset credentials for the specified cluster.
	ResetCredentials(ctx context.Context, in *ResetClusterCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Retrieves a list of hosts for the specified cluster.
	ListHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error)
	// Starts stopped Kafka cluster.
	Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Stops running Kafka cluster.
	Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Reschedule planned maintenance operation.
	RescheduleMaintenance(ctx context.Context, in *RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// List cluster operations.
	ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, ClusterService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ResetCredentials(ctx context.Context, in *ResetClusterCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, ClusterService_ResetCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error) {
	out := new(ListClusterHostsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListHosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) RescheduleMaintenance(ctx context.Context, in *RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, ClusterService_RescheduleMaintenance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error) {
	out := new(ListClusterOperationsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations must embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	// Returns the specified Apache Kafka® cluster.
	Get(context.Context, *GetClusterRequest) (*Cluster, error)
	// Retrieves a list of Apache Kafka® clusters that belong to the specified
	// project.
	List(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Creates an Apache Kafka® cluster in the specified project.
	Create(context.Context, *CreateClusterRequest) (*v1.Operation, error)
	// Updates the specified Apache Kafka® cluster.
	Update(context.Context, *UpdateClusterRequest) (*v1.Operation, error)
	// Deletes the specified Apache Kafka® cluster.
	Delete(context.Context, *DeleteClusterRequest) (*v1.Operation, error)
	// Reset credentials for the specified cluster.
	ResetCredentials(context.Context, *ResetClusterCredentialsRequest) (*v1.Operation, error)
	// Retrieves a list of hosts for the specified cluster.
	ListHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error)
	// Starts stopped Kafka cluster.
	Start(context.Context, *StartClusterRequest) (*v1.Operation, error)
	// Stops running Kafka cluster.
	Stop(context.Context, *StopClusterRequest) (*v1.Operation, error)
	// Reschedule planned maintenance operation.
	RescheduleMaintenance(context.Context, *RescheduleMaintenanceRequest) (*v1.Operation, error)
	// List cluster operations.
	ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error)
	mustEmbedUnimplementedClusterServiceServer()
}

// UnimplementedClusterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) Get(context.Context, *GetClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedClusterServiceServer) List(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedClusterServiceServer) Create(context.Context, *CreateClusterRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClusterServiceServer) Update(context.Context, *UpdateClusterRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClusterServiceServer) Delete(context.Context, *DeleteClusterRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClusterServiceServer) ResetCredentials(context.Context, *ResetClusterCredentialsRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCredentials not implemented")
}
func (UnimplementedClusterServiceServer) ListHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHosts not implemented")
}
func (UnimplementedClusterServiceServer) Start(context.Context, *StartClusterRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedClusterServiceServer) Stop(context.Context, *StopClusterRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedClusterServiceServer) RescheduleMaintenance(context.Context, *RescheduleMaintenanceRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RescheduleMaintenance not implemented")
}
func (UnimplementedClusterServiceServer) ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedClusterServiceServer) mustEmbedUnimplementedClusterServiceServer() {}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Get(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).List(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Create(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Update(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Delete(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ResetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetClusterCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ResetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ResetCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ResetCredentials(ctx, req.(*ResetClusterCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListHosts(ctx, req.(*ListClusterHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Start(ctx, req.(*StartClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Stop(ctx, req.(*StopClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_RescheduleMaintenance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RescheduleMaintenanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).RescheduleMaintenance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_RescheduleMaintenance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).RescheduleMaintenance(ctx, req.(*RescheduleMaintenanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListOperations(ctx, req.(*ListClusterOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.kafka.v1.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ClusterService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClusterService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ClusterService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClusterService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterService_Delete_Handler,
		},
		{
			MethodName: "ResetCredentials",
			Handler:    _ClusterService_ResetCredentials_Handler,
		},
		{
			MethodName: "ListHosts",
			Handler:    _ClusterService_ListHosts_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ClusterService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ClusterService_Stop_Handler,
		},
		{
			MethodName: "RescheduleMaintenance",
			Handler:    _ClusterService_RescheduleMaintenance_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ClusterService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/kafka/v1/cluster_service.proto",
}