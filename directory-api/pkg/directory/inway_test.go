// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
	storage_mock "go.nlx.io/nlx/directory-api/domain/directory/storage/mock"
	"go.nlx.io/nlx/directory-api/pkg/directory"
)

const testOrganizationName = "Test Organization Name"
const testOrganizationSerialNumber = "01234567890123456789"
const testServiceName = "Test Service Name"
const testNlxVersion127 = "v0.127.0"
const testNlxVersion128Prerelease = "v0.128.0-review-156525df5"
const testNlxVersion128 = "v0.128.0"

type testClock struct {
	timeToReturn time.Time
}

func (c *testClock) Now() time.Time {
	return c.timeToReturn
}

//nolint:funlen // adding the tests was the first step to make the functionality testable. making it less complex is out of scope for now.
func TestRegisterInway(t *testing.T) {
	now := time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC)

	tests := map[string]struct {
		setup        func(serviceMocks)
		context      context.Context
		request      *directoryapi.RegisterInwayRequest
		wantResponse *directoryapi.RegisterInwayResponse
		wantErr      error
	}{
		"when_an_unexpected_repository_error_occurs": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion127,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:      "localhost:443",
					Organization: organization,
					NlxVersion:   testNlxVersion127,
					CreatedAt:    now,
					UpdatedAt:    now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					Return(errors.New("arbitrary error"))
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost:443",
			},
			wantResponse: nil,
			wantErr:      status.Error(codes.Internal, "failed to register inway"),
		},
		"when_specifying_an_invalid_service_name": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion127,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:      "localhost:443",
					Organization: organization,
					NlxVersion:   testNlxVersion127,
					CreatedAt:    now,
					UpdatedAt:    now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					AnyTimes()

				mocks.repository.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost:443").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost:443",
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
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion127,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:      "localhost:443",
					Organization: organization,
					NlxVersion:   testNlxVersion127,
					CreatedAt:    now,
					UpdatedAt:    now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					Return(nil).
					AnyTimes()

				mocks.repository.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost:443").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost:443",
				Services:     generateListOfServices(251),
			},
			wantResponse: nil,
			wantErr:      status.Error(codes.InvalidArgument, "inway registers more services than allowed (max. 250)"),
		},
		"when_registering_an_inway_without_services": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion127,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:      "localhost:443",
					Organization: organization,
					NlxVersion:   testNlxVersion127,
					CreatedAt:    now,
					UpdatedAt:    now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					AnyTimes()

				mocks.repository.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost:443").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost:443",
			},
			wantResponse: &directoryapi.RegisterInwayResponse{},
			wantErr:      nil,
		},
		"cannot_compute_management_api_proxy_address_when_registering_an_organization_inway_nlx_version_0_127_0_and_below": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion127,
			)),
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
			wantErr: status.Error(codes.InvalidArgument, "cannot compute inway proxy address"),
		},
		"when_inway_address_is_invalid": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion127,
			)),
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost",
			},
			wantResponse: nil,
			wantErr:      status.Error(codes.InvalidArgument, "validation failed: Address: must be a valid dial string."),
		},
		"when_management_api_proxy_address_is_invalid": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion128,
			)),
			request: &directoryapi.RegisterInwayRequest{
				IsOrganizationInway:       true,
				InwayAddress:              "localhost:443",
				ManagementApiProxyAddress: "localhost",
			},
			wantResponse: nil,
			wantErr:      status.Error(codes.InvalidArgument, "management API proxy address must use port 443 or 8443"),
		},
		"when_inway_address_port_is_invalid": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion128,
			)),
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost:444",
			},
			wantResponse: nil,
			wantErr:      status.Error(codes.InvalidArgument, "inway address must use port 443 or 8443"),
		},
		"when_management_api_proxy_address_port_is_invalid": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion128,
			)),
			request: &directoryapi.RegisterInwayRequest{
				IsOrganizationInway:       true,
				InwayAddress:              "localhost:443",
				ManagementApiProxyAddress: "localhost:444",
			},
			wantResponse: nil,
			wantErr:      status.Error(codes.InvalidArgument, "management API proxy address must use port 443 or 8443"),
		},
		//nolint:dupl // mocks are the same but the request is different.
		"happy_flow_when_registering_an_organization_inway_nlx_version_0_127_0_and_below": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion127,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:                   "localhost:443",
					ManagementAPIProxyAddress: "localhost:444",
					Organization:              organization,
					NlxVersion:                testNlxVersion127,
					IsOrganizationInway:       true,
					CreatedAt:                 now,
					UpdatedAt:                 now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					Return(nil).
					AnyTimes()

				mocks.repository.
					EXPECT().
					RegisterService(gomock.Any()).
					Return(nil)

				mocks.repository.
					EXPECT().
					SetOrganizationInway(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress:        "localhost:443",
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
		//nolint:dupl // mocks are the same but the request is different.
		"happy_flow_when_registering_an_organization_inway_nlx_version_0_128_0-review": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion128Prerelease,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:                   "localhost:443",
					ManagementAPIProxyAddress: "localhost:8443",
					Organization:              organization,
					NlxVersion:                testNlxVersion128Prerelease,
					IsOrganizationInway:       true,
					CreatedAt:                 now,
					UpdatedAt:                 now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					Return(nil).
					AnyTimes()

				mocks.repository.
					EXPECT().
					RegisterService(gomock.Any()).
					Return(nil)

				mocks.repository.
					EXPECT().
					SetOrganizationInway(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress:              "localhost:443",
				ManagementApiProxyAddress: "localhost:8443",
				IsOrganizationInway:       true,
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
		//nolint:dupl // mocks are the same but the request is different.
		"happy_flow_when_registering_an_organization_inway": {
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion128,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:                   "localhost:443",
					ManagementAPIProxyAddress: "localhost:8443",
					Organization:              organization,
					NlxVersion:                testNlxVersion128,
					IsOrganizationInway:       true,
					CreatedAt:                 now,
					UpdatedAt:                 now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					Return(nil).
					AnyTimes()

				mocks.repository.
					EXPECT().
					RegisterService(gomock.Any()).
					Return(nil)

				mocks.repository.
					EXPECT().
					SetOrganizationInway(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress:              "localhost:443",
				ManagementApiProxyAddress: "localhost:8443",
				IsOrganizationInway:       true,
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
			context: metadata.NewIncomingContext(context.Background(), metadata.Pairs(
				"nlx-version", testNlxVersion128,
			)),
			setup: func(mocks serviceMocks) {
				organization, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
				assert.NoError(t, err)

				inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
					Address:      "localhost:443",
					Organization: organization,
					NlxVersion:   testNlxVersion128,
					CreatedAt:    now,
					UpdatedAt:    now,
				})
				assert.NoError(t, err)

				mocks.repository.
					EXPECT().
					RegisterInway(inwayModel).
					Return(nil).
					AnyTimes()

				mocks.repository.
					EXPECT().
					RegisterService(gomock.Any()).Return(nil)

				mocks.repository.EXPECT().
					ClearIfSetAsOrganizationInway(gomock.Any(), testOrganizationSerialNumber, "localhost:443").
					Return(nil)
			},
			request: &directoryapi.RegisterInwayRequest{
				InwayAddress: "localhost:443",
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
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: now,
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.RegisterInway(tt.context, tt.request)

			// log is added to be able to compare the error message of the status object
			if err != nil && tt.wantErr == nil {
				t.Log(err.Error())
			}

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantResponse, got)
		})
	}
}

type serviceMocks struct {
	repository *storage_mock.MockRepository
}

func testGetOrganizationInformationFromRequest(context.Context) (*tls.OrganizationInformation, error) {
	return &tls.OrganizationInformation{
		Name:         testOrganizationName,
		SerialNumber: testOrganizationSerialNumber,
	}, nil
}

func newService(t *testing.T, version, termsOfServiceURL string, clock directory.Clock) (*directory.DirectoryService, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		repository: storage_mock.NewMockRepository(ctrl),
	}

	service := directory.New(&directory.NewDirectoryArgs{
		Logger:                                zap.NewNop(),
		TermsOfServiceURL:                     termsOfServiceURL,
		Repository:                            mocks.repository,
		HTTPClient:                            nil,
		Clock:                                 clock,
		Version:                               version,
		GetOrganizationInformationFromRequest: testGetOrganizationInformationFromRequest,
	})

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
