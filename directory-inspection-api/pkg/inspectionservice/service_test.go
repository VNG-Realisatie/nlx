// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inspectionservice_test

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

func TestInspectionService_ListServices(t *testing.T) {
	type fields struct {
		logger   *zap.Logger
		database database.DirectoryDatabase
	}
	type args struct {
		ctx context.Context
		req *inspectionapi.ListServicesRequest
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse *inspectionapi.ListServicesResponse
		expectedError    error
	}{
		{
			name: "failed to get services from the database",
			fields: fields{
				logger: zap.NewNop(),
				database: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().RegisterOutwayVersion(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					db.EXPECT().ListServices(gomock.Any(), "TODO").Return(nil, errors.New("arbitrary error")).AnyTimes()

					return db
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &inspectionapi.ListServicesRequest{},
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		{
			name: "happy flow",
			fields: fields{
				logger: zap.NewNop(),
				database: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().RegisterOutwayVersion(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					db.EXPECT().ListServices(gomock.Any(), "TODO").Return([]*database.Service{
						{
							Name: "Dummy Service Name",
						},
					}, nil).AnyTimes()

					return db
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &inspectionapi.ListServicesRequest{},
			},
			expectedResponse: &inspectionapi.ListServicesResponse{
				Services: []*inspectionapi.ListServicesResponse_Service{
					{
						ServiceName: "Dummy Service Name",
					},
				},
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := inspectionservice.New(tt.fields.logger, tt.fields.database)
			got, err := h.ListServices(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
