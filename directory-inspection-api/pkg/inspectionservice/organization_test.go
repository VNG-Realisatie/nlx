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
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

const testOrganizationSerialNumber = "01234567890123456789"

func testGetOrganisationInformationFromRequest(ctx context.Context) (*tls.OrganizationInformation, error) {
	return &tls.OrganizationInformation{
		SerialNumber: testOrganizationSerialNumber,
		Name:         "test organization name",
	}, nil
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
						SerialNumber: "01234567890123456789",
						Name:         "Dummy Organization Name",
					},
				}, nil)

				return db
			},
			expectedResponse: &inspectionapi.ListOrganizationsResponse{
				Organizations: []*inspectionapi.Organization{
					{
						SerialNumber: "01234567890123456789",
						Name:         "Dummy Organization Name",
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

			h := inspectionservice.New(zap.NewNop(), tt.db(ctrl), testGetOrganisationInformationFromRequest)
			got, err := h.ListOrganizations(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestInspectionService_GetOrganizationInway(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		req              *inspectionapi.GetOrganizationInwayRequest
		expectedResponse *inspectionapi.GetOrganizationInwayResponse
		expectedError    error
	}{
		{
			name: "organization_not_found",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					GetOrganizationInwayAddress(gomock.Any(), testOrganizationSerialNumber).
					Return("", database.ErrNoOrganization)

				return db
			},
			req:              &inspectionapi.GetOrganizationInwayRequest{OrganizationSerialNumber: testOrganizationSerialNumber},
			expectedResponse: nil,
			expectedError:    status.New(codes.NotFound, "organization has no inway").Err(),
		},
		{
			name: "organization_is_empty",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					GetOrganizationInwayAddress(gomock.Any(), testOrganizationSerialNumber).
					Times(0)

				return db
			},
			req:              &inspectionapi.GetOrganizationInwayRequest{OrganizationSerialNumber: ""},
			expectedResponse: nil,
			expectedError:    status.New(codes.InvalidArgument, "organization serial number is empty").Err(),
		},
		{
			name: "happy_flow",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					GetOrganizationInwayAddress(gomock.Any(), testOrganizationSerialNumber).
					Return("inway.nlx.local", nil)

				return db
			},
			req: &inspectionapi.GetOrganizationInwayRequest{OrganizationSerialNumber: testOrganizationSerialNumber},
			expectedResponse: &inspectionapi.GetOrganizationInwayResponse{
				Address: "inway.nlx.local",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := inspectionservice.New(zap.NewNop(), tt.db(ctrl), testGetOrganisationInformationFromRequest)
			got, err := h.GetOrganizationInway(context.Background(), tt.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
