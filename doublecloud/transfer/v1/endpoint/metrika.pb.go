// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/metrika.proto

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

type MetrikaStreamType int32

const (
	MetrikaStreamType_METRIKA_STREAM_TYPE_UNSPECIFIED MetrikaStreamType = 0
	MetrikaStreamType_METRIKA_STREAM_TYPE_VISITS      MetrikaStreamType = 1
	MetrikaStreamType_METRIKA_STREAM_TYPE_HITS_V2     MetrikaStreamType = 2
)

// Enum value maps for MetrikaStreamType.
var (
	MetrikaStreamType_name = map[int32]string{
		0: "METRIKA_STREAM_TYPE_UNSPECIFIED",
		1: "METRIKA_STREAM_TYPE_VISITS",
		2: "METRIKA_STREAM_TYPE_HITS_V2",
	}
	MetrikaStreamType_value = map[string]int32{
		"METRIKA_STREAM_TYPE_UNSPECIFIED": 0,
		"METRIKA_STREAM_TYPE_VISITS":      1,
		"METRIKA_STREAM_TYPE_HITS_V2":     2,
	}
)

func (x MetrikaStreamType) Enum() *MetrikaStreamType {
	p := new(MetrikaStreamType)
	*p = x
	return p
}

func (x MetrikaStreamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetrikaStreamType) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_metrika_proto_enumTypes[0].Descriptor()
}

func (MetrikaStreamType) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_metrika_proto_enumTypes[0]
}

func (x MetrikaStreamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetrikaStreamType.Descriptor instead.
func (MetrikaStreamType) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescGZIP(), []int{0}
}

type MetrikaStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    MetrikaStreamType `protobuf:"varint,1,opt,name=type,proto3,enum=doublecloud.transfer.v1.endpoint.MetrikaStreamType" json:"type,omitempty"`
	Columns []string          `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *MetrikaStream) Reset() {
	*x = MetrikaStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetrikaStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetrikaStream) ProtoMessage() {}

func (x *MetrikaStream) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetrikaStream.ProtoReflect.Descriptor instead.
func (*MetrikaStream) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescGZIP(), []int{0}
}

func (x *MetrikaStream) GetType() MetrikaStreamType {
	if x != nil {
		return x.Type
	}
	return MetrikaStreamType_METRIKA_STREAM_TYPE_UNSPECIFIED
}

func (x *MetrikaStream) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

type MetrikaSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CounterIds []int64          `protobuf:"varint,1,rep,packed,name=counter_ids,json=counterIds,proto3" json:"counter_ids,omitempty"`
	Token      *Secret          `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Streams    []*MetrikaStream `protobuf:"bytes,3,rep,name=streams,proto3" json:"streams,omitempty"`
}

func (x *MetrikaSource) Reset() {
	*x = MetrikaSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetrikaSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetrikaSource) ProtoMessage() {}

func (x *MetrikaSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetrikaSource.ProtoReflect.Descriptor instead.
func (*MetrikaSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescGZIP(), []int{1}
}

func (x *MetrikaSource) GetCounterIds() []int64 {
	if x != nil {
		return x.CounterIds
	}
	return nil
}

func (x *MetrikaSource) GetToken() *Secret {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *MetrikaSource) GetStreams() []*MetrikaStream {
	if x != nil {
		return x.Streams
	}
	return nil
}

var File_doublecloud_transfer_v1_endpoint_metrika_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x1a, 0x2d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x72, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x6b, 0x61, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x33, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x6b, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x6b,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x49, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x6b, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2a, 0x79, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x6b, 0x61, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x4b, 0x41, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a,
	0x1a, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x4b, 0x41, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x54, 0x53, 0x10, 0x01, 0x12, 0x1f, 0x0a,
	0x1b, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x4b, 0x41, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x54, 0x53, 0x5f, 0x56, 0x32, 0x10, 0x02, 0x42, 0x4e,
	0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_metrika_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_transfer_v1_endpoint_metrika_proto_goTypes = []interface{}{
	(MetrikaStreamType)(0), // 0: doublecloud.transfer.v1.endpoint.MetrikaStreamType
	(*MetrikaStream)(nil),  // 1: doublecloud.transfer.v1.endpoint.MetrikaStream
	(*MetrikaSource)(nil),  // 2: doublecloud.transfer.v1.endpoint.MetrikaSource
	(*Secret)(nil),         // 3: doublecloud.transfer.v1.endpoint.Secret
}
var file_doublecloud_transfer_v1_endpoint_metrika_proto_depIdxs = []int32{
	0, // 0: doublecloud.transfer.v1.endpoint.MetrikaStream.type:type_name -> doublecloud.transfer.v1.endpoint.MetrikaStreamType
	3, // 1: doublecloud.transfer.v1.endpoint.MetrikaSource.token:type_name -> doublecloud.transfer.v1.endpoint.Secret
	1, // 2: doublecloud.transfer.v1.endpoint.MetrikaSource.streams:type_name -> doublecloud.transfer.v1.endpoint.MetrikaStream
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_metrika_proto_init() }
func file_doublecloud_transfer_v1_endpoint_metrika_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_metrika_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetrikaStream); i {
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
		file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetrikaSource); i {
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
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_metrika_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_metrika_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_endpoint_metrika_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_metrika_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_metrika_proto = out.File
	file_doublecloud_transfer_v1_endpoint_metrika_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_metrika_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_metrika_proto_depIdxs = nil
}