// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/redshift_source.proto

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

type RedshiftSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Database string   `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	Schemas  []string `protobuf:"bytes,4,rep,name=schemas,proto3" json:"schemas,omitempty"`
	Username string   `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RedshiftSource) Reset() {
	*x = RedshiftSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedshiftSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedshiftSource) ProtoMessage() {}

func (x *RedshiftSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedshiftSource.ProtoReflect.Descriptor instead.
func (*RedshiftSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescGZIP(), []int{0}
}

func (x *RedshiftSource) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RedshiftSource) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RedshiftSource) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *RedshiftSource) GetSchemas() []string {
	if x != nil {
		return x.Schemas
	}
	return nil
}

func (x *RedshiftSource) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RedshiftSource) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x64, 0x73, 0x68,
	0x69, 0x66, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x52,
	0x65, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79,
	0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_goTypes = []interface{}{
	(*RedshiftSource)(nil), // 0: doublecloud.transfer.v1.endpoint.airbyte.RedshiftSource
}
var file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedshiftSource); i {
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
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_redshift_source_proto_depIdxs = nil
}
