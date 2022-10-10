// Code generated by MockGen. DO NOT EDIT.
// Source: api/external/external_grpc.pb.go

// Package mock_external is a generated GoMock package.
package mock_external

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	external "go.nlx.io/nlx/management-api/api/external"
)

// MockAccessRequestServiceClient is a mock of AccessRequestServiceClient interface.
type MockAccessRequestServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccessRequestServiceClientMockRecorder
}

// MockAccessRequestServiceClientMockRecorder is the mock recorder for MockAccessRequestServiceClient.
type MockAccessRequestServiceClientMockRecorder struct {
	mock *MockAccessRequestServiceClient
}

// NewMockAccessRequestServiceClient creates a new mock instance.
func NewMockAccessRequestServiceClient(ctrl *gomock.Controller) *MockAccessRequestServiceClient {
	mock := &MockAccessRequestServiceClient{ctrl: ctrl}
	mock.recorder = &MockAccessRequestServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessRequestServiceClient) EXPECT() *MockAccessRequestServiceClientMockRecorder {
	return m.recorder
}

// GetAccessProof mocks base method.
func (m *MockAccessRequestServiceClient) GetAccessProof(ctx context.Context, in *external.GetAccessProofRequest, opts ...grpc.CallOption) (*external.AccessProof, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessProof", varargs...)
	ret0, _ := ret[0].(*external.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessProof indicates an expected call of GetAccessProof.
func (mr *MockAccessRequestServiceClientMockRecorder) GetAccessProof(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessProof", reflect.TypeOf((*MockAccessRequestServiceClient)(nil).GetAccessProof), varargs...)
}

// GetAccessRequestState mocks base method.
func (m *MockAccessRequestServiceClient) GetAccessRequestState(ctx context.Context, in *external.GetAccessRequestStateRequest, opts ...grpc.CallOption) (*external.GetAccessRequestStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessRequestState", varargs...)
	ret0, _ := ret[0].(*external.GetAccessRequestStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessRequestState indicates an expected call of GetAccessRequestState.
func (mr *MockAccessRequestServiceClientMockRecorder) GetAccessRequestState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessRequestState", reflect.TypeOf((*MockAccessRequestServiceClient)(nil).GetAccessRequestState), varargs...)
}

// RequestAccess mocks base method.
func (m *MockAccessRequestServiceClient) RequestAccess(ctx context.Context, in *external.RequestAccessRequest, opts ...grpc.CallOption) (*external.RequestAccessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestAccess", varargs...)
	ret0, _ := ret[0].(*external.RequestAccessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestAccess indicates an expected call of RequestAccess.
func (mr *MockAccessRequestServiceClientMockRecorder) RequestAccess(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAccess", reflect.TypeOf((*MockAccessRequestServiceClient)(nil).RequestAccess), varargs...)
}

// MockAccessRequestServiceServer is a mock of AccessRequestServiceServer interface.
type MockAccessRequestServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccessRequestServiceServerMockRecorder
}

// MockAccessRequestServiceServerMockRecorder is the mock recorder for MockAccessRequestServiceServer.
type MockAccessRequestServiceServerMockRecorder struct {
	mock *MockAccessRequestServiceServer
}

// NewMockAccessRequestServiceServer creates a new mock instance.
func NewMockAccessRequestServiceServer(ctrl *gomock.Controller) *MockAccessRequestServiceServer {
	mock := &MockAccessRequestServiceServer{ctrl: ctrl}
	mock.recorder = &MockAccessRequestServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessRequestServiceServer) EXPECT() *MockAccessRequestServiceServerMockRecorder {
	return m.recorder
}

// GetAccessProof mocks base method.
func (m *MockAccessRequestServiceServer) GetAccessProof(arg0 context.Context, arg1 *external.GetAccessProofRequest) (*external.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessProof", arg0, arg1)
	ret0, _ := ret[0].(*external.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessProof indicates an expected call of GetAccessProof.
func (mr *MockAccessRequestServiceServerMockRecorder) GetAccessProof(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessProof", reflect.TypeOf((*MockAccessRequestServiceServer)(nil).GetAccessProof), arg0, arg1)
}

// GetAccessRequestState mocks base method.
func (m *MockAccessRequestServiceServer) GetAccessRequestState(arg0 context.Context, arg1 *external.GetAccessRequestStateRequest) (*external.GetAccessRequestStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessRequestState", arg0, arg1)
	ret0, _ := ret[0].(*external.GetAccessRequestStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessRequestState indicates an expected call of GetAccessRequestState.
func (mr *MockAccessRequestServiceServerMockRecorder) GetAccessRequestState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessRequestState", reflect.TypeOf((*MockAccessRequestServiceServer)(nil).GetAccessRequestState), arg0, arg1)
}

// RequestAccess mocks base method.
func (m *MockAccessRequestServiceServer) RequestAccess(arg0 context.Context, arg1 *external.RequestAccessRequest) (*external.RequestAccessResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestAccess", arg0, arg1)
	ret0, _ := ret[0].(*external.RequestAccessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestAccess indicates an expected call of RequestAccess.
func (mr *MockAccessRequestServiceServerMockRecorder) RequestAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAccess", reflect.TypeOf((*MockAccessRequestServiceServer)(nil).RequestAccess), arg0, arg1)
}

// mustEmbedUnimplementedAccessRequestServiceServer mocks base method.
func (m *MockAccessRequestServiceServer) mustEmbedUnimplementedAccessRequestServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAccessRequestServiceServer")
}

// mustEmbedUnimplementedAccessRequestServiceServer indicates an expected call of mustEmbedUnimplementedAccessRequestServiceServer.
func (mr *MockAccessRequestServiceServerMockRecorder) mustEmbedUnimplementedAccessRequestServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAccessRequestServiceServer", reflect.TypeOf((*MockAccessRequestServiceServer)(nil).mustEmbedUnimplementedAccessRequestServiceServer))
}

// MockUnsafeAccessRequestServiceServer is a mock of UnsafeAccessRequestServiceServer interface.
type MockUnsafeAccessRequestServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAccessRequestServiceServerMockRecorder
}

// MockUnsafeAccessRequestServiceServerMockRecorder is the mock recorder for MockUnsafeAccessRequestServiceServer.
type MockUnsafeAccessRequestServiceServerMockRecorder struct {
	mock *MockUnsafeAccessRequestServiceServer
}

// NewMockUnsafeAccessRequestServiceServer creates a new mock instance.
func NewMockUnsafeAccessRequestServiceServer(ctrl *gomock.Controller) *MockUnsafeAccessRequestServiceServer {
	mock := &MockUnsafeAccessRequestServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAccessRequestServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAccessRequestServiceServer) EXPECT() *MockUnsafeAccessRequestServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAccessRequestServiceServer mocks base method.
func (m *MockUnsafeAccessRequestServiceServer) mustEmbedUnimplementedAccessRequestServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAccessRequestServiceServer")
}

// mustEmbedUnimplementedAccessRequestServiceServer indicates an expected call of mustEmbedUnimplementedAccessRequestServiceServer.
func (mr *MockUnsafeAccessRequestServiceServerMockRecorder) mustEmbedUnimplementedAccessRequestServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAccessRequestServiceServer", reflect.TypeOf((*MockUnsafeAccessRequestServiceServer)(nil).mustEmbedUnimplementedAccessRequestServiceServer))
}

// MockDelegationServiceClient is a mock of DelegationServiceClient interface.
type MockDelegationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDelegationServiceClientMockRecorder
}

// MockDelegationServiceClientMockRecorder is the mock recorder for MockDelegationServiceClient.
type MockDelegationServiceClientMockRecorder struct {
	mock *MockDelegationServiceClient
}

// NewMockDelegationServiceClient creates a new mock instance.
func NewMockDelegationServiceClient(ctrl *gomock.Controller) *MockDelegationServiceClient {
	mock := &MockDelegationServiceClient{ctrl: ctrl}
	mock.recorder = &MockDelegationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelegationServiceClient) EXPECT() *MockDelegationServiceClientMockRecorder {
	return m.recorder
}

// ListOrders mocks base method.
func (m *MockDelegationServiceClient) ListOrders(ctx context.Context, in *external.ListOrdersRequest, opts ...grpc.CallOption) (*external.ListOrdersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrders", varargs...)
	ret0, _ := ret[0].(*external.ListOrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockDelegationServiceClientMockRecorder) ListOrders(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockDelegationServiceClient)(nil).ListOrders), varargs...)
}

// RequestClaim mocks base method.
func (m *MockDelegationServiceClient) RequestClaim(ctx context.Context, in *external.RequestClaimRequest, opts ...grpc.CallOption) (*external.RequestClaimResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestClaim", varargs...)
	ret0, _ := ret[0].(*external.RequestClaimResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestClaim indicates an expected call of RequestClaim.
func (mr *MockDelegationServiceClientMockRecorder) RequestClaim(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestClaim", reflect.TypeOf((*MockDelegationServiceClient)(nil).RequestClaim), varargs...)
}

// MockDelegationServiceServer is a mock of DelegationServiceServer interface.
type MockDelegationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockDelegationServiceServerMockRecorder
}

// MockDelegationServiceServerMockRecorder is the mock recorder for MockDelegationServiceServer.
type MockDelegationServiceServerMockRecorder struct {
	mock *MockDelegationServiceServer
}

// NewMockDelegationServiceServer creates a new mock instance.
func NewMockDelegationServiceServer(ctrl *gomock.Controller) *MockDelegationServiceServer {
	mock := &MockDelegationServiceServer{ctrl: ctrl}
	mock.recorder = &MockDelegationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelegationServiceServer) EXPECT() *MockDelegationServiceServerMockRecorder {
	return m.recorder
}

// ListOrders mocks base method.
func (m *MockDelegationServiceServer) ListOrders(arg0 context.Context, arg1 *external.ListOrdersRequest) (*external.ListOrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", arg0, arg1)
	ret0, _ := ret[0].(*external.ListOrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockDelegationServiceServerMockRecorder) ListOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockDelegationServiceServer)(nil).ListOrders), arg0, arg1)
}

// RequestClaim mocks base method.
func (m *MockDelegationServiceServer) RequestClaim(arg0 context.Context, arg1 *external.RequestClaimRequest) (*external.RequestClaimResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestClaim", arg0, arg1)
	ret0, _ := ret[0].(*external.RequestClaimResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestClaim indicates an expected call of RequestClaim.
func (mr *MockDelegationServiceServerMockRecorder) RequestClaim(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestClaim", reflect.TypeOf((*MockDelegationServiceServer)(nil).RequestClaim), arg0, arg1)
}

// mustEmbedUnimplementedDelegationServiceServer mocks base method.
func (m *MockDelegationServiceServer) mustEmbedUnimplementedDelegationServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDelegationServiceServer")
}

// mustEmbedUnimplementedDelegationServiceServer indicates an expected call of mustEmbedUnimplementedDelegationServiceServer.
func (mr *MockDelegationServiceServerMockRecorder) mustEmbedUnimplementedDelegationServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDelegationServiceServer", reflect.TypeOf((*MockDelegationServiceServer)(nil).mustEmbedUnimplementedDelegationServiceServer))
}

// MockUnsafeDelegationServiceServer is a mock of UnsafeDelegationServiceServer interface.
type MockUnsafeDelegationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeDelegationServiceServerMockRecorder
}

// MockUnsafeDelegationServiceServerMockRecorder is the mock recorder for MockUnsafeDelegationServiceServer.
type MockUnsafeDelegationServiceServerMockRecorder struct {
	mock *MockUnsafeDelegationServiceServer
}

// NewMockUnsafeDelegationServiceServer creates a new mock instance.
func NewMockUnsafeDelegationServiceServer(ctrl *gomock.Controller) *MockUnsafeDelegationServiceServer {
	mock := &MockUnsafeDelegationServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeDelegationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeDelegationServiceServer) EXPECT() *MockUnsafeDelegationServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedDelegationServiceServer mocks base method.
func (m *MockUnsafeDelegationServiceServer) mustEmbedUnimplementedDelegationServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDelegationServiceServer")
}

// mustEmbedUnimplementedDelegationServiceServer indicates an expected call of mustEmbedUnimplementedDelegationServiceServer.
func (mr *MockUnsafeDelegationServiceServerMockRecorder) mustEmbedUnimplementedDelegationServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDelegationServiceServer", reflect.TypeOf((*MockUnsafeDelegationServiceServer)(nil).mustEmbedUnimplementedDelegationServiceServer))
}
