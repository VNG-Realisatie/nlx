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
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // this is a test method
func TestUpdateOutgoingOrder(t *testing.T) {
	validFrom := time.Now().UTC()
	validUntil := time.Now().Add(1 * time.Hour).UTC()

	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	testPublicKeyPEM, err := certBundle.PublicKeyPEM()
	require.NoError(t, err)

	validOutgoingOrderRequest := func() api.UpdateOutgoingOrderRequest {
		return api.UpdateOutgoingOrderRequest{
			Reference:      "a-reference",
			Description:    "a-description",
			Delegatee:      "00000000000000000001",
			PublicKeyPem:   testPublicKeyPEM,
			ValidFrom:      timestamppb.New(validFrom),
			ValidUntil:     timestamppb.New(validUntil),
			AccessProofIds: []uint64{1},
		}
	}

	tests := map[string]struct {
		request *api.UpdateOutgoingOrderRequest
		setup   func(serviceMocks)
		ctx     context.Context
		wantErr error
	}{
		"missing_required_permission": {
			ctx: testCreateUserWithoutPermissionsContext(),
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Reference = ""
				return &request
			}(),
			setup:   func(mocks serviceMocks) {},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.outgoing_order.update\" to execute this request").Err(),
		},
		"when_providing_an_empty_reference": {
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Reference = ""
				return &request
			}(),
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID: 1,
					}, nil)
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid update outgoing order: Reference: cannot be blank."),
		},
		"when_providing_an_empty_description": {
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Description = ""
				return &request
			}(),
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID: 1,
					}, nil)
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid update outgoing order: Description: cannot be blank."),
		},
		"when_providing_a_description_which_is_too_long": {
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Description = "description-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcd"
				return &request
			}(),
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID: 1,
					}, nil)
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid update outgoing order: Description: the length must be between 1 and 100."),
		},
		"when_providing_an_end_date_which_is_before_the_start_date": {
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.ValidFrom = timestamppb.New(time.Now().Add(time.Millisecond))
				request.ValidUntil = timestamppb.New(time.Now())
				return &request
			}(),
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID: 1,
					}, nil)
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid update outgoing order: ValidUntil: order can not expire before the start date."),
		},
		"when_providing_an_empty_list_of_access_proof_ids": {
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.AccessProofIds = []uint64{}
				return &request
			}(),
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID: 1,
					}, nil)
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid update outgoing order: AccessProofIds: cannot be blank."),
		},
		"when_providing_an_invalid_public_key": {
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.PublicKeyPem = "invalid"
				return &request
			}(),
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID: 1,
					}, nil)
			},
			wantErr: status.Error(codes.InvalidArgument, "invalid update outgoing order: PublicKeyPEM: expect public key as pem."),
		},
		"when_updating_the_order_fails": {
			wantErr: status.Error(codes.Internal, "failed to update outgoing order"),
			ctx:     testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingUpdate(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", "a-reference", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "10000000000000000001",
								Name:         "organization-a",
							},
							Service: "service-a",
						},
					})

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID:           0,
						Reference:    "a-reference",
						Description:  "a-description",
						Delegatee:    "00000000000000000001",
						PublicKeyPEM: testPublicKeyPEM,
						ValidFrom:    validFrom,
						ValidUntil:   validUntil,
					}, nil)

				mocks.db.EXPECT().GetAccessProofs(gomock.Any(), []uint64{1}).Return([]*database.AccessProof{
					{
						ID: 1,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{Organization: database.Organization{
							SerialNumber: "10000000000000000001",
							Name:         "organization-a",
						},
							ServiceName: "service-a",
						},
					},
				}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingOrder(gomock.Any(), &database.UpdateOutgoingOrder{
						Reference:      "a-reference",
						Description:    "a-description",
						PublicKeyPEM:   testPublicKeyPEM,
						ValidFrom:      validFrom,
						ValidUntil:     validUntil,
						AccessProofIds: []uint64{1},
					}).
					Return(errors.New("arbitrary error"))
			},
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
		"when_the_order_is_not_found": {
			wantErr: status.Error(codes.NotFound, "could not find outgoing order in management database"),
			ctx:     testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(nil, database.ErrNotFound)
			},
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
		"when_updating_audit_log_fails": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID:           0,
						Reference:    "a-reference",
						Description:  "a-description",
						Delegatee:    "00000000000000000001",
						PublicKeyPEM: testPublicKeyPEM,
						ValidFrom:    validFrom,
						ValidUntil:   validUntil,
					}, nil)

				mocks.db.EXPECT().GetAccessProofs(gomock.Any(), []uint64{1}).Return([]*database.AccessProof{
					{
						ID: 1,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{Organization: database.Organization{
							SerialNumber: "10000000000000000001",
							Name:         "organization-a",
						},
							ServiceName: "service-a",
						},
					},
				}, nil)

				mocks.al.
					EXPECT().
					OrderOutgoingUpdate(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", "a-reference", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "10000000000000000001",
								Name:         "organization-a",
							},
							Service: "service-a",
						},
					}).
					Return(errors.New("arbitrary error"))
			},
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
			wantErr: status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingUpdate(gomock.Any(), "admin@example.com", "nlxctl", "00000000000000000001", "a-reference", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "10000000000000000001",
								Name:         "organization-a",
							},
							Service: "service-a",
						},
					})

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(&database.OutgoingOrder{
						ID:           0,
						Reference:    "a-reference",
						Description:  "a-description",
						Delegatee:    "00000000000000000001",
						PublicKeyPEM: testPublicKeyPEM,
						ValidFrom:    validFrom,
						ValidUntil:   validUntil,
					}, nil)

				mocks.db.EXPECT().GetAccessProofs(gomock.Any(), []uint64{1}).Return([]*database.AccessProof{
					{
						ID: 1,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{Organization: database.Organization{
							SerialNumber: "10000000000000000001",
							Name:         "organization-a",
						},
							ServiceName: "service-a",
						},
					},
				}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingOrder(gomock.Any(), &database.UpdateOutgoingOrder{
						Reference:      "a-reference",
						Description:    "a-description",
						PublicKeyPEM:   testPublicKeyPEM,
						ValidFrom:      validFrom,
						ValidUntil:     validUntil,
						AccessProofIds: []uint64{1},
					}).
					Return(nil)
			},
			request: func() *api.UpdateOutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
	}

	//nolint:dupl // this is a test method
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t, nil)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			response, err := service.UpdateOutgoingOrder(tt.ctx, tt.request)

			if tt.wantErr == nil {
				assert.IsType(t, &api.UpdateOutgoingOrderResponse{}, response)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, response)
			}
		})
	}
}
