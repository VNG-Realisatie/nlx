// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gogo/protobuf/types"
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

//nolint:funlen // adding the tests was the first step to make the functionality testable. making it less complex is out of scope for now.
func TestDirectoryRegistrationService_SetInsightConfiguration(t *testing.T) {
	type fields struct {
		logger                         *zap.Logger
		db                             func(ctrl *gomock.Controller) database.DirectoryDatabase
		getOrganisationNameFromRequest func(ctx context.Context) (string, error)
	}

	type args struct {
		ctx context.Context
		req *registrationapi.SetInsightConfigurationRequest
	}

	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse *types.Empty
		expectedError    error
	}{
		{
			name: "with an invalid organization name in the request",
			fields: fields{
				logger: zap.NewNop(),
				db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
					return mock.NewMockDirectoryDatabase(ctrl)
				},
				getOrganisationNameFromRequest: func(ctx context.Context) (string, error) {
					return testInvalidOrganizationName, nil
				},
			},
			args: args{
				ctx: context.Background(),
				req: &registrationapi.SetInsightConfigurationRequest{
					InsightAPIURL: "https://insight-api.url",
					IrmaServerURL: "https://irma-server-url",
				},
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.InvalidArgument, "Invalid organization name").Err(),
		},
		{
			name: "failed to communicate with the database",
			fields: fields{
				logger: zap.NewNop(),
				db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
					db := mock.NewMockDirectoryDatabase(ctrl)
					db.EXPECT().SetInsightConfiguration(
						gomock.Any(),
						testOrganizationName,
						"https://insight-api.url",
						"https://irma-server-url",
					).Return(errors.New("arbitrary  error")).AnyTimes()

					return db
				},
				getOrganisationNameFromRequest: testGetOrganizationNameFromRequest,
			},
			args: args{
				ctx: context.Background(),
				req: &registrationapi.SetInsightConfigurationRequest{
					InsightAPIURL: "https://insight-api.url",
					IrmaServerURL: "https://irma-server-url",
				},
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		{
			name: "happy flow",
			fields: fields{
				logger: zap.NewNop(),
				db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
					db := mock.NewMockDirectoryDatabase(ctrl)
					db.EXPECT().SetInsightConfiguration(
						gomock.Any(),
						testOrganizationName,
						"https://insight-api.url",
						"https://irma-server-url",
					)

					return db
				},
				getOrganisationNameFromRequest: testGetOrganizationNameFromRequest,
			},
			args: args{
				ctx: context.Background(),
				req: &registrationapi.SetInsightConfigurationRequest{
					InsightAPIURL: "https://insight-api.url",
					IrmaServerURL: "https://irma-server-url",
				},
			},
			expectedResponse: &types.Empty{},
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := registrationservice.New(tt.fields.logger, tt.fields.db(ctrl), nil, tt.fields.getOrganisationNameFromRequest)
			got, err := h.SetInsightConfiguration(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestDirectoryRegistrationService_SetOrganizationInway(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		address          string
		expectedResponse *types.Empty
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
			expectedResponse: &types.Empty{},
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
		db               func(ctrl *gomock.Controller) database.DirectoryDatabase
		address          string
		expectedResponse *types.Empty
		expectedError    error
	}{
		{
			name: "no_organization",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
					Return(database.ErrNoOrganization)

				return db
			},
			address:          "inway.nlx.local",
			expectedResponse: nil,
			expectedError:    status.New(codes.NotFound, "organization not found").Err(),
		},
		{
			name: "database_error",
			db: func(ctrl *gomock.Controller) database.DirectoryDatabase {
				db := mock.NewMockDirectoryDatabase(ctrl)
				db.EXPECT().
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
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
					ClearOrganizationInway(gomock.Any(), "Test Organization Name").
					Return(nil)

				return db
			},
			address:          "inway.nlx.local",
			expectedResponse: &types.Empty{},
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := registrationservice.New(zap.NewNop(), tt.db(ctrl), nil, testGetOrganizationNameFromRequest)

			got, err := s.ClearOrganizationInway(context.Background(), &types.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
