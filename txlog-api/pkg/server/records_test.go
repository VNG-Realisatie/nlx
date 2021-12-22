// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain"
	mock_txlog "go.nlx.io/nlx/txlog-api/domain/txlog/storage/mock"
	"go.nlx.io/nlx/txlog-api/pkg/server"
)

func TestListRecords(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		setup   func(context.Context, *mock_txlog.MockRepository)
		want    *api.ListRecordsResponse
		wantErr error
	}{
		"happy_flow": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {
				args := []*domain.NewRecordArgs{
					{
						Source:        createNewOrganization(t, "0001"),
						Destination:   createNewOrganization(t, "0002"),
						Direction:     domain.IN,
						Service:       createNewService(t, "test-service"),
						Order:         createNewOrder(t, "0003", "test-reference"),
						Data:          []byte(`{"test": "data"}`),
						CreatedAt:     now,
						TransactionID: "abcde",
					},
				}

				models := make([]*domain.Record, len(args))
				for i, arg := range args {
					var err error
					models[i], err = domain.NewRecord(arg)
					require.NoError(t, err)
				}

				mocks.
					EXPECT().
					ListRecords(ctx, int32(100)).
					Return(models, nil)
			},
			want: &api.ListRecordsResponse{
				Records: []*api.Record{
					{
						Source: &api.Organization{
							SerialNumber: "0001",
						},
						Destination: &api.Organization{
							SerialNumber: "0002",
						},
						Direction: api.Direction_IN,
						Service: &api.Service{
							Name: "test-service",
						},
						Order: &api.Order{
							Delegator: "0003",
							Reference: "test-reference",
						},
						Data:          `{"test": "data"}`,
						TransactionID: "abcde",
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
			service, mocks := newStorageRepository(t)
			tt.setup(context.Background(), mocks)

			got, err := service.ListRecords(context.Background(), &emptypb.Empty{})

			t.Log(err)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func newStorageRepository(t *testing.T) (s *server.TXLogService, m *mock_txlog.MockRepository) {
	logger := zaptest.Logger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	m = mock_txlog.NewMockRepository(ctrl)

	s = server.NewTXLogService(logger, m)

	return
}

func createNewOrganization(t *testing.T, serialNumber string) *domain.Organization {
	m, err := domain.NewOrganization(serialNumber)
	require.NoError(t, err)

	return m
}

func createNewService(t *testing.T, name string) *domain.Service {
	m, err := domain.NewService(name)
	require.NoError(t, err)

	return m
}

func createNewOrder(t *testing.T, delegator, reference string) *domain.Order {
	m, err := domain.NewOrder(&domain.NewOrderArgs{
		Delegator: delegator,
		Reference: reference,
	})
	require.NoError(t, err)

	return m
}