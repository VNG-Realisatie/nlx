// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
)

//nolint:funlen,dupl // its a unittest
func TestListAuditLogs(t *testing.T) {
	createTimestamp := func(ti time.Time) *timestamppb.Timestamp {
		return &timestamppb.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	const numberOfAuditLogs = 100

	tests := map[string]struct {
		ctx     context.Context
		setup   func(serviceMocks)
		req     *api.ListAuditLogsRequest
		want    *api.ListAuditLogsResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(serviceMocks) {},
			want:    nil,
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.audit_logs.read\" to execute this request").Err(),
		},
		"when_error_occurs_while_retrieving_organizations": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(nil, errors.New("arbitrary error"))
			},
			req:     &api.ListAuditLogsRequest{},
			want:    nil,
			wantErr: status.New(codes.Internal, "failed to retrieve audit logs").Err(),
		},
		"when_error_occurs_while_retrieving_logs": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.al.
					EXPECT().
					List(gomock.Any(), numberOfAuditLogs).
					Return(nil, errors.New("arbitrary error"))
			},
			req:     &api.ListAuditLogsRequest{},
			want:    nil,
			wantErr: status.New(codes.Internal, "failed to retrieve audit logs").Err(),
		},
		"when_no_logs_are_available": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.al.
					EXPECT().
					List(gomock.Any(), numberOfAuditLogs).
					Return([]*auditlog.Record{}, nil)
			},
			req: &api.ListAuditLogsRequest{},
			want: &api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{},
			},
			wantErr: nil,
		},
		"with_a_single_log_created_via_the_browser": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.al.
					EXPECT().
					List(gomock.Any(), numberOfAuditLogs).
					Return([]*auditlog.Record{
						{
							ID:           1,
							Username:     "admin@example.com",
							ActionType:   auditlog.LoginSuccess,
							UserAgent:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
							HasSucceeded: true,
							CreatedAt:    time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
						},
					}, nil)
			},
			req: &api.ListAuditLogsRequest{},
			want: &api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "admin@example.com",
						Action:          api.AuditLogRecord_ACTION_TYPE_LOGIN_SUCCESS,
						OperatingSystem: "Mac OS X",
						Browser:         "Safari",
						Client:          "NLX Management",
						Services:        []*api.AuditLogRecord_Service{},
						HasSucceeded:    true,
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					},
				},
			},
			wantErr: nil,
		},
		"with_a_single_log_created_via_nlxctl": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.al.
					EXPECT().
					List(gomock.Any(), numberOfAuditLogs).
					Return([]*auditlog.Record{
						{
							ID:           1,
							Username:     "",
							ActionType:   auditlog.LoginSuccess,
							UserAgent:    "nlxctl/1.x (Mac OS X)",
							HasSucceeded: true,
							CreatedAt:    time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
						},
					}, nil)
			},
			req: &api.ListAuditLogsRequest{},
			want: &api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "",
						Action:          api.AuditLogRecord_ACTION_TYPE_LOGIN_SUCCESS,
						OperatingSystem: "Mac OS X",
						Browser:         "",
						Client:          "nlxctl",
						Services:        []*api.AuditLogRecord_Service{},
						HasSucceeded:    true,
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					},
				},
			},
			wantErr: nil,
		},
		"with_a_single_log_created_via_the_browser_with_order_metadata": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
						},
					}, nil)

				mocks.al.
					EXPECT().
					List(gomock.Any(), numberOfAuditLogs).
					Return([]*auditlog.Record{
						{
							ID:           1,
							Username:     "admin@example.com",
							ActionType:   auditlog.OrderOutgoingRevoke,
							UserAgent:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
							CreatedAt:    time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
							HasSucceeded: true,
							Data: &auditlog.RecordData{
								Delegatee:            newStringPointer("00000000000000000001"),
								Reference:            newStringPointer("test-reference"),
								PublicKeyFingerprint: newStringPointer("public-key-fingerprint"),
							},
						},
					}, nil)
			},
			req: &api.ListAuditLogsRequest{},
			want: &api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "admin@example.com",
						Action:          api.AuditLogRecord_ACTION_TYPE_ORDER_OUTGOING_REVOKE,
						OperatingSystem: "Mac OS X",
						Browser:         "Safari",
						Client:          "NLX Management",
						HasSucceeded:    true,
						Services:        []*api.AuditLogRecord_Service{},
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
						Data: &api.AuditLogRecordMetadata{
							Delegatee: &external.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
							Reference:            "test-reference",
							PublicKeyFingerprint: "public-key-fingerprint",
						},
					},
				},
			},
			wantErr: nil,
		},
		"with_a_single_log_created_via_the_browser_with_inway_metadata": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.al.
					EXPECT().
					List(gomock.Any(), numberOfAuditLogs).
					Return([]*auditlog.Record{
						{
							ID:         1,
							Username:   "admin@example.com",
							ActionType: auditlog.InwayDelete,
							UserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
							CreatedAt:  time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
							Data: &auditlog.RecordData{
								InwayName: newStringPointer("my-inway"),
							},
							HasSucceeded: true,
						},
					}, nil)
			},
			req: &api.ListAuditLogsRequest{},
			want: &api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "admin@example.com",
						Action:          api.AuditLogRecord_ACTION_TYPE_INWAY_DELETE,
						OperatingSystem: "Mac OS X",
						Browser:         "Safari",
						Client:          "NLX Management",
						Services:        []*api.AuditLogRecord_Service{},
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
						HasSucceeded:    true,
						Data: &api.AuditLogRecordMetadata{
							InwayName: "my-inway",
						},
					},
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t, nil)

			tt.setup(mocks)

			actual, err := service.ListAuditLogs(tt.ctx, tt.req)
			assert.Equal(t, tt.want, actual)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func newStringPointer(s string) *string {
	return &s
}
