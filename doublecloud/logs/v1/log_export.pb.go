// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: doublecloud/logs/v1/log_export.proto

package logs

import (
	_ "github.com/doublecloud/go-genproto/doublecloud/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogExportStatus int32

const (
	LogExportStatus_LOG_EXPORT_STATUS_UNSPECIFIED LogExportStatus = 0
	LogExportStatus_LOG_EXPORT_STATUS_PENDING     LogExportStatus = 1
	LogExportStatus_LOG_EXPORT_STATUS_RUNNING     LogExportStatus = 2
	LogExportStatus_LOG_EXPORT_STATUS_STOPPED     LogExportStatus = 3
	LogExportStatus_LOG_EXPORT_STATUS_FAILED      LogExportStatus = 4
)

// Enum value maps for LogExportStatus.
var (
	LogExportStatus_name = map[int32]string{
		0: "LOG_EXPORT_STATUS_UNSPECIFIED",
		1: "LOG_EXPORT_STATUS_PENDING",
		2: "LOG_EXPORT_STATUS_RUNNING",
		3: "LOG_EXPORT_STATUS_STOPPED",
		4: "LOG_EXPORT_STATUS_FAILED",
	}
	LogExportStatus_value = map[string]int32{
		"LOG_EXPORT_STATUS_UNSPECIFIED": 0,
		"LOG_EXPORT_STATUS_PENDING":     1,
		"LOG_EXPORT_STATUS_RUNNING":     2,
		"LOG_EXPORT_STATUS_STOPPED":     3,
		"LOG_EXPORT_STATUS_FAILED":      4,
	}
)

func (x LogExportStatus) Enum() *LogExportStatus {
	p := new(LogExportStatus)
	*p = x
	return p
}

func (x LogExportStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogExportStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_logs_v1_log_export_proto_enumTypes[0].Descriptor()
}

func (LogExportStatus) Type() protoreflect.EnumType {
	return &file_doublecloud_logs_v1_log_export_proto_enumTypes[0]
}

func (x LogExportStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogExportStatus.Descriptor instead.
func (LogExportStatus) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{0}
}

// see:
// https://a.yandex-team.ru/arcadia/vendor/github.com/DataDog/datadog-api-client-go/v2/api/datadog/configuration.go#L128
type LogsTargetDatadog_DatadogHost int32

const (
	LogsTargetDatadog_DATADOG_HOST_UNSPECIFIED       LogsTargetDatadog_DatadogHost = 0
	LogsTargetDatadog_DATADOG_HOST_DATADOGHQ         LogsTargetDatadog_DatadogHost = 1
	LogsTargetDatadog_DATADOG_HOST_US3_DATADOGHQ     LogsTargetDatadog_DatadogHost = 2
	LogsTargetDatadog_DATADOG_HOST_US5_DATADOGHQ_COM LogsTargetDatadog_DatadogHost = 3
	LogsTargetDatadog_DATADOG_HOST_AP1_DATADOGHQ_COM LogsTargetDatadog_DatadogHost = 4
	LogsTargetDatadog_DATADOG_HOST_DATADOGHQ_EU      LogsTargetDatadog_DatadogHost = 5
	LogsTargetDatadog_DATADOG_HOST_DDOGGOV_COM       LogsTargetDatadog_DatadogHost = 6
)

// Enum value maps for LogsTargetDatadog_DatadogHost.
var (
	LogsTargetDatadog_DatadogHost_name = map[int32]string{
		0: "DATADOG_HOST_UNSPECIFIED",
		1: "DATADOG_HOST_DATADOGHQ",
		2: "DATADOG_HOST_US3_DATADOGHQ",
		3: "DATADOG_HOST_US5_DATADOGHQ_COM",
		4: "DATADOG_HOST_AP1_DATADOGHQ_COM",
		5: "DATADOG_HOST_DATADOGHQ_EU",
		6: "DATADOG_HOST_DDOGGOV_COM",
	}
	LogsTargetDatadog_DatadogHost_value = map[string]int32{
		"DATADOG_HOST_UNSPECIFIED":       0,
		"DATADOG_HOST_DATADOGHQ":         1,
		"DATADOG_HOST_US3_DATADOGHQ":     2,
		"DATADOG_HOST_US5_DATADOGHQ_COM": 3,
		"DATADOG_HOST_AP1_DATADOGHQ_COM": 4,
		"DATADOG_HOST_DATADOGHQ_EU":      5,
		"DATADOG_HOST_DDOGGOV_COM":       6,
	}
)

func (x LogsTargetDatadog_DatadogHost) Enum() *LogsTargetDatadog_DatadogHost {
	p := new(LogsTargetDatadog_DatadogHost)
	*p = x
	return p
}

func (x LogsTargetDatadog_DatadogHost) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogsTargetDatadog_DatadogHost) Descriptor() protoreflect.EnumDescriptor {
	return file_doublecloud_logs_v1_log_export_proto_enumTypes[1].Descriptor()
}

func (LogsTargetDatadog_DatadogHost) Type() protoreflect.EnumType {
	return &file_doublecloud_logs_v1_log_export_proto_enumTypes[1]
}

func (x LogsTargetDatadog_DatadogHost) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogsTargetDatadog_DatadogHost.Descriptor instead.
func (LogsTargetDatadog_DatadogHost) EnumDescriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{2, 0}
}

type LogsExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId    string          `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Name         string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description  string          `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Sources      []*LogSource    `protobuf:"bytes,5,rep,name=sources,proto3" json:"sources,omitempty"`
	Target       *LogsTarget     `protobuf:"bytes,6,opt,name=target,proto3" json:"target,omitempty"`
	Status       LogExportStatus `protobuf:"varint,7,opt,name=status,proto3,enum=doublecloud.logs.v1.LogExportStatus" json:"status,omitempty"`
	StatusReason string          `protobuf:"bytes,8,opt,name=status_reason,json=statusReason,proto3" json:"status_reason,omitempty"`
}

func (x *LogsExport) Reset() {
	*x = LogsExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsExport) ProtoMessage() {}

func (x *LogsExport) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsExport.ProtoReflect.Descriptor instead.
func (*LogsExport) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{0}
}

func (x *LogsExport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LogsExport) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *LogsExport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogsExport) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LogsExport) GetSources() []*LogSource {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *LogsExport) GetTarget() *LogsTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *LogsExport) GetStatus() LogExportStatus {
	if x != nil {
		return x.Status
	}
	return LogExportStatus_LOG_EXPORT_STATUS_UNSPECIFIED
}

func (x *LogsExport) GetStatusReason() string {
	if x != nil {
		return x.StatusReason
	}
	return ""
}

type LogsTargetS3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket             string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	BucketLayout       string `protobuf:"bytes,2,opt,name=bucket_layout,json=bucketLayout,proto3" json:"bucket_layout,omitempty"`
	AwsAccessKeyId     string `protobuf:"bytes,3,opt,name=aws_access_key_id,json=awsAccessKeyId,proto3" json:"aws_access_key_id,omitempty"`
	AwsSecretAccessKey string `protobuf:"bytes,4,opt,name=aws_secret_access_key,json=awsSecretAccessKey,proto3" json:"aws_secret_access_key,omitempty"`
	Region             string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Endpoint           string `protobuf:"bytes,6,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	DisableSsl         bool   `protobuf:"varint,7,opt,name=disable_ssl,json=disableSsl,proto3" json:"disable_ssl,omitempty"`
	SkipVerifySslCert  bool   `protobuf:"varint,8,opt,name=skip_verify_ssl_cert,json=skipVerifySslCert,proto3" json:"skip_verify_ssl_cert,omitempty"`
}

func (x *LogsTargetS3) Reset() {
	*x = LogsTargetS3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTargetS3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTargetS3) ProtoMessage() {}

func (x *LogsTargetS3) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTargetS3.ProtoReflect.Descriptor instead.
func (*LogsTargetS3) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{1}
}

func (x *LogsTargetS3) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *LogsTargetS3) GetBucketLayout() string {
	if x != nil {
		return x.BucketLayout
	}
	return ""
}

func (x *LogsTargetS3) GetAwsAccessKeyId() string {
	if x != nil {
		return x.AwsAccessKeyId
	}
	return ""
}

func (x *LogsTargetS3) GetAwsSecretAccessKey() string {
	if x != nil {
		return x.AwsSecretAccessKey
	}
	return ""
}

func (x *LogsTargetS3) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *LogsTargetS3) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *LogsTargetS3) GetDisableSsl() bool {
	if x != nil {
		return x.DisableSsl
	}
	return false
}

func (x *LogsTargetS3) GetSkipVerifySslCert() bool {
	if x != nil {
		return x.SkipVerifySslCert
	}
	return false
}

type LogsTargetDatadog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey      string                        `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	DatadogHost LogsTargetDatadog_DatadogHost `protobuf:"varint,2,opt,name=datadog_host,json=datadogHost,proto3,enum=doublecloud.logs.v1.LogsTargetDatadog_DatadogHost" json:"datadog_host,omitempty"`
}

func (x *LogsTargetDatadog) Reset() {
	*x = LogsTargetDatadog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTargetDatadog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTargetDatadog) ProtoMessage() {}

func (x *LogsTargetDatadog) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTargetDatadog.ProtoReflect.Descriptor instead.
func (*LogsTargetDatadog) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{2}
}

func (x *LogsTargetDatadog) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *LogsTargetDatadog) GetDatadogHost() LogsTargetDatadog_DatadogHost {
	if x != nil {
		return x.DatadogHost
	}
	return LogsTargetDatadog_DATADOG_HOST_UNSPECIFIED
}

type LogsTargetCoralogix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogsTargetCoralogix) Reset() {
	*x = LogsTargetCoralogix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTargetCoralogix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTargetCoralogix) ProtoMessage() {}

func (x *LogsTargetCoralogix) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTargetCoralogix.ProtoReflect.Descriptor instead.
func (*LogsTargetCoralogix) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{3}
}

func (x *LogsTargetCoralogix) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *LogsTargetCoralogix) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogsTargetDatadogCompatible struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey      string `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	DatadogHost string `protobuf:"bytes,2,opt,name=datadog_host,json=datadogHost,proto3" json:"datadog_host,omitempty"`
}

func (x *LogsTargetDatadogCompatible) Reset() {
	*x = LogsTargetDatadogCompatible{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTargetDatadogCompatible) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTargetDatadogCompatible) ProtoMessage() {}

func (x *LogsTargetDatadogCompatible) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTargetDatadogCompatible.ProtoReflect.Descriptor instead.
func (*LogsTargetDatadogCompatible) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{4}
}

func (x *LogsTargetDatadogCompatible) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *LogsTargetDatadogCompatible) GetDatadogHost() string {
	if x != nil {
		return x.DatadogHost
	}
	return ""
}

type LogsTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*LogsTarget_S3
	//	*LogsTarget_Datadog
	//	*LogsTarget_Coralogix
	//	*LogsTarget_DatadogCompatible
	Target isLogsTarget_Target `protobuf_oneof:"target"`
}

func (x *LogsTarget) Reset() {
	*x = LogsTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTarget) ProtoMessage() {}

func (x *LogsTarget) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_logs_v1_log_export_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTarget.ProtoReflect.Descriptor instead.
func (*LogsTarget) Descriptor() ([]byte, []int) {
	return file_doublecloud_logs_v1_log_export_proto_rawDescGZIP(), []int{5}
}

func (m *LogsTarget) GetTarget() isLogsTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *LogsTarget) GetS3() *LogsTargetS3 {
	if x, ok := x.GetTarget().(*LogsTarget_S3); ok {
		return x.S3
	}
	return nil
}

func (x *LogsTarget) GetDatadog() *LogsTargetDatadog {
	if x, ok := x.GetTarget().(*LogsTarget_Datadog); ok {
		return x.Datadog
	}
	return nil
}

func (x *LogsTarget) GetCoralogix() *LogsTargetCoralogix {
	if x, ok := x.GetTarget().(*LogsTarget_Coralogix); ok {
		return x.Coralogix
	}
	return nil
}

func (x *LogsTarget) GetDatadogCompatible() *LogsTargetDatadogCompatible {
	if x, ok := x.GetTarget().(*LogsTarget_DatadogCompatible); ok {
		return x.DatadogCompatible
	}
	return nil
}

type isLogsTarget_Target interface {
	isLogsTarget_Target()
}

type LogsTarget_S3 struct {
	S3 *LogsTargetS3 `protobuf:"bytes,1,opt,name=s3,proto3,oneof"`
}

type LogsTarget_Datadog struct {
	Datadog *LogsTargetDatadog `protobuf:"bytes,2,opt,name=datadog,proto3,oneof"`
}

type LogsTarget_Coralogix struct {
	Coralogix *LogsTargetCoralogix `protobuf:"bytes,3,opt,name=coralogix,proto3,oneof"`
}

type LogsTarget_DatadogCompatible struct {
	DatadogCompatible *LogsTargetDatadogCompatible `protobuf:"bytes,4,opt,name=datadog_compatible,json=datadogCompatible,proto3,oneof"`
}

func (*LogsTarget_S3) isLogsTarget_Target() {}

func (*LogsTarget_Datadog) isLogsTarget_Target() {}

func (*LogsTarget_Coralogix) isLogsTarget_Target() {}

func (*LogsTarget_DatadogCompatible) isLogsTarget_Target() {}

var File_doublecloud_logs_v1_log_export_proto protoreflect.FileDescriptor

var file_doublecloud_logs_v1_log_export_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc7, 0x02, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xaf, 0x02, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x61, 0x77, 0x73, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x61, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x73, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x73, 0x6c, 0x12, 0x2f, 0x0a, 0x14, 0x73, 0x6b,
	0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x73, 0x6c, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x6b, 0x69, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x53, 0x73, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x22, 0xf2, 0x02, 0x0a, 0x11,
	0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x55, 0x0a, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x48, 0x6f, 0x73,
	0x74, 0x22, 0xec, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x5f, 0x48, 0x4f, 0x53,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x48, 0x51, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x44,
	0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x55, 0x53, 0x33, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x48, 0x51, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x44,
	0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x55, 0x53, 0x35, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x48, 0x51, 0x5f, 0x43, 0x4f, 0x4d, 0x10, 0x03, 0x12,
	0x22, 0x0a, 0x1e, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x41, 0x50, 0x31, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x48, 0x51, 0x5f, 0x43, 0x4f,
	0x4d, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x5f, 0x48,
	0x4f, 0x53, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x48, 0x51, 0x5f, 0x45, 0x55,
	0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x41, 0x54, 0x41, 0x44, 0x4f, 0x47, 0x5f, 0x48, 0x4f,
	0x53, 0x54, 0x5f, 0x44, 0x44, 0x4f, 0x47, 0x47, 0x4f, 0x56, 0x5f, 0x43, 0x4f, 0x4d, 0x10, 0x06,
	0x22, 0x43, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x59, 0x0a, 0x1b, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74,
	0x69, 0x62, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x48, 0x6f, 0x73, 0x74,
	0x22, 0xbc, 0x02, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x33, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x33, 0x48, 0x00,
	0x52, 0x02, 0x73, 0x33, 0x12, 0x42, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x48, 0x00, 0x52,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x12, 0x48, 0x0a, 0x09, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x12, 0x61, 0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65,
	0x48, 0x00, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x69, 0x62, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2a,
	0xaf, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x4c, 0x4f, 0x47, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x4f, 0x47, 0x5f, 0x45, 0x58,
	0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x4f, 0x47, 0x5f, 0x45, 0x58, 0x50,
	0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x4f, 0x47, 0x5f, 0x45, 0x58, 0x50, 0x4f,
	0x52, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x4f, 0x47, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x04, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doublecloud_logs_v1_log_export_proto_rawDescOnce sync.Once
	file_doublecloud_logs_v1_log_export_proto_rawDescData = file_doublecloud_logs_v1_log_export_proto_rawDesc
)

func file_doublecloud_logs_v1_log_export_proto_rawDescGZIP() []byte {
	file_doublecloud_logs_v1_log_export_proto_rawDescOnce.Do(func() {
		file_doublecloud_logs_v1_log_export_proto_rawDescData = protoimpl.X.CompressGZIP(file_doublecloud_logs_v1_log_export_proto_rawDescData)
	})
	return file_doublecloud_logs_v1_log_export_proto_rawDescData
}

var file_doublecloud_logs_v1_log_export_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_doublecloud_logs_v1_log_export_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_doublecloud_logs_v1_log_export_proto_goTypes = []any{
	(LogExportStatus)(0),                // 0: doublecloud.logs.v1.LogExportStatus
	(LogsTargetDatadog_DatadogHost)(0),  // 1: doublecloud.logs.v1.LogsTargetDatadog.DatadogHost
	(*LogsExport)(nil),                  // 2: doublecloud.logs.v1.LogsExport
	(*LogsTargetS3)(nil),                // 3: doublecloud.logs.v1.LogsTargetS3
	(*LogsTargetDatadog)(nil),           // 4: doublecloud.logs.v1.LogsTargetDatadog
	(*LogsTargetCoralogix)(nil),         // 5: doublecloud.logs.v1.LogsTargetCoralogix
	(*LogsTargetDatadogCompatible)(nil), // 6: doublecloud.logs.v1.LogsTargetDatadogCompatible
	(*LogsTarget)(nil),                  // 7: doublecloud.logs.v1.LogsTarget
	(*LogSource)(nil),                   // 8: doublecloud.logs.v1.LogSource
}
var file_doublecloud_logs_v1_log_export_proto_depIdxs = []int32{
	8, // 0: doublecloud.logs.v1.LogsExport.sources:type_name -> doublecloud.logs.v1.LogSource
	7, // 1: doublecloud.logs.v1.LogsExport.target:type_name -> doublecloud.logs.v1.LogsTarget
	0, // 2: doublecloud.logs.v1.LogsExport.status:type_name -> doublecloud.logs.v1.LogExportStatus
	1, // 3: doublecloud.logs.v1.LogsTargetDatadog.datadog_host:type_name -> doublecloud.logs.v1.LogsTargetDatadog.DatadogHost
	3, // 4: doublecloud.logs.v1.LogsTarget.s3:type_name -> doublecloud.logs.v1.LogsTargetS3
	4, // 5: doublecloud.logs.v1.LogsTarget.datadog:type_name -> doublecloud.logs.v1.LogsTargetDatadog
	5, // 6: doublecloud.logs.v1.LogsTarget.coralogix:type_name -> doublecloud.logs.v1.LogsTargetCoralogix
	6, // 7: doublecloud.logs.v1.LogsTarget.datadog_compatible:type_name -> doublecloud.logs.v1.LogsTargetDatadogCompatible
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_doublecloud_logs_v1_log_export_proto_init() }
func file_doublecloud_logs_v1_log_export_proto_init() {
	if File_doublecloud_logs_v1_log_export_proto != nil {
		return
	}
	file_doublecloud_logs_v1_log_source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_doublecloud_logs_v1_log_export_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LogsExport); i {
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
		file_doublecloud_logs_v1_log_export_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LogsTargetS3); i {
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
		file_doublecloud_logs_v1_log_export_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LogsTargetDatadog); i {
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
		file_doublecloud_logs_v1_log_export_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LogsTargetCoralogix); i {
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
		file_doublecloud_logs_v1_log_export_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LogsTargetDatadogCompatible); i {
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
		file_doublecloud_logs_v1_log_export_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LogsTarget); i {
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
	file_doublecloud_logs_v1_log_export_proto_msgTypes[5].OneofWrappers = []any{
		(*LogsTarget_S3)(nil),
		(*LogsTarget_Datadog)(nil),
		(*LogsTarget_Coralogix)(nil),
		(*LogsTarget_DatadogCompatible)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_doublecloud_logs_v1_log_export_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_logs_v1_log_export_proto_goTypes,
		DependencyIndexes: file_doublecloud_logs_v1_log_export_proto_depIdxs,
		EnumInfos:         file_doublecloud_logs_v1_log_export_proto_enumTypes,
		MessageInfos:      file_doublecloud_logs_v1_log_export_proto_msgTypes,
	}.Build()
	File_doublecloud_logs_v1_log_export_proto = out.File
	file_doublecloud_logs_v1_log_export_proto_rawDesc = nil
	file_doublecloud_logs_v1_log_export_proto_goTypes = nil
	file_doublecloud_logs_v1_log_export_proto_depIdxs = nil
}
