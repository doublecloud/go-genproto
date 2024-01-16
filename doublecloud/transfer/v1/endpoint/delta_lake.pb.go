// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/delta_lake.proto

package endpoint

import (
	airbyte "github.com/doublecloud/go-genproto/doublecloud/transfer/v1/endpoint/airbyte"
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

type DeltaLakeSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider *airbyte.S3Source_Provider `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Table    *DeltaResultTable          `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Settings *DeltaSettings             `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *DeltaLakeSource) Reset() {
	*x = DeltaLakeSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaLakeSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaLakeSource) ProtoMessage() {}

func (x *DeltaLakeSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaLakeSource.ProtoReflect.Descriptor instead.
func (*DeltaLakeSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescGZIP(), []int{0}
}

func (x *DeltaLakeSource) GetProvider() *airbyte.S3Source_Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

func (x *DeltaLakeSource) GetTable() *DeltaResultTable {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *DeltaLakeSource) GetSettings() *DeltaSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type DeltaResultTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeltaResultTable) Reset() {
	*x = DeltaResultTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaResultTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaResultTable) ProtoMessage() {}

func (x *DeltaResultTable) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaResultTable.ProtoReflect.Descriptor instead.
func (*DeltaResultTable) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescGZIP(), []int{1}
}

type DeltaSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *DeltaSettings) Reset() {
	*x = DeltaSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaSettings) ProtoMessage() {}

func (x *DeltaSettings) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaSettings.ProtoReflect.Descriptor instead.
func (*DeltaSettings) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescGZIP(), []int{2}
}

func (x *DeltaSettings) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_delta_lake_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDesc = []byte{
	0x0a, 0x31, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x38, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f,
	0x73, 0x33, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x81, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4c, 0x61, 0x6b, 0x65, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65,
	0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_doublecloud_transfer_v1_endpoint_delta_lake_proto_goTypes = []interface{}{
	(*DeltaLakeSource)(nil),           // 0: doublecloud.transfer.v1.endpoint.DeltaLakeSource
	(*DeltaResultTable)(nil),          // 1: doublecloud.transfer.v1.endpoint.DeltaResultTable
	(*DeltaSettings)(nil),             // 2: doublecloud.transfer.v1.endpoint.DeltaSettings
	(*airbyte.S3Source_Provider)(nil), // 3: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Provider
}
var file_doublecloud_transfer_v1_endpoint_delta_lake_proto_depIdxs = []int32{
	3, // 0: doublecloud.transfer.v1.endpoint.DeltaLakeSource.provider:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Provider
	1, // 1: doublecloud.transfer.v1.endpoint.DeltaLakeSource.table:type_name -> doublecloud.transfer.v1.endpoint.DeltaResultTable
	2, // 2: doublecloud.transfer.v1.endpoint.DeltaLakeSource.settings:type_name -> doublecloud.transfer.v1.endpoint.DeltaSettings
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_delta_lake_proto_init() }
func file_doublecloud_transfer_v1_endpoint_delta_lake_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_delta_lake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaLakeSource); i {
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
		file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaResultTable); i {
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
		file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaSettings); i {
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
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_delta_lake_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_delta_lake_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_delta_lake_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_delta_lake_proto = out.File
	file_doublecloud_transfer_v1_endpoint_delta_lake_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_delta_lake_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_delta_lake_proto_depIdxs = nil
}
