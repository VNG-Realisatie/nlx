// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/outway"
	"go.nlx.io/nlx/management-api/pkg/server"
	"go.nlx.io/nlx/management-api/pkg/txlog"
	mock_txlog "go.nlx.io/nlx/management-api/pkg/txlog/mock"
	txlogapi "go.nlx.io/nlx/txlog-api/api"
)

func TestIsTXLogEnabled(t *testing.T) {
	tests := map[string]struct {
		client  txlog.Client
		enabled bool
	}{
		"returns_false_when_txlog_client_is_nil": {
			client:  nil,
			enabled: false,
		},

		"returns_true_when_txlog_client_is_not_nil": {
			client:  mock_txlog.NewMockClient(nil),
			enabled: true,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service := server.NewManagementService(
				&server.NewManagementServiceArgs{
					Logger:                     nil,
					DirectoryClient:            nil,
					TxlogClient:                tt.client,
					OrgCert:                    nil,
					InternalCert:               nil,
					ConfigDatabase:             nil,
					TxlogDatabase:              nil,
					AuditLogger:                nil,
					CreateManagementClientFunc: management.NewClient,
					CreateOutwayClientFunc:     outway.NewClient,
					Clock:                      nil,
				})

			response, err := service.IsTXLogEnabled(context.Background(), nil)
			assert.NoError(t, err)
			assert.Equal(t, tt.enabled, response.Enabled)
		})
	}
}

//nolint:funlen // this is a test
func TestTXLogListRecords(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		setup   func(context.Context, *txlogServiceMocks)
		want    *api.TXLogListRecordsResponse
		wantErr error
	}{
		"when_directory_client_errors": {
			setup: func(ctx context.Context, mocks *txlogServiceMocks) {
				mocks.d.
					EXPECT().
					ListOrganizations(ctx, &emptypb.Empty{}).
					Return(nil, errors.New("arbitrary error"))
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "txlog error"),
		},
		"when_txlogclient_errors": {
			setup: func(ctx context.Context, mocks *txlogServiceMocks) {
				mocks.d.
					EXPECT().
					ListOrganizations(ctx, &emptypb.Empty{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.m.
					EXPECT().
					ListRecords(ctx, &emptypb.Empty{}).
					Return(nil, errors.New("arbitrary error"))
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "txlog error"),
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks *txlogServiceMocks) {
				mocks.d.
					EXPECT().
					ListOrganizations(ctx, &emptypb.Empty{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
							{
								SerialNumber: "00000000000000000002",
								Name:         "Organization Two",
							},
							{
								SerialNumber: "00000000000000000003",
								Name:         "Organization Three",
							},
						}}, nil)

				mocks.m.
					EXPECT().
					ListRecords(ctx, &emptypb.Empty{}).
					Return(&txlogapi.ListRecordsResponse{
						Records: []*txlogapi.ListRecordsResponse_Record{
							{
								Source: &txlogapi.ListRecordsResponse_Record_Organization{
									SerialNumber: "00000000000000000001",
								},
								Destination: &txlogapi.ListRecordsResponse_Record_Organization{
									SerialNumber: "00000000000000000002",
								},
								Direction: txlogapi.ListRecordsResponse_Record_DIRECTION_IN,
								Service: &txlogapi.ListRecordsResponse_Record_Service{
									Name: "test-service",
								},
								Order: &txlogapi.ListRecordsResponse_Record_Order{
									Delegator: "00000000000000000003",
									Reference: "test-reference",
								},
								Data:          `{"test":"data"}`,
								TransactionID: "abcd",
								CreatedAt:     timestamppb.New(now),
							},
						},
					}, nil)
			},
			want: &api.TXLogListRecordsResponse{
				Records: []*api.TXLogRecord{
					{
						Source: &api.TXLogOrganization{
							SerialNumber: "00000000000000000001",
							Name:         "Organization One",
						},
						Destination: &api.TXLogOrganization{
							SerialNumber: "00000000000000000002",
							Name:         "Organization Two",
						},
						Direction: api.TXLogDirection_TX_LOG_DIRECTION_IN,
						Service: &api.TXLogService{
							Name: "test-service",
						},
						Order: &api.TXLogOrder{
							Delegator: &api.TXLogOrganization{
								SerialNumber: "00000000000000000003",
								Name:         "Organization Three",
							},
							Reference: "test-reference",
						},
						Data:          `{"test":"data"}`,
						TransactionId: "abcd",
						CreatedAt:     timestamppb.New(now),
					},
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newTXLogService(t)
			tt.setup(context.Background(), mocks)

			got, err := service.ListRecords(context.Background(), &emptypb.Empty{})

			t.Log(err)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type txlogServiceMocks struct {
	m *mock_txlog.MockClient
	d *mock_directory.MockClient
}

func newTXLogService(t *testing.T) (s *server.TXLogService, mocks *txlogServiceMocks) {
	logger := zaptest.NewLogger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	mocks = &txlogServiceMocks{
		m: mock_txlog.NewMockClient(ctrl),
		d: mock_directory.NewMockClient(ctrl),
	}

	s = server.NewTXLogService(logger, mocks.m, mocks.d)

	return
}
