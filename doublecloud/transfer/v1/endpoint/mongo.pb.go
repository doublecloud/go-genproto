// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/mongo.proto

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

type SrvMongo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname   string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	ReplicaSet string   `protobuf:"bytes,2,opt,name=replica_set,json=replicaSet,proto3" json:"replica_set,omitempty"`
	TlsMode    *TLSMode `protobuf:"bytes,3,opt,name=tls_mode,json=tlsMode,proto3" json:"tls_mode,omitempty"`
}

func (x *SrvMongo) Reset() {
	*x = SrvMongo{}
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SrvMongo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SrvMongo) ProtoMessage() {}

func (x *SrvMongo) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SrvMongo.ProtoReflect.Descriptor instead.
func (*SrvMongo) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{0}
}

func (x *SrvMongo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *SrvMongo) GetReplicaSet() string {
	if x != nil {
		return x.ReplicaSet
	}
	return ""
}

func (x *SrvMongo) GetTlsMode() *TLSMode {
	if x != nil {
		return x.TlsMode
	}
	return nil
}

type OnPremiseMongo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts      []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Port       int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ReplicaSet string   `protobuf:"bytes,5,opt,name=replica_set,json=replicaSet,proto3" json:"replica_set,omitempty"`
	TlsMode    *TLSMode `protobuf:"bytes,6,opt,name=tls_mode,json=tlsMode,proto3" json:"tls_mode,omitempty"`
}

func (x *OnPremiseMongo) Reset() {
	*x = OnPremiseMongo{}
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnPremiseMongo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPremiseMongo) ProtoMessage() {}

func (x *OnPremiseMongo) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPremiseMongo.ProtoReflect.Descriptor instead.
func (*OnPremiseMongo) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{1}
}

func (x *OnPremiseMongo) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *OnPremiseMongo) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *OnPremiseMongo) GetReplicaSet() string {
	if x != nil {
		return x.ReplicaSet
	}
	return ""
}

func (x *OnPremiseMongo) GetTlsMode() *TLSMode {
	if x != nil {
		return x.TlsMode
	}
	return nil
}

type MongoConnectionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Address:
	//
	//	*MongoConnectionOptions_OnPremise
	//	*MongoConnectionOptions_Srv
	Address isMongoConnectionOptions_Address `protobuf_oneof:"address"`
	// User name
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Password for user
	Password *Secret `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Database name associated with the credentials
	AuthSource string `protobuf:"bytes,5,opt,name=auth_source,json=authSource,proto3" json:"auth_source,omitempty"`
}

func (x *MongoConnectionOptions) Reset() {
	*x = MongoConnectionOptions{}
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MongoConnectionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoConnectionOptions) ProtoMessage() {}

func (x *MongoConnectionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoConnectionOptions.ProtoReflect.Descriptor instead.
func (*MongoConnectionOptions) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{2}
}

func (m *MongoConnectionOptions) GetAddress() isMongoConnectionOptions_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (x *MongoConnectionOptions) GetOnPremise() *OnPremiseMongo {
	if x, ok := x.GetAddress().(*MongoConnectionOptions_OnPremise); ok {
		return x.OnPremise
	}
	return nil
}

func (x *MongoConnectionOptions) GetSrv() *SrvMongo {
	if x, ok := x.GetAddress().(*MongoConnectionOptions_Srv); ok {
		return x.Srv
	}
	return nil
}

func (x *MongoConnectionOptions) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MongoConnectionOptions) GetPassword() *Secret {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *MongoConnectionOptions) GetAuthSource() string {
	if x != nil {
		return x.AuthSource
	}
	return ""
}

type isMongoConnectionOptions_Address interface {
	isMongoConnectionOptions_Address()
}

type MongoConnectionOptions_OnPremise struct {
	OnPremise *OnPremiseMongo `protobuf:"bytes,2,opt,name=on_premise,json=onPremise,proto3,oneof"`
}

type MongoConnectionOptions_Srv struct {
	Srv *SrvMongo `protobuf:"bytes,6,opt,name=srv,proto3,oneof"`
}

func (*MongoConnectionOptions_OnPremise) isMongoConnectionOptions_Address() {}

func (*MongoConnectionOptions_Srv) isMongoConnectionOptions_Address() {}

type MongoConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Connection:
	//
	//	*MongoConnection_ConnectionOptions
	Connection isMongoConnection_Connection `protobuf_oneof:"connection"`
}

func (x *MongoConnection) Reset() {
	*x = MongoConnection{}
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MongoConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoConnection) ProtoMessage() {}

func (x *MongoConnection) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoConnection.ProtoReflect.Descriptor instead.
func (*MongoConnection) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{3}
}

func (m *MongoConnection) GetConnection() isMongoConnection_Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (x *MongoConnection) GetConnectionOptions() *MongoConnectionOptions {
	if x, ok := x.GetConnection().(*MongoConnection_ConnectionOptions); ok {
		return x.ConnectionOptions
	}
	return nil
}

type isMongoConnection_Connection interface {
	isMongoConnection_Connection()
}

type MongoConnection_ConnectionOptions struct {
	ConnectionOptions *MongoConnectionOptions `protobuf:"bytes,3,opt,name=connection_options,json=connectionOptions,proto3,oneof"`
}

func (*MongoConnection_ConnectionOptions) isMongoConnection_Connection() {}

type MongoCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseName   string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	CollectionName string `protobuf:"bytes,2,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
}

func (x *MongoCollection) Reset() {
	*x = MongoCollection{}
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MongoCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoCollection) ProtoMessage() {}

func (x *MongoCollection) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoCollection.ProtoReflect.Descriptor instead.
func (*MongoCollection) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{4}
}

func (x *MongoCollection) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *MongoCollection) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

type MongoSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *MongoConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// List of collections for replication. Empty list implies replication of all
	// tables on the deployment. Allowed to use * as collection name.
	Collections []*MongoCollection `protobuf:"bytes,6,rep,name=collections,proto3" json:"collections,omitempty"`
	// List of forbidden collections for replication. Allowed to use * as collection
	// name for forbid all collections of concrete schema.
	ExcludedCollections []*MongoCollection `protobuf:"bytes,7,rep,name=excluded_collections,json=excludedCollections,proto3" json:"excluded_collections,omitempty"`
	// Read mode for mongo client
	SecondaryPreferredMode bool `protobuf:"varint,8,opt,name=secondary_preferred_mode,json=secondaryPreferredMode,proto3" json:"secondary_preferred_mode,omitempty"`
}

func (x *MongoSource) Reset() {
	*x = MongoSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MongoSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoSource) ProtoMessage() {}

func (x *MongoSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoSource.ProtoReflect.Descriptor instead.
func (*MongoSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{5}
}

func (x *MongoSource) GetConnection() *MongoConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *MongoSource) GetCollections() []*MongoCollection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *MongoSource) GetExcludedCollections() []*MongoCollection {
	if x != nil {
		return x.ExcludedCollections
	}
	return nil
}

func (x *MongoSource) GetSecondaryPreferredMode() bool {
	if x != nil {
		return x.SecondaryPreferredMode
	}
	return false
}

type MongoTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *MongoConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Database name
	Database      string        `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	CleanupPolicy CleanupPolicy `protobuf:"varint,6,opt,name=cleanup_policy,json=cleanupPolicy,proto3,enum=doublecloud.transfer.v1.endpoint.CleanupPolicy" json:"cleanup_policy,omitempty"`
}

func (x *MongoTarget) Reset() {
	*x = MongoTarget{}
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MongoTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoTarget) ProtoMessage() {}

func (x *MongoTarget) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoTarget.ProtoReflect.Descriptor instead.
func (*MongoTarget) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{6}
}

func (x *MongoTarget) GetConnection() *MongoConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *MongoTarget) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *MongoTarget) GetCleanupPolicy() CleanupPolicy {
	if x != nil {
		return x.CleanupPolicy
	}
	return CleanupPolicy_CLEANUP_POLICY_UNSPECIFIED
}

var File_doublecloud_transfer_v1_endpoint_mongo_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x1a, 0x2d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8d, 0x01, 0x0a, 0x08, 0x53, 0x72, 0x76, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x65, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x74, 0x6c, 0x73,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x54,
	0x4c, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x74, 0x6c, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x0e, 0x4f, 0x6e, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x73, 0x65, 0x4d, 0x6f, 0x6e,
	0x67, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x65, 0x74, 0x12, 0x44, 0x0a,
	0x08, 0x74, 0x6c, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x54, 0x4c, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x74, 0x6c, 0x73, 0x4d,
	0x6f, 0x64, 0x65, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x05, 0x22, 0xb7, 0x02, 0x0a, 0x16, 0x4d, 0x6f,
	0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x0a, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x6d, 0x69,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x6e, 0x50, 0x72,
	0x65, 0x6d, 0x69, 0x73, 0x65, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e,
	0x50, 0x72, 0x65, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x73, 0x72, 0x76, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x72, 0x76, 0x4d, 0x6f, 0x6e, 0x67, 0x6f,
	0x48, 0x00, 0x52, 0x03, 0x73, 0x72, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x22, 0x90, 0x01, 0x0a, 0x0f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x69, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52,
	0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x03, 0x22, 0x5f, 0x0a, 0x0f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xdb, 0x02, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d,
	0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0b, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x64, 0x0a, 0x14, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x13, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x4a,
	0x04, 0x08, 0x02, 0x10, 0x06, 0x22, 0xda, 0x01, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x51, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e,
	0x67, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0e, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x63,
	0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x06, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_doublecloud_transfer_v1_endpoint_mongo_proto_goTypes = []any{
	(*SrvMongo)(nil),               // 0: doublecloud.transfer.v1.endpoint.SrvMongo
	(*OnPremiseMongo)(nil),         // 1: doublecloud.transfer.v1.endpoint.OnPremiseMongo
	(*MongoConnectionOptions)(nil), // 2: doublecloud.transfer.v1.endpoint.MongoConnectionOptions
	(*MongoConnection)(nil),        // 3: doublecloud.transfer.v1.endpoint.MongoConnection
	(*MongoCollection)(nil),        // 4: doublecloud.transfer.v1.endpoint.MongoCollection
	(*MongoSource)(nil),            // 5: doublecloud.transfer.v1.endpoint.MongoSource
	(*MongoTarget)(nil),            // 6: doublecloud.transfer.v1.endpoint.MongoTarget
	(*TLSMode)(nil),                // 7: doublecloud.transfer.v1.endpoint.TLSMode
	(*Secret)(nil),                 // 8: doublecloud.transfer.v1.endpoint.Secret
	(CleanupPolicy)(0),             // 9: doublecloud.transfer.v1.endpoint.CleanupPolicy
}
var file_doublecloud_transfer_v1_endpoint_mongo_proto_depIdxs = []int32{
	7,  // 0: doublecloud.transfer.v1.endpoint.SrvMongo.tls_mode:type_name -> doublecloud.transfer.v1.endpoint.TLSMode
	7,  // 1: doublecloud.transfer.v1.endpoint.OnPremiseMongo.tls_mode:type_name -> doublecloud.transfer.v1.endpoint.TLSMode
	1,  // 2: doublecloud.transfer.v1.endpoint.MongoConnectionOptions.on_premise:type_name -> doublecloud.transfer.v1.endpoint.OnPremiseMongo
	0,  // 3: doublecloud.transfer.v1.endpoint.MongoConnectionOptions.srv:type_name -> doublecloud.transfer.v1.endpoint.SrvMongo
	8,  // 4: doublecloud.transfer.v1.endpoint.MongoConnectionOptions.password:type_name -> doublecloud.transfer.v1.endpoint.Secret
	2,  // 5: doublecloud.transfer.v1.endpoint.MongoConnection.connection_options:type_name -> doublecloud.transfer.v1.endpoint.MongoConnectionOptions
	3,  // 6: doublecloud.transfer.v1.endpoint.MongoSource.connection:type_name -> doublecloud.transfer.v1.endpoint.MongoConnection
	4,  // 7: doublecloud.transfer.v1.endpoint.MongoSource.collections:type_name -> doublecloud.transfer.v1.endpoint.MongoCollection
	4,  // 8: doublecloud.transfer.v1.endpoint.MongoSource.excluded_collections:type_name -> doublecloud.transfer.v1.endpoint.MongoCollection
	3,  // 9: doublecloud.transfer.v1.endpoint.MongoTarget.connection:type_name -> doublecloud.transfer.v1.endpoint.MongoConnection
	9,  // 10: doublecloud.transfer.v1.endpoint.MongoTarget.cleanup_policy:type_name -> doublecloud.transfer.v1.endpoint.CleanupPolicy
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_mongo_proto_init() }
func file_doublecloud_transfer_v1_endpoint_mongo_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_mongo_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_common_proto_init()
	file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[2].OneofWrappers = []any{
		(*MongoConnectionOptions_OnPremise)(nil),
		(*MongoConnectionOptions_Srv)(nil),
	}
	file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes[3].OneofWrappers = []any{
		(*MongoConnection_ConnectionOptions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_mongo_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_mongo_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_mongo_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_mongo_proto = out.File
	file_doublecloud_transfer_v1_endpoint_mongo_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_mongo_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_mongo_proto_depIdxs = nil
}
