// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: doublecloud/organizationmanager/v1/group_mapping.proto

package organizationmanager

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

// Group mapping represents which external (federated) groups should match which internal (cloud) groups
type GroupMappingItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// External group id (received from identity provider)
	ExternalGroupId string `protobuf:"bytes,1,opt,name=external_group_id,json=externalGroupId,proto3" json:"external_group_id,omitempty"`
	// Internal cloud group id
	InternalGroupId string `protobuf:"bytes,2,opt,name=internal_group_id,json=internalGroupId,proto3" json:"internal_group_id,omitempty"`
}

func (x *GroupMappingItem) Reset() {
	*x = GroupMappingItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMappingItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMappingItem) ProtoMessage() {}

func (x *GroupMappingItem) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMappingItem.ProtoReflect.Descriptor instead.
func (*GroupMappingItem) Descriptor() ([]byte, []int) {
	return file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *GroupMappingItem) GetExternalGroupId() string {
	if x != nil {
		return x.ExternalGroupId
	}
	return ""
}

func (x *GroupMappingItem) GetInternalGroupId() string {
	if x != nil {
		return x.InternalGroupId
	}
	return ""
}

// Group synchronization status for a specific federation
// Absence of this object for a federation means that there is no group synchronization set of for the federation.
type GroupMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Federation id
	FederationId string `protobuf:"bytes,1,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
	// Flag to show whether group synchronization should be enabled for this federation.
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *GroupMapping) Reset() {
	*x = GroupMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMapping) ProtoMessage() {}

func (x *GroupMapping) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMapping.ProtoReflect.Descriptor instead.
func (*GroupMapping) Descriptor() ([]byte, []int) {
	return file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescGZIP(), []int{1}
}

func (x *GroupMapping) GetFederationId() string {
	if x != nil {
		return x.FederationId
	}
	return ""
}

func (x *GroupMapping) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_doublecloud_organizationmanager_v1_group_mapping_proto protoreflect.FileDescriptor

var file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDesc = []byte{
	0x0a, 0x36, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x6a, 0x0a, 0x10,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x5b, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescOnce sync.Once
	file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescData = file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDesc
)

func file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescGZIP() []byte {
	file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescOnce.Do(func() {
		file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescData)
	})
	return file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDescData
}

var file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doublecloud_organizationmanager_v1_group_mapping_proto_goTypes = []interface{}{
	(*GroupMappingItem)(nil), // 0: doublecloud.organizationmanager.v1.GroupMappingItem
	(*GroupMapping)(nil),     // 1: doublecloud.organizationmanager.v1.GroupMapping
}
var file_doublecloud_organizationmanager_v1_group_mapping_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doublecloud_organizationmanager_v1_group_mapping_proto_init() }
func file_doublecloud_organizationmanager_v1_group_mapping_proto_init() {
	if File_doublecloud_organizationmanager_v1_group_mapping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMappingItem); i {
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
		file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMapping); i {
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
			RawDescriptor: file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_organizationmanager_v1_group_mapping_proto_goTypes,
		DependencyIndexes: file_doublecloud_organizationmanager_v1_group_mapping_proto_depIdxs,
		MessageInfos:      file_doublecloud_organizationmanager_v1_group_mapping_proto_msgTypes,
	}.Build()
	File_doublecloud_organizationmanager_v1_group_mapping_proto = out.File
	file_doublecloud_organizationmanager_v1_group_mapping_proto_rawDesc = nil
	file_doublecloud_organizationmanager_v1_group_mapping_proto_goTypes = nil
	file_doublecloud_organizationmanager_v1_group_mapping_proto_depIdxs = nil
}
