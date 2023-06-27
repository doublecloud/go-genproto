// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: doublecloud/network/v1/network_service.proto

package network

import (
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
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

type GetNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Network resource to return.
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (x *GetNetworkRequest) Reset() {
	*x = GetNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkRequest) ProtoMessage() {}

func (x *GetNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetNetworkRequest) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

type ListNetworksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the project to list Networks in.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Paging information of the request
	Paging *v1.Paging `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	// Fields to filter a network list
	Filter *ListNetworksRequest_Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListNetworksRequest) Reset() {
	*x = ListNetworksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworksRequest) ProtoMessage() {}

func (x *ListNetworksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworksRequest.ProtoReflect.Descriptor instead.
func (*ListNetworksRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListNetworksRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ListNetworksRequest) GetPaging() *v1.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

func (x *ListNetworksRequest) GetFilter() *ListNetworksRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListNetworksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Network resources.
	Networks []*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	// Pagination information of the response
	NextPage *v1.NextPage `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *ListNetworksResponse) Reset() {
	*x = ListNetworksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworksResponse) ProtoMessage() {}

func (x *ListNetworksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworksResponse.ProtoReflect.Descriptor instead.
func (*ListNetworksResponse) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListNetworksResponse) GetNetworks() []*Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *ListNetworksResponse) GetNextPage() *v1.NextPage {
	if x != nil {
		return x.NextPage
	}
	return nil
}

type CreateNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the project to create the Network in.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Type of the cloud where instances should be hosted.
	CloudType string `protobuf:"bytes,2,opt,name=cloud_type,json=cloudType,proto3" json:"cloud_type,omitempty"`
	// ID of the region to place instances.
	RegionId string `protobuf:"bytes,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// Name of the Network. The name must be unique within the project.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Network.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The IPv4 network range for the subnet, in CIDR notation.
	// For example, 10.0.0.0/16.
	Ipv4CidrBlock string `protobuf:"bytes,6,opt,name=ipv4_cidr_block,json=ipv4CidrBlock,proto3" json:"ipv4_cidr_block,omitempty"`
}

func (x *CreateNetworkRequest) Reset() {
	*x = CreateNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkRequest) ProtoMessage() {}

func (x *CreateNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkRequest.ProtoReflect.Descriptor instead.
func (*CreateNetworkRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNetworkRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateNetworkRequest) GetCloudType() string {
	if x != nil {
		return x.CloudType
	}
	return ""
}

func (x *CreateNetworkRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *CreateNetworkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNetworkRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateNetworkRequest) GetIpv4CidrBlock() string {
	if x != nil {
		return x.Ipv4CidrBlock
	}
	return ""
}

type DeleteNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Network to delete.
	// To get the Network ID, use a [NetworkService.List] request.
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (x *DeleteNetworkRequest) Reset() {
	*x = DeleteNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkRequest) ProtoMessage() {}

func (x *DeleteNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkRequest.ProtoReflect.Descriptor instead.
func (*DeleteNetworkRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNetworkRequest) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

type ImportNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the project to create the Network in.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name of the Network. The name must be unique within the project.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Network.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Provider specific parameters.
	//
	// Types that are assignable to Params:
	//
	//	*ImportNetworkRequest_Aws
	Params isImportNetworkRequest_Params `protobuf_oneof:"params"`
}

func (x *ImportNetworkRequest) Reset() {
	*x = ImportNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportNetworkRequest) ProtoMessage() {}

func (x *ImportNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportNetworkRequest.ProtoReflect.Descriptor instead.
func (*ImportNetworkRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{5}
}

func (x *ImportNetworkRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ImportNetworkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportNetworkRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *ImportNetworkRequest) GetParams() isImportNetworkRequest_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *ImportNetworkRequest) GetAws() *ImportAWSVPCRequest {
	if x, ok := x.GetParams().(*ImportNetworkRequest_Aws); ok {
		return x.Aws
	}
	return nil
}

type isImportNetworkRequest_Params interface {
	isImportNetworkRequest_Params()
}

type ImportNetworkRequest_Aws struct {
	// Import AWS VPC.
	Aws *ImportAWSVPCRequest `protobuf:"bytes,4,opt,name=aws,proto3,oneof"`
}

func (*ImportNetworkRequest_Aws) isImportNetworkRequest_Params() {}

type ImportAWSVPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the VPC.
	VpcId string `protobuf:"bytes,1,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	// ID of the region to place instances.
	RegionId string `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// ID of the VPC owner account.
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// IAM role ARN to use for resource creations.
	IamRoleArn string `protobuf:"bytes,4,opt,name=iam_role_arn,json=iamRoleArn,proto3" json:"iam_role_arn,omitempty"`
	// ID of the CloudFormation stack.
	StackId string `protobuf:"bytes,5,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	// CloudFormation template version.
	Version string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ImportAWSVPCRequest) Reset() {
	*x = ImportAWSVPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportAWSVPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportAWSVPCRequest) ProtoMessage() {}

func (x *ImportAWSVPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportAWSVPCRequest.ProtoReflect.Descriptor instead.
func (*ImportAWSVPCRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{6}
}

func (x *ImportAWSVPCRequest) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *ImportAWSVPCRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *ImportAWSVPCRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ImportAWSVPCRequest) GetIamRoleArn() string {
	if x != nil {
		return x.IamRoleArn
	}
	return ""
}

func (x *ImportAWSVPCRequest) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *ImportAWSVPCRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ListNetworksRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the cloud where instances should be hosted.
	CloudType *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=cloud_type,json=cloudType,proto3" json:"cloud_type,omitempty"`
	// ID of the region to place instances.
	RegionId *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// Status of the network.
	Status Network_NetworkStatus `protobuf:"varint,3,opt,name=status,proto3,enum=doublecloud.network.v1.Network_NetworkStatus" json:"status,omitempty"`
	// Flag to show if a network is external (BYOA feature)
	IsExternal *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=is_external,json=isExternal,proto3" json:"is_external,omitempty"`
}

func (x *ListNetworksRequest_Filter) Reset() {
	*x = ListNetworksRequest_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworksRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworksRequest_Filter) ProtoMessage() {}

func (x *ListNetworksRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_network_v1_network_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworksRequest_Filter.ProtoReflect.Descriptor instead.
func (*ListNetworksRequest_Filter) Descriptor() ([]byte, []int) {
	return file_doublecloud_network_v1_network_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListNetworksRequest_Filter) GetCloudType() *wrapperspb.StringValue {
	if x != nil {
		return x.CloudType
	}
	return nil
}

func (x *ListNetworksRequest_Filter) GetRegionId() *wrapperspb.StringValue {
	if x != nil {
		return x.RegionId
	}
	return nil
}

func (x *ListNetworksRequest_Filter) GetStatus() Network_NetworkStatus {
	if x != nil {
		return x.Status
	}
	return Network_NETWORK_STATUS_INVALID
}

func (x *ListNetworksRequest_Filter) GetIsExternal() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsExternal
	}
	return nil
}

var File_doublecloud_network_v1_network_service_proto protoreflect.FileDescriptor

var file_doublecloud_network_v1_network_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x22, 0xb7, 0x03,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x12, 0x4a, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x1a, 0x84, 0x02, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0a, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x73, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x35, 0x0a,
	0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x0f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x34, 0x43, 0x69, 0x64,
	0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x22, 0xb6, 0x01,
	0x0a, 0x14, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x03, 0x61,
	0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x57, 0x53, 0x56, 0x50, 0x43, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77, 0x73, 0x42, 0x08, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x13, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x41, 0x57, 0x53, 0x56, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x70, 0x63, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x61, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65,
	0x41, 0x72, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xbf, 0x03, 0x0a, 0x0e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x29, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x61,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x51, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2c,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x06, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x2c, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_network_v1_network_service_proto_rawDescOnce sync.Once
	file_doublecloud_network_v1_network_service_proto_rawDescData = file_doublecloud_network_v1_network_service_proto_rawDesc
)

func file_doublecloud_network_v1_network_service_proto_rawDescGZIP() []byte {
	file_doublecloud_network_v1_network_service_proto_rawDescOnce.Do(func() {
		file_doublecloud_network_v1_network_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_network_v1_network_service_proto_rawDescData)
	})
	return file_doublecloud_network_v1_network_service_proto_rawDescData
}

var file_doublecloud_network_v1_network_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_doublecloud_network_v1_network_service_proto_goTypes = []interface{}{
	(*GetNetworkRequest)(nil),          // 0: doublecloud.network.v1.GetNetworkRequest
	(*ListNetworksRequest)(nil),        // 1: doublecloud.network.v1.ListNetworksRequest
	(*ListNetworksResponse)(nil),       // 2: doublecloud.network.v1.ListNetworksResponse
	(*CreateNetworkRequest)(nil),       // 3: doublecloud.network.v1.CreateNetworkRequest
	(*DeleteNetworkRequest)(nil),       // 4: doublecloud.network.v1.DeleteNetworkRequest
	(*ImportNetworkRequest)(nil),       // 5: doublecloud.network.v1.ImportNetworkRequest
	(*ImportAWSVPCRequest)(nil),        // 6: doublecloud.network.v1.ImportAWSVPCRequest
	(*ListNetworksRequest_Filter)(nil), // 7: doublecloud.network.v1.ListNetworksRequest.Filter
	(*v1.Paging)(nil),                  // 8: doublecloud.v1.Paging
	(*Network)(nil),                    // 9: doublecloud.network.v1.Network
	(*v1.NextPage)(nil),                // 10: doublecloud.v1.NextPage
	(*wrapperspb.StringValue)(nil),     // 11: google.protobuf.StringValue
	(Network_NetworkStatus)(0),         // 12: doublecloud.network.v1.Network.NetworkStatus
	(*wrapperspb.BoolValue)(nil),       // 13: google.protobuf.BoolValue
	(*v1.Operation)(nil),               // 14: doublecloud.v1.Operation
}
var file_doublecloud_network_v1_network_service_proto_depIdxs = []int32{
	8,  // 0: doublecloud.network.v1.ListNetworksRequest.paging:type_name -> doublecloud.v1.Paging
	7,  // 1: doublecloud.network.v1.ListNetworksRequest.filter:type_name -> doublecloud.network.v1.ListNetworksRequest.Filter
	9,  // 2: doublecloud.network.v1.ListNetworksResponse.networks:type_name -> doublecloud.network.v1.Network
	10, // 3: doublecloud.network.v1.ListNetworksResponse.next_page:type_name -> doublecloud.v1.NextPage
	6,  // 4: doublecloud.network.v1.ImportNetworkRequest.aws:type_name -> doublecloud.network.v1.ImportAWSVPCRequest
	11, // 5: doublecloud.network.v1.ListNetworksRequest.Filter.cloud_type:type_name -> google.protobuf.StringValue
	11, // 6: doublecloud.network.v1.ListNetworksRequest.Filter.region_id:type_name -> google.protobuf.StringValue
	12, // 7: doublecloud.network.v1.ListNetworksRequest.Filter.status:type_name -> doublecloud.network.v1.Network.NetworkStatus
	13, // 8: doublecloud.network.v1.ListNetworksRequest.Filter.is_external:type_name -> google.protobuf.BoolValue
	0,  // 9: doublecloud.network.v1.NetworkService.Get:input_type -> doublecloud.network.v1.GetNetworkRequest
	1,  // 10: doublecloud.network.v1.NetworkService.List:input_type -> doublecloud.network.v1.ListNetworksRequest
	3,  // 11: doublecloud.network.v1.NetworkService.Create:input_type -> doublecloud.network.v1.CreateNetworkRequest
	4,  // 12: doublecloud.network.v1.NetworkService.Delete:input_type -> doublecloud.network.v1.DeleteNetworkRequest
	5,  // 13: doublecloud.network.v1.NetworkService.Import:input_type -> doublecloud.network.v1.ImportNetworkRequest
	9,  // 14: doublecloud.network.v1.NetworkService.Get:output_type -> doublecloud.network.v1.Network
	2,  // 15: doublecloud.network.v1.NetworkService.List:output_type -> doublecloud.network.v1.ListNetworksResponse
	14, // 16: doublecloud.network.v1.NetworkService.Create:output_type -> doublecloud.v1.Operation
	14, // 17: doublecloud.network.v1.NetworkService.Delete:output_type -> doublecloud.v1.Operation
	14, // 18: doublecloud.network.v1.NetworkService.Import:output_type -> doublecloud.v1.Operation
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_doublecloud_network_v1_network_service_proto_init() }
func file_doublecloud_network_v1_network_service_proto_init() {
	if File_doublecloud_network_v1_network_service_proto != nil {
		return
	}
	file_doublecloud_network_v1_network_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_network_v1_network_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkRequest); i {
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
		file_doublecloud_network_v1_network_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworksRequest); i {
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
		file_doublecloud_network_v1_network_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworksResponse); i {
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
		file_doublecloud_network_v1_network_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkRequest); i {
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
		file_doublecloud_network_v1_network_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkRequest); i {
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
		file_doublecloud_network_v1_network_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportNetworkRequest); i {
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
		file_doublecloud_network_v1_network_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportAWSVPCRequest); i {
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
		file_doublecloud_network_v1_network_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworksRequest_Filter); i {
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
	file_doublecloud_network_v1_network_service_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ImportNetworkRequest_Aws)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_network_v1_network_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doublecloud_network_v1_network_service_proto_goTypes,
		DependencyIndexes: file_doublecloud_network_v1_network_service_proto_depIdxs,
		MessageInfos:      file_doublecloud_network_v1_network_service_proto_msgTypes,
	}.Build()
	File_doublecloud_network_v1_network_service_proto = out.File
	file_doublecloud_network_v1_network_service_proto_rawDesc = nil
	file_doublecloud_network_v1_network_service_proto_goTypes = nil
	file_doublecloud_network_v1_network_service_proto_depIdxs = nil
}
