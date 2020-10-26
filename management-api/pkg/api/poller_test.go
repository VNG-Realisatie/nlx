package api

import (
	"context"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/directory"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
)

//nolint:funlen // covers all test-cases of access-proofs
func TestSyncAccessProof(t *testing.T) {
	tests := map[string]struct {
		setupMocks func(context.Context, *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client)
		request    *database.OutgoingAccessRequest
		wantErr    bool
	}{
		"returns_an_error_when_get_organization_inway_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				directoryClient := mock_directory.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(nil, errors.New("random error"))

				return nil, directoryClient, nil
			},
		},

		"returns_an_error_when_inway_address_is_invalid": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				directoryClient := mock_directory.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "invalid",
					}, nil)

				return nil, directoryClient, nil
			},
		},

		"returns_an_error_when_external_get_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
					ServiceName:      "service",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				directoryClient := mock_directory.NewMockClient(ctrl)
				managementClient := mock_management.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "localhost:8000",
					}, nil)

				managementClient.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(nil, errors.New("random error"))

				return nil, directoryClient, managementClient
			},
		},

		"returns_an_error_when_parsing_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
					ServiceName:      "service",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				directoryClient := mock_directory.NewMockClient(ctrl)
				managementClient := mock_management.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "localhost:8000",
					}, nil)

				managementClient.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{}, nil)

				return nil, directoryClient, managementClient
			},
		},

		"returns_an_error_when_database_getting_latest_access_proof_for_service_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
					ServiceName:      "service",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				db := mock_database.NewMockConfigDatabase(ctrl)
				directoryClient := mock_directory.NewMockClient(ctrl)
				managementClient := mock_management.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "localhost:8000",
					}, nil)

				managementClient.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						CreatedAt: types.TimestampNow(),
						RevokedAt: nil,
					}, nil)

				db.
					EXPECT().
					GetLatestAccessProofForService(ctx, "organization-a", "service").
					Return(nil, errors.New("random error"))

				return db, directoryClient, managementClient
			},
		},

		"returns_an_error_when_database_create_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
					ServiceName:      "service",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				ts := types.TimestampNow()
				t, _ := types.TimestampFromProto(ts)

				db := mock_database.NewMockConfigDatabase(ctrl)
				directoryClient := mock_directory.NewMockClient(ctrl)
				managementClient := mock_management.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "localhost:8000",
					}, nil)

				managementClient.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						CreatedAt:        ts,
						OrganizationName: "organization-a",
						ServiceName:      "service",
						RevokedAt:        nil,
					}, nil)

				db.
					EXPECT().
					GetLatestAccessProofForService(ctx, "organization-a", "service").
					Return(nil, database.ErrNotFound)

				db.
					EXPECT().
					CreateAccessProof(ctx, &database.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        t,
					}).
					Return(nil, errors.New("random error"))

				return db, directoryClient, managementClient
			},
		},

		"returns_an_error_when_database_revoke_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
					ServiceName:      "service",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				ts := types.TimestampNow()
				t, _ := types.TimestampFromProto(ts)

				db := mock_database.NewMockConfigDatabase(ctrl)
				directoryClient := mock_directory.NewMockClient(ctrl)
				managementClient := mock_management.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "localhost:8000",
					}, nil)

				managementClient.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        ts,
						RevokedAt:        ts,
					}, nil)

				db.
					EXPECT().
					GetLatestAccessProofForService(ctx, "organization-a", "service").
					Return(&database.AccessProof{
						ID:               "1",
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        t,
						RevokedAt:        time.Time{},
					}, nil)

				db.
					EXPECT().
					RevokeAccessProof(
						ctx,
						"organization-a",
						"service",
						"1",
						t,
					).
					Return(nil, errors.New("random error"))

				return db, directoryClient, managementClient
			},
		},

		"successfully_revokes_an_access_grant_when_its_revoked": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
					ServiceName:      "service",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				ts := types.TimestampNow()
				t, _ := types.TimestampFromProto(ts)

				db := mock_database.NewMockConfigDatabase(ctrl)
				directoryClient := mock_directory.NewMockClient(ctrl)
				managementClient := mock_management.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "localhost:8000",
					}, nil)

				managementClient.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        ts,
						RevokedAt:        ts,
					}, nil)

				db.
					EXPECT().
					GetLatestAccessProofForService(ctx, "organization-a", "service").
					Return(&database.AccessProof{
						ID:               "1",
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        t,
						RevokedAt:        time.Time{},
					}, nil)

				db.
					EXPECT().
					RevokeAccessProof(
						ctx,
						"organization-a",
						"service",
						"1",
						t,
					).
					Return(nil, nil)

				return db, directoryClient, managementClient
			},
		},

		"successfully_creates_an_access_proof_when_its_found": {
			wantErr: false,
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "organization-a",
					ServiceName:      "service",
				},
			},
			setupMocks: func(ctx context.Context, ctrl *gomock.Controller) (database.ConfigDatabase, directory.Client, management.Client) {
				ts := types.TimestampNow()
				t, _ := types.TimestampFromProto(ts)

				db := mock_database.NewMockConfigDatabase(ctrl)
				directoryClient := mock_directory.NewMockClient(ctrl)
				managementClient := mock_management.NewMockClient(ctrl)

				directoryClient.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "localhost:8000",
					}, nil)

				managementClient.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						CreatedAt:        ts,
						OrganizationName: "organization-a",
						ServiceName:      "service",
						RevokedAt:        nil,
					}, nil)

				db.
					EXPECT().
					GetLatestAccessProofForService(ctx, "organization-a", "service").
					Return(nil, database.ErrNotFound)

				db.
					EXPECT().
					CreateAccessProof(ctx, &database.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        t,
					}).
					Return(nil, nil)

				return db, directoryClient, managementClient
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, directoryClient, managementClient := tt.setupMocks(ctx, ctrl)

			poller := &accessProofPoller{
				logger:          zap.NewNop(),
				orgCert:         nil,
				configDatabase:  db,
				directoryClient: directoryClient,
				createManagementClientFunc: func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error) {
					return managementClient, nil
				},
			}

			err := poller.syncAccessProof(ctx, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
