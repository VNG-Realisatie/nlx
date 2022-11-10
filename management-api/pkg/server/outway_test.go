// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
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
					Name:                 "outway42.ip-context-required",
					PublicKeyPEM:         testPublicKeyPEM,
					PublicKeyFingerprint: certBundle.PublicKeyFingerprint(),
					SelfAddressAPI:       "self-address",
					Version:              "unknown",
				},
				request: &api.RegisterOutwayRequest{
					Name:         "outway42.ip-context-required",
					PublicKeyPem: testPublicKeyPEM,
					Version:      "unknown",
				},
				peer: &peer.Peer{Addr: nil},
			},
			wantErr: status.Error(codes.Internal, "peer addr is invalid"),
		},
		"when_providing_an_invalid_outway_name": {
			args: args{
				request: &api.RegisterOutwayRequest{
					Name:           "",
					PublicKeyPem:   testPublicKeyPEM,
					SelfAddressApi: "self-address",
					Version:        "unknown",
				},
				peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}},
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid outway: name: cannot be blank."),
		},
		"when_providing_an_empty_self_address": {
			args: args{
				request: &api.RegisterOutwayRequest{
					Name:           "outway42.basic",
					PublicKeyPem:   testPublicKeyPEM,
					SelfAddressApi: "",
					Version:        "unknown",
				},
				peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}},
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid outway: self_address_api: cannot be blank."),
		},
		"happy_flow_ipv4": {
			args: args{
				database: &database.Outway{
					Name:                 "outway42.basic",
					PublicKeyPEM:         testPublicKeyPEM,
					IPAddress:            mockIP(t, "127.1.1.1/32"),
					SelfAddressAPI:       "self-address",
					Version:              "unknown",
					PublicKeyFingerprint: "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
				},
				request: &api.RegisterOutwayRequest{
					Name:           "outway42.basic",
					PublicKeyPem:   testPublicKeyPEM,
					SelfAddressApi: "self-address",
					Version:        "unknown",
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
					SelfAddressAPI:       "self-address",
					Version:              "unknown",
				},
				request: &api.RegisterOutwayRequest{
					Name:           "outway42.ipv6",
					PublicKeyPem:   testPublicKeyPEM,
					SelfAddressApi: "self-address",
					Version:        "unknown",
				},
				peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t, nil)

			ctx := peer.NewContext(context.Background(), tt.args.peer)

			if tt.wantErr == nil {
				mocks.db.
					EXPECT().
					RegisterOutway(ctx, tt.args.database)
			}

			_, err := service.RegisterOutway(ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestListOutways(t *testing.T) {
	tests := map[string]struct {
		ctx     context.Context
		setup   func(*testing.T, serviceMocks)
		want    *api.ListOutwaysResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(t *testing.T, mocks serviceMocks) {},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.outways.read\" to execute this request").Err(),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, mocks serviceMocks) {
				mocks.db.EXPECT().ListOutways(gomock.Any()).Return([]*database.Outway{
					{Name: "outway42.test"},
					{Name: "outway43.test"},
					{
						Name:                 "outway.test",
						Version:              "1.0.0",
						IPAddress:            mockIP(t, "127.1.1.1/32"),
						PublicKeyPEM:         "mock-public-key-pem",
						PublicKeyFingerprint: "mock-public-key-fingerprint",
						SelfAddressAPI:       "self-address",
					},
				}, nil)
			},
			want: &api.ListOutwaysResponse{
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
						PublicKeyPem:         "mock-public-key-pem",
						PublicKeyFingerprint: "mock-public-key-fingerprint",
						SelfAddressApi:       "self-address",
					},
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, _, mocks := newService(t, nil)
			tt.setup(t, mocks)

			want, err := service.ListOutways(tt.ctx, nil)
			assert.Equal(t, tt.want, want)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteOutway(t *testing.T) {
	tests := map[string]struct {
		request *api.DeleteOutwayRequest
		ctx     context.Context
		setup   func(serviceMocks)
		wantErr error
	}{
		"failed_to_retrieve_user_info_from_context": {
			ctx: context.Background(),
			request: &api.DeleteOutwayRequest{
				Name: "my-outway",
			},
			wantErr: status.Error(codes.Internal, "could not retrieve user info to create audit log"),
		},
		"missing_required_permission": {
			ctx: testCreateUserWithoutPermissionsContext(),
			request: &api.DeleteOutwayRequest{
				Name: "my-outway",
			},
			wantErr: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.outway.delete\" to execute this request"),
		},
		"failed_to_create_audit_log": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.al.EXPECT().
					OutwayDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-outway").
					Return(fmt.Errorf("error"))
			},
			request: &api.DeleteOutwayRequest{
				Name: "my-outway",
			},
			wantErr: status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"failed_to_delete_outway_from_database": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().
					DeleteOutway(gomock.Any(), "my-outway").
					Return(fmt.Errorf("error"))

				mocks.al.EXPECT().
					OutwayDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-outway")
			},
			request: &api.DeleteOutwayRequest{
				Name: "my-outway",
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					DeleteOutway(gomock.Any(), "my-outway")

				mocks.al.
					EXPECT().
					OutwayDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-outway")
			},
			request: &api.DeleteOutwayRequest{
				Name: "my-outway",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t, nil)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			_, err := service.DeleteOutway(tt.ctx, tt.request)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
