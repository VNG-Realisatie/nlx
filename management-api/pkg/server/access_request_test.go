// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/permissions"
	"go.nlx.io/nlx/management-api/pkg/server"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // this is a test
func TestApproveIncomingAccessRequest(t *testing.T) {
	tests := map[string]struct {
		auditLog func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger
		ctx      func() context.Context
		setup    func(serviceMocks)
		request  *api.ApproveIncomingAccessRequestRequest
		response *emptypb.Empty
		err      error
	}{

		"unknown_access_request": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			ctx: testCreateAdminUserContext,
			request: &api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequest(gomock.Any(), uint(1)).Return(nil, database.ErrNotFound)
			},
			err: status.Error(codes.NotFound, "access request not found"),
		},
		"access_request_already_approved": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequest(gomock.Any(), uint(1)).Return(&database.IncomingAccessRequest{
					ServiceID: 1,
					Service: &database.Service{
						Name: "test-service",
					},
					State: database.IncomingAccessRequestApproved,
				}, nil)
			},
			ctx: testCreateAdminUserContext,
			request: &api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			err: status.Error(codes.AlreadyExists, "access request is already approved"),
		},
		"missing_required_permission": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			setup: func(mocks serviceMocks) {

			},
			ctx: func() context.Context {
				ctx := context.Background()
				return context.WithValue(ctx, domain.UserKey, &domain.User{
					Email:       "admin@example.com",
					Permissions: map[permissions.Permission]bool{},
				})
			},
			request: &api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			err: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.incoming_access_request.approve\" to execute this request"),
		},
		"update_state_fails": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestAccept(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", "test-organization", "test-service")
				return auditLogger
			},
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequest(gomock.Any(), uint(1)).Return(&database.IncomingAccessRequest{
					ID: 1,
					Organization: database.IncomingAccessRequestOrganization{
						SerialNumber: "00000000000000000001",
						Name:         "test-organization",
					},
					Service: &database.Service{
						Name: "test-service",
					},
				}, nil)
				mocks.db.EXPECT().UpdateIncomingAccessRequestState(gomock.Any(), uint(1), database.IncomingAccessRequestApproved).Return(fmt.Errorf("arbitrary error"))
			},
			ctx: testCreateAdminUserContext,
			request: &api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			response: nil,
			err:      status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestAccept(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", "test-organization", "test-service")
				return auditLogger
			},
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequest(gomock.Any(), uint(1)).Return(&database.IncomingAccessRequest{
					ID: 1,
					Organization: database.IncomingAccessRequestOrganization{
						SerialNumber: "00000000000000000001",
						Name:         "test-organization",
					},
					Service: &database.Service{
						Name: "test-service",
					},
				}, nil)

				mocks.db.EXPECT().UpdateIncomingAccessRequestState(gomock.Any(), uint(1), database.IncomingAccessRequestApproved).Return(nil)

				mocks.db.EXPECT().CreateAccessGrant(gomock.Any(), &database.IncomingAccessRequest{
					ID: 1,
					Organization: database.IncomingAccessRequestOrganization{
						SerialNumber: "00000000000000000001",
						Name:         "test-organization",
					},
					Service: &database.Service{
						Name: "test-service",
					},
				})
			},
			ctx: testCreateAdminUserContext,
			request: &api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			response: &emptypb.Empty{},
			err:      nil,
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			test.auditLog(*mocks.al)

			test.setup(mocks)

			actual, err := service.ApproveIncomingAccessRequest(test.ctx(), test.request)
			assert.Equal(t, test.response, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

func testCreateAdminUserContext() context.Context {
	ctx := context.Background()

	return context.WithValue(ctx, domain.UserKey, &domain.User{
		Email:     "admin@example.com",
		UserAgent: "nlxctl",
		Permissions: map[permissions.Permission]bool{
			permissions.ApproveIncomingAccessRequest: true,
			permissions.RejectIncomingAccessRequest:  true,
			permissions.ReadIncomingAccessRequests:   true,
			permissions.CreateOutgoingAccessRequest:  true,
			permissions.UpdateOutgoingAccessRequest:  true,
			permissions.SendOutgoingAccessRequest:    true,
			permissions.ReadAccessGrants:             true,
			permissions.RevokeAccessGrant:            true,
			permissions.ReadAuditLogs:                true,
			permissions.ReadFinanceReport:            true,
			permissions.ReadInway:                    true,
			permissions.UpdateInway:                  true,
			permissions.DeleteInway:                  true,
			permissions.ReadInways:                   true,
			permissions.CreateOutgoingOrder:          true,
			permissions.UpdateOutgoingOrder:          true,
			permissions.RevokeOutgoingOrder:          true,
			permissions.ReadOutgoingOrders:           true,
			permissions.ReadIncomingOrders:           true,
			permissions.SynchronizeIncomingOrders:    true,
			permissions.ReadOutways:                  true,
			permissions.DeleteOutway:                 true,
			permissions.CreateService:                true,
			permissions.ReadService:                  true,
			permissions.UpdateService:                true,
			permissions.DeleteService:                true,
			permissions.ReadServices:                 true,
			permissions.ReadServicesStatistics:       true,
			permissions.ReadOrganizationSettings:     true,
			permissions.UpdateOrganizationSettings:   true,
			permissions.AcceptTermsOfService:         true,
			permissions.ReadTermsOfServiceStatus:     true,
			permissions.ReadTransactionLogs:          true,
		},
	})
}

func testCreateUserWithoutPermissionsContext() context.Context {
	ctx := context.Background()
	ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
		"username":               "admin@example.com",
		"grpcgateway-user-agent": "nlxctl",
	}))

	return context.WithValue(ctx, domain.UserKey, &domain.User{
		Email:       "admin@example.com",
		Permissions: map[permissions.Permission]bool{},
	})
}

//nolint:funlen // this is a test
func TestRejectIncomingAccessRequest(t *testing.T) {
	tests := map[string]struct {
		auditLog         func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger
		ctx              context.Context
		setup            func(mocks serviceMocks)
		request          *api.RejectIncomingAccessRequestRequest
		accessRequest    *database.IncomingAccessRequest
		accessRequestErr error
		expectUpdateCall bool
		updateErr        error
		response         *emptypb.Empty
		err              error
	}{
		"missing_required_permission": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			ctx:   testCreateUserWithoutPermissionsContext(),
			setup: func(mocks serviceMocks) {},
			request: &api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			err: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.incoming_access_request.reject\" to execute this request"),
		},
		"unknown_access_request": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequest(gomock.Any(), uint(1)).Return(nil, database.ErrNotFound)
			},
			request: &api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			err: status.Error(codes.NotFound, "access request not found"),
		},
		"update_state_fails": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestReject(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", "test-organization", "test-service")
				return auditLogger
			},
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequest(gomock.Any(), uint(1)).Return(&database.IncomingAccessRequest{
					ID: 1,
					Organization: database.IncomingAccessRequestOrganization{
						SerialNumber: "00000000000000000001",
						Name:         "test-organization",
					},
					Service: &database.Service{
						Name: "other-service",
					},
				}, nil)

				mocks.db.EXPECT().UpdateIncomingAccessRequestState(gomock.Any(), uint(1), database.IncomingAccessRequestRejected).Return(fmt.Errorf("arbitrary error"))
			},
			ctx: testCreateAdminUserContext(),
			request: &api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			err: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			auditLog: func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestReject(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", "test-organization", "test-service")
				return auditLogger
			},
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequest(gomock.Any(), uint(1)).Return(&database.IncomingAccessRequest{
					ID: 1,
					Organization: database.IncomingAccessRequestOrganization{
						SerialNumber: "00000000000000000001",
						Name:         "test-organization",
					},
					Service: &database.Service{
						Name: "other-service",
					},
				}, nil)

				mocks.db.EXPECT().UpdateIncomingAccessRequestState(gomock.Any(), uint(1), database.IncomingAccessRequestRejected).Return(nil)
			},
			ctx: testCreateAdminUserContext(),
			request: &api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			response: &emptypb.Empty{},
			err:      nil,
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			test.setup(mocks)

			test.auditLog(*mocks.al)

			actual, err := service.RejectIncomingAccessRequest(test.ctx, test.request)
			assert.Equal(t, test.response, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

// nolint funlen: this is a test
func TestExternalRequestAccess(t *testing.T) {
	tests := map[string]struct {
		setup   func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context
		want    *external.RequestAccessResponse
		wantErr error
	}{
		"when_peer_context_is_missing": {
			setup: func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_the_service_does_not_exist": {
			setup: func(_ *testing.T, db *mock_database.MockConfigDatabase, _ *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(nil, database.ErrNotFound)

				return ctx
			},
			wantErr: server.ErrServiceDoesNotExist,
		},
		"when_creating_the_access_request_errors": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, orgCert *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{
						ID:   1,
						Name: "Service",
					}, nil)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, gomock.Any(), "service", orgCert.PublicKeyFingerprint()).
					Return(nil, database.ErrNotFound)

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(nil, errors.New("error"))

				return ctx
			},
			wantErr: status.Error(codes.Internal, "failed to create access request"),
		},
		"returns_error_when_a_active_access_request_already_exists": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, orgCert *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{
						ID:   1,
						Name: "Service",
					}, nil)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, gomock.Any(), "service", orgCert.PublicKeyFingerprint()).
					Return(nil, database.ErrNotFound)

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(nil, database.ErrActiveAccessRequest)

				return ctx
			},
			wantErr: status.Error(codes.AlreadyExists, "an active access request already exists"),
		},
		"happy_flow": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, orgCert *tls.CertificateBundle) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{
						ID:   1,
						Name: "Service",
					}, nil)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, gomock.Any(), "service", orgCert.PublicKeyFingerprint()).
					Return(nil, database.ErrNotFound)

				pem, err := orgCert.PublicKeyPEM()
				require.NoError(t, err)

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, &database.IncomingAccessRequest{
						ServiceID: 1,
						Organization: database.IncomingAccessRequestOrganization{
							SerialNumber: certBundle.Certificate().Subject.SerialNumber,
							Name:         certBundle.Certificate().Subject.Organization[0],
						},
						State:                database.IncomingAccessRequestReceived,
						PublicKeyPEM:         pem,
						PublicKeyFingerprint: orgCert.PublicKeyFingerprint(),
					}).
					Return(&database.IncomingAccessRequest{
						ID: 42,
					}, nil)

				return ctx
			},
			want: &external.RequestAccessResponse{
				ReferenceId: 42,
			},
		},
		"happy_flow_existing_incoming_access_request": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, orgCert *tls.CertificateBundle) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{
						ID:   1,
						Name: "Service",
					}, nil)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, gomock.Any(), "service", orgCert.PublicKeyFingerprint()).
					Return(&database.IncomingAccessRequest{
						ID:        43,
						ServiceID: 1,
						Organization: database.IncomingAccessRequestOrganization{
							SerialNumber: certBundle.Certificate().Subject.SerialNumber,
							Name:         certBundle.Certificate().Subject.Organization[0],
						},
						State:                database.IncomingAccessRequestApproved,
						PublicKeyFingerprint: certBundle.PublicKeyFingerprint(),
					}, nil)

				return ctx
			},
			want: &external.RequestAccessResponse{
				ReferenceId:        43,
				AccessRequestState: api.AccessRequestState_APPROVED,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, orgBundle, mocks := newService(t)
			ctx := tt.setup(t, mocks.db, orgBundle)

			pem, err := orgBundle.PublicKeyPEM()
			assert.NoError(t, err)

			result, err := service.RequestAccess(ctx, &external.RequestAccessRequest{
				ServiceName:  "service",
				PublicKeyPem: pem,
			})

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.Equal(t, tt.want, result)
			}
		})
	}
}

// nolint funlen: this is a test
func TestExternalGetAccessRequestState(t *testing.T) {
	tests := map[string]struct {
		setup   func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context
		want    *external.GetAccessRequestStateResponse
		wantErr error
	}{
		"when_peer_context_is_missing": {
			setup: func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_retrieving_the_service_fails": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(nil, errors.New("error"))

				return ctx
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"when_retrieving_state_errors": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{}, nil)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, certBundle.Certificate().Subject.SerialNumber, "service", "public-key-fingerprint").
					Return(nil, errors.New("error"))

				return ctx
			},
			wantErr: status.Error(codes.Internal, "failed to retrieve access request"),
		},
		"when_the_service_does_not_exists": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(nil, database.ErrNotFound)

				return ctx
			},
			wantErr: server.ErrServiceDoesNotExist,
		},
		"happy_flow": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{}, nil)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(
						ctx,
						certBundle.Certificate().Subject.SerialNumber,
						"service",
						"public-key-fingerprint",
					).
					Return(&database.IncomingAccessRequest{
						State: database.IncomingAccessRequestReceived,
					}, nil)

				return ctx
			},
			want: &external.GetAccessRequestStateResponse{
				State: api.AccessRequestState_RECEIVED,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, certBundle, mocks := newService(t)
			ctx := tt.setup(t, mocks.db, certBundle)

			actual, err := service.GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
				ServiceName:          "service",
				PublicKeyFingerprint: "public-key-fingerprint",
			})

			assert.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}

// nolint funlen: this is a test
func TestListIncomingAccessRequests(t *testing.T) {
	currentTime := time.Now()

	tests := map[string]struct {
		setup   func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context
		ctx     context.Context
		want    *api.ListIncomingAccessRequestsResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx: testCreateUserWithoutPermissionsContext(),
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)
				return ctx
			},
			wantErr: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.incoming_access_requests.read\" to execute this request"),
		},
		"when_retrieving_the_service_fails": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(gomock.Any(), "service").
					Return(nil, errors.New("error"))

				return ctx
			},
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"when_the_service_does_not_exists": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(gomock.Any(), "service").
					Return(nil, database.ErrNotFound)

				return ctx
			},
			wantErr: server.ErrServiceDoesNotExist,
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetService(gomock.Any(), "service").
					Return(&database.Service{}, nil)

				db.
					EXPECT().
					ListIncomingAccessRequests(
						gomock.Any(),
						"service",
					).
					Return([]*database.IncomingAccessRequest{
						{
							ID:        1,
							ServiceID: 1,
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "test-organization",
							},
							Service: &database.Service{
								ID:          1,
								Name:        "service-name",
								EndpointURL: "https://example.com",
								CreatedAt:   currentTime,
								UpdatedAt:   currentTime,
							},
							State:                database.IncomingAccessRequestReceived,
							PublicKeyFingerprint: "public-key-fingerprint",
							CreatedAt:            currentTime,
							UpdatedAt:            currentTime,
						},
					}, nil)

				return ctx
			},
			want: &api.ListIncomingAccessRequestsResponse{
				AccessRequests: []*api.IncomingAccessRequest{
					{
						Id: 1,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
						ServiceName:          "service-name",
						State:                api.AccessRequestState_RECEIVED,
						CreatedAt:            timestamppb.New(currentTime),
						UpdatedAt:            timestamppb.New(currentTime),
						PublicKeyFingerprint: "public-key-fingerprint",
					},
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, certBundle, mocks := newService(t)
			_ = tt.setup(t, mocks.db, certBundle)

			actual, err := service.ListIncomingAccessRequests(tt.ctx, &api.ListIncomingAccessRequestsRequest{
				ServiceName: "service",
			})

			assert.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
