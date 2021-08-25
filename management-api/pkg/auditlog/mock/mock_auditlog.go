// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/auditlog/logger.go

// Package mock_auditlog is a generated GoMock package.
package mock_auditlog

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auditlog "go.nlx.io/nlx/management-api/pkg/auditlog"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// AccessGrantRevoke mocks base method.
func (m *MockLogger) AccessGrantRevoke(ctx context.Context, userName, userAgent, organization, serviceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessGrantRevoke", ctx, userName, userAgent, organization, serviceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// AccessGrantRevoke indicates an expected call of AccessGrantRevoke.
func (mr *MockLoggerMockRecorder) AccessGrantRevoke(ctx, userName, userAgent, organization, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessGrantRevoke", reflect.TypeOf((*MockLogger)(nil).AccessGrantRevoke), ctx, userName, userAgent, organization, serviceName)
}

// IncomingAccessRequestAccept mocks base method.
func (m *MockLogger) IncomingAccessRequestAccept(ctx context.Context, userName, userAgent, organization, service string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncomingAccessRequestAccept", ctx, userName, userAgent, organization, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncomingAccessRequestAccept indicates an expected call of IncomingAccessRequestAccept.
func (mr *MockLoggerMockRecorder) IncomingAccessRequestAccept(ctx, userName, userAgent, organization, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncomingAccessRequestAccept", reflect.TypeOf((*MockLogger)(nil).IncomingAccessRequestAccept), ctx, userName, userAgent, organization, service)
}

// IncomingAccessRequestReject mocks base method.
func (m *MockLogger) IncomingAccessRequestReject(ctx context.Context, userName, userAgent, organization, service string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncomingAccessRequestReject", ctx, userName, userAgent, organization, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncomingAccessRequestReject indicates an expected call of IncomingAccessRequestReject.
func (mr *MockLoggerMockRecorder) IncomingAccessRequestReject(ctx, userName, userAgent, organization, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncomingAccessRequestReject", reflect.TypeOf((*MockLogger)(nil).IncomingAccessRequestReject), ctx, userName, userAgent, organization, service)
}

// ListAll mocks base method.
func (m *MockLogger) ListAll(ctx context.Context) ([]*auditlog.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", ctx)
	ret0, _ := ret[0].([]*auditlog.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockLoggerMockRecorder) ListAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockLogger)(nil).ListAll), ctx)
}

// LoginFail mocks base method.
func (m *MockLogger) LoginFail(ctx context.Context, userAgent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginFail", ctx, userAgent)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoginFail indicates an expected call of LoginFail.
func (mr *MockLoggerMockRecorder) LoginFail(ctx, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginFail", reflect.TypeOf((*MockLogger)(nil).LoginFail), ctx, userAgent)
}

// LoginSuccess mocks base method.
func (m *MockLogger) LoginSuccess(ctx context.Context, userName, userAgent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginSuccess", ctx, userName, userAgent)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoginSuccess indicates an expected call of LoginSuccess.
func (mr *MockLoggerMockRecorder) LoginSuccess(ctx, userName, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginSuccess", reflect.TypeOf((*MockLogger)(nil).LoginSuccess), ctx, userName, userAgent)
}

// LogoutSuccess mocks base method.
func (m *MockLogger) LogoutSuccess(ctx context.Context, userName, userAgent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogoutSuccess", ctx, userName, userAgent)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogoutSuccess indicates an expected call of LogoutSuccess.
func (mr *MockLoggerMockRecorder) LogoutSuccess(ctx, userName, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutSuccess", reflect.TypeOf((*MockLogger)(nil).LogoutSuccess), ctx, userName, userAgent)
}

// OrderCreate mocks base method.
func (m *MockLogger) OrderCreate(ctx context.Context, userName, userAgent, delegatee string, services []auditlog.RecordService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderCreate", ctx, userName, userAgent, delegatee, services)
	ret0, _ := ret[0].(error)
	return ret0
}

// OrderCreate indicates an expected call of OrderCreate.
func (mr *MockLoggerMockRecorder) OrderCreate(ctx, userName, userAgent, delegatee, services interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderCreate", reflect.TypeOf((*MockLogger)(nil).OrderCreate), ctx, userName, userAgent, delegatee, services)
}

// OrderOutgoingRevoke mocks base method.
func (m *MockLogger) OrderOutgoingRevoke(ctx context.Context, userName, userAgent, delegatee, reference string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderOutgoingRevoke", ctx, userName, userAgent, delegatee, reference)
	ret0, _ := ret[0].(error)
	return ret0
}

// OrderOutgoingRevoke indicates an expected call of OrderOutgoingRevoke.
func (mr *MockLoggerMockRecorder) OrderOutgoingRevoke(ctx, userName, userAgent, delegatee, reference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderOutgoingRevoke", reflect.TypeOf((*MockLogger)(nil).OrderOutgoingRevoke), ctx, userName, userAgent, delegatee, reference)
}

// OrganizationSettingsUpdate mocks base method.
func (m *MockLogger) OrganizationSettingsUpdate(ctx context.Context, userName, userAgent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationSettingsUpdate", ctx, userName, userAgent)
	ret0, _ := ret[0].(error)
	return ret0
}

// OrganizationSettingsUpdate indicates an expected call of OrganizationSettingsUpdate.
func (mr *MockLoggerMockRecorder) OrganizationSettingsUpdate(ctx, userName, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationSettingsUpdate", reflect.TypeOf((*MockLogger)(nil).OrganizationSettingsUpdate), ctx, userName, userAgent)
}

// OutgoingAccessRequestCreate mocks base method.
func (m *MockLogger) OutgoingAccessRequestCreate(ctx context.Context, userName, userAgent, organization, service string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutgoingAccessRequestCreate", ctx, userName, userAgent, organization, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// OutgoingAccessRequestCreate indicates an expected call of OutgoingAccessRequestCreate.
func (mr *MockLoggerMockRecorder) OutgoingAccessRequestCreate(ctx, userName, userAgent, organization, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutgoingAccessRequestCreate", reflect.TypeOf((*MockLogger)(nil).OutgoingAccessRequestCreate), ctx, userName, userAgent, organization, service)
}

// ServiceCreate mocks base method.
func (m *MockLogger) ServiceCreate(ctx context.Context, userName, userAgent, serviceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceCreate", ctx, userName, userAgent, serviceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServiceCreate indicates an expected call of ServiceCreate.
func (mr *MockLoggerMockRecorder) ServiceCreate(ctx, userName, userAgent, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceCreate", reflect.TypeOf((*MockLogger)(nil).ServiceCreate), ctx, userName, userAgent, serviceName)
}

// ServiceDelete mocks base method.
func (m *MockLogger) ServiceDelete(ctx context.Context, userName, userAgent, serviceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceDelete", ctx, userName, userAgent, serviceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServiceDelete indicates an expected call of ServiceDelete.
func (mr *MockLoggerMockRecorder) ServiceDelete(ctx, userName, userAgent, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceDelete", reflect.TypeOf((*MockLogger)(nil).ServiceDelete), ctx, userName, userAgent, serviceName)
}

// ServiceUpdate mocks base method.
func (m *MockLogger) ServiceUpdate(ctx context.Context, userName, userAgent, serviceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceUpdate", ctx, userName, userAgent, serviceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServiceUpdate indicates an expected call of ServiceUpdate.
func (mr *MockLoggerMockRecorder) ServiceUpdate(ctx, userName, userAgent, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceUpdate", reflect.TypeOf((*MockLogger)(nil).ServiceUpdate), ctx, userName, userAgent, serviceName)
}
