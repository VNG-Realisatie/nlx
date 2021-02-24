package server

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
	mock_txlogdb "go.nlx.io/nlx/management-api/pkg/txlogdb/mock"
)

func TestIsFinanceEnabled(t *testing.T) {
	service := NewManagementService(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	response, err := service.IsFinanceEnabled(context.Background(), nil)
	assert.NoError(t, err)
	assert.False(t, response.Enabled)

	service = NewManagementService(
		nil,
		nil,
		nil,
		nil,
		nil,
		mock_txlogdb.NewMockTxlogDatabase(nil),
		nil,
	)

	response, err = service.IsFinanceEnabled(context.Background(), nil)
	assert.NoError(t, err)
	assert.True(t, response.Enabled)
}

func TestDownloadFinanceExport(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		wantErr   bool
		setupMock func(*mock_database.MockConfigDatabase, *mock_txlogdb.MockTxlogDatabase)
	}{
		"returns_error_when_list_services_returns_an_error": {
			wantErr: true,
			setupMock: func(db *mock_database.MockConfigDatabase, _logDB *mock_txlogdb.MockTxlogDatabase) {
				db.EXPECT().
					ListServices(ctx).
					Return(nil, errors.New("random error"))
			},
		},

		"returns_error_when_filtering_transaction_log_records_returns_an_error": {
			wantErr: true,
			setupMock: func(db *mock_database.MockConfigDatabase, logDB *mock_txlogdb.MockTxlogDatabase) {
				db.EXPECT().
					ListServices(ctx).
					Return([]*database.Service{}, nil)

				logDB.EXPECT().
					FilterRecords(ctx, &txlogdb.Filters{
						Destination: "NLX Intermediate CA",
						Direction:   transactionlog.DirectionIn,
					}).
					Return(nil, errors.New("random error"))
			},
		},

		"returns_csv_data_when_successful": {
			setupMock: func(db *mock_database.MockConfigDatabase, logDB *mock_txlogdb.MockTxlogDatabase) {
				db.EXPECT().
					ListServices(ctx).
					Return([]*database.Service{
						{
							Name:         "Test",
							RequestCosts: 150,
						},
					}, nil)

				logDB.EXPECT().
					FilterRecords(ctx, &txlogdb.Filters{
						Destination: "NLX Intermediate CA",
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
			ctrl := gomock.NewController(t)
			txlogDB := mock_txlogdb.NewMockTxlogDatabase(ctrl)
			db := mock_database.NewMockConfigDatabase(ctrl)

			pkiDir := filepath.Join("..", "..", "..", "testing", "pki")
			bundle, _ := common_tls.NewBundleFromFiles(
				filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root.pem"),
			)

			service := NewManagementService(
				zap.NewNop(),
				nil,
				nil,
				bundle,
				db,
				txlogDB,
				nil,
			)

			tt.setupMock(db, txlogDB)

			response, err := service.DownloadFinanceExport(ctx, nil)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, response.Data)
			}
		})
	}
}
