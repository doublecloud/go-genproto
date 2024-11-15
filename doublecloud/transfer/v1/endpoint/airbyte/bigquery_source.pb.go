// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/bigquery_source.proto

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

type BigQuerySource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId       string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DatasetId       string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	CredentialsJson string `protobuf:"bytes,3,opt,name=credentials_json,json=credentialsJson,proto3" json:"credentials_json,omitempty"`
}

func (x *BigQuerySource) Reset() {
	*x = BigQuerySource{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BigQuerySource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQuerySource) ProtoMessage() {}

func (x *BigQuerySource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQuerySource.ProtoReflect.Descriptor instead.
func (*BigQuerySource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescGZIP(), []int{0}
}

func (x *BigQuerySource) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BigQuerySource) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *BigQuerySource) GetCredentialsJson() string {
	if x != nil {
		return x.CredentialsJson
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0x79, 0x0a, 0x0e, 0x42, 0x69,
	0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72,
	0x62, 0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69,
	0x72, 0x62, 0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_goTypes = []any{
	(*BigQuerySource)(nil), // 0: doublecloud.transfer.v1.endpoint.airbyte.BigQuerySource
}
var file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_bigquery_source_proto_depIdxs = nil
}
