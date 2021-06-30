// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inway_mock "go.nlx.io/nlx/directory-registration-api/domain/inway/mock"
	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-registration-api/pkg/registrationservice"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

const testServiceName = "Test Service Name"

//nolint:funlen // adding the tests was the first step to make the functionality testable. making it less complex is out of scope for now.
func TestDirectoryRegistrationService_RegisterInway(t *testing.T) {
	tests := map[string]struct {
		setup        func(serviceMocks)
		request      *registrationapi.RegisterInwayRequest
		wantResponse *registrationapi.RegisterInwayResponse
		wantErr      error
	}{
		"when_an_unexpected_repository_error_occurs": {
			setup: func(mocks serviceMocks) {
				mocks.ir.
					EXPECT().
					Register(gomock.Any()).
					Return(errors.New("arbitrary error"))
			},
			request: &registrationapi.RegisterInwayRequest{
				InwayName:    "my-inway",
				InwayAddress: "localhost",
			},
			wantResponse: nil,
			wantErr:      status.New(codes.Internal, "failed to register inway").Err(),
		},
		"when_specifying_an_invalid_service_name": {
			setup: func(mocks serviceMocks) {
				mocks.ir.
					EXPECT().
					Register(gomock.Any()).
					AnyTimes()
			},
			request: &registrationapi.RegisterInwayRequest{
				InwayName:    "my-inway",
				InwayAddress: "localhost",
				Services: []*registrationapi.RegisterInwayRequest_RegisterService{
					{
						Name: "../../test",
					},
				},
			},
			wantResponse: nil,
			wantErr:      status.New(codes.InvalidArgument, "validation for service named '../../test' failed: Name: must be in a valid format.").Err(),
		},
		"when_registering_an_inway_with_amount_of_services_which_exceed_the_maximum": {
			setup: func(mocks serviceMocks) {
				mocks.ir.
					EXPECT().
					Register(gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			request: &registrationapi.RegisterInwayRequest{
				InwayName:    "my-inway",
				InwayAddress: "localhost",
				Services:     generateListOfServices(251),
			},
			wantResponse: nil,
			wantErr:      status.New(codes.InvalidArgument, "inway registers more services than allowed (max. 250)").Err(),
		},
		"when_registering_an_inway_without_services": {
			setup: func(mocks serviceMocks) {
				mocks.ir.
					EXPECT().
					Register(gomock.Any()).
					AnyTimes()
			},
			request: &registrationapi.RegisterInwayRequest{
				InwayName:    "my-inway",
				InwayAddress: "localhost",
			},
			wantResponse: &registrationapi.RegisterInwayResponse{},
			wantErr:      nil,
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.ir.
					EXPECT().
					Register(gomock.Any()).
					Return(nil).
					AnyTimes()

				mocks.db.
					EXPECT().
					RegisterService(gomock.Eq(&database.RegisterServiceParams{
						OrganizationName: testOrganizationName,
						Name:             testServiceName,
						Internal:         false,
						MonthlyCosts:     500,
						RequestCosts:     100,
						OneTimeCosts:     50,
					})).Return(nil)
			},
			request: &registrationapi.RegisterInwayRequest{
				InwayName:    "my-inway",
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
			wantResponse: &registrationapi.RegisterInwayResponse{},
			wantErr:      nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.RegisterInway(context.Background(), tt.request)

			assert.Equal(t, tt.wantResponse, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

type serviceMocks struct {
	db *mock.MockDirectoryDatabase
	ir *inway_mock.MockRepository
}

func newService(t *testing.T) (*registrationservice.DirectoryRegistrationService, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		db: mock.NewMockDirectoryDatabase(ctrl),
		ir: inway_mock.NewMockRepository(ctrl),
	}

	service := registrationservice.New(
		zap.NewNop(),
		mocks.db,
		mocks.ir,
		nil,
		testGetOrganizationNameFromRequest,
	)

	return service, mocks
}

func generateListOfServices(amount int) []*registrationapi.RegisterInwayRequest_RegisterService {
	var result = make([]*registrationapi.RegisterInwayRequest_RegisterService, amount)

	for i := 0; i < amount; i++ {
		result[i] = &registrationapi.RegisterInwayRequest_RegisterService{
			Name: fmt.Sprintf("Service number %d", i+1),
		}
	}

	return result
}
