// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/mssql_source.proto

package endpoint_airbyte

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

type MSSQLSource_MSSQLReplicationMethod int32

const (
	MSSQLSource_MSSQL_REPLICATION_METHOD_UNSPECIFIED MSSQLSource_MSSQLReplicationMethod = 0
	MSSQLSource_STANDARD                             MSSQLSource_MSSQLReplicationMethod = 1
	MSSQLSource_CDC                                  MSSQLSource_MSSQLReplicationMethod = 2
)

// Enum value maps for MSSQLSource_MSSQLReplicationMethod.
var (
	MSSQLSource_MSSQLReplicationMethod_name = map[int32]string{
		0: "MSSQL_REPLICATION_METHOD_UNSPECIFIED",
		1: "STANDARD",
		2: "CDC",
	}
	MSSQLSource_MSSQLReplicationMethod_value = map[string]int32{
		"MSSQL_REPLICATION_METHOD_UNSPECIFIED": 0,
		"STANDARD":                             1,
		"CDC":                                  2,
	}
)

func (x MSSQLSource_MSSQLReplicationMethod) Enum() *MSSQLSource_MSSQLReplicationMethod {
	p := new(MSSQLSource_MSSQLReplicationMethod)
	*p = x
	return p
}

func (x MSSQLSource_MSSQLReplicationMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSSQLSource_MSSQLReplicationMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_enumTypes[0].Descriptor()
}

func (MSSQLSource_MSSQLReplicationMethod) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_enumTypes[0]
}

func (x MSSQLSource_MSSQLReplicationMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSSQLSource_MSSQLReplicationMethod.Descriptor instead.
func (MSSQLSource_MSSQLReplicationMethod) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescGZIP(), []int{0, 0}
}

type MSSQLSource struct {
	state             protoimpl.MessageState             `protogen:"open.v1"`
	Host              string                             `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port              int64                              `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Database          string                             `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	Username          string                             `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password          string                             `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	ReplicationMethod MSSQLSource_MSSQLReplicationMethod `protobuf:"varint,6,opt,name=replication_method,json=replicationMethod,proto3,enum=doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource_MSSQLReplicationMethod" json:"replication_method,omitempty"`
	SslMethod         *MSSQLSource_SSLConfig             `protobuf:"bytes,7,opt,name=ssl_method,json=sslMethod,proto3" json:"ssl_method,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *MSSQLSource) Reset() {
	*x = MSSQLSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MSSQLSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSSQLSource) ProtoMessage() {}

func (x *MSSQLSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSSQLSource.ProtoReflect.Descriptor instead.
func (*MSSQLSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescGZIP(), []int{0}
}

func (x *MSSQLSource) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *MSSQLSource) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *MSSQLSource) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *MSSQLSource) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MSSQLSource) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MSSQLSource) GetReplicationMethod() MSSQLSource_MSSQLReplicationMethod {
	if x != nil {
		return x.ReplicationMethod
	}
	return MSSQLSource_MSSQL_REPLICATION_METHOD_UNSPECIFIED
}

func (x *MSSQLSource) GetSslMethod() *MSSQLSource_SSLConfig {
	if x != nil {
		return x.SslMethod
	}
	return nil
}

type MSSQLSource_SSLUnencrypted struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MSSQLSource_SSLUnencrypted) Reset() {
	*x = MSSQLSource_SSLUnencrypted{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MSSQLSource_SSLUnencrypted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSSQLSource_SSLUnencrypted) ProtoMessage() {}

func (x *MSSQLSource_SSLUnencrypted) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSSQLSource_SSLUnencrypted.ProtoReflect.Descriptor instead.
func (*MSSQLSource_SSLUnencrypted) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescGZIP(), []int{0, 0}
}

type MSSQLSource_SSLEncryptedTrusted struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MSSQLSource_SSLEncryptedTrusted) Reset() {
	*x = MSSQLSource_SSLEncryptedTrusted{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MSSQLSource_SSLEncryptedTrusted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSSQLSource_SSLEncryptedTrusted) ProtoMessage() {}

func (x *MSSQLSource_SSLEncryptedTrusted) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSSQLSource_SSLEncryptedTrusted.ProtoReflect.Descriptor instead.
func (*MSSQLSource_SSLEncryptedTrusted) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescGZIP(), []int{0, 1}
}

type MSSQLSource_SSLEncryptedVerifyCert struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	HostNameInCertificate string                 `protobuf:"bytes,1,opt,name=host_name_in_certificate,json=hostNameInCertificate,proto3" json:"host_name_in_certificate,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *MSSQLSource_SSLEncryptedVerifyCert) Reset() {
	*x = MSSQLSource_SSLEncryptedVerifyCert{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MSSQLSource_SSLEncryptedVerifyCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSSQLSource_SSLEncryptedVerifyCert) ProtoMessage() {}

func (x *MSSQLSource_SSLEncryptedVerifyCert) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSSQLSource_SSLEncryptedVerifyCert.ProtoReflect.Descriptor instead.
func (*MSSQLSource_SSLEncryptedVerifyCert) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescGZIP(), []int{0, 2}
}

func (x *MSSQLSource_SSLEncryptedVerifyCert) GetHostNameInCertificate() string {
	if x != nil {
		return x.HostNameInCertificate
	}
	return ""
}

type MSSQLSource_SSLConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to SslMethod:
	//
	//	*MSSQLSource_SSLConfig_Unencrypted
	//	*MSSQLSource_SSLConfig_EncryptedTrustServerCertificate
	//	*MSSQLSource_SSLConfig_EncryptedVerifyCertificate
	SslMethod     isMSSQLSource_SSLConfig_SslMethod `protobuf_oneof:"ssl_method"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MSSQLSource_SSLConfig) Reset() {
	*x = MSSQLSource_SSLConfig{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MSSQLSource_SSLConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSSQLSource_SSLConfig) ProtoMessage() {}

func (x *MSSQLSource_SSLConfig) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSSQLSource_SSLConfig.ProtoReflect.Descriptor instead.
func (*MSSQLSource_SSLConfig) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescGZIP(), []int{0, 3}
}

func (x *MSSQLSource_SSLConfig) GetSslMethod() isMSSQLSource_SSLConfig_SslMethod {
	if x != nil {
		return x.SslMethod
	}
	return nil
}

func (x *MSSQLSource_SSLConfig) GetUnencrypted() *MSSQLSource_SSLUnencrypted {
	if x != nil {
		if x, ok := x.SslMethod.(*MSSQLSource_SSLConfig_Unencrypted); ok {
			return x.Unencrypted
		}
	}
	return nil
}

func (x *MSSQLSource_SSLConfig) GetEncryptedTrustServerCertificate() *MSSQLSource_SSLEncryptedTrusted {
	if x != nil {
		if x, ok := x.SslMethod.(*MSSQLSource_SSLConfig_EncryptedTrustServerCertificate); ok {
			return x.EncryptedTrustServerCertificate
		}
	}
	return nil
}

func (x *MSSQLSource_SSLConfig) GetEncryptedVerifyCertificate() *MSSQLSource_SSLEncryptedVerifyCert {
	if x != nil {
		if x, ok := x.SslMethod.(*MSSQLSource_SSLConfig_EncryptedVerifyCertificate); ok {
			return x.EncryptedVerifyCertificate
		}
	}
	return nil
}

type isMSSQLSource_SSLConfig_SslMethod interface {
	isMSSQLSource_SSLConfig_SslMethod()
}

type MSSQLSource_SSLConfig_Unencrypted struct {
	Unencrypted *MSSQLSource_SSLUnencrypted `protobuf:"bytes,1,opt,name=unencrypted,proto3,oneof"`
}

type MSSQLSource_SSLConfig_EncryptedTrustServerCertificate struct {
	EncryptedTrustServerCertificate *MSSQLSource_SSLEncryptedTrusted `protobuf:"bytes,2,opt,name=encrypted_trust_server_certificate,json=encryptedTrustServerCertificate,proto3,oneof"`
}

type MSSQLSource_SSLConfig_EncryptedVerifyCertificate struct {
	EncryptedVerifyCertificate *MSSQLSource_SSLEncryptedVerifyCert `protobuf:"bytes,3,opt,name=encrypted_verify_certificate,json=encryptedVerifyCertificate,proto3,oneof"`
}

func (*MSSQLSource_SSLConfig_Unencrypted) isMSSQLSource_SSLConfig_SslMethod() {}

func (*MSSQLSource_SSLConfig_EncryptedTrustServerCertificate) isMSSQLSource_SSLConfig_SslMethod() {}

func (*MSSQLSource_SSLConfig_EncryptedVerifyCertificate) isMSSQLSource_SSLConfig_SslMethod() {}

var File_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDesc = string([]byte{
	0x0a, 0x3b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x6d, 0x73, 0x73, 0x71, 0x6c,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0xf1, 0x07, 0x0a, 0x0b, 0x4d, 0x53, 0x53, 0x51,
	0x4c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x7b, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x4c, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x4d, 0x53, 0x53, 0x51, 0x4c,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x53, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x11, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x5e, 0x0a, 0x0a, 0x73, 0x73, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e,
	0x4d, 0x53, 0x53, 0x51, 0x4c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x53, 0x4c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x73, 0x73, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x1a, 0x10, 0x0a, 0x0e, 0x53, 0x53, 0x4c, 0x55, 0x6e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x1a, 0x15, 0x0a, 0x13, 0x53, 0x53, 0x4c, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x51, 0x0a, 0x16, 0x53, 0x53, 0x4c,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0xb1, 0x03, 0x0a,
	0x09, 0x53, 0x53, 0x4c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x68, 0x0a, 0x0b, 0x75, 0x6e,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x44, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x4d, 0x53, 0x53, 0x51, 0x4c,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x53, 0x4c, 0x55, 0x6e, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x6e, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x98, 0x01, 0x0a, 0x22, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x49, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x4d, 0x53, 0x53,
	0x51, 0x4c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x53, 0x4c, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x1f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x72, 0x75, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x90, 0x01, 0x0a, 0x1c, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74,
	0x65, 0x2e, 0x4d, 0x53, 0x53, 0x51, 0x4c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x53,
	0x4c, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x65, 0x72, 0x74, 0x48, 0x00, 0x52, 0x1a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x73, 0x73, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x22, 0x59, 0x0a, 0x16, 0x4d, 0x53, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x28, 0x0a, 0x24, 0x4d, 0x53,
	0x53, 0x51, 0x4c, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x44, 0x43, 0x10, 0x02, 0x42, 0x5e, 0x5a, 0x5c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescData []byte
)

func file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDesc)))
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_goTypes = []any{
	(MSSQLSource_MSSQLReplicationMethod)(0),    // 0: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.MSSQLReplicationMethod
	(*MSSQLSource)(nil),                        // 1: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource
	(*MSSQLSource_SSLUnencrypted)(nil),         // 2: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLUnencrypted
	(*MSSQLSource_SSLEncryptedTrusted)(nil),    // 3: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLEncryptedTrusted
	(*MSSQLSource_SSLEncryptedVerifyCert)(nil), // 4: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLEncryptedVerifyCert
	(*MSSQLSource_SSLConfig)(nil),              // 5: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLConfig
}
var file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_depIdxs = []int32{
	0, // 0: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.replication_method:type_name -> doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.MSSQLReplicationMethod
	5, // 1: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.ssl_method:type_name -> doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLConfig
	2, // 2: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLConfig.unencrypted:type_name -> doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLUnencrypted
	3, // 3: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLConfig.encrypted_trust_server_certificate:type_name -> doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLEncryptedTrusted
	4, // 4: doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLConfig.encrypted_verify_certificate:type_name -> doublecloud.transfer.v1.endpoint.airbyte.MSSQLSource.SSLEncryptedVerifyCert
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes[4].OneofWrappers = []any{
		(*MSSQLSource_SSLConfig_Unencrypted)(nil),
		(*MSSQLSource_SSLConfig_EncryptedTrustServerCertificate)(nil),
		(*MSSQLSource_SSLConfig_EncryptedVerifyCertificate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_mssql_source_proto_depIdxs = nil
}
