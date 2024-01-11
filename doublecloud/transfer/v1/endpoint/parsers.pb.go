// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/parsers.proto

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

type ProtoMessagePackageType int32

const (
	ProtoMessagePackageType_PROTO_MESSAGE_PACKAGE_TYPE_UNSPECIFIED ProtoMessagePackageType = 0
	ProtoMessagePackageType_PROTOSEQ                               ProtoMessagePackageType = 1
	ProtoMessagePackageType_REPEATED                               ProtoMessagePackageType = 2
	ProtoMessagePackageType_SINGLE_MESSAGE                         ProtoMessagePackageType = 3
)

// Enum value maps for ProtoMessagePackageType.
var (
	ProtoMessagePackageType_name = map[int32]string{
		0: "PROTO_MESSAGE_PACKAGE_TYPE_UNSPECIFIED",
		1: "PROTOSEQ",
		2: "REPEATED",
		3: "SINGLE_MESSAGE",
	}
	ProtoMessagePackageType_value = map[string]int32{
		"PROTO_MESSAGE_PACKAGE_TYPE_UNSPECIFIED": 0,
		"PROTOSEQ":                               1,
		"REPEATED":                               2,
		"SINGLE_MESSAGE":                         3,
	}
)

func (x ProtoMessagePackageType) Enum() *ProtoMessagePackageType {
	p := new(ProtoMessagePackageType)
	*p = x
	return p
}

func (x ProtoMessagePackageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoMessagePackageType) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_enumTypes[0].Descriptor()
}

func (ProtoMessagePackageType) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_parsers_proto_enumTypes[0]
}

func (x ProtoMessagePackageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtoMessagePackageType.Descriptor instead.
func (ProtoMessagePackageType) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{0}
}

type Parser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Parser:
	//
	//	*Parser_JsonParser
	//	*Parser_TskvParser
	Parser isParser_Parser `protobuf_oneof:"parser"`
}

func (x *Parser) Reset() {
	*x = Parser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parser) ProtoMessage() {}

func (x *Parser) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parser.ProtoReflect.Descriptor instead.
func (*Parser) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{0}
}

func (m *Parser) GetParser() isParser_Parser {
	if m != nil {
		return m.Parser
	}
	return nil
}

func (x *Parser) GetJsonParser() *GenericParserCommon {
	if x, ok := x.GetParser().(*Parser_JsonParser); ok {
		return x.JsonParser
	}
	return nil
}

func (x *Parser) GetTskvParser() *GenericParserCommon {
	if x, ok := x.GetParser().(*Parser_TskvParser); ok {
		return x.TskvParser
	}
	return nil
}

type isParser_Parser interface {
	isParser_Parser()
}

type Parser_JsonParser struct {
	JsonParser *GenericParserCommon `protobuf:"bytes,1,opt,name=json_parser,json=jsonParser,proto3,oneof"`
}

type Parser_TskvParser struct {
	TskvParser *GenericParserCommon `protobuf:"bytes,6,opt,name=tskv_parser,json=tskvParser,proto3,oneof"`
}

func (*Parser_JsonParser) isParser_Parser() {}

func (*Parser_TskvParser) isParser_Parser() {}

type GenericParserCommon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSchema *DataSchema `protobuf:"bytes,1,opt,name=data_schema,json=dataSchema,proto3" json:"data_schema,omitempty"`
	// Allow null keys, if no - null keys will be putted to unparsed data
	NullKeysAllowed bool `protobuf:"varint,2,opt,name=null_keys_allowed,json=nullKeysAllowed,proto3" json:"null_keys_allowed,omitempty"`
	// Will add _rest column for all unknown fields
	AddRestColumn bool `protobuf:"varint,3,opt,name=add_rest_column,json=addRestColumn,proto3" json:"add_rest_column,omitempty"`
}

func (x *GenericParserCommon) Reset() {
	*x = GenericParserCommon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericParserCommon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericParserCommon) ProtoMessage() {}

func (x *GenericParserCommon) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericParserCommon.ProtoReflect.Descriptor instead.
func (*GenericParserCommon) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{1}
}

func (x *GenericParserCommon) GetDataSchema() *DataSchema {
	if x != nil {
		return x.DataSchema
	}
	return nil
}

func (x *GenericParserCommon) GetNullKeysAllowed() bool {
	if x != nil {
		return x.NullKeysAllowed
	}
	return false
}

func (x *GenericParserCommon) GetAddRestColumn() bool {
	if x != nil {
		return x.AddRestColumn
	}
	return false
}

type ProtoParser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoDesc      *ProtoDesc              `protobuf:"bytes,1,opt,name=proto_desc,json=protoDesc,proto3" json:"proto_desc,omitempty"`
	MsgPackageType ProtoMessagePackageType `protobuf:"varint,2,opt,name=msg_package_type,json=msgPackageType,proto3,enum=doublecloud.transfer.v1.endpoint.ProtoMessagePackageType" json:"msg_package_type,omitempty"`
	// If package type is repeated - name of message containing repeated target field;
	// else - name of target message
	MsgName        string           `protobuf:"bytes,3,opt,name=msg_name,json=msgName,proto3" json:"msg_name,omitempty"`
	PrimaryKeys    []string         `protobuf:"bytes,4,rep,name=primary_keys,json=primaryKeys,proto3" json:"primary_keys,omitempty"`
	IncludedFields *ProtoDataSchema `protobuf:"bytes,5,opt,name=included_fields,json=includedFields,proto3" json:"included_fields,omitempty"`
	// Allow null keys, if no - null keys will be putted to unparsed data
	NullKeysAllowed bool `protobuf:"varint,6,opt,name=null_keys_allowed,json=nullKeysAllowed,proto3" json:"null_keys_allowed,omitempty"`
}

func (x *ProtoParser) Reset() {
	*x = ProtoParser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoParser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoParser) ProtoMessage() {}

func (x *ProtoParser) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoParser.ProtoReflect.Descriptor instead.
func (*ProtoParser) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoParser) GetProtoDesc() *ProtoDesc {
	if x != nil {
		return x.ProtoDesc
	}
	return nil
}

func (x *ProtoParser) GetMsgPackageType() ProtoMessagePackageType {
	if x != nil {
		return x.MsgPackageType
	}
	return ProtoMessagePackageType_PROTO_MESSAGE_PACKAGE_TYPE_UNSPECIFIED
}

func (x *ProtoParser) GetMsgName() string {
	if x != nil {
		return x.MsgName
	}
	return ""
}

func (x *ProtoParser) GetPrimaryKeys() []string {
	if x != nil {
		return x.PrimaryKeys
	}
	return nil
}

func (x *ProtoParser) GetIncludedFields() *ProtoDataSchema {
	if x != nil {
		return x.IncludedFields
	}
	return nil
}

func (x *ProtoParser) GetNullKeysAllowed() bool {
	if x != nil {
		return x.NullKeysAllowed
	}
	return false
}

type ProtoDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Desc:
	//
	//	*ProtoDesc_DescFile
	Desc isProtoDesc_Desc `protobuf_oneof:"desc"`
}

func (x *ProtoDesc) Reset() {
	*x = ProtoDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoDesc) ProtoMessage() {}

func (x *ProtoDesc) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoDesc.ProtoReflect.Descriptor instead.
func (*ProtoDesc) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{3}
}

func (m *ProtoDesc) GetDesc() isProtoDesc_Desc {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (x *ProtoDesc) GetDescFile() []byte {
	if x, ok := x.GetDesc().(*ProtoDesc_DescFile); ok {
		return x.DescFile
	}
	return nil
}

type isProtoDesc_Desc interface {
	isProtoDesc_Desc()
}

type ProtoDesc_DescFile struct {
	DescFile []byte `protobuf:"bytes,1,opt,name=desc_file,json=descFile,proto3,oneof"`
}

func (*ProtoDesc_DescFile) isProtoDesc_Desc() {}

type ProtoDataSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Schema:
	//
	//	*ProtoDataSchema_ColParamsList
	Schema isProtoDataSchema_Schema `protobuf_oneof:"schema"`
}

func (x *ProtoDataSchema) Reset() {
	*x = ProtoDataSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoDataSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoDataSchema) ProtoMessage() {}

func (x *ProtoDataSchema) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoDataSchema.ProtoReflect.Descriptor instead.
func (*ProtoDataSchema) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{4}
}

func (m *ProtoDataSchema) GetSchema() isProtoDataSchema_Schema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (x *ProtoDataSchema) GetColParamsList() *ColumnParamsList {
	if x, ok := x.GetSchema().(*ProtoDataSchema_ColParamsList); ok {
		return x.ColParamsList
	}
	return nil
}

type isProtoDataSchema_Schema interface {
	isProtoDataSchema_Schema()
}

type ProtoDataSchema_ColParamsList struct {
	ColParamsList *ColumnParamsList `protobuf:"bytes,1,opt,name=col_params_list,json=colParamsList,proto3,oneof"`
}

func (*ProtoDataSchema_ColParamsList) isProtoDataSchema_Schema() {}

type ColumnParamsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColParams []*ColumnParams `protobuf:"bytes,1,rep,name=col_params,json=colParams,proto3" json:"col_params,omitempty"`
}

func (x *ColumnParamsList) Reset() {
	*x = ColumnParamsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnParamsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnParamsList) ProtoMessage() {}

func (x *ColumnParamsList) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnParamsList.ProtoReflect.Descriptor instead.
func (*ColumnParamsList) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{5}
}

func (x *ColumnParamsList) GetColParams() []*ColumnParams {
	if x != nil {
		return x.ColParams
	}
	return nil
}

type ColumnParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required bool   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *ColumnParams) Reset() {
	*x = ColumnParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnParams) ProtoMessage() {}

func (x *ColumnParams) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnParams.ProtoReflect.Descriptor instead.
func (*ColumnParams) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP(), []int{6}
}

func (x *ColumnParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnParams) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

var File_doublecloud_transfer_v1_endpoint_parsers_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x1a, 0x2d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x0b,
	0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x6a, 0x73, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x0b, 0x74, 0x73, 0x6b, 0x76, 0x5f, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x73, 0x6b, 0x76, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x42, 0x08, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6e, 0x75,
	0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x61, 0x64, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x52, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x84, 0x03, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x63, 0x0a, 0x10, 0x6d, 0x73, 0x67, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x6d, 0x73, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x5a, 0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6e, 0x75, 0x6c,
	0x6c, 0x4b, 0x65, 0x79, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x32, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x09, 0x64, 0x65, 0x73,
	0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08,
	0x64, 0x65, 0x73, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x22, 0x79, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x5c, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x61, 0x0a, 0x10, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x4d, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x3e,
	0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2a, 0x75,
	0x0a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x52, 0x4f,
	0x54, 0x4f, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x41,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x53, 0x45,
	0x51, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x10, 0x03, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_parsers_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_doublecloud_transfer_v1_endpoint_parsers_proto_goTypes = []interface{}{
	(ProtoMessagePackageType)(0), // 0: doublecloud.transfer.v1.endpoint.ProtoMessagePackageType
	(*Parser)(nil),               // 1: doublecloud.transfer.v1.endpoint.Parser
	(*GenericParserCommon)(nil),  // 2: doublecloud.transfer.v1.endpoint.GenericParserCommon
	(*ProtoParser)(nil),          // 3: doublecloud.transfer.v1.endpoint.ProtoParser
	(*ProtoDesc)(nil),            // 4: doublecloud.transfer.v1.endpoint.ProtoDesc
	(*ProtoDataSchema)(nil),      // 5: doublecloud.transfer.v1.endpoint.ProtoDataSchema
	(*ColumnParamsList)(nil),     // 6: doublecloud.transfer.v1.endpoint.ColumnParamsList
	(*ColumnParams)(nil),         // 7: doublecloud.transfer.v1.endpoint.ColumnParams
	(*DataSchema)(nil),           // 8: doublecloud.transfer.v1.endpoint.DataSchema
}
var file_doublecloud_transfer_v1_endpoint_parsers_proto_depIdxs = []int32{
	2, // 0: doublecloud.transfer.v1.endpoint.Parser.json_parser:type_name -> doublecloud.transfer.v1.endpoint.GenericParserCommon
	2, // 1: doublecloud.transfer.v1.endpoint.Parser.tskv_parser:type_name -> doublecloud.transfer.v1.endpoint.GenericParserCommon
	8, // 2: doublecloud.transfer.v1.endpoint.GenericParserCommon.data_schema:type_name -> doublecloud.transfer.v1.endpoint.DataSchema
	4, // 3: doublecloud.transfer.v1.endpoint.ProtoParser.proto_desc:type_name -> doublecloud.transfer.v1.endpoint.ProtoDesc
	0, // 4: doublecloud.transfer.v1.endpoint.ProtoParser.msg_package_type:type_name -> doublecloud.transfer.v1.endpoint.ProtoMessagePackageType
	5, // 5: doublecloud.transfer.v1.endpoint.ProtoParser.included_fields:type_name -> doublecloud.transfer.v1.endpoint.ProtoDataSchema
	6, // 6: doublecloud.transfer.v1.endpoint.ProtoDataSchema.col_params_list:type_name -> doublecloud.transfer.v1.endpoint.ColumnParamsList
	7, // 7: doublecloud.transfer.v1.endpoint.ColumnParamsList.col_params:type_name -> doublecloud.transfer.v1.endpoint.ColumnParams
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_parsers_proto_init() }
func file_doublecloud_transfer_v1_endpoint_parsers_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_parsers_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parser); i {
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
		file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericParserCommon); i {
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
		file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoParser); i {
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
		file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoDesc); i {
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
		file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoDataSchema); i {
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
		file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnParamsList); i {
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
		file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnParams); i {
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
	file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Parser_JsonParser)(nil),
		(*Parser_TskvParser)(nil),
	}
	file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ProtoDesc_DescFile)(nil),
	}
	file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ProtoDataSchema_ColParamsList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_parsers_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_parsers_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_endpoint_parsers_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_parsers_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_parsers_proto = out.File
	file_doublecloud_transfer_v1_endpoint_parsers_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_parsers_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_parsers_proto_depIdxs = nil
}
