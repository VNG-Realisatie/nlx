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
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // this is a test method
func TestCreateOutgoingOrder(t *testing.T) {
	validFrom := time.Now().UTC()
	validUntil := time.Now().Add(1 * time.Hour).UTC()

	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	testPublicKeyPEM, err := certBundle.PublicKeyPEM()
	require.NoError(t, err)

	validOutgoingOrderRequest := func() api.CreateOutgoingOrderRequest {
		return api.CreateOutgoingOrderRequest{
			Reference:      "a-reference",
			Description:    "a-description",
			Delegatee:      "00000000000000000001",
			PublicKeyPEM:   testPublicKeyPEM,
			ValidFrom:      timestamppb.New(validFrom),
			ValidUntil:     timestamppb.New(validUntil),
			AccessProofIds: []uint64{1, 2, 3},
		}
	}

	tests := map[string]struct {
		request *api.CreateOutgoingOrderRequest
		setup   func(serviceMocks)
		ctx     context.Context
		wantErr error
	}{
		"missing_required_permission": {
			setup:   func(mocks serviceMocks) {},
			ctx:     testCreateUserWithoutPermissionsContext(),
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.outgoing_order.create\" to execute this request").Err(),
		},
		"when_providing_an_empty_reference": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Reference = ""
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Reference: cannot be blank."),
		},
		"when_providing_a_reference_which_is_too_long": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Reference = "reference-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdea"
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Reference: the length must be between 1 and 100."),
		},
		"when_providing_an_empty_description": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Description = ""
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Description: cannot be blank."),
		},
		"when_providing_a_description_which_is_too_long": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Description = "description-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcd"
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Description: the length must be between 1 and 100."),
		},
		"when_providing_an_invalid_delegatee": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Delegatee = "00000000000000000000000001_too_long"
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Delegatee: organization serial number must be in a valid format: too long, max 20 bytes."),
		},
		"when_providing_an_end_date_which_is_before_the_start_date": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.ValidFrom = timestamppb.New(time.Now().Add(time.Millisecond))
				request.ValidUntil = timestamppb.New(time.Now())
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: ValidUntil: order can not expire before the start date."),
		},
		"when_providing_an_empty_list_of_access_proofs": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.AccessProofIds = []uint64{}
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: AccessProofIds: cannot be blank."),
		},
		"when_providing_an_invalid_public_key": {
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.PublicKeyPEM = "invalid"
				return &request
			}(),
			ctx:     testCreateAdminUserContext(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: PublicKeyPEM: expect public key as pem."),
		},
		"when_a_record_with_the_same_reference_for_the_same_organization_already_exists": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetAccessProofs(gomock.Any(), []uint64{1, 2, 3}).
					Return([]*database.AccessProof{
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "organization-a",
								},
								ServiceName: "service-a",
							},
						},
					}, nil)

				mocks.al.
					EXPECT().
					OrderCreate(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "organization-a",
							},
							Service: "service-a",
						},
					})

				mocks.db.
					EXPECT().
					CreateOutgoingOrder(gomock.Any(), &database.CreateOutgoingOrder{
						Reference:      "a-reference",
						Description:    "a-description",
						PublicKeyPEM:   testPublicKeyPEM,
						Delegatee:      "00000000000000000001",
						ValidFrom:      validFrom,
						ValidUntil:     validUntil,
						AccessProofIds: []uint64{1, 2, 3},
					}).
					Return(database.ErrDuplicateOutgoingOrder)
			},
			ctx: testCreateAdminUserContext(),
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "an order with reference a-reference for 00000000000000000001 already exist"),
		},
		"when_creating_the_order_fails": {
			wantErr: status.Error(codes.Internal, "failed to create outgoing order"),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetAccessProofs(gomock.Any(), []uint64{1, 2, 3}).
					Return([]*database.AccessProof{
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "organization-a",
								},
								ServiceName: "service-a",
							},
						},
					}, nil)

				mocks.al.
					EXPECT().
					OrderCreate(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "organization-a",
							},
							Service: "service-a",
						},
					})

				mocks.db.
					EXPECT().
					CreateOutgoingOrder(gomock.Any(), &database.CreateOutgoingOrder{
						Reference:      "a-reference",
						Description:    "a-description",
						PublicKeyPEM:   testPublicKeyPEM,
						Delegatee:      "00000000000000000001",
						ValidFrom:      validFrom,
						ValidUntil:     validUntil,
						AccessProofIds: []uint64{1, 2, 3},
					}).
					Return(errors.New("arbitrary error"))
			},
			ctx: testCreateAdminUserContext(),
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
		"when_get_access_proofs_fails": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetAccessProofs(gomock.Any(), []uint64{1, 2, 3}).
					Return(nil, errors.New("arbitrary error"))
			},
			ctx: testCreateAdminUserContext(),
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
			wantErr: status.Error(codes.Internal, "could not retrieve access proofs"),
		},
		"when_creating_audit_log_fails": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetAccessProofs(gomock.Any(), []uint64{1, 2, 3}).
					Return([]*database.AccessProof{
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "organization-a",
								},
								ServiceName: "service-a",
							},
						},
					}, nil)

				mocks.al.
					EXPECT().
					OrderCreate(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "organization-a",
							},
							Service: "service-a",
						},
					}).
					Return(errors.New("arbitrary error"))
			},
			ctx: testCreateAdminUserContext(),
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
			wantErr: status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"when_order_has_duplicate_services": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetAccessProofs(gomock.Any(), []uint64{1, 2, 3}).
					Return([]*database.AccessProof{
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "organization-a",
								},
								ServiceName: "service-a",
							},
						},
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "organization-a",
								},
								ServiceName: "service-a",
							},
						},
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000002",
									Name:         "organization-b",
								},
								ServiceName: "service-b",
							},
						},
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000002",
									Name:         "organization-b",
								},
								ServiceName: "service-c",
							},
						},
					}, nil)
			},
			ctx: testCreateAdminUserContext(),
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
			wantErr: status.Error(codes.Internal, "cannot create order with duplicate services"),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetAccessProofs(gomock.Any(), []uint64{1, 2, 3}).
					Return([]*database.AccessProof{
						{
							OutgoingAccessRequest: &database.OutgoingAccessRequest{
								Organization: database.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "organization-a",
								},
								ServiceName: "service-a",
							},
						},
					}, nil)

				mocks.al.
					EXPECT().
					OrderCreate(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "organization-a",
							},
							Service: "service-a",
						},
					})

				mocks.db.
					EXPECT().
					CreateOutgoingOrder(gomock.Any(), &database.CreateOutgoingOrder{
						Reference:      "a-reference",
						Description:    "a-description",
						PublicKeyPEM:   testPublicKeyPEM,
						Delegatee:      "00000000000000000001",
						ValidFrom:      validFrom,
						ValidUntil:     validUntil,
						AccessProofIds: []uint64{1, 2, 3},
					}).
					Return(nil)
			},
			ctx: testCreateAdminUserContext(),
			request: func() *api.CreateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			response, err := service.CreateOutgoingOrder(tt.ctx, tt.request)

			if tt.wantErr == nil {
				assert.IsType(t, &emptypb.Empty{}, response)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, response)
			}
		})
	}
}
