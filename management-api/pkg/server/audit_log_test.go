// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"github.com/fgrosse/zaptest"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path/filepath"
	"testing"

	"go.nlx.io/nlx/management-api/api"
)

func newManagementService(t *testing.T) (s *server.ManagementService, auditLogger *mock_auditlog.MockLogger) {
	logger := zaptest.Logger(t)
	proc := process.NewProcess(logger)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")
	bundle, err := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	assert.NoError(t, err)

	auditLogger = mock_auditlog.NewMockLogger(ctrl)

	s = server.NewManagementService(
		logger,
		proc,
		mock_directory.NewMockClient(ctrl),
		bundle,
		mock_database.NewMockConfigDatabase(ctrl),
		auditLogger,
	)

	return
}

//nolint:funlen // its a unittest
func TestListAuditLogs(t *testing.T) {
	tests := map[string]struct {
		auditLogs    []*api.AuditLogRecord
		auditLogsErr error
		req          *types.Empty
		expectedRes  *api.ListAuditLogsResponse
		expectedErr  error
	}{
		"when_error_occurs_while_retrieving_logs": {
			[]*api.AuditLogRecord{},
			errors.New("arbitrary error"),
			&types.Empty{},
			nil,
			status.New(codes.Internal, "failed to retrieve audit logs").Err(),
		},
		"when_no_logs_are_available": {
			[]*api.AuditLogRecord{},
			nil,
			&types.Empty{},
			&api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{},
			},
			nil,
		},
		"with_a_single_log": {
			[]*api.AuditLogRecord{},
			nil,
			&types.Empty{},
			&api.ListAuditLogsResponse{
				AuditLogs: []*api.AuditLogRecord{},
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
