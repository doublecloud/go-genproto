// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: doublecloud/logs/v1/log_source.proto

package logs

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

type LogSourceType int32

const (
	LogSourceType_LOG_SOURCE_TYPE_INVALID    LogSourceType = 0
	LogSourceType_LOG_SOURCE_TYPE_CLICKHOUSE LogSourceType = 1
	LogSourceType_LOG_SOURCE_TYPE_KAFKA      LogSourceType = 2
	LogSourceType_LOG_SOURCE_TYPE_TRANSFER   LogSourceType = 3
)

// Enum value maps for LogSourceType.
var (
	LogSourceType_name = map[int32]string{
		0: "LOG_SOURCE_TYPE_INVALID",
		1: "LOG_SOURCE_TYPE_CLICKHOUSE",
		2: "LOG_SOURCE_TYPE_KAFKA",
		3: "LOG_SOURCE_TYPE_TRANSFER",
	}
	LogSourceType_value = map[string]int32{
		"LOG_SOURCE_TYPE_INVALID":    0,
		"LOG_SOURCE_TYPE_CLICKHOUSE": 1,
		"LOG_SOURCE_TYPE_KAFKA":      2,
		"LOG_SOURCE_TYPE_TRANSFER":   3,
	}
)

func (x LogSourceType) Enum() *LogSourceType {
	p := new(LogSourceType)
	*p = x
	return p
}

func (x LogSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_logs_v1_log_source_proto_enumTypes[0].Descriptor()
}

func (LogSourceType) Type() protoreflect.EnumType {
	return &file_doublecloud_logs_v1_log_source_proto_enumTypes[0]
}

func (x LogSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogSourceType.Descriptor instead.
func (LogSourceType) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_source_proto_rawDescGZIP(), []int{0}
}

type LogSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source service that produces logs
	Type LogSourceType `protobuf:"varint,1,opt,name=type,proto3,enum=doublecloud.logs.v1.LogSourceType" json:"type,omitempty"`
	// Log source resource id.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LogSource) Reset() {
	*x = LogSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSource) ProtoMessage() {}

func (x *LogSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSource.ProtoReflect.Descriptor instead.
func (*LogSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_source_proto_rawDescGZIP(), []int{0}
}

func (x *LogSource) GetType() LogSourceType {
	if x != nil {
		return x.Type
	}
	return LogSourceType_LOG_SOURCE_TYPE_INVALID
}

func (x *LogSource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_doublecloud_logs_v1_log_source_proto protoreflect.FileDescriptor

var file_doublecloud_logs_v1_log_source_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x53, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x2a, 0x85, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x1e, 0x0a, 0x1a, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x48, 0x4f, 0x55, 0x53, 0x45, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x4f,
	0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x03, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_logs_v1_log_source_proto_rawDescOnce sync.Once
	file_doublecloud_logs_v1_log_source_proto_rawDescData = file_doublecloud_logs_v1_log_source_proto_rawDesc
)

func file_doublecloud_logs_v1_log_source_proto_rawDescGZIP() []byte {
	file_doublecloud_logs_v1_log_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_logs_v1_log_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_logs_v1_log_source_proto_rawDescData)
	})
	return file_doublecloud_logs_v1_log_source_proto_rawDescData
}

var file_doublecloud_logs_v1_log_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_logs_v1_log_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_logs_v1_log_source_proto_goTypes = []interface{}{
	(LogSourceType)(0), // 0: doublecloud.logs.v1.LogSourceType
	(*LogSource)(nil),  // 1: doublecloud.logs.v1.LogSource
}
var file_doublecloud_logs_v1_log_source_proto_depIdxs = []int32{
	0, // 0: doublecloud.logs.v1.LogSource.type:type_name -> doublecloud.logs.v1.LogSourceType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_doublecloud_logs_v1_log_source_proto_init() }
func file_doublecloud_logs_v1_log_source_proto_init() {
	if File_doublecloud_logs_v1_log_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_logs_v1_log_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSource); i {
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
			RawDescriptor: file_doublecloud_logs_v1_log_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_logs_v1_log_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_logs_v1_log_source_proto_depIdxs,
		EnumInfos:         file_doublecloud_logs_v1_log_source_proto_enumTypes,
		MessageInfos:      file_doublecloud_logs_v1_log_source_proto_msgTypes,
	}.Build()
	File_doublecloud_logs_v1_log_source_proto = out.File
	file_doublecloud_logs_v1_log_source_proto_rawDesc = nil
	file_doublecloud_logs_v1_log_source_proto_goTypes = nil
	file_doublecloud_logs_v1_log_source_proto_depIdxs = nil
}
