// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/server"
	mock_txlog "go.nlx.io/nlx/management-api/pkg/txlog/mock"
	txlogapi "go.nlx.io/nlx/txlog-api/api"
)

func TestTXLogListRecords(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		setup   func(context.Context, *mock_txlog.MockClient)
		want    *api.TXLogListRecordsResponse
		wantErr error
	}{
		"when_txlog_error": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockClient) {
				mocks.
					EXPECT().
					ListRecords(ctx, &emptypb.Empty{}).
					Return(nil, errors.New("arbitrary error"))
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "txlog error"),
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockClient) {
				mocks.
					EXPECT().
					ListRecords(ctx, &emptypb.Empty{}).
					Return(&txlogapi.ListRecordsResponse{
						Records: []*txlogapi.Record{
							{
								Source: &txlogapi.Organization{
									SerialNumber: "0001",
								},
								Destination: &txlogapi.Organization{
									SerialNumber: "0002",
								},
								Direction: txlogapi.Direction_IN,
								Service: &txlogapi.Service{
									Name: "test-service",
								},
								Order: &txlogapi.Order{
									Delegator: "0003",
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
							SerialNumber: "0001",
						},
						Destination: &api.TXLogOrganization{
							SerialNumber: "0002",
						},
						Direction: api.TXLogDirection_IN,
						Service: &api.TXLogService{
							Name: "test-service",
						},
						Order: &api.TXLogOrder{
							Delegator: "0003",
							Reference: "test-reference",
						},
						Data:          `{"test":"data"}`,
						TransactionID: "abcd",
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

func newTXLogService(t *testing.T) (s *server.TXLogService, m *mock_txlog.MockClient) {
	logger := zaptest.Logger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	m = mock_txlog.NewMockClient(ctrl)

	s = server.NewTXLogService(logger, m)

	return
}
