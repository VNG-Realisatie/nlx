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
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
	"go.nlx.io/nlx/directory-api/pkg/directory"
)

func TestClearOrganizationInway(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *directoryapi.ClearOrganizationInwayResponse
		expectedError    error
	}{
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), testOrganizationSerialNumber).
					Return(errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, directory.ErrMessageDatabase),
		},
		"when_the_organization_is_not_present": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), testOrganizationSerialNumber).
					Return(storage.ErrNotFound)
			},
			expectedResponse: &directoryapi.ClearOrganizationInwayResponse{},
			expectedError:    nil,
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), testOrganizationSerialNumber).
					Return(nil)
			},
			expectedResponse: &directoryapi.ClearOrganizationInwayResponse{},
			expectedError:    nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: time.Now(),
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.ClearOrganizationInway(context.Background(), &directoryapi.ClearOrganizationInwayRequest{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestSetOrganizationEmailAddress(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		req              *directoryapi.SetOrganizationContactDetailsRequest
		expectedResponse *directoryapi.SetOrganizationContactDetailsResponse
		expectedError    error
	}{
		"invalid_email_address": {
			setup: func(mocks serviceMocks) {
			},
			req: &directoryapi.SetOrganizationContactDetailsRequest{
				EmailAddress: "invalidemail",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.InvalidArgument, "email_address: must be a valid email address."),
		},
		"database_error": {
			setup: func(mocks serviceMocks) {
				org, err := domain.NewOrganization("Test Organization Name", "01234567890123456789")
				require.NoError(t, err)

				mocks.repository.
					EXPECT().
					SetOrganizationEmailAddress(gomock.Any(), org, "mock@email.com").
					Return(errors.New("arbitrary error"))
			},
			req: &directoryapi.SetOrganizationContactDetailsRequest{
				EmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, directory.ErrMessageDatabase),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				org, err := domain.NewOrganization("Test Organization Name", "01234567890123456789")
				require.NoError(t, err)

				mocks.repository.
					EXPECT().
					SetOrganizationEmailAddress(gomock.Any(), org, "mock@email.com").
					Return(nil)
			},
			req: &directoryapi.SetOrganizationContactDetailsRequest{
				EmailAddress: "mock@email.com",
			},
			expectedResponse: &directoryapi.SetOrganizationContactDetailsResponse{},
			expectedError:    nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: time.Now(),
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.SetOrganizationContactDetails(context.Background(), tt.req)

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
				mocks.repository.
					EXPECT().
					ListOrganizations(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, directory.ErrMessageDatabase),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				organizationA, _ := domain.NewOrganization("org-a", "00000000000000000001")
				organizationB, _ := domain.NewOrganization("org-b", "00000000000000000002")

				mocks.repository.
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
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: time.Now(),
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.ListOrganizations(context.Background(), &directoryapi.ListOrganizationsRequest{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestGetOrganizationManagementAPIProxyAddress(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *directoryapi.GetOrganizationManagementAPIProxyAddressResponse
		expectedError    error
	}{
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					GetOrganizationInwayManagementAPIProxyAddress(gomock.Any(), "00000000000000000001").
					Return("", fmt.Errorf("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, directory.ErrMessageDatabase).Err(),
		},
		"organization_inway_not_set": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					GetOrganizationInwayManagementAPIProxyAddress(gomock.Any(), "00000000000000000001").
					Return("", storage.ErrNotFound)
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.NotFound, directory.ErrMessageOrganizationInwayNotFound),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					GetOrganizationInwayManagementAPIProxyAddress(gomock.Any(), "00000000000000000001").
					Return("mockaddress.nl:443", nil)
			},
			expectedResponse: &directoryapi.GetOrganizationManagementAPIProxyAddressResponse{Address: "mockaddress.nl:443"},
			expectedError:    nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: time.Now(),
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.GetOrganizationManagementAPIProxyAddress(context.Background(), &directoryapi.GetOrganizationManagementAPIProxyAddressRequest{
				OrganizationSerialNumber: "00000000000000000001",
			})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestGetOrganizationInway(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		expectedResponse *directoryapi.GetOrganizationInwayResponse
		expectedError    error
	}{
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					GetOrganizationInwayAddress(gomock.Any(), "00000000000000000001").
					Return("", fmt.Errorf("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, directory.ErrMessageDatabase),
		},
		"organization_inway_not_set": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					GetOrganizationInwayAddress(gomock.Any(), "00000000000000000001").
					Return("", storage.ErrNotFound)
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.NotFound, directory.ErrMessageOrganizationInwayNotFound),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.repository.
					EXPECT().
					GetOrganizationInwayAddress(gomock.Any(), "00000000000000000001").
					Return("mockaddress.nl:443", nil)
			},
			expectedResponse: &directoryapi.GetOrganizationInwayResponse{Address: "mockaddress.nl:443"},
			expectedError:    nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newService(t, testNlxVersion128, "", &testClock{
				timeToReturn: time.Now(),
			})

			if tt.setup != nil {
				tt.setup(mocks)
			}

			got, err := service.GetOrganizationInway(context.Background(), &directoryapi.GetOrganizationInwayRequest{
				OrganizationSerialNumber: "00000000000000000001",
			})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
