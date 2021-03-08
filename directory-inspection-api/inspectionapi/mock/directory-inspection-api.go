// Code generated by MockGen. DO NOT EDIT.
// Source: inspectionapi.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	inspectionapi "go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// MockDirectoryInspectionClient is a mock of DirectoryInspectionClient interface
type MockDirectoryInspectionClient struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryInspectionClientMockRecorder
}

// MockDirectoryInspectionClientMockRecorder is the mock recorder for MockDirectoryInspectionClient
type MockDirectoryInspectionClientMockRecorder struct {
	mock *MockDirectoryInspectionClient
}

// NewMockDirectoryInspectionClient creates a new mock instance
func NewMockDirectoryInspectionClient(ctrl *gomock.Controller) *MockDirectoryInspectionClient {
	mock := &MockDirectoryInspectionClient{ctrl: ctrl}
	mock.recorder = &MockDirectoryInspectionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirectoryInspectionClient) EXPECT() *MockDirectoryInspectionClientMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockDirectoryInspectionClient) ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*inspectionapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*inspectionapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockDirectoryInspectionClientMockRecorder) ListServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockDirectoryInspectionClient)(nil).ListServices), varargs...)
}

// ListOrganizations mocks base method
func (m *MockDirectoryInspectionClient) ListOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*inspectionapi.ListOrganizationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizations", varargs...)
	ret0, _ := ret[0].(*inspectionapi.ListOrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations
func (mr *MockDirectoryInspectionClientMockRecorder) ListOrganizations(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockDirectoryInspectionClient)(nil).ListOrganizations), varargs...)
}

// GetOrganizationInway mocks base method
func (m *MockDirectoryInspectionClient) GetOrganizationInway(ctx context.Context, in *inspectionapi.GetOrganizationInwayRequest, opts ...grpc.CallOption) (*inspectionapi.GetOrganizationInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationInway", varargs...)
	ret0, _ := ret[0].(*inspectionapi.GetOrganizationInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInway indicates an expected call of GetOrganizationInway
func (mr *MockDirectoryInspectionClientMockRecorder) GetOrganizationInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInway", reflect.TypeOf((*MockDirectoryInspectionClient)(nil).GetOrganizationInway), varargs...)
}

// MockDirectoryInspectionServer is a mock of DirectoryInspectionServer interface
type MockDirectoryInspectionServer struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryInspectionServerMockRecorder
}

// MockDirectoryInspectionServerMockRecorder is the mock recorder for MockDirectoryInspectionServer
type MockDirectoryInspectionServerMockRecorder struct {
	mock *MockDirectoryInspectionServer
}

// NewMockDirectoryInspectionServer creates a new mock instance
func NewMockDirectoryInspectionServer(ctrl *gomock.Controller) *MockDirectoryInspectionServer {
	mock := &MockDirectoryInspectionServer{ctrl: ctrl}
	mock.recorder = &MockDirectoryInspectionServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirectoryInspectionServer) EXPECT() *MockDirectoryInspectionServerMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockDirectoryInspectionServer) ListServices(arg0 context.Context, arg1 *emptypb.Empty) (*inspectionapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", arg0, arg1)
	ret0, _ := ret[0].(*inspectionapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockDirectoryInspectionServerMockRecorder) ListServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockDirectoryInspectionServer)(nil).ListServices), arg0, arg1)
}

// ListOrganizations mocks base method
func (m *MockDirectoryInspectionServer) ListOrganizations(arg0 context.Context, arg1 *emptypb.Empty) (*inspectionapi.ListOrganizationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrganizations", arg0, arg1)
	ret0, _ := ret[0].(*inspectionapi.ListOrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations
func (mr *MockDirectoryInspectionServerMockRecorder) ListOrganizations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockDirectoryInspectionServer)(nil).ListOrganizations), arg0, arg1)
}

// GetOrganizationInway mocks base method
func (m *MockDirectoryInspectionServer) GetOrganizationInway(arg0 context.Context, arg1 *inspectionapi.GetOrganizationInwayRequest) (*inspectionapi.GetOrganizationInwayResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizationInway", arg0, arg1)
	ret0, _ := ret[0].(*inspectionapi.GetOrganizationInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInway indicates an expected call of GetOrganizationInway
func (mr *MockDirectoryInspectionServerMockRecorder) GetOrganizationInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInway", reflect.TypeOf((*MockDirectoryInspectionServer)(nil).GetOrganizationInway), arg0, arg1)
}
