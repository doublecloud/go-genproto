// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: doublecloud/v1/paging.proto

package doublecloud

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

type Paging struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [NextPage.token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [NextPage.token]
	// returned by a previous list request.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Paging) Reset() {
	*x = Paging{}
	mi := &file_doublecloud_v1_paging_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paging) ProtoMessage() {}

func (x *Paging) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_paging_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paging.ProtoReflect.Descriptor instead.
func (*Paging) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_paging_proto_rawDescGZIP(), []int{0}
}

func (x *Paging) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Paging) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type NextPage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This token allows you to get the next page of results for list requests. If the
	// number of results
	// is larger than [Paging.page_size], use the [token] as the value
	// for the [Paging.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [token] to continue paging through the results.
	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NextPage) Reset() {
	*x = NextPage{}
	mi := &file_doublecloud_v1_paging_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NextPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextPage) ProtoMessage() {}

func (x *NextPage) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_paging_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextPage.ProtoReflect.Descriptor instead.
func (*NextPage) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_paging_proto_rawDescGZIP(), []int{1}
}

func (x *NextPage) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_doublecloud_v1_paging_proto protoreflect.FileDescriptor

var file_doublecloud_v1_paging_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x44, 0x0a,
	0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x20, 0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_doublecloud_v1_paging_proto_rawDescOnce sync.Once
	file_doublecloud_v1_paging_proto_rawDescData []byte
)

func file_doublecloud_v1_paging_proto_rawDescGZIP() []byte {
	file_doublecloud_v1_paging_proto_rawDescOnce.Do(func() {
		file_doublecloud_v1_paging_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_v1_paging_proto_rawDesc), len(file_doublecloud_v1_paging_proto_rawDesc)))
	})
	return file_doublecloud_v1_paging_proto_rawDescData
}

var file_doublecloud_v1_paging_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_v1_paging_proto_goTypes = []any{
	(*Paging)(nil),   // 0: doublecloud.v1.Paging
	(*NextPage)(nil), // 1: doublecloud.v1.NextPage
}
var file_doublecloud_v1_paging_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doublecloud_v1_paging_proto_init() }
func file_doublecloud_v1_paging_proto_init() {
	if File_doublecloud_v1_paging_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_v1_paging_proto_rawDesc), len(file_doublecloud_v1_paging_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_v1_paging_proto_goTypes,
		DependencyIndexes: file_doublecloud_v1_paging_proto_depIdxs,
		MessageInfos:      file_doublecloud_v1_paging_proto_msgTypes,
	}.Build()
	File_doublecloud_v1_paging_proto = out.File
	file_doublecloud_v1_paging_proto_goTypes = nil
	file_doublecloud_v1_paging_proto_depIdxs = nil
}
