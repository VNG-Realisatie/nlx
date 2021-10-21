// Code generated by MockGen. DO NOT EDIT.
// Source: /home/ronald/go/src/github.com/cloudflare/cfssl/signer/signer.go

// Package mock_signer is a generated GoMock package.
package mock_signer

import (
	x509 "crypto/x509"
	http "net/http"
	reflect "reflect"

	certdb "github.com/cloudflare/cfssl/certdb"
	config "github.com/cloudflare/cfssl/config"
	info "github.com/cloudflare/cfssl/info"
	signer "github.com/cloudflare/cfssl/signer"
	gomock "github.com/golang/mock/gomock"
)

// MockSigner is a mock of Signer interface
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockSigner) Info(arg0 info.Req) (*info.Resp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(*info.Resp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockSignerMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockSigner)(nil).Info), arg0)
}

// Policy mocks base method
func (m *MockSigner) Policy() *config.Signing {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policy")
	ret0, _ := ret[0].(*config.Signing)
	return ret0
}

// Policy indicates an expected call of Policy
func (mr *MockSignerMockRecorder) Policy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policy", reflect.TypeOf((*MockSigner)(nil).Policy))
}

// SetDBAccessor mocks base method
func (m *MockSigner) SetDBAccessor(arg0 certdb.Accessor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDBAccessor", arg0)
}

// SetDBAccessor indicates an expected call of SetDBAccessor
func (mr *MockSignerMockRecorder) SetDBAccessor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDBAccessor", reflect.TypeOf((*MockSigner)(nil).SetDBAccessor), arg0)
}

// GetDBAccessor mocks base method
func (m *MockSigner) GetDBAccessor() certdb.Accessor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBAccessor")
	ret0, _ := ret[0].(certdb.Accessor)
	return ret0
}

// GetDBAccessor indicates an expected call of GetDBAccessor
func (mr *MockSignerMockRecorder) GetDBAccessor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBAccessor", reflect.TypeOf((*MockSigner)(nil).GetDBAccessor))
}

// SetPolicy mocks base method
func (m *MockSigner) SetPolicy(arg0 *config.Signing) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPolicy", arg0)
}

// SetPolicy indicates an expected call of SetPolicy
func (mr *MockSignerMockRecorder) SetPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPolicy", reflect.TypeOf((*MockSigner)(nil).SetPolicy), arg0)
}

// SigAlgo mocks base method
func (m *MockSigner) SigAlgo() x509.SignatureAlgorithm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SigAlgo")
	ret0, _ := ret[0].(x509.SignatureAlgorithm)
	return ret0
}

// SigAlgo indicates an expected call of SigAlgo
func (mr *MockSignerMockRecorder) SigAlgo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigAlgo", reflect.TypeOf((*MockSigner)(nil).SigAlgo))
}

// Sign mocks base method
func (m *MockSigner) Sign(req signer.SignRequest) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", req)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockSignerMockRecorder) Sign(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockSigner)(nil).Sign), req)
}

// SetReqModifier mocks base method
func (m *MockSigner) SetReqModifier(arg0 func(*http.Request, []byte)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReqModifier", arg0)
}

// SetReqModifier indicates an expected call of SetReqModifier
func (mr *MockSignerMockRecorder) SetReqModifier(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReqModifier", reflect.TypeOf((*MockSigner)(nil).SetReqModifier), arg0)
}
