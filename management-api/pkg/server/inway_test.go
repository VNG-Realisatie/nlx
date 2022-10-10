// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestRegisterInway(t *testing.T) {

	tests := map[string]struct {
		peer       *peer.Peer
		setupMocks func(mocks serviceMocks)
		request    *api.Inway
		wantErr    error
		want       *api.Inway
	}{
		"peer_does_not_contain_address": {
			peer:    &peer.Peer{Addr: nil},
			request: &api.Inway{Name: "inway42.basic"},
			wantErr: status.Error(codes.Internal, "peer addr is invalid"),
		},
		"register_inway_database_call_fails": {
			peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}},
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().RegisterInway(gomock.Any(), &database.Inway{Name: "inway42.basic", IPAddress: "127.1.1.1", CreatedAt: fixtureTime, UpdatedAt: fixtureTime}).Return(fmt.Errorf("arbitrary error"))
			},
			request: &api.Inway{Name: "inway42.basic"},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"happy_flow_address_from_peer": {
			peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}},
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().RegisterInway(gomock.Any(), &database.Inway{Name: "inway42.basic", IPAddress: "127.1.1.1", CreatedAt: fixtureTime, UpdatedAt: fixtureTime})
			},
			request: &api.Inway{Name: "inway42.basic"},
			want: &api.Inway{
				Name:      "inway42.basic",
				IpAddress: "127.1.1.1",
			},
		},
		"happy_flow_ipv6_address_from_peer": {
			peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}},
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().RegisterInway(gomock.Any(), &database.Inway{Name: "inway42.basic", IPAddress: "::1", CreatedAt: fixtureTime, UpdatedAt: fixtureTime})
			},
			request: &api.Inway{Name: "inway42.basic"},
			want: &api.Inway{
				Name:      "inway42.basic",
				IpAddress: "::1",
			},
		},
		"happy_flow_ip_address_from_request_ignored": {
			peer: &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}},
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().RegisterInway(gomock.Any(), &database.Inway{Name: "inway42.basic", IPAddress: "127.1.1.1", CreatedAt: fixtureTime, UpdatedAt: fixtureTime})
			},
			request: &api.Inway{Name: "inway42.basic", IpAddress: "127.2.2.2"},
			want: &api.Inway{
				Name:      "inway42.basic",
				IpAddress: "127.1.1.1",
			},
		},
	}

	for name, test := range tests {
		tt := test

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			if tt.setupMocks != nil {
				tt.setupMocks(mocks)
			}

			ctx := peer.NewContext(context.Background(), tt.peer)

			response, err := service.RegisterInway(ctx, tt.request)

			assert.Equal(t, tt.want, response)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetInway(t *testing.T) {
	tests := map[string]struct {
		ctx        context.Context
		setupMocks func(mocks serviceMocks)
		request    *api.GetInwayRequest
		wantErr    error
		want       *api.Inway
	}{
		"missing_required_permission": {
			ctx:        testCreateUserWithoutPermissionsContext(),
			setupMocks: func(mocks serviceMocks) {},
			request: &api.GetInwayRequest{
				Name: "inway42.test",
			},
			wantErr: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.inway.read\" to execute this request"),
		},
		"not_found": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42.test").Return(nil, database.ErrNotFound)
			},
			request: &api.GetInwayRequest{
				Name: "inway42.test",
			},
			wantErr: status.Error(codes.NotFound, "inway not found"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42.test").Return(&database.Inway{
					Name:      "inway42.test",
					IPAddress: "",
					Services: []*database.Service{{
						Name: "forty-two",
					}},
				}, nil)
			},
			request: &api.GetInwayRequest{
				Name: "inway42.test",
			},
			want: &api.Inway{
				Name:     "inway42.test",
				Services: []*api.Inway_Service{{Name: "forty-two"}},
			},
		},
	}

	for name, test := range tests {
		tt := test

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			tt.setupMocks(mocks)

			response, err := service.GetInway(tt.ctx, tt.request)

			assert.Equal(t, tt.want, response)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

//nolint:funlen // this is a test
func TestUpdateInway(t *testing.T) {
	tests := map[string]struct {
		ctx        context.Context
		setupMocks func(mocks serviceMocks)
		request    *api.UpdateInwayRequest
		wantErr    error
		want       *api.Inway
	}{
		"missing_required_permission": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42.test").Return(nil, database.ErrNotFound)
			},
			request: &api.UpdateInwayRequest{
				Name: "inway42.test",
				Inway: &api.Inway{
					Name: "inway42.test",
				},
			},
			wantErr: status.Error(codes.NotFound, "inway with the name inway42.test does not exist"),
		},
		"inway_not_found": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42.test").Return(nil, database.ErrNotFound)
			},
			request: &api.UpdateInwayRequest{
				Name: "inway42.test",
				Inway: &api.Inway{
					Name: "inway42.test",
				},
			},
			wantErr: status.Error(codes.NotFound, "inway with the name inway42.test does not exist"),
		},
		"get_inway_from_database_fails": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42.test").Return(nil, fmt.Errorf("arbitrary error"))
			},
			request: &api.UpdateInwayRequest{
				Name: "inway42.test",
				Inway: &api.Inway{
					Name: "inway42.test",
				},
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"update_inway_fails": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mockInway := &database.Inway{
					ID:   1,
					Name: "inway42.test",
				}
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42.test").Return(mockInway, nil)

				mocks.db.EXPECT().UpdateInway(gomock.Any(), mockInway).Return(fmt.Errorf("arbitrary error"))
			},
			request: &api.UpdateInwayRequest{
				Name: "inway42.test",
				Inway: &api.Inway{
					Name: "inway42.test",
				},
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mockInway := &database.Inway{
					ID:   1,
					Name: "inway42.test",
				}
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42.test").Return(mockInway, nil)

				mocks.db.EXPECT().UpdateInway(gomock.Any(), mockInway)
			},
			request: &api.UpdateInwayRequest{
				Name: "inway42.test",
				Inway: &api.Inway{
					Name: "inway42.test",
				},
			},
			want: &api.Inway{
				Name: "inway42.test",
			},
		},
	}

	for name, test := range tests {
		tt := test

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			tt.setupMocks(mocks)

			response, err := service.UpdateInway(tt.ctx, tt.request)

			assert.Equal(t, tt.want, response)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestDeleteInway(t *testing.T) {
	tests := map[string]struct {
		request *api.DeleteInwayRequest
		ctx     context.Context
		setup   func(serviceMocks)
		wantErr error
	}{
		"failed_to_retrieve_user_info_from_context": {
			ctx: context.Background(),
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
			wantErr: status.Error(codes.Internal, "could not retrieve user info to create audit log"),
		},
		"missing_required_permission": {
			ctx: testCreateUserWithoutPermissionsContext(),
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
			wantErr: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.inway.delete\" to execute this request"),
		},
		"failed_to_create_audit_log": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.al.EXPECT().
					InwayDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-inway").
					Return(fmt.Errorf("error"))
			},
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
			wantErr: status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"failed_to_delete_inway_from_database": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().
					DeleteInway(gomock.Any(), "my-inway").
					Return(fmt.Errorf("error"))

				mocks.al.EXPECT().
					InwayDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-inway")
			},
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					DeleteInway(gomock.Any(), "my-inway")

				mocks.al.
					EXPECT().
					InwayDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-inway")
			},
			request: &api.DeleteInwayRequest{
				Name: "my-inway",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			_, err := service.DeleteInway(tt.ctx, tt.request)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestListInways(t *testing.T) {
	tests := map[string]struct {
		ctx     context.Context
		setup   func(serviceMocks)
		want    *api.ListInwaysResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(mocks serviceMocks) {},
			wantErr: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.inways.read\" to execute this request"),
		},
		"database_call_fails": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().ListInways(gomock.Any()).Return(nil, fmt.Errorf("arbitrary error"))
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().ListInways(gomock.Any()).Return([]*database.Inway{
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
				}, nil)
			},
			want: &api.ListInwaysResponse{
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
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			response, err := service.ListInways(tt.ctx, &api.ListInwaysRequest{})

			assert.Equal(t, tt.want, response)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

//nolint:funlen // this is a test
func TestGetInwayConfig(t *testing.T) {
	tests := map[string]struct {
		ctx        context.Context
		setupMocks func(mocks serviceMocks)
		request    *api.GetInwayConfigRequest
		wantErr    error
		want       *api.GetInwayConfigResponse
	}{
		"when_get_inway_database_call_fails": {
			ctx: context.Background(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42").Return(nil, errors.New("arbitrary error"))
			},
			request: &api.GetInwayConfigRequest{
				Name: "inway42",
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"when_get_settings_database_call_fails": {
			ctx: context.Background(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42").Return(&database.Inway{
					Services: []*database.Service{
						{
							Name: "service1",
						},
						{
							Name: "service2",
						},
					},
				}, nil)

				mocks.db.EXPECT().GetSettings(gomock.Any()).Return(nil, errors.New("arbitrary error"))
			},
			request: &api.GetInwayConfigRequest{
				Name: "inway42",
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"when_inway_not_found": {
			ctx: context.Background(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42").Return(nil, database.ErrNotFound)
			},
			request: &api.GetInwayConfigRequest{
				Name: "inway42",
			},
			wantErr: status.Error(codes.NotFound, "inway not found"),
		},
		"happy_flow_when_not_org_inway": {
			ctx: context.Background(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42").Return(&database.Inway{
					Name: "inway42",
					Services: []*database.Service{
						{
							Name: "service1",
						},
					},
				}, nil)

				settings, err := domain.NewSettings("inway99", "test@example.com")
				assert.NoError(t, err)

				mocks.db.EXPECT().GetSettings(gomock.Any()).Return(settings, nil)

				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "service1").Return([]*database.AccessGrant{
					{
						IncomingAccessRequest: &database.IncomingAccessRequest{
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: "1111",
								Name:         "org1",
							},
							PublicKeyFingerprint: "abc",
							PublicKeyPEM:         "def",
						},
					},
				}, nil)
			},
			request: &api.GetInwayConfigRequest{
				Name: "inway42",
			},
			want: &api.GetInwayConfigResponse{
				IsOrganizationInway: false,
				Services: []*api.GetInwayConfigResponse_Service{
					{
						Name: "service1",
						AuthorizationSettings: &api.GetInwayConfigResponse_Service_AuthorizationSettings{
							Authorizations: []*api.GetInwayConfigResponse_Service_AuthorizationSettings_Authorization{
								{
									Organization: &external.Organization{
										SerialNumber: "1111",
										Name:         "org1",
									},
									PublicKeyHash: "abc",
									PublicKeyPem:  "def",
								},
							},
						},
					},
				},
			},
		},
		"happy_flow": {
			ctx: context.Background(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetInway(gomock.Any(), "inway42").Return(&database.Inway{
					Name: "inway42",
					Services: []*database.Service{
						{
							Name: "service1",
						},
						{
							Name: "service2",
						},
					},
				}, nil)

				settings, err := domain.NewSettings("inway42", "test@example.com")
				assert.NoError(t, err)

				mocks.db.EXPECT().GetSettings(gomock.Any()).Return(settings, nil)

				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "service1").Return([]*database.AccessGrant{
					{
						IncomingAccessRequest: &database.IncomingAccessRequest{
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: "1111",
								Name:         "org1",
							},
							PublicKeyFingerprint: "abc",
							PublicKeyPEM:         "def",
						},
					},
					{
						IncomingAccessRequest: &database.IncomingAccessRequest{
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: "2222",
								Name:         "org2",
							},
							PublicKeyFingerprint: "uvw",
							PublicKeyPEM:         "xyz",
						},
					},
				}, nil)

				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "service2").Return([]*database.AccessGrant{
					{
						IncomingAccessRequest: &database.IncomingAccessRequest{
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: "3333",
								Name:         "org3",
							},
							PublicKeyFingerprint: "ghi",
							PublicKeyPEM:         "jkl",
						},
					},
					{
						IncomingAccessRequest: &database.IncomingAccessRequest{
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: "4444",
								Name:         "org4",
							},
							PublicKeyFingerprint: "mno",
							PublicKeyPEM:         "pqr",
						},
					},
				}, nil)
			},
			request: &api.GetInwayConfigRequest{
				Name: "inway42",
			},
			want: &api.GetInwayConfigResponse{
				IsOrganizationInway: true,
				Services: []*api.GetInwayConfigResponse_Service{
					{
						Name: "service1",
						AuthorizationSettings: &api.GetInwayConfigResponse_Service_AuthorizationSettings{
							Authorizations: []*api.GetInwayConfigResponse_Service_AuthorizationSettings_Authorization{
								{
									Organization: &external.Organization{
										SerialNumber: "1111",
										Name:         "org1",
									},
									PublicKeyHash: "abc",
									PublicKeyPem:  "def",
								},
								{
									Organization: &external.Organization{
										SerialNumber: "2222",
										Name:         "org2",
									},
									PublicKeyHash: "uvw",
									PublicKeyPem:  "xyz",
								},
							},
						},
					},
					{
						Name: "service2",
						AuthorizationSettings: &api.GetInwayConfigResponse_Service_AuthorizationSettings{
							Authorizations: []*api.GetInwayConfigResponse_Service_AuthorizationSettings_Authorization{
								{
									Organization: &external.Organization{
										SerialNumber: "3333",
										Name:         "org3",
									},
									PublicKeyHash: "ghi",
									PublicKeyPem:  "jkl",
								},
								{
									Organization: &external.Organization{
										SerialNumber: "4444",
										Name:         "org4",
									},
									PublicKeyHash: "mno",
									PublicKeyPem:  "pqr",
								},
							},
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		tt := test

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			tt.setupMocks(mocks)

			response, err := service.GetInwayConfig(tt.ctx, tt.request)

			assert.Equal(t, tt.want, response)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
