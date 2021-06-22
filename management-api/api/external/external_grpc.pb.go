// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package external

import (
	context "context"
	api "go.nlx.io/nlx/management-api/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccessRequestServiceClient is the client API for AccessRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessRequestServiceClient interface {
	RequestAccess(ctx context.Context, in *RequestAccessRequest, opts ...grpc.CallOption) (*RequestAccessResponse, error)
	GetAccessRequestState(ctx context.Context, in *GetAccessRequestStateRequest, opts ...grpc.CallOption) (*GetAccessRequestStateResponse, error)
	GetAccessProof(ctx context.Context, in *GetAccessProofRequest, opts ...grpc.CallOption) (*api.AccessProof, error)
}

type accessRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessRequestServiceClient(cc grpc.ClientConnInterface) AccessRequestServiceClient {
	return &accessRequestServiceClient{cc}
}

func (c *accessRequestServiceClient) RequestAccess(ctx context.Context, in *RequestAccessRequest, opts ...grpc.CallOption) (*RequestAccessResponse, error) {
	out := new(RequestAccessResponse)
	err := c.cc.Invoke(ctx, "/nlx.management.external.AccessRequestService/RequestAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRequestServiceClient) GetAccessRequestState(ctx context.Context, in *GetAccessRequestStateRequest, opts ...grpc.CallOption) (*GetAccessRequestStateResponse, error) {
	out := new(GetAccessRequestStateResponse)
	err := c.cc.Invoke(ctx, "/nlx.management.external.AccessRequestService/GetAccessRequestState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRequestServiceClient) GetAccessProof(ctx context.Context, in *GetAccessProofRequest, opts ...grpc.CallOption) (*api.AccessProof, error) {
	out := new(api.AccessProof)
	err := c.cc.Invoke(ctx, "/nlx.management.external.AccessRequestService/GetAccessProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessRequestServiceServer is the server API for AccessRequestService service.
// All implementations must embed UnimplementedAccessRequestServiceServer
// for forward compatibility
type AccessRequestServiceServer interface {
	RequestAccess(context.Context, *RequestAccessRequest) (*RequestAccessResponse, error)
	GetAccessRequestState(context.Context, *GetAccessRequestStateRequest) (*GetAccessRequestStateResponse, error)
	GetAccessProof(context.Context, *GetAccessProofRequest) (*api.AccessProof, error)
	mustEmbedUnimplementedAccessRequestServiceServer()
}

// UnimplementedAccessRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccessRequestServiceServer struct {
}

func (UnimplementedAccessRequestServiceServer) RequestAccess(context.Context, *RequestAccessRequest) (*RequestAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAccess not implemented")
}
func (UnimplementedAccessRequestServiceServer) GetAccessRequestState(context.Context, *GetAccessRequestStateRequest) (*GetAccessRequestStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessRequestState not implemented")
}
func (UnimplementedAccessRequestServiceServer) GetAccessProof(context.Context, *GetAccessProofRequest) (*api.AccessProof, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessProof not implemented")
}
func (UnimplementedAccessRequestServiceServer) mustEmbedUnimplementedAccessRequestServiceServer() {}

// UnsafeAccessRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessRequestServiceServer will
// result in compilation errors.
type UnsafeAccessRequestServiceServer interface {
	mustEmbedUnimplementedAccessRequestServiceServer()
}

func RegisterAccessRequestServiceServer(s grpc.ServiceRegistrar, srv AccessRequestServiceServer) {
	s.RegisterService(&AccessRequestService_ServiceDesc, srv)
}

func _AccessRequestService_RequestAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRequestServiceServer).RequestAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.management.external.AccessRequestService/RequestAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRequestServiceServer).RequestAccess(ctx, req.(*RequestAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRequestService_GetAccessRequestState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessRequestStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRequestServiceServer).GetAccessRequestState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.management.external.AccessRequestService/GetAccessRequestState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRequestServiceServer).GetAccessRequestState(ctx, req.(*GetAccessRequestStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRequestService_GetAccessProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRequestServiceServer).GetAccessProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.management.external.AccessRequestService/GetAccessProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRequestServiceServer).GetAccessProof(ctx, req.(*GetAccessProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessRequestService_ServiceDesc is the grpc.ServiceDesc for AccessRequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessRequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nlx.management.external.AccessRequestService",
	HandlerType: (*AccessRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAccess",
			Handler:    _AccessRequestService_RequestAccess_Handler,
		},
		{
			MethodName: "GetAccessRequestState",
			Handler:    _AccessRequestService_GetAccessRequestState_Handler,
		},
		{
			MethodName: "GetAccessProof",
			Handler:    _AccessRequestService_GetAccessProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external.proto",
}

// DelegationServiceClient is the client API for DelegationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DelegationServiceClient interface {
	RequestClaim(ctx context.Context, in *RequestClaimRequest, opts ...grpc.CallOption) (*RequestClaimResponse, error)
	ListOrders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrdersResponse, error)
}

type delegationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDelegationServiceClient(cc grpc.ClientConnInterface) DelegationServiceClient {
	return &delegationServiceClient{cc}
}

func (c *delegationServiceClient) RequestClaim(ctx context.Context, in *RequestClaimRequest, opts ...grpc.CallOption) (*RequestClaimResponse, error) {
	out := new(RequestClaimResponse)
	err := c.cc.Invoke(ctx, "/nlx.management.external.DelegationService/RequestClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegationServiceClient) ListOrders(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, "/nlx.management.external.DelegationService/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DelegationServiceServer is the server API for DelegationService service.
// All implementations must embed UnimplementedDelegationServiceServer
// for forward compatibility
type DelegationServiceServer interface {
	RequestClaim(context.Context, *RequestClaimRequest) (*RequestClaimResponse, error)
	ListOrders(context.Context, *emptypb.Empty) (*ListOrdersResponse, error)
	mustEmbedUnimplementedDelegationServiceServer()
}

// UnimplementedDelegationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDelegationServiceServer struct {
}

func (UnimplementedDelegationServiceServer) RequestClaim(context.Context, *RequestClaimRequest) (*RequestClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestClaim not implemented")
}
func (UnimplementedDelegationServiceServer) ListOrders(context.Context, *emptypb.Empty) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedDelegationServiceServer) mustEmbedUnimplementedDelegationServiceServer() {}

// UnsafeDelegationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DelegationServiceServer will
// result in compilation errors.
type UnsafeDelegationServiceServer interface {
	mustEmbedUnimplementedDelegationServiceServer()
}

func RegisterDelegationServiceServer(s grpc.ServiceRegistrar, srv DelegationServiceServer) {
	s.RegisterService(&DelegationService_ServiceDesc, srv)
}

func _DelegationService_RequestClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegationServiceServer).RequestClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.management.external.DelegationService/RequestClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegationServiceServer).RequestClaim(ctx, req.(*RequestClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegationService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegationServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlx.management.external.DelegationService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegationServiceServer).ListOrders(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DelegationService_ServiceDesc is the grpc.ServiceDesc for DelegationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DelegationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nlx.management.external.DelegationService",
	HandlerType: (*DelegationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestClaim",
			Handler:    _DelegationService_RequestClaim_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _DelegationService_ListOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external.proto",
}
