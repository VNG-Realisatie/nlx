// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inspectionservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gogo/protobuf/types"
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

const testOrganizationName = "Test Organization Name"

func testGetOrganizationNameFromRequest(ctx context.Context) (string, error) {
	return testOrganizationName, nil
}

func TestInspectionService_ListOrganizations(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		expectedResponse *inspectionapi.ListOrganizationsResponse
		expectedError    error
	}{
		{
			name: "failed to get organizations from the db",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().ListOrganizations(gomock.Any()).Return(nil, errors.New("arbitrary error"))

				return db
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Database error.").Err(),
		},
		{
			name: "happy flow",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().ListOrganizations(gomock.Any()).Return([]*database.Organization{
					{
						Name: "Dummy Organization Name",
					},
				}, nil)

				return db
			},
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

			h := inspectionservice.New(zap.NewNop(), tt.db(ctrl), testGetOrganizationNameFromRequest)
			got, err := h.ListOrganizations(context.Background(), &types.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
