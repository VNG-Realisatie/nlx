// Code generated by MockGen. DO NOT EDIT.
// Source: api/directoryapi_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

// MockDirectoryClient is a mock of DirectoryClient interface.
type MockDirectoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryClientMockRecorder
}

// MockDirectoryClientMockRecorder is the mock recorder for MockDirectoryClient.
type MockDirectoryClientMockRecorder struct {
	mock *MockDirectoryClient
}

// NewMockDirectoryClient creates a new mock instance.
func NewMockDirectoryClient(ctrl *gomock.Controller) *MockDirectoryClient {
	mock := &MockDirectoryClient{ctrl: ctrl}
	mock.recorder = &MockDirectoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectoryClient) EXPECT() *MockDirectoryClientMockRecorder {
	return m.recorder
}

// ClearOrganizationInway mocks base method.
func (m *MockDirectoryClient) ClearOrganizationInway(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClearOrganizationInway", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearOrganizationInway indicates an expected call of ClearOrganizationInway.
func (mr *MockDirectoryClientMockRecorder) ClearOrganizationInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearOrganizationInway", reflect.TypeOf((*MockDirectoryClient)(nil).ClearOrganizationInway), varargs...)
}

// GetOrganizationInway mocks base method.
func (m *MockDirectoryClient) GetOrganizationInway(ctx context.Context, in *directoryapi.GetOrganizationInwayRequest, opts ...grpc.CallOption) (*directoryapi.GetOrganizationInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationInway", varargs...)
	ret0, _ := ret[0].(*directoryapi.GetOrganizationInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInway indicates an expected call of GetOrganizationInway.
func (mr *MockDirectoryClientMockRecorder) GetOrganizationInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInway", reflect.TypeOf((*MockDirectoryClient)(nil).GetOrganizationInway), varargs...)
}

// ListInOutwayStatistics mocks base method.
func (m *MockDirectoryClient) ListInOutwayStatistics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*directoryapi.ListInOutwayStatisticsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInOutwayStatistics", varargs...)
	ret0, _ := ret[0].(*directoryapi.ListInOutwayStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInOutwayStatistics indicates an expected call of ListInOutwayStatistics.
func (mr *MockDirectoryClientMockRecorder) ListInOutwayStatistics(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInOutwayStatistics", reflect.TypeOf((*MockDirectoryClient)(nil).ListInOutwayStatistics), varargs...)
}

// ListOrganizations mocks base method.
func (m *MockDirectoryClient) ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*directoryapi.ListOrganizationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizations", varargs...)
	ret0, _ := ret[0].(*directoryapi.ListOrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations.
func (mr *MockDirectoryClientMockRecorder) ListOrganizations(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockDirectoryClient)(nil).ListOrganizations), varargs...)
}

// ListServices mocks base method.
func (m *MockDirectoryClient) ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*directoryapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*directoryapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockDirectoryClientMockRecorder) ListServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockDirectoryClient)(nil).ListServices), varargs...)
}

// RegisterInway mocks base method.
func (m *MockDirectoryClient) RegisterInway(ctx context.Context, in *directoryapi.RegisterInwayRequest, opts ...grpc.CallOption) (*directoryapi.RegisterInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterInway", varargs...)
	ret0, _ := ret[0].(*directoryapi.RegisterInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInway indicates an expected call of RegisterInway.
func (mr *MockDirectoryClientMockRecorder) RegisterInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInway", reflect.TypeOf((*MockDirectoryClient)(nil).RegisterInway), varargs...)
}

// MockDirectoryServer is a mock of DirectoryServer interface.
type MockDirectoryServer struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryServerMockRecorder
}

// MockDirectoryServerMockRecorder is the mock recorder for MockDirectoryServer.
type MockDirectoryServerMockRecorder struct {
	mock *MockDirectoryServer
}

// NewMockDirectoryServer creates a new mock instance.
func NewMockDirectoryServer(ctrl *gomock.Controller) *MockDirectoryServer {
	mock := &MockDirectoryServer{ctrl: ctrl}
	mock.recorder = &MockDirectoryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectoryServer) EXPECT() *MockDirectoryServerMockRecorder {
	return m.recorder
}

// ClearOrganizationInway mocks base method.
func (m *MockDirectoryServer) ClearOrganizationInway(arg0 context.Context, arg1 *emptypb.Empty) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearOrganizationInway", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearOrganizationInway indicates an expected call of ClearOrganizationInway.
func (mr *MockDirectoryServerMockRecorder) ClearOrganizationInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearOrganizationInway", reflect.TypeOf((*MockDirectoryServer)(nil).ClearOrganizationInway), arg0, arg1)
}

// GetOrganizationInway mocks base method.
func (m *MockDirectoryServer) GetOrganizationInway(arg0 context.Context, arg1 *directoryapi.GetOrganizationInwayRequest) (*directoryapi.GetOrganizationInwayResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizationInway", arg0, arg1)
	ret0, _ := ret[0].(*directoryapi.GetOrganizationInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInway indicates an expected call of GetOrganizationInway.
func (mr *MockDirectoryServerMockRecorder) GetOrganizationInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInway", reflect.TypeOf((*MockDirectoryServer)(nil).GetOrganizationInway), arg0, arg1)
}

// ListInOutwayStatistics mocks base method.
func (m *MockDirectoryServer) ListInOutwayStatistics(arg0 context.Context, arg1 *emptypb.Empty) (*directoryapi.ListInOutwayStatisticsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInOutwayStatistics", arg0, arg1)
	ret0, _ := ret[0].(*directoryapi.ListInOutwayStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInOutwayStatistics indicates an expected call of ListInOutwayStatistics.
func (mr *MockDirectoryServerMockRecorder) ListInOutwayStatistics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInOutwayStatistics", reflect.TypeOf((*MockDirectoryServer)(nil).ListInOutwayStatistics), arg0, arg1)
}

// ListOrganizations mocks base method.
func (m *MockDirectoryServer) ListOrganizations(arg0 context.Context, arg1 *emptypb.Empty) (*directoryapi.ListOrganizationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrganizations", arg0, arg1)
	ret0, _ := ret[0].(*directoryapi.ListOrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations.
func (mr *MockDirectoryServerMockRecorder) ListOrganizations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockDirectoryServer)(nil).ListOrganizations), arg0, arg1)
}

// ListServices mocks base method.
func (m *MockDirectoryServer) ListServices(arg0 context.Context, arg1 *emptypb.Empty) (*directoryapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", arg0, arg1)
	ret0, _ := ret[0].(*directoryapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockDirectoryServerMockRecorder) ListServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockDirectoryServer)(nil).ListServices), arg0, arg1)
}

// RegisterInway mocks base method.
func (m *MockDirectoryServer) RegisterInway(arg0 context.Context, arg1 *directoryapi.RegisterInwayRequest) (*directoryapi.RegisterInwayResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInway", arg0, arg1)
	ret0, _ := ret[0].(*directoryapi.RegisterInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInway indicates an expected call of RegisterInway.
func (mr *MockDirectoryServerMockRecorder) RegisterInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInway", reflect.TypeOf((*MockDirectoryServer)(nil).RegisterInway), arg0, arg1)
}

// mustEmbedUnimplementedDirectoryServer mocks base method.
func (m *MockDirectoryServer) mustEmbedUnimplementedDirectoryServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDirectoryServer")
}

// mustEmbedUnimplementedDirectoryServer indicates an expected call of mustEmbedUnimplementedDirectoryServer.
func (mr *MockDirectoryServerMockRecorder) mustEmbedUnimplementedDirectoryServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDirectoryServer", reflect.TypeOf((*MockDirectoryServer)(nil).mustEmbedUnimplementedDirectoryServer))
}

// MockUnsafeDirectoryServer is a mock of UnsafeDirectoryServer interface.
type MockUnsafeDirectoryServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeDirectoryServerMockRecorder
}

// MockUnsafeDirectoryServerMockRecorder is the mock recorder for MockUnsafeDirectoryServer.
type MockUnsafeDirectoryServerMockRecorder struct {
	mock *MockUnsafeDirectoryServer
}

// NewMockUnsafeDirectoryServer creates a new mock instance.
func NewMockUnsafeDirectoryServer(ctrl *gomock.Controller) *MockUnsafeDirectoryServer {
	mock := &MockUnsafeDirectoryServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeDirectoryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeDirectoryServer) EXPECT() *MockUnsafeDirectoryServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedDirectoryServer mocks base method.
func (m *MockUnsafeDirectoryServer) mustEmbedUnimplementedDirectoryServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDirectoryServer")
}

// mustEmbedUnimplementedDirectoryServer indicates an expected call of mustEmbedUnimplementedDirectoryServer.
func (mr *MockUnsafeDirectoryServerMockRecorder) mustEmbedUnimplementedDirectoryServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDirectoryServer", reflect.TypeOf((*MockUnsafeDirectoryServer)(nil).mustEmbedUnimplementedDirectoryServer))
}