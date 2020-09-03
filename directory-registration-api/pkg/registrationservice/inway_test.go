// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-registration-api/pkg/registrationservice"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

const testServiceName = "Test Service Name"

//nolint:funlen // adding the tests was the first step to make the functionality testable. making it less complex is out of scope for now.
func TestDirectoryRegistrationService_RegisterInway(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		req              *registrationapi.RegisterInwayRequest
		expectedResponse *registrationapi.RegisterInwayResponse
		expectedError    error
	}{
		{
			name: "failed to communicate with the database",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().InsertAvailability(gomock.Any()).Return(errors.New("arbitrary error"))

				return db
			},
			req: &registrationapi.RegisterInwayRequest{
				InwayAddress: "",
				Services: []*registrationapi.RegisterInwayRequest_RegisterService{
					{
						Name: testServiceName,
					},
				},
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		{
			name: "happy flow",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().InsertAvailability(gomock.Eq(&database.InsertAvailabilityParams{
					OrganizationName: testOrganizationName,
					ServiceName:      testServiceName,
					ServiceInternal:  false,
					NlxVersion:       "unknown",
				})).Return(nil)

				return db
			},
			req: &registrationapi.RegisterInwayRequest{
				InwayAddress: "",
				Services: []*registrationapi.RegisterInwayRequest_RegisterService{
					{
						Name: testServiceName,
					},
				},
			},
			expectedResponse: &registrationapi.RegisterInwayResponse{},
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := registrationservice.New(zap.NewNop(), tt.db(ctrl), nil, testGetOrganizationNameFromRequest)
			got, err := h.RegisterInway(context.Background(), tt.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
