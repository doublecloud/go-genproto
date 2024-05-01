// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/s3_source.proto

package endpoint_airbyte

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

type S3Source_Jsonl_UnexpectedFieldBehavior int32

const (
	S3Source_Jsonl_UNEXPECTED_FIELD_BEHAVIOR_UNSPECIFIED S3Source_Jsonl_UnexpectedFieldBehavior = 0
	S3Source_Jsonl_UNEXPECTED_FIELD_BEHAVIOR_IGNORE      S3Source_Jsonl_UnexpectedFieldBehavior = 1
	S3Source_Jsonl_UNEXPECTED_FIELD_BEHAVIOR_INFER       S3Source_Jsonl_UnexpectedFieldBehavior = 2
	S3Source_Jsonl_UNEXPECTED_FIELD_BEHAVIOR_ERROR       S3Source_Jsonl_UnexpectedFieldBehavior = 3
)

// Enum value maps for S3Source_Jsonl_UnexpectedFieldBehavior.
var (
	S3Source_Jsonl_UnexpectedFieldBehavior_name = map[int32]string{
		0: "UNEXPECTED_FIELD_BEHAVIOR_UNSPECIFIED",
		1: "UNEXPECTED_FIELD_BEHAVIOR_IGNORE",
		2: "UNEXPECTED_FIELD_BEHAVIOR_INFER",
		3: "UNEXPECTED_FIELD_BEHAVIOR_ERROR",
	}
	S3Source_Jsonl_UnexpectedFieldBehavior_value = map[string]int32{
		"UNEXPECTED_FIELD_BEHAVIOR_UNSPECIFIED": 0,
		"UNEXPECTED_FIELD_BEHAVIOR_IGNORE":      1,
		"UNEXPECTED_FIELD_BEHAVIOR_INFER":       2,
		"UNEXPECTED_FIELD_BEHAVIOR_ERROR":       3,
	}
)

func (x S3Source_Jsonl_UnexpectedFieldBehavior) Enum() *S3Source_Jsonl_UnexpectedFieldBehavior {
	p := new(S3Source_Jsonl_UnexpectedFieldBehavior)
	*p = x
	return p
}

func (x S3Source_Jsonl_UnexpectedFieldBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (S3Source_Jsonl_UnexpectedFieldBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_enumTypes[0].Descriptor()
}

func (S3Source_Jsonl_UnexpectedFieldBehavior) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_enumTypes[0]
}

func (x S3Source_Jsonl_UnexpectedFieldBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use S3Source_Jsonl_UnexpectedFieldBehavior.Descriptor instead.
func (S3Source_Jsonl_UnexpectedFieldBehavior) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0, 4, 0}
}

type S3Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataset     string             `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	PathPattern string             `protobuf:"bytes,2,opt,name=path_pattern,json=pathPattern,proto3" json:"path_pattern,omitempty"`
	Schema      string             `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Format      *S3Source_Format   `protobuf:"bytes,5,opt,name=format,proto3" json:"format,omitempty"`
	Provider    *S3Source_Provider `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *S3Source) Reset() {
	*x = S3Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Source) ProtoMessage() {}

func (x *S3Source) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Source.ProtoReflect.Descriptor instead.
func (*S3Source) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0}
}

func (x *S3Source) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *S3Source) GetPathPattern() string {
	if x != nil {
		return x.PathPattern
	}
	return ""
}

func (x *S3Source) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *S3Source) GetFormat() *S3Source_Format {
	if x != nil {
		return x.Format
	}
	return nil
}

func (x *S3Source) GetProvider() *S3Source_Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

type S3Source_Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Format:
	//
	//	*S3Source_Format_Csv
	//	*S3Source_Format_Parquet
	//	*S3Source_Format_Avro
	//	*S3Source_Format_Jsonl
	Format isS3Source_Format_Format `protobuf_oneof:"format"`
}

func (x *S3Source_Format) Reset() {
	*x = S3Source_Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Source_Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Source_Format) ProtoMessage() {}

func (x *S3Source_Format) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Source_Format.ProtoReflect.Descriptor instead.
func (*S3Source_Format) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0, 0}
}

func (m *S3Source_Format) GetFormat() isS3Source_Format_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (x *S3Source_Format) GetCsv() *S3Source_Csv {
	if x, ok := x.GetFormat().(*S3Source_Format_Csv); ok {
		return x.Csv
	}
	return nil
}

func (x *S3Source_Format) GetParquet() *S3Source_Parquet {
	if x, ok := x.GetFormat().(*S3Source_Format_Parquet); ok {
		return x.Parquet
	}
	return nil
}

func (x *S3Source_Format) GetAvro() *S3Source_Avro {
	if x, ok := x.GetFormat().(*S3Source_Format_Avro); ok {
		return x.Avro
	}
	return nil
}

func (x *S3Source_Format) GetJsonl() *S3Source_Jsonl {
	if x, ok := x.GetFormat().(*S3Source_Format_Jsonl); ok {
		return x.Jsonl
	}
	return nil
}

type isS3Source_Format_Format interface {
	isS3Source_Format_Format()
}

type S3Source_Format_Csv struct {
	Csv *S3Source_Csv `protobuf:"bytes,4,opt,name=csv,proto3,oneof"`
}

type S3Source_Format_Parquet struct {
	Parquet *S3Source_Parquet `protobuf:"bytes,5,opt,name=parquet,proto3,oneof"`
}

type S3Source_Format_Avro struct {
	Avro *S3Source_Avro `protobuf:"bytes,6,opt,name=avro,proto3,oneof"`
}

type S3Source_Format_Jsonl struct {
	Jsonl *S3Source_Jsonl `protobuf:"bytes,7,opt,name=jsonl,proto3,oneof"`
}

func (*S3Source_Format_Csv) isS3Source_Format_Format() {}

func (*S3Source_Format_Parquet) isS3Source_Format_Format() {}

func (*S3Source_Format_Avro) isS3Source_Format_Format() {}

func (*S3Source_Format_Jsonl) isS3Source_Format_Format() {}

type S3Source_Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket             string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	AwsAccessKeyId     string `protobuf:"bytes,2,opt,name=aws_access_key_id,json=awsAccessKeyId,proto3" json:"aws_access_key_id,omitempty"`
	AwsSecretAccessKey string `protobuf:"bytes,3,opt,name=aws_secret_access_key,json=awsSecretAccessKey,proto3" json:"aws_secret_access_key,omitempty"`
	PathPrefix         string `protobuf:"bytes,4,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	Endpoint           string `protobuf:"bytes,5,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	UseSsl             bool   `protobuf:"varint,6,opt,name=use_ssl,json=useSsl,proto3" json:"use_ssl,omitempty"`
	VerifySslCert      bool   `protobuf:"varint,7,opt,name=verify_ssl_cert,json=verifySslCert,proto3" json:"verify_ssl_cert,omitempty"`
}

func (x *S3Source_Provider) Reset() {
	*x = S3Source_Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Source_Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Source_Provider) ProtoMessage() {}

func (x *S3Source_Provider) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Source_Provider.ProtoReflect.Descriptor instead.
func (*S3Source_Provider) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0, 1}
}

func (x *S3Source_Provider) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3Source_Provider) GetAwsAccessKeyId() string {
	if x != nil {
		return x.AwsAccessKeyId
	}
	return ""
}

func (x *S3Source_Provider) GetAwsSecretAccessKey() string {
	if x != nil {
		return x.AwsSecretAccessKey
	}
	return ""
}

func (x *S3Source_Provider) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

func (x *S3Source_Provider) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *S3Source_Provider) GetUseSsl() bool {
	if x != nil {
		return x.UseSsl
	}
	return false
}

func (x *S3Source_Provider) GetVerifySslCert() bool {
	if x != nil {
		return x.VerifySslCert
	}
	return false
}

type S3Source_Csv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delimiter               string `protobuf:"bytes,2,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	QuoteChar               string `protobuf:"bytes,3,opt,name=quote_char,json=quoteChar,proto3" json:"quote_char,omitempty"`
	EscapeChar              string `protobuf:"bytes,4,opt,name=escape_char,json=escapeChar,proto3" json:"escape_char,omitempty"`
	Encoding                string `protobuf:"bytes,5,opt,name=encoding,proto3" json:"encoding,omitempty"`
	DoubleQuote             bool   `protobuf:"varint,6,opt,name=double_quote,json=doubleQuote,proto3" json:"double_quote,omitempty"`
	NewlinesInValues        bool   `protobuf:"varint,7,opt,name=newlines_in_values,json=newlinesInValues,proto3" json:"newlines_in_values,omitempty"`
	BlockSize               int64  `protobuf:"varint,8,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	AdditionalReaderOptions string `protobuf:"bytes,9,opt,name=additional_reader_options,json=additionalReaderOptions,proto3" json:"additional_reader_options,omitempty"`
	AdvancedOptions         string `protobuf:"bytes,10,opt,name=advanced_options,json=advancedOptions,proto3" json:"advanced_options,omitempty"`
}

func (x *S3Source_Csv) Reset() {
	*x = S3Source_Csv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Source_Csv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Source_Csv) ProtoMessage() {}

func (x *S3Source_Csv) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Source_Csv.ProtoReflect.Descriptor instead.
func (*S3Source_Csv) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0, 2}
}

func (x *S3Source_Csv) GetDelimiter() string {
	if x != nil {
		return x.Delimiter
	}
	return ""
}

func (x *S3Source_Csv) GetQuoteChar() string {
	if x != nil {
		return x.QuoteChar
	}
	return ""
}

func (x *S3Source_Csv) GetEscapeChar() string {
	if x != nil {
		return x.EscapeChar
	}
	return ""
}

func (x *S3Source_Csv) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *S3Source_Csv) GetDoubleQuote() bool {
	if x != nil {
		return x.DoubleQuote
	}
	return false
}

func (x *S3Source_Csv) GetNewlinesInValues() bool {
	if x != nil {
		return x.NewlinesInValues
	}
	return false
}

func (x *S3Source_Csv) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *S3Source_Csv) GetAdditionalReaderOptions() string {
	if x != nil {
		return x.AdditionalReaderOptions
	}
	return ""
}

func (x *S3Source_Csv) GetAdvancedOptions() string {
	if x != nil {
		return x.AdvancedOptions
	}
	return ""
}

type S3Source_Avro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S3Source_Avro) Reset() {
	*x = S3Source_Avro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Source_Avro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Source_Avro) ProtoMessage() {}

func (x *S3Source_Avro) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Source_Avro.ProtoReflect.Descriptor instead.
func (*S3Source_Avro) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0, 3}
}

type S3Source_Jsonl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewlinesInValues        bool                                   `protobuf:"varint,2,opt,name=newlines_in_values,json=newlinesInValues,proto3" json:"newlines_in_values,omitempty"`
	UnexpectedFieldBehavior S3Source_Jsonl_UnexpectedFieldBehavior `protobuf:"varint,3,opt,name=unexpected_field_behavior,json=unexpectedFieldBehavior,proto3,enum=doublecloud.transfer.v1.endpoint.airbyte.S3Source_Jsonl_UnexpectedFieldBehavior" json:"unexpected_field_behavior,omitempty"`
	BlockSize               int64                                  `protobuf:"varint,4,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
}

func (x *S3Source_Jsonl) Reset() {
	*x = S3Source_Jsonl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Source_Jsonl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Source_Jsonl) ProtoMessage() {}

func (x *S3Source_Jsonl) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Source_Jsonl.ProtoReflect.Descriptor instead.
func (*S3Source_Jsonl) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0, 4}
}

func (x *S3Source_Jsonl) GetNewlinesInValues() bool {
	if x != nil {
		return x.NewlinesInValues
	}
	return false
}

func (x *S3Source_Jsonl) GetUnexpectedFieldBehavior() S3Source_Jsonl_UnexpectedFieldBehavior {
	if x != nil {
		return x.UnexpectedFieldBehavior
	}
	return S3Source_Jsonl_UNEXPECTED_FIELD_BEHAVIOR_UNSPECIFIED
}

func (x *S3Source_Jsonl) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

type S3Source_Parquet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BufferSize int64    `protobuf:"varint,2,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	Columns    []string `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	BatchSize  int64    `protobuf:"varint,4,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
}

func (x *S3Source_Parquet) Reset() {
	*x = S3Source_Parquet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Source_Parquet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Source_Parquet) ProtoMessage() {}

func (x *S3Source_Parquet) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Source_Parquet.ProtoReflect.Descriptor instead.
func (*S3Source_Parquet) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP(), []int{0, 5}
}

func (x *S3Source_Parquet) GetBufferSize() int64 {
	if x != nil {
		return x.BufferSize
	}
	return 0
}

func (x *S3Source_Parquet) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *S3Source_Parquet) GetBatchSize() int64 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

var File_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDesc = []byte{
	0x0a, 0x38, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x73, 0x33, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72,
	0x62, 0x79, 0x74, 0x65, 0x22, 0xe7, 0x0d, 0x0a, 0x08, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x51, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74,
	0x65, 0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x57, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61,
	0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x1a, 0xdd, 0x02, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x4a, 0x0a,
	0x03, 0x63, 0x73, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69,
	0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43,
	0x73, 0x76, 0x48, 0x00, 0x52, 0x03, 0x63, 0x73, 0x76, 0x12, 0x56, 0x0a, 0x07, 0x70, 0x61, 0x72,
	0x71, 0x75, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69,
	0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65,
	0x74, 0x12, 0x4d, 0x0a, 0x04, 0x61, 0x76, 0x72, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x41, 0x76, 0x72, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x61, 0x76, 0x72, 0x6f,
	0x12, 0x50, 0x0a, 0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x6a, 0x73, 0x6f,
	0x6e, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x04, 0x1a, 0xfe, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x61, 0x77, 0x73, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x61, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x73, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x73, 0x65, 0x53, 0x73, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x73, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x73, 0x6c, 0x43,
	0x65, 0x72, 0x74, 0x1a, 0xdc, 0x02, 0x0a, 0x03, 0x43, 0x73, 0x76, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x73, 0x63, 0x61,
	0x70, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x73, 0x63, 0x61, 0x70, 0x65, 0x43, 0x68, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x77, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6e, 0x65, 0x77, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x49, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x1a, 0x06, 0x0a, 0x04, 0x41, 0x76, 0x72, 0x6f, 0x1a, 0xa0, 0x03, 0x0a, 0x05, 0x4a,
	0x73, 0x6f, 0x6e, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x77, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x6e, 0x65, 0x77, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x49, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x19, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x50, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74,
	0x65, 0x2e, 0x53, 0x33, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x6c,
	0x2e, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x17, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x17, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x25,
	0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x55, 0x4e, 0x45, 0x58, 0x50,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x42, 0x45, 0x48, 0x41,
	0x56, 0x49, 0x4f, 0x52, 0x5f, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x10, 0x01, 0x12, 0x23, 0x0a,
	0x1f, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x52, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x1a, 0x69, 0x0a,
	0x07, 0x50, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69,
	0x7a, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x42, 0x5e,
	0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_goTypes = []interface{}{
	(S3Source_Jsonl_UnexpectedFieldBehavior)(0), // 0: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Jsonl.UnexpectedFieldBehavior
	(*S3Source)(nil),          // 1: doublecloud.transfer.v1.endpoint.airbyte.S3Source
	(*S3Source_Format)(nil),   // 2: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Format
	(*S3Source_Provider)(nil), // 3: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Provider
	(*S3Source_Csv)(nil),      // 4: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Csv
	(*S3Source_Avro)(nil),     // 5: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Avro
	(*S3Source_Jsonl)(nil),    // 6: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Jsonl
	(*S3Source_Parquet)(nil),  // 7: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Parquet
}
var file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_depIdxs = []int32{
	2, // 0: doublecloud.transfer.v1.endpoint.airbyte.S3Source.format:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Format
	3, // 1: doublecloud.transfer.v1.endpoint.airbyte.S3Source.provider:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Provider
	4, // 2: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Format.csv:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Csv
	7, // 3: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Format.parquet:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Parquet
	5, // 4: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Format.avro:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Avro
	6, // 5: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Format.jsonl:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Jsonl
	0, // 6: doublecloud.transfer.v1.endpoint.airbyte.S3Source.Jsonl.unexpected_field_behavior:type_name -> doublecloud.transfer.v1.endpoint.airbyte.S3Source.Jsonl.UnexpectedFieldBehavior
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Source); i {
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
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Source_Format); i {
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
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Source_Provider); i {
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
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Source_Csv); i {
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
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Source_Avro); i {
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
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Source_Jsonl); i {
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
		file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Source_Parquet); i {
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
	file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*S3Source_Format_Csv)(nil),
		(*S3Source_Format_Parquet)(nil),
		(*S3Source_Format_Avro)(nil),
		(*S3Source_Format_Jsonl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_s3_source_proto_depIdxs = nil
}
