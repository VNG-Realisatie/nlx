// Code generated by MockGen. DO NOT EDIT.
// Source: domain/txlog/storage/repository.go

// Package txlog_mock is a generated GoMock package.
package txlog_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	domain "go.nlx.io/nlx/txlog-api/domain"
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

// CreateRecord mocks base method.
func (m *MockRepository) CreateRecord(ctx context.Context, record *domain.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", ctx, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockRepositoryMockRecorder) CreateRecord(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockRepository)(nil).CreateRecord), ctx, record)
}

// ListRecords mocks base method.
func (m *MockRepository) ListRecords(ctx context.Context, limit int32) ([]*domain.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecords", ctx, limit)
	ret0, _ := ret[0].([]*domain.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecords indicates an expected call of ListRecords.
func (mr *MockRepositoryMockRecorder) ListRecords(ctx, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecords", reflect.TypeOf((*MockRepository)(nil).ListRecords), ctx, limit)
}

// Shutdown mocks base method.
func (m *MockRepository) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockRepositoryMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRepository)(nil).Shutdown))
}
