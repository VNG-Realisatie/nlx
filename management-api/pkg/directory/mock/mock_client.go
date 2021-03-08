// Code generated by MockGen. DO NOT EDIT.
// Source: go.nlx.io/nlx/management-api/pkg/directory (interfaces: Client)

// Package mock_directory is a generated GoMock package.
package mock_directory

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	inspectionapi "go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	registrationapi "go.nlx.io/nlx/directory-registration-api/registrationapi"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ClearOrganizationInway mocks base method
func (m *MockClient) ClearOrganizationInway(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClearOrganizationInway", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearOrganizationInway indicates an expected call of ClearOrganizationInway
func (mr *MockClientMockRecorder) ClearOrganizationInway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearOrganizationInway", reflect.TypeOf((*MockClient)(nil).ClearOrganizationInway), varargs...)
}

// GetOrganizationInway mocks base method
func (m *MockClient) GetOrganizationInway(arg0 context.Context, arg1 *inspectionapi.GetOrganizationInwayRequest, arg2 ...grpc.CallOption) (*inspectionapi.GetOrganizationInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationInway", varargs...)
	ret0, _ := ret[0].(*inspectionapi.GetOrganizationInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInway indicates an expected call of GetOrganizationInway
func (mr *MockClientMockRecorder) GetOrganizationInway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInway", reflect.TypeOf((*MockClient)(nil).GetOrganizationInway), varargs...)
}

// ListInOutwayStatistics mocks base method
func (m *MockClient) ListInOutwayStatistics(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*inspectionapi.ListInOutwayStatisticsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInOutwayStatistics", varargs...)
	ret0, _ := ret[0].(*inspectionapi.ListInOutwayStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInOutwayStatistics indicates an expected call of ListInOutwayStatistics
func (mr *MockClientMockRecorder) ListInOutwayStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInOutwayStatistics", reflect.TypeOf((*MockClient)(nil).ListInOutwayStatistics), varargs...)
}

// ListOrganizations mocks base method
func (m *MockClient) ListOrganizations(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*inspectionapi.ListOrganizationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizations", varargs...)
	ret0, _ := ret[0].(*inspectionapi.ListOrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations
func (mr *MockClientMockRecorder) ListOrganizations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockClient)(nil).ListOrganizations), varargs...)
}

// ListServices mocks base method
func (m *MockClient) ListServices(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*inspectionapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*inspectionapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockClient)(nil).ListServices), varargs...)
}

// RegisterInway mocks base method
func (m *MockClient) RegisterInway(arg0 context.Context, arg1 *registrationapi.RegisterInwayRequest, arg2 ...grpc.CallOption) (*registrationapi.RegisterInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterInway", varargs...)
	ret0, _ := ret[0].(*registrationapi.RegisterInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInway indicates an expected call of RegisterInway
func (mr *MockClientMockRecorder) RegisterInway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInway", reflect.TypeOf((*MockClient)(nil).RegisterInway), varargs...)
}

// SetInsightConfiguration mocks base method
func (m *MockClient) SetInsightConfiguration(arg0 context.Context, arg1 *registrationapi.SetInsightConfigurationRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetInsightConfiguration", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetInsightConfiguration indicates an expected call of SetInsightConfiguration
func (mr *MockClientMockRecorder) SetInsightConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInsightConfiguration", reflect.TypeOf((*MockClient)(nil).SetInsightConfiguration), varargs...)
}

// SetOrganizationInway mocks base method
func (m *MockClient) SetOrganizationInway(arg0 context.Context, arg1 *registrationapi.SetOrganizationInwayRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetOrganizationInway", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOrganizationInway indicates an expected call of SetOrganizationInway
func (mr *MockClientMockRecorder) SetOrganizationInway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrganizationInway", reflect.TypeOf((*MockClient)(nil).SetOrganizationInway), varargs...)
}
