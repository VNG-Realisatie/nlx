// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // code is not actually duplicated, the linter has lost it's mind
package server_test

import (
	"context"
	"database/sql"
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

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // this is a test method
func TestListOutgoingOrders(t *testing.T) {
	validFrom := time.Now()
	validUntil := time.Now().Add(time.Hour)

	accessProofCreatedAt := time.Now()
	accessRequestCreatedAt := time.Now()

	tests := map[string]struct {
		setup        func(serviceMocks)
		wantResponse *api.ListOutgoingOrdersResponse
		wantErr      error
	}{
		"when_retrieval_of_orders_from_database_fails": {
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().ListOutgoingOrders(gomock.Any()).Return(nil, errors.New("arbitrary error"))
			},
			wantErr: status.Error(codes.Internal, "failed to retrieve outgoing orders"),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListOutgoingOrders(gomock.Any()).
					Return([]*database.OutgoingOrder{
						{
							ID:           1,
							Reference:    "reference",
							Description:  "description",
							Delegatee:    "10000000000000000001",
							PublicKeyPEM: "public_key",
							ValidFrom:    validFrom,
							ValidUntil:   validUntil,
							RevokedAt:    sql.NullTime{},
							OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
								{
									OutgoingOrderID: 1,
									AccessProofID:   10,
									AccessProof: &database.AccessProof{
										ID:        10,
										CreatedAt: accessProofCreatedAt,
										OutgoingAccessRequest: &database.OutgoingAccessRequest{
											ID:                   100,
											Organization:         database.Organization{
												Name: ""
											},
											ServiceName:          "",
											ReferenceID:          0,
											State:                "",
											LockID:               nil,
											LockExpiresAt:        sql.NullTime{},
											PublicKeyFingerprint: "",
											ErrorCode:            0,
											ErrorCause:           "",
											ErrorStackTrace:      nil,
											CreatedAt:            time.Time{},
											UpdatedAt:            time.Time{},
											SynchronizeAt:        time.Time{},
										},
									},
								},
							},
						},
					}, nil)
			},
			wantResponse: &api.ListOutgoingOrdersResponse{
				Orders: []*api.OutgoingOrder{
					{
						Reference:    "reference",
						Description:  "description",
						Delegatee:    "10000000000000000001",
						PublicKeyPem: "public_key",
						ValidFrom:    timestamppb.New(validFrom),
						ValidUntil:   timestamppb.New(validUntil),
						AccessProofs: []*api.AccessProof{{
							Id: 10,
						}},
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

			response, err := service.ListOutgoingOrders(context.Background(), &emptypb.Empty{})

			if tt.wantErr == nil {
				assert.Equal(t, tt.wantResponse, response)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, response)
			}
		})
	}
}

//nolint:funlen // this is a test method
func TestListIncomingOrders(t *testing.T) {
	validFrom := time.Now()
	validUntil := time.Now().Add(time.Hour)
	revokedAt := time.Now()

	tests := map[string]struct {
		setup        func(serviceMocks)
		wantResponse *api.ListIncomingOrdersResponse
		wantErr      error
	}{
		"when_retrieval_of_orders_from_database_fails": {
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().ListIncomingOrders(gomock.Any()).Return(nil, errors.New("arbitrary error"))
			},
			wantErr: status.Error(codes.Internal, "failed to retrieve received orders"),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				services := []domain.IncomingOrderService{
					domain.NewIncomingOrderService("service-a", "00000000000000000001", "organization-a"),
				}
				model, _ := domain.NewIncomingOrder(
					&domain.NewIncomingOrderArgs{
						Reference:   "reference",
						Description: "description",
						Delegator:   "00000000000000000001",
						RevokedAt:   nil,
						ValidFrom:   validFrom,
						ValidUntil:  validUntil,
						Services:    services,
					},
				)

				mocks.db.
					EXPECT().
					ListIncomingOrders(gomock.Any()).
					Return([]*domain.IncomingOrder{
						model,
					}, nil)
			},
			wantResponse: &api.ListIncomingOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "reference",
						Description: "description",
						Delegator:   "00000000000000000001",
						ValidFrom:   timestamppb.New(validFrom),
						ValidUntil:  timestamppb.New(validUntil),
						Services: []*api.OrderService{
							{
								Service: "service-a",
								Organization: &api.Organization{
									Name:         "organization-a",
									SerialNumber: "00000000000000000001",
								},
							},
						},
					},
				},
			},
		},

		"happy_flow_revoked": {
			setup: func(mocks serviceMocks) {
				services := []domain.IncomingOrderService{
					domain.NewIncomingOrderService("service-a", "00000000000000000001", "organization-a"),
				}
				model, _ := domain.NewIncomingOrder(
					&domain.NewIncomingOrderArgs{
						Reference:   "reference",
						Description: "description",
						Delegator:   "00000000000000000001",
						RevokedAt:   &revokedAt,
						ValidFrom:   validFrom,
						ValidUntil:  validUntil,
						Services:    services,
					},
				)

				mocks.db.
					EXPECT().
					ListIncomingOrders(gomock.Any()).
					Return([]*domain.IncomingOrder{
						model,
					}, nil)
			},
			wantResponse: &api.ListIncomingOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "reference",
						Description: "description",
						Delegator:   "00000000000000000001",
						ValidFrom:   timestamppb.New(validFrom),
						ValidUntil:  timestamppb.New(validUntil),
						RevokedAt:   timestamppb.New(revokedAt),
						Services: []*api.OrderService{
							{
								Service: "service-a",
								Organization: &api.Organization{
									Name:         "organization-a",
									SerialNumber: "00000000000000000001",
								},
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

			response, err := service.ListIncomingOrders(context.Background(), &emptypb.Empty{})

			if tt.wantErr == nil {
				assert.Equal(t, tt.wantResponse, response)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, response)
			}
		})
	}
}

//nolint:funlen // this is a test method
func TestListOrders(t *testing.T) {
	validFrom := time.Now()
	validUntil := time.Now().Add(time.Hour)
	revokedAt := time.Now()

	tests := map[string]struct {
		setup   func(*testing.T, serviceMocks, *common_tls.CertificateBundle) context.Context
		want    *external.ListOrdersResponse
		wantErr error
	}{
		"when_retrieval_of_orders_from_database_fails": {
			setup: func(_ *testing.T, mocks serviceMocks, certBundle *common_tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				mocks.db.
					EXPECT().
					ListOutgoingOrdersByOrganization(ctx, certBundle.Certificate().Subject.SerialNumber).
					Return(nil, errors.New("arbitrary error"))

				return ctx
			},
			wantErr: status.Error(codes.Internal, "failed to retrieve external orders"),
		},
		"happy_flow": {
			setup: func(t *testing.T, mocks serviceMocks, certBundle *common_tls.CertificateBundle) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				requesterCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), requesterCertBundle)

				requesterOrganizationSerialNumber := requesterCertBundle.Certificate().Subject.SerialNumber

				mocks.db.
					EXPECT().
					ListOutgoingOrdersByOrganization(ctx, requesterOrganizationSerialNumber).
					Return([]*database.OutgoingOrder{
						{
							Reference:   "reference",
							Description: "description",
							Delegatee:   "10000000000000000001",
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							RevokedAt:   sql.NullTime{},
							OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
								{
									AccessProof: &database.AccessProof{
										OutgoingAccessRequest: &database.OutgoingAccessRequest{
											ServiceName: "service-a",
											Organization: database.Organization{
												Name:         "organization-a",
												SerialNumber: "00000000000000000001"},
										},
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			want: &external.ListOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "reference",
						Description: "description",
						Delegator:   "00000000000000000001",
						ValidFrom:   timestamppb.New(validFrom),
						ValidUntil:  timestamppb.New(validUntil),
						Services: []*api.OrderService{
							{
								Service: "service-a",
								Organization: &api.Organization{
									Name:         "organization-a",
									SerialNumber: "00000000000000000001",
								},
							},
						},
					},
				},
			},
		},
		"happy_flow_revoked": {
			setup: func(t *testing.T, mocks serviceMocks, certBundle *common_tls.CertificateBundle) context.Context {
				pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

				requesterCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
				require.NoError(t, err)

				ctx := setProxyMetadataWithCertBundle(t, context.Background(), requesterCertBundle)

				requesterOrganizationSerialNumber := requesterCertBundle.Certificate().Subject.SerialNumber

				mocks.db.
					EXPECT().
					ListOutgoingOrdersByOrganization(ctx, requesterOrganizationSerialNumber).
					Return([]*database.OutgoingOrder{
						{
							Reference:   "reference",
							Description: "description",
							Delegatee:   "10000000000000000001",
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							RevokedAt: sql.NullTime{
								Valid: true,
								Time:  revokedAt,
							},
							OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
								{
									AccessProof: &database.AccessProof{
										OutgoingAccessRequest: &database.OutgoingAccessRequest{
											ServiceName: "service-a",
											Organization: database.Organization{
												Name:         "organization-a",
												SerialNumber: "00000000000000000001"},
										},
									},
								},
							},
						},
					}, nil)

				return ctx
			},
			want: &external.ListOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "reference",
						Description: "description",
						Delegator:   "00000000000000000001",
						ValidFrom:   timestamppb.New(validFrom),
						ValidUntil:  timestamppb.New(validUntil),
						RevokedAt:   timestamppb.New(revokedAt),
						Services: []*api.OrderService{
							{
								Service: "service-a",
								Organization: &api.Organization{
									Name:         "organization-a",
									SerialNumber: "00000000000000000001",
								},
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
			t.Parallel()

			service, certBundle, mocks := newService(t)
			ctx := tt.setup(t, mocks, certBundle)

			response, err := service.ListOrders(ctx, &emptypb.Empty{})

			if tt.wantErr == nil {
				assert.Equal(t, tt.want, response)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, response)
			}
		})
	}
}
