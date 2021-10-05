// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.7
// source: inspectionapi.proto

package inspectionapi

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Inway_State int32

const (
	Inway_UNKNOWN Inway_State = 0
	Inway_UP      Inway_State = 1
	Inway_DOWN    Inway_State = 2
)

// Enum value maps for Inway_State.
var (
	Inway_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "UP",
		2: "DOWN",
	}
	Inway_State_value = map[string]int32{
		"UNKNOWN": 0,
		"UP":      1,
		"DOWN":    2,
	}
)

func (x Inway_State) Enum() *Inway_State {
	p := new(Inway_State)
	*p = x
	return p
}

func (x Inway_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Inway_State) Descriptor() protoreflect.EnumDescriptor {
	return file_inspectionapi_proto_enumTypes[0].Descriptor()
}

func (Inway_State) Type() protoreflect.EnumType {
	return &file_inspectionapi_proto_enumTypes[0]
}

func (x Inway_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Inway_State.Descriptor instead.
func (Inway_State) EnumDescriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{0, 0}
}

type ListInOutwayStatisticsResponse_Statistics_Type int32

const (
	ListInOutwayStatisticsResponse_Statistics_INWAY  ListInOutwayStatisticsResponse_Statistics_Type = 0
	ListInOutwayStatisticsResponse_Statistics_OUTWAY ListInOutwayStatisticsResponse_Statistics_Type = 1
)

// Enum value maps for ListInOutwayStatisticsResponse_Statistics_Type.
var (
	ListInOutwayStatisticsResponse_Statistics_Type_name = map[int32]string{
		0: "INWAY",
		1: "OUTWAY",
	}
	ListInOutwayStatisticsResponse_Statistics_Type_value = map[string]int32{
		"INWAY":  0,
		"OUTWAY": 1,
	}
)

func (x ListInOutwayStatisticsResponse_Statistics_Type) Enum() *ListInOutwayStatisticsResponse_Statistics_Type {
	p := new(ListInOutwayStatisticsResponse_Statistics_Type)
	*p = x
	return p
}

func (x ListInOutwayStatisticsResponse_Statistics_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListInOutwayStatisticsResponse_Statistics_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_inspectionapi_proto_enumTypes[1].Descriptor()
}

func (ListInOutwayStatisticsResponse_Statistics_Type) Type() protoreflect.EnumType {
	return &file_inspectionapi_proto_enumTypes[1]
}

func (x ListInOutwayStatisticsResponse_Statistics_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListInOutwayStatisticsResponse_Statistics_Type.Descriptor instead.
func (ListInOutwayStatisticsResponse_Statistics_Type) EnumDescriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{7, 0, 0}
}

type Inway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	State   Inway_State `protobuf:"varint,2,opt,name=state,proto3,enum=inspectionapi.Inway_State" json:"state,omitempty"`
}

func (x *Inway) Reset() {
	*x = Inway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inway) ProtoMessage() {}

func (x *Inway) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inway.ProtoReflect.Descriptor instead.
func (*Inway) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{0}
}

func (x *Inway) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Inway) GetState() Inway_State {
	if x != nil {
		return x.State
	}
	return Inway_UNKNOWN
}

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber string `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{1}
}

func (x *Organization) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Costs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OneTime int32 `protobuf:"varint,1,opt,name=one_time,json=oneTime,proto3" json:"one_time,omitempty"`
	Monthly int32 `protobuf:"varint,2,opt,name=monthly,proto3" json:"monthly,omitempty"`
	Request int32 `protobuf:"varint,3,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *Costs) Reset() {
	*x = Costs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Costs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Costs) ProtoMessage() {}

func (x *Costs) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Costs.ProtoReflect.Descriptor instead.
func (*Costs) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{2}
}

func (x *Costs) GetOneTime() int32 {
	if x != nil {
		return x.OneTime
	}
	return 0
}

func (x *Costs) GetMonthly() int32 {
	if x != nil {
		return x.Monthly
	}
	return 0
}

func (x *Costs) GetRequest() int32 {
	if x != nil {
		return x.Request
	}
	return 0
}

type ListServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ListServicesResponse_Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListServicesResponse) Reset() {
	*x = ListServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesResponse) ProtoMessage() {}

func (x *ListServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesResponse.ProtoReflect.Descriptor instead.
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{3}
}

func (x *ListServicesResponse) GetServices() []*ListServicesResponse_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type ListOrganizationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organizations []*Organization `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty"`
}

func (x *ListOrganizationsResponse) Reset() {
	*x = ListOrganizationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrganizationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsResponse) ProtoMessage() {}

func (x *ListOrganizationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsResponse.ProtoReflect.Descriptor instead.
func (*ListOrganizationsResponse) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrganizationsResponse) GetOrganizations() []*Organization {
	if x != nil {
		return x.Organizations
	}
	return nil
}

type GetOrganizationInwayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationSerialNumber string `protobuf:"bytes,1,opt,name=organization_serial_number,json=organizationSerialNumber,proto3" json:"organization_serial_number,omitempty"`
}

func (x *GetOrganizationInwayRequest) Reset() {
	*x = GetOrganizationInwayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrganizationInwayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrganizationInwayRequest) ProtoMessage() {}

func (x *GetOrganizationInwayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrganizationInwayRequest.ProtoReflect.Descriptor instead.
func (*GetOrganizationInwayRequest) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrganizationInwayRequest) GetOrganizationSerialNumber() string {
	if x != nil {
		return x.OrganizationSerialNumber
	}
	return ""
}

type GetOrganizationInwayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetOrganizationInwayResponse) Reset() {
	*x = GetOrganizationInwayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrganizationInwayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrganizationInwayResponse) ProtoMessage() {}

func (x *GetOrganizationInwayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrganizationInwayResponse.ProtoReflect.Descriptor instead.
func (*GetOrganizationInwayResponse) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrganizationInwayResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ListInOutwayStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []*ListInOutwayStatisticsResponse_Statistics `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *ListInOutwayStatisticsResponse) Reset() {
	*x = ListInOutwayStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInOutwayStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInOutwayStatisticsResponse) ProtoMessage() {}

func (x *ListInOutwayStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInOutwayStatisticsResponse.ProtoReflect.Descriptor instead.
func (*ListInOutwayStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{7}
}

func (x *ListInOutwayStatisticsResponse) GetVersions() []*ListInOutwayStatisticsResponse_Statistics {
	if x != nil {
		return x.Versions
	}
	return nil
}

type ListServicesResponse_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DocumentationUrl     string        `protobuf:"bytes,2,opt,name=documentation_url,json=documentationUrl,proto3" json:"documentation_url,omitempty"`
	ApiSpecificationType string        `protobuf:"bytes,3,opt,name=api_specification_type,json=apiSpecificationType,proto3" json:"api_specification_type,omitempty"`
	Internal             bool          `protobuf:"varint,4,opt,name=internal,proto3" json:"internal,omitempty"`
	PublicSupportContact string        `protobuf:"bytes,5,opt,name=public_support_contact,json=publicSupportContact,proto3" json:"public_support_contact,omitempty"`
	Inways               []*Inway      `protobuf:"bytes,6,rep,name=inways,proto3" json:"inways,omitempty"`
	Costs                *Costs        `protobuf:"bytes,7,opt,name=costs,proto3" json:"costs,omitempty"`
	Organization         *Organization `protobuf:"bytes,8,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *ListServicesResponse_Service) Reset() {
	*x = ListServicesResponse_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesResponse_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesResponse_Service) ProtoMessage() {}

func (x *ListServicesResponse_Service) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesResponse_Service.ProtoReflect.Descriptor instead.
func (*ListServicesResponse_Service) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListServicesResponse_Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListServicesResponse_Service) GetDocumentationUrl() string {
	if x != nil {
		return x.DocumentationUrl
	}
	return ""
}

func (x *ListServicesResponse_Service) GetApiSpecificationType() string {
	if x != nil {
		return x.ApiSpecificationType
	}
	return ""
}

func (x *ListServicesResponse_Service) GetInternal() bool {
	if x != nil {
		return x.Internal
	}
	return false
}

func (x *ListServicesResponse_Service) GetPublicSupportContact() string {
	if x != nil {
		return x.PublicSupportContact
	}
	return ""
}

func (x *ListServicesResponse_Service) GetInways() []*Inway {
	if x != nil {
		return x.Inways
	}
	return nil
}

func (x *ListServicesResponse_Service) GetCosts() *Costs {
	if x != nil {
		return x.Costs
	}
	return nil
}

func (x *ListServicesResponse_Service) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

type ListInOutwayStatisticsResponse_Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    ListInOutwayStatisticsResponse_Statistics_Type `protobuf:"varint,1,opt,name=type,proto3,enum=inspectionapi.ListInOutwayStatisticsResponse_Statistics_Type" json:"type,omitempty"`
	Version string                                         `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Amount  uint32                                         `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ListInOutwayStatisticsResponse_Statistics) Reset() {
	*x = ListInOutwayStatisticsResponse_Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspectionapi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInOutwayStatisticsResponse_Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInOutwayStatisticsResponse_Statistics) ProtoMessage() {}

func (x *ListInOutwayStatisticsResponse_Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_inspectionapi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInOutwayStatisticsResponse_Statistics.ProtoReflect.Descriptor instead.
func (*ListInOutwayStatisticsResponse_Statistics) Descriptor() ([]byte, []int) {
	return file_inspectionapi_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListInOutwayStatisticsResponse_Statistics) GetType() ListInOutwayStatisticsResponse_Statistics_Type {
	if x != nil {
		return x.Type
	}
	return ListInOutwayStatisticsResponse_Statistics_INWAY
}

func (x *ListInOutwayStatisticsResponse_Statistics) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ListInOutwayStatisticsResponse_Statistics) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_inspectionapi_proto protoreflect.FileDescriptor

var file_inspectionapi_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7b, 0x0a, 0x05, 0x49, 0x6e, 0x77, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x6e, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x22, 0x47, 0x0a, 0x0c,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x05, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xcf, 0x03,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0xed, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x16,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x70,
	0x69, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x34,
	0x0a, 0x16, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x69, 0x6e, 0x77, 0x61, 0x79, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x77, 0x61, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x77, 0x61,
	0x79, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x3f,
	0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x5e, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x5b, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c,
	0x0a, 0x1a, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x18, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xa9, 0x02, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x4f, 0x75, 0x74, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x4f, 0x75, 0x74, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0xb0, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x51,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x1d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x49,
	0x4e, 0x57, 0x41, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x55, 0x54, 0x57, 0x41, 0x59,
	0x10, 0x01, 0x32, 0xf3, 0x03, 0x0a, 0x13, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x80, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x2d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x71, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x77, 0x61, 0x79, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x4f, 0x75, 0x74,
	0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2d, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x77,
	0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_inspectionapi_proto_rawDescOnce sync.Once
	file_inspectionapi_proto_rawDescData = file_inspectionapi_proto_rawDesc
)

func file_inspectionapi_proto_rawDescGZIP() []byte {
	file_inspectionapi_proto_rawDescOnce.Do(func() {
		file_inspectionapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_inspectionapi_proto_rawDescData)
	})
	return file_inspectionapi_proto_rawDescData
}

var file_inspectionapi_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_inspectionapi_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_inspectionapi_proto_goTypes = []interface{}{
	(Inway_State)(0), // 0: inspectionapi.Inway.State
	(ListInOutwayStatisticsResponse_Statistics_Type)(0), // 1: inspectionapi.ListInOutwayStatisticsResponse.Statistics.Type
	(*Inway)(nil),                                     // 2: inspectionapi.Inway
	(*Organization)(nil),                              // 3: inspectionapi.Organization
	(*Costs)(nil),                                     // 4: inspectionapi.Costs
	(*ListServicesResponse)(nil),                      // 5: inspectionapi.ListServicesResponse
	(*ListOrganizationsResponse)(nil),                 // 6: inspectionapi.ListOrganizationsResponse
	(*GetOrganizationInwayRequest)(nil),               // 7: inspectionapi.GetOrganizationInwayRequest
	(*GetOrganizationInwayResponse)(nil),              // 8: inspectionapi.GetOrganizationInwayResponse
	(*ListInOutwayStatisticsResponse)(nil),            // 9: inspectionapi.ListInOutwayStatisticsResponse
	(*ListServicesResponse_Service)(nil),              // 10: inspectionapi.ListServicesResponse.Service
	(*ListInOutwayStatisticsResponse_Statistics)(nil), // 11: inspectionapi.ListInOutwayStatisticsResponse.Statistics
	(*emptypb.Empty)(nil),                             // 12: google.protobuf.Empty
}
var file_inspectionapi_proto_depIdxs = []int32{
	0,  // 0: inspectionapi.Inway.state:type_name -> inspectionapi.Inway.State
	10, // 1: inspectionapi.ListServicesResponse.services:type_name -> inspectionapi.ListServicesResponse.Service
	3,  // 2: inspectionapi.ListOrganizationsResponse.organizations:type_name -> inspectionapi.Organization
	11, // 3: inspectionapi.ListInOutwayStatisticsResponse.versions:type_name -> inspectionapi.ListInOutwayStatisticsResponse.Statistics
	2,  // 4: inspectionapi.ListServicesResponse.Service.inways:type_name -> inspectionapi.Inway
	4,  // 5: inspectionapi.ListServicesResponse.Service.costs:type_name -> inspectionapi.Costs
	3,  // 6: inspectionapi.ListServicesResponse.Service.organization:type_name -> inspectionapi.Organization
	1,  // 7: inspectionapi.ListInOutwayStatisticsResponse.Statistics.type:type_name -> inspectionapi.ListInOutwayStatisticsResponse.Statistics.Type
	12, // 8: inspectionapi.DirectoryInspection.ListServices:input_type -> google.protobuf.Empty
	12, // 9: inspectionapi.DirectoryInspection.ListOrganizations:input_type -> google.protobuf.Empty
	7,  // 10: inspectionapi.DirectoryInspection.GetOrganizationInway:input_type -> inspectionapi.GetOrganizationInwayRequest
	12, // 11: inspectionapi.DirectoryInspection.ListInOutwayStatistics:input_type -> google.protobuf.Empty
	5,  // 12: inspectionapi.DirectoryInspection.ListServices:output_type -> inspectionapi.ListServicesResponse
	6,  // 13: inspectionapi.DirectoryInspection.ListOrganizations:output_type -> inspectionapi.ListOrganizationsResponse
	8,  // 14: inspectionapi.DirectoryInspection.GetOrganizationInway:output_type -> inspectionapi.GetOrganizationInwayResponse
	9,  // 15: inspectionapi.DirectoryInspection.ListInOutwayStatistics:output_type -> inspectionapi.ListInOutwayStatisticsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_inspectionapi_proto_init() }
func file_inspectionapi_proto_init() {
	if File_inspectionapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inspectionapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inway); i {
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
		file_inspectionapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
		file_inspectionapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Costs); i {
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
		file_inspectionapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesResponse); i {
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
		file_inspectionapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrganizationsResponse); i {
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
		file_inspectionapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrganizationInwayRequest); i {
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
		file_inspectionapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrganizationInwayResponse); i {
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
		file_inspectionapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInOutwayStatisticsResponse); i {
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
		file_inspectionapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesResponse_Service); i {
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
		file_inspectionapi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInOutwayStatisticsResponse_Statistics); i {
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
			RawDescriptor: file_inspectionapi_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inspectionapi_proto_goTypes,
		DependencyIndexes: file_inspectionapi_proto_depIdxs,
		EnumInfos:         file_inspectionapi_proto_enumTypes,
		MessageInfos:      file_inspectionapi_proto_msgTypes,
	}.Build()
	File_inspectionapi_proto = out.File
	file_inspectionapi_proto_rawDesc = nil
	file_inspectionapi_proto_goTypes = nil
	file_inspectionapi_proto_depIdxs = nil
}
