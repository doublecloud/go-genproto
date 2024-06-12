// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: doublecloud/kafka/v1/version.proto

package kafka

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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the version.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the version.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Whether version is deprecated.
	Deprecated bool `protobuf:"varint,3,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	// List of versions that can be updated from current.
	UpdatableTo []string `protobuf:"bytes,4,rep,name=updatable_to,json=updatableTo,proto3" json:"updatable_to,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_kafka_v1_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Version) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *Version) GetUpdatableTo() []string {
	if x != nil {
		return x.UpdatableTo
	}
	return nil
}

var File_doublecloud_kafka_v1_version_proto protoreflect.FileDescriptor

var file_doublecloud_kafka_v1_version_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x70, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_kafka_v1_version_proto_rawDescOnce sync.Once
	file_doublecloud_kafka_v1_version_proto_rawDescData = file_doublecloud_kafka_v1_version_proto_rawDesc
)

func file_doublecloud_kafka_v1_version_proto_rawDescGZIP() []byte {
	file_doublecloud_kafka_v1_version_proto_rawDescOnce.Do(func() {
		file_doublecloud_kafka_v1_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_kafka_v1_version_proto_rawDescData)
	})
	return file_doublecloud_kafka_v1_version_proto_rawDescData
}

var file_doublecloud_kafka_v1_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_kafka_v1_version_proto_goTypes = []any{
	(*Version)(nil), // 0: doublecloud.kafka.v1.Version
}
var file_doublecloud_kafka_v1_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doublecloud_kafka_v1_version_proto_init() }
func file_doublecloud_kafka_v1_version_proto_init() {
	if File_doublecloud_kafka_v1_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_kafka_v1_version_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_doublecloud_kafka_v1_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_kafka_v1_version_proto_goTypes,
		DependencyIndexes: file_doublecloud_kafka_v1_version_proto_depIdxs,
		MessageInfos:      file_doublecloud_kafka_v1_version_proto_msgTypes,
	}.Build()
	File_doublecloud_kafka_v1_version_proto = out.File
	file_doublecloud_kafka_v1_version_proto_rawDesc = nil
	file_doublecloud_kafka_v1_version_proto_goTypes = nil
	file_doublecloud_kafka_v1_version_proto_depIdxs = nil
}
