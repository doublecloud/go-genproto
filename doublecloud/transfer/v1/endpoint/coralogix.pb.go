// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/coralogix.proto

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

type CoralogixSeverity int32

const (
	CoralogixSeverity_CORALOGIX_SEVERITY_UNSPECIFIED CoralogixSeverity = 0
	CoralogixSeverity_CORALOGIX_SEVERITY_DEBUG       CoralogixSeverity = 1
	CoralogixSeverity_CORALOGIX_SEVERITY_VERBOSE     CoralogixSeverity = 2
	CoralogixSeverity_CORALOGIX_SEVERITY_INFO        CoralogixSeverity = 3
	CoralogixSeverity_CORALOGIX_SEVERITY_WARN        CoralogixSeverity = 4
	CoralogixSeverity_CORALOGIX_SEVERITY_ERROR       CoralogixSeverity = 5
	CoralogixSeverity_CORALOGIX_SEVERITY_CRITICAL    CoralogixSeverity = 6
)

// Enum value maps for CoralogixSeverity.
var (
	CoralogixSeverity_name = map[int32]string{
		0: "CORALOGIX_SEVERITY_UNSPECIFIED",
		1: "CORALOGIX_SEVERITY_DEBUG",
		2: "CORALOGIX_SEVERITY_VERBOSE",
		3: "CORALOGIX_SEVERITY_INFO",
		4: "CORALOGIX_SEVERITY_WARN",
		5: "CORALOGIX_SEVERITY_ERROR",
		6: "CORALOGIX_SEVERITY_CRITICAL",
	}
	CoralogixSeverity_value = map[string]int32{
		"CORALOGIX_SEVERITY_UNSPECIFIED": 0,
		"CORALOGIX_SEVERITY_DEBUG":       1,
		"CORALOGIX_SEVERITY_VERBOSE":     2,
		"CORALOGIX_SEVERITY_INFO":        3,
		"CORALOGIX_SEVERITY_WARN":        4,
		"CORALOGIX_SEVERITY_ERROR":       5,
		"CORALOGIX_SEVERITY_CRITICAL":    6,
	}
)

func (x CoralogixSeverity) Enum() *CoralogixSeverity {
	p := new(CoralogixSeverity)
	*p = x
	return p
}

func (x CoralogixSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoralogixSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_coralogix_proto_enumTypes[0].Descriptor()
}

func (CoralogixSeverity) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_coralogix_proto_enumTypes[0]
}

func (x CoralogixSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoralogixSeverity.Descriptor instead.
func (CoralogixSeverity) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescGZIP(), []int{0}
}

type CoralogixTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token           string                         `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Domain          string                         `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	ApplicationName string                         `protobuf:"bytes,3,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	ColumnMapping   *CoralogixTarget_ColumnMapping `protobuf:"bytes,5,opt,name=column_mapping,json=columnMapping,proto3" json:"column_mapping,omitempty"`
}

func (x *CoralogixTarget) Reset() {
	*x = CoralogixTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoralogixTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoralogixTarget) ProtoMessage() {}

func (x *CoralogixTarget) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoralogixTarget.ProtoReflect.Descriptor instead.
func (*CoralogixTarget) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescGZIP(), []int{0}
}

func (x *CoralogixTarget) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CoralogixTarget) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CoralogixTarget) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *CoralogixTarget) GetColumnMapping() *CoralogixTarget_ColumnMapping {
	if x != nil {
		return x.ColumnMapping
	}
	return nil
}

type CoralogixTarget_SeverityMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value CoralogixSeverity `protobuf:"varint,2,opt,name=value,proto3,enum=doublecloud.transfer.v1.endpoint.CoralogixSeverity" json:"value,omitempty"`
}

func (x *CoralogixTarget_SeverityMap) Reset() {
	*x = CoralogixTarget_SeverityMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoralogixTarget_SeverityMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoralogixTarget_SeverityMap) ProtoMessage() {}

func (x *CoralogixTarget_SeverityMap) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoralogixTarget_SeverityMap.ProtoReflect.Descriptor instead.
func (*CoralogixTarget_SeverityMap) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CoralogixTarget_SeverityMap) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CoralogixTarget_SeverityMap) GetValue() CoralogixSeverity {
	if x != nil {
		return x.Value
	}
	return CoralogixSeverity_CORALOGIX_SEVERITY_UNSPECIFIED
}

type CoralogixTarget_ColumnMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp       string                         `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Severity        string                         `protobuf:"bytes,2,opt,name=severity,proto3" json:"severity,omitempty"`
	MessageTemplate string                         `protobuf:"bytes,3,opt,name=message_template,json=messageTemplate,proto3" json:"message_template,omitempty"`
	ClassName       string                         `protobuf:"bytes,4,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	MethodName      string                         `protobuf:"bytes,5,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	ThreadId        string                         `protobuf:"bytes,6,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	Category        string                         `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	Subsystem       string                         `protobuf:"bytes,8,opt,name=subsystem,proto3" json:"subsystem,omitempty"`
	Host            string                         `protobuf:"bytes,9,opt,name=host,proto3" json:"host,omitempty"`
	KnownSeverities []*CoralogixTarget_SeverityMap `protobuf:"bytes,10,rep,name=known_severities,json=knownSeverities,proto3" json:"known_severities,omitempty"`
}

func (x *CoralogixTarget_ColumnMapping) Reset() {
	*x = CoralogixTarget_ColumnMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoralogixTarget_ColumnMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoralogixTarget_ColumnMapping) ProtoMessage() {}

func (x *CoralogixTarget_ColumnMapping) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoralogixTarget_ColumnMapping.ProtoReflect.Descriptor instead.
func (*CoralogixTarget_ColumnMapping) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CoralogixTarget_ColumnMapping) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetMessageTemplate() string {
	if x != nil {
		return x.MessageTemplate
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetSubsystem() string {
	if x != nil {
		return x.Subsystem
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *CoralogixTarget_ColumnMapping) GetKnownSeverities() []*CoralogixTarget_SeverityMap {
	if x != nil {
		return x.KnownSeverities
	}
	return nil
}

var File_doublecloud_transfer_v1_endpoint_coralogix_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDesc = []byte{
	0x0a, 0x30, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0xcc, 0x05, 0x0a, 0x0f, 0x43, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x66, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x6c, 0x0a, 0x0b, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x89, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x68, 0x0a, 0x10, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x4d,
	0x61, 0x70, 0x52, 0x0f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2a, 0xee, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x4f, 0x52,
	0x41, 0x4c, 0x4f, 0x47, 0x49, 0x58, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x43, 0x4f, 0x52, 0x41, 0x4c, 0x4f, 0x47, 0x49, 0x58, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x43,
	0x4f, 0x52, 0x41, 0x4c, 0x4f, 0x47, 0x49, 0x58, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x43,
	0x4f, 0x52, 0x41, 0x4c, 0x4f, 0x47, 0x49, 0x58, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x52, 0x41,
	0x4c, 0x4f, 0x47, 0x49, 0x58, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x57,
	0x41, 0x52, 0x4e, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x52, 0x41, 0x4c, 0x4f, 0x47,
	0x49, 0x58, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x52, 0x41, 0x4c, 0x4f, 0x47, 0x49, 0x58,
	0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43,
	0x41, 0x4c, 0x10, 0x06, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_coralogix_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_doublecloud_transfer_v1_endpoint_coralogix_proto_goTypes = []interface{}{
	(CoralogixSeverity)(0),                // 0: doublecloud.transfer.v1.endpoint.CoralogixSeverity
	(*CoralogixTarget)(nil),               // 1: doublecloud.transfer.v1.endpoint.CoralogixTarget
	(*CoralogixTarget_SeverityMap)(nil),   // 2: doublecloud.transfer.v1.endpoint.CoralogixTarget.SeverityMap
	(*CoralogixTarget_ColumnMapping)(nil), // 3: doublecloud.transfer.v1.endpoint.CoralogixTarget.ColumnMapping
}
var file_doublecloud_transfer_v1_endpoint_coralogix_proto_depIdxs = []int32{
	3, // 0: doublecloud.transfer.v1.endpoint.CoralogixTarget.column_mapping:type_name -> doublecloud.transfer.v1.endpoint.CoralogixTarget.ColumnMapping
	0, // 1: doublecloud.transfer.v1.endpoint.CoralogixTarget.SeverityMap.value:type_name -> doublecloud.transfer.v1.endpoint.CoralogixSeverity
	2, // 2: doublecloud.transfer.v1.endpoint.CoralogixTarget.ColumnMapping.known_severities:type_name -> doublecloud.transfer.v1.endpoint.CoralogixTarget.SeverityMap
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_coralogix_proto_init() }
func file_doublecloud_transfer_v1_endpoint_coralogix_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_coralogix_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoralogixTarget); i {
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
		file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoralogixTarget_SeverityMap); i {
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
		file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoralogixTarget_ColumnMapping); i {
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
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_coralogix_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_coralogix_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_endpoint_coralogix_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_coralogix_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_coralogix_proto = out.File
	file_doublecloud_transfer_v1_endpoint_coralogix_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_coralogix_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_coralogix_proto_depIdxs = nil
}
