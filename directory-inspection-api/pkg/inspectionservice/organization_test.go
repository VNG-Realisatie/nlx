// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inspectionservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

func TestInspectionService_ListOrganizations(t *testing.T) {
	type fields struct {
		logger   *zap.Logger
		database database.DirectoryDatabase
	}
	type args struct {
		ctx context.Context
		req *inspectionapi.ListOrganizationsRequest
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse *inspectionapi.ListOrganizationsResponse
		expectedError    error
	}{
		{
			name: "failed to get organizations from the database",
			fields: fields{
				logger: zap.NewNop(),
				database: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().ListOrganizations(gomock.Any()).Return(nil, errors.New("arbitrary error")).AnyTimes()

					return db
				}(),
			},
			args:             args{},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Database error.").Err(),
		},
		{
			name: "happy flow",
			fields: fields{
				logger: zap.NewNop(),
				database: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().ListOrganizations(gomock.Any()).Return([]*database.Organization{
						{
							Name: "Dummy Organization Name",
						},
					}, nil).AnyTimes()

					return db
				}(),
			},
			args: args{},
			expectedResponse: &inspectionapi.ListOrganizationsResponse{
				Organizations: []*inspectionapi.ListOrganizationsResponse_Organization{
					{
						Name: "Dummy Organization Name",
					},
				},
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			h := inspectionservice.New(tt.fields.logger, tt.fields.database)
			got, err := h.ListOrganizations(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
