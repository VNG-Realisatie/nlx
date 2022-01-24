// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

//nolint:funlen // this is a test
func TestListServices(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *directoryapi.ListServicesResponse
		expectedError    error
	}{
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					ListServices(gomock.Any(), "01234567890123456789").
					Return(nil, errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Database error.").Err(),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				organizationA, _ := domain.NewOrganization("org-a", "00000000000000000001")
				organizationB, _ := domain.NewOrganization("org-b", "00000000000000000002")

				serviceA, _ := domain.NewService(&domain.NewServiceArgs{
					Name:                 "service-a",
					Organization:         organizationA,
					Internal:             false,
					DocumentationURL:     "http://docs.nlx.io",
					APISpecificationType: domain.OpenAPI2,
					PublicSupportContact: "public@support.com",
					TechSupportContact:   "tech@support.com",
					Costs: &domain.NewServiceCostsArgs{
						OneTime: 1,
						Monthly: 2,
						Request: 3,
					},
					Inways: []*domain.NewServiceInwayArgs{
						{
							Address: "inway-address-a",
							State:   domain.InwayUP,
						},
					},
				})

				serviceB, _ := domain.NewService(&domain.NewServiceArgs{
					Name:                 "service-b",
					Organization:         organizationB,
					Internal:             false,
					DocumentationURL:     "http://docs.nlx.io",
					APISpecificationType: domain.OpenAPI3,
					PublicSupportContact: "public@support.com",
					TechSupportContact:   "tech@support.com",
					Costs: &domain.NewServiceCostsArgs{
						OneTime: 1,
						Monthly: 2,
						Request: 3,
					},
					Inways: []*domain.NewServiceInwayArgs{
						{
							Address: "inway-address-b",
							State:   domain.InwayDOWN,
						},
						{
							Address: "inway-address-c",
							State:   domain.InwayUP,
						},
					},
				})

				mocks.r.
					EXPECT().
					ListServices(gomock.Any(), "01234567890123456789").
					Return([]*domain.Service{serviceA, serviceB}, nil)
			},
			expectedResponse: &directoryapi.ListServicesResponse{
				Services: []*directoryapi.ListServicesResponse_Service{
					{
						Name:                 "service-a",
						Internal:             false,
						DocumentationUrl:     "http://docs.nlx.io",
						ApiSpecificationType: "OpenAPI2",
						PublicSupportContact: "public@support.com",
						Organization: &directoryapi.Organization{
							Name:         "org-a",
							SerialNumber: "00000000000000000001",
						},
						Inways: []*directoryapi.Inway{
							{
								Address: "inway-address-a",
								State:   directoryapi.Inway_UP,
							},
						},
						Costs: &directoryapi.ListServicesResponse_Costs{
							OneTime: 1,
							Monthly: 2,
							Request: 3,
						},
					},
					{
						Name:                 "service-b",
						Internal:             false,
						DocumentationUrl:     "http://docs.nlx.io",
						ApiSpecificationType: "OpenAPI3",
						PublicSupportContact: "public@support.com",
						Inways: []*directoryapi.Inway{
							{
								Address: "inway-address-b",
								State:   directoryapi.Inway_DOWN,
							},
							{
								Address: "inway-address-c",
								State:   directoryapi.Inway_UP,
							},
						},
						Organization: &directoryapi.Organization{
							Name:         "org-b",
							SerialNumber: "00000000000000000002",
						},
						Costs: &directoryapi.ListServicesResponse_Costs{
							OneTime: 1,
							Monthly: 2,
							Request: 3,
						},
					},
				},
			},
			expectedError: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, "")

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.ListServices(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
