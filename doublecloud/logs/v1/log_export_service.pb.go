// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: doublecloud/logs/v1/log_export_service.proto

package logs

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

type ListExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string     `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Paging    *v1.Paging `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *ListExportRequest) Reset() {
	*x = ListExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExportRequest) ProtoMessage() {}

func (x *ListExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExportRequest.ProtoReflect.Descriptor instead.
func (*ListExportRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListExportRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ListExportRequest) GetPaging() *v1.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type ListExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exports  []*LogsExport `protobuf:"bytes,1,rep,name=exports,proto3" json:"exports,omitempty"`
	NextPage *v1.NextPage  `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *ListExportResponse) Reset() {
	*x = ListExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExportResponse) ProtoMessage() {}

func (x *ListExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExportResponse.ProtoReflect.Descriptor instead.
func (*ListExportResponse) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListExportResponse) GetExports() []*LogsExport {
	if x != nil {
		return x.Exports
	}
	return nil
}

func (x *ListExportResponse) GetNextPage() *v1.NextPage {
	if x != nil {
		return x.NextPage
	}
	return nil
}

type GetExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetExportRequest) Reset() {
	*x = GetExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExportRequest) ProtoMessage() {}

func (x *GetExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExportRequest.ProtoReflect.Descriptor instead.
func (*GetExportRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetExportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId   string       `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Sources     []*LogSource `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
	Target      *LogsTarget  `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *CreateExportRequest) Reset() {
	*x = CreateExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExportRequest) ProtoMessage() {}

func (x *CreateExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExportRequest.ProtoReflect.Descriptor instead.
func (*CreateExportRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateExportRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateExportRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateExportRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateExportRequest) GetSources() []*LogSource {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *CreateExportRequest) GetTarget() *LogsTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

type UpdateExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Sources     []*LogSource `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
	Target      *LogsTarget  `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *UpdateExportRequest) Reset() {
	*x = UpdateExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExportRequest) ProtoMessage() {}

func (x *UpdateExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExportRequest.ProtoReflect.Descriptor instead.
func (*UpdateExportRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateExportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateExportRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateExportRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateExportRequest) GetSources() []*LogSource {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *UpdateExportRequest) GetTarget() *LogsTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

type DeleteExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteExportRequest) Reset() {
	*x = DeleteExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExportRequest) ProtoMessage() {}

func (x *DeleteExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExportRequest.ProtoReflect.Descriptor instead.
func (*DeleteExportRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteExportRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_doublecloud_logs_v1_log_export_service_proto protoreflect.FileDescriptor

var file_doublecloud_logs_v1_log_export_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x22, 0x86, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x01,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0xce, 0x01,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x25,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa7, 0x03, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x25, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x57, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x26, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_logs_v1_log_export_service_proto_rawDescOnce sync.Once
	file_doublecloud_logs_v1_log_export_service_proto_rawDescData = file_doublecloud_logs_v1_log_export_service_proto_rawDesc
)

func file_doublecloud_logs_v1_log_export_service_proto_rawDescGZIP() []byte {
	file_doublecloud_logs_v1_log_export_service_proto_rawDescOnce.Do(func() {
		file_doublecloud_logs_v1_log_export_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_logs_v1_log_export_service_proto_rawDescData)
	})
	return file_doublecloud_logs_v1_log_export_service_proto_rawDescData
}

var file_doublecloud_logs_v1_log_export_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_doublecloud_logs_v1_log_export_service_proto_goTypes = []interface{}{
	(*ListExportRequest)(nil),   // 0: doublecloud.logs.v1.ListExportRequest
	(*ListExportResponse)(nil),  // 1: doublecloud.logs.v1.ListExportResponse
	(*GetExportRequest)(nil),    // 2: doublecloud.logs.v1.GetExportRequest
	(*CreateExportRequest)(nil), // 3: doublecloud.logs.v1.CreateExportRequest
	(*UpdateExportRequest)(nil), // 4: doublecloud.logs.v1.UpdateExportRequest
	(*DeleteExportRequest)(nil), // 5: doublecloud.logs.v1.DeleteExportRequest
	(*v1.Paging)(nil),           // 6: doublecloud.v1.Paging
	(*LogsExport)(nil),          // 7: doublecloud.logs.v1.LogsExport
	(*v1.NextPage)(nil),         // 8: doublecloud.v1.NextPage
	(*LogSource)(nil),           // 9: doublecloud.logs.v1.LogSource
	(*LogsTarget)(nil),          // 10: doublecloud.logs.v1.LogsTarget
	(*v1.Operation)(nil),        // 11: doublecloud.v1.Operation
}
var file_doublecloud_logs_v1_log_export_service_proto_depIdxs = []int32{
	6,  // 0: doublecloud.logs.v1.ListExportRequest.paging:type_name -> doublecloud.v1.Paging
	7,  // 1: doublecloud.logs.v1.ListExportResponse.exports:type_name -> doublecloud.logs.v1.LogsExport
	8,  // 2: doublecloud.logs.v1.ListExportResponse.next_page:type_name -> doublecloud.v1.NextPage
	9,  // 3: doublecloud.logs.v1.CreateExportRequest.sources:type_name -> doublecloud.logs.v1.LogSource
	10, // 4: doublecloud.logs.v1.CreateExportRequest.target:type_name -> doublecloud.logs.v1.LogsTarget
	9,  // 5: doublecloud.logs.v1.UpdateExportRequest.sources:type_name -> doublecloud.logs.v1.LogSource
	10, // 6: doublecloud.logs.v1.UpdateExportRequest.target:type_name -> doublecloud.logs.v1.LogsTarget
	2,  // 7: doublecloud.logs.v1.LogExportService.Get:input_type -> doublecloud.logs.v1.GetExportRequest
	0,  // 8: doublecloud.logs.v1.LogExportService.List:input_type -> doublecloud.logs.v1.ListExportRequest
	3,  // 9: doublecloud.logs.v1.LogExportService.Create:input_type -> doublecloud.logs.v1.CreateExportRequest
	4,  // 10: doublecloud.logs.v1.LogExportService.Update:input_type -> doublecloud.logs.v1.UpdateExportRequest
	5,  // 11: doublecloud.logs.v1.LogExportService.Delete:input_type -> doublecloud.logs.v1.DeleteExportRequest
	7,  // 12: doublecloud.logs.v1.LogExportService.Get:output_type -> doublecloud.logs.v1.LogsExport
	1,  // 13: doublecloud.logs.v1.LogExportService.List:output_type -> doublecloud.logs.v1.ListExportResponse
	11, // 14: doublecloud.logs.v1.LogExportService.Create:output_type -> doublecloud.v1.Operation
	11, // 15: doublecloud.logs.v1.LogExportService.Update:output_type -> doublecloud.v1.Operation
	11, // 16: doublecloud.logs.v1.LogExportService.Delete:output_type -> doublecloud.v1.Operation
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_doublecloud_logs_v1_log_export_service_proto_init() }
func file_doublecloud_logs_v1_log_export_service_proto_init() {
	if File_doublecloud_logs_v1_log_export_service_proto != nil {
		return
	}
	file_doublecloud_logs_v1_log_export_proto_init()
	file_doublecloud_logs_v1_log_source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_logs_v1_log_export_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExportRequest); i {
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
		file_doublecloud_logs_v1_log_export_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExportResponse); i {
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
		file_doublecloud_logs_v1_log_export_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExportRequest); i {
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
		file_doublecloud_logs_v1_log_export_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExportRequest); i {
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
		file_doublecloud_logs_v1_log_export_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExportRequest); i {
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
		file_doublecloud_logs_v1_log_export_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExportRequest); i {
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
			RawDescriptor: file_doublecloud_logs_v1_log_export_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doublecloud_logs_v1_log_export_service_proto_goTypes,
		DependencyIndexes: file_doublecloud_logs_v1_log_export_service_proto_depIdxs,
		MessageInfos:      file_doublecloud_logs_v1_log_export_service_proto_msgTypes,
	}.Build()
	File_doublecloud_logs_v1_log_export_service_proto = out.File
	file_doublecloud_logs_v1_log_export_service_proto_rawDesc = nil
	file_doublecloud_logs_v1_log_export_service_proto_goTypes = nil
	file_doublecloud_logs_v1_log_export_service_proto_depIdxs = nil
}
