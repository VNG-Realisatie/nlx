// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
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
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test method
func TestListIssuedOrders(t *testing.T) {
	validFrom := time.Now()
	validUntil := time.Now().Add(time.Hour)

	tests := map[string]struct {
		setup        func(serviceMocks)
		wantResponse *api.ListIssuedOrdersResponse
		wantErr      error
	}{
		"when_retrieval_of_orders_from_database_fails": {
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().ListIssuedOrders(gomock.Any()).Return(nil, errors.New("arbitrary error"))
			},
			wantErr: status.Error(codes.Internal, "failed to retrieve issued orders"),
		},
		"happy_path": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListIssuedOrders(gomock.Any()).
					Return([]*database.Order{
						{
							Reference:   "reference",
							Description: "description",
							Delegatee:   "saas-organization-x",
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							Services: []database.OrderService{
								{
									Service:      "service-a",
									Organization: "organization-a",
								},
							},
						},
					}, nil)
			},
			wantResponse: &api.ListIssuedOrdersResponse{
				Orders: []*api.Order{
					{
						Reference:   "reference",
						Description: "description",
						Delegatee:   "saas-organization-x",
						ValidFrom:   timestampProto(validFrom),
						ValidUntil:  timestampProto(validUntil),
						Services: []*api.Order_Service{
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

			response, err := service.ListIssuedOrders(context.Background(), &emptypb.Empty{})

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

	tests := map[string]struct {
		setup        func(serviceMocks)
		wantResponse *external.ListOrdersResponse
		wantErr      error
	}{
		"when_retrieval_of_orders_from_database_fails": {
			setup: func(mocks serviceMocks) {
				mocks.db.EXPECT().
					ListOrdersByOrganization(gomock.Any(), "organization-a").
					Return(nil, errors.New("arbitrary error"))
			},
			wantErr: status.Error(codes.Internal, "failed to retrieve issued orders"),
		},
		"happy_path": {
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListOrdersByOrganization(gomock.Any(), "organization-a").
					Return([]*database.Order{
						{
							Reference:   "reference",
							Description: "description",
							Delegatee:   "saas-organization-x",
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							Services: []database.OrderService{
								{
									Service:      "service-a",
									Organization: "organization-a",
								},
							},
						},
					}, nil)
			},
			wantResponse: &external.ListOrdersResponse{
				Orders: []*api.Order{
					{
						Reference:   "reference",
						Description: "description",
						Delegatee:   "saas-organization-x",
						ValidFrom:   timestampProto(validFrom),
						ValidUntil:  timestampProto(validUntil),
						Services: []*api.Order_Service{
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
