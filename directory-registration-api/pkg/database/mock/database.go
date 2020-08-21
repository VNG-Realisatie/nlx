// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/database/database.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	database "go.nlx.io/nlx/directory-registration-api/pkg/database"
	reflect "reflect"
)

// MockDirectoryDatabase is a mock of DirectoryDatabase interface
type MockDirectoryDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryDatabaseMockRecorder
}

// MockDirectoryDatabaseMockRecorder is the mock recorder for MockDirectoryDatabase
type MockDirectoryDatabaseMockRecorder struct {
	mock *MockDirectoryDatabase
}

// NewMockDirectoryDatabase creates a new mock instance
func NewMockDirectoryDatabase(ctrl *gomock.Controller) *MockDirectoryDatabase {
	mock := &MockDirectoryDatabase{ctrl: ctrl}
	mock.recorder = &MockDirectoryDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirectoryDatabase) EXPECT() *MockDirectoryDatabaseMockRecorder {
	return m.recorder
}

// SetInsightConfiguration mocks base method
func (m *MockDirectoryDatabase) SetInsightConfiguration(ctx context.Context, organizationName, insightAPIURL, irmaServerURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInsightConfiguration", ctx, organizationName, insightAPIURL, irmaServerURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInsightConfiguration indicates an expected call of SetInsightConfiguration
func (mr *MockDirectoryDatabaseMockRecorder) SetInsightConfiguration(ctx, organizationName, insightAPIURL, irmaServerURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInsightConfiguration", reflect.TypeOf((*MockDirectoryDatabase)(nil).SetInsightConfiguration), ctx, organizationName, insightAPIURL, irmaServerURL)
}

// InsertAvailability mocks base method
func (m *MockDirectoryDatabase) InsertAvailability(params *database.InsertAvailabilityParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAvailability", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAvailability indicates an expected call of InsertAvailability
func (mr *MockDirectoryDatabaseMockRecorder) InsertAvailability(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAvailability", reflect.TypeOf((*MockDirectoryDatabase)(nil).InsertAvailability), params)
}
