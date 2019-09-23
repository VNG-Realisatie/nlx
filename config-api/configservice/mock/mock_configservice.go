// Code generated by MockGen. DO NOT EDIT.
// Source: config-database.go

// Package mock_configservice is a generated GoMock package.
package mock_configservice

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	configapi "go.nlx.io/nlx/config-api/configapi"
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
func (m *MockConfigDatabase) ListServices(ctx context.Context) ([]*configapi.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", ctx)
	ret0, _ := ret[0].([]*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockConfigDatabaseMockRecorder) ListServices(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockConfigDatabase)(nil).ListServices), ctx)
}

// GetService mocks base method
func (m *MockConfigDatabase) GetService(ctx context.Context, name string) (*configapi.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", ctx, name)
	ret0, _ := ret[0].(*configapi.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockConfigDatabaseMockRecorder) GetService(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockConfigDatabase)(nil).GetService), ctx, name)
}

// CreateService mocks base method
func (m *MockConfigDatabase) CreateService(ctx context.Context, service *configapi.Service) error {
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
func (m *MockConfigDatabase) UpdateService(ctx context.Context, name string, service *configapi.Service) error {
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
func (m *MockConfigDatabase) ListInways(ctx context.Context) ([]*configapi.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInways", ctx)
	ret0, _ := ret[0].([]*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInways indicates an expected call of ListInways
func (mr *MockConfigDatabaseMockRecorder) ListInways(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInways", reflect.TypeOf((*MockConfigDatabase)(nil).ListInways), ctx)
}

// GetInway mocks base method
func (m *MockConfigDatabase) GetInway(ctx context.Context, name string) (*configapi.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInway", ctx, name)
	ret0, _ := ret[0].(*configapi.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInway indicates an expected call of GetInway
func (mr *MockConfigDatabaseMockRecorder) GetInway(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInway", reflect.TypeOf((*MockConfigDatabase)(nil).GetInway), ctx, name)
}

// CreateInway mocks base method
func (m *MockConfigDatabase) CreateInway(ctx context.Context, inway *configapi.Inway) error {
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
func (m *MockConfigDatabase) UpdateInway(ctx context.Context, name string, inway *configapi.Inway) error {
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
func (m *MockConfigDatabase) PutInsightConfiguration(ctx context.Context, configuration *configapi.InsightConfiguration) error {
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
func (m *MockConfigDatabase) GetInsightConfiguration(ctx context.Context) (*configapi.InsightConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInsightConfiguration", ctx)
	ret0, _ := ret[0].(*configapi.InsightConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightConfiguration indicates an expected call of GetInsightConfiguration
func (mr *MockConfigDatabaseMockRecorder) GetInsightConfiguration(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightConfiguration", reflect.TypeOf((*MockConfigDatabase)(nil).GetInsightConfiguration), ctx)
}
