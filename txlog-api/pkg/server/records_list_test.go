// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain"
	mock_txlog "go.nlx.io/nlx/txlog-api/domain/txlog/storage/mock"
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
						SourceOrganization:      "0001",
						DestinationOrganization: "0002",
						Direction:               domain.IN,
						ServiceName:             "test-service",
						Delegator:               "0003",
						OrderReference:          "test-reference",
						Data:                    []byte(`{"test": "data"}`),
						CreatedAt:               now,
						TransactionID:           "abcde",
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
				Records: []*api.ListRecordsResponse_Record{
					{
						Source: &api.ListRecordsResponse_Record_Organization{
							SerialNumber: "0001",
						},
						Destination: &api.ListRecordsResponse_Record_Organization{
							SerialNumber: "0002",
						},
						Direction: api.ListRecordsResponse_Record_IN,
						Service: &api.ListRecordsResponse_Record_Service{
							Name: "test-service",
						},
						Order: &api.ListRecordsResponse_Record_Order{
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
