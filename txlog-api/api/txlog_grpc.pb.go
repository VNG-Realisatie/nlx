// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: txlog.proto

package api

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TXLogServiceClient is the client API for TXLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TXLogServiceClient interface {
	ListRecords(ctx context.Context, in *ListRecordsRequest, opts ...grpc.CallOption) (*ListRecordsResponse, error)
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
}

type tXLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTXLogServiceClient(cc grpc.ClientConnInterface) TXLogServiceClient {
	return &tXLogServiceClient{cc}
}

func (c *tXLogServiceClient) ListRecords(ctx context.Context, in *ListRecordsRequest, opts ...grpc.CallOption) (*ListRecordsResponse, error) {
	out := new(ListRecordsResponse)
	err := c.cc.Invoke(ctx, "/nlx.txlog.TXLogService/ListRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tXLogServiceClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, "/nlx.txlog.TXLogService/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TXLogServiceServer is the server API for TXLogService service.
// All implementations must embed UnimplementedTXLogServiceServer
// for forward compatibility
type TXLogServiceServer interface {
	ListRecords(context.Context, *ListRecordsRequest) (*ListRecordsResponse, error)
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	mustEmbedUnimplementedTXLogServiceServer()
}

// UnimplementedTXLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTXLogServiceServer struct {
}

func (UnimplementedTXLogServiceServer) ListRecords(context.Context, *ListRecordsRequest) (*ListRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecords not implemented")
}
func (UnimplementedTXLogServiceServer) CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedTXLogServiceServer) mustEmbedUnimplementedTXLogServiceServer() {}

// UnsafeTXLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TXLogServiceServer will
// result in compilation errors.
type UnsafeTXLogServiceServer interface {
	mustEmbedUnimplementedTXLogServiceServer()
}

func RegisterTXLogServiceServer(s grpc.ServiceRegistrar, srv TXLogServiceServer) {
	s.RegisterService(&TXLogService_ServiceDesc, srv)
}

func _TXLogService_ListRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXLogServiceServer).ListRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.txlog.TXLogService/ListRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXLogServiceServer).ListRecords(ctx, req.(*ListRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TXLogService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXLogServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.txlog.TXLogService/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXLogServiceServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TXLogService_ServiceDesc is the grpc.ServiceDesc for TXLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TXLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nlx.txlog.TXLogService",
	HandlerType: (*TXLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRecords",
			Handler:    _TXLogService_ListRecords_Handler,
		},
		{
			MethodName: "CreateRecord",
			Handler:    _TXLogService_CreateRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "txlog.proto",
}
