// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/database.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "go.nlx.io/nlx/directory-registration-api/pkg/database"
)

// MockDirectoryDatabase is a mock of DirectoryDatabase interface.
type MockDirectoryDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryDatabaseMockRecorder
}

// MockDirectoryDatabaseMockRecorder is the mock recorder for MockDirectoryDatabase.
type MockDirectoryDatabaseMockRecorder struct {
	mock *MockDirectoryDatabase
}

// NewMockDirectoryDatabase creates a new mock instance.
func NewMockDirectoryDatabase(ctrl *gomock.Controller) *MockDirectoryDatabase {
	mock := &MockDirectoryDatabase{ctrl: ctrl}
	mock.recorder = &MockDirectoryDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectoryDatabase) EXPECT() *MockDirectoryDatabaseMockRecorder {
	return m.recorder
}

// ClearOrganizationInway mocks base method.
func (m *MockDirectoryDatabase) ClearOrganizationInway(ctx context.Context, organizationName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearOrganizationInway", ctx, organizationName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearOrganizationInway indicates an expected call of ClearOrganizationInway.
func (mr *MockDirectoryDatabaseMockRecorder) ClearOrganizationInway(ctx, organizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearOrganizationInway", reflect.TypeOf((*MockDirectoryDatabase)(nil).ClearOrganizationInway), ctx, organizationName)
}

// RegisterInway mocks base method.
func (m *MockDirectoryDatabase) RegisterInway(params *database.RegisterInwayParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInway", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterInway indicates an expected call of RegisterInway.
func (mr *MockDirectoryDatabaseMockRecorder) RegisterInway(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInway", reflect.TypeOf((*MockDirectoryDatabase)(nil).RegisterInway), params)
}

// RegisterService mocks base method.
func (m *MockDirectoryDatabase) RegisterService(params *database.RegisterServiceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterService", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterService indicates an expected call of RegisterService.
func (mr *MockDirectoryDatabaseMockRecorder) RegisterService(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterService", reflect.TypeOf((*MockDirectoryDatabase)(nil).RegisterService), params)
}

// SetOrganizationInway mocks base method.
func (m *MockDirectoryDatabase) SetOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrganizationInway", ctx, organizationName, inwayAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOrganizationInway indicates an expected call of SetOrganizationInway.
func (mr *MockDirectoryDatabaseMockRecorder) SetOrganizationInway(ctx, organizationName, inwayAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrganizationInway", reflect.TypeOf((*MockDirectoryDatabase)(nil).SetOrganizationInway), ctx, organizationName, inwayAddress)
}
