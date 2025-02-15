// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: doublecloud/clickhouse/v1/backup.proto

package clickhouse

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Backup_Type int32

const (
	Backup_TYPE_INVALID Backup_Type = 0
	// Backup created by automated daily schedule
	Backup_TYPE_AUTOMATED Backup_Type = 1
	// Backup created by user request
	Backup_TYPE_MANUAL Backup_Type = 2
)

// Enum value maps for Backup_Type.
var (
	Backup_Type_name = map[int32]string{
		0: "TYPE_INVALID",
		1: "TYPE_AUTOMATED",
		2: "TYPE_MANUAL",
	}
	Backup_Type_value = map[string]int32{
		"TYPE_INVALID":   0,
		"TYPE_AUTOMATED": 1,
		"TYPE_MANUAL":    2,
	}
)

func (x Backup_Type) Enum() *Backup_Type {
	p := new(Backup_Type)
	*p = x
	return p
}

func (x Backup_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Backup_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_clickhouse_v1_backup_proto_enumTypes[0].Descriptor()
}

func (Backup_Type) Type() protoreflect.EnumType {
	return &file_doublecloud_clickhouse_v1_backup_proto_enumTypes[0]
}

func (x Backup_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Backup_Type.Descriptor instead.
func (Backup_Type) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_backup_proto_rawDescGZIP(), []int{0, 0}
}

type Backup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the project that the backup belongs to.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// User defined or autogenerated backup name. Contains shard name until we
	// implement consistent multishard backups.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The time when the backup was created (i.e. when the backup operation completed).
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time when the backup operation was started.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// ID of the associated ClickHouse cluster.
	SourceClusterId string `protobuf:"bytes,6,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Size of backup in bytes
	Size int64 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	// How this backup was created (manual/automatic/etc...)
	Type          Backup_Type `protobuf:"varint,8,opt,name=type,proto3,enum=doublecloud.clickhouse.v1.Backup_Type" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Backup) Reset() {
	*x = Backup{}
	mi := &file_doublecloud_clickhouse_v1_backup_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Backup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backup) ProtoMessage() {}

func (x *Backup) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_clickhouse_v1_backup_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backup.ProtoReflect.Descriptor instead.
func (*Backup) Descriptor() ([]byte, []int) {
	return file_doublecloud_clickhouse_v1_backup_proto_rawDescGZIP(), []int{0}
}

func (x *Backup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Backup) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Backup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Backup) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Backup) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Backup) GetSourceClusterId() string {
	if x != nil {
		return x.SourceClusterId
	}
	return ""
}

func (x *Backup) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Backup) GetType() Backup_Type {
	if x != nil {
		return x.Type
	}
	return Backup_TYPE_INVALID
}

var File_doublecloud_clickhouse_v1_backup_proto protoreflect.FileDescriptor

var file_doublecloud_clickhouse_v1_backup_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4e,
	0x55, 0x41, 0x4c, 0x10, 0x02, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_doublecloud_clickhouse_v1_backup_proto_rawDescOnce sync.Once
	file_doublecloud_clickhouse_v1_backup_proto_rawDescData []byte
)

func file_doublecloud_clickhouse_v1_backup_proto_rawDescGZIP() []byte {
	file_doublecloud_clickhouse_v1_backup_proto_rawDescOnce.Do(func() {
		file_doublecloud_clickhouse_v1_backup_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_clickhouse_v1_backup_proto_rawDesc), len(file_doublecloud_clickhouse_v1_backup_proto_rawDesc)))
	})
	return file_doublecloud_clickhouse_v1_backup_proto_rawDescData
}

var file_doublecloud_clickhouse_v1_backup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_clickhouse_v1_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_clickhouse_v1_backup_proto_goTypes = []any{
	(Backup_Type)(0),              // 0: doublecloud.clickhouse.v1.Backup.Type
	(*Backup)(nil),                // 1: doublecloud.clickhouse.v1.Backup
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_doublecloud_clickhouse_v1_backup_proto_depIdxs = []int32{
	2, // 0: doublecloud.clickhouse.v1.Backup.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: doublecloud.clickhouse.v1.Backup.start_time:type_name -> google.protobuf.Timestamp
	0, // 2: doublecloud.clickhouse.v1.Backup.type:type_name -> doublecloud.clickhouse.v1.Backup.Type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_doublecloud_clickhouse_v1_backup_proto_init() }
func file_doublecloud_clickhouse_v1_backup_proto_init() {
	if File_doublecloud_clickhouse_v1_backup_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_clickhouse_v1_backup_proto_rawDesc), len(file_doublecloud_clickhouse_v1_backup_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_clickhouse_v1_backup_proto_goTypes,
		DependencyIndexes: file_doublecloud_clickhouse_v1_backup_proto_depIdxs,
		EnumInfos:         file_doublecloud_clickhouse_v1_backup_proto_enumTypes,
		MessageInfos:      file_doublecloud_clickhouse_v1_backup_proto_msgTypes,
	}.Build()
	File_doublecloud_clickhouse_v1_backup_proto = out.File
	file_doublecloud_clickhouse_v1_backup_proto_goTypes = nil
	file_doublecloud_clickhouse_v1_backup_proto_depIdxs = nil
}
