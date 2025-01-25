// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: doublecloud/network/v1/network_connection_service.proto

package network

import (
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetNetworkConnectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Network Connection resource to return.
	NetworkConnectionId string `protobuf:"bytes,1,opt,name=network_connection_id,json=networkConnectionId,proto3" json:"network_connection_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *GetNetworkConnectionRequest) Reset() {
	*x = GetNetworkConnectionRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNetworkConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkConnectionRequest) ProtoMessage() {}

func (x *GetNetworkConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkConnectionRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetNetworkConnectionRequest) GetNetworkConnectionId() string {
	if x != nil {
		return x.NetworkConnectionId
	}
	return ""
}

type ListNetworkConnectionsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the project to list Networks Connections in.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Paging information of the request
	Paging        *v1.Paging `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNetworkConnectionsRequest) Reset() {
	*x = ListNetworkConnectionsRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNetworkConnectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworkConnectionsRequest) ProtoMessage() {}

func (x *ListNetworkConnectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworkConnectionsRequest.ProtoReflect.Descriptor instead.
func (*ListNetworkConnectionsRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListNetworkConnectionsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ListNetworkConnectionsRequest) GetPaging() *v1.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type ListNetworkConnectionsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of Network Connection resources.
	NetworkConnections []*NetworkConnection `protobuf:"bytes,1,rep,name=network_connections,json=networkConnections,proto3" json:"network_connections,omitempty"`
	// Pagination information of the response
	NextPage      *v1.NextPage `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNetworkConnectionsResponse) Reset() {
	*x = ListNetworkConnectionsResponse{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNetworkConnectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworkConnectionsResponse) ProtoMessage() {}

func (x *ListNetworkConnectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworkConnectionsResponse.ProtoReflect.Descriptor instead.
func (*ListNetworkConnectionsResponse) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListNetworkConnectionsResponse) GetNetworkConnections() []*NetworkConnection {
	if x != nil {
		return x.NetworkConnections
	}
	return nil
}

func (x *ListNetworkConnectionsResponse) GetNextPage() *v1.NextPage {
	if x != nil {
		return x.NextPage
	}
	return nil
}

type CreateNetworkConnectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Network.
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Provider specific parameters.
	//
	// Types that are valid to be assigned to Params:
	//
	//	*CreateNetworkConnectionRequest_Aws
	//	*CreateNetworkConnectionRequest_Google
	Params isCreateNetworkConnectionRequest_Params `protobuf_oneof:"params"`
	// Description of the Network Connection.
	Description   string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNetworkConnectionRequest) Reset() {
	*x = CreateNetworkConnectionRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNetworkConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkConnectionRequest) ProtoMessage() {}

func (x *CreateNetworkConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateNetworkConnectionRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNetworkConnectionRequest) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *CreateNetworkConnectionRequest) GetParams() isCreateNetworkConnectionRequest_Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *CreateNetworkConnectionRequest) GetAws() *CreateAWSNetworkConnectionRequest {
	if x != nil {
		if x, ok := x.Params.(*CreateNetworkConnectionRequest_Aws); ok {
			return x.Aws
		}
	}
	return nil
}

func (x *CreateNetworkConnectionRequest) GetGoogle() *CreateGoogleNetworkConnectionRequest {
	if x != nil {
		if x, ok := x.Params.(*CreateNetworkConnectionRequest_Google); ok {
			return x.Google
		}
	}
	return nil
}

func (x *CreateNetworkConnectionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type isCreateNetworkConnectionRequest_Params interface {
	isCreateNetworkConnectionRequest_Params()
}

type CreateNetworkConnectionRequest_Aws struct {
	// Connect DoubleCloud Network with AWS VPC.
	Aws *CreateAWSNetworkConnectionRequest `protobuf:"bytes,3,opt,name=aws,proto3,oneof"`
}

type CreateNetworkConnectionRequest_Google struct {
	// Connect DoubleCloud Network with Google Cloud VPC.
	Google *CreateGoogleNetworkConnectionRequest `protobuf:"bytes,4,opt,name=google,proto3,oneof"`
}

func (*CreateNetworkConnectionRequest_Aws) isCreateNetworkConnectionRequest_Params() {}

func (*CreateNetworkConnectionRequest_Google) isCreateNetworkConnectionRequest_Params() {}

type CreateAWSNetworkConnectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*CreateAWSNetworkConnectionRequest_Peering
	Type          isCreateAWSNetworkConnectionRequest_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAWSNetworkConnectionRequest) Reset() {
	*x = CreateAWSNetworkConnectionRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAWSNetworkConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAWSNetworkConnectionRequest) ProtoMessage() {}

func (x *CreateAWSNetworkConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAWSNetworkConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateAWSNetworkConnectionRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAWSNetworkConnectionRequest) GetType() isCreateAWSNetworkConnectionRequest_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *CreateAWSNetworkConnectionRequest) GetPeering() *CreateAWSNetworkConnectionPeeringRequest {
	if x != nil {
		if x, ok := x.Type.(*CreateAWSNetworkConnectionRequest_Peering); ok {
			return x.Peering
		}
	}
	return nil
}

type isCreateAWSNetworkConnectionRequest_Type interface {
	isCreateAWSNetworkConnectionRequest_Type()
}

type CreateAWSNetworkConnectionRequest_Peering struct {
	// Connect AWS VPC using VPC Peering.
	Peering *CreateAWSNetworkConnectionPeeringRequest `protobuf:"bytes,1,opt,name=peering,proto3,oneof"`
}

func (*CreateAWSNetworkConnectionRequest_Peering) isCreateAWSNetworkConnectionRequest_Type() {}

type CreateAWSNetworkConnectionPeeringRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the VPC.
	VpcId string `protobuf:"bytes,1,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	// ID of the VPC owner account.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// ID of the AWS region.
	RegionId string `protobuf:"bytes,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// IPv4 CIDR block. Data Cloud will create route to this CIDR using Peering
	// Connection.
	Ipv4CidrBlock string `protobuf:"bytes,4,opt,name=ipv4_cidr_block,json=ipv4CidrBlock,proto3" json:"ipv4_cidr_block,omitempty"`
	// IPv6 CIDR block. Data Cloud will create route to this CIDR using Peering
	// Connection.
	Ipv6CidrBlock string `protobuf:"bytes,5,opt,name=ipv6_cidr_block,json=ipv6CidrBlock,proto3" json:"ipv6_cidr_block,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAWSNetworkConnectionPeeringRequest) Reset() {
	*x = CreateAWSNetworkConnectionPeeringRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAWSNetworkConnectionPeeringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAWSNetworkConnectionPeeringRequest) ProtoMessage() {}

func (x *CreateAWSNetworkConnectionPeeringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAWSNetworkConnectionPeeringRequest.ProtoReflect.Descriptor instead.
func (*CreateAWSNetworkConnectionPeeringRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAWSNetworkConnectionPeeringRequest) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *CreateAWSNetworkConnectionPeeringRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CreateAWSNetworkConnectionPeeringRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *CreateAWSNetworkConnectionPeeringRequest) GetIpv4CidrBlock() string {
	if x != nil {
		return x.Ipv4CidrBlock
	}
	return ""
}

func (x *CreateAWSNetworkConnectionPeeringRequest) GetIpv6CidrBlock() string {
	if x != nil {
		return x.Ipv6CidrBlock
	}
	return ""
}

type CreateGoogleNetworkConnectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of this peering. The name must comply with RFC1035. Specifically, the name
	// must be 1-63 characters long and match regular expression
	// `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter,
	// and all the following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The URL of the peer network. It can be either full URL or partial URL. Partial
	// URL must contain project.
	PeerNetworkUrl string `protobuf:"bytes,2,opt,name=peer_network_url,json=peerNetworkUrl,proto3" json:"peer_network_url,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateGoogleNetworkConnectionRequest) Reset() {
	*x = CreateGoogleNetworkConnectionRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGoogleNetworkConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoogleNetworkConnectionRequest) ProtoMessage() {}

func (x *CreateGoogleNetworkConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoogleNetworkConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateGoogleNetworkConnectionRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateGoogleNetworkConnectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoogleNetworkConnectionRequest) GetPeerNetworkUrl() string {
	if x != nil {
		return x.PeerNetworkUrl
	}
	return ""
}

type DeleteNetworkConnectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Network Connection to delete.
	// To get the Network Connection ID, use a [NetworkConnectionService.List] request.
	NetworkConnectionId string `protobuf:"bytes,1,opt,name=network_connection_id,json=networkConnectionId,proto3" json:"network_connection_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *DeleteNetworkConnectionRequest) Reset() {
	*x = DeleteNetworkConnectionRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteNetworkConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkConnectionRequest) ProtoMessage() {}

func (x *DeleteNetworkConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkConnectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteNetworkConnectionRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteNetworkConnectionRequest) GetNetworkConnectionId() string {
	if x != nil {
		return x.NetworkConnectionId
	}
	return ""
}

type UpdateNetworkConnectionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Network Connection to update.
	NetworkConnectionId string `protobuf:"bytes,1,opt,name=network_connection_id,json=networkConnectionId,proto3" json:"network_connection_id,omitempty"`
	// Description of the Network Connection.
	Description   string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNetworkConnectionRequest) Reset() {
	*x = UpdateNetworkConnectionRequest{}
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNetworkConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNetworkConnectionRequest) ProtoMessage() {}

func (x *UpdateNetworkConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNetworkConnectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateNetworkConnectionRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateNetworkConnectionRequest) GetNetworkConnectionId() string {
	if x != nil {
		return x.NetworkConnectionId
	}
	return ""
}

func (x *UpdateNetworkConnectionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_doublecloud_network_v1_network_connection_service_proto protoreflect.FileDescriptor

var file_doublecloud_network_v1_network_connection_service_proto_rawDesc = string([]byte{
	0x0a, 0x37, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x1a, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x51, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x6e, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x22, 0xb3, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x35, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x9e, 0x02, 0x0a, 0x1e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x03, 0x61, 0x77,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x56, 0x0a, 0x06, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x04, 0x08,
	0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0x89, 0x01, 0x0a, 0x21, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x5c, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x57, 0x53, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x06, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x28, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x57, 0x53, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x63, 0x69,
	0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x69, 0x70, 0x76, 0x34, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a,
	0x0f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x43, 0x69, 0x64, 0x72,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x64, 0x0a, 0x24, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x65,
	0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x72, 0x6c, 0x22, 0x54, 0x0a, 0x1e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x76, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb2, 0x03, 0x0a, 0x18, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x33, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x75, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x36,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_doublecloud_network_v1_network_connection_service_proto_rawDescOnce sync.Once
	file_doublecloud_network_v1_network_connection_service_proto_rawDescData []byte
)

func file_doublecloud_network_v1_network_connection_service_proto_rawDescGZIP() []byte {
	file_doublecloud_network_v1_network_connection_service_proto_rawDescOnce.Do(func() {
		file_doublecloud_network_v1_network_connection_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_network_v1_network_connection_service_proto_rawDesc), len(file_doublecloud_network_v1_network_connection_service_proto_rawDesc)))
	})
	return file_doublecloud_network_v1_network_connection_service_proto_rawDescData
}

var file_doublecloud_network_v1_network_connection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_doublecloud_network_v1_network_connection_service_proto_goTypes = []any{
	(*GetNetworkConnectionRequest)(nil),              // 0: doublecloud.network.v1.GetNetworkConnectionRequest
	(*ListNetworkConnectionsRequest)(nil),            // 1: doublecloud.network.v1.ListNetworkConnectionsRequest
	(*ListNetworkConnectionsResponse)(nil),           // 2: doublecloud.network.v1.ListNetworkConnectionsResponse
	(*CreateNetworkConnectionRequest)(nil),           // 3: doublecloud.network.v1.CreateNetworkConnectionRequest
	(*CreateAWSNetworkConnectionRequest)(nil),        // 4: doublecloud.network.v1.CreateAWSNetworkConnectionRequest
	(*CreateAWSNetworkConnectionPeeringRequest)(nil), // 5: doublecloud.network.v1.CreateAWSNetworkConnectionPeeringRequest
	(*CreateGoogleNetworkConnectionRequest)(nil),     // 6: doublecloud.network.v1.CreateGoogleNetworkConnectionRequest
	(*DeleteNetworkConnectionRequest)(nil),           // 7: doublecloud.network.v1.DeleteNetworkConnectionRequest
	(*UpdateNetworkConnectionRequest)(nil),           // 8: doublecloud.network.v1.UpdateNetworkConnectionRequest
	(*v1.Paging)(nil),                                // 9: doublecloud.v1.Paging
	(*NetworkConnection)(nil),                        // 10: doublecloud.network.v1.NetworkConnection
	(*v1.NextPage)(nil),                              // 11: doublecloud.v1.NextPage
	(*v1.Operation)(nil),                             // 12: doublecloud.v1.Operation
}
var file_doublecloud_network_v1_network_connection_service_proto_depIdxs = []int32{
	9,  // 0: doublecloud.network.v1.ListNetworkConnectionsRequest.paging:type_name -> doublecloud.v1.Paging
	10, // 1: doublecloud.network.v1.ListNetworkConnectionsResponse.network_connections:type_name -> doublecloud.network.v1.NetworkConnection
	11, // 2: doublecloud.network.v1.ListNetworkConnectionsResponse.next_page:type_name -> doublecloud.v1.NextPage
	4,  // 3: doublecloud.network.v1.CreateNetworkConnectionRequest.aws:type_name -> doublecloud.network.v1.CreateAWSNetworkConnectionRequest
	6,  // 4: doublecloud.network.v1.CreateNetworkConnectionRequest.google:type_name -> doublecloud.network.v1.CreateGoogleNetworkConnectionRequest
	5,  // 5: doublecloud.network.v1.CreateAWSNetworkConnectionRequest.peering:type_name -> doublecloud.network.v1.CreateAWSNetworkConnectionPeeringRequest
	0,  // 6: doublecloud.network.v1.NetworkConnectionService.Get:input_type -> doublecloud.network.v1.GetNetworkConnectionRequest
	1,  // 7: doublecloud.network.v1.NetworkConnectionService.List:input_type -> doublecloud.network.v1.ListNetworkConnectionsRequest
	3,  // 8: doublecloud.network.v1.NetworkConnectionService.Create:input_type -> doublecloud.network.v1.CreateNetworkConnectionRequest
	7,  // 9: doublecloud.network.v1.NetworkConnectionService.Delete:input_type -> doublecloud.network.v1.DeleteNetworkConnectionRequest
	10, // 10: doublecloud.network.v1.NetworkConnectionService.Get:output_type -> doublecloud.network.v1.NetworkConnection
	2,  // 11: doublecloud.network.v1.NetworkConnectionService.List:output_type -> doublecloud.network.v1.ListNetworkConnectionsResponse
	12, // 12: doublecloud.network.v1.NetworkConnectionService.Create:output_type -> doublecloud.v1.Operation
	12, // 13: doublecloud.network.v1.NetworkConnectionService.Delete:output_type -> doublecloud.v1.Operation
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_doublecloud_network_v1_network_connection_service_proto_init() }
func file_doublecloud_network_v1_network_connection_service_proto_init() {
	if File_doublecloud_network_v1_network_connection_service_proto != nil {
		return
	}
	file_doublecloud_network_v1_network_connection_proto_init()
	file_doublecloud_network_v1_network_connection_service_proto_msgTypes[3].OneofWrappers = []any{
		(*CreateNetworkConnectionRequest_Aws)(nil),
		(*CreateNetworkConnectionRequest_Google)(nil),
	}
	file_doublecloud_network_v1_network_connection_service_proto_msgTypes[4].OneofWrappers = []any{
		(*CreateAWSNetworkConnectionRequest_Peering)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_network_v1_network_connection_service_proto_rawDesc), len(file_doublecloud_network_v1_network_connection_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doublecloud_network_v1_network_connection_service_proto_goTypes,
		DependencyIndexes: file_doublecloud_network_v1_network_connection_service_proto_depIdxs,
		MessageInfos:      file_doublecloud_network_v1_network_connection_service_proto_msgTypes,
	}.Build()
	File_doublecloud_network_v1_network_connection_service_proto = out.File
	file_doublecloud_network_v1_network_connection_service_proto_goTypes = nil
	file_doublecloud_network_v1_network_connection_service_proto_depIdxs = nil
}
