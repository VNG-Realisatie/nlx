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
		setup            func(serviceMocks)
		req              *registrationapi.RegisterInwayRequest
		expectedResponse *registrationapi.RegisterInwayResponse
		expectedError    error
	}{
		{
			name: "failed to communicate with the database",
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					InsertAvailability(gomock.Any()).
					Return(errors.New("arbitrary error"))
			},
			req: &registrationapi.RegisterInwayRequest{
				InwayAddress: "localhost",
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
			name: "invalid params",
			req: &registrationapi.RegisterInwayRequest{
				InwayAddress: "localhost",
				Services: []*registrationapi.RegisterInwayRequest_RegisterService{
					{
						Name: "../../test",
					},
				},
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.InvalidArgument, "validation failed: ServiceName: must be in a valid format.").Err(),
		},
		{
			name: "happy flow",
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().InsertAvailability(gomock.Eq(&database.InsertAvailabilityParams{
					OrganizationName:    testOrganizationName,
					ServiceName:         testServiceName,
					ServiceInternal:     false,
					RequestInwayAddress: "localhost",
					NlxVersion:          "unknown",
					MonthlyCosts:        500,
					RequestCosts:        100,
					OneTimeCosts:        50,
				})).Return(nil)
			},
			req: &registrationapi.RegisterInwayRequest{
				InwayAddress: "localhost",
				Services: []*registrationapi.RegisterInwayRequest_RegisterService{
					{
						Name:         testServiceName,
						MonthlyCosts: int32(500),
						RequestCosts: int32(100),
						OneTimeCosts: int32(50),
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
			service, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.RegisterInway(context.Background(), tt.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

type serviceMocks struct {
	db *mock.MockDirectoryDatabase
}

func newService(t *testing.T) (*registrationservice.DirectoryRegistrationService, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		db: mock.NewMockDirectoryDatabase(ctrl),
	}

	service := registrationservice.New(
		zap.NewNop(),
		mocks.db,
		nil,
		testGetOrganizationNameFromRequest,
	)

	return service, mocks
}
