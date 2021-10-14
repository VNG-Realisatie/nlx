// Copyright © VNG Realisatie 2020
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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/diagnostics"
	common_testing "go.nlx.io/nlx/common/testing"
	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func createTimestamp(ti time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: ti.Unix(),
		Nanos:   int32(ti.Nanosecond()),
	}
}

func TestCreateAccessRequest(t *testing.T) {
	tests := map[string]struct {
		setup   func(*testing.T, serviceMocks, *tls.CertificateBundle) context.Context
		request *api.CreateAccessRequestRequest
		want    *api.OutgoingAccessRequest
		wantErr error
	}{
		"without_an_active_access_request": {
			setup: func(t *testing.T, mocks serviceMocks, certBundle *tls.CertificateBundle) context.Context {
				ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
					"username":               "Jane Doe",
					"grpcgateway-user-agent": "nlxctl",
				}))

				mocks.al.
					EXPECT().
					OutgoingAccessRequestCreate(ctx, "Jane Doe", "nlxctl", "00000000000000000001", "test-service").
					Return(nil)

				publicKeyPEM, err := certBundle.PublicKeyPEM()
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					CreateOutgoingAccessRequest(ctx, &database.OutgoingAccessRequest{
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
						},
						ServiceName:          "test-service",
						PublicKeyPEM:         publicKeyPEM,
						PublicKeyFingerprint: certBundle.PublicKeyFingerprint(),
						State:                database.OutgoingAccessRequestCreated,
					}).
					Return(&database.OutgoingAccessRequest{
						ID: 1,
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
						ServiceName:          "test-service",
						PublicKeyPEM:         publicKeyPEM,
						PublicKeyFingerprint: certBundle.PublicKeyFingerprint(),
						State:                database.OutgoingAccessRequestCreated,
						CreatedAt:            time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
						UpdatedAt:            time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
					}, nil)

				return ctx
			},
			request: &api.CreateAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			want: &api.OutgoingAccessRequest{
				Id: 1,
				Organization: &api.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName: "test-service",
				State:       api.AccessRequestState_CREATED,
				CreatedAt:   createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
				UpdatedAt:   createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
			},
		},
		"with_an_active_access_request": {
			setup: func(t *testing.T, mocks serviceMocks, certBundle *tls.CertificateBundle) context.Context {
				ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
					"username":               "Jane Doe",
					"grpcgateway-user-agent": "nlxctl",
				}))

				mocks.al.
					EXPECT().
					OutgoingAccessRequestCreate(ctx, "Jane Doe", "nlxctl", "00000000000000000001", "test-service").
					Return(nil)

				publicKeyPEM, err := certBundle.PublicKeyPEM()
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					CreateOutgoingAccessRequest(ctx, &database.OutgoingAccessRequest{
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
						},
						ServiceName:          "test-service",
						PublicKeyPEM:         publicKeyPEM,
						PublicKeyFingerprint: certBundle.PublicKeyFingerprint(),
						State:                database.OutgoingAccessRequestCreated,
					}).
					Return(nil, database.ErrActiveAccessRequest)

				return ctx
			},
			request: &api.CreateAccessRequestRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			wantErr: status.New(codes.AlreadyExists, "there is already an active access request").Err(),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, certBundle, mocks := newService(t)
			ctx := tt.setup(t, mocks, certBundle)

			actual, err := service.CreateAccessRequest(ctx, tt.request)

			assert.Equal(t, tt.want, actual)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

//nolint:funlen // this is a test
func TestSendAccessRequest(t *testing.T) {
	tests := []struct {
		name             string
		request          *api.SendAccessRequestRequest
		accessRequest    *database.OutgoingAccessRequest
		accessRequestErr error
		updateMock       func(mock *gomock.Call)
		response         *api.OutgoingAccessRequest
		responseErr      error
	}{
		{
			"non_existing",
			&api.SendAccessRequestRequest{
				AccessRequestID: 1,
			},
			nil,
			database.ErrNotFound,
			func(mock *gomock.Call) {
				mock.MaxTimes(0)
			},
			nil,
			status.New(codes.NotFound, "access request not found").Err(),
		},
		{
			"database_error",
			&api.SendAccessRequestRequest{
				AccessRequestID: 1,
			},
			nil,
			errors.New("an error"),
			func(mock *gomock.Call) {
				mock.MaxTimes(0)
			},
			nil,
			status.New(codes.Internal, "database error").Err(),
		},
		{
			"update_failed",
			&api.SendAccessRequestRequest{
				AccessRequestID: 1,
			},
			&database.OutgoingAccessRequest{

				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName: "test-service",
				State:       database.OutgoingAccessRequestCreated,
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(errors.New("an error"))
			},
			nil,
			status.New(codes.Internal, "database error").Err(),
		},
		{
			"created_state",
			&api.SendAccessRequestRequest{
				AccessRequestID: 1,
			},
			&database.OutgoingAccessRequest{
				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName: "test-service",
				State:       database.OutgoingAccessRequestCreated,
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(nil)
			},
			&api.OutgoingAccessRequest{
				Id: 1,
				Organization: &api.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName: "test-service",
				State:       api.AccessRequestState_CREATED,
				CreatedAt:   createTimestamp(time.Time{}),
				UpdatedAt:   createTimestamp(time.Time{}),
			},
			nil,
		},
		{
			"failed_state",
			&api.SendAccessRequestRequest{
				AccessRequestID: 1,
			},
			&database.OutgoingAccessRequest{
				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName: "test-service",
				State:       database.OutgoingAccessRequestFailed,
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(nil)
			},
			&api.OutgoingAccessRequest{
				Id: 1,
				Organization: &api.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName: "test-service",
				State:       api.AccessRequestState_CREATED,
				CreatedAt:   createTimestamp(time.Time{}),
				UpdatedAt:   createTimestamp(time.Time{}),
			},
			nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			service, _, mocks := newService(t)
			ctx := context.Background()

			mocks.db.
				EXPECT().
				GetOutgoingAccessRequest(ctx, uint(test.request.AccessRequestID)).
				Return(test.accessRequest, test.accessRequestErr)

			updateMock := mocks.db.
				EXPECT().
				UpdateOutgoingAccessRequestState(ctx, uint(test.request.AccessRequestID), database.OutgoingAccessRequestCreated, uint(0), nil).
				Do(func(_ context.Context, _ uint, state database.OutgoingAccessRequestState, _ uint, errorDetails *diagnostics.ErrorDetails) error {
					test.accessRequest.State = state
					return nil
				})
			test.updateMock(updateMock)

			response, err := service.SendAccessRequest(ctx, test.request)

			assert.Equal(t, test.response, response)
			assert.Equal(t, test.responseErr, err)
		})
	}
}

//nolint:funlen // this is a test
func TestApproveIncomingAccessRequest(t *testing.T) {
	tests := []struct {
		name             string
		auditLog         func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger
		ctx              context.Context
		request          *api.ApproveIncomingAccessRequestRequest
		accessRequest    *database.IncomingAccessRequest
		accessRequestErr error
		expectUpdateCall bool
		updateErr        error
		response         *emptypb.Empty
		err              error
	}{
		{
			"unknown_access_request",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			context.Background(),
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},

			nil,
			database.ErrNotFound,
			false,
			nil,
			nil,
			status.Error(codes.NotFound, "access request not found"),
		},
		{
			"access_request_already_approved",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			context.Background(),
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			&database.IncomingAccessRequest{
				ServiceID: 1,
				Service: &database.Service{
					Name: "test-service",
				},
				State: database.IncomingAccessRequestApproved,
			},
			nil,
			false,
			nil,
			nil,
			status.Error(codes.AlreadyExists, "access request is already approved"),
		},
		{
			"update_state_fails",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestAccept(gomock.Any(), "Jane Doe", "nlxctl", "test-organization", "test-service")
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			&database.IncomingAccessRequest{
				Organization: database.IncomingAccessRequestOrganization{
					Name: "test-organization",
				},
				Service: &database.Service{
					Name: "test-service",
				},
			},
			nil,
			true,
			errors.New("arbitrary error"),
			nil,
			status.Error(codes.Internal, "database error"),
		},
		{
			"happy_flow",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestAccept(gomock.Any(), "Jane Doe", "nlxctl", "test-organization", "test-service")
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			&database.IncomingAccessRequest{
				Organization: database.IncomingAccessRequestOrganization{
					Name: "test-organization",
				},
				Service: &database.Service{
					Name: "test-service",
				},
			},
			nil,
			true,
			nil,
			&emptypb.Empty{},
			nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			service, _, mocks := newService(t)
			test.auditLog(*mocks.al)

			mocks.db.EXPECT().GetIncomingAccessRequest(test.ctx, uint(test.request.AccessRequestID)).Return(test.accessRequest, test.accessRequestErr)

			if test.response != nil {
				mocks.db.EXPECT().CreateAccessGrant(test.ctx, test.accessRequest)
			}

			if test.expectUpdateCall {
				mocks.db.EXPECT().UpdateIncomingAccessRequestState(test.ctx, test.accessRequest.ID, database.IncomingAccessRequestApproved).Return(test.updateErr)
			}

			actual, err := service.ApproveIncomingAccessRequest(test.ctx, test.request)
			assert.Equal(t, test.response, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

//nolint:funlen // this is a test
func TestRejectIncomingAccessRequest(t *testing.T) {
	tests := []struct {
		name             string
		auditLog         func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger
		ctx              context.Context
		request          *api.RejectIncomingAccessRequestRequest
		accessRequest    *database.IncomingAccessRequest
		accessRequestErr error
		expectUpdateCall bool
		updateErr        error
		response         *emptypb.Empty
		err              error
	}{
		{
			"unknown_access_request",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			nil,
			database.ErrNotFound,
			false,
			nil,
			nil,
			status.Error(codes.NotFound, "access request not found"),
		},
		{
			"update_state_fails",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestReject(gomock.Any(), "Jane Doe", "nlxctl", "test-organization", "test-service")
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			&database.IncomingAccessRequest{
				Organization: database.IncomingAccessRequestOrganization{
					Name: "test-organization",
				},
				Service: &database.Service{
					Name: "other-service",
				},
			},
			nil,
			true,
			errors.New("arbitrary error"),
			nil,
			status.Error(codes.Internal, "database error"),
		},
		{
			"happy_flow",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().IncomingAccessRequestReject(gomock.Any(), "Jane Doe", "nlxctl", "test-organization", "test-service")
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: 1,
			},
			&database.IncomingAccessRequest{
				Organization: database.IncomingAccessRequestOrganization{
					Name: "test-organization",
				},
				Service: &database.Service{
					Name: "other-service",
				},
			},
			nil,
			true,
			nil,
			&emptypb.Empty{},
			nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			service, _, mocks := newService(t)

			test.auditLog(*mocks.al)

			mocks.db.EXPECT().GetIncomingAccessRequest(test.ctx, uint(test.request.AccessRequestID)).Return(test.accessRequest, test.accessRequestErr)

			if test.expectUpdateCall {
				mocks.db.EXPECT().UpdateIncomingAccessRequestState(test.ctx, test.accessRequest.ID, database.IncomingAccessRequestRejected).Return(test.updateErr)
			}

			actual, err := service.RejectIncomingAccessRequest(test.ctx, test.request)

			assert.Equal(t, test.response, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestExternalRequestAccess(t *testing.T) {
	tests := map[string]struct {
		setup   func(*testing.T, *mock_database.MockConfigDatabase) context.Context
		want    *external.RequestAccessResponse
		wantErr error
	}{
		"when_peer_context_is_missing": {
			setup: func(*testing.T, *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_the_service_does_not_exist": {
			setup: func(_ *testing.T, db *mock_database.MockConfigDatabase) context.Context {
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
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase) context.Context {
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
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(nil, errors.New("error"))

				return ctx
			},
			wantErr: status.Error(codes.Internal, "failed to create access request"),
		},
		"returns_error_when_a_active_access_request_already_exists": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase) context.Context {
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
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(nil, database.ErrActiveAccessRequest)

				return ctx
			},
			wantErr: status.Error(codes.AlreadyExists, "an active access request already exists"),
		},
		"happy_flow": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				publicKeyPEM, err := certBundle.PublicKeyPEM()
				require.NoError(t, err)

				db.
					EXPECT().
					GetService(ctx, "service").
					Return(&database.Service{
						ID:   1,
						Name: "Service",
					}, nil)

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, &database.IncomingAccessRequest{
						ServiceID: 1,
						Organization: database.IncomingAccessRequestOrganization{
							Name:         certBundle.Certificate().Subject.Organization[0],
							SerialNumber: certBundle.Certificate().Subject.SerialNumber,
						},
						State:                database.IncomingAccessRequestReceived,
						PublicKeyPEM:         publicKeyPEM,
						PublicKeyFingerprint: certBundle.PublicKeyFingerprint(),
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
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			ctx := tt.setup(t, mocks.db)

			result, err := service.RequestAccess(ctx, &external.RequestAccessRequest{
				ServiceName: "service",
			})

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.Equal(t, tt.want, result)
			}
		})
	}
}

//nolint funlen: this is a test
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
					GetLatestIncomingAccessRequest(ctx, certBundle.Certificate().Subject.SerialNumber, "service").
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
				ServiceName: "service",
			})

			assert.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
