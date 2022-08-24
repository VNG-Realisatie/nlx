// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain"
	mock_txlog "go.nlx.io/nlx/txlog-api/domain/txlog/storage/mock"
)

//nolint:funlen // this is a test
func TestCreateRecord(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, *mock_txlog.MockRepository)
		req     *api.CreateRecordRequest
		want    *emptypb.Empty
		wantErr error
	}{
		"without_source_org": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {},
			req: &api.CreateRecordRequest{
				DestOrganization: "00000000000000000002",
				Direction:        api.CreateRecordRequest_IN,
				ServiceName:      "test-service",
				LogrecordID:      "42",
				Data:             `{"request-path":"/get"}`,
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "invalid record: SourceOrganization: cannot be blank."),
		},
		"without_destination_org": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
				DestOrganization:   "",
				Direction:          api.CreateRecordRequest_IN,
				ServiceName:        "test-service",
				LogrecordID:        "42",
				Data:               `{"request-path":"/get"}`,
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "invalid record: DestinationOrganization: cannot be blank."),
		},
		"without_service": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
				DestOrganization:   "00000000000000000002",
				Direction:          api.CreateRecordRequest_IN,
				LogrecordID:        "42",
				Data:               `{"request-path":"/get"}`,
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "invalid record: ServiceName: cannot be blank."),
		},
		"incomplete_order_missing_reference": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
				DestOrganization:   "00000000000000000002",
				Direction:          api.CreateRecordRequest_IN,
				LogrecordID:        "42",
				Data:               `{"request-path":"/get"}`,
				Delegator:          "00000000000000000003",
				ServiceName:        "test-service",
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "invalid record: empty order reference, both the delegator and order reference should be provided"),
		},
		"incomplete_order_missing_delegator": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
				DestOrganization:   "00000000000000000002",
				Direction:          api.CreateRecordRequest_IN,
				LogrecordID:        "42",
				Data:               `{"request-path":"/get"}`,
				OrderReference:     "test-reference",
				ServiceName:        "test-service",
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "invalid record: empty delegator, both the delegator and order reference should be provided"),
		},
		"db_call_fails": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {
				model, err := domain.NewRecord(&domain.NewRecordArgs{
					SourceOrganization:      "00000000000000000001",
					DestinationOrganization: "00000000000000000002",
					Direction:               domain.IN,
					ServiceName:             "test-service",
					TransactionID:           "42",
					OrderReference:          "test-reference",
					Delegator:               "00000000000000000003",
					Data:                    []byte(`{"request-path":"/get"}`),
					CreatedAt:               fixedTestClockTime,
					DataSubjects:            map[string]string{"foo": "bar"},
				})
				require.NoError(t, err)

				mocks.
					EXPECT().
					CreateRecord(ctx, model).
					Return(errors.New("arbitrary error"))
			},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
				DestOrganization:   "00000000000000000002",
				Direction:          api.CreateRecordRequest_IN,
				ServiceName:        "test-service",
				LogrecordID:        "42",
				Delegator:          "00000000000000000003",
				OrderReference:     "test-reference",
				Data:               `{"request-path":"/get"}`,
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "storage error"),
		},
		"happy_flow_without_order": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {
				model, err := domain.NewRecord(&domain.NewRecordArgs{
					SourceOrganization:      "00000000000000000001",
					DestinationOrganization: "00000000000000000002",
					Direction:               domain.IN,
					ServiceName:             "test-service",
					TransactionID:           "42",
					Data:                    []byte(`{"request-path":"/get"}`),
					CreatedAt:               fixedTestClockTime,
					DataSubjects:            map[string]string{"foo": "bar"},
				})
				require.NoError(t, err)

				mocks.
					EXPECT().
					CreateRecord(ctx, model).
					Return(nil)
			},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
				DestOrganization:   "00000000000000000002",
				Direction:          api.CreateRecordRequest_IN,
				ServiceName:        "test-service",
				LogrecordID:        "42",
				Data:               `{"request-path":"/get"}`,
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {
				model, err := domain.NewRecord(&domain.NewRecordArgs{
					SourceOrganization:      "00000000000000000001",
					DestinationOrganization: "00000000000000000002",
					Direction:               domain.IN,
					ServiceName:             "test-service",
					TransactionID:           "42",
					Delegator:               "00000000000000000003",
					OrderReference:          "test-reference",
					Data:                    []byte(`{"request-path":"/get"}`),
					CreatedAt:               fixedTestClockTime,
					DataSubjects:            map[string]string{"foo": "bar"},
				})
				require.NoError(t, err)

				mocks.
					EXPECT().
					CreateRecord(ctx, model).
					Return(nil)
			},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
				DestOrganization:   "00000000000000000002",
				Direction:          api.CreateRecordRequest_IN,
				ServiceName:        "test-service",
				LogrecordID:        "42",
				Delegator:          "00000000000000000003",
				OrderReference:     "test-reference",
				Data:               `{"request-path":"/get"}`,
				DataSubjects: []*api.CreateRecordRequest_DataSubject{
					{
						Key:   "foo",
						Value: "bar",
					},
				},
			},
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newStorageRepository(t)
			tt.setup(context.Background(), mocks)

			got, err := service.CreateRecord(context.Background(), tt.req)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
