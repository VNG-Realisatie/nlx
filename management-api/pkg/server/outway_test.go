// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"net"
	"path/filepath"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

func mockIP(t *testing.T, ip string) pgtype.Inet {
	_, ipNet, err := net.ParseCIDR(ip)
	require.NoError(t, err)

	return pgtype.Inet{
		Status: pgtype.Present,
		IPNet:  ipNet,
	}
}

//nolint:funlen // this is a test function
func TestRegisterOutway(t *testing.T) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	testPublicKeyPEM, err := certBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		peer     *peer.Peer
		database *database.Outway
		request  *api.RegisterOutwayRequest
	}

	var tests = map[string]struct {
		setup   func(serviceMocks)
		args    args
		wantErr error
	}{
		"when_the_connection_context_does_not_contain_an_address": {
			args: args{
				database: &database.Outway{
					Name:         "outway42.ip-context-required",
					PublicKeyPEM: testPublicKeyPEM,
					Version:      "unknown",
				},
				request: &api.RegisterOutwayRequest{
					Name:         "outway42.ip-context-required",
					PublicKeyPEM: testPublicKeyPEM,
					Version:      "unknown",
				},
				peer: &peer.Peer{Addr: nil},
			},
			wantErr: status.Error(codes.Internal, "peer addr is invalid"),
		},
		"when_providing_an_invalid_outway_name": {
			args: args{
				request: &api.RegisterOutwayRequest{
					Name:         "",
					PublicKeyPEM: testPublicKeyPEM,
					Version:      "unknown",
				},
				peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}},
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid outway: name: cannot be blank."),
		},
		"happy_flow_ipv4": {
			args: args{
				database: &database.Outway{
					Name:                 "outway42.basic",
					PublicKeyPEM:         testPublicKeyPEM,
					IPAddress:            mockIP(t, "127.1.1.1/32"),
					Version:              "unknown",
					PublicKeyFingerprint: "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
				},
				request: &api.RegisterOutwayRequest{
					Name:         "outway42.basic",
					PublicKeyPEM: testPublicKeyPEM,
					Version:      "unknown",
				},
				peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}},
			},
		},
		"happy_flow_ipv6": {
			args: args{
				database: &database.Outway{
					Name:                 "outway42.ipv6",
					IPAddress:            mockIP(t, "::1/32"),
					PublicKeyPEM:         testPublicKeyPEM,
					PublicKeyFingerprint: "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
					Version:              "unknown",
				},
				request: &api.RegisterOutwayRequest{
					Name:         "outway42.ipv6",
					PublicKeyPEM: testPublicKeyPEM,
					Version:      "unknown",
				},
				peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			logger := zaptest.Logger(t)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := peer.NewContext(context.Background(), tt.args.peer)
			mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)

			if tt.wantErr == nil {
				mockDatabase.EXPECT().RegisterOutway(ctx, tt.args.database)
			}

			service := server.NewManagementService(
				logger,
				mock_directory.NewMockClient(mockCtrl),
				nil,
				nil,
				mockDatabase,
				nil,
				mock_auditlog.NewMockLogger(mockCtrl),
				management.NewClient,
			)

			_, err := service.RegisterOutway(ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestListOutways(t *testing.T) {
	logger := zap.NewNop()
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)

	mockListOutways := []*database.Outway{
		{Name: "outway42.test"},
		{Name: "outway43.test"},
		{
			Name:                 "outway.test",
			Version:              "1.0.0",
			IPAddress:            mockIP(t, "127.1.1.1/32"),
			PublicKeyPEM:         "mock-public-key-pem",
			PublicKeyFingerprint: "mock-public-key-fingerprint",
		},
	}

	mockDatabase.EXPECT().ListOutways(ctx).Return(mockListOutways, nil)

	service := server.NewManagementService(
		logger,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		nil,
		mockDatabase,
		nil,
		mock_auditlog.NewMockLogger(mockCtrl),
		management.NewClient,
	)
	actualResponse, err := service.ListOutways(ctx, &api.ListOutwaysRequest{})
	assert.NoError(t, err)

	expectedResponse := &api.ListOutwaysResponse{
		Outways: []*api.Outway{
			{
				Name: "outway42.test",
			},
			{
				Name: "outway43.test",
			},
			{
				Name:                 "outway.test",
				IpAddress:            "127.1.1.1",
				Version:              "1.0.0",
				PublicKeyPEM:         "mock-public-key-pem",
				PublicKeyFingerprint: "mock-public-key-fingerprint",
			},
		},
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
