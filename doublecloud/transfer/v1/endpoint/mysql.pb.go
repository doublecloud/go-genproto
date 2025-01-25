// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/mysql.proto

package endpoint

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

type OnPremiseMysql struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Database port
	Port  int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Hosts []string `protobuf:"bytes,5,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// TLS settings for server connection. Disabled by default.
	TlsMode       *TLSMode `protobuf:"bytes,6,opt,name=tls_mode,json=tlsMode,proto3" json:"tls_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OnPremiseMysql) Reset() {
	*x = OnPremiseMysql{}
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnPremiseMysql) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPremiseMysql) ProtoMessage() {}

func (x *OnPremiseMysql) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPremiseMysql.ProtoReflect.Descriptor instead.
func (*OnPremiseMysql) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescGZIP(), []int{0}
}

func (x *OnPremiseMysql) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *OnPremiseMysql) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *OnPremiseMysql) GetTlsMode() *TLSMode {
	if x != nil {
		return x.TlsMode
	}
	return nil
}

type MysqlConnection struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Connection:
	//
	//	*MysqlConnection_OnPremise
	Connection    isMysqlConnection_Connection `protobuf_oneof:"connection"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MysqlConnection) Reset() {
	*x = MysqlConnection{}
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MysqlConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MysqlConnection) ProtoMessage() {}

func (x *MysqlConnection) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MysqlConnection.ProtoReflect.Descriptor instead.
func (*MysqlConnection) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescGZIP(), []int{1}
}

func (x *MysqlConnection) GetConnection() isMysqlConnection_Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *MysqlConnection) GetOnPremise() *OnPremiseMysql {
	if x != nil {
		if x, ok := x.Connection.(*MysqlConnection_OnPremise); ok {
			return x.OnPremise
		}
	}
	return nil
}

type isMysqlConnection_Connection interface {
	isMysqlConnection_Connection()
}

type MysqlConnection_OnPremise struct {
	// Connection options for on-premise MySQL
	OnPremise *OnPremiseMysql `protobuf:"bytes,2,opt,name=on_premise,json=onPremise,proto3,oneof"`
}

func (*MysqlConnection_OnPremise) isMysqlConnection_Connection() {}

type MysqlObjectTransferSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Views
	//
	// CREATE VIEW ...
	View ObjectTransferStage `protobuf:"varint,1,opt,name=view,proto3,enum=doublecloud.transfer.v1.endpoint.ObjectTransferStage" json:"view,omitempty"`
	// Routines
	//
	// CREATE PROCEDURE ... ; CREATE FUNCTION ... ;
	Routine ObjectTransferStage `protobuf:"varint,2,opt,name=routine,proto3,enum=doublecloud.transfer.v1.endpoint.ObjectTransferStage" json:"routine,omitempty"`
	// Triggers
	//
	// CREATE TRIGGER ...
	Trigger       ObjectTransferStage `protobuf:"varint,3,opt,name=trigger,proto3,enum=doublecloud.transfer.v1.endpoint.ObjectTransferStage" json:"trigger,omitempty"`
	Tables        ObjectTransferStage `protobuf:"varint,4,opt,name=tables,proto3,enum=doublecloud.transfer.v1.endpoint.ObjectTransferStage" json:"tables,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MysqlObjectTransferSettings) Reset() {
	*x = MysqlObjectTransferSettings{}
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MysqlObjectTransferSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MysqlObjectTransferSettings) ProtoMessage() {}

func (x *MysqlObjectTransferSettings) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MysqlObjectTransferSettings.ProtoReflect.Descriptor instead.
func (*MysqlObjectTransferSettings) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescGZIP(), []int{2}
}

func (x *MysqlObjectTransferSettings) GetView() ObjectTransferStage {
	if x != nil {
		return x.View
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *MysqlObjectTransferSettings) GetRoutine() ObjectTransferStage {
	if x != nil {
		return x.Routine
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *MysqlObjectTransferSettings) GetTrigger() ObjectTransferStage {
	if x != nil {
		return x.Trigger
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *MysqlObjectTransferSettings) GetTables() ObjectTransferStage {
	if x != nil {
		return x.Tables
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

type MysqlSource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Database connection settings
	Connection *MysqlConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Database name
	//
	// You can leave it empty, then it will be possible to transfer tables from several
	// databases at the same time from this source.
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// User for database access.
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Password for database access.
	Password *Secret `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Database timezone
	//
	// Is used for parsing timestamps for saving source timezones. Accepts values from
	// IANA timezone database. Default: local timezone.
	Timezone string `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// Schema migration
	//
	// Select database objects to be transferred during activation or deactivation.
	ObjectTransferSettings *MysqlObjectTransferSettings `protobuf:"bytes,11,opt,name=object_transfer_settings,json=objectTransferSettings,proto3" json:"object_transfer_settings,omitempty"`
	IncludeTablesRegex     []string                     `protobuf:"bytes,12,rep,name=include_tables_regex,json=includeTablesRegex,proto3" json:"include_tables_regex,omitempty"`
	ExcludeTablesRegex     []string                     `protobuf:"bytes,13,rep,name=exclude_tables_regex,json=excludeTablesRegex,proto3" json:"exclude_tables_regex,omitempty"`
	// Database for service tables
	//
	// Default: data source database. Here created technical tables (__tm_keeper,
	// __tm_gtid_keeper).
	ServiceDatabase string `protobuf:"bytes,15,opt,name=service_database,json=serviceDatabase,proto3" json:"service_database,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MysqlSource) Reset() {
	*x = MysqlSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MysqlSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MysqlSource) ProtoMessage() {}

func (x *MysqlSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MysqlSource.ProtoReflect.Descriptor instead.
func (*MysqlSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescGZIP(), []int{3}
}

func (x *MysqlSource) GetConnection() *MysqlConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *MysqlSource) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *MysqlSource) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MysqlSource) GetPassword() *Secret {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *MysqlSource) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *MysqlSource) GetObjectTransferSettings() *MysqlObjectTransferSettings {
	if x != nil {
		return x.ObjectTransferSettings
	}
	return nil
}

func (x *MysqlSource) GetIncludeTablesRegex() []string {
	if x != nil {
		return x.IncludeTablesRegex
	}
	return nil
}

func (x *MysqlSource) GetExcludeTablesRegex() []string {
	if x != nil {
		return x.ExcludeTablesRegex
	}
	return nil
}

func (x *MysqlSource) GetServiceDatabase() string {
	if x != nil {
		return x.ServiceDatabase
	}
	return ""
}

type MysqlTarget struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Database connection settings
	Connection *MysqlConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Database name
	//
	// Allowed to leave it empty, then the tables will be created in databases with the
	// same names as on the source. If this field is empty, then you must fill below db
	// schema for service table.
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// User for database access.
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Password for database access.
	Password *Secret `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Default: NO_AUTO_VALUE_ON_ZERO,NO_DIR_IN_CREATE,NO_ENGINE_SUBSTITUTION.
	SqlMode string `protobuf:"bytes,5,opt,name=sql_mode,json=sqlMode,proto3" json:"sql_mode,omitempty"`
	// Disable constraints checks
	//
	// Recommend to disable for increase replication speed, but if schema contain
	// cascading operations we don't recommend to disable. This option set
	// FOREIGN_KEY_CHECKS=0 and UNIQUE_CHECKS=0.
	SkipConstraintChecks bool `protobuf:"varint,6,opt,name=skip_constraint_checks,json=skipConstraintChecks,proto3" json:"skip_constraint_checks,omitempty"`
	// Database timezone
	//
	// Is used for parsing timestamps for saving source timezones. Accepts values from
	// IANA timezone database. Default: local timezone.
	Timezone string `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// Cleanup policy
	//
	// Cleanup policy for activate, reactivate and reupload processes. Default is
	// DISABLED.
	CleanupPolicy CleanupPolicy `protobuf:"varint,8,opt,name=cleanup_policy,json=cleanupPolicy,proto3,enum=doublecloud.transfer.v1.endpoint.CleanupPolicy" json:"cleanup_policy,omitempty"`
	// Database schema for service table
	//
	// Default: db name. Here created technical tables (__tm_keeper, __tm_gtid_keeper).
	ServiceDatabase string `protobuf:"bytes,15,opt,name=service_database,json=serviceDatabase,proto3" json:"service_database,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,16,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MysqlTarget) Reset() {
	*x = MysqlTarget{}
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MysqlTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MysqlTarget) ProtoMessage() {}

func (x *MysqlTarget) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MysqlTarget.ProtoReflect.Descriptor instead.
func (*MysqlTarget) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescGZIP(), []int{4}
}

func (x *MysqlTarget) GetConnection() *MysqlConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *MysqlTarget) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *MysqlTarget) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MysqlTarget) GetPassword() *Secret {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *MysqlTarget) GetSqlMode() string {
	if x != nil {
		return x.SqlMode
	}
	return ""
}

func (x *MysqlTarget) GetSkipConstraintChecks() bool {
	if x != nil {
		return x.SkipConstraintChecks
	}
	return false
}

func (x *MysqlTarget) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *MysqlTarget) GetCleanupPolicy() CleanupPolicy {
	if x != nil {
		return x.CleanupPolicy
	}
	return CleanupPolicy_CLEANUP_POLICY_UNSPECIFIED
}

func (x *MysqlTarget) GetServiceDatabase() string {
	if x != nil {
		return x.ServiceDatabase
	}
	return ""
}

func (x *MysqlTarget) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

var File_doublecloud_transfer_v1_endpoint_mysql_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x1a, 0x2d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8c, 0x01, 0x0a, 0x0e, 0x4f, 0x6e, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x73, 0x65, 0x4d, 0x79, 0x73,
	0x71, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x08,
	0x74, 0x6c, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x54, 0x4c, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x74, 0x6c, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x05, 0x22, 0x78,
	0x0a, 0x0f, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x51, 0x0a, 0x0a, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x6e, 0x50, 0x72, 0x65, 0x6d, 0x69,
	0x73, 0x65, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e, 0x50, 0x72, 0x65,
	0x6d, 0x69, 0x73, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0xd9, 0x02, 0x0a, 0x1b, 0x4d, 0x79, 0x73,
	0x71, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x49, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x04, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x4f, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x07, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x07, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x06, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x22, 0x8c, 0x04, 0x0a, 0x0b, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x79, 0x73, 0x71,
	0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x77, 0x0a, 0x18, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d,
	0x79, 0x73, 0x71, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x16, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0b, 0x4a, 0x04, 0x08,
	0x0e, 0x10, 0x0f, 0x22, 0xf5, 0x03, 0x0a, 0x0b, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x51, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x71, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x71, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x6b, 0x69, 0x70, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x56, 0x0a, 0x0e, 0x63, 0x6c, 0x65,
	0x61, 0x6e, 0x75, 0x70, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x0d, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0f, 0x42, 0x4e, 0x5a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescData []byte
)

func file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDesc)))
	})
	return file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_doublecloud_transfer_v1_endpoint_mysql_proto_goTypes = []any{
	(*OnPremiseMysql)(nil),              // 0: doublecloud.transfer.v1.endpoint.OnPremiseMysql
	(*MysqlConnection)(nil),             // 1: doublecloud.transfer.v1.endpoint.MysqlConnection
	(*MysqlObjectTransferSettings)(nil), // 2: doublecloud.transfer.v1.endpoint.MysqlObjectTransferSettings
	(*MysqlSource)(nil),                 // 3: doublecloud.transfer.v1.endpoint.MysqlSource
	(*MysqlTarget)(nil),                 // 4: doublecloud.transfer.v1.endpoint.MysqlTarget
	(*TLSMode)(nil),                     // 5: doublecloud.transfer.v1.endpoint.TLSMode
	(ObjectTransferStage)(0),            // 6: doublecloud.transfer.v1.endpoint.ObjectTransferStage
	(*Secret)(nil),                      // 7: doublecloud.transfer.v1.endpoint.Secret
	(CleanupPolicy)(0),                  // 8: doublecloud.transfer.v1.endpoint.CleanupPolicy
}
var file_doublecloud_transfer_v1_endpoint_mysql_proto_depIdxs = []int32{
	5,  // 0: doublecloud.transfer.v1.endpoint.OnPremiseMysql.tls_mode:type_name -> doublecloud.transfer.v1.endpoint.TLSMode
	0,  // 1: doublecloud.transfer.v1.endpoint.MysqlConnection.on_premise:type_name -> doublecloud.transfer.v1.endpoint.OnPremiseMysql
	6,  // 2: doublecloud.transfer.v1.endpoint.MysqlObjectTransferSettings.view:type_name -> doublecloud.transfer.v1.endpoint.ObjectTransferStage
	6,  // 3: doublecloud.transfer.v1.endpoint.MysqlObjectTransferSettings.routine:type_name -> doublecloud.transfer.v1.endpoint.ObjectTransferStage
	6,  // 4: doublecloud.transfer.v1.endpoint.MysqlObjectTransferSettings.trigger:type_name -> doublecloud.transfer.v1.endpoint.ObjectTransferStage
	6,  // 5: doublecloud.transfer.v1.endpoint.MysqlObjectTransferSettings.tables:type_name -> doublecloud.transfer.v1.endpoint.ObjectTransferStage
	1,  // 6: doublecloud.transfer.v1.endpoint.MysqlSource.connection:type_name -> doublecloud.transfer.v1.endpoint.MysqlConnection
	7,  // 7: doublecloud.transfer.v1.endpoint.MysqlSource.password:type_name -> doublecloud.transfer.v1.endpoint.Secret
	2,  // 8: doublecloud.transfer.v1.endpoint.MysqlSource.object_transfer_settings:type_name -> doublecloud.transfer.v1.endpoint.MysqlObjectTransferSettings
	1,  // 9: doublecloud.transfer.v1.endpoint.MysqlTarget.connection:type_name -> doublecloud.transfer.v1.endpoint.MysqlConnection
	7,  // 10: doublecloud.transfer.v1.endpoint.MysqlTarget.password:type_name -> doublecloud.transfer.v1.endpoint.Secret
	8,  // 11: doublecloud.transfer.v1.endpoint.MysqlTarget.cleanup_policy:type_name -> doublecloud.transfer.v1.endpoint.CleanupPolicy
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_mysql_proto_init() }
func file_doublecloud_transfer_v1_endpoint_mysql_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_mysql_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_init()
	file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes[1].OneofWrappers = []any{
		(*MysqlConnection_OnPremise)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_mysql_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_mysql_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_mysql_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_mysql_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_mysql_proto = out.File
	file_doublecloud_transfer_v1_endpoint_mysql_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_mysql_proto_depIdxs = nil
}
