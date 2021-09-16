// Code generated by MockGen. DO NOT EDIT.
// Source: domain/directory/repository.go

// Package directory_mock is a generated GoMock package.
package directory_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "go.nlx.io/nlx/directory-registration-api/domain"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// ClearIfSetAsOrganizationInway mocks base method.
func (m *MockRepository) ClearIfSetAsOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearIfSetAsOrganizationInway", ctx, organizationName, inwayAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearIfSetAsOrganizationInway indicates an expected call of ClearIfSetAsOrganizationInway.
func (mr *MockRepositoryMockRecorder) ClearIfSetAsOrganizationInway(ctx, organizationName, inwayAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearIfSetAsOrganizationInway", reflect.TypeOf((*MockRepository)(nil).ClearIfSetAsOrganizationInway), ctx, organizationName, inwayAddress)
}

// ClearOrganizationInway mocks base method.
func (m *MockRepository) ClearOrganizationInway(ctx context.Context, organizationName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearOrganizationInway", ctx, organizationName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearOrganizationInway indicates an expected call of ClearOrganizationInway.
func (mr *MockRepositoryMockRecorder) ClearOrganizationInway(ctx, organizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearOrganizationInway", reflect.TypeOf((*MockRepository)(nil).ClearOrganizationInway), ctx, organizationName)
}

// GetInway mocks base method.
func (m *MockRepository) GetInway(name, organization string) (*domain.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInway", name, organization)
	ret0, _ := ret[0].(*domain.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInway indicates an expected call of GetInway.
func (mr *MockRepositoryMockRecorder) GetInway(name, organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInway", reflect.TypeOf((*MockRepository)(nil).GetInway), name, organization)
}

// GetOrganizationInwayAddress mocks base method.
func (m *MockRepository) GetOrganizationInwayAddress(ctx context.Context, organizationName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizationInwayAddress", ctx, organizationName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInwayAddress indicates an expected call of GetOrganizationInwayAddress.
func (mr *MockRepositoryMockRecorder) GetOrganizationInwayAddress(ctx, organizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInwayAddress", reflect.TypeOf((*MockRepository)(nil).GetOrganizationInwayAddress), ctx, organizationName)
}

// GetService mocks base method.
func (m *MockRepository) GetService(id uint) (*domain.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", id)
	ret0, _ := ret[0].(*domain.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockRepositoryMockRecorder) GetService(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockRepository)(nil).GetService), id)
}

// RegisterInway mocks base method.
func (m *MockRepository) RegisterInway(arg0 *domain.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInway", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterInway indicates an expected call of RegisterInway.
func (mr *MockRepositoryMockRecorder) RegisterInway(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInway", reflect.TypeOf((*MockRepository)(nil).RegisterInway), arg0)
}

// RegisterService mocks base method.
func (m *MockRepository) RegisterService(arg0 *domain.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterService indicates an expected call of RegisterService.
func (mr *MockRepositoryMockRecorder) RegisterService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterService", reflect.TypeOf((*MockRepository)(nil).RegisterService), arg0)
}

// SetOrganizationInway mocks base method.
func (m *MockRepository) SetOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrganizationInway", ctx, organizationName, inwayAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOrganizationInway indicates an expected call of SetOrganizationInway.
func (mr *MockRepositoryMockRecorder) SetOrganizationInway(ctx, organizationName, inwayAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrganizationInway", reflect.TypeOf((*MockRepository)(nil).SetOrganizationInway), ctx, organizationName, inwayAddress)
}
