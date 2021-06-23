// Code generated by MockGen. DO NOT EDIT.
// Source: domain/inway/repository.go

// Package inway_mock is a generated GoMock package.
package inway_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	inway "go.nlx.io/nlx/directory-registration-api/domain/inway"
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

// Register mocks base method.
func (m *MockRepository) Register(i *inway.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockRepositoryMockRecorder) Register(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRepository)(nil).Register), i)
}
