// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: doublecloud/v1/cluster.proto

package doublecloud

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterStatus int32

const (
	// Invalid cluster status value.
	ClusterStatus_CLUSTER_STATUS_INVALID ClusterStatus = 0
	// Cluster is running and all cluster hosts are healthy.
	ClusterStatus_CLUSTER_STATUS_ALIVE ClusterStatus = 1
	// Cluster is working below capacity (at least one host in the cluster is not
	// healthy).
	ClusterStatus_CLUSTER_STATUS_DEGRADED ClusterStatus = 2
	// Cluster is inoperable (all hosts in the cluster are inoperable).
	ClusterStatus_CLUSTER_STATUS_DEAD ClusterStatus = 3
	// State of the cluster is unknown (all hosts in the cluster are in unknown state).
	ClusterStatus_CLUSTER_STATUS_UNKNOWN ClusterStatus = 4
	// Cluster is being created.
	ClusterStatus_CLUSTER_STATUS_CREATING ClusterStatus = 5
	// Cluster is being updated.
	ClusterStatus_CLUSTER_STATUS_UPDATING ClusterStatus = 6
	// Cluster is stopping.
	ClusterStatus_CLUSTER_STATUS_STOPPING ClusterStatus = 7
	// Cluster stopped.
	ClusterStatus_CLUSTER_STATUS_STOPPED ClusterStatus = 8
	// Cluster is starting.
	ClusterStatus_CLUSTER_STATUS_STARTING ClusterStatus = 9
	// Cluster encountered a problem and cannot operate.
	ClusterStatus_CLUSTER_STATUS_ERROR ClusterStatus = 10
)

// Enum value maps for ClusterStatus.
var (
	ClusterStatus_name = map[int32]string{
		0:  "CLUSTER_STATUS_INVALID",
		1:  "CLUSTER_STATUS_ALIVE",
		2:  "CLUSTER_STATUS_DEGRADED",
		3:  "CLUSTER_STATUS_DEAD",
		4:  "CLUSTER_STATUS_UNKNOWN",
		5:  "CLUSTER_STATUS_CREATING",
		6:  "CLUSTER_STATUS_UPDATING",
		7:  "CLUSTER_STATUS_STOPPING",
		8:  "CLUSTER_STATUS_STOPPED",
		9:  "CLUSTER_STATUS_STARTING",
		10: "CLUSTER_STATUS_ERROR",
	}
	ClusterStatus_value = map[string]int32{
		"CLUSTER_STATUS_INVALID":  0,
		"CLUSTER_STATUS_ALIVE":    1,
		"CLUSTER_STATUS_DEGRADED": 2,
		"CLUSTER_STATUS_DEAD":     3,
		"CLUSTER_STATUS_UNKNOWN":  4,
		"CLUSTER_STATUS_CREATING": 5,
		"CLUSTER_STATUS_UPDATING": 6,
		"CLUSTER_STATUS_STOPPING": 7,
		"CLUSTER_STATUS_STOPPED":  8,
		"CLUSTER_STATUS_STARTING": 9,
		"CLUSTER_STATUS_ERROR":    10,
	}
)

func (x ClusterStatus) Enum() *ClusterStatus {
	p := new(ClusterStatus)
	*p = x
	return p
}

func (x ClusterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_v1_cluster_proto_enumTypes[0].Descriptor()
}

func (ClusterStatus) Type() protoreflect.EnumType {
	return &file_doublecloud_v1_cluster_proto_enumTypes[0]
}

func (x ClusterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterStatus.Descriptor instead.
func (ClusterStatus) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{0}
}

type HostStatus int32

const (
	// Host is in unknown state (we have no data)
	HostStatus_HOST_STATUS_INVALID HostStatus = 0
	// Host is alive and well (all services are alive)
	HostStatus_HOST_STATUS_ALIVE HostStatus = 1
	// Host is inoperable (it cannot perform any of its essential functions)
	HostStatus_HOST_STATUS_DEAD HostStatus = 2
	// Host is partially alive (it can perform some of its essential functions)
	HostStatus_HOST_STATUS_DEGRADED HostStatus = 3
)

// Enum value maps for HostStatus.
var (
	HostStatus_name = map[int32]string{
		0: "HOST_STATUS_INVALID",
		1: "HOST_STATUS_ALIVE",
		2: "HOST_STATUS_DEAD",
		3: "HOST_STATUS_DEGRADED",
	}
	HostStatus_value = map[string]int32{
		"HOST_STATUS_INVALID":  0,
		"HOST_STATUS_ALIVE":    1,
		"HOST_STATUS_DEAD":     2,
		"HOST_STATUS_DEGRADED": 3,
	}
)

func (x HostStatus) Enum() *HostStatus {
	p := new(HostStatus)
	*p = x
	return p
}

func (x HostStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HostStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_v1_cluster_proto_enumTypes[1].Descriptor()
}

func (HostStatus) Type() protoreflect.EnumType {
	return &file_doublecloud_v1_cluster_proto_enumTypes[1]
}

func (x HostStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HostStatus.Descriptor instead.
func (HostStatus) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{1}
}

type Access_DataService int32

const (
	Access_DATA_SERVICE_INVALID Access_DataService = 0
	// Allow access for DataLens
	Access_DATA_SERVICE_VISUALIZATION Access_DataService = 1
	// Allow access for DataTransfer
	Access_DATA_SERVICE_TRANSFER Access_DataService = 2
)

// Enum value maps for Access_DataService.
var (
	Access_DataService_name = map[int32]string{
		0: "DATA_SERVICE_INVALID",
		1: "DATA_SERVICE_VISUALIZATION",
		2: "DATA_SERVICE_TRANSFER",
	}
	Access_DataService_value = map[string]int32{
		"DATA_SERVICE_INVALID":       0,
		"DATA_SERVICE_VISUALIZATION": 1,
		"DATA_SERVICE_TRANSFER":      2,
	}
)

func (x Access_DataService) Enum() *Access_DataService {
	p := new(Access_DataService)
	*p = x
	return p
}

func (x Access_DataService) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Access_DataService) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_v1_cluster_proto_enumTypes[2].Descriptor()
}

func (Access_DataService) Type() protoreflect.EnumType {
	return &file_doublecloud_v1_cluster_proto_enumTypes[2]
}

func (x Access_DataService) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Access_DataService.Descriptor instead.
func (Access_DataService) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{0, 0}
}

// Access settings
type Access struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of ipv4 IPs (in 'CIDR' notation) of external networks that are allowed
	// to connect to the instance
	Ipv4CidrBlocks *Access_CidrBlockList `protobuf:"bytes,1,opt,name=ipv4_cidr_blocks,json=ipv4CidrBlocks,proto3" json:"ipv4_cidr_blocks,omitempty"`
	// The list of ipv6 IPs (in 'CIDR' notation) of external networks that are allowed
	// to connect to the instance
	Ipv6CidrBlocks *Access_CidrBlockList   `protobuf:"bytes,2,opt,name=ipv6_cidr_blocks,json=ipv6CidrBlocks,proto3" json:"ipv6_cidr_blocks,omitempty"`
	DataServices   *Access_DataServiceList `protobuf:"bytes,3,opt,name=data_services,json=dataServices,proto3" json:"data_services,omitempty"`
}

func (x *Access) Reset() {
	*x = Access{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access) ProtoMessage() {}

func (x *Access) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access.ProtoReflect.Descriptor instead.
func (*Access) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Access) GetIpv4CidrBlocks() *Access_CidrBlockList {
	if x != nil {
		return x.Ipv4CidrBlocks
	}
	return nil
}

func (x *Access) GetIpv6CidrBlocks() *Access_CidrBlockList {
	if x != nil {
		return x.Ipv6CidrBlocks
	}
	return nil
}

func (x *Access) GetDataServices() *Access_DataServiceList {
	if x != nil {
		return x.DataServices
	}
	return nil
}

type DataEncryption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// deprecated. Encryption is always enabled.
	//
	// Deprecated: Marked as deprecated in doublecloud/v1/cluster.proto.
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *DataEncryption) Reset() {
	*x = DataEncryption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_v1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataEncryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataEncryption) ProtoMessage() {}

func (x *DataEncryption) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataEncryption.ProtoReflect.Descriptor instead.
func (*DataEncryption) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in doublecloud/v1/cluster.proto.
func (x *DataEncryption) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

// CIDR subnet with description.
type Access_CidrBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Access_CidrBlock) Reset() {
	*x = Access_CidrBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_v1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access_CidrBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access_CidrBlock) ProtoMessage() {}

func (x *Access_CidrBlock) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access_CidrBlock.ProtoReflect.Descriptor instead.
func (*Access_CidrBlock) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Access_CidrBlock) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Access_CidrBlock) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Wrapper for list of CidrBlock
type Access_CidrBlockList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Access_CidrBlock `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Access_CidrBlockList) Reset() {
	*x = Access_CidrBlockList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_v1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access_CidrBlockList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access_CidrBlockList) ProtoMessage() {}

func (x *Access_CidrBlockList) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access_CidrBlockList.ProtoReflect.Descriptor instead.
func (*Access_CidrBlockList) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Access_CidrBlockList) GetValues() []*Access_CidrBlock {
	if x != nil {
		return x.Values
	}
	return nil
}

// Wrapper for list of DataServices
type Access_DataServiceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []Access_DataService `protobuf:"varint,1,rep,packed,name=values,proto3,enum=doublecloud.v1.Access_DataService" json:"values,omitempty"`
}

func (x *Access_DataServiceList) Reset() {
	*x = Access_DataServiceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_v1_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access_DataServiceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access_DataServiceList) ProtoMessage() {}

func (x *Access_DataServiceList) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access_DataServiceList.ProtoReflect.Descriptor instead.
func (*Access_DataServiceList) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_cluster_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Access_DataServiceList) GetValues() []Access_DataService {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_doublecloud_v1_cluster_proto protoreflect.FileDescriptor

var file_doublecloud_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8,
	0x04, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x69, 0x70, 0x76,
	0x34, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x69, 0x64, 0x72,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0e, 0x69, 0x70, 0x76, 0x34, 0x43,
	0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x69, 0x70, 0x76,
	0x36, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x69, 0x64, 0x72,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0e, 0x69, 0x70, 0x76, 0x36, 0x43,
	0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x43, 0x0a, 0x09, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x49, 0x0a, 0x0d, 0x43,
	0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x4d, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1e,
	0x0a, 0x1a, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x56,
	0x49, 0x53, 0x55, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x02, 0x22, 0x4a, 0x0a, 0x0e, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x2a, 0xc1, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4c, 0x55, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a,
	0x17, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x44, 0x45, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c,
	0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41,
	0x44, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x12,
	0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x55,
	0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x50,
	0x50, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12,
	0x18, 0x0a, 0x14, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0a, 0x2a, 0x6c, 0x0a, 0x0a, 0x48, 0x6f, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x47,
	0x52, 0x41, 0x44, 0x45, 0x44, 0x10, 0x03, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_v1_cluster_proto_rawDescOnce sync.Once
	file_doublecloud_v1_cluster_proto_rawDescData = file_doublecloud_v1_cluster_proto_rawDesc
)

func file_doublecloud_v1_cluster_proto_rawDescGZIP() []byte {
	file_doublecloud_v1_cluster_proto_rawDescOnce.Do(func() {
		file_doublecloud_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_v1_cluster_proto_rawDescData)
	})
	return file_doublecloud_v1_cluster_proto_rawDescData
}

var file_doublecloud_v1_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_doublecloud_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_doublecloud_v1_cluster_proto_goTypes = []interface{}{
	(ClusterStatus)(0),             // 0: doublecloud.v1.ClusterStatus
	(HostStatus)(0),                // 1: doublecloud.v1.HostStatus
	(Access_DataService)(0),        // 2: doublecloud.v1.Access.DataService
	(*Access)(nil),                 // 3: doublecloud.v1.Access
	(*DataEncryption)(nil),         // 4: doublecloud.v1.DataEncryption
	(*Access_CidrBlock)(nil),       // 5: doublecloud.v1.Access.CidrBlock
	(*Access_CidrBlockList)(nil),   // 6: doublecloud.v1.Access.CidrBlockList
	(*Access_DataServiceList)(nil), // 7: doublecloud.v1.Access.DataServiceList
	(*wrapperspb.BoolValue)(nil),   // 8: google.protobuf.BoolValue
}
var file_doublecloud_v1_cluster_proto_depIdxs = []int32{
	6, // 0: doublecloud.v1.Access.ipv4_cidr_blocks:type_name -> doublecloud.v1.Access.CidrBlockList
	6, // 1: doublecloud.v1.Access.ipv6_cidr_blocks:type_name -> doublecloud.v1.Access.CidrBlockList
	7, // 2: doublecloud.v1.Access.data_services:type_name -> doublecloud.v1.Access.DataServiceList
	8, // 3: doublecloud.v1.DataEncryption.enabled:type_name -> google.protobuf.BoolValue
	5, // 4: doublecloud.v1.Access.CidrBlockList.values:type_name -> doublecloud.v1.Access.CidrBlock
	2, // 5: doublecloud.v1.Access.DataServiceList.values:type_name -> doublecloud.v1.Access.DataService
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_doublecloud_v1_cluster_proto_init() }
func file_doublecloud_v1_cluster_proto_init() {
	if File_doublecloud_v1_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access); i {
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
		file_doublecloud_v1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataEncryption); i {
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
		file_doublecloud_v1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access_CidrBlock); i {
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
		file_doublecloud_v1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access_CidrBlockList); i {
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
		file_doublecloud_v1_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access_DataServiceList); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_v1_cluster_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_v1_cluster_proto_goTypes,
		DependencyIndexes: file_doublecloud_v1_cluster_proto_depIdxs,
		EnumInfos:         file_doublecloud_v1_cluster_proto_enumTypes,
		MessageInfos:      file_doublecloud_v1_cluster_proto_msgTypes,
	}.Build()
	File_doublecloud_v1_cluster_proto = out.File
	file_doublecloud_v1_cluster_proto_rawDesc = nil
	file_doublecloud_v1_cluster_proto_goTypes = nil
	file_doublecloud_v1_cluster_proto_depIdxs = nil
}
