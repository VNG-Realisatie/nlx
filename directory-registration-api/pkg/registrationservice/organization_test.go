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
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-registration-api/pkg/registrationservice"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func TestDirectoryRegistrationService_SetOrganizationInway(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		address          string
		expectedResponse *emptypb.Empty
		expectedError    error
	}{
		{
			name: "empty_address",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					SetOrganizationInway(gomock.Any(), nil, nil).
					Times(0)

				return db
			},
			address:          "",
			expectedResponse: nil,
			expectedError:    status.New(codes.InvalidArgument, "address is empty").Err(),
		},
		{
			name: "no_inway_with_address",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					SetOrganizationInway(gomock.Any(), "Test Organization Name", "inway.nlx.local").
					Return(database.ErrNoInwayWithAddress)

				return db
			},
			address:          "inway.nlx.local",
			expectedResponse: nil,
			expectedError:    status.New(codes.NotFound, "inway with address not found").Err(),
		},
		{
			name: "database_error",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					SetOrganizationInway(gomock.Any(), "Test Organization Name", "inway.nlx.local").
					Return(errors.New("arbitrary error"))

				return db
			},
			address:          "inway.nlx.local",
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		{
			name: "happy_flow",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					SetOrganizationInway(gomock.Any(), "Test Organization Name", "inway.nlx.local").
					Return(nil)

				return db
			},
			address:          "inway.nlx.local",
			expectedResponse: &emptypb.Empty{},
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := registrationservice.New(zap.NewNop(), tt.db(ctrl), nil, testGetOrganizationNameFromRequest)

			got, err := s.SetOrganizationInway(context.Background(), &registrationapi.SetOrganizationInwayRequest{Address: tt.address})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestDirectoryRegistrationService_ClearOrganizationInway(t *testing.T) {
	tests := []struct {
		name             string
		setup            func(serviceMocks)
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		address          string
		expectedResponse *emptypb.Empty
		expectedError    error
	}{
		{
			name: "no_organization",
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
					Return(database.ErrOrganizationNotFound)
			},
			address:          "inway.nlx.local",
			expectedResponse: nil,
			expectedError:    status.New(codes.NotFound, "organization not found").Err(),
		},
		{
			name: "database_error",
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
					Return(errors.New("arbitrary error"))
			},
			address:          "inway.nlx.local",
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		{
			name: "happy_flow",
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
					Return(nil)
			},
			address:          "inway.nlx.local",
			expectedResponse: &emptypb.Empty{},
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

			got, err := service.ClearOrganizationInway(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
