// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
	mock_txlogdb "go.nlx.io/nlx/management-api/pkg/txlogdb/mock"
)

func TestIsFinanceEnabled(t *testing.T) {
	tests := map[string]struct {
		disableTxLogDB bool
		enabled        bool
	}{
		"returns_false_when_database_is_nil": {
			disableTxLogDB: true,
			enabled:        false,
		},

		"returns_true_when_database_is_not_nil": {
			disableTxLogDB: false,
			enabled:        true,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			var service *server.ManagementService
			if tt.disableTxLogDB {
				service, _, _ = newServiceWithoutTXLog(t)
			} else {
				service, _, _ = newService(t)
			}

			response, err := service.IsFinanceEnabled(context.Background(), nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.enabled, response.Enabled)
		})
	}
}

func TestDownloadFinanceExport(t *testing.T) {
	tests := map[string]struct {
		ctx       context.Context
		wantErr   error
		setupMock func(*mock_database.MockConfigDatabase, *mock_txlogdb.MockTxlogDatabase)
	}{
		"missing_required_permission": {
			ctx:       testCreateUserWithoutPermissionsContext(),
			wantErr:   status.Error(codes.PermissionDenied, "user needs the permission \"permissions.finance_report.read\" to execute this request"),
			setupMock: func(db *mock_database.MockConfigDatabase, _logDB *mock_txlogdb.MockTxlogDatabase) {},
		},
		"returns_error_when_list_services_returns_an_error": {
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.Internal, "database error"),
			setupMock: func(db *mock_database.MockConfigDatabase, _logDB *mock_txlogdb.MockTxlogDatabase) {
				db.EXPECT().
					ListServices(gomock.Any()).
					Return(nil, errors.New("random error"))
			},
		},
		"returns_error_when_filtering_transaction_log_records_returns_an_error": {
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.Internal, "database error"),
			setupMock: func(db *mock_database.MockConfigDatabase, logDB *mock_txlogdb.MockTxlogDatabase) {
				db.EXPECT().
					ListServices(gomock.Any()).
					Return([]*database.Service{}, nil)

				logDB.EXPECT().
					FilterRecords(gomock.Any(), &txlogdb.Filters{
						Destination: "00000000000000000001",
						Direction:   transactionlog.DirectionIn,
					}).
					Return(nil, fmt.Errorf("arbitrary error"))
			},
		},
		"returns_csv_data_when_successful": {
			ctx: testCreateAdminUserContext(),
			setupMock: func(db *mock_database.MockConfigDatabase, logDB *mock_txlogdb.MockTxlogDatabase) {
				db.EXPECT().
					ListServices(gomock.Any()).
					Return([]*database.Service{
						{
							Name:         "Test",
							RequestCosts: 150,
						},
					}, nil)

				logDB.EXPECT().
					FilterRecords(gomock.Any(), &txlogdb.Filters{
						Destination: "00000000000000000001",
						Direction:   transactionlog.DirectionIn,
					}).
					Return([]txlogdb.Record{
						{
							Direction:    transactionlog.DirectionIn,
							Source:       "Org1",
							Destination:  "Me",
							ServiceName:  "Test",
							RequestCount: 100,
							CreatedAt:    time.Now(),
						},
					}, nil)
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setupMock(mocks.db, mocks.dbTxLog)

			_, err := service.DownloadFinanceExport(tt.ctx, nil)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
