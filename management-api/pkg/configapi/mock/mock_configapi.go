// Code generated by MockGen. DO NOT EDIT.
// Source: configapi.pb.go

// Package mock_configapi is a generated GoMock package.
package mock_configapi

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	configapi "go.nlx.io/nlx/management-api/pkg/configapi"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockConfigApiClient is a mock of ConfigApiClient interface
type MockConfigApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockConfigApiClientMockRecorder
}

// MockConfigApiClientMockRecorder is the mock recorder for MockConfigApiClient
type MockConfigApiClientMockRecorder struct {
	mock *MockConfigApiClient
}

// NewMockConfigApiClient creates a new mock instance
func NewMockConfigApiClient(ctrl *gomock.Controller) *MockConfigApiClient {
	mock := &MockConfigApiClient{ctrl: ctrl}
	mock.recorder = &MockConfigApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigApiClient) EXPECT() *MockConfigApiClientMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockConfigApiClient) ListServices(ctx context.Context, in *configapi.ListServicesRequest, opts ...grpc.CallOption) (*configapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*configapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockConfigApiClientMockRecorder) ListServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockConfigApiClient)(nil).ListServices), varargs...)
}

// GetService mocks base method
func (m *MockConfigApiClient) GetService(ctx context.Context, in *configapi.GetServiceRequest, opts ...grpc.CallOption) (*configapi.Service, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetService", varargs...)
	ret0, _ := ret[0].(*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockConfigApiClientMockRecorder) GetService(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockConfigApiClient)(nil).GetService), varargs...)
}

// CreateService mocks base method
func (m *MockConfigApiClient) CreateService(ctx context.Context, in *configapi.Service, opts ...grpc.CallOption) (*configapi.Service, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateService", varargs...)
	ret0, _ := ret[0].(*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService
func (mr *MockConfigApiClientMockRecorder) CreateService(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockConfigApiClient)(nil).CreateService), varargs...)
}

// UpdateService mocks base method
func (m *MockConfigApiClient) UpdateService(ctx context.Context, in *configapi.UpdateServiceRequest, opts ...grpc.CallOption) (*configapi.Service, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateService", varargs...)
	ret0, _ := ret[0].(*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockConfigApiClientMockRecorder) UpdateService(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockConfigApiClient)(nil).UpdateService), varargs...)
}

// DeleteService mocks base method
func (m *MockConfigApiClient) DeleteService(ctx context.Context, in *configapi.DeleteServiceRequest, opts ...grpc.CallOption) (*configapi.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteService", varargs...)
	ret0, _ := ret[0].(*configapi.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockConfigApiClientMockRecorder) DeleteService(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockConfigApiClient)(nil).DeleteService), varargs...)
}

// ListInways mocks base method
func (m *MockConfigApiClient) ListInways(ctx context.Context, in *configapi.ListInwaysRequest, opts ...grpc.CallOption) (*configapi.ListInwaysResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInways", varargs...)
	ret0, _ := ret[0].(*configapi.ListInwaysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInways indicates an expected call of ListInways
func (mr *MockConfigApiClientMockRecorder) ListInways(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInways", reflect.TypeOf((*MockConfigApiClient)(nil).ListInways), varargs...)
}

// GetInway mocks base method
func (m *MockConfigApiClient) GetInway(ctx context.Context, in *configapi.GetInwayRequest, opts ...grpc.CallOption) (*configapi.Inway, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInway", varargs...)
	ret0, _ := ret[0].(*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInway indicates an expected call of GetInway
func (mr *MockConfigApiClientMockRecorder) GetInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInway", reflect.TypeOf((*MockConfigApiClient)(nil).GetInway), varargs...)
}

// CreateInway mocks base method
func (m *MockConfigApiClient) CreateInway(ctx context.Context, in *configapi.Inway, opts ...grpc.CallOption) (*configapi.Inway, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateInway", varargs...)
	ret0, _ := ret[0].(*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInway indicates an expected call of CreateInway
func (mr *MockConfigApiClientMockRecorder) CreateInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInway", reflect.TypeOf((*MockConfigApiClient)(nil).CreateInway), varargs...)
}

// UpdateInway mocks base method
func (m *MockConfigApiClient) UpdateInway(ctx context.Context, in *configapi.UpdateInwayRequest, opts ...grpc.CallOption) (*configapi.Inway, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateInway", varargs...)
	ret0, _ := ret[0].(*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInway indicates an expected call of UpdateInway
func (mr *MockConfigApiClientMockRecorder) UpdateInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInway", reflect.TypeOf((*MockConfigApiClient)(nil).UpdateInway), varargs...)
}

// DeleteInway mocks base method
func (m *MockConfigApiClient) DeleteInway(ctx context.Context, in *configapi.DeleteInwayRequest, opts ...grpc.CallOption) (*configapi.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteInway", varargs...)
	ret0, _ := ret[0].(*configapi.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInway indicates an expected call of DeleteInway
func (mr *MockConfigApiClientMockRecorder) DeleteInway(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInway", reflect.TypeOf((*MockConfigApiClient)(nil).DeleteInway), varargs...)
}

// PutInsightConfiguration mocks base method
func (m *MockConfigApiClient) PutInsightConfiguration(ctx context.Context, in *configapi.InsightConfiguration, opts ...grpc.CallOption) (*configapi.InsightConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutInsightConfiguration", varargs...)
	ret0, _ := ret[0].(*configapi.InsightConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutInsightConfiguration indicates an expected call of PutInsightConfiguration
func (mr *MockConfigApiClientMockRecorder) PutInsightConfiguration(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutInsightConfiguration", reflect.TypeOf((*MockConfigApiClient)(nil).PutInsightConfiguration), varargs...)
}

// GetInsightConfiguration mocks base method
func (m *MockConfigApiClient) GetInsightConfiguration(ctx context.Context, in *configapi.Empty, opts ...grpc.CallOption) (*configapi.InsightConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsightConfiguration", varargs...)
	ret0, _ := ret[0].(*configapi.InsightConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightConfiguration indicates an expected call of GetInsightConfiguration
func (mr *MockConfigApiClientMockRecorder) GetInsightConfiguration(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightConfiguration", reflect.TypeOf((*MockConfigApiClient)(nil).GetInsightConfiguration), varargs...)
}

// ListOutgoingAccessRequests mocks base method
func (m *MockConfigApiClient) ListOutgoingAccessRequests(ctx context.Context, in *configapi.ListOutgoingAccessRequestsRequest, opts ...grpc.CallOption) (*configapi.ListOutgoingAccessRequestsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOutgoingAccessRequests", varargs...)
	ret0, _ := ret[0].(*configapi.ListOutgoingAccessRequestsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOutgoingAccessRequests indicates an expected call of ListOutgoingAccessRequests
func (mr *MockConfigApiClientMockRecorder) ListOutgoingAccessRequests(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutgoingAccessRequests", reflect.TypeOf((*MockConfigApiClient)(nil).ListOutgoingAccessRequests), varargs...)
}

// CreateAccessRequest mocks base method
func (m *MockConfigApiClient) CreateAccessRequest(ctx context.Context, in *configapi.CreateAccessRequestRequest, opts ...grpc.CallOption) (*configapi.AccessRequest, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccessRequest", varargs...)
	ret0, _ := ret[0].(*configapi.AccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessRequest indicates an expected call of CreateAccessRequest
func (mr *MockConfigApiClientMockRecorder) CreateAccessRequest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessRequest", reflect.TypeOf((*MockConfigApiClient)(nil).CreateAccessRequest), varargs...)
}

// MockConfigApiServer is a mock of ConfigApiServer interface
type MockConfigApiServer struct {
	ctrl     *gomock.Controller
	recorder *MockConfigApiServerMockRecorder
}

// MockConfigApiServerMockRecorder is the mock recorder for MockConfigApiServer
type MockConfigApiServerMockRecorder struct {
	mock *MockConfigApiServer
}

// NewMockConfigApiServer creates a new mock instance
func NewMockConfigApiServer(ctrl *gomock.Controller) *MockConfigApiServer {
	mock := &MockConfigApiServer{ctrl: ctrl}
	mock.recorder = &MockConfigApiServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigApiServer) EXPECT() *MockConfigApiServerMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockConfigApiServer) ListServices(arg0 context.Context, arg1 *configapi.ListServicesRequest) (*configapi.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", arg0, arg1)
	ret0, _ := ret[0].(*configapi.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockConfigApiServerMockRecorder) ListServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockConfigApiServer)(nil).ListServices), arg0, arg1)
}

// GetService mocks base method
func (m *MockConfigApiServer) GetService(arg0 context.Context, arg1 *configapi.GetServiceRequest) (*configapi.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockConfigApiServerMockRecorder) GetService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockConfigApiServer)(nil).GetService), arg0, arg1)
}

// CreateService mocks base method
func (m *MockConfigApiServer) CreateService(arg0 context.Context, arg1 *configapi.Service) (*configapi.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService
func (mr *MockConfigApiServerMockRecorder) CreateService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockConfigApiServer)(nil).CreateService), arg0, arg1)
}

// UpdateService mocks base method
func (m *MockConfigApiServer) UpdateService(arg0 context.Context, arg1 *configapi.UpdateServiceRequest) (*configapi.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockConfigApiServerMockRecorder) UpdateService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockConfigApiServer)(nil).UpdateService), arg0, arg1)
}

// DeleteService mocks base method
func (m *MockConfigApiServer) DeleteService(arg0 context.Context, arg1 *configapi.DeleteServiceRequest) (*configapi.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockConfigApiServerMockRecorder) DeleteService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockConfigApiServer)(nil).DeleteService), arg0, arg1)
}

// ListInways mocks base method
func (m *MockConfigApiServer) ListInways(arg0 context.Context, arg1 *configapi.ListInwaysRequest) (*configapi.ListInwaysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInways", arg0, arg1)
	ret0, _ := ret[0].(*configapi.ListInwaysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInways indicates an expected call of ListInways
func (mr *MockConfigApiServerMockRecorder) ListInways(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInways", reflect.TypeOf((*MockConfigApiServer)(nil).ListInways), arg0, arg1)
}

// GetInway mocks base method
func (m *MockConfigApiServer) GetInway(arg0 context.Context, arg1 *configapi.GetInwayRequest) (*configapi.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInway", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInway indicates an expected call of GetInway
func (mr *MockConfigApiServerMockRecorder) GetInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInway", reflect.TypeOf((*MockConfigApiServer)(nil).GetInway), arg0, arg1)
}

// CreateInway mocks base method
func (m *MockConfigApiServer) CreateInway(arg0 context.Context, arg1 *configapi.Inway) (*configapi.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInway", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInway indicates an expected call of CreateInway
func (mr *MockConfigApiServerMockRecorder) CreateInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInway", reflect.TypeOf((*MockConfigApiServer)(nil).CreateInway), arg0, arg1)
}

// UpdateInway mocks base method
func (m *MockConfigApiServer) UpdateInway(arg0 context.Context, arg1 *configapi.UpdateInwayRequest) (*configapi.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInway", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInway indicates an expected call of UpdateInway
func (mr *MockConfigApiServerMockRecorder) UpdateInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInway", reflect.TypeOf((*MockConfigApiServer)(nil).UpdateInway), arg0, arg1)
}

// DeleteInway mocks base method
func (m *MockConfigApiServer) DeleteInway(arg0 context.Context, arg1 *configapi.DeleteInwayRequest) (*configapi.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInway", arg0, arg1)
	ret0, _ := ret[0].(*configapi.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInway indicates an expected call of DeleteInway
func (mr *MockConfigApiServerMockRecorder) DeleteInway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInway", reflect.TypeOf((*MockConfigApiServer)(nil).DeleteInway), arg0, arg1)
}

// PutInsightConfiguration mocks base method
func (m *MockConfigApiServer) PutInsightConfiguration(arg0 context.Context, arg1 *configapi.InsightConfiguration) (*configapi.InsightConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutInsightConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*configapi.InsightConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutInsightConfiguration indicates an expected call of PutInsightConfiguration
func (mr *MockConfigApiServerMockRecorder) PutInsightConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutInsightConfiguration", reflect.TypeOf((*MockConfigApiServer)(nil).PutInsightConfiguration), arg0, arg1)
}

// GetInsightConfiguration mocks base method
func (m *MockConfigApiServer) GetInsightConfiguration(arg0 context.Context, arg1 *configapi.Empty) (*configapi.InsightConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInsightConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*configapi.InsightConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightConfiguration indicates an expected call of GetInsightConfiguration
func (mr *MockConfigApiServerMockRecorder) GetInsightConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightConfiguration", reflect.TypeOf((*MockConfigApiServer)(nil).GetInsightConfiguration), arg0, arg1)
}

// ListOutgoingAccessRequests mocks base method
func (m *MockConfigApiServer) ListOutgoingAccessRequests(arg0 context.Context, arg1 *configapi.ListOutgoingAccessRequestsRequest) (*configapi.ListOutgoingAccessRequestsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutgoingAccessRequests", arg0, arg1)
	ret0, _ := ret[0].(*configapi.ListOutgoingAccessRequestsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOutgoingAccessRequests indicates an expected call of ListOutgoingAccessRequests
func (mr *MockConfigApiServerMockRecorder) ListOutgoingAccessRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutgoingAccessRequests", reflect.TypeOf((*MockConfigApiServer)(nil).ListOutgoingAccessRequests), arg0, arg1)
}

// CreateAccessRequest mocks base method
func (m *MockConfigApiServer) CreateAccessRequest(arg0 context.Context, arg1 *configapi.CreateAccessRequestRequest) (*configapi.AccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessRequest", arg0, arg1)
	ret0, _ := ret[0].(*configapi.AccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessRequest indicates an expected call of CreateAccessRequest
func (mr *MockConfigApiServerMockRecorder) CreateAccessRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessRequest", reflect.TypeOf((*MockConfigApiServer)(nil).CreateAccessRequest), arg0, arg1)
}
