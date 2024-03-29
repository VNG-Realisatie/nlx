// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_outway is a generated GoMock package.
package mock_outway

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHTTPService is a mock of HTTPService interface.
type MockHTTPService struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPServiceMockRecorder
}

// MockHTTPServiceMockRecorder is the mock recorder for MockHTTPService.
type MockHTTPServiceMockRecorder struct {
	mock *MockHTTPService
}

// NewMockHTTPService creates a new mock instance.
func NewMockHTTPService(ctrl *gomock.Controller) *MockHTTPService {
	mock := &MockHTTPService{ctrl: ctrl}
	mock.recorder = &MockHTTPServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPService) EXPECT() *MockHTTPServiceMockRecorder {
	return m.recorder
}

// FullName mocks base method.
func (m *MockHTTPService) FullName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FullName indicates an expected call of FullName.
func (mr *MockHTTPServiceMockRecorder) FullName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullName", reflect.TypeOf((*MockHTTPService)(nil).FullName))
}

// GetInwayAddresses mocks base method.
func (m *MockHTTPService) GetInwayAddresses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInwayAddresses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetInwayAddresses indicates an expected call of GetInwayAddresses.
func (mr *MockHTTPServiceMockRecorder) GetInwayAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInwayAddresses", reflect.TypeOf((*MockHTTPService)(nil).GetInwayAddresses))
}

// ProxyHTTPRequest mocks base method.
func (m *MockHTTPService) ProxyHTTPRequest(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProxyHTTPRequest", w, r)
}

// ProxyHTTPRequest indicates an expected call of ProxyHTTPRequest.
func (mr *MockHTTPServiceMockRecorder) ProxyHTTPRequest(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyHTTPRequest", reflect.TypeOf((*MockHTTPService)(nil).ProxyHTTPRequest), w, r)
}
