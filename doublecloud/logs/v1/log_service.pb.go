// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: doublecloud/logs/v1/log_service.proto

package logs

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

type ReadLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Selector:
	//
	//	*ReadLogsRequest_Paging
	//	*ReadLogsRequest_Criteria
	Selector isReadLogsRequest_Selector `protobuf_oneof:"selector"`
}

func (x *ReadLogsRequest) Reset() {
	*x = ReadLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadLogsRequest) ProtoMessage() {}

func (x *ReadLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadLogsRequest.ProtoReflect.Descriptor instead.
func (*ReadLogsRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_service_proto_rawDescGZIP(), []int{0}
}

func (m *ReadLogsRequest) GetSelector() isReadLogsRequest_Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (x *ReadLogsRequest) GetPaging() *v1.NextPage {
	if x, ok := x.GetSelector().(*ReadLogsRequest_Paging); ok {
		return x.Paging
	}
	return nil
}

func (x *ReadLogsRequest) GetCriteria() *Criteria {
	if x, ok := x.GetSelector().(*ReadLogsRequest_Criteria); ok {
		return x.Criteria
	}
	return nil
}

type isReadLogsRequest_Selector interface {
	isReadLogsRequest_Selector()
}

type ReadLogsRequest_Paging struct {
	// Paging information from previous calls, it contains Criteria info.
	Paging *v1.NextPage `protobuf:"bytes,1,opt,name=paging,proto3,oneof"`
}

type ReadLogsRequest_Criteria struct {
	// Log selection criteria.
	Criteria *Criteria `protobuf:"bytes,2,opt,name=criteria,proto3,oneof"`
}

func (*ReadLogsRequest_Paging) isReadLogsRequest_Selector() {}

func (*ReadLogsRequest_Criteria) isReadLogsRequest_Selector() {}

type ReadLogRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One of the requested log records.
	Record *LogRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	// Pagination information of the response, containing criteria.
	NextPage *v1.NextPage `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *ReadLogRecord) Reset() {
	*x = ReadLogRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadLogRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadLogRecord) ProtoMessage() {}

func (x *ReadLogRecord) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadLogRecord.ProtoReflect.Descriptor instead.
func (*ReadLogRecord) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReadLogRecord) GetRecord() *LogRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *ReadLogRecord) GetNextPage() *v1.NextPage {
	if x != nil {
		return x.NextPage
	}
	return nil
}

var File_doublecloud_logs_v1_log_service_proto protoreflect.FileDescriptor

var file_doublecloud_logs_v1_log_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x61,
	0x64, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x12, 0x3b, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x48, 0x00, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x42, 0x0a, 0x0a,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x7e, 0x0a, 0x0d, 0x52, 0x65, 0x61,
	0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x32, 0x60, 0x0a, 0x0a, 0x4c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x24, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_doublecloud_logs_v1_log_service_proto_rawDescOnce sync.Once
	file_doublecloud_logs_v1_log_service_proto_rawDescData = file_doublecloud_logs_v1_log_service_proto_rawDesc
)

func file_doublecloud_logs_v1_log_service_proto_rawDescGZIP() []byte {
	file_doublecloud_logs_v1_log_service_proto_rawDescOnce.Do(func() {
		file_doublecloud_logs_v1_log_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_logs_v1_log_service_proto_rawDescData)
	})
	return file_doublecloud_logs_v1_log_service_proto_rawDescData
}

var file_doublecloud_logs_v1_log_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_logs_v1_log_service_proto_goTypes = []interface{}{
	(*ReadLogsRequest)(nil), // 0: doublecloud.logs.v1.ReadLogsRequest
	(*ReadLogRecord)(nil),   // 1: doublecloud.logs.v1.ReadLogRecord
	(*v1.NextPage)(nil),     // 2: doublecloud.v1.NextPage
	(*Criteria)(nil),        // 3: doublecloud.logs.v1.Criteria
	(*LogRecord)(nil),       // 4: doublecloud.logs.v1.LogRecord
}
var file_doublecloud_logs_v1_log_service_proto_depIdxs = []int32{
	2, // 0: doublecloud.logs.v1.ReadLogsRequest.paging:type_name -> doublecloud.v1.NextPage
	3, // 1: doublecloud.logs.v1.ReadLogsRequest.criteria:type_name -> doublecloud.logs.v1.Criteria
	4, // 2: doublecloud.logs.v1.ReadLogRecord.record:type_name -> doublecloud.logs.v1.LogRecord
	2, // 3: doublecloud.logs.v1.ReadLogRecord.next_page:type_name -> doublecloud.v1.NextPage
	0, // 4: doublecloud.logs.v1.LogService.Read:input_type -> doublecloud.logs.v1.ReadLogsRequest
	1, // 5: doublecloud.logs.v1.LogService.Read:output_type -> doublecloud.logs.v1.ReadLogRecord
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_doublecloud_logs_v1_log_service_proto_init() }
func file_doublecloud_logs_v1_log_service_proto_init() {
	if File_doublecloud_logs_v1_log_service_proto != nil {
		return
	}
	file_doublecloud_logs_v1_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_logs_v1_log_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadLogsRequest); i {
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
		file_doublecloud_logs_v1_log_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadLogRecord); i {
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
	file_doublecloud_logs_v1_log_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ReadLogsRequest_Paging)(nil),
		(*ReadLogsRequest_Criteria)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_logs_v1_log_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doublecloud_logs_v1_log_service_proto_goTypes,
		DependencyIndexes: file_doublecloud_logs_v1_log_service_proto_depIdxs,
		MessageInfos:      file_doublecloud_logs_v1_log_service_proto_msgTypes,
	}.Build()
	File_doublecloud_logs_v1_log_service_proto = out.File
	file_doublecloud_logs_v1_log_service_proto_rawDesc = nil
	file_doublecloud_logs_v1_log_service_proto_goTypes = nil
	file_doublecloud_logs_v1_log_service_proto_depIdxs = nil
}
