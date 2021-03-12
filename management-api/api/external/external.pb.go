// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: external.proto

package external

import (
	proto "github.com/golang/protobuf/proto"
	api "go.nlx.io/nlx/management-api/api"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RequestAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *RequestAccessRequest) Reset() {
	*x = RequestAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAccessRequest) ProtoMessage() {}

func (x *RequestAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAccessRequest.ProtoReflect.Descriptor instead.
func (*RequestAccessRequest) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAccessRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type GetAccessRequestStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *GetAccessRequestStateRequest) Reset() {
	*x = GetAccessRequestStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessRequestStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessRequestStateRequest) ProtoMessage() {}

func (x *GetAccessRequestStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessRequestStateRequest.ProtoReflect.Descriptor instead.
func (*GetAccessRequestStateRequest) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccessRequestStateRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type GetAccessRequestStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State api.AccessRequestState `protobuf:"varint,1,opt,name=state,proto3,enum=nlx.management.AccessRequestState" json:"state,omitempty"`
}

func (x *GetAccessRequestStateResponse) Reset() {
	*x = GetAccessRequestStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessRequestStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessRequestStateResponse) ProtoMessage() {}

func (x *GetAccessRequestStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessRequestStateResponse.ProtoReflect.Descriptor instead.
func (*GetAccessRequestStateResponse) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccessRequestStateResponse) GetState() api.AccessRequestState {
	if x != nil {
		return x.State
	}
	return api.AccessRequestState_UNSPECIFIED
}

type GetAccessProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *GetAccessProofRequest) Reset() {
	*x = GetAccessProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessProofRequest) ProtoMessage() {}

func (x *GetAccessProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessProofRequest.ProtoReflect.Descriptor instead.
func (*GetAccessProofRequest) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccessProofRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type RequestAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceId uint64 `protobuf:"varint,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
}

func (x *RequestAccessResponse) Reset() {
	*x = RequestAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAccessResponse) ProtoMessage() {}

func (x *RequestAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAccessResponse.ProtoReflect.Descriptor instead.
func (*RequestAccessResponse) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{4}
}

func (x *RequestAccessResponse) GetReferenceId() uint64 {
	if x != nil {
		return x.ReferenceId
	}
	return 0
}

type RequestClaimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderReference string `protobuf:"bytes,1,opt,name=order_reference,json=orderReference,proto3" json:"order_reference,omitempty"`
}

func (x *RequestClaimRequest) Reset() {
	*x = RequestClaimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestClaimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestClaimRequest) ProtoMessage() {}

func (x *RequestClaimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestClaimRequest.ProtoReflect.Descriptor instead.
func (*RequestClaimRequest) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{5}
}

func (x *RequestClaimRequest) GetOrderReference() string {
	if x != nil {
		return x.OrderReference
	}
	return ""
}

type RequestClaimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claim string `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *RequestClaimResponse) Reset() {
	*x = RequestClaimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestClaimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestClaimResponse) ProtoMessage() {}

func (x *RequestClaimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestClaimResponse.ProtoReflect.Descriptor instead.
func (*RequestClaimResponse) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{6}
}

func (x *RequestClaimResponse) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

var File_external_proto protoreflect.FileDescriptor

var file_external_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x10, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x14, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6e, 0x6c, 0x78, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3a, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x13,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x14,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x32, 0xee, 0x02, 0x0a, 0x14, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e,
	0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x2e,
	0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x32, 0x80, 0x01, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x12, 0x2c, 0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x6e, 0x6c, 0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x6e, 0x6c, 0x78, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_external_proto_rawDescOnce sync.Once
	file_external_proto_rawDescData = file_external_proto_rawDesc
)

func file_external_proto_rawDescGZIP() []byte {
	file_external_proto_rawDescOnce.Do(func() {
		file_external_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_proto_rawDescData)
	})
	return file_external_proto_rawDescData
}

var file_external_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_external_proto_goTypes = []interface{}{
	(*RequestAccessRequest)(nil),          // 0: nlx.management.external.RequestAccessRequest
	(*GetAccessRequestStateRequest)(nil),  // 1: nlx.management.external.GetAccessRequestStateRequest
	(*GetAccessRequestStateResponse)(nil), // 2: nlx.management.external.GetAccessRequestStateResponse
	(*GetAccessProofRequest)(nil),         // 3: nlx.management.external.GetAccessProofRequest
	(*RequestAccessResponse)(nil),         // 4: nlx.management.external.RequestAccessResponse
	(*RequestClaimRequest)(nil),           // 5: nlx.management.external.RequestClaimRequest
	(*RequestClaimResponse)(nil),          // 6: nlx.management.external.RequestClaimResponse
	(api.AccessRequestState)(0),           // 7: nlx.management.AccessRequestState
	(*api.AccessProof)(nil),               // 8: nlx.management.AccessProof
}
var file_external_proto_depIdxs = []int32{
	7, // 0: nlx.management.external.GetAccessRequestStateResponse.state:type_name -> nlx.management.AccessRequestState
	0, // 1: nlx.management.external.AccessRequestService.RequestAccess:input_type -> nlx.management.external.RequestAccessRequest
	1, // 2: nlx.management.external.AccessRequestService.GetAccessRequestState:input_type -> nlx.management.external.GetAccessRequestStateRequest
	3, // 3: nlx.management.external.AccessRequestService.GetAccessProof:input_type -> nlx.management.external.GetAccessProofRequest
	5, // 4: nlx.management.external.DelegationService.RequestClaim:input_type -> nlx.management.external.RequestClaimRequest
	4, // 5: nlx.management.external.AccessRequestService.RequestAccess:output_type -> nlx.management.external.RequestAccessResponse
	2, // 6: nlx.management.external.AccessRequestService.GetAccessRequestState:output_type -> nlx.management.external.GetAccessRequestStateResponse
	8, // 7: nlx.management.external.AccessRequestService.GetAccessProof:output_type -> nlx.management.AccessProof
	6, // 8: nlx.management.external.DelegationService.RequestClaim:output_type -> nlx.management.external.RequestClaimResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_external_proto_init() }
func file_external_proto_init() {
	if File_external_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAccessRequest); i {
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
		file_external_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessRequestStateRequest); i {
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
		file_external_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessRequestStateResponse); i {
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
		file_external_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessProofRequest); i {
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
		file_external_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAccessResponse); i {
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
		file_external_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestClaimRequest); i {
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
		file_external_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestClaimResponse); i {
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
			RawDescriptor: file_external_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_external_proto_goTypes,
		DependencyIndexes: file_external_proto_depIdxs,
		MessageInfos:      file_external_proto_msgTypes,
	}.Build()
	File_external_proto = out.File
	file_external_proto_rawDesc = nil
	file_external_proto_goTypes = nil
	file_external_proto_depIdxs = nil
}
