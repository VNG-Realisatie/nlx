// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/diagnostics"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

var testPublicKeyPEM = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEArN5xGkM73tJsCpKny59e5lXNRY+eT0sbWyEGsR1qIPRKmLSiRHl3
xMsovn5mo6jN3eeK/Q4wKd6Ae5XGzP63pTG6U5KVVB74eQxSFfV3UEOrDaJ78X5m
BZO+Ku21V2QFr44tvMh5IZDX3RbMB/4Kad6sapmSF00HWrqTVMkrEsZ98DTb5nwG
Lh3kISnct4tLyVSpsl9s1rtkSgGUcs1TIvWxS2D2mOsSL1HRdUNcFQmzchbfG87k
XPvicoOISAZDJKDqWp3iuH0gJpQ+XMBfmcD90I7Z/cRQjWP3P93B3V06cJkd00cE
IRcIQqF8N+lE01H88Fi+wePhZRy92NP54wIDAQAB
-----END RSA PUBLIC KEY-----
`

func newService(t *testing.T) (s *server.ManagementService, db *mock_database.MockConfigDatabase, auditLogger *mock_auditlog.MockLogger, bundle *common_tls.CertificateBundle) {
	logger := zaptest.Logger(t)
	proc := process.NewProcess(logger)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	db = mock_database.NewMockConfigDatabase(ctrl)
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")
	bundle, err := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	assert.NoError(t, err)

	auditLogger = mock_auditlog.NewMockLogger(ctrl)

	s = server.NewManagementService(logger, proc, mock_directory.NewMockClient(ctrl), bundle, db, nil, auditLogger)

	return s, db, auditLogger, bundle
}

func createTimestamp(ti time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: ti.Unix(),
		Nanos:   int32(ti.Nanosecond()),
	}
}

func TestCreateAccessRequest(t *testing.T) {
	tests := map[string]struct {
		auditLog    func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger
		ctx         context.Context
		req         *api.CreateAccessRequestRequest
		ar          *database.OutgoingAccessRequest
		returnReq   *database.OutgoingAccessRequest
		returnErr   error
		expectedReq *api.OutgoingAccessRequest
		expectedErr error
	}{
		"without_an_active_access_request": {
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().OutgoingAccessRequestCreate(gomock.Any(), "Jane Doe", "nlxctl", "test-organization", "test-service")
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.OutgoingAccessRequest{
				OrganizationName:     "test-organization",
				ServiceName:          "test-service",
				PublicKeyPEM:         testPublicKeyPEM,
				PublicKeyFingerprint: "60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=",
				State:                database.OutgoingAccessRequestCreated,
			},
			&database.OutgoingAccessRequest{
				ID:                   1,
				OrganizationName:     "test-organization",
				ServiceName:          "test-service",
				PublicKeyPEM:         testPublicKeyPEM,
				PublicKeyFingerprint: "60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=",
				State:                database.OutgoingAccessRequestCreated,
				CreatedAt:            time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
				UpdatedAt:            time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
			},
			nil,
			&api.OutgoingAccessRequest{
				Id:               1,
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
				UpdatedAt:        createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
			},
			nil,
		},
		"with_an_active_access_request": {
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().OutgoingAccessRequestCreate(gomock.Any(), "Jane Doe", "nlxctl", "test-organization", "test-service")
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.OutgoingAccessRequest{
				OrganizationName:     "test-organization",
				ServiceName:          "test-service",
				PublicKeyPEM:         testPublicKeyPEM,
				PublicKeyFingerprint: "60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=",
				State:                database.OutgoingAccessRequestCreated,
			},
			nil,
			database.ErrActiveAccessRequest,
			nil,
			status.New(codes.AlreadyExists, "there is already an active access request").Err(),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db, auditLogger, _ := newService(t)

			tt.auditLog(*auditLogger)

			db.EXPECT().CreateOutgoingAccessRequest(tt.ctx, tt.ar).
				Return(tt.returnReq, tt.returnErr)
			actual, err := service.CreateAccessRequest(tt.ctx, tt.req)

			assert.Equal(t, tt.expectedReq, actual)
			assert.Equal(t, tt.expectedErr, err)
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
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  1,
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
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  1,
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
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  1,
			},
			&database.OutgoingAccessRequest{

				ID:               1,
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            database.OutgoingAccessRequestCreated,
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
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  1,
			},
			&database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            database.OutgoingAccessRequestCreated,
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(nil)
			},
			&api.OutgoingAccessRequest{
				Id:               1,
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        createTimestamp(time.Time{}),
				UpdatedAt:        createTimestamp(time.Time{}),
			},
			nil,
		},
		{
			"failed_state",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  1,
			},
			&database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            database.OutgoingAccessRequestFailed,
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(nil)
			},
			&api.OutgoingAccessRequest{
				Id:               1,
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        createTimestamp(time.Time{}),
				UpdatedAt:        createTimestamp(time.Time{}),
			},
			nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			service, db, _, _ := newService(t)
			ctx := context.Background()

			db.EXPECT().GetOutgoingAccessRequest(ctx, uint(test.request.AccessRequestID)).
				Return(test.accessRequest, test.accessRequestErr)

			updateMock := db.EXPECT().UpdateOutgoingAccessRequestState(ctx, uint(test.request.AccessRequestID), database.OutgoingAccessRequestCreated, uint(0), nil).
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
				OrganizationName: "test-organization",
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
				OrganizationName: "test-organization",
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
			service, db, auditLogger, _ := newService(t)
			test.auditLog(*auditLogger)

			db.EXPECT().GetIncomingAccessRequest(test.ctx, uint(test.request.AccessRequestID)).Return(test.accessRequest, test.accessRequestErr)

			if test.response != nil {
				db.EXPECT().CreateAccessGrant(test.ctx, test.accessRequest)
			}

			if test.expectUpdateCall {
				db.EXPECT().UpdateIncomingAccessRequestState(test.ctx, test.accessRequest.ID, database.IncomingAccessRequestApproved).Return(test.updateErr)
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
				OrganizationName: "test-organization",
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
				OrganizationName: "test-organization",
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
			service, db, auditLogger, _ := newService(t)

			test.auditLog(*auditLogger)

			db.EXPECT().GetIncomingAccessRequest(test.ctx, uint(test.request.AccessRequestID)).Return(test.accessRequest, test.accessRequestErr)

			if test.expectUpdateCall {
				db.EXPECT().UpdateIncomingAccessRequestState(test.ctx, test.accessRequest.ID, database.IncomingAccessRequestRejected).Return(test.updateErr)
			}

			actual, err := service.RejectIncomingAccessRequest(test.ctx, test.request)

			assert.Equal(t, test.response, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

func setProxyMetadata(ctx context.Context) context.Context {
	md := metadata.Pairs(
		"nlx-organization", "organization-a",
		"nlx-public-key-der", "ZHVtbXktcHVibGljLWtleQo=",
		"nlx-public-key-fingerprint", "1655A0AB68576280",
	)

	return metadata.NewIncomingContext(ctx, md)
}

func TestExternalRequestAccess(t *testing.T) {
	tests := map[string]struct {
		wantCode codes.Code
		setup    func(*mock_database.MockConfigDatabase) context.Context
	}{
		"errors_when_peer_context_is_missing": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
		},
		"service does not exist": {
			wantCode: codes.NotFound,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.EXPECT().GetService(ctx, "service").Return(nil, database.ErrNotFound)

				return ctx
			},
		},

		"returns_error_when_creating_access_request_errors": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())
				db.EXPECT().GetService(ctx, "service").Return(&database.Service{
					ID:   1,
					Name: "Service",
				}, nil)

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(nil, errors.New("error"))

				return ctx
			},
		},

		"returns_empty_when_creating_the_access_request_succeeds": {
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.EXPECT().GetService(ctx, "service").Return(&database.Service{
					ID:   1,
					Name: "Service",
				}, nil)

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(&database.IncomingAccessRequest{}, nil)

				return ctx
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db, _, _ := newService(t)
			ctx := tt.setup(db)

			_, err := service.RequestAccess(ctx, &external.RequestAccessRequest{
				ServiceName: "service",
			})

			if tt.wantCode > 0 {
				assert.Error(t, err)

				st, ok := status.FromError(err)

				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExternalGetAccessRequestState(t *testing.T) {
	tests := map[string]struct {
		want     api.AccessRequestState
		wantCode codes.Code
		setup    func(*mock_database.MockConfigDatabase) context.Context
	}{
		"errors_when_peer_context_is_missing": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
		},

		"returns_error_when_retrieving_state_errors": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "organization-a", "service").
					Return(nil, errors.New("error"))

				return ctx
			},
		},

		"returns_the_right_state_when_the_request_exists": {
			want: api.AccessRequestState_RECEIVED,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "organization-a", "service").
					Return(&database.IncomingAccessRequest{
						State: database.IncomingAccessRequestReceived,
					}, nil)

				return ctx
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db, _, _ := newService(t)
			ctx := tt.setup(db)

			response, err := service.GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
				ServiceName: "service",
			})

			if tt.wantCode > 0 {
				assert.Error(t, err)

				st, ok := status.FromError(err)

				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
			} else {
				assert.NoError(t, err)

				if assert.NotNil(t, response) {
					assert.Equal(t, tt.want, response.State)
				}
			}
		})
	}
}
