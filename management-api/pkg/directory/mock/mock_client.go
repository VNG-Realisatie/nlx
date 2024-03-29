// Code generated by MockGen. DO NOT EDIT.
// Source: go.nlx.io/nlx/management-api/pkg/directory (interfaces: Client)

// Package mock_directory is a generated GoMock package.
package mock_directory

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ClearOrganizationInway mocks base method.
func (m *MockClient) ClearOrganizationInway(arg0 context.Context, arg1 *directoryapi.ClearOrganizationInwayRequest, arg2 ...grpc.CallOption) (*directoryapi.ClearOrganizationInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClearOrganizationInway", varargs...)
	ret0, _ := ret[0].(*directoryapi.ClearOrganizationInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearOrganizationInway indicates an expected call of ClearOrganizationInway.
func (mr *MockClientMockRecorder) ClearOrganizationInway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearOrganizationInway", reflect.TypeOf((*MockClient)(nil).ClearOrganizationInway), varargs...)
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// GetOrganizationInway mocks base method.
func (m *MockClient) GetOrganizationInway(arg0 context.Context, arg1 *directoryapi.GetOrganizationInwayRequest, arg2 ...grpc.CallOption) (*directoryapi.GetOrganizationInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationInway", varargs...)
	ret0, _ := ret[0].(*directoryapi.GetOrganizationInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInway indicates an expected call of GetOrganizationInway.
func (mr *MockClientMockRecorder) GetOrganizationInway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInway", reflect.TypeOf((*MockClient)(nil).GetOrganizationInway), varargs...)
}

// GetOrganizationInwayProxyAddress mocks base method.
func (m *MockClient) GetOrganizationInwayProxyAddress(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizationInwayProxyAddress", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInwayProxyAddress indicates an expected call of GetOrganizationInwayProxyAddress.
func (mr *MockClientMockRecorder) GetOrganizationInwayProxyAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInwayProxyAddress", reflect.TypeOf((*MockClient)(nil).GetOrganizationInwayProxyAddress), arg0, arg1)
}

// GetOrganizationManagementAPIProxyAddress mocks base method.
func (m *MockClient) GetOrganizationManagementAPIProxyAddress(arg0 context.Context, arg1 *directoryapi.GetOrganizationManagementAPIProxyAddressRequest, arg2 ...grpc.CallOption) (*directoryapi.GetOrganizationManagementAPIProxyAddressResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationManagementAPIProxyAddress", varargs...)
	ret0, _ := ret[0].(*directoryapi.GetOrganizationManagementAPIProxyAddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationManagementAPIProxyAddress indicates an expected call of GetOrganizationManagementAPIProxyAddress.
func (mr *MockClientMockRecorder) GetOrganizationManagementAPIProxyAddress(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationManagementAPIProxyAddress", reflect.TypeOf((*MockClient)(nil).GetOrganizationManagementAPIProxyAddress), varargs...)
}

// GetTermsOfService mocks base method.
func (m *MockClient) GetTermsOfService(arg0 context.Context, arg1 *directoryapi.GetTermsOfServiceRequest, arg2 ...grpc.CallOption) (*directoryapi.GetTermsOfServiceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTermsOfService", varargs...)
	ret0, _ := ret[0].(*directoryapi.GetTermsOfServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTermsOfService indicates an expected call of GetTermsOfService.
func (mr *MockClientMockRecorder) GetTermsOfService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTermsOfService", reflect.TypeOf((*MockClient)(nil).GetTermsOfService), varargs...)
}

// GetVersion mocks base method.
func (m *MockClient) GetVersion(arg0 context.Context, arg1 *directoryapi.GetVersionRequest, arg2 ...grpc.CallOption) (*directoryapi.GetVersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVersion", varargs...)
	ret0, _ := ret[0].(*directoryapi.GetVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockClientMockRecorder) GetVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockClient)(nil).GetVersion), varargs...)
}

// ListInOutwayStatistics mocks base method.
func (m *MockClient) ListInOutwayStatistics(arg0 context.Context, arg1 *directoryapi.ListInOutwayStatisticsRequest, arg2 ...grpc.CallOption) (*directoryapi.ListInOutwayStatisticsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInOutwayStatistics", varargs...)
	ret0, _ := ret[0].(*directoryapi.ListInOutwayStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInOutwayStatistics indicates an expected call of ListInOutwayStatistics.
func (mr *MockClientMockRecorder) ListInOutwayStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInOutwayStatistics", reflect.TypeOf((*MockClient)(nil).ListInOutwayStatistics), varargs...)
}

// ListOrganizations mocks base method.
func (m *MockClient) ListOrganizations(arg0 context.Context, arg1 *directoryapi.ListOrganizationsRequest, arg2 ...grpc.CallOption) (*directoryapi.ListOrganizationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizations", varargs...)
	ret0, _ := ret[0].(*directoryapi.ListOrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations.
func (mr *MockClientMockRecorder) ListOrganizations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockClient)(nil).ListOrganizations), varargs...)
}

// ListParticipants mocks base method.
func (m *MockClient) ListParticipants(arg0 context.Context, arg1 *directoryapi.ListParticipantsRequest, arg2 ...grpc.CallOption) (*directoryapi.ListParticipantsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListParticipants", varargs...)
	ret0, _ := ret[0].(*directoryapi.ListParticipantsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListParticipants indicates an expected call of ListParticipants.
func (mr *MockClientMockRecorder) ListParticipants(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListParticipants", reflect.TypeOf((*MockClient)(nil).ListParticipants), varargs...)
}

// ListServices mocks base method.
func (m *MockClient) ListServices(arg0 context.Context, arg1 *directoryapi.ListServicesRequest, arg2 ...grpc.CallOption) (*directoryapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*directoryapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockClient)(nil).ListServices), varargs...)
}

// RegisterInway mocks base method.
func (m *MockClient) RegisterInway(arg0 context.Context, arg1 *directoryapi.RegisterInwayRequest, arg2 ...grpc.CallOption) (*directoryapi.RegisterInwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterInway", varargs...)
	ret0, _ := ret[0].(*directoryapi.RegisterInwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInway indicates an expected call of RegisterInway.
func (mr *MockClientMockRecorder) RegisterInway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInway", reflect.TypeOf((*MockClient)(nil).RegisterInway), varargs...)
}

// RegisterOutway mocks base method.
func (m *MockClient) RegisterOutway(arg0 context.Context, arg1 *directoryapi.RegisterOutwayRequest, arg2 ...grpc.CallOption) (*directoryapi.RegisterOutwayResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterOutway", varargs...)
	ret0, _ := ret[0].(*directoryapi.RegisterOutwayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterOutway indicates an expected call of RegisterOutway.
func (mr *MockClientMockRecorder) RegisterOutway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOutway", reflect.TypeOf((*MockClient)(nil).RegisterOutway), varargs...)
}

// SetOrganizationContactDetails mocks base method.
func (m *MockClient) SetOrganizationContactDetails(arg0 context.Context, arg1 *directoryapi.SetOrganizationContactDetailsRequest, arg2 ...grpc.CallOption) (*directoryapi.SetOrganizationContactDetailsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetOrganizationContactDetails", varargs...)
	ret0, _ := ret[0].(*directoryapi.SetOrganizationContactDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOrganizationContactDetails indicates an expected call of SetOrganizationContactDetails.
func (mr *MockClientMockRecorder) SetOrganizationContactDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrganizationContactDetails", reflect.TypeOf((*MockClient)(nil).SetOrganizationContactDetails), varargs...)
}
