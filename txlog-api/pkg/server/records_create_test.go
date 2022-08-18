// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain"
	mock_txlog "go.nlx.io/nlx/txlog-api/domain/txlog/storage/mock"
)

func TestCreateRecord(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, *mock_txlog.MockRepository)
		want    *emptypb.Empty
		wantErr error
	}{
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
					CreateRecord(ctx, gomock.Eq(model)).
					Return(nil)
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

			got, err := service.CreateRecord(
				context.Background(),
				&api.CreateRecordRequest{
					SourceOrganization: "00000000000000000001",
					DestOrganization:   "00000000000000000002",
					Direction:          api.CreateRecordRequest_IN,
					ServiceName:        "test-service",
					LogrecordID:        "42",
					Delegator:          "00000000000000000003",
					OrderReference:     "test-reference",
					Data:               `{"request-path":"/get"}`,
					DataSubjects:       map[string]string{"foo": "bar"},
				})

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
