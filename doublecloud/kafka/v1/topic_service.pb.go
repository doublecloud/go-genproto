// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: doublecloud/kafka/v1/topic_service.proto

package kafka

import (
	v1 "github.com/doublecloud/go-genproto/doublecloud/v1"
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

type GetTopicRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Apache Kafka cluster that the topic belongs to.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the Apache Kafka Topic resource to return.
	// To get the name of the topic use a [TopicService.List] request.
	TopicName     string `protobuf:"bytes,2,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTopicRequest) Reset() {
	*x = GetTopicRequest{}
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicRequest) ProtoMessage() {}

func (x *GetTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicRequest.ProtoReflect.Descriptor instead.
func (*GetTopicRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_topic_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTopicRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetTopicRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

type ListTopicsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Apache Kafka cluster to list topics in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Paging information of the request
	Paging        *v1.Paging `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTopicsRequest) Reset() {
	*x = ListTopicsRequest{}
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTopicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopicsRequest) ProtoMessage() {}

func (x *ListTopicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopicsRequest.ProtoReflect.Descriptor instead.
func (*ListTopicsRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_topic_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListTopicsRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ListTopicsRequest) GetPaging() *v1.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type ListTopicsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of Apache Kafka Topic resources.
	Topics []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	// Pagination information of the response
	NextPage      *v1.NextPage `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTopicsResponse) Reset() {
	*x = ListTopicsResponse{}
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTopicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopicsResponse) ProtoMessage() {}

func (x *ListTopicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopicsResponse.ProtoReflect.Descriptor instead.
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_topic_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTopicsResponse) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *ListTopicsResponse) GetNextPage() *v1.NextPage {
	if x != nil {
		return x.NextPage
	}
	return nil
}

type CreateTopicRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the Apache Kafka cluster to create a topic in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Configuration of the topic to create.
	TopicSpec     *TopicSpec `protobuf:"bytes,2,opt,name=topic_spec,json=topicSpec,proto3" json:"topic_spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTopicRequest) Reset() {
	*x = CreateTopicRequest{}
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicRequest) ProtoMessage() {}

func (x *CreateTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicRequest.ProtoReflect.Descriptor instead.
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_topic_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTopicRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateTopicRequest) GetTopicSpec() *TopicSpec {
	if x != nil {
		return x.TopicSpec
	}
	return nil
}

type UpdateTopicRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the Apache Kafka cluster to update a topic in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Name of the topic to update.
	// To get the name of the topic use a [TopicService.List] request.
	TopicName string `protobuf:"bytes,2,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// Required. Configuration of the topic to create.
	TopicSpec     *TopicSpec `protobuf:"bytes,3,opt,name=topic_spec,json=topicSpec,proto3" json:"topic_spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTopicRequest) Reset() {
	*x = UpdateTopicRequest{}
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTopicRequest) ProtoMessage() {}

func (x *UpdateTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTopicRequest.ProtoReflect.Descriptor instead.
func (*UpdateTopicRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_topic_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTopicRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UpdateTopicRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *UpdateTopicRequest) GetTopicSpec() *TopicSpec {
	if x != nil {
		return x.TopicSpec
	}
	return nil
}

type DeleteTopicRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the Apache Kafka cluster to delete a topic in.
	// To get the cluster ID, use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Name of the topic to delete.
	// To get the name of the topic, use a [TopicService.List] request.
	TopicName     string `protobuf:"bytes,2,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTopicRequest) Reset() {
	*x = DeleteTopicRequest{}
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTopicRequest) ProtoMessage() {}

func (x *DeleteTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doublecloud_kafka_v1_topic_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTopicRequest.ProtoReflect.Descriptor instead.
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) {
	return file_doublecloud_kafka_v1_topic_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTopicRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteTopicRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

var File_doublecloud_kafka_v1_topic_service_proto protoreflect.FileDescriptor

var file_doublecloud_kafka_v1_topic_service_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x62, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x12, 0x35, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x73, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0a,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x22, 0x92, 0x01, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x70, 0x65,
	0x63, 0x22, 0x52, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xa1, 0x03, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x25, 0x2e,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x59, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_doublecloud_kafka_v1_topic_service_proto_rawDescOnce sync.Once
	file_doublecloud_kafka_v1_topic_service_proto_rawDescData []byte
)

func file_doublecloud_kafka_v1_topic_service_proto_rawDescGZIP() []byte {
	file_doublecloud_kafka_v1_topic_service_proto_rawDescOnce.Do(func() {
		file_doublecloud_kafka_v1_topic_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doublecloud_kafka_v1_topic_service_proto_rawDesc), len(file_doublecloud_kafka_v1_topic_service_proto_rawDesc)))
	})
	return file_doublecloud_kafka_v1_topic_service_proto_rawDescData
}

var file_doublecloud_kafka_v1_topic_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_doublecloud_kafka_v1_topic_service_proto_goTypes = []any{
	(*GetTopicRequest)(nil),    // 0: doublecloud.kafka.v1.GetTopicRequest
	(*ListTopicsRequest)(nil),  // 1: doublecloud.kafka.v1.ListTopicsRequest
	(*ListTopicsResponse)(nil), // 2: doublecloud.kafka.v1.ListTopicsResponse
	(*CreateTopicRequest)(nil), // 3: doublecloud.kafka.v1.CreateTopicRequest
	(*UpdateTopicRequest)(nil), // 4: doublecloud.kafka.v1.UpdateTopicRequest
	(*DeleteTopicRequest)(nil), // 5: doublecloud.kafka.v1.DeleteTopicRequest
	(*v1.Paging)(nil),          // 6: doublecloud.v1.Paging
	(*Topic)(nil),              // 7: doublecloud.kafka.v1.Topic
	(*v1.NextPage)(nil),        // 8: doublecloud.v1.NextPage
	(*TopicSpec)(nil),          // 9: doublecloud.kafka.v1.TopicSpec
	(*v1.Operation)(nil),       // 10: doublecloud.v1.Operation
}
var file_doublecloud_kafka_v1_topic_service_proto_depIdxs = []int32{
	6,  // 0: doublecloud.kafka.v1.ListTopicsRequest.paging:type_name -> doublecloud.v1.Paging
	7,  // 1: doublecloud.kafka.v1.ListTopicsResponse.topics:type_name -> doublecloud.kafka.v1.Topic
	8,  // 2: doublecloud.kafka.v1.ListTopicsResponse.next_page:type_name -> doublecloud.v1.NextPage
	9,  // 3: doublecloud.kafka.v1.CreateTopicRequest.topic_spec:type_name -> doublecloud.kafka.v1.TopicSpec
	9,  // 4: doublecloud.kafka.v1.UpdateTopicRequest.topic_spec:type_name -> doublecloud.kafka.v1.TopicSpec
	0,  // 5: doublecloud.kafka.v1.TopicService.Get:input_type -> doublecloud.kafka.v1.GetTopicRequest
	1,  // 6: doublecloud.kafka.v1.TopicService.List:input_type -> doublecloud.kafka.v1.ListTopicsRequest
	3,  // 7: doublecloud.kafka.v1.TopicService.Create:input_type -> doublecloud.kafka.v1.CreateTopicRequest
	4,  // 8: doublecloud.kafka.v1.TopicService.Update:input_type -> doublecloud.kafka.v1.UpdateTopicRequest
	5,  // 9: doublecloud.kafka.v1.TopicService.Delete:input_type -> doublecloud.kafka.v1.DeleteTopicRequest
	7,  // 10: doublecloud.kafka.v1.TopicService.Get:output_type -> doublecloud.kafka.v1.Topic
	2,  // 11: doublecloud.kafka.v1.TopicService.List:output_type -> doublecloud.kafka.v1.ListTopicsResponse
	10, // 12: doublecloud.kafka.v1.TopicService.Create:output_type -> doublecloud.v1.Operation
	10, // 13: doublecloud.kafka.v1.TopicService.Update:output_type -> doublecloud.v1.Operation
	10, // 14: doublecloud.kafka.v1.TopicService.Delete:output_type -> doublecloud.v1.Operation
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_doublecloud_kafka_v1_topic_service_proto_init() }
func file_doublecloud_kafka_v1_topic_service_proto_init() {
	if File_doublecloud_kafka_v1_topic_service_proto != nil {
		return
	}
	file_doublecloud_kafka_v1_topic_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doublecloud_kafka_v1_topic_service_proto_rawDesc), len(file_doublecloud_kafka_v1_topic_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doublecloud_kafka_v1_topic_service_proto_goTypes,
		DependencyIndexes: file_doublecloud_kafka_v1_topic_service_proto_depIdxs,
		MessageInfos:      file_doublecloud_kafka_v1_topic_service_proto_msgTypes,
	}.Build()
	File_doublecloud_kafka_v1_topic_service_proto = out.File
	file_doublecloud_kafka_v1_topic_service_proto_goTypes = nil
	file_doublecloud_kafka_v1_topic_service_proto_depIdxs = nil
}
