// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: doublecloud/logs/v1/log.proto

package logs

import (
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SortOrder int32

const (
	// Not specified, so default order = SORT_ORDER_ASCENDING applied.
	SortOrder_SORT_ORDER_INVALID SortOrder = 0
	// Stream logs from oldest to newest.
	SortOrder_SORT_ORDER_ASCENDING SortOrder = 1
	// Stream logs staring from newest.
	SortOrder_SORT_ORDER_DESCENDING SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "SORT_ORDER_INVALID",
		1: "SORT_ORDER_ASCENDING",
		2: "SORT_ORDER_DESCENDING",
	}
	SortOrder_value = map[string]int32{
		"SORT_ORDER_INVALID":    0,
		"SORT_ORDER_ASCENDING":  1,
		"SORT_ORDER_DESCENDING": 2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_logs_v1_log_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_doublecloud_logs_v1_log_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_proto_rawDescGZIP(), []int{0}
}

type LogLevel int32

const (
	LogLevel_LOG_LEVEL_DEFAULT LogLevel = 0
	LogLevel_LOG_LEVEL_DEBUG   LogLevel = 2
	LogLevel_LOG_LEVEL_INFO    LogLevel = 3
	LogLevel_LOG_LEVEL_WARN    LogLevel = 4
	LogLevel_LOG_LEVEL_ERROR   LogLevel = 5
	LogLevel_LOG_LEVEL_FATAL   LogLevel = 6
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "LOG_LEVEL_DEFAULT",
		2: "LOG_LEVEL_DEBUG",
		3: "LOG_LEVEL_INFO",
		4: "LOG_LEVEL_WARN",
		5: "LOG_LEVEL_ERROR",
		6: "LOG_LEVEL_FATAL",
	}
	LogLevel_value = map[string]int32{
		"LOG_LEVEL_DEFAULT": 0,
		"LOG_LEVEL_DEBUG":   2,
		"LOG_LEVEL_INFO":    3,
		"LOG_LEVEL_WARN":    4,
		"LOG_LEVEL_ERROR":   5,
		"LOG_LEVEL_FATAL":   6,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_logs_v1_log_proto_enumTypes[1].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_doublecloud_logs_v1_log_proto_enumTypes[1]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_proto_rawDescGZIP(), []int{1}
}

type Criteria struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Log sources to read from.
	Sources []*LogSource `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	// Start timestamp for the logs request.
	FromTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	// End timestamp for the logs request.
	// If streaming and this field is not set, all existing logs will be sent and then
	// the new ones as
	// they appear. In essence it has 'tail -f' semantics.
	ToTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	// Filter on log level, empty list means no filter.
	Levels []LogLevel `protobuf:"varint,4,rep,packed,name=levels,proto3,enum=doublecloud.logs.v1.LogLevel" json:"levels,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The filtering language allows you to:
	//
	// - pinpoint searches by leveraging parameter values and comparison operators.
	// - sophisticated filters using logical operators (AND, OR, NOT).
	// Examples:
	// timestamp > "2023-06-08T00:00:00Z"
	// message = "warning" AND message <> "error dialing endpoint"
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// Order of the log messages in the stream. Ignored if to_time is not specified.
	SortOrder     SortOrder    `protobuf:"varint,6,opt,name=sort_order,json=sortOrder,proto3,enum=doublecloud.logs.v1.SortOrder" json:"sort_order,omitempty"`
	Paging        *v1.NextPage `protobuf:"bytes,7,opt,name=paging,proto3" json:"paging,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Criteria) Reset() {
	*x = Criteria{}
	mi := &file_doublecloud_logs_v1_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Criteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Criteria) ProtoMessage() {}

func (x *Criteria) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Criteria.ProtoReflect.Descriptor instead.
func (*Criteria) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *Criteria) GetSources() []*LogSource {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *Criteria) GetFromTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FromTime
	}
	return nil
}

func (x *Criteria) GetToTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ToTime
	}
	return nil
}

func (x *Criteria) GetLevels() []LogLevel {
	if x != nil {
		return x.Levels
	}
	return nil
}

func (x *Criteria) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *Criteria) GetSortOrder() SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return SortOrder_SORT_ORDER_INVALID
}

func (x *Criteria) GetPaging() *v1.NextPage {
	if x != nil {
		return x.Paging
	}
	return nil
}

type LogRecord struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Origin of the log message.
	Source *LogSource `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Optional field. Set if applicable for source type. For example host for
	// databases.
	Instance string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	// Timestamp of log string.
	RecordTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=record_time,json=recordTime,proto3" json:"record_time,omitempty"`
	// Severity of the message.
	Level LogLevel `protobuf:"varint,4,opt,name=level,proto3,enum=doublecloud.logs.v1.LogLevel" json:"level,omitempty"`
	// Log message
	Message       string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogRecord) Reset() {
	*x = LogRecord{}
	mi := &file_doublecloud_logs_v1_log_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRecord) ProtoMessage() {}

func (x *LogRecord) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRecord.ProtoReflect.Descriptor instead.
func (*LogRecord) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogRecord) GetSource() *LogSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *LogRecord) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *LogRecord) GetRecordTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordTime
	}
	return nil
}

func (x *LogRecord) GetLevel() LogLevel {
	if x != nil {
		return x.Level
	}
	return LogLevel_LOG_LEVEL_DEFAULT
}

func (x *LogRecord) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_doublecloud_logs_v1_log_proto protoreflect.FileDescriptor

var file_doublecloud_logs_v1_log_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x02, 0x0a, 0x08, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x38, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x37, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a,
	0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0xeb, 0x01,
	0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x58, 0x0a, 0x09, 0x53,
	0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x52, 0x54,
	0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41,
	0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x88, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f,
	0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x57, 0x41, 0x52, 0x4e, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4c,
	0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x06,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_logs_v1_log_proto_rawDescOnce sync.Once
	file_doublecloud_logs_v1_log_proto_rawDescData = file_doublecloud_logs_v1_log_proto_rawDesc
)

func file_doublecloud_logs_v1_log_proto_rawDescGZIP() []byte {
	file_doublecloud_logs_v1_log_proto_rawDescOnce.Do(func() {
		file_doublecloud_logs_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_logs_v1_log_proto_rawDescData)
	})
	return file_doublecloud_logs_v1_log_proto_rawDescData
}

var file_doublecloud_logs_v1_log_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_doublecloud_logs_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_logs_v1_log_proto_goTypes = []any{
	(SortOrder)(0),                // 0: doublecloud.logs.v1.SortOrder
	(LogLevel)(0),                 // 1: doublecloud.logs.v1.LogLevel
	(*Criteria)(nil),              // 2: doublecloud.logs.v1.Criteria
	(*LogRecord)(nil),             // 3: doublecloud.logs.v1.LogRecord
	(*LogSource)(nil),             // 4: doublecloud.logs.v1.LogSource
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*v1.NextPage)(nil),           // 6: doublecloud.v1.NextPage
}
var file_doublecloud_logs_v1_log_proto_depIdxs = []int32{
	4, // 0: doublecloud.logs.v1.Criteria.sources:type_name -> doublecloud.logs.v1.LogSource
	5, // 1: doublecloud.logs.v1.Criteria.from_time:type_name -> google.protobuf.Timestamp
	5, // 2: doublecloud.logs.v1.Criteria.to_time:type_name -> google.protobuf.Timestamp
	1, // 3: doublecloud.logs.v1.Criteria.levels:type_name -> doublecloud.logs.v1.LogLevel
	0, // 4: doublecloud.logs.v1.Criteria.sort_order:type_name -> doublecloud.logs.v1.SortOrder
	6, // 5: doublecloud.logs.v1.Criteria.paging:type_name -> doublecloud.v1.NextPage
	4, // 6: doublecloud.logs.v1.LogRecord.source:type_name -> doublecloud.logs.v1.LogSource
	5, // 7: doublecloud.logs.v1.LogRecord.record_time:type_name -> google.protobuf.Timestamp
	1, // 8: doublecloud.logs.v1.LogRecord.level:type_name -> doublecloud.logs.v1.LogLevel
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_doublecloud_logs_v1_log_proto_init() }
func file_doublecloud_logs_v1_log_proto_init() {
	if File_doublecloud_logs_v1_log_proto != nil {
		return
	}
	file_doublecloud_logs_v1_log_source_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_logs_v1_log_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_logs_v1_log_proto_goTypes,
		DependencyIndexes: file_doublecloud_logs_v1_log_proto_depIdxs,
		EnumInfos:         file_doublecloud_logs_v1_log_proto_enumTypes,
		MessageInfos:      file_doublecloud_logs_v1_log_proto_msgTypes,
	}.Build()
	File_doublecloud_logs_v1_log_proto = out.File
	file_doublecloud_logs_v1_log_proto_rawDesc = nil
	file_doublecloud_logs_v1_log_proto_goTypes = nil
	file_doublecloud_logs_v1_log_proto_depIdxs = nil
}
