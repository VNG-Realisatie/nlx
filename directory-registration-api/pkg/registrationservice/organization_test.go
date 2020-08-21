// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice_test

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-registration-api/pkg/registrationservice"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func TestDirectoryRegistrationService_SetInsightConfiguration(t *testing.T) {
	type fields struct {
		logger                         *zap.Logger
		db                             database.DirectoryDatabase
		httpClient                     *http.Client
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
		expectedResponse *registrationapi.Empty
		expectedError    error
	}{
		{
			name: "with an invalid organization name in the request",
			fields: fields{
				logger: zap.NewNop(),
				db: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					return db
				}(),
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
				db: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().SetInsightConfiguration(
						gomock.Any(),
						testOrganizationName,
						"https://insight-api.url",
						"https://irma-server-url",
					).Return(errors.New("arbitrary  error")).AnyTimes()

					return db
				}(),
				getOrganisationNameFromRequest: func(ctx context.Context) (string, error) {
					return testOrganizationName, nil
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
			expectedError:    status.New(codes.Internal, "database error").Err(),
		},
		{
			name: "happy flow",
			fields: fields{
				logger: zap.NewNop(),
				db: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().SetInsightConfiguration(
						gomock.Any(),
						testOrganizationName,
						"https://insight-api.url",
						"https://irma-server-url",
					)

					return db
				}(),
				getOrganisationNameFromRequest: func(ctx context.Context) (string, error) {
					return testOrganizationName, nil
				},
			},
			args: args{
				ctx: context.Background(),
				req: &registrationapi.SetInsightConfigurationRequest{
					InsightAPIURL: "https://insight-api.url",
					IrmaServerURL: "https://irma-server-url",
				},
			},
			expectedResponse: &registrationapi.Empty{},
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			h := registrationservice.New(tt.fields.logger, tt.fields.db, tt.fields.httpClient, tt.fields.getOrganisationNameFromRequest)
			got, err := h.SetInsightConfiguration(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
