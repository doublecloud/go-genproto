// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/aws_cloud_trail_source.proto

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

type AWSCloudTrailSource struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AwsKeyId      string                 `protobuf:"bytes,1,opt,name=aws_key_id,json=awsKeyId,proto3" json:"aws_key_id,omitempty"`
	AwsSecretKey  string                 `protobuf:"bytes,2,opt,name=aws_secret_key,json=awsSecretKey,proto3" json:"aws_secret_key,omitempty"`
	AwsRegionName string                 `protobuf:"bytes,3,opt,name=aws_region_name,json=awsRegionName,proto3" json:"aws_region_name,omitempty"`
	StartDate     string                 `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AWSCloudTrailSource) Reset() {
	*x = AWSCloudTrailSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AWSCloudTrailSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSCloudTrailSource) ProtoMessage() {}

func (x *AWSCloudTrailSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSCloudTrailSource.ProtoReflect.Descriptor instead.
func (*AWSCloudTrailSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDescGZIP(), []int{0}
}

func (x *AWSCloudTrailSource) GetAwsKeyId() string {
	if x != nil {
		return x.AwsKeyId
	}
	return ""
}

func (x *AWSCloudTrailSource) GetAwsSecretKey() string {
	if x != nil {
		return x.AwsSecretKey
	}
	return ""
}

func (x *AWSCloudTrailSource) GetAwsRegionName() string {
	if x != nil {
		return x.AwsRegionName
	}
	return ""
}

func (x *AWSCloudTrailSource) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDesc = string([]byte{
	0x0a, 0x45, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74,
	0x65, 0x22, 0xa0, 0x01, 0x0a, 0x13, 0x41, 0x57, 0x53, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x77, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x77, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a,
	0x0f, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x77, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62,
	0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDescData []byte
)

func file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDesc)))
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_goTypes = []any{
	(*AWSCloudTrailSource)(nil), // 0: doublecloud.transfer.v1.endpoint.airbyte.AWSCloudTrailSource
}
var file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_aws_cloud_trail_source_proto_depIdxs = nil
}
