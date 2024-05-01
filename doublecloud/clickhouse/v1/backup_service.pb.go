// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: doublecloud/clickhouse/v1/backup_service.proto

package clickhouse

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

type GetBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the backup to return.
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
}

func (x *GetBackupRequest) Reset() {
	*x = GetBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBackupRequest) ProtoMessage() {}

func (x *GetBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBackupRequest.ProtoReflect.Descriptor instead.
func (*GetBackupRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_backup_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBackupRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

type ListBackupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the project to list backups in.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Paging information of the request
	Paging *v1.Paging `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *ListBackupsRequest) Reset() {
	*x = ListBackupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsRequest) ProtoMessage() {}

func (x *ListBackupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsRequest.ProtoReflect.Descriptor instead.
func (*ListBackupsRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_backup_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListBackupsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ListBackupsRequest) GetPaging() *v1.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type ListBackupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested list of backups.
	Backups []*Backup `protobuf:"bytes,1,rep,name=backups,proto3" json:"backups,omitempty"`
	// Pagination information of the response
	NextPage *v1.NextPage `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *ListBackupsResponse) Reset() {
	*x = ListBackupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsResponse) ProtoMessage() {}

func (x *ListBackupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsResponse.ProtoReflect.Descriptor instead.
func (*ListBackupsResponse) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_backup_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBackupsResponse) GetBackups() []*Backup {
	if x != nil {
		return x.Backups
	}
	return nil
}

func (x *ListBackupsResponse) GetNextPage() *v1.NextPage {
	if x != nil {
		return x.NextPage
	}
	return nil
}

type CreateBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the ClickHouse cluster to back up.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Optional. Name of the ClickHouse cluster backup.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBackupRequest) Reset() {
	*x = CreateBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBackupRequest) ProtoMessage() {}

func (x *CreateBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBackupRequest.ProtoReflect.Descriptor instead.
func (*CreateBackupRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_backup_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBackupRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateBackupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the backup to delete.
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
}

func (x *DeleteBackupRequest) Reset() {
	*x = DeleteBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBackupRequest) ProtoMessage() {}

func (x *DeleteBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBackupRequest.ProtoReflect.Descriptor instead.
func (*DeleteBackupRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_backup_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBackupRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

var File_doublecloud_clickhouse_v1_backup_service_proto protoreflect.FileDescriptor

var file_doublecloud_clickhouse_v1_backup_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x63, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64,
	0x32, 0xf7, 0x02, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x65, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2d, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x2e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_clickhouse_v1_backup_service_proto_rawDescOnce sync.Once
	file_doublecloud_clickhouse_v1_backup_service_proto_rawDescData = file_doublecloud_clickhouse_v1_backup_service_proto_rawDesc
)

func file_doublecloud_clickhouse_v1_backup_service_proto_rawDescGZIP() []byte {
	file_doublecloud_clickhouse_v1_backup_service_proto_rawDescOnce.Do(func() {
		file_doublecloud_clickhouse_v1_backup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_clickhouse_v1_backup_service_proto_rawDescData)
	})
	return file_doublecloud_clickhouse_v1_backup_service_proto_rawDescData
}

var file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_doublecloud_clickhouse_v1_backup_service_proto_goTypes = []interface{}{
	(*GetBackupRequest)(nil),    // 0: doublecloud.clickhouse.v1.GetBackupRequest
	(*ListBackupsRequest)(nil),  // 1: doublecloud.clickhouse.v1.ListBackupsRequest
	(*ListBackupsResponse)(nil), // 2: doublecloud.clickhouse.v1.ListBackupsResponse
	(*CreateBackupRequest)(nil), // 3: doublecloud.clickhouse.v1.CreateBackupRequest
	(*DeleteBackupRequest)(nil), // 4: doublecloud.clickhouse.v1.DeleteBackupRequest
	(*v1.Paging)(nil),           // 5: doublecloud.v1.Paging
	(*Backup)(nil),              // 6: doublecloud.clickhouse.v1.Backup
	(*v1.NextPage)(nil),         // 7: doublecloud.v1.NextPage
	(*v1.Operation)(nil),        // 8: doublecloud.v1.Operation
}
var file_doublecloud_clickhouse_v1_backup_service_proto_depIdxs = []int32{
	5, // 0: doublecloud.clickhouse.v1.ListBackupsRequest.paging:type_name -> doublecloud.v1.Paging
	6, // 1: doublecloud.clickhouse.v1.ListBackupsResponse.backups:type_name -> doublecloud.clickhouse.v1.Backup
	7, // 2: doublecloud.clickhouse.v1.ListBackupsResponse.next_page:type_name -> doublecloud.v1.NextPage
	0, // 3: doublecloud.clickhouse.v1.BackupService.Get:input_type -> doublecloud.clickhouse.v1.GetBackupRequest
	1, // 4: doublecloud.clickhouse.v1.BackupService.List:input_type -> doublecloud.clickhouse.v1.ListBackupsRequest
	3, // 5: doublecloud.clickhouse.v1.BackupService.Create:input_type -> doublecloud.clickhouse.v1.CreateBackupRequest
	4, // 6: doublecloud.clickhouse.v1.BackupService.Delete:input_type -> doublecloud.clickhouse.v1.DeleteBackupRequest
	6, // 7: doublecloud.clickhouse.v1.BackupService.Get:output_type -> doublecloud.clickhouse.v1.Backup
	2, // 8: doublecloud.clickhouse.v1.BackupService.List:output_type -> doublecloud.clickhouse.v1.ListBackupsResponse
	8, // 9: doublecloud.clickhouse.v1.BackupService.Create:output_type -> doublecloud.v1.Operation
	8, // 10: doublecloud.clickhouse.v1.BackupService.Delete:output_type -> doublecloud.v1.Operation
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_doublecloud_clickhouse_v1_backup_service_proto_init() }
func file_doublecloud_clickhouse_v1_backup_service_proto_init() {
	if File_doublecloud_clickhouse_v1_backup_service_proto != nil {
		return
	}
	file_doublecloud_clickhouse_v1_backup_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBackupRequest); i {
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
		file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsRequest); i {
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
		file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsResponse); i {
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
		file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBackupRequest); i {
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
		file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBackupRequest); i {
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
			RawDescriptor: file_doublecloud_clickhouse_v1_backup_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doublecloud_clickhouse_v1_backup_service_proto_goTypes,
		DependencyIndexes: file_doublecloud_clickhouse_v1_backup_service_proto_depIdxs,
		MessageInfos:      file_doublecloud_clickhouse_v1_backup_service_proto_msgTypes,
	}.Build()
	File_doublecloud_clickhouse_v1_backup_service_proto = out.File
	file_doublecloud_clickhouse_v1_backup_service_proto_rawDesc = nil
	file_doublecloud_clickhouse_v1_backup_service_proto_goTypes = nil
	file_doublecloud_clickhouse_v1_backup_service_proto_depIdxs = nil
}
