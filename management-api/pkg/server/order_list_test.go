// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // code is not actually duplicated, the linter has lost it's mind
package server_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

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
									Service:      "service-a",
									Organization: "organization-a",
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
						ValidFrom:   timestampProto(validFrom),
						ValidUntil:  timestampProto(validUntil),
						Services: []*api.OrderService{
							{
								Service:      "service-a",
								Organization: "organization-a",
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
					domain.NewIncomingOrderService("service-a", "organization-a"),
				}
				model, _ := domain.NewIncomingOrder(
					"reference",
					"description",
					"nlx-test",
					nil,
					validFrom,
					validUntil,
					services,
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
						ValidFrom:   timestampProto(validFrom),
						ValidUntil:  timestampProto(validUntil),
						Services: []*api.OrderService{
							{
								Service:      "service-a",
								Organization: "organization-a",
							},
						},
					},
				},
			},
		},

		"happy_flow_revoked": {
			setup: func(mocks serviceMocks) {
				services := []domain.IncomingOrderService{
					domain.NewIncomingOrderService("service-a", "organization-a"),
				}
				model, _ := domain.NewIncomingOrder(
					"reference",
					"description",
					"nlx-test",
					&revokedAt,
					validFrom,
					validUntil,
					services,
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
						ValidFrom:   timestampProto(validFrom),
						ValidUntil:  timestampProto(validUntil),
						RevokedAt:   timestampProto(revokedAt),
						Services: []*api.OrderService{
							{
								Service:      "service-a",
								Organization: "organization-a",
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
		setup        func(serviceMocks)
		wantResponse *external.ListOrdersResponse
		wantErr      error
	}{
		"when_retrieval_of_orders_from_database_fails": {
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().
					ListOutgoingOrdersByOrganization(gomock.Any(), "organization-a").
					Return(nil, errors.New("arbitrary error"))
			},
			wantErr: status.Error(codes.Internal, "failed to retrieve external orders"),
		},
		"happy_flow": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListOutgoingOrdersByOrganization(gomock.Any(), "organization-a").
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
									Service:      "service-a",
									Organization: "organization-a",
								},
							},
						},
					}, nil)
			},
			wantResponse: &external.ListOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "reference",
						Description: "description",
						Delegator:   "nlx-test",
						ValidFrom:   timestampProto(validFrom),
						ValidUntil:  timestampProto(validUntil),
						Services: []*api.OrderService{
							{
								Service:      "service-a",
								Organization: "organization-a",
							},
						},
					},
				},
			},
		},
		"happy_flow_revoked": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListOutgoingOrdersByOrganization(gomock.Any(), "organization-a").
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
									Service:      "service-a",
									Organization: "organization-a",
								},
							},
						},
					}, nil)
			},
			wantResponse: &external.ListOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "reference",
						Description: "description",
						Delegator:   "nlx-test",
						ValidFrom:   timestampProto(validFrom),
						ValidUntil:  timestampProto(validUntil),
						RevokedAt:   timestampProto(revokedAt),
						Services: []*api.OrderService{
							{
								Service:      "service-a",
								Organization: "organization-a",
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
			ctx := setProxyMetadata(context.Background())

			if tt.setup != nil {
				tt.setup(mocks)
			}

			response, err := service.ListOrders(ctx, &emptypb.Empty{})

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
