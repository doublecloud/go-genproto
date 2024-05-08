// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: doublecloud/kafka/v1/config.proto

package kafka

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

// Configuration for Apache Kafka® brokers in the cluster.
type KafkaConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The largest record batch size allowed by Kafka. Default value: 1048588.
	MessageMaxBytes *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=message_max_bytes,json=messageMaxBytes,proto3" json:"message_max_bytes,omitempty"`
	// The number of bytes of messages to attempt to fetch for each partition. Default
	// value: 1048576.
	ReplicaFetchMaxBytes *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=replica_fetch_max_bytes,json=replicaFetchMaxBytes,proto3" json:"replica_fetch_max_bytes,omitempty"`
	// The maximum size of the log before deleting it. Default is -1.
	LogRetentionBytes *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=log_retention_bytes,json=logRetentionBytes,proto3" json:"log_retention_bytes,omitempty"`
	// The number of hours to keep a log file before deleting it (in hours), tertiary
	// to log.retention.ms property. Default: 168 hours.
	LogRetentionHours *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=log_retention_hours,json=logRetentionHours,proto3" json:"log_retention_hours,omitempty"`
	// The number of minutes to keep a log file before deleting it (in minutes),
	// secondary to log.retention.ms property. If not set, the value in
	// log.retention.hours is used.
	LogRetentionMinutes *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=log_retention_minutes,json=logRetentionMinutes,proto3" json:"log_retention_minutes,omitempty"`
	// The number of milliseconds to keep a log file before deleting it (in
	// milliseconds), If not set, the value in log.retention.minutes is used. If set to
	// -1, no time limit is applied.
	LogRetentionMs *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=log_retention_ms,json=logRetentionMs,proto3" json:"log_retention_ms,omitempty"`
}

func (x *KafkaConfig) Reset() {
	*x = KafkaConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_kafka_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaConfig) ProtoMessage() {}

func (x *KafkaConfig) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaConfig.ProtoReflect.Descriptor instead.
func (*KafkaConfig) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *KafkaConfig) GetMessageMaxBytes() *wrapperspb.Int64Value {
	if x != nil {
		return x.MessageMaxBytes
	}
	return nil
}

func (x *KafkaConfig) GetReplicaFetchMaxBytes() *wrapperspb.Int64Value {
	if x != nil {
		return x.ReplicaFetchMaxBytes
	}
	return nil
}

func (x *KafkaConfig) GetLogRetentionBytes() *wrapperspb.Int64Value {
	if x != nil {
		return x.LogRetentionBytes
	}
	return nil
}

func (x *KafkaConfig) GetLogRetentionHours() *wrapperspb.Int64Value {
	if x != nil {
		return x.LogRetentionHours
	}
	return nil
}

func (x *KafkaConfig) GetLogRetentionMinutes() *wrapperspb.Int64Value {
	if x != nil {
		return x.LogRetentionMinutes
	}
	return nil
}

func (x *KafkaConfig) GetLogRetentionMs() *wrapperspb.Int64Value {
	if x != nil {
		return x.LogRetentionMs
	}
	return nil
}

// Configuration for Schema Registry.
type SchemaRegistryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Is Schema Registry enabled for this cluster.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *SchemaRegistryConfig) Reset() {
	*x = SchemaRegistryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_kafka_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaRegistryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaRegistryConfig) ProtoMessage() {}

func (x *SchemaRegistryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaRegistryConfig.ProtoReflect.Descriptor instead.
func (*SchemaRegistryConfig) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *SchemaRegistryConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// Configuration for REST API.
type RestAPIConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Is REST API enabled for this cluster.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *RestAPIConfig) Reset() {
	*x = RestAPIConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_kafka_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestAPIConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestAPIConfig) ProtoMessage() {}

func (x *RestAPIConfig) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestAPIConfig.ProtoReflect.Descriptor instead.
func (*RestAPIConfig) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *RestAPIConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_doublecloud_kafka_v1_config_proto protoreflect.FileDescriptor

var file_doublecloud_kafka_v1_config_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x0b, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x11, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x52, 0x0a, 0x17, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x14, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4d, 0x61,
	0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x11, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73,
	0x12, 0x4f, 0x0a, 0x15, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x45, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x74, 0x41, 0x50, 0x49, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31,
	0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_kafka_v1_config_proto_rawDescOnce sync.Once
	file_doublecloud_kafka_v1_config_proto_rawDescData = file_doublecloud_kafka_v1_config_proto_rawDesc
)

func file_doublecloud_kafka_v1_config_proto_rawDescGZIP() []byte {
	file_doublecloud_kafka_v1_config_proto_rawDescOnce.Do(func() {
		file_doublecloud_kafka_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_kafka_v1_config_proto_rawDescData)
	})
	return file_doublecloud_kafka_v1_config_proto_rawDescData
}

var file_doublecloud_kafka_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_doublecloud_kafka_v1_config_proto_goTypes = []interface{}{
	(*KafkaConfig)(nil),           // 0: doublecloud.kafka.v1.KafkaConfig
	(*SchemaRegistryConfig)(nil),  // 1: doublecloud.kafka.v1.SchemaRegistryConfig
	(*RestAPIConfig)(nil),         // 2: doublecloud.kafka.v1.RestAPIConfig
	(*wrapperspb.Int64Value)(nil), // 3: google.protobuf.Int64Value
}
var file_doublecloud_kafka_v1_config_proto_depIdxs = []int32{
	3, // 0: doublecloud.kafka.v1.KafkaConfig.message_max_bytes:type_name -> google.protobuf.Int64Value
	3, // 1: doublecloud.kafka.v1.KafkaConfig.replica_fetch_max_bytes:type_name -> google.protobuf.Int64Value
	3, // 2: doublecloud.kafka.v1.KafkaConfig.log_retention_bytes:type_name -> google.protobuf.Int64Value
	3, // 3: doublecloud.kafka.v1.KafkaConfig.log_retention_hours:type_name -> google.protobuf.Int64Value
	3, // 4: doublecloud.kafka.v1.KafkaConfig.log_retention_minutes:type_name -> google.protobuf.Int64Value
	3, // 5: doublecloud.kafka.v1.KafkaConfig.log_retention_ms:type_name -> google.protobuf.Int64Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_doublecloud_kafka_v1_config_proto_init() }
func file_doublecloud_kafka_v1_config_proto_init() {
	if File_doublecloud_kafka_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_kafka_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaConfig); i {
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
		file_doublecloud_kafka_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaRegistryConfig); i {
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
		file_doublecloud_kafka_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestAPIConfig); i {
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
			RawDescriptor: file_doublecloud_kafka_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_kafka_v1_config_proto_goTypes,
		DependencyIndexes: file_doublecloud_kafka_v1_config_proto_depIdxs,
		MessageInfos:      file_doublecloud_kafka_v1_config_proto_msgTypes,
	}.Build()
	File_doublecloud_kafka_v1_config_proto = out.File
	file_doublecloud_kafka_v1_config_proto_rawDesc = nil
	file_doublecloud_kafka_v1_config_proto_goTypes = nil
	file_doublecloud_kafka_v1_config_proto_depIdxs = nil
}
