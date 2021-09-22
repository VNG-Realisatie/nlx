// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/directory-registration-api/adapters"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func TestDirectoryRegistrationService_SetOrganizationInway(t *testing.T) {
	tests := map[string]struct {
		setup            func(serviceMocks)
		address          string
		expectedResponse *emptypb.Empty
		expectedError    error
	}{
		"empty_address": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					SetOrganizationInway(gomock.Any(), nil, nil).
					Times(0)
			},
			address:          "",
			expectedResponse: nil,
			expectedError:    status.New(codes.InvalidArgument, "address is empty").Err(),
		},
		"no_inway_with_address": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					SetOrganizationInway(gomock.Any(), "Test Organization Name", "inway.nlx.local").
					Return(adapters.ErrNoInwayWithAddress)
			},
			address:          "inway.nlx.local",
			expectedResponse: nil,
			expectedError:    status.New(codes.NotFound, "inway with address not found").Err(),
		},
		"database_error": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					SetOrganizationInway(gomock.Any(), "Test Organization Name", "inway.nlx.local").
					Return(errors.New("arbitrary error"))
			},
			address:          "inway.nlx.local",
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					SetOrganizationInway(gomock.Any(), "Test Organization Name", "inway.nlx.local").
					Return(nil)
			},
			address:          "inway.nlx.local",
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

			got, err := service.SetOrganizationInway(context.Background(), &registrationapi.SetOrganizationInwayRequest{Address: tt.address})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

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
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
					Return(errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		"when_the_organization_is_not_present": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
					Return(adapters.ErrOrganizationNotFound)
			},
			expectedResponse: &emptypb.Empty{},
			expectedError:    nil,
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.r.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
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
