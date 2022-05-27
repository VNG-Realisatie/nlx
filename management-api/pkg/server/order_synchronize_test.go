package server_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // its a unit test
func TestSynchronizeOrders(t *testing.T) {
	validFrom := time.Now().UTC()
	validUntil := time.Now().Add(time.Hour).UTC()
	revokedAt := time.Now().UTC()

	tests := map[string]struct {
		setup   func(mocks serviceMocks)
		ctx     context.Context
		wantErr bool
		want    *api.SynchronizeOrdersResponse
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(mocks serviceMocks) {},
			wantErr: true,
		},
		"synchronize_fails_when_directory_list_organization_errors": {
			wantErr: true,
			ctx:     testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.EXPECT().
					ListOrganizations(gomock.Any(), &emptypb.Empty{}).
					Return(nil, errors.New("directory fails"))
			},
		},

		"synchronize_does_not_fail_when_directory_get_organization_inway_proxy_address_errors": {
			want: &api.SynchronizeOrdersResponse{Orders: []*api.IncomingOrder{}},
			ctx:  testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.EXPECT().
					ListOrganizations(gomock.Any(), &emptypb.Empty{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "nlx-test",
							},
						},
					}, nil)

				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("", errors.New("directory fails"))
			},
		},

		"synchronization_does_not_fail_when_management_list_orders_errors": {
			want: &api.SynchronizeOrdersResponse{Orders: []*api.IncomingOrder{}},
			ctx:  testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.EXPECT().
					ListOrganizations(gomock.Any(), &emptypb.Empty{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "nlx-test",
							},
						},
					}, nil)

				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("localhost:1234", nil)

				mocks.mc.EXPECT().
					ListOrders(gomock.Any(), &emptypb.Empty{}).
					Return(nil, errors.New("management fails"))

				mocks.mc.EXPECT().Close().Return(nil)
			},
		},

		"synchronization_fails_when_database_synchronize_order_error": {
			wantErr: true,
			ctx:     testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &emptypb.Empty{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "nlx-test",
							},
						},
					}, nil)

				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("localhost:1234", nil)

				mocks.mc.EXPECT().
					ListOrders(gomock.Any(), &emptypb.Empty{}).
					Return(&external.ListOrdersResponse{
						Orders: []*api.IncomingOrder{
							{
								Reference:   "ref-order-1",
								Description: "Order number 1",
								Delegator: &api.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "Organization One",
								},
								RevokedAt:  nil,
								ValidFrom:  timestamppb.New(validFrom),
								ValidUntil: timestamppb.New(validUntil),
								Services: []*api.OrderService{
									{
										Organization: &api.Organization{
											Name:         "organization-a",
											SerialNumber: "00000000000000000001",
										},
										Service: "service-1",
									},
								},
							},
						},
					}, nil)

				mocks.db.EXPECT().
					SynchronizeOrders(gomock.Any(), []*database.IncomingOrder{
						{
							Reference:   "ref-order-1",
							Description: "Order number 1",
							Delegator:   "00000000000000000001",
							RevokedAt:   sql.NullTime{},
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							Services: []database.IncomingOrderService{
								{
									Organization: database.IncomingOrderServiceOrganization{
										Name:         "organization-a",
										SerialNumber: "00000000000000000001",
									},
									Service: "service-1",
								},
							},
						},
					}).
					Return(errors.New("database error"))

				mocks.mc.EXPECT().Close().Return(nil)
			},
		},

		"synchronization_succeeds_on_happy_flow": {
			want: &api.SynchronizeOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "ref-order-1",
						Description: "Order number 1",
						Delegator: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "Organization One",
						},
						RevokedAt:  nil,
						ValidFrom:  timestamppb.New(validFrom),
						ValidUntil: timestamppb.New(validUntil),
						Services: []*api.OrderService{
							{
								Organization: &api.Organization{
									Name:         "Organization One",
									SerialNumber: "00000000000000000001",
								},
								Service: "service-1",
							},
						},
					},
				},
			},
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.EXPECT().
					ListOrganizations(gomock.Any(), &emptypb.Empty{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
						},
					}, nil)

				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("localhost:1234", nil)

				mocks.mc.EXPECT().
					ListOrders(gomock.Any(), &emptypb.Empty{}).
					Return(&external.ListOrdersResponse{
						Orders: []*api.IncomingOrder{
							{
								Reference:   "ref-order-1",
								Description: "Order number 1",
								Delegator: &api.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "Organization One",
								},
								RevokedAt:  nil,
								ValidFrom:  timestamppb.New(validFrom),
								ValidUntil: timestamppb.New(validUntil),
								Services: []*api.OrderService{
									{
										Organization: &api.Organization{
											Name:         "Organization One",
											SerialNumber: "00000000000000000001",
										},
										Service: "service-1",
									},
								},
							},
						},
					}, nil)

				mocks.db.EXPECT().
					SynchronizeOrders(gomock.Any(), []*database.IncomingOrder{
						{
							Reference:   "ref-order-1",
							Description: "Order number 1",
							Delegator:   "00000000000000000001",
							RevokedAt:   sql.NullTime{},
							ValidFrom:   validFrom,
							ValidUntil:  validUntil,
							Services: []database.IncomingOrderService{
								{
									Organization: database.IncomingOrderServiceOrganization{
										Name:         "Organization One",
										SerialNumber: "00000000000000000001",
									},
									Service: "service-1",
								},
							},
						},
					}).
					Return(nil)

				mocks.mc.EXPECT().Close().Return(nil)
			},
		},

		"synchronization_succeeds_on_happy_flow_revoked_orders": {
			want: &api.SynchronizeOrdersResponse{
				Orders: []*api.IncomingOrder{
					{
						Reference:   "ref-order-1",
						Description: "Order number 1",
						Delegator: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "Organization One",
						},
						RevokedAt:  timestamppb.New(revokedAt),
						ValidFrom:  timestamppb.New(validFrom),
						ValidUntil: timestamppb.New(validUntil),
						Services: []*api.OrderService{
							{
								Organization: &api.Organization{
									Name:         "Organization One",
									SerialNumber: "00000000000000000001",
								},
								Service: "service-1",
							},
						},
					},
				},
			},
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &emptypb.Empty{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("localhost:1234", nil)

				mocks.mc.
					EXPECT().
					ListOrders(gomock.Any(), &emptypb.Empty{}).
					Return(&external.ListOrdersResponse{
						Orders: []*api.IncomingOrder{
							{
								Reference:   "ref-order-1",
								Description: "Order number 1",
								Delegator: &api.Organization{
									SerialNumber: "00000000000000000001",
									Name:         "Organization One",
								},
								RevokedAt:  timestamppb.New(revokedAt),
								ValidFrom:  timestamppb.New(validFrom),
								ValidUntil: timestamppb.New(validUntil),
								Services: []*api.OrderService{
									{
										Organization: &api.Organization{
											Name:         "Organization One",
											SerialNumber: "00000000000000000001",
										},
										Service: "service-1",
									},
								},
							},
						},
					}, nil)

				mocks.db.
					EXPECT().
					SynchronizeOrders(gomock.Any(), []*database.IncomingOrder{
						{
							Reference:   "ref-order-1",
							Description: "Order number 1",
							Delegator:   "00000000000000000001",
							RevokedAt: sql.NullTime{
								Valid: true,
								Time:  revokedAt,
							},
							ValidFrom:  validFrom,
							ValidUntil: validUntil,
							Services: []database.IncomingOrderService{
								{
									Organization: database.IncomingOrderServiceOrganization{
										Name:         "Organization One",
										SerialNumber: "00000000000000000001",
									},
									Service: "service-1",
								},
							},
						},
					}).
					Return(nil)

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			tt.setup(mocks)

			orders, err := service.SynchronizeOrders(tt.ctx, &emptypb.Empty{})

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, orders)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, orders, tt.want)
			}
		})
	}
}
