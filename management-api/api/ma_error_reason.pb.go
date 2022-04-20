// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: ma_error_reason.proto

package api

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	// Do not use this default value.
	ErrorReason_ERROR_REASON_UNSPECIFIED ErrorReason = 0
	// The order that is being used is revoked
	ErrorReason_ORDER_REVOKED ErrorReason = 1
	// The order could not be found
	ErrorReason_ORDER_NOT_FOUND ErrorReason = 2
	// Something went wrong while trying to retrieve the claim
	ErrorReason_UNABLE_TO_RETRIEVE_CLAIM ErrorReason = 3
	// Something went wrong while trying to sign the claim
	ErrorReason_UNABLE_TO_SIGN_CLAIM ErrorReason = 4
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "ERROR_REASON_UNSPECIFIED",
		1: "ORDER_REVOKED",
		2: "ORDER_NOT_FOUND",
		3: "UNABLE_TO_RETRIEVE_CLAIM",
		4: "UNABLE_TO_SIGN_CLAIM",
	}
	ErrorReason_value = map[string]int32{
		"ERROR_REASON_UNSPECIFIED": 0,
		"ORDER_REVOKED":            1,
		"ORDER_NOT_FOUND":          2,
		"UNABLE_TO_RETRIEVE_CLAIM": 3,
		"UNABLE_TO_SIGN_CLAIM":     4,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_ma_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_ma_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_ma_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_ma_error_reason_proto protoreflect.FileDescriptor

var file_ma_error_reason_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x6c, 0x78, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x8b, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x1c, 0x0a,
	0x18, 0x55, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x49,
	0x45, 0x56, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x55,
	0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x43, 0x4c,
	0x41, 0x49, 0x4d, 0x10, 0x04, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x6f, 0x2e, 0x6e, 0x6c, 0x78, 0x2e,
	0x69, 0x6f, 0x2f, 0x6e, 0x6c, 0x78, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ma_error_reason_proto_rawDescOnce sync.Once
	file_ma_error_reason_proto_rawDescData = file_ma_error_reason_proto_rawDesc
)

func file_ma_error_reason_proto_rawDescGZIP() []byte {
	file_ma_error_reason_proto_rawDescOnce.Do(func() {
		file_ma_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_ma_error_reason_proto_rawDescData)
	})
	return file_ma_error_reason_proto_rawDescData
}

var file_ma_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ma_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: nlx.management.ErrorReason
}
var file_ma_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ma_error_reason_proto_init() }
func file_ma_error_reason_proto_init() {
	if File_ma_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ma_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ma_error_reason_proto_goTypes,
		DependencyIndexes: file_ma_error_reason_proto_depIdxs,
		EnumInfos:         file_ma_error_reason_proto_enumTypes,
	}.Build()
	File_ma_error_reason_proto = out.File
	file_ma_error_reason_proto_rawDesc = nil
	file_ma_error_reason_proto_goTypes = nil
	file_ma_error_reason_proto_depIdxs = nil
}
