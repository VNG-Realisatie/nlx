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

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

func TestInspectionService_List(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		expectedResponse *inspectionapi.ListServicesResponse
		expectedError    error
	}{
		{
			name: "failed to get services from the db",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().RegisterOutwayVersion(gomock.Any(), gomock.Any()).Times(0)
				db.EXPECT().ListServices(gomock.Any(), testOrganizationSerialNumber).Return(nil, errors.New("arbitrary error"))

				return db
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "db error").Err(),
		},
		{
			name: "happy flow",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().RegisterOutwayVersion(gomock.Any(), gomock.Any()).Times(0)
				db.EXPECT().ListServices(gomock.Any(), testOrganizationSerialNumber).Return([]*database.Service{
					{
						Name: "Dummy Service Name",
						Costs: &database.ServiceCosts{
							Monthly: 1,
							Request: 5,
							OneTime: 250,
						},
					},
				}, nil)

				return db
			},
			expectedResponse: &inspectionapi.ListServicesResponse{
				Services: []*inspectionapi.ListServicesResponse_Service{
					{
						Name: "Dummy Service Name",
						Costs: &inspectionapi.Costs{
							Monthly: 1,
							Request: 5,
							OneTime: 250,
						},
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
			got, err := h.ListServices(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
