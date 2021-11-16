// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package directoryapi

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DirectoryClient is the client API for Directory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryClient interface {
	RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error)
	ClearOrganizationInway(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error)
	ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrganizationsResponse, error)
	GetOrganizationInway(ctx context.Context, in *GetOrganizationInwayRequest, opts ...grpc.CallOption) (*GetOrganizationInwayResponse, error)
	ListInOutwayStatistics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInOutwayStatisticsResponse, error)
}

type directoryClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryClient(cc grpc.ClientConnInterface) DirectoryClient {
	return &directoryClient{cc}
}

func (c *directoryClient) RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error) {
	out := new(RegisterInwayResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.Directory/RegisterInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) ClearOrganizationInway(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/directoryapi.Directory/ClearOrganizationInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.Directory/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrganizationsResponse, error) {
	out := new(ListOrganizationsResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.Directory/ListOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) GetOrganizationInway(ctx context.Context, in *GetOrganizationInwayRequest, opts ...grpc.CallOption) (*GetOrganizationInwayResponse, error) {
	out := new(GetOrganizationInwayResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.Directory/GetOrganizationInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) ListInOutwayStatistics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInOutwayStatisticsResponse, error) {
	out := new(ListInOutwayStatisticsResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.Directory/ListInOutwayStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryServer is the server API for Directory service.
// All implementations must embed UnimplementedDirectoryServer
// for forward compatibility
type DirectoryServer interface {
	RegisterInway(context.Context, *RegisterInwayRequest) (*RegisterInwayResponse, error)
	ClearOrganizationInway(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ListServices(context.Context, *emptypb.Empty) (*ListServicesResponse, error)
	ListOrganizations(context.Context, *emptypb.Empty) (*ListOrganizationsResponse, error)
	GetOrganizationInway(context.Context, *GetOrganizationInwayRequest) (*GetOrganizationInwayResponse, error)
	ListInOutwayStatistics(context.Context, *emptypb.Empty) (*ListInOutwayStatisticsResponse, error)
	mustEmbedUnimplementedDirectoryServer()
}

// UnimplementedDirectoryServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryServer struct {
}

func (UnimplementedDirectoryServer) RegisterInway(context.Context, *RegisterInwayRequest) (*RegisterInwayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInway not implemented")
}
func (UnimplementedDirectoryServer) ClearOrganizationInway(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearOrganizationInway not implemented")
}
func (UnimplementedDirectoryServer) ListServices(context.Context, *emptypb.Empty) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedDirectoryServer) ListOrganizations(context.Context, *emptypb.Empty) (*ListOrganizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizations not implemented")
}
func (UnimplementedDirectoryServer) GetOrganizationInway(context.Context, *GetOrganizationInwayRequest) (*GetOrganizationInwayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationInway not implemented")
}
func (UnimplementedDirectoryServer) ListInOutwayStatistics(context.Context, *emptypb.Empty) (*ListInOutwayStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInOutwayStatistics not implemented")
}
func (UnimplementedDirectoryServer) mustEmbedUnimplementedDirectoryServer() {}

// UnsafeDirectoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryServer will
// result in compilation errors.
type UnsafeDirectoryServer interface {
	mustEmbedUnimplementedDirectoryServer()
}

func RegisterDirectoryServer(s grpc.ServiceRegistrar, srv DirectoryServer) {
	s.RegisterService(&Directory_ServiceDesc, srv)
}

func _Directory_RegisterInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInwayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).RegisterInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/RegisterInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).RegisterInway(ctx, req.(*RegisterInwayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_ClearOrganizationInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).ClearOrganizationInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/ClearOrganizationInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).ClearOrganizationInway(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).ListServices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_ListOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).ListOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/ListOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).ListOrganizations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_GetOrganizationInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationInwayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).GetOrganizationInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/GetOrganizationInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).GetOrganizationInway(ctx, req.(*GetOrganizationInwayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_ListInOutwayStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).ListInOutwayStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/ListInOutwayStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).ListInOutwayStatistics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Directory_ServiceDesc is the grpc.ServiceDesc for Directory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Directory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "directoryapi.Directory",
	HandlerType: (*DirectoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInway",
			Handler:    _Directory_RegisterInway_Handler,
		},
		{
			MethodName: "ClearOrganizationInway",
			Handler:    _Directory_ClearOrganizationInway_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _Directory_ListServices_Handler,
		},
		{
			MethodName: "ListOrganizations",
			Handler:    _Directory_ListOrganizations_Handler,
		},
		{
			MethodName: "GetOrganizationInway",
			Handler:    _Directory_GetOrganizationInway_Handler,
		},
		{
			MethodName: "ListInOutwayStatistics",
			Handler:    _Directory_ListInOutwayStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directoryapi.proto",
}

// DirectoryRegistrationClient is the client API for DirectoryRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryRegistrationClient interface {
	RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error)
	ClearOrganizationInway(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error)
	ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrganizationsResponse, error)
	GetOrganizationInway(ctx context.Context, in *GetOrganizationInwayRequest, opts ...grpc.CallOption) (*GetOrganizationInwayResponse, error)
	ListInOutwayStatistics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInOutwayStatisticsResponse, error)
}

type directoryRegistrationClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryRegistrationClient(cc grpc.ClientConnInterface) DirectoryRegistrationClient {
	return &directoryRegistrationClient{cc}
}

func (c *directoryRegistrationClient) RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error) {
	out := new(RegisterInwayResponse)
	err := c.cc.Invoke(ctx, "/registrationapi.DirectoryRegistration/RegisterInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryRegistrationClient) ClearOrganizationInway(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/registrationapi.DirectoryRegistration/ClearOrganizationInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryRegistrationClient) ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/registrationapi.DirectoryRegistration/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryRegistrationClient) ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrganizationsResponse, error) {
	out := new(ListOrganizationsResponse)
	err := c.cc.Invoke(ctx, "/registrationapi.DirectoryRegistration/ListOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryRegistrationClient) GetOrganizationInway(ctx context.Context, in *GetOrganizationInwayRequest, opts ...grpc.CallOption) (*GetOrganizationInwayResponse, error) {
	out := new(GetOrganizationInwayResponse)
	err := c.cc.Invoke(ctx, "/registrationapi.DirectoryRegistration/GetOrganizationInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryRegistrationClient) ListInOutwayStatistics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInOutwayStatisticsResponse, error) {
	out := new(ListInOutwayStatisticsResponse)
	err := c.cc.Invoke(ctx, "/registrationapi.DirectoryRegistration/ListInOutwayStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryRegistrationServer is the server API for DirectoryRegistration service.
// All implementations must embed UnimplementedDirectoryRegistrationServer
// for forward compatibility
type DirectoryRegistrationServer interface {
	RegisterInway(context.Context, *RegisterInwayRequest) (*RegisterInwayResponse, error)
	ClearOrganizationInway(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ListServices(context.Context, *emptypb.Empty) (*ListServicesResponse, error)
	ListOrganizations(context.Context, *emptypb.Empty) (*ListOrganizationsResponse, error)
	GetOrganizationInway(context.Context, *GetOrganizationInwayRequest) (*GetOrganizationInwayResponse, error)
	ListInOutwayStatistics(context.Context, *emptypb.Empty) (*ListInOutwayStatisticsResponse, error)
	mustEmbedUnimplementedDirectoryRegistrationServer()
}

// UnimplementedDirectoryRegistrationServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryRegistrationServer struct {
}

func (UnimplementedDirectoryRegistrationServer) RegisterInway(context.Context, *RegisterInwayRequest) (*RegisterInwayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInway not implemented")
}
func (UnimplementedDirectoryRegistrationServer) ClearOrganizationInway(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearOrganizationInway not implemented")
}
func (UnimplementedDirectoryRegistrationServer) ListServices(context.Context, *emptypb.Empty) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedDirectoryRegistrationServer) ListOrganizations(context.Context, *emptypb.Empty) (*ListOrganizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizations not implemented")
}
func (UnimplementedDirectoryRegistrationServer) GetOrganizationInway(context.Context, *GetOrganizationInwayRequest) (*GetOrganizationInwayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationInway not implemented")
}
func (UnimplementedDirectoryRegistrationServer) ListInOutwayStatistics(context.Context, *emptypb.Empty) (*ListInOutwayStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInOutwayStatistics not implemented")
}
func (UnimplementedDirectoryRegistrationServer) mustEmbedUnimplementedDirectoryRegistrationServer() {}

// UnsafeDirectoryRegistrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryRegistrationServer will
// result in compilation errors.
type UnsafeDirectoryRegistrationServer interface {
	mustEmbedUnimplementedDirectoryRegistrationServer()
}

func RegisterDirectoryRegistrationServer(s grpc.ServiceRegistrar, srv DirectoryRegistrationServer) {
	s.RegisterService(&DirectoryRegistration_ServiceDesc, srv)
}

func _DirectoryRegistration_RegisterInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInwayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServer).RegisterInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryRegistration/RegisterInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServer).RegisterInway(ctx, req.(*RegisterInwayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryRegistration_ClearOrganizationInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServer).ClearOrganizationInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryRegistration/ClearOrganizationInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServer).ClearOrganizationInway(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryRegistration_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryRegistration/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServer).ListServices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryRegistration_ListOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServer).ListOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryRegistration/ListOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServer).ListOrganizations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryRegistration_GetOrganizationInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationInwayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServer).GetOrganizationInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryRegistration/GetOrganizationInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServer).GetOrganizationInway(ctx, req.(*GetOrganizationInwayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryRegistration_ListInOutwayStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryRegistrationServer).ListInOutwayStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryRegistration/ListInOutwayStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryRegistrationServer).ListInOutwayStatistics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectoryRegistration_ServiceDesc is the grpc.ServiceDesc for DirectoryRegistration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectoryRegistration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registrationapi.DirectoryRegistration",
	HandlerType: (*DirectoryRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInway",
			Handler:    _DirectoryRegistration_RegisterInway_Handler,
		},
		{
			MethodName: "ClearOrganizationInway",
			Handler:    _DirectoryRegistration_ClearOrganizationInway_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _DirectoryRegistration_ListServices_Handler,
		},
		{
			MethodName: "ListOrganizations",
			Handler:    _DirectoryRegistration_ListOrganizations_Handler,
		},
		{
			MethodName: "GetOrganizationInway",
			Handler:    _DirectoryRegistration_GetOrganizationInway_Handler,
		},
		{
			MethodName: "ListInOutwayStatistics",
			Handler:    _DirectoryRegistration_ListInOutwayStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directoryapi.proto",
}

// DirectoryInspectionClient is the client API for DirectoryInspection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryInspectionClient interface {
	RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error)
	ClearOrganizationInway(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error)
	ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrganizationsResponse, error)
	GetOrganizationInway(ctx context.Context, in *GetOrganizationInwayRequest, opts ...grpc.CallOption) (*GetOrganizationInwayResponse, error)
	ListInOutwayStatistics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInOutwayStatisticsResponse, error)
}

type directoryInspectionClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryInspectionClient(cc grpc.ClientConnInterface) DirectoryInspectionClient {
	return &directoryInspectionClient{cc}
}

func (c *directoryInspectionClient) RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error) {
	out := new(RegisterInwayResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.DirectoryInspection/RegisterInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryInspectionClient) ClearOrganizationInway(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/directoryapi.DirectoryInspection/ClearOrganizationInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryInspectionClient) ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.DirectoryInspection/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryInspectionClient) ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListOrganizationsResponse, error) {
	out := new(ListOrganizationsResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.DirectoryInspection/ListOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryInspectionClient) GetOrganizationInway(ctx context.Context, in *GetOrganizationInwayRequest, opts ...grpc.CallOption) (*GetOrganizationInwayResponse, error) {
	out := new(GetOrganizationInwayResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.DirectoryInspection/GetOrganizationInway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryInspectionClient) ListInOutwayStatistics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInOutwayStatisticsResponse, error) {
	out := new(ListInOutwayStatisticsResponse)
	err := c.cc.Invoke(ctx, "/directoryapi.DirectoryInspection/ListInOutwayStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryInspectionServer is the server API for DirectoryInspection service.
// All implementations must embed UnimplementedDirectoryInspectionServer
// for forward compatibility
type DirectoryInspectionServer interface {
	RegisterInway(context.Context, *RegisterInwayRequest) (*RegisterInwayResponse, error)
	ClearOrganizationInway(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ListServices(context.Context, *emptypb.Empty) (*ListServicesResponse, error)
	ListOrganizations(context.Context, *emptypb.Empty) (*ListOrganizationsResponse, error)
	GetOrganizationInway(context.Context, *GetOrganizationInwayRequest) (*GetOrganizationInwayResponse, error)
	ListInOutwayStatistics(context.Context, *emptypb.Empty) (*ListInOutwayStatisticsResponse, error)
	mustEmbedUnimplementedDirectoryInspectionServer()
}

// UnimplementedDirectoryInspectionServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryInspectionServer struct {
}

func (UnimplementedDirectoryInspectionServer) RegisterInway(context.Context, *RegisterInwayRequest) (*RegisterInwayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInway not implemented")
}
func (UnimplementedDirectoryInspectionServer) ClearOrganizationInway(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearOrganizationInway not implemented")
}
func (UnimplementedDirectoryInspectionServer) ListServices(context.Context, *emptypb.Empty) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedDirectoryInspectionServer) ListOrganizations(context.Context, *emptypb.Empty) (*ListOrganizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizations not implemented")
}
func (UnimplementedDirectoryInspectionServer) GetOrganizationInway(context.Context, *GetOrganizationInwayRequest) (*GetOrganizationInwayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationInway not implemented")
}
func (UnimplementedDirectoryInspectionServer) ListInOutwayStatistics(context.Context, *emptypb.Empty) (*ListInOutwayStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInOutwayStatistics not implemented")
}
func (UnimplementedDirectoryInspectionServer) mustEmbedUnimplementedDirectoryInspectionServer() {}

// UnsafeDirectoryInspectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryInspectionServer will
// result in compilation errors.
type UnsafeDirectoryInspectionServer interface {
	mustEmbedUnimplementedDirectoryInspectionServer()
}

func RegisterDirectoryInspectionServer(s grpc.ServiceRegistrar, srv DirectoryInspectionServer) {
	s.RegisterService(&DirectoryInspection_ServiceDesc, srv)
}

func _DirectoryInspection_RegisterInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInwayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryInspectionServer).RegisterInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryInspection/RegisterInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryInspectionServer).RegisterInway(ctx, req.(*RegisterInwayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryInspection_ClearOrganizationInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryInspectionServer).ClearOrganizationInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryInspection/ClearOrganizationInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryInspectionServer).ClearOrganizationInway(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryInspection_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryInspectionServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryInspection/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryInspectionServer).ListServices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryInspection_ListOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryInspectionServer).ListOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryInspection/ListOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryInspectionServer).ListOrganizations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryInspection_GetOrganizationInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationInwayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryInspectionServer).GetOrganizationInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryInspection/GetOrganizationInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryInspectionServer).GetOrganizationInway(ctx, req.(*GetOrganizationInwayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryInspection_ListInOutwayStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryInspectionServer).ListInOutwayStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.DirectoryInspection/ListInOutwayStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryInspectionServer).ListInOutwayStatistics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectoryInspection_ServiceDesc is the grpc.ServiceDesc for DirectoryInspection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectoryInspection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspectionapi.DirectoryInspection",
	HandlerType: (*DirectoryInspectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInway",
			Handler:    _DirectoryInspection_RegisterInway_Handler,
		},
		{
			MethodName: "ClearOrganizationInway",
			Handler:    _DirectoryInspection_ClearOrganizationInway_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _DirectoryInspection_ListServices_Handler,
		},
		{
			MethodName: "ListOrganizations",
			Handler:    _DirectoryInspection_ListOrganizations_Handler,
		},
		{
			MethodName: "GetOrganizationInway",
			Handler:    _DirectoryInspection_GetOrganizationInway_Handler,
		},
		{
			MethodName: "ListInOutwayStatistics",
			Handler:    _DirectoryInspection_ListInOutwayStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directoryapi.proto",
}
