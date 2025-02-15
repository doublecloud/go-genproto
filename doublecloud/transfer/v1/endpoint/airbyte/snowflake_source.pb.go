// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: doublecloud/transfer/v1/endpoint/airbyte/snowflake_source.proto

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

type SnowflakeSource struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Host          string                       `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Role          string                       `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Warehouse     string                       `protobuf:"bytes,3,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	Database      string                       `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	Schema        string                       `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
	JdbcUrlParams string                       `protobuf:"bytes,6,opt,name=jdbc_url_params,json=jdbcUrlParams,proto3" json:"jdbc_url_params,omitempty"`
	Credentials   *SnowflakeSource_Credentials `protobuf:"bytes,7,opt,name=credentials,proto3" json:"credentials,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SnowflakeSource) Reset() {
	*x = SnowflakeSource{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnowflakeSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnowflakeSource) ProtoMessage() {}

func (x *SnowflakeSource) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnowflakeSource.ProtoReflect.Descriptor instead.
func (*SnowflakeSource) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescGZIP(), []int{0}
}

func (x *SnowflakeSource) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SnowflakeSource) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *SnowflakeSource) GetWarehouse() string {
	if x != nil {
		return x.Warehouse
	}
	return ""
}

func (x *SnowflakeSource) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *SnowflakeSource) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *SnowflakeSource) GetJdbcUrlParams() string {
	if x != nil {
		return x.JdbcUrlParams
	}
	return ""
}

func (x *SnowflakeSource) GetCredentials() *SnowflakeSource_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type SnowflakeSource_Credentials struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Credentials:
	//
	//	*SnowflakeSource_Credentials_Oauth
	//	*SnowflakeSource_Credentials_BasicAuth_
	Credentials   isSnowflakeSource_Credentials_Credentials `protobuf_oneof:"credentials"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SnowflakeSource_Credentials) Reset() {
	*x = SnowflakeSource_Credentials{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnowflakeSource_Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnowflakeSource_Credentials) ProtoMessage() {}

func (x *SnowflakeSource_Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnowflakeSource_Credentials.ProtoReflect.Descriptor instead.
func (*SnowflakeSource_Credentials) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SnowflakeSource_Credentials) GetCredentials() isSnowflakeSource_Credentials_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *SnowflakeSource_Credentials) GetOauth() *SnowflakeSource_Credentials_OAuth {
	if x != nil {
		if x, ok := x.Credentials.(*SnowflakeSource_Credentials_Oauth); ok {
			return x.Oauth
		}
	}
	return nil
}

func (x *SnowflakeSource_Credentials) GetBasicAuth() *SnowflakeSource_Credentials_BasicAuth {
	if x != nil {
		if x, ok := x.Credentials.(*SnowflakeSource_Credentials_BasicAuth_); ok {
			return x.BasicAuth
		}
	}
	return nil
}

type isSnowflakeSource_Credentials_Credentials interface {
	isSnowflakeSource_Credentials_Credentials()
}

type SnowflakeSource_Credentials_Oauth struct {
	Oauth *SnowflakeSource_Credentials_OAuth `protobuf:"bytes,1,opt,name=oauth,proto3,oneof"`
}

type SnowflakeSource_Credentials_BasicAuth_ struct {
	BasicAuth *SnowflakeSource_Credentials_BasicAuth `protobuf:"bytes,2,opt,name=basic_auth,json=basicAuth,proto3,oneof"`
}

func (*SnowflakeSource_Credentials_Oauth) isSnowflakeSource_Credentials_Credentials() {}

func (*SnowflakeSource_Credentials_BasicAuth_) isSnowflakeSource_Credentials_Credentials() {}

type SnowflakeSource_Credentials_OAuth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret  string                 `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AccessToken   string                 `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SnowflakeSource_Credentials_OAuth) Reset() {
	*x = SnowflakeSource_Credentials_OAuth{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnowflakeSource_Credentials_OAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnowflakeSource_Credentials_OAuth) ProtoMessage() {}

func (x *SnowflakeSource_Credentials_OAuth) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnowflakeSource_Credentials_OAuth.ProtoReflect.Descriptor instead.
func (*SnowflakeSource_Credentials_OAuth) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SnowflakeSource_Credentials_OAuth) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SnowflakeSource_Credentials_OAuth) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *SnowflakeSource_Credentials_OAuth) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SnowflakeSource_Credentials_OAuth) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type SnowflakeSource_Credentials_BasicAuth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SnowflakeSource_Credentials_BasicAuth) Reset() {
	*x = SnowflakeSource_Credentials_BasicAuth{}
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnowflakeSource_Credentials_BasicAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnowflakeSource_Credentials_BasicAuth) ProtoMessage() {}

func (x *SnowflakeSource_Credentials_BasicAuth) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnowflakeSource_Credentials_BasicAuth.ProtoReflect.Descriptor instead.
func (*SnowflakeSource_Credentials_BasicAuth) Descriptor() ([]byte, []int) {
	return file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SnowflakeSource_Credentials_BasicAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SnowflakeSource_Credentials_BasicAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto protoreflect.FileDescriptor

var file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDesc = string([]byte{
	0x0a, 0x3f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66,
	0x6c, 0x61, 0x6b, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x22, 0xf7, 0x05, 0x0a, 0x0f,
	0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x6a, 0x64, 0x62,
	0x63, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6a, 0x64, 0x62, 0x63, 0x55, 0x72, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x67, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74,
	0x65, 0x2e, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0xd8, 0x03, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x63, 0x0a, 0x05, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72,
	0x62, 0x79, 0x74, 0x65, 0x2e, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x12,
	0x70, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x69, 0x72, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x53,
	0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x09, 0x62, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74,
	0x68, 0x1a, 0x97, 0x01, 0x0a, 0x05, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x1a, 0x49, 0x0a, 0x09, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x69, 0x72,
	0x62, 0x79, 0x74, 0x65, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x69,
	0x72, 0x62, 0x79, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescOnce sync.Once
	file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescData []byte
)

func file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescGZIP() []byte {
	file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescOnce.Do(func() {
		file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDesc)))
	})
	return file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDescData
}

var file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_goTypes = []any{
	(*SnowflakeSource)(nil),                       // 0: doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource
	(*SnowflakeSource_Credentials)(nil),           // 1: doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials
	(*SnowflakeSource_Credentials_OAuth)(nil),     // 2: doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials.OAuth
	(*SnowflakeSource_Credentials_BasicAuth)(nil), // 3: doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials.BasicAuth
}
var file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_depIdxs = []int32{
	1, // 0: doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.credentials:type_name -> doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials
	2, // 1: doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials.oauth:type_name -> doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials.OAuth
	3, // 2: doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials.basic_auth:type_name -> doublecloud.transfer.v1.endpoint.airbyte.SnowflakeSource.Credentials.BasicAuth
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_init() }
func file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_init() {
	if File_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto != nil {
		return
	}
	file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes[1].OneofWrappers = []any{
		(*SnowflakeSource_Credentials_Oauth)(nil),
		(*SnowflakeSource_Credentials_BasicAuth_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDesc), len(file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_goTypes,
		DependencyIndexes: file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_depIdxs,
		MessageInfos:      file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_msgTypes,
	}.Build()
	File_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto = out.File
	file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_goTypes = nil
	file_doublecloud_transfer_v1_endpoint_airbyte_snowflake_source_proto_depIdxs = nil
}
