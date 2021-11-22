// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

type args struct {
	peer     *peer.Peer
	database *database.Inway
	request  *api.Inway
}

var createInwayTests = []struct {
	name string
	args args
	want *api.Inway
}{
	{
		name: "ip address from context",
		args: args{
			database: &database.Inway{Name: "inway42.basic", IPAddress: "127.1.1.1"},
			request:  &api.Inway{Name: "inway42.basic"},
			peer:     &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}},
		},
		want: &api.Inway{
			Name:      "inway42.basic",
			IpAddress: "127.1.1.1",
		},
	},
	{
		name: "ip address from request is ignored",
		args: args{
			database: &database.Inway{Name: "inway42.ignore-ip", IPAddress: "127.1.1.1"},
			request:  &api.Inway{Name: "inway42.ignore-ip", IpAddress: "127.2.2.2"},
			peer:     &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}},
		},
		want: &api.Inway{
			Name:      "inway42.ignore-ip",
			IpAddress: "127.1.1.1",
		},
	},
	{
		name: "the connection context must contain an address",
		args: args{
			database: &database.Inway{Name: "inway42.ip-context-required"},
			request:  &api.Inway{Name: "inway42.ip-context-required"},
			peer:     &peer.Peer{Addr: nil},
		},
	},
	{
		name: "ipv6",
		args: args{
			database: &database.Inway{Name: "inway42.ipv6", IPAddress: "::1"},
			request:  &api.Inway{Name: "inway42.ipv6"},
			peer:     &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}},
		},
		want: &api.Inway{
			Name:      "inway42.ipv6",
			IpAddress: "::1",
		},
	},
}

func TestRegisterInway(t *testing.T) {
	for _, tt := range createInwayTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			logger := zaptest.Logger(t)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := peer.NewContext(context.Background(), tt.args.peer)
			mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
			if tt.want != nil {
				mockDatabase.EXPECT().RegisterInway(ctx, tt.args.database)
			}
			service := server.NewManagementService(
				logger,
				mock_directory.NewMockClient(mockCtrl),
				nil,
				mockDatabase,
				nil,
				mock_auditlog.NewMockLogger(mockCtrl),
				management.NewClient,
			)

			response, err := service.RegisterInway(ctx, tt.args.request)
			if tt.want != nil {
				assert.NoError(t, err, "could not create inway")
				assert.Equal(t, tt.want, response)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetInway(t *testing.T) {
	logger := zap.NewNop()
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(
		logger,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		mockDatabase,
		nil,
		mock_auditlog.NewMockLogger(mockCtrl),
		management.NewClient,
	)

	getInwayRequest := &api.GetInwayRequest{
		Name: "inway42.test",
	}

	mockDatabase.EXPECT().GetInway(ctx, "inway42.test").Return(nil, database.ErrNotFound)

	_, actualError := service.GetInway(ctx, getInwayRequest)
	expectedError := status.Error(codes.NotFound, "inway not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	mockInwayResponse := &database.Inway{
		Name:      "inway42.test",
		IPAddress: "",
		Services: []*database.Service{{
			Name: "forty-two",
		}},
	}

	mockDatabase.EXPECT().GetInway(ctx, "inway42.test").Return(mockInwayResponse, nil)

	getInwayResponse, err := service.GetInway(ctx, getInwayRequest)
	assert.NoError(t, err)

	expectedResponse := &api.Inway{
		Name:     "inway42.test",
		Services: []*api.Inway_Service{{Name: "forty-two"}},
	}

	assert.Equal(t, expectedResponse, getInwayResponse)
}

func TestUpdateInway(t *testing.T) {
	logger := zap.NewNop()
	ctx := context.Background()

	mockInway := &database.Inway{
		ID:   1,
		Name: "inway42.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().UpdateInway(ctx, mockInway)
	mockDatabase.EXPECT().GetInway(ctx, "inway42.test").Return(mockInway, nil)

	service := server.NewManagementService(
		logger,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		mockDatabase,
		nil,
		mock_auditlog.NewMockLogger(mockCtrl),
		management.NewClient,
	)

	updateInwayRequest := &api.UpdateInwayRequest{
		Name: "inway42.test",
		Inway: &api.Inway{
			Name: "inway42.test",
		},
	}

	updateInwayResponse, err := service.UpdateInway(ctx, updateInwayRequest)
	assert.NoError(t, err)

	expectedResponse := &api.Inway{
		Name: "inway42.test",
	}

	assert.Equal(t, expectedResponse, updateInwayResponse)
}

func TestDeleteInway(t *testing.T) {
	tests := map[string]struct {
		request       *api.DeleteInwayRequest
		ctx           context.Context
		setup         func(*common_tls.CertificateBundle, serviceMocks)
		expectedError error
	}{
		"failed_to_retrieve_user_info_from_context": {
			ctx: context.Background(),
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
			expectedError: status.Error(codes.Internal, "could not retrieve user info to create audit log"),
		},
		"failed_to_create_audit_log": {
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.al.EXPECT().
					InwayDelete(gomock.Any(), "Jane Doe", "nlxctl", "my-inway").
					Return(fmt.Errorf("error"))
			},
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
			expectedError: status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"failed_to_delete_inway_from_database": {
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.EXPECT().
					DeleteInway(gomock.Any(), "my-inway").
					Return(fmt.Errorf("error"))

				mocks.al.EXPECT().
					InwayDelete(gomock.Any(), "Jane Doe", "nlxctl", "my-inway")
			},
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
			expectedError: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					DeleteInway(gomock.Any(), "my-inway")

				mocks.al.
					EXPECT().
					InwayDelete(gomock.Any(), "Jane Doe", "nlxctl", "my-inway")
			},
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, bundle, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(bundle, mocks)
			}

			_, err := service.DeleteInway(tt.ctx, tt.request)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestListInways(t *testing.T) {
	logger := zap.NewNop()
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)

	mockListInways := []*database.Inway{
		{Name: "inway42.test"},
		{Name: "inway43.test"},
		{
			Name:        "inway.test",
			Version:     "1.0.0",
			Hostname:    "inway.test.local",
			SelfAddress: "inway.nlx",
			Services: []*database.Service{
				{
					Name: "mock-service",
					Inways: []*database.Inway{
						{
							Name: "inway.test",
						},
					},
				},
			},
		},
	}

	mockDatabase.EXPECT().ListInways(ctx).Return(mockListInways, nil)

	service := server.NewManagementService(
		logger,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		mockDatabase,
		nil,
		mock_auditlog.NewMockLogger(mockCtrl),
		management.NewClient,
	)
	actualResponse, err := service.ListInways(ctx, &api.ListInwaysRequest{})
	assert.NoError(t, err)

	expectedResponse := &api.ListInwaysResponse{
		Inways: []*api.Inway{
			{
				Name: "inway42.test",
			},
			{
				Name: "inway43.test",
			},
			{
				Name:        "inway.test",
				Version:     "1.0.0",
				Hostname:    "inway.test.local",
				SelfAddress: "inway.nlx",
				Services: []*api.Inway_Service{
					{
						Name: "mock-service",
					},
				},
			},
		},
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
