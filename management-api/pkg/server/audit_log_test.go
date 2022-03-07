// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/outway"
	"go.nlx.io/nlx/management-api/pkg/server"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

func newManagementService(t *testing.T) (s *server.ManagementService, auditLogger *mock_auditlog.MockLogger) {
	logger := zaptest.Logger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	orgCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	internalCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.NLXTestInternal)
	require.NoError(t, err)

	assert.NoError(t, err)

	auditLogger = mock_auditlog.NewMockLogger(ctrl)

	s = server.NewManagementService(
		logger,
		mock_directory.NewMockClient(ctrl),
		nil,
		orgCert,
		internalCert,
		mock_database.NewMockConfigDatabase(ctrl),
		nil,
		auditLogger,
		management.NewClient,
		outway.NewClient,
	)

	return s, auditLogger
}

//nolint:funlen,dupl // its a unittest
func TestListAuditLogs(t *testing.T) {
	createTimestamp := func(ti time.Time) *timestamppb.Timestamp {
		return &timestamppb.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	tests := map[string]struct {
		auditLogs    []*auditlog.Record
		auditLogsErr error
		req          *emptypb.Empty
		expectedRes  *api.ListAuditLogsResponse
		expectedErr  error
	}{
		"when_error_occurs_while_retrieving_logs": {
			[]*auditlog.Record{},
			errors.New("arbitrary error"),
			&emptypb.Empty{},
			nil,
			status.New(codes.Internal, "failed to retrieve audit logs").Err(),
		},
		"when_no_logs_are_available": {
			[]*auditlog.Record{},
			nil,
			&emptypb.Empty{},
			&api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{},
			},
			nil,
		},
		"with_a_single_log_created_via_the_browser": {
			[]*auditlog.Record{
				{
					ID:         1,
					Username:   "Jane Doe",
					ActionType: auditlog.LoginSuccess,
					UserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
					CreatedAt:  time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
				},
			},
			nil,
			&emptypb.Empty{},
			&api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "Jane Doe",
						Action:          api.AuditLogRecord_loginSuccess,
						OperatingSystem: "Mac OS X",
						Browser:         "Safari",
						Client:          "NLX Management",
						Services:        []*api.AuditLogRecord_Service{},
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					},
				},
			},
			nil,
		},
		"with_a_single_log_created_via_nlxctl": {
			[]*auditlog.Record{
				{
					ID:         1,
					Username:   "",
					ActionType: auditlog.LoginSuccess,
					UserAgent:  "nlxctl/1.x (Mac OS X)",
					CreatedAt:  time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
				},
			},
			nil,
			&emptypb.Empty{},
			&api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "",
						Action:          api.AuditLogRecord_loginSuccess,
						OperatingSystem: "Mac OS X",
						Browser:         "",
						Client:          "nlxctl",
						Services:        []*api.AuditLogRecord_Service{},
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					},
				},
			},
			nil,
		},
		"with_a_single_log_created_via_the_browser_with_order_metadata": {
			[]*auditlog.Record{
				{
					ID:         1,
					Username:   "Jane Doe",
					ActionType: auditlog.OrderOutgoingRevoke,
					UserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
					CreatedAt:  time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
					Data: &auditlog.RecordData{
						Delegatee: newStringPointer("00000000000000000001"),
						Reference: newStringPointer("test-reference"),
					},
				},
			},
			nil,
			&emptypb.Empty{},
			&api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "Jane Doe",
						Action:          api.AuditLogRecord_orderOutgoingRevoke,
						OperatingSystem: "Mac OS X",
						Browser:         "Safari",
						Client:          "NLX Management",
						Services:        []*api.AuditLogRecord_Service{},
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
						Data: &api.AuditLogRecordMetadata{
							Delegatee: "00000000000000000001",
							Reference: "test-reference",
						},
					},
				},
			},
			nil,
		},
		"with_a_single_log_created_via_the_browser_with_inway_metadata": {
			[]*auditlog.Record{
				{
					ID:         1,
					Username:   "Jane Doe",
					ActionType: auditlog.InwayDelete,
					UserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
					CreatedAt:  time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
					Data: &auditlog.RecordData{
						InwayName: newStringPointer("my-inway"),
					},
				},
			},
			nil,
			&emptypb.Empty{},
			&api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{
					{
						Id:              1,
						User:            "Jane Doe",
						Action:          api.AuditLogRecord_inwayDelete,
						OperatingSystem: "Mac OS X",
						Browser:         "Safari",
						Client:          "NLX Management",
						Services:        []*api.AuditLogRecord_Service{},
						CreatedAt:       createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
						Data: &api.AuditLogRecordMetadata{
							InwayName: "my-inway",
						},
					},
				},
			},
			nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, auditLogger := newManagementService(t)
			ctx := context.Background()

			auditLogger.EXPECT().
				ListAll(ctx).
				Return(tt.auditLogs, tt.auditLogsErr)

			actual, err := service.ListAuditLogs(ctx, tt.req)
			assert.Equal(t, tt.expectedRes, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func newStringPointer(s string) *string {
	return &s
}
