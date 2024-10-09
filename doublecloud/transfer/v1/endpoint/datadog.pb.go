// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/datadog.proto

package endpoint

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

type DatadogTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientApiKey  string                       `protobuf:"bytes,1,opt,name=client_api_key,json=clientApiKey,proto3" json:"client_api_key,omitempty"`
	Host          string                       `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	ColumnMapping *DatadogTarget_ColumnMapping `protobuf:"bytes,3,opt,name=column_mapping,json=columnMapping,proto3" json:"column_mapping,omitempty"`
}

func (x *DatadogTarget) Reset() {
	*x = DatadogTarget{}
	mi := &file_doublecloud_transfer_v1_endpoint_datadog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatadogTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatadogTarget) ProtoMessage() {}

func (x *DatadogTarget) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_datadog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatadogTarget.ProtoReflect.Descriptor instead.
func (*DatadogTarget) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescGZIP(), []int{0}
}

func (x *DatadogTarget) GetClientApiKey() string {
	if x != nil {
		return x.ClientApiKey
	}
	return ""
}

func (x *DatadogTarget) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *DatadogTarget) GetColumnMapping() *DatadogTarget_ColumnMapping {
	if x != nil {
		return x.ColumnMapping
	}
	return nil
}

type DatadogTarget_ColumnMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source          string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Tags            []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Host            string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	MessageTemplate string   `protobuf:"bytes,4,opt,name=message_template,json=messageTemplate,proto3" json:"message_template,omitempty"`
	Service         string   `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *DatadogTarget_ColumnMapping) Reset() {
	*x = DatadogTarget_ColumnMapping{}
	mi := &file_doublecloud_transfer_v1_endpoint_datadog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DatadogTarget_ColumnMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatadogTarget_ColumnMapping) ProtoMessage() {}

func (x *DatadogTarget_ColumnMapping) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_datadog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatadogTarget_ColumnMapping.ProtoReflect.Descriptor instead.
func (*DatadogTarget_ColumnMapping) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DatadogTarget_ColumnMapping) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *DatadogTarget_ColumnMapping) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *DatadogTarget_ColumnMapping) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *DatadogTarget_ColumnMapping) GetMessageTemplate() string {
	if x != nil {
		return x.MessageTemplate
	}
	return ""
}

func (x *DatadogTarget_ColumnMapping) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_datadog_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x22, 0xc6, 0x02, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x64,
	0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x1a, 0x94, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_datadog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_transfer_v1_endpoint_datadog_proto_goTypes = []any{
	(*DatadogTarget)(nil),               // 0: doublecloud.transfer.v1.endpoint.DatadogTarget
	(*DatadogTarget_ColumnMapping)(nil), // 1: doublecloud.transfer.v1.endpoint.DatadogTarget.ColumnMapping
}
var file_doublecloud_transfer_v1_endpoint_datadog_proto_depIdxs = []int32{
	1, // 0: doublecloud.transfer.v1.endpoint.DatadogTarget.column_mapping:type_name -> doublecloud.transfer.v1.endpoint.DatadogTarget.ColumnMapping
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_datadog_proto_init() }
func file_doublecloud_transfer_v1_endpoint_datadog_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_datadog_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_datadog_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_datadog_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_datadog_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_datadog_proto = out.File
	file_doublecloud_transfer_v1_endpoint_datadog_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_datadog_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_datadog_proto_depIdxs = nil
}
