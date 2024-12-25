// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/hubspot_source.proto

package endpoint_airbyte

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HubspotSource struct {
	state                     protoimpl.MessageState     `protogen:"open.v1"`
	StartDate                 string                     `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	Credentials               *HubspotSource_Credentials `protobuf:"bytes,2,opt,name=credentials,proto3" json:"credentials,omitempty"`
	EnableExperimentalStreams bool                       `protobuf:"varint,3,opt,name=enable_experimental_streams,json=enableExperimentalStreams,proto3" json:"enable_experimental_streams,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *HubspotSource) Reset() {
	*x = HubspotSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HubspotSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HubspotSource) ProtoMessage() {}

func (x *HubspotSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HubspotSource.ProtoReflect.Descriptor instead.
func (*HubspotSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescGZIP(), []int{0}
}

func (x *HubspotSource) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *HubspotSource) GetCredentials() *HubspotSource_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *HubspotSource) GetEnableExperimentalStreams() bool {
	if x != nil {
		return x.EnableExperimentalStreams
	}
	return false
}

type HubspotSource_Credentials struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to AuthMethod:
	//
	//	*HubspotSource_Credentials_PrivateApp_
	AuthMethod    isHubspotSource_Credentials_AuthMethod `protobuf_oneof:"auth_method"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HubspotSource_Credentials) Reset() {
	*x = HubspotSource_Credentials{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HubspotSource_Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HubspotSource_Credentials) ProtoMessage() {}

func (x *HubspotSource_Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HubspotSource_Credentials.ProtoReflect.Descriptor instead.
func (*HubspotSource_Credentials) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HubspotSource_Credentials) GetAuthMethod() isHubspotSource_Credentials_AuthMethod {
	if x != nil {
		return x.AuthMethod
	}
	return nil
}

func (x *HubspotSource_Credentials) GetPrivateApp() *HubspotSource_Credentials_PrivateApp {
	if x != nil {
		if x, ok := x.AuthMethod.(*HubspotSource_Credentials_PrivateApp_); ok {
			return x.PrivateApp
		}
	}
	return nil
}

type isHubspotSource_Credentials_AuthMethod interface {
	isHubspotSource_Credentials_AuthMethod()
}

type HubspotSource_Credentials_PrivateApp_ struct {
	PrivateApp *HubspotSource_Credentials_PrivateApp `protobuf:"bytes,1,opt,name=private_app,json=privateApp,proto3,oneof"`
}

func (*HubspotSource_Credentials_PrivateApp_) isHubspotSource_Credentials_AuthMethod() {}

type HubspotSource_Credentials_PrivateApp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HubspotSource_Credentials_PrivateApp) Reset() {
	*x = HubspotSource_Credentials_PrivateApp{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HubspotSource_Credentials_PrivateApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HubspotSource_Credentials_PrivateApp) ProtoMessage() {}

func (x *HubspotSource_Credentials_PrivateApp) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HubspotSource_Credentials_PrivateApp.ProtoReflect.Descriptor instead.
func (*HubspotSource_Credentials_PrivateApp) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *HubspotSource_Credentials_PrivateApp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x68, 0x75, 0x62, 0x73, 0x70,
	0x6f, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0x98, 0x03, 0x0a, 0x0d, 0x48, 0x75,
	0x62, 0x73, 0x70, 0x6f, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x65, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x43, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x48, 0x75, 0x62, 0x73, 0x70,
	0x6f, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x1a, 0xc0, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x71, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74,
	0x65, 0x2e, 0x48, 0x75, 0x62, 0x73, 0x70, 0x6f, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x1a, 0x2f, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62,
	0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_goTypes = []any{
	(*HubspotSource)(nil),                        // 0: doublecloud.transfer.v1.endpoint.airbyte.HubspotSource
	(*HubspotSource_Credentials)(nil),            // 1: doublecloud.transfer.v1.endpoint.airbyte.HubspotSource.Credentials
	(*HubspotSource_Credentials_PrivateApp)(nil), // 2: doublecloud.transfer.v1.endpoint.airbyte.HubspotSource.Credentials.PrivateApp
}
var file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_depIdxs = []int32{
	1, // 0: doublecloud.transfer.v1.endpoint.airbyte.HubspotSource.credentials:type_name -> doublecloud.transfer.v1.endpoint.airbyte.HubspotSource.Credentials
	2, // 1: doublecloud.transfer.v1.endpoint.airbyte.HubspotSource.Credentials.private_app:type_name -> doublecloud.transfer.v1.endpoint.airbyte.HubspotSource.Credentials.PrivateApp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes[1].OneofWrappers = []any{
		(*HubspotSource_Credentials_PrivateApp_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_hubspot_source_proto_depIdxs = nil
}
