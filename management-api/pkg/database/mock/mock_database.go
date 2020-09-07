// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	database "go.nlx.io/nlx/management-api/pkg/database"
	reflect "reflect"
)

// MockConfigDatabase is a mock of ConfigDatabase interface
type MockConfigDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockConfigDatabaseMockRecorder
}

// MockConfigDatabaseMockRecorder is the mock recorder for MockConfigDatabase
type MockConfigDatabaseMockRecorder struct {
	mock *MockConfigDatabase
}

// NewMockConfigDatabase creates a new mock instance
func NewMockConfigDatabase(ctrl *gomock.Controller) *MockConfigDatabase {
	mock := &MockConfigDatabase{ctrl: ctrl}
	mock.recorder = &MockConfigDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigDatabase) EXPECT() *MockConfigDatabaseMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockConfigDatabase) ListServices(ctx context.Context) ([]*database.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", ctx)
	ret0, _ := ret[0].([]*database.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockConfigDatabaseMockRecorder) ListServices(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockConfigDatabase)(nil).ListServices), ctx)
}

// GetService mocks base method
func (m *MockConfigDatabase) GetService(ctx context.Context, name string) (*database.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", ctx, name)
	ret0, _ := ret[0].(*database.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockConfigDatabaseMockRecorder) GetService(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockConfigDatabase)(nil).GetService), ctx, name)
}

// CreateService mocks base method
func (m *MockConfigDatabase) CreateService(ctx context.Context, service *database.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", ctx, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService
func (mr *MockConfigDatabaseMockRecorder) CreateService(ctx, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockConfigDatabase)(nil).CreateService), ctx, service)
}

// UpdateService mocks base method
func (m *MockConfigDatabase) UpdateService(ctx context.Context, name string, service *database.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", ctx, name, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockConfigDatabaseMockRecorder) UpdateService(ctx, name, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateService), ctx, name, service)
}

// DeleteService mocks base method
func (m *MockConfigDatabase) DeleteService(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockConfigDatabaseMockRecorder) DeleteService(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockConfigDatabase)(nil).DeleteService), ctx, name)
}

// ListInways mocks base method
func (m *MockConfigDatabase) ListInways(ctx context.Context) ([]*database.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInways", ctx)
	ret0, _ := ret[0].([]*database.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInways indicates an expected call of ListInways
func (mr *MockConfigDatabaseMockRecorder) ListInways(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInways", reflect.TypeOf((*MockConfigDatabase)(nil).ListInways), ctx)
}

// GetInway mocks base method
func (m *MockConfigDatabase) GetInway(ctx context.Context, name string) (*database.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInway", ctx, name)
	ret0, _ := ret[0].(*database.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInway indicates an expected call of GetInway
func (mr *MockConfigDatabaseMockRecorder) GetInway(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInway", reflect.TypeOf((*MockConfigDatabase)(nil).GetInway), ctx, name)
}

// CreateInway mocks base method
func (m *MockConfigDatabase) CreateInway(ctx context.Context, inway *database.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInway", ctx, inway)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInway indicates an expected call of CreateInway
func (mr *MockConfigDatabaseMockRecorder) CreateInway(ctx, inway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInway", reflect.TypeOf((*MockConfigDatabase)(nil).CreateInway), ctx, inway)
}

// UpdateInway mocks base method
func (m *MockConfigDatabase) UpdateInway(ctx context.Context, name string, inway *database.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInway", ctx, name, inway)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInway indicates an expected call of UpdateInway
func (mr *MockConfigDatabaseMockRecorder) UpdateInway(ctx, name, inway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInway", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateInway), ctx, name, inway)
}

// DeleteInway mocks base method
func (m *MockConfigDatabase) DeleteInway(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInway", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInway indicates an expected call of DeleteInway
func (mr *MockConfigDatabaseMockRecorder) DeleteInway(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInway", reflect.TypeOf((*MockConfigDatabase)(nil).DeleteInway), ctx, name)
}

// PutInsightConfiguration mocks base method
func (m *MockConfigDatabase) PutInsightConfiguration(ctx context.Context, configuration *database.InsightConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutInsightConfiguration", ctx, configuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutInsightConfiguration indicates an expected call of PutInsightConfiguration
func (mr *MockConfigDatabaseMockRecorder) PutInsightConfiguration(ctx, configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutInsightConfiguration", reflect.TypeOf((*MockConfigDatabase)(nil).PutInsightConfiguration), ctx, configuration)
}

// GetInsightConfiguration mocks base method
func (m *MockConfigDatabase) GetInsightConfiguration(ctx context.Context) (*database.InsightConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInsightConfiguration", ctx)
	ret0, _ := ret[0].(*database.InsightConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightConfiguration indicates an expected call of GetInsightConfiguration
func (mr *MockConfigDatabaseMockRecorder) GetInsightConfiguration(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightConfiguration", reflect.TypeOf((*MockConfigDatabase)(nil).GetInsightConfiguration), ctx)
}

// ListAllOutgoingAccessRequests mocks base method
func (m *MockConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) ([]*database.AccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllOutgoingAccessRequests", ctx)
	ret0, _ := ret[0].([]*database.AccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllOutgoingAccessRequests indicates an expected call of ListAllOutgoingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListAllOutgoingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllOutgoingAccessRequests), ctx)
}

// ListOutgoingAccessRequests mocks base method
func (m *MockConfigDatabase) ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*database.AccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutgoingAccessRequests", ctx, organizationName, serviceName)
	ret0, _ := ret[0].([]*database.AccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOutgoingAccessRequests indicates an expected call of ListOutgoingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListOutgoingAccessRequests(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListOutgoingAccessRequests), ctx, organizationName, serviceName)
}

// GetLatestOutgoingAccessRequest mocks base method
func (m *MockConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*database.AccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestOutgoingAccessRequest", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.AccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestOutgoingAccessRequest indicates an expected call of GetLatestOutgoingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) GetLatestOutgoingAccessRequest(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestOutgoingAccessRequest), ctx, organizationName, serviceName)
}

// ListAllLatestOutgoingAccessRequests mocks base method
func (m *MockConfigDatabase) ListAllLatestOutgoingAccessRequests(ctx context.Context) (map[string]*database.AccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllLatestOutgoingAccessRequests", ctx)
	ret0, _ := ret[0].(map[string]*database.AccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllLatestOutgoingAccessRequests indicates an expected call of ListAllLatestOutgoingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListAllLatestOutgoingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllLatestOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllLatestOutgoingAccessRequests), ctx)
}

// CreateAccessRequest mocks base method
func (m *MockConfigDatabase) CreateAccessRequest(ctx context.Context, accessRequest *database.AccessRequest) (*database.AccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(*database.AccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessRequest indicates an expected call of CreateAccessRequest
func (mr *MockConfigDatabaseMockRecorder) CreateAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).CreateAccessRequest), ctx, accessRequest)
}

// UpdateAccessRequestState mocks base method
func (m *MockConfigDatabase) UpdateAccessRequestState(ctx context.Context, accessRequest *database.AccessRequest, state database.AccessRequestState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccessRequestState", ctx, accessRequest, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccessRequestState indicates an expected call of UpdateAccessRequestState
func (mr *MockConfigDatabaseMockRecorder) UpdateAccessRequestState(ctx, accessRequest, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessRequestState", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateAccessRequestState), ctx, accessRequest, state)
}

// GetSettings mocks base method
func (m *MockConfigDatabase) GetSettings(ctx context.Context) (*database.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", ctx)
	ret0, _ := ret[0].(*database.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings
func (mr *MockConfigDatabaseMockRecorder) GetSettings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockConfigDatabase)(nil).GetSettings), ctx)
}

// UpdateSettings mocks base method
func (m *MockConfigDatabase) UpdateSettings(ctx context.Context, settings *database.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSettings", ctx, settings)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettings indicates an expected call of UpdateSettings
func (mr *MockConfigDatabaseMockRecorder) UpdateSettings(ctx, settings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettings", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateSettings), ctx, settings)
}
