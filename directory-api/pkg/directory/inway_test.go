// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory_test

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

	"go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	storage_mock "go.nlx.io/nlx/directory-api/domain/directory/storage/mock"
	"go.nlx.io/nlx/directory-api/pkg/directory"
)

const testServiceName = "Test Service Name"

//nolint:funlen // adding the tests was the first step to make the functionality testable. making it less complex is out of scope for now.
func TestRegisterInway(t *testing.T) {
	tests := map[string]struct {
		setup        func(serviceMocks)
		request      *directoryapi.RegisterInwayRequest
		wantResponse *directoryapi.RegisterInwayResponse
		wantErr      error
	}{
		"when_an_unexpected_repository_error_occurs": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					RegisterInway(gomock.Any()).
					Return(errors.New("arbitrary error"))
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost",
			},
			wantResponse: nil,
			wantErr:      status.New(codes.Internal, "failed to register inway").Err(),
		},
		"when_specifying_an_invalid_service_name": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					RegisterInway(gomock.Any()).
					AnyTimes()

				mocks.r.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost",
				Services: []*directoryapi.RegisterInwayRequest_RegisterService{
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
				mocks.r.
					EXPECT().
					RegisterInway(gomock.Any()).
					Return(nil).
					AnyTimes()

				mocks.r.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost",
				Services:     generateListOfServices(251),
			},
			wantResponse: nil,
			wantErr:      status.New(codes.InvalidArgument, "inway registers more services than allowed (max. 250)").Err(),
		},
		"when_registering_an_inway_without_services": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					RegisterInway(gomock.Any()).
					AnyTimes()

				mocks.r.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost",
			},
			wantResponse: &directoryapi.RegisterInwayResponse{},
			wantErr:      nil,
		},
		"happy_flow_is_org_inway": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					RegisterInway(gomock.Any()).
					Return(nil).
					AnyTimes()

				mocks.r.
					EXPECT().
					RegisterService(gomock.Any()).
					Return(nil)

				mocks.r.
					EXPECT().
					SetOrganizationInway(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress:        "localhost",
				IsOrganizationInway: true,
				Services: []*directoryapi.RegisterInwayRequest_RegisterService{
					{
						Name:         testServiceName,
						MonthlyCosts: 500,
						RequestCosts: 100,
						OneTimeCosts: 50,
					},
				},
			},
			wantResponse: &directoryapi.RegisterInwayResponse{},
			wantErr:      nil,
		},
		"happy_flow_not_org_inway": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					RegisterInway(gomock.Any()).
					Return(nil).
					AnyTimes()

				mocks.r.
					EXPECT().
					RegisterService(gomock.Any()).Return(nil)

				mocks.r.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost",
				Services: []*directoryapi.RegisterInwayRequest_RegisterService{
					{
						Name:         testServiceName,
						MonthlyCosts: 500,
						RequestCosts: 100,
						OneTimeCosts: 50,
					},
				},
			},
			wantResponse: &directoryapi.RegisterInwayResponse{},
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
	r *storage_mock.MockRepository
}

const testOrganizationName = "Test Organization Name"
const testOrganizationSerialNumber = "01234567890123456789"

func testGetOrganizationInformationFromRequest(context.Context) (*tls.OrganizationInformation, error) {
	return &tls.OrganizationInformation{
		Name:         testOrganizationName,
		SerialNumber: testOrganizationSerialNumber,
	}, nil
}
func newService(t *testing.T) (*directory.DirectoryService, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		r: storage_mock.NewMockRepository(ctrl),
	}

	service := directory.New(
		zap.NewNop(),
		mocks.r,
		nil,
		testGetOrganizationInformationFromRequest,
	)

	return service, mocks
}

func generateListOfServices(amount int) []*directoryapi.RegisterInwayRequest_RegisterService {
	var result = make([]*directoryapi.RegisterInwayRequest_RegisterService, amount)

	for i := 0; i < amount; i++ {
		result[i] = &directoryapi.RegisterInwayRequest_RegisterService{
			Name: fmt.Sprintf("Service number %d", i+1),
		}
	}

	return result
}
