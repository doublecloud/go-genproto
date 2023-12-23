// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/amazon_ads_source.proto

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

type AmazonAdsSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string    `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string    `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Scope        string    `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	RefreshToken string    `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	StartDate    string    `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	Region       string    `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	Profiles     []float64 `protobuf:"fixed64,7,rep,packed,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *AmazonAdsSource) Reset() {
	*x = AmazonAdsSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmazonAdsSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmazonAdsSource) ProtoMessage() {}

func (x *AmazonAdsSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmazonAdsSource.ProtoReflect.Descriptor instead.
func (*AmazonAdsSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescGZIP(), []int{0}
}

func (x *AmazonAdsSource) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AmazonAdsSource) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *AmazonAdsSource) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *AmazonAdsSource) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AmazonAdsSource) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *AmazonAdsSource) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AmazonAdsSource) GetProfiles() []float64 {
	if x != nil {
		return x.Profiles
	}
	return nil
}

var File_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDesc = []byte{
	0x0a, 0x40, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x61, 0x6d, 0x61, 0x7a, 0x6f,
	0x6e, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0xe1, 0x01, 0x0a,
	0x0f, 0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x41, 0x64, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x3b,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_goTypes = []interface{}{
	(*AmazonAdsSource)(nil), // 0: doublecloud.transfer.v1.endpoint.airbyte.AmazonAdsSource
}
var file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmazonAdsSource); i {
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
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_amazon_ads_source_proto_depIdxs = nil
}
