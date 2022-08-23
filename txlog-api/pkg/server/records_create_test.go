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
			wantErr: status.Error(codes.InvalidArgument, "source organization: error validating organization serial number: cannot be empty"),
		},
		"without_destination_org": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {},
			req: &api.CreateRecordRequest{
				SourceOrganization: "00000000000000000001",
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
			wantErr: status.Error(codes.InvalidArgument, "destination organization: error validating organization serial number: cannot be empty"),
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
			wantErr: status.Error(codes.InvalidArgument, "service: cannot be blank"),
		},
		"without_order": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {
				model, err := domain.NewRecord(&domain.NewRecordArgs{
					Source:        createNewOrganization(t, "00000000000000000001"),
					Destination:   createNewOrganization(t, "00000000000000000002"),
					Direction:     domain.IN,
					Service:       createNewService(t, "test-service"),
					TransactionID: "42",
					Data:          []byte(`{"request-path":"/get"}`),
					CreatedAt:     fixedTestClockTime,
					DataSubjects:  map[string]string{"foo": "bar"},
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
		"db_call_fails": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {
				model, err := domain.NewRecord(&domain.NewRecordArgs{
					Source:        createNewOrganization(t, "00000000000000000001"),
					Destination:   createNewOrganization(t, "00000000000000000002"),
					Direction:     domain.IN,
					Service:       createNewService(t, "test-service"),
					TransactionID: "42",
					Order:         createNewOrder(t, "00000000000000000003", "test-reference"),
					Data:          []byte(`{"request-path":"/get"}`),
					CreatedAt:     fixedTestClockTime,
					DataSubjects:  map[string]string{"foo": "bar"},
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
		"happy_flow": {
			setup: func(ctx context.Context, mocks *mock_txlog.MockRepository) {
				model, err := domain.NewRecord(&domain.NewRecordArgs{
					Source:        createNewOrganization(t, "00000000000000000001"),
					Destination:   createNewOrganization(t, "00000000000000000002"),
					Direction:     domain.IN,
					Service:       createNewService(t, "test-service"),
					TransactionID: "42",
					Order:         createNewOrder(t, "00000000000000000003", "test-reference"),
					Data:          []byte(`{"request-path":"/get"}`),
					CreatedAt:     fixedTestClockTime,
					DataSubjects:  map[string]string{"foo": "bar"},
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
