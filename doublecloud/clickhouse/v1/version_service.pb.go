// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: doublecloud/clickhouse/v1/version_service.proto

package clickhouse

import (
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
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

type ListVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information of the request
	Paging *v1.Paging `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *ListVersionsRequest) Reset() {
	*x = ListVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_clickhouse_v1_version_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVersionsRequest) ProtoMessage() {}

func (x *ListVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_version_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVersionsRequest.ProtoReflect.Descriptor instead.
func (*ListVersionsRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_version_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListVersionsRequest) GetPaging() *v1.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type ListVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested list of available versions.
	Versions []*Version `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
	// Pagination information of the response
	NextPage *v1.NextPage `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *ListVersionsResponse) Reset() {
	*x = ListVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_clickhouse_v1_version_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVersionsResponse) ProtoMessage() {}

func (x *ListVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_version_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVersionsResponse.ProtoReflect.Descriptor instead.
func (*ListVersionsResponse) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_version_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListVersionsResponse) GetVersions() []*Version {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *ListVersionsResponse) GetNextPage() *v1.NextPage {
	if x != nil {
		return x.NextPage
	}
	return nil
}

var File_doublecloud_clickhouse_v1_version_service_proto protoreflect.FileDescriptor

var file_doublecloud_clickhouse_v1_version_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x45, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x8d, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x32, 0x79, 0x0a, 0x0e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_clickhouse_v1_version_service_proto_rawDescOnce sync.Once
	file_doublecloud_clickhouse_v1_version_service_proto_rawDescData = file_doublecloud_clickhouse_v1_version_service_proto_rawDesc
)

func file_doublecloud_clickhouse_v1_version_service_proto_rawDescGZIP() []byte {
	file_doublecloud_clickhouse_v1_version_service_proto_rawDescOnce.Do(func() {
		file_doublecloud_clickhouse_v1_version_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_clickhouse_v1_version_service_proto_rawDescData)
	})
	return file_doublecloud_clickhouse_v1_version_service_proto_rawDescData
}

var file_doublecloud_clickhouse_v1_version_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_clickhouse_v1_version_service_proto_goTypes = []interface{}{
	(*ListVersionsRequest)(nil),  // 0: doublecloud.clickhouse.v1.ListVersionsRequest
	(*ListVersionsResponse)(nil), // 1: doublecloud.clickhouse.v1.ListVersionsResponse
	(*v1.Paging)(nil),            // 2: doublecloud.v1.Paging
	(*Version)(nil),              // 3: doublecloud.clickhouse.v1.Version
	(*v1.NextPage)(nil),          // 4: doublecloud.v1.NextPage
}
var file_doublecloud_clickhouse_v1_version_service_proto_depIdxs = []int32{
	2, // 0: doublecloud.clickhouse.v1.ListVersionsRequest.paging:type_name -> doublecloud.v1.Paging
	3, // 1: doublecloud.clickhouse.v1.ListVersionsResponse.versions:type_name -> doublecloud.clickhouse.v1.Version
	4, // 2: doublecloud.clickhouse.v1.ListVersionsResponse.next_page:type_name -> doublecloud.v1.NextPage
	0, // 3: doublecloud.clickhouse.v1.VersionService.List:input_type -> doublecloud.clickhouse.v1.ListVersionsRequest
	1, // 4: doublecloud.clickhouse.v1.VersionService.List:output_type -> doublecloud.clickhouse.v1.ListVersionsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_doublecloud_clickhouse_v1_version_service_proto_init() }
func file_doublecloud_clickhouse_v1_version_service_proto_init() {
	if File_doublecloud_clickhouse_v1_version_service_proto != nil {
		return
	}
	file_doublecloud_clickhouse_v1_version_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_clickhouse_v1_version_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVersionsRequest); i {
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
		file_doublecloud_clickhouse_v1_version_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVersionsResponse); i {
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
			RawDescriptor: file_doublecloud_clickhouse_v1_version_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doublecloud_clickhouse_v1_version_service_proto_goTypes,
		DependencyIndexes: file_doublecloud_clickhouse_v1_version_service_proto_depIdxs,
		MessageInfos:      file_doublecloud_clickhouse_v1_version_service_proto_msgTypes,
	}.Build()
	File_doublecloud_clickhouse_v1_version_service_proto = out.File
	file_doublecloud_clickhouse_v1_version_service_proto_rawDesc = nil
	file_doublecloud_clickhouse_v1_version_service_proto_goTypes = nil
	file_doublecloud_clickhouse_v1_version_service_proto_depIdxs = nil
}
