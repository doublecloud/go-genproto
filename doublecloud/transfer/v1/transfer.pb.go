// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/transfer.proto

package transfer

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

type TransferType int32

const (
	TransferType_TRANSFER_TYPE_UNSPECIFIED TransferType = 0
	// Snapshot and increment
	TransferType_SNAPSHOT_AND_INCREMENT TransferType = 1
	// Snapshot
	TransferType_SNAPSHOT_ONLY TransferType = 2
	// Increment
	TransferType_INCREMENT_ONLY TransferType = 3
)

// Enum value maps for TransferType.
var (
	TransferType_name = map[int32]string{
		0: "TRANSFER_TYPE_UNSPECIFIED",
		1: "SNAPSHOT_AND_INCREMENT",
		2: "SNAPSHOT_ONLY",
		3: "INCREMENT_ONLY",
	}
	TransferType_value = map[string]int32{
		"TRANSFER_TYPE_UNSPECIFIED": 0,
		"SNAPSHOT_AND_INCREMENT":    1,
		"SNAPSHOT_ONLY":             2,
		"INCREMENT_ONLY":            3,
	}
)

func (x TransferType) Enum() *TransferType {
	p := new(TransferType)
	*p = x
	return p
}

func (x TransferType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferType) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_transfer_proto_enumTypes[0].Descriptor()
}

func (TransferType) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_transfer_proto_enumTypes[0]
}

func (x TransferType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferType.Descriptor instead.
func (TransferType) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_transfer_proto_rawDescGZIP(), []int{0}
}

type TransferStatus int32

const (
	TransferStatus_TRANSFER_STATUS_UNSPECIFIED TransferStatus = 0
	// Transfer does some work before running
	TransferStatus_CREATING TransferStatus = 1
	// Transfer created but not started by user
	TransferStatus_CREATED TransferStatus = 2
	// Transfer currently doing replication work
	TransferStatus_RUNNING TransferStatus = 3
	// Transfer shutdown
	TransferStatus_STOPPING TransferStatus = 4
	// Transfer stopped by user
	TransferStatus_STOPPED TransferStatus = 5
	// Transfer stopped by system
	TransferStatus_ERROR TransferStatus = 6
	// Transfer copy snapshot
	TransferStatus_SNAPSHOTTING TransferStatus = 7
	// Transfer reach terminal phase
	TransferStatus_DONE TransferStatus = 8
)

// Enum value maps for TransferStatus.
var (
	TransferStatus_name = map[int32]string{
		0: "TRANSFER_STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "CREATED",
		3: "RUNNING",
		4: "STOPPING",
		5: "STOPPED",
		6: "ERROR",
		7: "SNAPSHOTTING",
		8: "DONE",
	}
	TransferStatus_value = map[string]int32{
		"TRANSFER_STATUS_UNSPECIFIED": 0,
		"CREATING":                    1,
		"CREATED":                     2,
		"RUNNING":                     3,
		"STOPPING":                    4,
		"STOPPED":                     5,
		"ERROR":                       6,
		"SNAPSHOTTING":                7,
		"DONE":                        8,
	}
)

func (x TransferStatus) Enum() *TransferStatus {
	p := new(TransferStatus)
	*p = x
	return p
}

func (x TransferStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_transfer_proto_enumTypes[1].Descriptor()
}

func (TransferStatus) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_transfer_proto_enumTypes[1]
}

func (x TransferStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferStatus.Descriptor instead.
func (TransferStatus) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_transfer_proto_rawDescGZIP(), []int{1}
}

// Transfer core entity
type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId   string            `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Name        string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels      map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Source      *Endpoint         `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Target      *Endpoint         `protobuf:"bytes,8,opt,name=target,proto3" json:"target,omitempty"`
	Status      TransferStatus    `protobuf:"varint,10,opt,name=status,proto3,enum=doublecloud.transfer.v1.TransferStatus" json:"status,omitempty"`
	Type        TransferType      `protobuf:"varint,12,opt,name=type,proto3,enum=doublecloud.transfer.v1.TransferType" json:"type,omitempty"`
	Warning     string            `protobuf:"bytes,15,opt,name=warning,proto3" json:"warning,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_transfer_v1_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *Transfer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transfer) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Transfer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Transfer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Transfer) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Transfer) GetSource() *Endpoint {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Transfer) GetTarget() *Endpoint {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Transfer) GetStatus() TransferStatus {
	if x != nil {
		return x.Status
	}
	return TransferStatus_TRANSFER_STATUS_UNSPECIFIED
}

func (x *Transfer) GetType() TransferType {
	if x != nil {
		return x.Type
	}
	return TransferType_TRANSFER_TYPE_UNSPECIFIED
}

func (x *Transfer) GetWarning() string {
	if x != nil {
		return x.Warning
	}
	return ""
}

var File_doublecloud_transfer_v1_transfer_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_transfer_proto_rawDesc = []byte{
	0x0a, 0x26, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x26, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x08, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x70, 0x0a, 0x0c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4e, 0x41, 0x50,
	0x53, 0x48, 0x4f, 0x54, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x43, 0x52, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54,
	0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x43, 0x52, 0x45,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x03, 0x2a, 0x9b, 0x01, 0x0a, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x1b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x08, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_transfer_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_transfer_proto_rawDescData = file_doublecloud_transfer_v1_transfer_proto_rawDesc
)

func file_doublecloud_transfer_v1_transfer_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_transfer_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_transfer_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_transfer_proto_rawDescData
}

var file_doublecloud_transfer_v1_transfer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_doublecloud_transfer_v1_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_transfer_v1_transfer_proto_goTypes = []interface{}{
	(TransferType)(0),   // 0: doublecloud.transfer.v1.TransferType
	(TransferStatus)(0), // 1: doublecloud.transfer.v1.TransferStatus
	(*Transfer)(nil),    // 2: doublecloud.transfer.v1.Transfer
	nil,                 // 3: doublecloud.transfer.v1.Transfer.LabelsEntry
	(*Endpoint)(nil),    // 4: doublecloud.transfer.v1.Endpoint
}
var file_doublecloud_transfer_v1_transfer_proto_depIdxs = []int32{
	3, // 0: doublecloud.transfer.v1.Transfer.labels:type_name -> doublecloud.transfer.v1.Transfer.LabelsEntry
	4, // 1: doublecloud.transfer.v1.Transfer.source:type_name -> doublecloud.transfer.v1.Endpoint
	4, // 2: doublecloud.transfer.v1.Transfer.target:type_name -> doublecloud.transfer.v1.Endpoint
	1, // 3: doublecloud.transfer.v1.Transfer.status:type_name -> doublecloud.transfer.v1.TransferStatus
	0, // 4: doublecloud.transfer.v1.Transfer.type:type_name -> doublecloud.transfer.v1.TransferType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_transfer_proto_init() }
func file_doublecloud_transfer_v1_transfer_proto_init() {
	if File_doublecloud_transfer_v1_transfer_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_transfer_v1_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
			RawDescriptor: file_doublecloud_transfer_v1_transfer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_transfer_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_transfer_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_transfer_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_transfer_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_transfer_proto = out.File
	file_doublecloud_transfer_v1_transfer_proto_rawDesc = nil
	file_doublecloud_transfer_v1_transfer_proto_goTypes = nil
	file_doublecloud_transfer_v1_transfer_proto_depIdxs = nil
}
