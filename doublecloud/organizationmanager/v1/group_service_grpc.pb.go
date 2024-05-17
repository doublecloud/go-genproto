// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: doublecloud/organizationmanager/v1/group_service.proto

package organizationmanager

import (
	context "context"
	v11 "github.com/doublecloud/go-genproto/doublecloud/access/v1"
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
	GroupService_Get_FullMethodName                  = "/doublecloud.organizationmanager.v1.GroupService/Get"
	GroupService_List_FullMethodName                 = "/doublecloud.organizationmanager.v1.GroupService/List"
	GroupService_Create_FullMethodName               = "/doublecloud.organizationmanager.v1.GroupService/Create"
	GroupService_Update_FullMethodName               = "/doublecloud.organizationmanager.v1.GroupService/Update"
	GroupService_Delete_FullMethodName               = "/doublecloud.organizationmanager.v1.GroupService/Delete"
	GroupService_ListOperations_FullMethodName       = "/doublecloud.organizationmanager.v1.GroupService/ListOperations"
	GroupService_ListMembers_FullMethodName          = "/doublecloud.organizationmanager.v1.GroupService/ListMembers"
	GroupService_UpdateMembers_FullMethodName        = "/doublecloud.organizationmanager.v1.GroupService/UpdateMembers"
	GroupService_ListAccessBindings_FullMethodName   = "/doublecloud.organizationmanager.v1.GroupService/ListAccessBindings"
	GroupService_SetAccessBindings_FullMethodName    = "/doublecloud.organizationmanager.v1.GroupService/SetAccessBindings"
	GroupService_UpdateAccessBindings_FullMethodName = "/doublecloud.organizationmanager.v1.GroupService/UpdateAccessBindings"
)

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	// Returns the specified Group resource.
	//
	// To get the list of available Group resources, make a [List] request.
	Get(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// Retrieves the list of group resources.
	List(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	// Creates a group in the specified organization.
	Create(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Updates the specified group.
	Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Deletes the specified group.
	Delete(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Lists operations for the specified group.
	ListOperations(ctx context.Context, in *ListGroupOperationsRequest, opts ...grpc.CallOption) (*ListGroupOperationsResponse, error)
	// List group active members.
	ListMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error)
	// Update group members.
	UpdateMembers(ctx context.Context, in *UpdateGroupMembersRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Lists access bindings for the specified group.
	ListAccessBindings(ctx context.Context, in *v11.ListAccessBindingsRequest, opts ...grpc.CallOption) (*v11.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified group.
	SetAccessBindings(ctx context.Context, in *v11.SetAccessBindingsRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Updates access bindings for the specified group.
	UpdateAccessBindings(ctx context.Context, in *v11.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) Get(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, GroupService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) List(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, GroupService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Create(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Delete(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListOperations(ctx context.Context, in *ListGroupOperationsRequest, opts ...grpc.CallOption) (*ListGroupOperationsResponse, error) {
	out := new(ListGroupOperationsResponse)
	err := c.cc.Invoke(ctx, GroupService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error) {
	out := new(ListGroupMembersResponse)
	err := c.cc.Invoke(ctx, GroupService_ListMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateMembers(ctx context.Context, in *UpdateGroupMembersRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupService_UpdateMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListAccessBindings(ctx context.Context, in *v11.ListAccessBindingsRequest, opts ...grpc.CallOption) (*v11.ListAccessBindingsResponse, error) {
	out := new(v11.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, GroupService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) SetAccessBindings(ctx context.Context, in *v11.SetAccessBindingsRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateAccessBindings(ctx context.Context, in *v11.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	// Returns the specified Group resource.
	//
	// To get the list of available Group resources, make a [List] request.
	Get(context.Context, *GetGroupRequest) (*Group, error)
	// Retrieves the list of group resources.
	List(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	// Creates a group in the specified organization.
	Create(context.Context, *CreateGroupRequest) (*v1.Operation, error)
	// Updates the specified group.
	Update(context.Context, *UpdateGroupRequest) (*v1.Operation, error)
	// Deletes the specified group.
	Delete(context.Context, *DeleteGroupRequest) (*v1.Operation, error)
	// Lists operations for the specified group.
	ListOperations(context.Context, *ListGroupOperationsRequest) (*ListGroupOperationsResponse, error)
	// List group active members.
	ListMembers(context.Context, *ListGroupMembersRequest) (*ListGroupMembersResponse, error)
	// Update group members.
	UpdateMembers(context.Context, *UpdateGroupMembersRequest) (*v1.Operation, error)
	// Lists access bindings for the specified group.
	ListAccessBindings(context.Context, *v11.ListAccessBindingsRequest) (*v11.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified group.
	SetAccessBindings(context.Context, *v11.SetAccessBindingsRequest) (*v1.Operation, error)
	// Updates access bindings for the specified group.
	UpdateAccessBindings(context.Context, *v11.UpdateAccessBindingsRequest) (*v1.Operation, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) Get(context.Context, *GetGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGroupServiceServer) List(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGroupServiceServer) Create(context.Context, *CreateGroupRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGroupServiceServer) Update(context.Context, *UpdateGroupRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGroupServiceServer) Delete(context.Context, *DeleteGroupRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGroupServiceServer) ListOperations(context.Context, *ListGroupOperationsRequest) (*ListGroupOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedGroupServiceServer) ListMembers(context.Context, *ListGroupMembersRequest) (*ListGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedGroupServiceServer) UpdateMembers(context.Context, *UpdateGroupMembersRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMembers not implemented")
}
func (UnimplementedGroupServiceServer) ListAccessBindings(context.Context, *v11.ListAccessBindingsRequest) (*v11.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedGroupServiceServer) SetAccessBindings(context.Context, *v11.SetAccessBindingsRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedGroupServiceServer) UpdateAccessBindings(context.Context, *v11.UpdateAccessBindingsRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Get(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).List(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Create(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Update(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Delete(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListOperations(ctx, req.(*ListGroupOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_ListMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListMembers(ctx, req.(*ListGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_UpdateMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateMembers(ctx, req.(*UpdateGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListAccessBindings(ctx, req.(*v11.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).SetAccessBindings(ctx, req.(*v11.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateAccessBindings(ctx, req.(*v11.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doublecloud.organizationmanager.v1.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GroupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GroupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _GroupService_ListOperations_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _GroupService_ListMembers_Handler,
		},
		{
			MethodName: "UpdateMembers",
			Handler:    _GroupService_UpdateMembers_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _GroupService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _GroupService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _GroupService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doublecloud/organizationmanager/v1/group_service.proto",
}
