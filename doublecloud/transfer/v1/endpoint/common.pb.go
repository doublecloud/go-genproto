// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/common.proto

package endpoint

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ObjectTransferStage int32

const (
	ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED ObjectTransferStage = 0
	// Before data transfer
	ObjectTransferStage_BEFORE_DATA ObjectTransferStage = 1
	// After data transfer
	ObjectTransferStage_AFTER_DATA ObjectTransferStage = 2
	// Don't copy
	ObjectTransferStage_NEVER ObjectTransferStage = 3
)

// Enum value maps for ObjectTransferStage.
var (
	ObjectTransferStage_name = map[int32]string{
		0: "OBJECT_TRANSFER_STAGE_UNSPECIFIED",
		1: "BEFORE_DATA",
		2: "AFTER_DATA",
		3: "NEVER",
	}
	ObjectTransferStage_value = map[string]int32{
		"OBJECT_TRANSFER_STAGE_UNSPECIFIED": 0,
		"BEFORE_DATA":                       1,
		"AFTER_DATA":                        2,
		"NEVER":                             3,
	}
)

func (x ObjectTransferStage) Enum() *ObjectTransferStage {
	p := new(ObjectTransferStage)
	*p = x
	return p
}

func (x ObjectTransferStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectTransferStage) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes[0].Descriptor()
}

func (ObjectTransferStage) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes[0]
}

func (x ObjectTransferStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectTransferStage.Descriptor instead.
func (ObjectTransferStage) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{0}
}

type CleanupPolicy int32

const (
	CleanupPolicy_CLEANUP_POLICY_UNSPECIFIED CleanupPolicy = 0
	// Don't cleanup
	CleanupPolicy_DISABLED CleanupPolicy = 1
	// Drop
	CleanupPolicy_DROP CleanupPolicy = 2
	// Truncate
	CleanupPolicy_TRUNCATE CleanupPolicy = 3
)

// Enum value maps for CleanupPolicy.
var (
	CleanupPolicy_name = map[int32]string{
		0: "CLEANUP_POLICY_UNSPECIFIED",
		1: "DISABLED",
		2: "DROP",
		3: "TRUNCATE",
	}
	CleanupPolicy_value = map[string]int32{
		"CLEANUP_POLICY_UNSPECIFIED": 0,
		"DISABLED":                   1,
		"DROP":                       2,
		"TRUNCATE":                   3,
	}
)

func (x CleanupPolicy) Enum() *CleanupPolicy {
	p := new(CleanupPolicy)
	*p = x
	return p
}

func (x CleanupPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CleanupPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes[1].Descriptor()
}

func (CleanupPolicy) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes[1]
}

func (x CleanupPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CleanupPolicy.Descriptor instead.
func (CleanupPolicy) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{1}
}

type ColumnType int32

const (
	ColumnType_COLUMN_TYPE_UNSPECIFIED ColumnType = 0
	ColumnType_INT32                   ColumnType = 1
	ColumnType_INT16                   ColumnType = 2
	ColumnType_INT8                    ColumnType = 3
	ColumnType_UINT64                  ColumnType = 4
	ColumnType_UINT32                  ColumnType = 5
	ColumnType_UINT16                  ColumnType = 6
	ColumnType_UINT8                   ColumnType = 7
	ColumnType_DOUBLE                  ColumnType = 8
	ColumnType_BOOLEAN                 ColumnType = 9
	ColumnType_STRING                  ColumnType = 10
	ColumnType_UTF8                    ColumnType = 11
	ColumnType_ANY                     ColumnType = 12
	ColumnType_DATETIME                ColumnType = 13
	ColumnType_INT64                   ColumnType = 14
)

// Enum value maps for ColumnType.
var (
	ColumnType_name = map[int32]string{
		0:  "COLUMN_TYPE_UNSPECIFIED",
		1:  "INT32",
		2:  "INT16",
		3:  "INT8",
		4:  "UINT64",
		5:  "UINT32",
		6:  "UINT16",
		7:  "UINT8",
		8:  "DOUBLE",
		9:  "BOOLEAN",
		10: "STRING",
		11: "UTF8",
		12: "ANY",
		13: "DATETIME",
		14: "INT64",
	}
	ColumnType_value = map[string]int32{
		"COLUMN_TYPE_UNSPECIFIED": 0,
		"INT32":                   1,
		"INT16":                   2,
		"INT8":                    3,
		"UINT64":                  4,
		"UINT32":                  5,
		"UINT16":                  6,
		"UINT8":                   7,
		"DOUBLE":                  8,
		"BOOLEAN":                 9,
		"STRING":                  10,
		"UTF8":                    11,
		"ANY":                     12,
		"DATETIME":                13,
		"INT64":                   14,
	}
)

func (x ColumnType) Enum() *ColumnType {
	p := new(ColumnType)
	*p = x
	return p
}

func (x ColumnType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColumnType) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes[2].Descriptor()
}

func (ColumnType) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes[2]
}

func (x ColumnType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColumnType.Descriptor instead.
func (ColumnType) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{2}
}

type AltName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source table name
	FromName string `protobuf:"bytes,1,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	// Target table name
	ToName string `protobuf:"bytes,2,opt,name=to_name,json=toName,proto3" json:"to_name,omitempty"`
}

func (x *AltName) Reset() {
	*x = AltName{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AltName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AltName) ProtoMessage() {}

func (x *AltName) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AltName.ProtoReflect.Descriptor instead.
func (*AltName) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{0}
}

func (x *AltName) GetFromName() string {
	if x != nil {
		return x.FromName
	}
	return ""
}

func (x *AltName) GetToName() string {
	if x != nil {
		return x.ToName
	}
	return ""
}

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Secret_Raw
	Value isSecret_Value `protobuf_oneof:"value"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[1]
	if x != nil {
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
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{1}
}

func (m *Secret) GetValue() isSecret_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Secret) GetRaw() string {
	if x, ok := x.GetValue().(*Secret_Raw); ok {
		return x.Raw
	}
	return ""
}

type isSecret_Value interface {
	isSecret_Value()
}

type Secret_Raw struct {
	// Raw secret value
	Raw string `protobuf:"bytes,1,opt,name=raw,proto3,oneof"`
}

func (*Secret_Raw) isSecret_Value() {}

type ColSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     ColumnType `protobuf:"varint,2,opt,name=type,proto3,enum=doublecloud.transfer.v1.endpoint.ColumnType" json:"type,omitempty"`
	Key      bool       `protobuf:"varint,3,opt,name=key,proto3" json:"key,omitempty"`
	Required bool       `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	Path     string     `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ColSchema) Reset() {
	*x = ColSchema{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColSchema) ProtoMessage() {}

func (x *ColSchema) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColSchema.ProtoReflect.Descriptor instead.
func (*ColSchema) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{2}
}

func (x *ColSchema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColSchema) GetType() ColumnType {
	if x != nil {
		return x.Type
	}
	return ColumnType_COLUMN_TYPE_UNSPECIFIED
}

func (x *ColSchema) GetKey() bool {
	if x != nil {
		return x.Key
	}
	return false
}

func (x *ColSchema) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *ColSchema) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type TLSMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TlsMode:
	//
	//	*TLSMode_Disabled
	//	*TLSMode_Enabled
	TlsMode isTLSMode_TlsMode `protobuf_oneof:"tls_mode"`
}

func (x *TLSMode) Reset() {
	*x = TLSMode{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLSMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSMode) ProtoMessage() {}

func (x *TLSMode) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSMode.ProtoReflect.Descriptor instead.
func (*TLSMode) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{3}
}

func (m *TLSMode) GetTlsMode() isTLSMode_TlsMode {
	if m != nil {
		return m.TlsMode
	}
	return nil
}

func (x *TLSMode) GetDisabled() *emptypb.Empty {
	if x, ok := x.GetTlsMode().(*TLSMode_Disabled); ok {
		return x.Disabled
	}
	return nil
}

func (x *TLSMode) GetEnabled() *TLSConfig {
	if x, ok := x.GetTlsMode().(*TLSMode_Enabled); ok {
		return x.Enabled
	}
	return nil
}

type isTLSMode_TlsMode interface {
	isTLSMode_TlsMode()
}

type TLSMode_Disabled struct {
	Disabled *emptypb.Empty `protobuf:"bytes,1,opt,name=disabled,proto3,oneof"`
}

type TLSMode_Enabled struct {
	Enabled *TLSConfig `protobuf:"bytes,2,opt,name=enabled,proto3,oneof"`
}

func (*TLSMode_Disabled) isTLSMode_TlsMode() {}

func (*TLSMode_Enabled) isTLSMode_TlsMode() {}

type TLSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CA certificate
	//
	// X.509 certificate of the certificate authority which issued the server's
	// certificate, in PEM format. When CA certificate is specified TLS is used to
	// connect to the server.
	CaCertificate string `protobuf:"bytes,1,opt,name=ca_certificate,json=caCertificate,proto3" json:"ca_certificate,omitempty"`
}

func (x *TLSConfig) Reset() {
	*x = TLSConfig{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSConfig) ProtoMessage() {}

func (x *TLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSConfig.ProtoReflect.Descriptor instead.
func (*TLSConfig) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{4}
}

func (x *TLSConfig) GetCaCertificate() string {
	if x != nil {
		return x.CaCertificate
	}
	return ""
}

type ColumnValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*ColumnValue_StringValue
	Value isColumnValue_Value `protobuf_oneof:"value"`
}

func (x *ColumnValue) Reset() {
	*x = ColumnValue{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColumnValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnValue) ProtoMessage() {}

func (x *ColumnValue) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnValue.ProtoReflect.Descriptor instead.
func (*ColumnValue) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{5}
}

func (m *ColumnValue) GetValue() isColumnValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ColumnValue) GetStringValue() string {
	if x, ok := x.GetValue().(*ColumnValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

type isColumnValue_Value interface {
	isColumnValue_Value()
}

type ColumnValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*ColumnValue_StringValue) isColumnValue_Value() {}

type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{6}
}

func (x *HeaderValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HeaderValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DataTransformationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cloud function
	CloudFunction string `protobuf:"bytes,1,opt,name=cloud_function,json=cloudFunction,proto3" json:"cloud_function,omitempty"`
	// Number of retries
	NumberOfRetries int64 `protobuf:"varint,2,opt,name=number_of_retries,json=numberOfRetries,proto3" json:"number_of_retries,omitempty"`
	// Buffer size for function
	BufferSize string `protobuf:"bytes,3,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	// Flush interval
	BufferFlushInterval string `protobuf:"bytes,4,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	// Invocation timeout
	InvocationTimeout string         `protobuf:"bytes,5,opt,name=invocation_timeout,json=invocationTimeout,proto3" json:"invocation_timeout,omitempty"`
	CloudFunctionUrl  string         `protobuf:"bytes,9,opt,name=cloud_function_url,json=cloudFunctionUrl,proto3" json:"cloud_function_url,omitempty"`
	Headers           []*HeaderValue `protobuf:"bytes,10,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *DataTransformationOptions) Reset() {
	*x = DataTransformationOptions{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataTransformationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTransformationOptions) ProtoMessage() {}

func (x *DataTransformationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTransformationOptions.ProtoReflect.Descriptor instead.
func (*DataTransformationOptions) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{7}
}

func (x *DataTransformationOptions) GetCloudFunction() string {
	if x != nil {
		return x.CloudFunction
	}
	return ""
}

func (x *DataTransformationOptions) GetNumberOfRetries() int64 {
	if x != nil {
		return x.NumberOfRetries
	}
	return 0
}

func (x *DataTransformationOptions) GetBufferSize() string {
	if x != nil {
		return x.BufferSize
	}
	return ""
}

func (x *DataTransformationOptions) GetBufferFlushInterval() string {
	if x != nil {
		return x.BufferFlushInterval
	}
	return ""
}

func (x *DataTransformationOptions) GetInvocationTimeout() string {
	if x != nil {
		return x.InvocationTimeout
	}
	return ""
}

func (x *DataTransformationOptions) GetCloudFunctionUrl() string {
	if x != nil {
		return x.CloudFunctionUrl
	}
	return ""
}

func (x *DataTransformationOptions) GetHeaders() []*HeaderValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

type FieldList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Column schema
	Fields []*ColSchema `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *FieldList) Reset() {
	*x = FieldList{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldList) ProtoMessage() {}

func (x *FieldList) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldList.ProtoReflect.Descriptor instead.
func (*FieldList) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{8}
}

func (x *FieldList) GetFields() []*ColSchema {
	if x != nil {
		return x.Fields
	}
	return nil
}

type DataSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Schema:
	//
	//	*DataSchema_JsonFields
	//	*DataSchema_Fields
	Schema isDataSchema_Schema `protobuf_oneof:"schema"`
}

func (x *DataSchema) Reset() {
	*x = DataSchema{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSchema) ProtoMessage() {}

func (x *DataSchema) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSchema.ProtoReflect.Descriptor instead.
func (*DataSchema) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{9}
}

func (m *DataSchema) GetSchema() isDataSchema_Schema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (x *DataSchema) GetJsonFields() string {
	if x, ok := x.GetSchema().(*DataSchema_JsonFields); ok {
		return x.JsonFields
	}
	return ""
}

func (x *DataSchema) GetFields() *FieldList {
	if x, ok := x.GetSchema().(*DataSchema_Fields); ok {
		return x.Fields
	}
	return nil
}

type isDataSchema_Schema interface {
	isDataSchema_Schema()
}

type DataSchema_JsonFields struct {
	JsonFields string `protobuf:"bytes,1,opt,name=json_fields,json=jsonFields,proto3,oneof"`
}

type DataSchema_Fields struct {
	Fields *FieldList `protobuf:"bytes,2,opt,name=fields,proto3,oneof"`
}

func (*DataSchema_JsonFields) isDataSchema_Schema() {}

func (*DataSchema_Fields) isDataSchema_Schema() {}

// No authentication
type NoAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoAuth) Reset() {
	*x = NoAuth{}
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoAuth) ProtoMessage() {}

func (x *NoAuth) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoAuth.ProtoReflect.Descriptor instead.
func (*NoAuth) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP(), []int{10}
}

var File_doublecloud_transfer_v1_endpoint_common_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_common_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f,
	0x0a, 0x07, 0x41, 0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x25, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x42, 0x07, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x94, 0x01, 0x0a,
	0x07, 0x54, 0x4c, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x47,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x74, 0x6c, 0x73, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x22, 0x32, 0x0a, 0x09, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xef, 0x02, 0x0a, 0x19,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a,
	0x15, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x47,
	0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x09, 0x22, 0x56, 0x0a,
	0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0x80, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x21, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6a, 0x73, 0x6f,
	0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x45, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x42, 0x08,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x08, 0x0a, 0x06, 0x4e, 0x6f, 0x41, 0x75,
	0x74, 0x68, 0x2a, 0x68, 0x0a, 0x13, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x4f, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x45, 0x46, 0x4f, 0x52, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x55, 0x0a, 0x0d,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x0a,
	0x1a, 0x43, 0x4c, 0x45, 0x41, 0x4e, 0x55, 0x50, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x52, 0x4f, 0x50, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x55, 0x4e, 0x43, 0x41, 0x54,
	0x45, 0x10, 0x03, 0x2a, 0xc9, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e,
	0x54, 0x31, 0x36, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x31,
	0x36, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x07, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f,
	0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x54, 0x46, 0x38, 0x10, 0x0b, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x4e, 0x59, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49,
	0x4d, 0x45, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x0e, 0x42,
	0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_common_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_common_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_common_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_common_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_common_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_common_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_common_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_doublecloud_transfer_v1_endpoint_common_proto_goTypes = []any{
	(ObjectTransferStage)(0),          // 0: doublecloud.transfer.v1.endpoint.ObjectTransferStage
	(CleanupPolicy)(0),                // 1: doublecloud.transfer.v1.endpoint.CleanupPolicy
	(ColumnType)(0),                   // 2: doublecloud.transfer.v1.endpoint.ColumnType
	(*AltName)(nil),                   // 3: doublecloud.transfer.v1.endpoint.AltName
	(*Secret)(nil),                    // 4: doublecloud.transfer.v1.endpoint.Secret
	(*ColSchema)(nil),                 // 5: doublecloud.transfer.v1.endpoint.ColSchema
	(*TLSMode)(nil),                   // 6: doublecloud.transfer.v1.endpoint.TLSMode
	(*TLSConfig)(nil),                 // 7: doublecloud.transfer.v1.endpoint.TLSConfig
	(*ColumnValue)(nil),               // 8: doublecloud.transfer.v1.endpoint.ColumnValue
	(*HeaderValue)(nil),               // 9: doublecloud.transfer.v1.endpoint.HeaderValue
	(*DataTransformationOptions)(nil), // 10: doublecloud.transfer.v1.endpoint.DataTransformationOptions
	(*FieldList)(nil),                 // 11: doublecloud.transfer.v1.endpoint.FieldList
	(*DataSchema)(nil),                // 12: doublecloud.transfer.v1.endpoint.DataSchema
	(*NoAuth)(nil),                    // 13: doublecloud.transfer.v1.endpoint.NoAuth
	(*emptypb.Empty)(nil),             // 14: google.protobuf.Empty
}
var file_doublecloud_transfer_v1_endpoint_common_proto_depIdxs = []int32{
	2,  // 0: doublecloud.transfer.v1.endpoint.ColSchema.type:type_name -> doublecloud.transfer.v1.endpoint.ColumnType
	14, // 1: doublecloud.transfer.v1.endpoint.TLSMode.disabled:type_name -> google.protobuf.Empty
	7,  // 2: doublecloud.transfer.v1.endpoint.TLSMode.enabled:type_name -> doublecloud.transfer.v1.endpoint.TLSConfig
	9,  // 3: doublecloud.transfer.v1.endpoint.DataTransformationOptions.headers:type_name -> doublecloud.transfer.v1.endpoint.HeaderValue
	5,  // 4: doublecloud.transfer.v1.endpoint.FieldList.fields:type_name -> doublecloud.transfer.v1.endpoint.ColSchema
	11, // 5: doublecloud.transfer.v1.endpoint.DataSchema.fields:type_name -> doublecloud.transfer.v1.endpoint.FieldList
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_common_proto_init() }
func file_doublecloud_transfer_v1_endpoint_common_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_common_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[1].OneofWrappers = []any{
		(*Secret_Raw)(nil),
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[3].OneofWrappers = []any{
		(*TLSMode_Disabled)(nil),
		(*TLSMode_Enabled)(nil),
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[5].OneofWrappers = []any{
		(*ColumnValue_StringValue)(nil),
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes[9].OneofWrappers = []any{
		(*DataSchema_JsonFields)(nil),
		(*DataSchema_Fields)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_common_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_common_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_endpoint_common_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_common_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_common_proto = out.File
	file_doublecloud_transfer_v1_endpoint_common_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_common_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_common_proto_depIdxs = nil
}
