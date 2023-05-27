// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: doublecloud/visualization/v1/workbook.proto

package visualization

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlainSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *PlainSecret) Reset() {
	*x = PlainSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlainSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlainSecret) ProtoMessage() {}

func (x *PlainSecret) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlainSecret.ProtoReflect.Descriptor instead.
func (*PlainSecret) Descriptor() ([]byte, []int) {
	return file_doublecloud_visualization_v1_workbook_proto_rawDescGZIP(), []int{0}
}

func (x *PlainSecret) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Secret:
	//
	//	*Secret_PlainSecret
	Secret isSecret_Secret `protobuf_oneof:"secret"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_doublecloud_visualization_v1_workbook_proto_rawDescGZIP(), []int{1}
}

func (m *Secret) GetSecret() isSecret_Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (x *Secret) GetPlainSecret() *PlainSecret {
	if x, ok := x.GetSecret().(*Secret_PlainSecret); ok {
		return x.PlainSecret
	}
	return nil
}

type isSecret_Secret interface {
	isSecret_Secret()
}

type Secret_PlainSecret struct {
	PlainSecret *PlainSecret `protobuf:"bytes,1,opt,name=plain_secret,json=plainSecret,proto3,oneof"`
}

func (*Secret_PlainSecret) isSecret_Secret() {}

type Workbook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Workbook config according [spec](dcdoc://visualization-configs/Workbook)
	Config *structpb.Value `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Workbook) Reset() {
	*x = Workbook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workbook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workbook) ProtoMessage() {}

func (x *Workbook) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workbook.ProtoReflect.Descriptor instead.
func (*Workbook) Descriptor() ([]byte, []int) {
	return file_doublecloud_visualization_v1_workbook_proto_rawDescGZIP(), []int{2}
}

func (x *Workbook) GetConfig() *structpb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dataset config according [spec](dcdoc://visualization-configs/Dataset)
	Config *structpb.Value `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_doublecloud_visualization_v1_workbook_proto_rawDescGZIP(), []int{3}
}

func (x *Dataset) GetConfig() *structpb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Connection config according [spec](dcdoc://visualization-configs/Connection)
	Config *structpb.Value `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_visualization_v1_workbook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_doublecloud_visualization_v1_workbook_proto_rawDescGZIP(), []int{4}
}

func (x *Connection) GetConfig() *structpb.Value {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_doublecloud_visualization_v1_workbook_proto protoreflect.FileDescriptor

var file_doublecloud_visualization_v1_workbook_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x6c, 0x61,
	0x69, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x22, 0x62, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x4e, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6c, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x3a, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x6f, 0x6f, 0x6b,
	0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x39, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3c, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x73,
	0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_doublecloud_visualization_v1_workbook_proto_rawDescOnce sync.Once
	file_doublecloud_visualization_v1_workbook_proto_rawDescData = file_doublecloud_visualization_v1_workbook_proto_rawDesc
)

func file_doublecloud_visualization_v1_workbook_proto_rawDescGZIP() []byte {
	file_doublecloud_visualization_v1_workbook_proto_rawDescOnce.Do(func() {
		file_doublecloud_visualization_v1_workbook_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_visualization_v1_workbook_proto_rawDescData)
	})
	return file_doublecloud_visualization_v1_workbook_proto_rawDescData
}

var file_doublecloud_visualization_v1_workbook_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_doublecloud_visualization_v1_workbook_proto_goTypes = []interface{}{
	(*PlainSecret)(nil),    // 0: doublecloud.visualization.v1.PlainSecret
	(*Secret)(nil),         // 1: doublecloud.visualization.v1.Secret
	(*Workbook)(nil),       // 2: doublecloud.visualization.v1.Workbook
	(*Dataset)(nil),        // 3: doublecloud.visualization.v1.Dataset
	(*Connection)(nil),     // 4: doublecloud.visualization.v1.Connection
	(*structpb.Value)(nil), // 5: google.protobuf.Value
}
var file_doublecloud_visualization_v1_workbook_proto_depIdxs = []int32{
	0, // 0: doublecloud.visualization.v1.Secret.plain_secret:type_name -> doublecloud.visualization.v1.PlainSecret
	5, // 1: doublecloud.visualization.v1.Workbook.config:type_name -> google.protobuf.Value
	5, // 2: doublecloud.visualization.v1.Dataset.config:type_name -> google.protobuf.Value
	5, // 3: doublecloud.visualization.v1.Connection.config:type_name -> google.protobuf.Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_doublecloud_visualization_v1_workbook_proto_init() }
func file_doublecloud_visualization_v1_workbook_proto_init() {
	if File_doublecloud_visualization_v1_workbook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_visualization_v1_workbook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlainSecret); i {
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
		file_doublecloud_visualization_v1_workbook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_doublecloud_visualization_v1_workbook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workbook); i {
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
		file_doublecloud_visualization_v1_workbook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
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
		file_doublecloud_visualization_v1_workbook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
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
	file_doublecloud_visualization_v1_workbook_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Secret_PlainSecret)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_visualization_v1_workbook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_visualization_v1_workbook_proto_goTypes,
		DependencyIndexes: file_doublecloud_visualization_v1_workbook_proto_depIdxs,
		MessageInfos:      file_doublecloud_visualization_v1_workbook_proto_msgTypes,
	}.Build()
	File_doublecloud_visualization_v1_workbook_proto = out.File
	file_doublecloud_visualization_v1_workbook_proto_rawDesc = nil
	file_doublecloud_visualization_v1_workbook_proto_goTypes = nil
	file_doublecloud_visualization_v1_workbook_proto_depIdxs = nil
}