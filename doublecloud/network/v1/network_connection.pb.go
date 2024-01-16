// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: doublecloud/network/v1/network_connection.proto

package network

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NetworkConnection_NetworkConnectionStatus int32

const (
	NetworkConnection_NETWORK_CONNECTION_STATUS_INVALID  NetworkConnection_NetworkConnectionStatus = 0
	NetworkConnection_NETWORK_CONNECTION_STATUS_CREATING NetworkConnection_NetworkConnectionStatus = 1
	NetworkConnection_NETWORK_CONNECTION_STATUS_PENDING  NetworkConnection_NetworkConnectionStatus = 2
	NetworkConnection_NETWORK_CONNECTION_STATUS_ACTIVE   NetworkConnection_NetworkConnectionStatus = 3
	NetworkConnection_NETWORK_CONNECTION_STATUS_DELETING NetworkConnection_NetworkConnectionStatus = 4
	NetworkConnection_NETWORK_CONNECTION_STATUS_ERROR    NetworkConnection_NetworkConnectionStatus = 5
)

// Enum value maps for NetworkConnection_NetworkConnectionStatus.
var (
	NetworkConnection_NetworkConnectionStatus_name = map[int32]string{
		0: "NETWORK_CONNECTION_STATUS_INVALID",
		1: "NETWORK_CONNECTION_STATUS_CREATING",
		2: "NETWORK_CONNECTION_STATUS_PENDING",
		3: "NETWORK_CONNECTION_STATUS_ACTIVE",
		4: "NETWORK_CONNECTION_STATUS_DELETING",
		5: "NETWORK_CONNECTION_STATUS_ERROR",
	}
	NetworkConnection_NetworkConnectionStatus_value = map[string]int32{
		"NETWORK_CONNECTION_STATUS_INVALID":  0,
		"NETWORK_CONNECTION_STATUS_CREATING": 1,
		"NETWORK_CONNECTION_STATUS_PENDING":  2,
		"NETWORK_CONNECTION_STATUS_ACTIVE":   3,
		"NETWORK_CONNECTION_STATUS_DELETING": 4,
		"NETWORK_CONNECTION_STATUS_ERROR":    5,
	}
)

func (x NetworkConnection_NetworkConnectionStatus) Enum() *NetworkConnection_NetworkConnectionStatus {
	p := new(NetworkConnection_NetworkConnectionStatus)
	*p = x
	return p
}

func (x NetworkConnection_NetworkConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkConnection_NetworkConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_network_v1_network_connection_proto_enumTypes[0].Descriptor()
}

func (NetworkConnection_NetworkConnectionStatus) Type() protoreflect.EnumType {
	return &file_doublecloud_network_v1_network_connection_proto_enumTypes[0]
}

func (x NetworkConnection_NetworkConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkConnection_NetworkConnectionStatus.Descriptor instead.
func (NetworkConnection_NetworkConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_proto_rawDescGZIP(), []int{0, 0}
}

// A network connection resource.
type NetworkConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the network connection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the network.
	NetworkId string `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Provider specific info.
	//
	// Types that are assignable to ConnectionInfo:
	//
	//	*NetworkConnection_Aws
	//	*NetworkConnection_Google
	ConnectionInfo isNetworkConnection_ConnectionInfo `protobuf_oneof:"connection_info"`
	// Creation timestamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Description of the network connection.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// Current state of the network connection.
	Status NetworkConnection_NetworkConnectionStatus `protobuf:"varint,9,opt,name=status,proto3,enum=doublecloud.network.v1.NetworkConnection_NetworkConnectionStatus" json:"status,omitempty"`
	// The reason for the status change.
	StatusReason string `protobuf:"bytes,10,opt,name=status_reason,json=statusReason,proto3" json:"status_reason,omitempty"`
}

func (x *NetworkConnection) Reset() {
	*x = NetworkConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkConnection) ProtoMessage() {}

func (x *NetworkConnection) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkConnection.ProtoReflect.Descriptor instead.
func (*NetworkConnection) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkConnection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NetworkConnection) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (m *NetworkConnection) GetConnectionInfo() isNetworkConnection_ConnectionInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

func (x *NetworkConnection) GetAws() *AWSNetworkConnectionInfo {
	if x, ok := x.GetConnectionInfo().(*NetworkConnection_Aws); ok {
		return x.Aws
	}
	return nil
}

func (x *NetworkConnection) GetGoogle() *GoogleNetworkConnectionInfo {
	if x, ok := x.GetConnectionInfo().(*NetworkConnection_Google); ok {
		return x.Google
	}
	return nil
}

func (x *NetworkConnection) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *NetworkConnection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NetworkConnection) GetStatus() NetworkConnection_NetworkConnectionStatus {
	if x != nil {
		return x.Status
	}
	return NetworkConnection_NETWORK_CONNECTION_STATUS_INVALID
}

func (x *NetworkConnection) GetStatusReason() string {
	if x != nil {
		return x.StatusReason
	}
	return ""
}

type isNetworkConnection_ConnectionInfo interface {
	isNetworkConnection_ConnectionInfo()
}

type NetworkConnection_Aws struct {
	// AWS connection info.
	Aws *AWSNetworkConnectionInfo `protobuf:"bytes,4,opt,name=aws,proto3,oneof"`
}

type NetworkConnection_Google struct {
	// Google Cloud connection info.
	Google *GoogleNetworkConnectionInfo `protobuf:"bytes,5,opt,name=google,proto3,oneof"`
}

func (*NetworkConnection_Aws) isNetworkConnection_ConnectionInfo() {}

func (*NetworkConnection_Google) isNetworkConnection_ConnectionInfo() {}

type AWSNetworkConnectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*AWSNetworkConnectionInfo_Peering
	Type isAWSNetworkConnectionInfo_Type `protobuf_oneof:"type"`
}

func (x *AWSNetworkConnectionInfo) Reset() {
	*x = AWSNetworkConnectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSNetworkConnectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSNetworkConnectionInfo) ProtoMessage() {}

func (x *AWSNetworkConnectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSNetworkConnectionInfo.ProtoReflect.Descriptor instead.
func (*AWSNetworkConnectionInfo) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_proto_rawDescGZIP(), []int{1}
}

func (m *AWSNetworkConnectionInfo) GetType() isAWSNetworkConnectionInfo_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *AWSNetworkConnectionInfo) GetPeering() *AWSNetworkConnectionPeeringInfo {
	if x, ok := x.GetType().(*AWSNetworkConnectionInfo_Peering); ok {
		return x.Peering
	}
	return nil
}

type isAWSNetworkConnectionInfo_Type interface {
	isAWSNetworkConnectionInfo_Type()
}

type AWSNetworkConnectionInfo_Peering struct {
	// VPC Peering connection info.
	Peering *AWSNetworkConnectionPeeringInfo `protobuf:"bytes,1,opt,name=peering,proto3,oneof"`
}

func (*AWSNetworkConnectionInfo_Peering) isAWSNetworkConnectionInfo_Type() {}

type AWSNetworkConnectionPeeringInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// VPC ID.
	VpcId string `protobuf:"bytes,1,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	// ID of the VPC owner account.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// ID of the AWS region.
	RegionId string `protobuf:"bytes,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// Customer IPv4 CIDR block.
	Ipv4CidrBlock string `protobuf:"bytes,4,opt,name=ipv4_cidr_block,json=ipv4CidrBlock,proto3" json:"ipv4_cidr_block,omitempty"`
	// Customer IPv6 CIDR block.
	Ipv6CidrBlock string `protobuf:"bytes,5,opt,name=ipv6_cidr_block,json=ipv6CidrBlock,proto3" json:"ipv6_cidr_block,omitempty"`
	// Peering Connection ID.
	PeeringConnectionId string `protobuf:"bytes,6,opt,name=peering_connection_id,json=peeringConnectionId,proto3" json:"peering_connection_id,omitempty"`
	// Managed AWS IPv4 CIDR block. Customer will create route to this CIDR using
	// Peering Connection.
	ManagedIpv4CidrBlock string `protobuf:"bytes,7,opt,name=managed_ipv4_cidr_block,json=managedIpv4CidrBlock,proto3" json:"managed_ipv4_cidr_block,omitempty"`
	// managed AWS IPv6 CIDR block. Customer will create route to this CIDR using
	// Peering Connection.
	ManagedIpv6CidrBlock string `protobuf:"bytes,8,opt,name=managed_ipv6_cidr_block,json=managedIpv6CidrBlock,proto3" json:"managed_ipv6_cidr_block,omitempty"`
}

func (x *AWSNetworkConnectionPeeringInfo) Reset() {
	*x = AWSNetworkConnectionPeeringInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSNetworkConnectionPeeringInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSNetworkConnectionPeeringInfo) ProtoMessage() {}

func (x *AWSNetworkConnectionPeeringInfo) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSNetworkConnectionPeeringInfo.ProtoReflect.Descriptor instead.
func (*AWSNetworkConnectionPeeringInfo) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_proto_rawDescGZIP(), []int{2}
}

func (x *AWSNetworkConnectionPeeringInfo) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *AWSNetworkConnectionPeeringInfo) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *AWSNetworkConnectionPeeringInfo) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *AWSNetworkConnectionPeeringInfo) GetIpv4CidrBlock() string {
	if x != nil {
		return x.Ipv4CidrBlock
	}
	return ""
}

func (x *AWSNetworkConnectionPeeringInfo) GetIpv6CidrBlock() string {
	if x != nil {
		return x.Ipv6CidrBlock
	}
	return ""
}

func (x *AWSNetworkConnectionPeeringInfo) GetPeeringConnectionId() string {
	if x != nil {
		return x.PeeringConnectionId
	}
	return ""
}

func (x *AWSNetworkConnectionPeeringInfo) GetManagedIpv4CidrBlock() string {
	if x != nil {
		return x.ManagedIpv4CidrBlock
	}
	return ""
}

func (x *AWSNetworkConnectionPeeringInfo) GetManagedIpv6CidrBlock() string {
	if x != nil {
		return x.ManagedIpv6CidrBlock
	}
	return ""
}

type GoogleNetworkConnectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of this peering.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The URL of the peer network.
	PeerNetworkUrl string `protobuf:"bytes,2,opt,name=peer_network_url,json=peerNetworkUrl,proto3" json:"peer_network_url,omitempty"`
	// The URL of the managed GCP network.
	ManagedNetworkUrl string `protobuf:"bytes,3,opt,name=managed_network_url,json=managedNetworkUrl,proto3" json:"managed_network_url,omitempty"`
}

func (x *GoogleNetworkConnectionInfo) Reset() {
	*x = GoogleNetworkConnectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleNetworkConnectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleNetworkConnectionInfo) ProtoMessage() {}

func (x *GoogleNetworkConnectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_connection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleNetworkConnectionInfo.ProtoReflect.Descriptor instead.
func (*GoogleNetworkConnectionInfo) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_connection_proto_rawDescGZIP(), []int{3}
}

func (x *GoogleNetworkConnectionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoogleNetworkConnectionInfo) GetPeerNetworkUrl() string {
	if x != nil {
		return x.PeerNetworkUrl
	}
	return ""
}

func (x *GoogleNetworkConnectionInfo) GetManagedNetworkUrl() string {
	if x != nil {
		return x.ManagedNetworkUrl
	}
	return ""
}

var File_doublecloud_network_v1_network_connection_proto protoreflect.FileDescriptor

var file_doublecloud_network_v1_network_connection_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x05, 0x0a, 0x11, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12,
	0x44, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x4d, 0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x06, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x82, 0x02, 0x0a, 0x17, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x25, 0x0a, 0x21, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x25,
	0x0a, 0x21, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x42, 0x11, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x77, 0x0a, 0x18, 0x41, 0x57, 0x53, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x53, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57,
	0x53, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xe6, 0x02, 0x0a, 0x1f, 0x41, 0x57, 0x53, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x70, 0x76, 0x34, 0x5f,
	0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x69, 0x70, 0x76, 0x34, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x26, 0x0a, 0x0f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x43, 0x69,
	0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x65, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x63, 0x69, 0x64, 0x72,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x70, 0x76, 0x34, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x69, 0x70,
	0x76, 0x36, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x70, 0x76, 0x36,
	0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x8b, 0x01, 0x0a, 0x1b, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x65, 0x72, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x55, 0x72, 0x6c, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_network_v1_network_connection_proto_rawDescOnce sync.Once
	file_doublecloud_network_v1_network_connection_proto_rawDescData = file_doublecloud_network_v1_network_connection_proto_rawDesc
)

func file_doublecloud_network_v1_network_connection_proto_rawDescGZIP() []byte {
	file_doublecloud_network_v1_network_connection_proto_rawDescOnce.Do(func() {
		file_doublecloud_network_v1_network_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_network_v1_network_connection_proto_rawDescData)
	})
	return file_doublecloud_network_v1_network_connection_proto_rawDescData
}

var file_doublecloud_network_v1_network_connection_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_network_v1_network_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_doublecloud_network_v1_network_connection_proto_goTypes = []interface{}{
	(NetworkConnection_NetworkConnectionStatus)(0), // 0: doublecloud.network.v1.NetworkConnection.NetworkConnectionStatus
	(*NetworkConnection)(nil),                      // 1: doublecloud.network.v1.NetworkConnection
	(*AWSNetworkConnectionInfo)(nil),               // 2: doublecloud.network.v1.AWSNetworkConnectionInfo
	(*AWSNetworkConnectionPeeringInfo)(nil),        // 3: doublecloud.network.v1.AWSNetworkConnectionPeeringInfo
	(*GoogleNetworkConnectionInfo)(nil),            // 4: doublecloud.network.v1.GoogleNetworkConnectionInfo
	(*timestamppb.Timestamp)(nil),                  // 5: google.protobuf.Timestamp
}
var file_doublecloud_network_v1_network_connection_proto_depIdxs = []int32{
	2, // 0: doublecloud.network.v1.NetworkConnection.aws:type_name -> doublecloud.network.v1.AWSNetworkConnectionInfo
	4, // 1: doublecloud.network.v1.NetworkConnection.google:type_name -> doublecloud.network.v1.GoogleNetworkConnectionInfo
	5, // 2: doublecloud.network.v1.NetworkConnection.create_time:type_name -> google.protobuf.Timestamp
	0, // 3: doublecloud.network.v1.NetworkConnection.status:type_name -> doublecloud.network.v1.NetworkConnection.NetworkConnectionStatus
	3, // 4: doublecloud.network.v1.AWSNetworkConnectionInfo.peering:type_name -> doublecloud.network.v1.AWSNetworkConnectionPeeringInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_doublecloud_network_v1_network_connection_proto_init() }
func file_doublecloud_network_v1_network_connection_proto_init() {
	if File_doublecloud_network_v1_network_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_network_v1_network_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkConnection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_doublecloud_network_v1_network_connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSNetworkConnectionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_doublecloud_network_v1_network_connection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSNetworkConnectionPeeringInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_doublecloud_network_v1_network_connection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleNetworkConnectionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_doublecloud_network_v1_network_connection_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NetworkConnection_Aws)(nil),
		(*NetworkConnection_Google)(nil),
	}
	file_doublecloud_network_v1_network_connection_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AWSNetworkConnectionInfo_Peering)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_network_v1_network_connection_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_network_v1_network_connection_proto_goTypes,
		DependencyIndexes: file_doublecloud_network_v1_network_connection_proto_depIdxs,
		EnumInfos:         file_doublecloud_network_v1_network_connection_proto_enumTypes,
		MessageInfos:      file_doublecloud_network_v1_network_connection_proto_msgTypes,
	}.Build()
	File_doublecloud_network_v1_network_connection_proto = out.File
	file_doublecloud_network_v1_network_connection_proto_rawDesc = nil
	file_doublecloud_network_v1_network_connection_proto_goTypes = nil
	file_doublecloud_network_v1_network_connection_proto_depIdxs = nil
}
