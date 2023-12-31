// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.1
// source: loopify/v1/api.proto

package loopify

import (
	shared "github.com/weflux/loopify/protocol/shared"
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

type SurveyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExpectReplies uint32 `protobuf:"varint,2,opt,name=expect_replies,proto3" json:"expect_replies,omitempty"` // first, all
	Timeout       uint32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Topic         string `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Command       string `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	ContentType   string `protobuf:"bytes,6,opt,name=content_type,proto3" json:"content_type,omitempty"`
	PayloadText   string `protobuf:"bytes,7,opt,name=payload_text,proto3" json:"payload_text,omitempty"`
	PayloadBytes  []byte `protobuf:"bytes,8,opt,name=payload_bytes,proto3" json:"payload_bytes,omitempty"`
}

func (x *SurveyRequest) Reset() {
	*x = SurveyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyRequest) ProtoMessage() {}

func (x *SurveyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyRequest.ProtoReflect.Descriptor instead.
func (*SurveyRequest) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *SurveyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SurveyRequest) GetExpectReplies() uint32 {
	if x != nil {
		return x.ExpectReplies
	}
	return 0
}

func (x *SurveyRequest) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *SurveyRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SurveyRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *SurveyRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SurveyRequest) GetPayloadText() string {
	if x != nil {
		return x.PayloadText
	}
	return ""
}

func (x *SurveyRequest) GetPayloadBytes() []byte {
	if x != nil {
		return x.PayloadBytes
	}
	return nil
}

type SurveyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error   *Error                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Command string                `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Results []*SurveyReply_Result `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SurveyReply) Reset() {
	*x = SurveyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyReply) ProtoMessage() {}

func (x *SurveyReply) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyReply.ProtoReflect.Descriptor instead.
func (*SurveyReply) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *SurveyReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SurveyReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SurveyReply) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *SurveyReply) GetResults() []*SurveyReply_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic        string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Retain       bool   `protobuf:"varint,2,opt,name=retain,proto3" json:"retain,omitempty"`
	Qos          int32  `protobuf:"varint,3,opt,name=qos,proto3" json:"qos,omitempty"`
	Type         string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	ContentType  string `protobuf:"bytes,5,opt,name=content_type,proto3" json:"content_type,omitempty"`
	PayloadText  string `protobuf:"bytes,6,opt,name=payload_text,proto3" json:"payload_text,omitempty"`
	PayloadBytes []byte `protobuf:"bytes,7,opt,name=payload_bytes,proto3" json:"payload_bytes,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *PublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishRequest) GetRetain() bool {
	if x != nil {
		return x.Retain
	}
	return false
}

func (x *PublishRequest) GetQos() int32 {
	if x != nil {
		return x.Qos
	}
	return 0
}

func (x *PublishRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PublishRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *PublishRequest) GetPayloadText() string {
	if x != nil {
		return x.PayloadText
	}
	return ""
}

func (x *PublishRequest) GetPayloadBytes() []byte {
	if x != nil {
		return x.PayloadBytes
	}
	return nil
}

type PublishReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PublishReply) Reset() {
	*x = PublishReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReply) ProtoMessage() {}

func (x *PublishReply) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReply.ProtoReflect.Descriptor instead.
func (*PublishReply) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *PublishReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	Qos    int32  `protobuf:"varint,3,opt,name=qos,proto3" json:"qos,omitempty"`
	Broker string `protobuf:"bytes,4,opt,name=broker,proto3" json:"broker,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeRequest) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *SubscribeRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *SubscribeRequest) GetQos() int32 {
	if x != nil {
		return x.Qos
	}
	return 0
}

func (x *SubscribeRequest) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReply.ProtoReflect.Descriptor instead.
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	Broker string `protobuf:"bytes,3,opt,name=broker,proto3" json:"broker,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{6}
}

func (x *UnsubscribeRequest) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *UnsubscribeRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *UnsubscribeRequest) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

type UnsubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UnsubscribeReply) Reset() {
	*x = UnsubscribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeReply) ProtoMessage() {}

func (x *UnsubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeReply.ProtoReflect.Descriptor instead.
func (*UnsubscribeReply) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{7}
}

func (x *UnsubscribeReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{8}
}

func (x *DisconnectRequest) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *DisconnectRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DisconnectRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type DisconnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DisconnectReply) Reset() {
	*x = DisconnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectReply) ProtoMessage() {}

func (x *DisconnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectReply.ProtoReflect.Descriptor instead.
func (*DisconnectReply) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{9}
}

func (x *DisconnectReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type SurveyReply_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        *Error           `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Metadata     *shared.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ContentType  string           `protobuf:"bytes,5,opt,name=content_type,proto3" json:"content_type,omitempty"`
	PayloadText  string           `protobuf:"bytes,6,opt,name=payload_text,proto3" json:"payload_text,omitempty"`
	PayloadBytes []byte           `protobuf:"bytes,7,opt,name=payload_bytes,proto3" json:"payload_bytes,omitempty"`
}

func (x *SurveyReply_Result) Reset() {
	*x = SurveyReply_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loopify_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurveyReply_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurveyReply_Result) ProtoMessage() {}

func (x *SurveyReply_Result) ProtoReflect() protoreflect.Message {
	mi := &file_loopify_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurveyReply_Result.ProtoReflect.Descriptor instead.
func (*SurveyReply_Result) Descriptor() ([]byte, []int) {
	return file_loopify_v1_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SurveyReply_Result) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SurveyReply_Result) GetMetadata() *shared.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SurveyReply_Result) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SurveyReply_Result) GetPayloadText() string {
	if x != nil {
		return x.PayloadText
	}
	return ""
}

func (x *SurveyReply_Result) GetPayloadBytes() []byte {
	if x != nil {
		return x.PayloadBytes
	}
	return nil
}

var File_loopify_v1_api_proto protoreflect.FileDescriptor

var file_loopify_v1_api_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xff, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x22, 0xf2, 0x02, 0x0a, 0x0b, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x1a, 0xd5, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x37, 0x0a,
	0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c,
	0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x6f,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x5c, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x3b, 0x0a,
	0x10, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x57, 0x0a, 0x11, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x66, 0x6c, 0x75, 0x78, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loopify_v1_api_proto_rawDescOnce sync.Once
	file_loopify_v1_api_proto_rawDescData = file_loopify_v1_api_proto_rawDesc
)

func file_loopify_v1_api_proto_rawDescGZIP() []byte {
	file_loopify_v1_api_proto_rawDescOnce.Do(func() {
		file_loopify_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_loopify_v1_api_proto_rawDescData)
	})
	return file_loopify_v1_api_proto_rawDescData
}

var file_loopify_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_loopify_v1_api_proto_goTypes = []interface{}{
	(*SurveyRequest)(nil),      // 0: loopify.v1.SurveyRequest
	(*SurveyReply)(nil),        // 1: loopify.v1.SurveyReply
	(*PublishRequest)(nil),     // 2: loopify.v1.PublishRequest
	(*PublishReply)(nil),       // 3: loopify.v1.PublishReply
	(*SubscribeRequest)(nil),   // 4: loopify.v1.SubscribeRequest
	(*SubscribeReply)(nil),     // 5: loopify.v1.SubscribeReply
	(*UnsubscribeRequest)(nil), // 6: loopify.v1.UnsubscribeRequest
	(*UnsubscribeReply)(nil),   // 7: loopify.v1.UnsubscribeReply
	(*DisconnectRequest)(nil),  // 8: loopify.v1.DisconnectRequest
	(*DisconnectReply)(nil),    // 9: loopify.v1.DisconnectReply
	(*SurveyReply_Result)(nil), // 10: loopify.v1.SurveyReply.Result
	(*Error)(nil),              // 11: loopify.v1.Error
	(*shared.Metadata)(nil),    // 12: loopify.shared.Metadata
}
var file_loopify_v1_api_proto_depIdxs = []int32{
	11, // 0: loopify.v1.SurveyReply.error:type_name -> loopify.v1.Error
	10, // 1: loopify.v1.SurveyReply.results:type_name -> loopify.v1.SurveyReply.Result
	11, // 2: loopify.v1.PublishReply.error:type_name -> loopify.v1.Error
	11, // 3: loopify.v1.SubscribeReply.error:type_name -> loopify.v1.Error
	11, // 4: loopify.v1.UnsubscribeReply.error:type_name -> loopify.v1.Error
	11, // 5: loopify.v1.DisconnectReply.error:type_name -> loopify.v1.Error
	11, // 6: loopify.v1.SurveyReply.Result.error:type_name -> loopify.v1.Error
	12, // 7: loopify.v1.SurveyReply.Result.metadata:type_name -> loopify.shared.Metadata
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_loopify_v1_api_proto_init() }
func file_loopify_v1_api_proto_init() {
	if File_loopify_v1_api_proto != nil {
		return
	}
	file_loopify_v1_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_loopify_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyRequest); i {
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
		file_loopify_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyReply); i {
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
		file_loopify_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_loopify_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReply); i {
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
		file_loopify_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_loopify_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReply); i {
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
		file_loopify_v1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_loopify_v1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeReply); i {
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
		file_loopify_v1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_loopify_v1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectReply); i {
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
		file_loopify_v1_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurveyReply_Result); i {
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
			RawDescriptor: file_loopify_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loopify_v1_api_proto_goTypes,
		DependencyIndexes: file_loopify_v1_api_proto_depIdxs,
		MessageInfos:      file_loopify_v1_api_proto_msgTypes,
	}.Build()
	File_loopify_v1_api_proto = out.File
	file_loopify_v1_api_proto_rawDesc = nil
	file_loopify_v1_api_proto_goTypes = nil
	file_loopify_v1_api_proto_depIdxs = nil
}
