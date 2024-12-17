// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/jira_source.proto

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

type JiraSource_Expand int32

const (
	JiraSource_EXPAND_UNSPECIFIED JiraSource_Expand = 0
	JiraSource_RENDERED_FIELDS    JiraSource_Expand = 1
	JiraSource_TRANSITIONS        JiraSource_Expand = 2
	JiraSource_CHANGELOG          JiraSource_Expand = 3
)

// Enum value maps for JiraSource_Expand.
var (
	JiraSource_Expand_name = map[int32]string{
		0: "EXPAND_UNSPECIFIED",
		1: "RENDERED_FIELDS",
		2: "TRANSITIONS",
		3: "CHANGELOG",
	}
	JiraSource_Expand_value = map[string]int32{
		"EXPAND_UNSPECIFIED": 0,
		"RENDERED_FIELDS":    1,
		"TRANSITIONS":        2,
		"CHANGELOG":          3,
	}
)

func (x JiraSource_Expand) Enum() *JiraSource_Expand {
	p := new(JiraSource_Expand)
	*p = x
	return p
}

func (x JiraSource_Expand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JiraSource_Expand) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_enumTypes[0].Descriptor()
}

func (JiraSource_Expand) Type() protoreflect.EnumType {
	return &file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_enumTypes[0]
}

func (x JiraSource_Expand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JiraSource_Expand.Descriptor instead.
func (JiraSource_Expand) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescGZIP(), []int{0, 0}
}

type JiraSource struct {
	state                     protoimpl.MessageState `protogen:"open.v1"`
	ApiToken                  string                 `protobuf:"bytes,1,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	Domain                    string                 `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Email                     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Projects                  []string               `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty"`
	StartDate                 string                 `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	IssuesStreamExpandWith    []JiraSource_Expand    `protobuf:"varint,6,rep,packed,name=issues_stream_expand_with,json=issuesStreamExpandWith,proto3,enum=doublecloud.transfer.v1.endpoint.airbyte.JiraSource_Expand" json:"issues_stream_expand_with,omitempty"`
	EnableExperimentalStreams bool                   `protobuf:"varint,7,opt,name=enable_experimental_streams,json=enableExperimentalStreams,proto3" json:"enable_experimental_streams,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *JiraSource) Reset() {
	*x = JiraSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JiraSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JiraSource) ProtoMessage() {}

func (x *JiraSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JiraSource.ProtoReflect.Descriptor instead.
func (*JiraSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescGZIP(), []int{0}
}

func (x *JiraSource) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

func (x *JiraSource) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *JiraSource) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *JiraSource) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *JiraSource) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *JiraSource) GetIssuesStreamExpandWith() []JiraSource_Expand {
	if x != nil {
		return x.IssuesStreamExpandWith
	}
	return nil
}

func (x *JiraSource) GetEnableExperimentalStreams() bool {
	if x != nil {
		return x.EnableExperimentalStreams
	}
	return false
}

var File_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x6a, 0x69, 0x72, 0x61, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61,
	0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0xa1, 0x03, 0x0a, 0x0a, 0x4a, 0x69, 0x72, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x76, 0x0a, 0x19, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x70,
	0x61, 0x6e, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x4a, 0x69, 0x72, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x16, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x57,
	0x69, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x1b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x22, 0x55, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a,
	0x12, 0x45, 0x58, 0x50, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x45,
	0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x4c, 0x4f, 0x47, 0x10, 0x03, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescData = file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDesc
)

func file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescData)
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_goTypes = []any{
	(JiraSource_Expand)(0), // 0: doublecloud.transfer.v1.endpoint.airbyte.JiraSource.Expand
	(*JiraSource)(nil),     // 1: doublecloud.transfer.v1.endpoint.airbyte.JiraSource
}
var file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_depIdxs = []int32{
	0, // 0: doublecloud.transfer.v1.endpoint.airbyte.JiraSource.issues_stream_expand_with:type_name -> doublecloud.transfer.v1.endpoint.airbyte.JiraSource.Expand
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_depIdxs,
		EnumInfos:         file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_enumTypes,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_rawDesc = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_jira_source_proto_depIdxs = nil
}
