// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/linkedin_ads_source.proto

package endpoint_airbyte

import (
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

type LinkedinAdsSource struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	StartDate     string                         `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	AccountIds    []int64                        `protobuf:"varint,2,rep,packed,name=account_ids,json=accountIds,proto3" json:"account_ids,omitempty"`
	Credentials   *LinkedinAdsSource_Credentials `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LinkedinAdsSource) Reset() {
	*x = LinkedinAdsSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LinkedinAdsSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedinAdsSource) ProtoMessage() {}

func (x *LinkedinAdsSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedinAdsSource.ProtoReflect.Descriptor instead.
func (*LinkedinAdsSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescGZIP(), []int{0}
}

func (x *LinkedinAdsSource) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *LinkedinAdsSource) GetAccountIds() []int64 {
	if x != nil {
		return x.AccountIds
	}
	return nil
}

func (x *LinkedinAdsSource) GetCredentials() *LinkedinAdsSource_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type LinkedinAdsSource_Credentials struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Credentials:
	//
	//	*LinkedinAdsSource_Credentials_Oauth
	//	*LinkedinAdsSource_Credentials_AccessToken
	Credentials   isLinkedinAdsSource_Credentials_Credentials `protobuf_oneof:"credentials"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LinkedinAdsSource_Credentials) Reset() {
	*x = LinkedinAdsSource_Credentials{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LinkedinAdsSource_Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedinAdsSource_Credentials) ProtoMessage() {}

func (x *LinkedinAdsSource_Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedinAdsSource_Credentials.ProtoReflect.Descriptor instead.
func (*LinkedinAdsSource_Credentials) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LinkedinAdsSource_Credentials) GetCredentials() isLinkedinAdsSource_Credentials_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *LinkedinAdsSource_Credentials) GetOauth() *LinkedinAdsSource_Credentials_OAuth {
	if x != nil {
		if x, ok := x.Credentials.(*LinkedinAdsSource_Credentials_Oauth); ok {
			return x.Oauth
		}
	}
	return nil
}

func (x *LinkedinAdsSource_Credentials) GetAccessToken() string {
	if x != nil {
		if x, ok := x.Credentials.(*LinkedinAdsSource_Credentials_AccessToken); ok {
			return x.AccessToken
		}
	}
	return ""
}

type isLinkedinAdsSource_Credentials_Credentials interface {
	isLinkedinAdsSource_Credentials_Credentials()
}

type LinkedinAdsSource_Credentials_Oauth struct {
	Oauth *LinkedinAdsSource_Credentials_OAuth `protobuf:"bytes,1,opt,name=oauth,proto3,oneof"`
}

type LinkedinAdsSource_Credentials_AccessToken struct {
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3,oneof"`
}

func (*LinkedinAdsSource_Credentials_Oauth) isLinkedinAdsSource_Credentials_Credentials() {}

func (*LinkedinAdsSource_Credentials_AccessToken) isLinkedinAdsSource_Credentials_Credentials() {}

type LinkedinAdsSource_Credentials_OAuth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret  string                 `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LinkedinAdsSource_Credentials_OAuth) Reset() {
	*x = LinkedinAdsSource_Credentials_OAuth{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LinkedinAdsSource_Credentials_OAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedinAdsSource_Credentials_OAuth) ProtoMessage() {}

func (x *LinkedinAdsSource_Credentials_OAuth) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedinAdsSource_Credentials_OAuth.ProtoReflect.Descriptor instead.
func (*LinkedinAdsSource_Credentials_OAuth) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *LinkedinAdsSource_Credentials_OAuth) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *LinkedinAdsSource_Credentials_OAuth) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *LinkedinAdsSource_Credentials_OAuth) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDesc = string([]byte{
	0x0a, 0x42, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x69, 0x6e, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0xd9,
	0x03, 0x0a, 0x11, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x41, 0x64, 0x73, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x73, 0x12, 0x69, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72,
	0x62, 0x79, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x41, 0x64, 0x73,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a,
	0x98, 0x02, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x65, 0x0a, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x69, 0x6e, 0x41, 0x64, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52,
	0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x6e, 0x0a, 0x05, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescData []byte
)

func file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDesc)))
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_goTypes = []any{
	(*LinkedinAdsSource)(nil),                   // 0: doublecloud.transfer.v1.endpoint.airbyte.LinkedinAdsSource
	(*LinkedinAdsSource_Credentials)(nil),       // 1: doublecloud.transfer.v1.endpoint.airbyte.LinkedinAdsSource.Credentials
	(*LinkedinAdsSource_Credentials_OAuth)(nil), // 2: doublecloud.transfer.v1.endpoint.airbyte.LinkedinAdsSource.Credentials.OAuth
}
var file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_depIdxs = []int32{
	1, // 0: doublecloud.transfer.v1.endpoint.airbyte.LinkedinAdsSource.credentials:type_name -> doublecloud.transfer.v1.endpoint.airbyte.LinkedinAdsSource.Credentials
	2, // 1: doublecloud.transfer.v1.endpoint.airbyte.LinkedinAdsSource.Credentials.oauth:type_name -> doublecloud.transfer.v1.endpoint.airbyte.LinkedinAdsSource.Credentials.OAuth
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes[1].OneofWrappers = []any{
		(*LinkedinAdsSource_Credentials_Oauth)(nil),
		(*LinkedinAdsSource_Credentials_AccessToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_linkedin_ads_source_proto_depIdxs = nil
}
