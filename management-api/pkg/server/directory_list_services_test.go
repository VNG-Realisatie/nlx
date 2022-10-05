// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/environment"
	"go.nlx.io/nlx/management-api/pkg/server"
)

// nolint dupl: this is a test
func TestListDirectoryServices(t *testing.T) {
	logger := zap.NewNop()
	env := &environment.Environment{}
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	clientServices := []*directoryapi.ListServicesResponse_Service{
		{
			Name: "test-service-1",
			Organization: &directoryapi.Organization{
				SerialNumber: "00000000000000000001",
				Name:         "Organization One",
			},
			ApiSpecificationType: "OpenAPI3",
			DocumentationUrl:     "https://example.com",
			PublicSupportContact: "test@example.com",
			Costs: &directoryapi.ListServicesResponse_Costs{
				OneTime: 1,
				Monthly: 5,
				Request: 250,
			},
		},
	}

	client := mock_directory.NewMockClient(mockCtrl)
	client.EXPECT().
		ListServices(ctx, &directoryapi.ListServicesRequest{}).
		Return(&directoryapi.ListServicesResponse{Services: clientServices}, nil)

	client.
		EXPECT().
		ListOrganizations(ctx, &directoryapi.ListOrganizationsRequest{}).
		Return(&directoryapi.ListOrganizationsResponse{
			Organizations: []*directoryapi.Organization{
				{
					SerialNumber: "00000000000000000001",
					Name:         "Organization One",
				},
			},
		}, nil)

	db := mock_database.NewMockConfigDatabase(mockCtrl)
	service := clientServices[0]

	db.EXPECT().
		ListLatestOutgoingAccessRequests(ctx, service.Organization.SerialNumber, service.Name).
		Return([]*database.OutgoingAccessRequest{
			{
				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
				},
				ServiceName:          "test-service-1",
				State:                database.OutgoingAccessRequestReceived,
				PublicKeyFingerprint: "public-key-fingerprint",
				CreatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
				UpdatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			},
		}, nil)

	db.EXPECT().
		GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
		Return(&database.AccessProof{
			ID:                      1,
			AccessRequestOutgoingID: 1,
			OutgoingAccessRequest: &database.OutgoingAccessRequest{
				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
				},
				ServiceName:          "test-service-1",
				State:                database.OutgoingAccessRequestReceived,
				PublicKeyFingerprint: "public-key-fingerprint",
				CreatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
				UpdatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			},

			CreatedAt: time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		}, nil)

	directoryService := server.NewDirectoryService(
		logger,
		env,
		client,
		db,
	)
	response, err := directoryService.ListServices(ctx, &emptypb.Empty{})
	assert.NoError(t, err)

	expected := []*api.DirectoryService{
		{
			ServiceName: "test-service-1",
			Organization: &api.Organization{
				SerialNumber: "00000000000000000001",
				Name:         "Organization One",
			},
			ApiSpecificationType: "OpenAPI3",
			DocumentationUrl:     "https://example.com",
			PublicSupportContact: "test@example.com",
			State:                api.DirectoryService_STATE_UNSPECIFIED,
			OneTimeCosts:         1,
			MonthlyCosts:         5,
			RequestCosts:         250,
			//nolint dupl: this is a test
			AccessStates: []*api.DirectoryService_AccessState{
				{
					AccessRequest: &api.OutgoingAccessRequest{
						Id: 1,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "Organization One",
						},
						ServiceName:          "test-service-1",
						PublicKeyFingerprint: "public-key-fingerprint",
						State:                api.AccessRequestState_ACCESS_REQUEST_STATE_RECEIVED,
						CreatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
						UpdatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
					},
					AccessProof: &api.AccessProof{
						Id: 1,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "Organization One",
						},
						ServiceName:          "test-service-1",
						AccessRequestId:      1,
						PublicKeyFingerprint: "public-key-fingerprint",
						CreatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
					},
				},
			},
		},
	}

	assert.EqualValues(t, expected, response.Services)
}
