// Code generated by MockGen. DO NOT EDIT.
// Source: go.nlx.io/nlx/management-api/pkg/management (interfaces: Client)

// Package mock_management is a generated GoMock package.
package mock_management

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "go.nlx.io/nlx/management-api/api"
	external "go.nlx.io/nlx/management-api/api/external"
	grpc "google.golang.org/grpc"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// GetAccessProof mocks base method.
func (m *MockClient) GetAccessProof(arg0 context.Context, arg1 *external.GetAccessProofRequest, arg2 ...grpc.CallOption) (*api.AccessProof, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessProof", varargs...)
	ret0, _ := ret[0].(*api.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessProof indicates an expected call of GetAccessProof.
func (mr *MockClientMockRecorder) GetAccessProof(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessProof", reflect.TypeOf((*MockClient)(nil).GetAccessProof), varargs...)
}

// GetAccessRequestState mocks base method.
func (m *MockClient) GetAccessRequestState(arg0 context.Context, arg1 *external.GetAccessRequestStateRequest, arg2 ...grpc.CallOption) (*external.GetAccessRequestStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessRequestState", varargs...)
	ret0, _ := ret[0].(*external.GetAccessRequestStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessRequestState indicates an expected call of GetAccessRequestState.
func (mr *MockClientMockRecorder) GetAccessRequestState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessRequestState", reflect.TypeOf((*MockClient)(nil).GetAccessRequestState), varargs...)
}

// RequestAccess mocks base method.
func (m *MockClient) RequestAccess(arg0 context.Context, arg1 *external.RequestAccessRequest, arg2 ...grpc.CallOption) (*external.RequestAccessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestAccess", varargs...)
	ret0, _ := ret[0].(*external.RequestAccessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestAccess indicates an expected call of RequestAccess.
func (mr *MockClientMockRecorder) RequestAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAccess", reflect.TypeOf((*MockClient)(nil).RequestAccess), varargs...)
}

// RequestClaim mocks base method.
func (m *MockClient) RequestClaim(arg0 context.Context, arg1 *external.RequestClaimRequest, arg2 ...grpc.CallOption) (*external.RequestClaimResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestClaim", varargs...)
	ret0, _ := ret[0].(*external.RequestClaimResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestClaim indicates an expected call of RequestClaim.
func (mr *MockClientMockRecorder) RequestClaim(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestClaim", reflect.TypeOf((*MockClient)(nil).RequestClaim), varargs...)
}
