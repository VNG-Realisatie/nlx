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

	common_testing "go.nlx.io/nlx/common/testing"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test method
func TestListOutgoingOrders(t *testing.T) {
	validFrom := time.Now()
	validUntil := time.Now().Add(time.Hour)

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
							Reference:   "reference",
							Description: "description",
							Delegatee:   "saas-organization-x",
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							RevokedAt:   sql.NullTime{},
							Services: []database.OutgoingOrderService{
								{
									Service: "service-a",
									Organization: database.OutgoingOrderServiceOrganization{
										Name:         "a-organization",
										SerialNumber: "00000000000000000001",
									},
								},
							},
						},
					}, nil)
			},
			wantResponse: &api.ListOutgoingOrdersResponse{
				Orders: []*api.OutgoingOrder{
					{
						Reference:   "reference",
						Description: "description",
						Delegatee:   "saas-organization-x",
						ValidFrom:   timestamppb.New(validFrom),
						ValidUntil:  timestamppb.New(validUntil),
						Services: []*api.OrderService{
							{
								Service: "service-a",
								Organization: &api.Organization{
									Name:         "a-organization",
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
						Delegator:   "nlx-test",
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
						Delegator:   "nlx-test",
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
						Delegator:   "nlx-test",
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
						Delegator:   "nlx-test",
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
					ListOutgoingOrdersByOrganization(ctx, certBundle.Certificate().Subject.Organization[0]).
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

				requesterOrganizationName := requesterCertBundle.Certificate().Subject.Organization[0]

				mocks.db.
					EXPECT().
					ListOutgoingOrdersByOrganization(ctx, requesterOrganizationName).
					Return([]*database.OutgoingOrder{
						{
							Reference:   "reference",
							Description: "description",
							Delegatee:   "saas-organization-x",
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							RevokedAt:   sql.NullTime{},
							Services: []database.OutgoingOrderService{
								{
									Service: "service-a",
									Organization: database.OutgoingOrderServiceOrganization{
										Name:         "organization-a",
										SerialNumber: "00000000000000000001",
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
						Delegator:   "nlx-test",
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

				requesterOrganizationName := requesterCertBundle.Certificate().Subject.Organization[0]

				mocks.db.
					EXPECT().
					ListOutgoingOrdersByOrganization(ctx, requesterOrganizationName).
					Return([]*database.OutgoingOrder{
						{
							Reference:   "reference",
							Description: "description",
							Delegatee:   "saas-organization-x",
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							RevokedAt: sql.NullTime{
								Valid: true,
								Time:  revokedAt,
							},
							Services: []database.OutgoingOrderService{
								{
									Service: "service-a",
									Organization: database.OutgoingOrderServiceOrganization{
										Name:         "organization-a",
										SerialNumber: "00000000000000000001",
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
						Delegator:   "nlx-test",
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
