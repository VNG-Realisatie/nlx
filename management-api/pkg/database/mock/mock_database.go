// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	diagnostics "go.nlx.io/nlx/common/diagnostics"
	database "go.nlx.io/nlx/management-api/pkg/database"
)

// MockConfigDatabase is a mock of ConfigDatabase interface.
type MockConfigDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockConfigDatabaseMockRecorder
}

// MockConfigDatabaseMockRecorder is the mock recorder for MockConfigDatabase.
type MockConfigDatabaseMockRecorder struct {
	mock *MockConfigDatabase
}

// NewMockConfigDatabase creates a new mock instance.
func NewMockConfigDatabase(ctrl *gomock.Controller) *MockConfigDatabase {
	mock := &MockConfigDatabase{ctrl: ctrl}
	mock.recorder = &MockConfigDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigDatabase) EXPECT() *MockConfigDatabaseMockRecorder {
	return m.recorder
}

// CreateAccessGrant mocks base method.
func (m *MockConfigDatabase) CreateAccessGrant(ctx context.Context, accessRequest *database.IncomingAccessRequest) (*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessGrant", ctx, accessRequest)
	ret0, _ := ret[0].(*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessGrant indicates an expected call of CreateAccessGrant.
func (mr *MockConfigDatabaseMockRecorder) CreateAccessGrant(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessGrant", reflect.TypeOf((*MockConfigDatabase)(nil).CreateAccessGrant), ctx, accessRequest)
}

// CreateAccessProof mocks base method.
func (m *MockConfigDatabase) CreateAccessProof(ctx context.Context, accessRequest *database.OutgoingAccessRequest) (*database.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessProof", ctx, accessRequest)
	ret0, _ := ret[0].(*database.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessProof indicates an expected call of CreateAccessProof.
func (mr *MockConfigDatabaseMockRecorder) CreateAccessProof(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessProof", reflect.TypeOf((*MockConfigDatabase)(nil).CreateAccessProof), ctx, accessRequest)
}

// CreateAuditLogRecord mocks base method.
func (m *MockConfigDatabase) CreateAuditLogRecord(ctx context.Context, auditLogRecord *database.AuditLogRecord) (*database.AuditLogRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuditLogRecord", ctx, auditLogRecord)
	ret0, _ := ret[0].(*database.AuditLogRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuditLogRecord indicates an expected call of CreateAuditLogRecord.
func (mr *MockConfigDatabaseMockRecorder) CreateAuditLogRecord(ctx, auditLogRecord interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuditLogRecord", reflect.TypeOf((*MockConfigDatabase)(nil).CreateAuditLogRecord), ctx, auditLogRecord)
}

// CreateIncomingAccessRequest mocks base method.
func (m *MockConfigDatabase) CreateIncomingAccessRequest(ctx context.Context, accessRequest *database.IncomingAccessRequest) (*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncomingAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIncomingAccessRequest indicates an expected call of CreateIncomingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) CreateIncomingAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncomingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).CreateIncomingAccessRequest), ctx, accessRequest)
}

// CreateInway mocks base method.
func (m *MockConfigDatabase) CreateInway(ctx context.Context, inway *database.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInway", ctx, inway)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInway indicates an expected call of CreateInway.
func (mr *MockConfigDatabaseMockRecorder) CreateInway(ctx, inway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInway", reflect.TypeOf((*MockConfigDatabase)(nil).CreateInway), ctx, inway)
}

// CreateOrder mocks base method.
func (m *MockConfigDatabase) CreateOrder(ctx context.Context, order *database.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockConfigDatabaseMockRecorder) CreateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockConfigDatabase)(nil).CreateOrder), ctx, order)
}

// CreateOutgoingAccessRequest mocks base method.
func (m *MockConfigDatabase) CreateOutgoingAccessRequest(ctx context.Context, accessRequest *database.OutgoingAccessRequest) (*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOutgoingAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOutgoingAccessRequest indicates an expected call of CreateOutgoingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) CreateOutgoingAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).CreateOutgoingAccessRequest), ctx, accessRequest)
}

// CreateService mocks base method.
func (m *MockConfigDatabase) CreateService(ctx context.Context, service *database.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", ctx, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService.
func (mr *MockConfigDatabaseMockRecorder) CreateService(ctx, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockConfigDatabase)(nil).CreateService), ctx, service)
}

// CreateServiceWithInways mocks base method.
func (m *MockConfigDatabase) CreateServiceWithInways(ctx context.Context, service *database.Service, inwayNames []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceWithInways", ctx, service, inwayNames)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateServiceWithInways indicates an expected call of CreateServiceWithInways.
func (mr *MockConfigDatabaseMockRecorder) CreateServiceWithInways(ctx, service, inwayNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceWithInways", reflect.TypeOf((*MockConfigDatabase)(nil).CreateServiceWithInways), ctx, service, inwayNames)
}

// CreateUser mocks base method.
func (m *MockConfigDatabase) CreateUser(ctx context.Context, email string, roleNames []string) (*database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, email, roleNames)
	ret0, _ := ret[0].(*database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockConfigDatabaseMockRecorder) CreateUser(ctx, email, roleNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockConfigDatabase)(nil).CreateUser), ctx, email, roleNames)
}

// DeleteInway mocks base method.
func (m *MockConfigDatabase) DeleteInway(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInway", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInway indicates an expected call of DeleteInway.
func (mr *MockConfigDatabaseMockRecorder) DeleteInway(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInway", reflect.TypeOf((*MockConfigDatabase)(nil).DeleteInway), ctx, name)
}

// DeleteService mocks base method.
func (m *MockConfigDatabase) DeleteService(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockConfigDatabaseMockRecorder) DeleteService(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockConfigDatabase)(nil).DeleteService), ctx, name)
}

// GetAccessProofForOutgoingAccessRequest mocks base method.
func (m *MockConfigDatabase) GetAccessProofForOutgoingAccessRequest(ctx context.Context, accessRequestID uint) (*database.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessProofForOutgoingAccessRequest", ctx, accessRequestID)
	ret0, _ := ret[0].(*database.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessProofForOutgoingAccessRequest indicates an expected call of GetAccessProofForOutgoingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) GetAccessProofForOutgoingAccessRequest(ctx, accessRequestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessProofForOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetAccessProofForOutgoingAccessRequest), ctx, accessRequestID)
}

// GetIncomingAccessRequest mocks base method.
func (m *MockConfigDatabase) GetIncomingAccessRequest(ctx context.Context, id uint) (*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncomingAccessRequest", ctx, id)
	ret0, _ := ret[0].(*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncomingAccessRequest indicates an expected call of GetIncomingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) GetIncomingAccessRequest(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncomingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetIncomingAccessRequest), ctx, id)
}

// GetIncomingAccessRequestCountByService mocks base method.
func (m *MockConfigDatabase) GetIncomingAccessRequestCountByService(ctx context.Context) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncomingAccessRequestCountByService", ctx)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncomingAccessRequestCountByService indicates an expected call of GetIncomingAccessRequestCountByService.
func (mr *MockConfigDatabaseMockRecorder) GetIncomingAccessRequestCountByService(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncomingAccessRequestCountByService", reflect.TypeOf((*MockConfigDatabase)(nil).GetIncomingAccessRequestCountByService), ctx)
}

// GetInway mocks base method.
func (m *MockConfigDatabase) GetInway(ctx context.Context, name string) (*database.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInway", ctx, name)
	ret0, _ := ret[0].(*database.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInway indicates an expected call of GetInway.
func (mr *MockConfigDatabaseMockRecorder) GetInway(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInway", reflect.TypeOf((*MockConfigDatabase)(nil).GetInway), ctx, name)
}

// GetLatestAccessGrantForService mocks base method.
func (m *MockConfigDatabase) GetLatestAccessGrantForService(ctx context.Context, organizationName, serviceName string) (*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestAccessGrantForService", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestAccessGrantForService indicates an expected call of GetLatestAccessGrantForService.
func (mr *MockConfigDatabaseMockRecorder) GetLatestAccessGrantForService(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestAccessGrantForService", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestAccessGrantForService), ctx, organizationName, serviceName)
}

// GetLatestAccessProofForService mocks base method.
func (m *MockConfigDatabase) GetLatestAccessProofForService(ctx context.Context, organizationName, serviceName string) (*database.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestAccessProofForService", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestAccessProofForService indicates an expected call of GetLatestAccessProofForService.
func (mr *MockConfigDatabaseMockRecorder) GetLatestAccessProofForService(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestAccessProofForService", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestAccessProofForService), ctx, organizationName, serviceName)
}

// GetLatestIncomingAccessRequest mocks base method.
func (m *MockConfigDatabase) GetLatestIncomingAccessRequest(ctx context.Context, organizationName, serviceName string) (*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestIncomingAccessRequest", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestIncomingAccessRequest indicates an expected call of GetLatestIncomingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) GetLatestIncomingAccessRequest(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestIncomingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestIncomingAccessRequest), ctx, organizationName, serviceName)
}

// GetLatestOutgoingAccessRequest mocks base method.
func (m *MockConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestOutgoingAccessRequest", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestOutgoingAccessRequest indicates an expected call of GetLatestOutgoingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) GetLatestOutgoingAccessRequest(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestOutgoingAccessRequest), ctx, organizationName, serviceName)
}

// GetOrderByReference mocks base method.
func (m *MockConfigDatabase) GetOrderByReference(ctx context.Context, reference string) (*database.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByReference", ctx, reference)
	ret0, _ := ret[0].(*database.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByReference indicates an expected call of GetOrderByReference.
func (mr *MockConfigDatabaseMockRecorder) GetOrderByReference(ctx, reference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByReference", reflect.TypeOf((*MockConfigDatabase)(nil).GetOrderByReference), ctx, reference)
}

// GetOutgoingAccessRequest mocks base method.
func (m *MockConfigDatabase) GetOutgoingAccessRequest(ctx context.Context, id uint) (*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutgoingAccessRequest", ctx, id)
	ret0, _ := ret[0].(*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutgoingAccessRequest indicates an expected call of GetOutgoingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) GetOutgoingAccessRequest(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetOutgoingAccessRequest), ctx, id)
}

// GetService mocks base method.
func (m *MockConfigDatabase) GetService(ctx context.Context, name string) (*database.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", ctx, name)
	ret0, _ := ret[0].(*database.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockConfigDatabaseMockRecorder) GetService(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockConfigDatabase)(nil).GetService), ctx, name)
}

// GetSettings mocks base method.
func (m *MockConfigDatabase) GetSettings(ctx context.Context) (*database.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", ctx)
	ret0, _ := ret[0].(*database.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockConfigDatabaseMockRecorder) GetSettings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockConfigDatabase)(nil).GetSettings), ctx)
}

// GetUser mocks base method.
func (m *MockConfigDatabase) GetUser(ctx context.Context, email string) (*database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(*database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockConfigDatabaseMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockConfigDatabase)(nil).GetUser), ctx, email)
}

// ListAccessGrantsForService mocks base method.
func (m *MockConfigDatabase) ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessGrantsForService", ctx, serviceName)
	ret0, _ := ret[0].([]*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessGrantsForService indicates an expected call of ListAccessGrantsForService.
func (mr *MockConfigDatabaseMockRecorder) ListAccessGrantsForService(ctx, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessGrantsForService", reflect.TypeOf((*MockConfigDatabase)(nil).ListAccessGrantsForService), ctx, serviceName)
}

// ListAllIncomingAccessRequests mocks base method.
func (m *MockConfigDatabase) ListAllIncomingAccessRequests(ctx context.Context) ([]*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllIncomingAccessRequests", ctx)
	ret0, _ := ret[0].([]*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllIncomingAccessRequests indicates an expected call of ListAllIncomingAccessRequests.
func (mr *MockConfigDatabaseMockRecorder) ListAllIncomingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllIncomingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllIncomingAccessRequests), ctx)
}

// ListAllOutgoingAccessRequests mocks base method.
func (m *MockConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) ([]*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllOutgoingAccessRequests", ctx)
	ret0, _ := ret[0].([]*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllOutgoingAccessRequests indicates an expected call of ListAllOutgoingAccessRequests.
func (mr *MockConfigDatabaseMockRecorder) ListAllOutgoingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllOutgoingAccessRequests), ctx)
}

// ListAuditLogRecords mocks base method.
func (m *MockConfigDatabase) ListAuditLogRecords(ctx context.Context) ([]*database.AuditLogRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuditLogRecords", ctx)
	ret0, _ := ret[0].([]*database.AuditLogRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuditLogRecords indicates an expected call of ListAuditLogRecords.
func (mr *MockConfigDatabaseMockRecorder) ListAuditLogRecords(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuditLogRecords", reflect.TypeOf((*MockConfigDatabase)(nil).ListAuditLogRecords), ctx)
}

// ListIncomingAccessRequests mocks base method.
func (m *MockConfigDatabase) ListIncomingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIncomingAccessRequests", ctx, organizationName, serviceName)
	ret0, _ := ret[0].([]*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIncomingAccessRequests indicates an expected call of ListIncomingAccessRequests.
func (mr *MockConfigDatabaseMockRecorder) ListIncomingAccessRequests(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncomingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListIncomingAccessRequests), ctx, organizationName, serviceName)
}

// ListInways mocks base method.
func (m *MockConfigDatabase) ListInways(ctx context.Context) ([]*database.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInways", ctx)
	ret0, _ := ret[0].([]*database.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInways indicates an expected call of ListInways.
func (mr *MockConfigDatabaseMockRecorder) ListInways(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInways", reflect.TypeOf((*MockConfigDatabase)(nil).ListInways), ctx)
}

// ListOutgoingAccessRequests mocks base method.
func (m *MockConfigDatabase) ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutgoingAccessRequests", ctx, organizationName, serviceName)
	ret0, _ := ret[0].([]*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOutgoingAccessRequests indicates an expected call of ListOutgoingAccessRequests.
func (mr *MockConfigDatabaseMockRecorder) ListOutgoingAccessRequests(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListOutgoingAccessRequests), ctx, organizationName, serviceName)
}

// ListServices mocks base method.
func (m *MockConfigDatabase) ListServices(ctx context.Context) ([]*database.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", ctx)
	ret0, _ := ret[0].([]*database.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockConfigDatabaseMockRecorder) ListServices(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockConfigDatabase)(nil).ListServices), ctx)
}

// PutInsightConfiguration mocks base method.
func (m *MockConfigDatabase) PutInsightConfiguration(ctx context.Context, irmaServerURL, insightAPIURL string) (*database.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutInsightConfiguration", ctx, irmaServerURL, insightAPIURL)
	ret0, _ := ret[0].(*database.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutInsightConfiguration indicates an expected call of PutInsightConfiguration.
func (mr *MockConfigDatabaseMockRecorder) PutInsightConfiguration(ctx, irmaServerURL, insightAPIURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutInsightConfiguration", reflect.TypeOf((*MockConfigDatabase)(nil).PutInsightConfiguration), ctx, irmaServerURL, insightAPIURL)
}

// PutOrganizationInway mocks base method.
func (m *MockConfigDatabase) PutOrganizationInway(ctx context.Context, inwayID *uint) (*database.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutOrganizationInway", ctx, inwayID)
	ret0, _ := ret[0].(*database.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutOrganizationInway indicates an expected call of PutOrganizationInway.
func (mr *MockConfigDatabaseMockRecorder) PutOrganizationInway(ctx, inwayID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutOrganizationInway", reflect.TypeOf((*MockConfigDatabase)(nil).PutOrganizationInway), ctx, inwayID)
}

// RevokeAccessGrant mocks base method.
func (m *MockConfigDatabase) RevokeAccessGrant(ctx context.Context, id uint, revokedAt time.Time) (*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccessGrant", ctx, id, revokedAt)
	ret0, _ := ret[0].(*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAccessGrant indicates an expected call of RevokeAccessGrant.
func (mr *MockConfigDatabaseMockRecorder) RevokeAccessGrant(ctx, id, revokedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccessGrant", reflect.TypeOf((*MockConfigDatabase)(nil).RevokeAccessGrant), ctx, id, revokedAt)
}

// RevokeAccessProof mocks base method.
func (m *MockConfigDatabase) RevokeAccessProof(ctx context.Context, id uint, revokedAt time.Time) (*database.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccessProof", ctx, id, revokedAt)
	ret0, _ := ret[0].(*database.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAccessProof indicates an expected call of RevokeAccessProof.
func (mr *MockConfigDatabaseMockRecorder) RevokeAccessProof(ctx, id, revokedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccessProof", reflect.TypeOf((*MockConfigDatabase)(nil).RevokeAccessProof), ctx, id, revokedAt)
}

// TakePendingOutgoingAccessRequest mocks base method.
func (m *MockConfigDatabase) TakePendingOutgoingAccessRequest(ctx context.Context) (*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakePendingOutgoingAccessRequest", ctx)
	ret0, _ := ret[0].(*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakePendingOutgoingAccessRequest indicates an expected call of TakePendingOutgoingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) TakePendingOutgoingAccessRequest(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakePendingOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).TakePendingOutgoingAccessRequest), ctx)
}

// UnlockOutgoingAccessRequest mocks base method.
func (m *MockConfigDatabase) UnlockOutgoingAccessRequest(ctx context.Context, accessRequest *database.OutgoingAccessRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockOutgoingAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlockOutgoingAccessRequest indicates an expected call of UnlockOutgoingAccessRequest.
func (mr *MockConfigDatabaseMockRecorder) UnlockOutgoingAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).UnlockOutgoingAccessRequest), ctx, accessRequest)
}

// UpdateIncomingAccessRequestState mocks base method.
func (m *MockConfigDatabase) UpdateIncomingAccessRequestState(ctx context.Context, id uint, state database.IncomingAccessRequestState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIncomingAccessRequestState", ctx, id, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIncomingAccessRequestState indicates an expected call of UpdateIncomingAccessRequestState.
func (mr *MockConfigDatabaseMockRecorder) UpdateIncomingAccessRequestState(ctx, id, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIncomingAccessRequestState", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateIncomingAccessRequestState), ctx, id, state)
}

// UpdateInway mocks base method.
func (m *MockConfigDatabase) UpdateInway(ctx context.Context, inway *database.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInway", ctx, inway)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInway indicates an expected call of UpdateInway.
func (mr *MockConfigDatabaseMockRecorder) UpdateInway(ctx, inway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInway", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateInway), ctx, inway)
}

// UpdateOutgoingAccessRequestState mocks base method.
func (m *MockConfigDatabase) UpdateOutgoingAccessRequestState(ctx context.Context, id uint, state database.OutgoingAccessRequestState, referenceID uint, err *diagnostics.ErrorDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOutgoingAccessRequestState", ctx, id, state, referenceID, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOutgoingAccessRequestState indicates an expected call of UpdateOutgoingAccessRequestState.
func (mr *MockConfigDatabaseMockRecorder) UpdateOutgoingAccessRequestState(ctx, id, state, referenceID, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOutgoingAccessRequestState", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateOutgoingAccessRequestState), ctx, id, state, referenceID, err)
}

// UpdateService mocks base method.
func (m *MockConfigDatabase) UpdateService(ctx context.Context, service *database.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", ctx, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockConfigDatabaseMockRecorder) UpdateService(ctx, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateService), ctx, service)
}

// UpdateServiceWithInways mocks base method.
func (m *MockConfigDatabase) UpdateServiceWithInways(ctx context.Context, service *database.Service, inwayNames []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceWithInways", ctx, service, inwayNames)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateServiceWithInways indicates an expected call of UpdateServiceWithInways.
func (mr *MockConfigDatabaseMockRecorder) UpdateServiceWithInways(ctx, service, inwayNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceWithInways", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateServiceWithInways), ctx, service, inwayNames)
}
