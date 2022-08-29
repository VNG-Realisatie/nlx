// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpc_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain/record"
	txlog_mock "go.nlx.io/nlx/txlog-api/domain/record/mock"
)

func TestListRecords(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		setup   func(context.Context, *txlog_mock.MockRepository)
		want    *api.ListRecordsResponse
		wantErr error
	}{
		"database_error": {
			setup: func(ctx context.Context, mocks *txlog_mock.MockRepository) {
				mocks.
					EXPECT().
					ListRecords(ctx, gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "storage error"),
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks *txlog_mock.MockRepository) {
				records := []*record.Record{
					mustNewRecord(t, &record.NewRecordArgs{
						SourceOrganization:      "0001",
						DestinationOrganization: "0002",
						Direction:               record.IN,
						ServiceName:             "test-service",
						Delegator:               "0003",
						OrderReference:          "test-reference",
						Data:                    []byte(`{"test": "data"}`),
						CreatedAt:               now,
						TransactionID:           "abcde",
					}),
				}

				mocks.
					EXPECT().
					ListRecords(ctx, uint(100)).
					Return(records, nil)
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
			service, mocks := newService(t)
			tt.setup(context.Background(), mocks)

			got, err := service.ListRecords(context.Background(), &emptypb.Empty{})

			t.Log(err)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func mustNewRecord(t *testing.T, args *record.NewRecordArgs) *record.Record {
	result, err := record.NewRecord(args)
	assert.NoError(t, err)

	return result
}
