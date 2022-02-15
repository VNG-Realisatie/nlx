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

	validOutgoingOrderRequest := func() api.OutgoingOrderRequest {
		return api.OutgoingOrderRequest{
			Reference:    "a-reference",
			Description:  "a-description",
			Delegatee:    "00000000000000000001",
			PublicKeyPEM: testPublicKeyPEM,
			ValidFrom:    timestamppb.New(validFrom),
			ValidUntil:   timestamppb.New(validUntil),
		}
	}

	tests := map[string]struct {
		request *api.OutgoingOrderRequest
		setup   func(serviceMocks)
		wantErr error
	}{
		"when_providing_an_empty_reference": {
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Reference = ""
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Reference: cannot be blank."),
		},
		"when_providing_an_empty_delegatee": {
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Delegatee = ""
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Delegatee: organization serial number must be in a valid format: cannot be empty."),
		},
		"when_providing_an_empty_description": {
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Description = ""
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Description: cannot be blank."),
		},
		"when_providing_a_description_which_is_too_long": {
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.Description = "description-with-length-of-101-chars-abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcd"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Description: the length must be between 1 and 100."),
		},
		"when_providing_an_end_date_which_is_before_the_start_date": {
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.ValidFrom = timestamppb.New(time.Now().Add(time.Millisecond))
				request.ValidUntil = timestamppb.New(time.Now())
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: ValidUntil: order can not expire before the start date."),
		},
		"when_providing_an_empty_list_of_access_proof_ids": {
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.AccessProofIds = []uint64{}
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: Services: cannot be blank."),
		},
		"when_providing_an_invalid_public_key": {
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				request.PublicKeyPEM = "invalid"
				return &request
			}(),
			wantErr: status.Error(codes.InvalidArgument, "invalid outgoing order: PublicKeyPEM: expect public key as pem."),
		},
		"when_updating_the_order_fails": {
			wantErr: status.Error(codes.Internal, "failed to update outgoing order"),
			setup: func(mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingUpdate(gomock.Any(), "Jane Doe", "nlxctl", "00000000000000000001", "a-reference", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "10000000000000000001",
								Name:         "a-organization",
							},
							Service: "a-service",
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
						Services: []database.OutgoingOrderService{
							{
								Organization: database.OutgoingOrderServiceOrganization{
									SerialNumber: "10000000000000000001",
									Name:         "a-organization",
								},
								Service: "a-service",
							},
						},
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingOrder(gomock.Any(), &database.OutgoingOrder{
						Reference:    "a-reference",
						Description:  "a-description",
						Delegatee:    "00000000000000000001",
						PublicKeyPEM: testPublicKeyPEM,
						ValidFrom:    validFrom,
						ValidUntil:   validUntil,
						Services: []database.OutgoingOrderService{
							{
								Organization: database.OutgoingOrderServiceOrganization{
									SerialNumber: "10000000000000000001",
									Name:         "a-organization",
								},
								Service: "a-service",
							},
						},
					}).
					Return(errors.New("arbitrary error"))
			},
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
		"when_the_order_is_not_found": {
			wantErr: status.Error(codes.NotFound, "could not find outgoing order in management database"),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), gomock.Any()).
					Return(nil, database.ErrNotFound)
			},
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
		"when_updating_audit_log_fails": {
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
						Services: []database.OutgoingOrderService{
							{
								Organization: database.OutgoingOrderServiceOrganization{
									SerialNumber: "10000000000000000001",
									Name:         "a-organization",
								},
								Service: "a-service",
							},
						},
					}, nil)

				mocks.al.
					EXPECT().
					OrderOutgoingUpdate(gomock.Any(), "Jane Doe", "nlxctl", "00000000000000000001", "a-reference", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "10000000000000000001",
								Name:         "a-organization",
							},
							Service: "a-service",
						},
					}).
					Return(errors.New("arbitrary error"))
			},
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
			wantErr: status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					UpdateOutgoingOrder(gomock.Any(), &database.OutgoingOrder{
						Reference:    "a-reference",
						Description:  "a-description",
						Delegatee:    "00000000000000000001",
						PublicKeyPEM: testPublicKeyPEM,
						ValidFrom:    validFrom,
						ValidUntil:   validUntil,
						Services: []database.OutgoingOrderService{
							{
								Organization: database.OutgoingOrderServiceOrganization{
									SerialNumber: "10000000000000000001",
									Name:         "a-organization",
								},
								Service: "a-service",
							},
						},
					}).
					Return(nil)

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
						Services: []database.OutgoingOrderService{
							{
								Organization: database.OutgoingOrderServiceOrganization{
									SerialNumber: "10000000000000000001",
									Name:         "a-organization",
								},
								Service: "a-service",
							},
						},
					}, nil)

				mocks.al.
					EXPECT().
					OrderOutgoingUpdate(gomock.Any(), "Jane Doe", "nlxctl", "00000000000000000001", "a-reference", []auditlog.RecordService{
						{
							Organization: auditlog.RecordServiceOrganization{
								SerialNumber: "10000000000000000001",
								Name:         "a-organization",
							},
							Service: "a-service",
						},
					})
			},
			request: func() *api.OutgoingOrderRequest {
				request := validOutgoingOrderRequest()
				return &request
			}(),
		},
	}

	//nolint:dupl // this is a test method
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			}))

			response, err := service.UpdateOutgoingOrder(ctx, tt.request)

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
