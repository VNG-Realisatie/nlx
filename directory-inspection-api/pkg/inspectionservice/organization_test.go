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
		logger                         *zap.Logger
		db                             func(ctrl *gomock.Controller) database.DirectoryDatabase
		getOrganisationNameFromRequest func(ctx context.Context) (string, error)
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
			name: "failed to get organizations from the db",
			fields: fields{
				logger: zap.NewNop(),
				db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
					db := mock.NewMockDirectoryDatabase(ctrl)
					db.EXPECT().ListOrganizations(gomock.Any()).Return(nil, errors.New("arbitrary error"))

					return db
				},
				getOrganisationNameFromRequest: func(ctx context.Context) (string, error) {
					return testOrganizationName, nil
				},
			},
			args:             args{},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Database error.").Err(),
		},
		{
			name: "happy flow",
			fields: fields{
				logger: zap.NewNop(),
				db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
					db := mock.NewMockDirectoryDatabase(ctrl)
					db.EXPECT().ListOrganizations(gomock.Any()).Return([]*database.Organization{
						{
							Name: "Dummy Organization Name",
						},
					}, nil)

					return db
				},
				getOrganisationNameFromRequest: func(ctx context.Context) (string, error) {
					return testOrganizationName, nil
				},
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := inspectionservice.New(tt.fields.logger, tt.fields.db(ctrl), tt.fields.getOrganisationNameFromRequest)
			got, err := h.ListOrganizations(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
