// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: doublecloud/v1/maintenance.proto

package doublecloud

import (
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
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

type RescheduleType int32

const (
	RescheduleType_RESCHEDULE_TYPE_INVALID               RescheduleType = 0
	RescheduleType_RESCHEDULE_TYPE_IMMEDIATE             RescheduleType = 1
	RescheduleType_RESCHEDULE_TYPE_NEXT_AVAILABLE_WINDOW RescheduleType = 2
	RescheduleType_RESCHEDULE_TYPE_SPECIFIC_TIME         RescheduleType = 3
)

// Enum value maps for RescheduleType.
var (
	RescheduleType_name = map[int32]string{
		0: "RESCHEDULE_TYPE_INVALID",
		1: "RESCHEDULE_TYPE_IMMEDIATE",
		2: "RESCHEDULE_TYPE_NEXT_AVAILABLE_WINDOW",
		3: "RESCHEDULE_TYPE_SPECIFIC_TIME",
	}
	RescheduleType_value = map[string]int32{
		"RESCHEDULE_TYPE_INVALID":               0,
		"RESCHEDULE_TYPE_IMMEDIATE":             1,
		"RESCHEDULE_TYPE_NEXT_AVAILABLE_WINDOW": 2,
		"RESCHEDULE_TYPE_SPECIFIC_TIME":         3,
	}
)

func (x RescheduleType) Enum() *RescheduleType {
	p := new(RescheduleType)
	*p = x
	return p
}

func (x RescheduleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RescheduleType) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_v1_maintenance_proto_enumTypes[0].Descriptor()
}

func (RescheduleType) Type() protoreflect.EnumType {
	return &file_doublecloud_v1_maintenance_proto_enumTypes[0]
}

func (x RescheduleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RescheduleType.Descriptor instead.
func (RescheduleType) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_v1_maintenance_proto_rawDescGZIP(), []int{0}
}

type MaintenanceWindow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Policy:
	//
	//	*MaintenanceWindow_Anytime
	//	*MaintenanceWindow_WeeklyMaintenanceWindow
	Policy        isMaintenanceWindow_Policy `protobuf_oneof:"policy"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MaintenanceWindow) Reset() {
	*x = MaintenanceWindow{}
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceWindow) ProtoMessage() {}

func (x *MaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceWindow.ProtoReflect.Descriptor instead.
func (*MaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_maintenance_proto_rawDescGZIP(), []int{0}
}

func (x *MaintenanceWindow) GetPolicy() isMaintenanceWindow_Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *MaintenanceWindow) GetAnytime() *AnytimeMaintenanceWindow {
	if x != nil {
		if x, ok := x.Policy.(*MaintenanceWindow_Anytime); ok {
			return x.Anytime
		}
	}
	return nil
}

func (x *MaintenanceWindow) GetWeeklyMaintenanceWindow() *WeeklyMaintenanceWindow {
	if x != nil {
		if x, ok := x.Policy.(*MaintenanceWindow_WeeklyMaintenanceWindow); ok {
			return x.WeeklyMaintenanceWindow
		}
	}
	return nil
}

type isMaintenanceWindow_Policy interface {
	isMaintenanceWindow_Policy()
}

type MaintenanceWindow_Anytime struct {
	Anytime *AnytimeMaintenanceWindow `protobuf:"bytes,1,opt,name=anytime,proto3,oneof"`
}

type MaintenanceWindow_WeeklyMaintenanceWindow struct {
	WeeklyMaintenanceWindow *WeeklyMaintenanceWindow `protobuf:"bytes,2,opt,name=weekly_maintenance_window,json=weeklyMaintenanceWindow,proto3,oneof"`
}

func (*MaintenanceWindow_Anytime) isMaintenanceWindow_Policy() {}

func (*MaintenanceWindow_WeeklyMaintenanceWindow) isMaintenanceWindow_Policy() {}

type AnytimeMaintenanceWindow struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnytimeMaintenanceWindow) Reset() {
	*x = AnytimeMaintenanceWindow{}
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnytimeMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnytimeMaintenanceWindow) ProtoMessage() {}

func (x *AnytimeMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnytimeMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*AnytimeMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_maintenance_proto_rawDescGZIP(), []int{1}
}

type WeeklyMaintenanceWindow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Day   dayofweek.DayOfWeek    `protobuf:"varint,1,opt,name=day,proto3,enum=google.type.DayOfWeek" json:"day,omitempty"`
	// Hour of the day in UTC (1 - 24).
	Hour          int64 `protobuf:"varint,2,opt,name=hour,proto3" json:"hour,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WeeklyMaintenanceWindow) Reset() {
	*x = WeeklyMaintenanceWindow{}
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeeklyMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklyMaintenanceWindow) ProtoMessage() {}

func (x *WeeklyMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklyMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*WeeklyMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_maintenance_proto_rawDescGZIP(), []int{2}
}

func (x *WeeklyMaintenanceWindow) GetDay() dayofweek.DayOfWeek {
	if x != nil {
		return x.Day
	}
	return dayofweek.DayOfWeek(0)
}

func (x *WeeklyMaintenanceWindow) GetHour() int64 {
	if x != nil {
		return x.Hour
	}
	return 0
}

type MaintenanceOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Length of info string should be limited by 256 characters
	Info                      string                 `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	ScheduledMaintenanceTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=scheduled_maintenance_time,json=scheduledMaintenanceTime,proto3" json:"scheduled_maintenance_time,omitempty"`
	DeadlineMaintenanceTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=deadline_maintenance_time,json=deadlineMaintenanceTime,proto3" json:"deadline_maintenance_time,omitempty"`
	NextMaintenanceWindowTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=next_maintenance_window_time,json=nextMaintenanceWindowTime,proto3" json:"next_maintenance_window_time,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *MaintenanceOperation) Reset() {
	*x = MaintenanceOperation{}
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaintenanceOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceOperation) ProtoMessage() {}

func (x *MaintenanceOperation) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_v1_maintenance_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceOperation.ProtoReflect.Descriptor instead.
func (*MaintenanceOperation) Descriptor() ([]byte, []int) {
	return file_doublecloud_v1_maintenance_proto_rawDescGZIP(), []int{3}
}

func (x *MaintenanceOperation) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *MaintenanceOperation) GetScheduledMaintenanceTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledMaintenanceTime
	}
	return nil
}

func (x *MaintenanceOperation) GetDeadlineMaintenanceTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeadlineMaintenanceTime
	}
	return nil
}

func (x *MaintenanceOperation) GetNextMaintenanceWindowTime() *timestamppb.Timestamp {
	if x != nil {
		return x.NextMaintenanceWindowTime
	}
	return nil
}

var File_doublecloud_v1_maintenance_proto protoreflect.FileDescriptor

var file_doublecloud_v1_maintenance_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x64, 0x61, 0x79, 0x6f, 0x66, 0x77, 0x65, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xca, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x44, 0x0a, 0x07, 0x61, 0x6e, 0x79, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65,
	0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x48, 0x00, 0x52, 0x07, 0x61, 0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x65, 0x0a, 0x19,
	0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x48, 0x00, 0x52, 0x17, 0x77, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x1a, 0x0a,
	0x18, 0x41, 0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x57, 0x0a, 0x17, 0x57, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x12, 0x28, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x68, 0x6f,
	0x75, 0x72, 0x22, 0xb9, 0x02, 0x0a, 0x14, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x58, 0x0a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x18, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x19, 0x64, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x17, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x5b, 0x0a, 0x1c, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x19, 0x6e, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x9a,
	0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1d,
	0x0a, 0x19, 0x52, 0x45, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x29, 0x0a,
	0x25, 0x52, 0x45, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x53, 0x43,
	0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x3b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_v1_maintenance_proto_rawDescOnce sync.Once
	file_doublecloud_v1_maintenance_proto_rawDescData = file_doublecloud_v1_maintenance_proto_rawDesc
)

func file_doublecloud_v1_maintenance_proto_rawDescGZIP() []byte {
	file_doublecloud_v1_maintenance_proto_rawDescOnce.Do(func() {
		file_doublecloud_v1_maintenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_v1_maintenance_proto_rawDescData)
	})
	return file_doublecloud_v1_maintenance_proto_rawDescData
}

var file_doublecloud_v1_maintenance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_v1_maintenance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_doublecloud_v1_maintenance_proto_goTypes = []any{
	(RescheduleType)(0),              // 0: doublecloud.v1.RescheduleType
	(*MaintenanceWindow)(nil),        // 1: doublecloud.v1.MaintenanceWindow
	(*AnytimeMaintenanceWindow)(nil), // 2: doublecloud.v1.AnytimeMaintenanceWindow
	(*WeeklyMaintenanceWindow)(nil),  // 3: doublecloud.v1.WeeklyMaintenanceWindow
	(*MaintenanceOperation)(nil),     // 4: doublecloud.v1.MaintenanceOperation
	(dayofweek.DayOfWeek)(0),         // 5: google.type.DayOfWeek
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_doublecloud_v1_maintenance_proto_depIdxs = []int32{
	2, // 0: doublecloud.v1.MaintenanceWindow.anytime:type_name -> doublecloud.v1.AnytimeMaintenanceWindow
	3, // 1: doublecloud.v1.MaintenanceWindow.weekly_maintenance_window:type_name -> doublecloud.v1.WeeklyMaintenanceWindow
	5, // 2: doublecloud.v1.WeeklyMaintenanceWindow.day:type_name -> google.type.DayOfWeek
	6, // 3: doublecloud.v1.MaintenanceOperation.scheduled_maintenance_time:type_name -> google.protobuf.Timestamp
	6, // 4: doublecloud.v1.MaintenanceOperation.deadline_maintenance_time:type_name -> google.protobuf.Timestamp
	6, // 5: doublecloud.v1.MaintenanceOperation.next_maintenance_window_time:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_doublecloud_v1_maintenance_proto_init() }
func file_doublecloud_v1_maintenance_proto_init() {
	if File_doublecloud_v1_maintenance_proto != nil {
		return
	}
	file_doublecloud_v1_maintenance_proto_msgTypes[0].OneofWrappers = []any{
		(*MaintenanceWindow_Anytime)(nil),
		(*MaintenanceWindow_WeeklyMaintenanceWindow)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_v1_maintenance_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_v1_maintenance_proto_goTypes,
		DependencyIndexes: file_doublecloud_v1_maintenance_proto_depIdxs,
		EnumInfos:         file_doublecloud_v1_maintenance_proto_enumTypes,
		MessageInfos:      file_doublecloud_v1_maintenance_proto_msgTypes,
	}.Build()
	File_doublecloud_v1_maintenance_proto = out.File
	file_doublecloud_v1_maintenance_proto_rawDesc = nil
	file_doublecloud_v1_maintenance_proto_goTypes = nil
	file_doublecloud_v1_maintenance_proto_depIdxs = nil
}
