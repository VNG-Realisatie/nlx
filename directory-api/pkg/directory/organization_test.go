// Copyright © VNG Realisatie 2020
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
	storage "go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func TestDirectoryRegistrationService_ClearOrganizationInway(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *emptypb.Empty
		expectedError    error
	}{
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), testOrganizationSerialNumber).
					Return(errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		"when_the_organization_is_not_present": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), testOrganizationSerialNumber).
					Return(storage.ErrOrganizationNotFound)
			},
			expectedResponse: &emptypb.Empty{},
			expectedError:    nil,
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), testOrganizationSerialNumber).
					Return(nil)
			},
			expectedResponse: &emptypb.Empty{},
			expectedError:    nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.ClearOrganizationInway(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestListOrganizations(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *directoryapi.ListOrganizationsResponse
		expectedError    error
	}{
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					ListOrganizations(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "Database error.").Err(),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				organizationA, _ := domain.NewOrganization("org-a", "00000000000000000001")
				organizationB, _ := domain.NewOrganization("org-b", "00000000000000000002")

				mocks.r.
					EXPECT().
					ListOrganizations(gomock.Any()).
					Return([]*domain.Organization{organizationA, organizationB}, nil)
			},
			expectedResponse: &directoryapi.ListOrganizationsResponse{
				Organizations: []*directoryapi.Organization{
					{
						Name:         "org-a",
						SerialNumber: "00000000000000000001",
					},
					{
						Name:         "org-b",
						SerialNumber: "00000000000000000002",
					},
				},
			},
			expectedError: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.ListOrganizations(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
