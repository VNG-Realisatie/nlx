// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen,dupl // this is a test
func Test_SendAccessRequest(t *testing.T) {
	now := time.Now()

	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	testPublicKeyPEM, err := certBundle.PublicKeyPEM()
	require.NoError(t, err)

	testPublicKeyFingerprint := certBundle.PublicKeyFingerprint()

	tests := map[string]struct {
		ctx        context.Context
		setupMocks func(mocks serviceMocks)
		req        *api.SendAccessRequestRequest
		want       *api.SendAccessRequestResponse
		wantErr    error
	}{
		"missing_required_permission": {
			ctx:        testCreateUserWithoutPermissionsContext(),
			setupMocks: func(mocks serviceMocks) {},
			req:        &api.SendAccessRequestRequest{},
			want:       nil,
			wantErr:    status.New(codes.PermissionDenied, "user needs the permission \"permissions.outgoing_access_request.send\" to execute this request").Err(),
		},
		"invalid_fingerprint_provided": {
			ctx:        testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {},
			req: &api.SendAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
				PublicKeyPEM:             "arbitrary-pem",
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "invalid public key format").Err(),
		},
		"failed_to_retrieve_organization_inway_proxy_address": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("", errors.New("arbitrary error"))

				mocks.db.
					EXPECT().
					CreateOutgoingAccessRequest(gomock.Any(), &database.OutgoingAccessRequest{
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "",
						},
						ReferenceID:          0,
						ServiceName:          "my-service",
						PublicKeyPEM:         testPublicKeyPEM,
						PublicKeyFingerprint: testPublicKeyFingerprint,
						State:                database.OutgoingAccessRequestFailed,
						ErrorCause:           "The organization is not available.",
					}).
					Return(&database.OutgoingAccessRequest{
						ID: 42,
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "my-organization",
						},
						ServiceName:          "my-service",
						State:                database.OutgoingAccessRequestFailed,
						CreatedAt:            now,
						UpdatedAt:            now,
						ErrorCause:           "The organization is not available.",
						PublicKeyPEM:         testPublicKeyPEM,
						PublicKeyFingerprint: testPublicKeyFingerprint,
					}, nil)

				mocks.al.
					EXPECT().
					OutgoingAccessRequestCreate(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"my-service",
					).
					Return(nil)
			},
			req: &api.SendAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
				PublicKeyPEM:             testPublicKeyPEM,
			},
			want: &api.SendAccessRequestResponse{
				OutgoingAccessRequest: &api.OutgoingAccessRequest{
					Id: 42,
					Organization: &api.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "my-organization",
					},
					ServiceName:          "my-service",
					State:                api.AccessRequestState_ACCESS_REQUEST_STATE_FAILED,
					CreatedAt:            timestamppb.New(now),
					UpdatedAt:            timestamppb.New(now),
					PublicKeyFingerprint: testPublicKeyFingerprint,
					ErrorDetails: &api.ErrorDetails{
						Code:       api.ErrorCode_INTERNAL,
						Cause:      "The organization is not available.",
						StackTrace: nil,
					},
				},
			},
			wantErr: nil,
		},
		"failed_to_request_access": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName:  "my-service",
						PublicKeyPem: testPublicKeyPEM,
					}).Return(nil, errors.New("arbitrary"))

				mocks.db.
					EXPECT().
					CreateOutgoingAccessRequest(gomock.Any(), &database.OutgoingAccessRequest{
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "",
						},
						ReferenceID:          0,
						ServiceName:          "my-service",
						PublicKeyPEM:         testPublicKeyPEM,
						PublicKeyFingerprint: testPublicKeyFingerprint,
						State:                database.OutgoingAccessRequestFailed,
						ErrorCause:           "The organization is not available.",
					}).
					Return(&database.OutgoingAccessRequest{
						ID: 42,
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "my-organization",
						},
						ServiceName:          "my-service",
						State:                database.OutgoingAccessRequestFailed,
						CreatedAt:            now,
						UpdatedAt:            now,
						ErrorCause:           "The organization is not available.",
						PublicKeyPEM:         testPublicKeyPEM,
						PublicKeyFingerprint: testPublicKeyFingerprint,
					}, nil)

				mocks.al.
					EXPECT().
					OutgoingAccessRequestCreate(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"my-service",
					).
					Return(nil)
			},
			req: &api.SendAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
				PublicKeyPEM:             testPublicKeyPEM,
			},
			want: &api.SendAccessRequestResponse{
				OutgoingAccessRequest: &api.OutgoingAccessRequest{
					Id: 42,
					Organization: &api.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "my-organization",
					},
					ServiceName:          "my-service",
					State:                api.AccessRequestState_ACCESS_REQUEST_STATE_FAILED,
					CreatedAt:            timestamppb.New(now),
					UpdatedAt:            timestamppb.New(now),
					PublicKeyFingerprint: testPublicKeyFingerprint,
					ErrorDetails: &api.ErrorDetails{
						Code:       api.ErrorCode_INTERNAL,
						Cause:      "The organization is not available.",
						StackTrace: nil,
					},
				},
			},
			wantErr: nil,
		},
		"outgoing_access_request_already_present": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName:  "my-service",
						PublicKeyPem: testPublicKeyPEM,
					}).
					Return(&external.RequestAccessResponse{
						ReferenceId: 1,
					}, nil)

				mocks.db.
					EXPECT().CreateOutgoingAccessRequest(gomock.Any(), &database.OutgoingAccessRequest{
					Organization: database.Organization{
						SerialNumber: "00000000000000000001",
					},
					ReferenceID:          1,
					ServiceName:          "my-service",
					PublicKeyPEM:         testPublicKeyPEM,
					PublicKeyFingerprint: testPublicKeyFingerprint,
					State:                database.OutgoingAccessRequestReceived,
				}).Return(nil, database.ErrActiveAccessRequest)
			},
			req: &api.SendAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
				PublicKeyPEM:             testPublicKeyPEM,
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "internal").Err(),
		},
		"failed_to_create_outgoing_access_request": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName:  "my-service",
						PublicKeyPem: testPublicKeyPEM,
					}).
					Return(&external.RequestAccessResponse{
						ReferenceId: 1,
					}, nil)

				mocks.db.
					EXPECT().
					CreateOutgoingAccessRequest(gomock.Any(), &database.OutgoingAccessRequest{
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
						},
						ReferenceID:          1,
						ServiceName:          "my-service",
						PublicKeyPEM:         testPublicKeyPEM,
						PublicKeyFingerprint: testPublicKeyFingerprint,
						State:                database.OutgoingAccessRequestReceived,
					}).
					Return(nil, errors.New("arbitrary error"))
			},
			req: &api.SendAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
				PublicKeyPEM:             testPublicKeyPEM,
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "internal").Err(),
		},
		"failed_to_create_audit_log": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName:  "my-service",
						PublicKeyPem: testPublicKeyPEM,
					}).
					Return(&external.RequestAccessResponse{
						ReferenceId: 1,
					}, nil)

				mocks.db.
					EXPECT().
					CreateOutgoingAccessRequest(gomock.Any(), &database.OutgoingAccessRequest{
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
						},
						ReferenceID:          1,
						ServiceName:          "my-service",
						PublicKeyPEM:         testPublicKeyPEM,
						PublicKeyFingerprint: testPublicKeyFingerprint,
						State:                database.OutgoingAccessRequestReceived,
					}).
					Return(&database.OutgoingAccessRequest{}, nil)

				mocks.al.
					EXPECT().
					OutgoingAccessRequestCreate(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"my-service",
					).
					Return(errors.New("arbitrary error"))
			},
			req: &api.SendAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
				PublicKeyPEM:             testPublicKeyPEM,
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "internal").Err(),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName:  "my-service",
						PublicKeyPem: testPublicKeyPEM,
					}).
					Return(&external.RequestAccessResponse{
						ReferenceId: 1,
					}, nil)

				mocks.db.
					EXPECT().
					CreateOutgoingAccessRequest(gomock.Any(), &database.OutgoingAccessRequest{
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
						},
						ReferenceID:          1,
						ServiceName:          "my-service",
						PublicKeyPEM:         testPublicKeyPEM,
						PublicKeyFingerprint: testPublicKeyFingerprint,
						State:                database.OutgoingAccessRequestReceived,
					}).
					Return(&database.OutgoingAccessRequest{
						ID: 42,
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "my-organization",
						},
						ServiceName:          "my-service",
						State:                database.OutgoingAccessRequestReceived,
						CreatedAt:            now,
						UpdatedAt:            now,
						PublicKeyFingerprint: testPublicKeyFingerprint,
					}, nil)

				mocks.al.
					EXPECT().
					OutgoingAccessRequestCreate(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"my-service",
					).
					Return(nil)
			},
			req: &api.SendAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
				PublicKeyPEM:             testPublicKeyPEM,
			},
			want: &api.SendAccessRequestResponse{
				OutgoingAccessRequest: &api.OutgoingAccessRequest{
					Id: 42,
					Organization: &api.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "my-organization",
					},
					ServiceName:          "my-service",
					State:                api.AccessRequestState_ACCESS_REQUEST_STATE_RECEIVED,
					CreatedAt:            timestamppb.New(now),
					UpdatedAt:            timestamppb.New(now),
					PublicKeyFingerprint: testPublicKeyFingerprint,
				},
			},
			wantErr: nil,
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			test.setupMocks(mocks)
			got, err := service.SendAccessRequest(test.ctx, test.req)

			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
