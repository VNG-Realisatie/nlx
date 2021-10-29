// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/server"
)

// nolint funlen: this is a test function
func TestRetrieveClaim(t *testing.T) {
	tests := map[string]struct {
		request *api.RetrieveClaimForOrderRequest
		setup   func(*testing.T, *server.ManagementService, serviceMocks) context.Context
		want    func(*testing.T, *common_tls.CertificateBundle, *external.RequestClaimResponse)
		wantErr error
	}{
		"when_providing_an_empty_order_reference": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference: "",
			},
			setup: func(*testing.T, *server.ManagementService, serviceMocks) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.InvalidArgument, "an order reference must be provided"),
		},
		"when_providing_an_empty_order_organization_name": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference:                "arbitrary-order-reference",
				OrderOrganizationSerialNumber: "",
			},
			setup: func(*testing.T, *server.ManagementService, serviceMocks) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.InvalidArgument, "an organization serial number of the order must be provided"),
		},
		"when_getting_the_organization_inway_proxy_address_fails": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference:                "order-reference-a",
				OrderOrganizationSerialNumber: "organization-a", // @TODO serial number
			},
			setup: func(t *testing.T, service *server.ManagementService, mocks serviceMocks) context.Context {
				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), gomock.Any()).
					Return("", errors.New("arbitrary error"))

				return setProxyMetadata(t, context.Background())
			},
			wantErr: status.Error(codes.Internal, "unable to retrieve claim"),
		},
		// nolint dupl: linter is unable to detect difference in error message
		"when_creating_the_management_client_fails": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference:                "order-reference-a",
				OrderOrganizationSerialNumber: "organization-a",
			},
			setup: func(_ *testing.T, service *server.ManagementService, mocks serviceMocks) context.Context {
				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), gomock.Any()).
					Return("inway-address", nil)

				mocks.mc.EXPECT().Close()

				mocks.mc.EXPECT().
					RequestClaim(gomock.Any(), &external.RequestClaimRequest{
						OrderReference: "order-reference-a",
					}).
					Return(nil, errors.New("arbitrary error"))

				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "unable to retrieve claim"),
		},
		// nolint dupl: linter is unable to detect difference in error message
		"when_order_is_revoked": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference:                "order-reference-a",
				OrderOrganizationSerialNumber: "organization-a",
			},
			setup: func(_ *testing.T, service *server.ManagementService, mocks serviceMocks) context.Context {
				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), gomock.Any()).
					Return("inway-address", nil)

				mocks.mc.EXPECT().Close()

				mocks.mc.EXPECT().
					RequestClaim(gomock.Any(), &external.RequestClaimRequest{
						OrderReference: "order-reference-a",
					}).
					Return(nil, status.Errorf(codes.Unauthenticated, "order is revoked"))

				return context.Background()
			},
			wantErr: status.Error(codes.Unauthenticated, "order is revoked"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			ctx := tt.setup(t, service, mocks)

			_, err := service.RetrieveClaimForOrder(ctx, tt.request)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestRetrieveClaimHappyFlow(t *testing.T) {
	service, _, mocks := newService(t)
	ctx := setProxyMetadata(t, context.Background())

	mocks.dc.EXPECT().
		GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
		Return("inway-proxy-address", nil)

	mocks.mc.EXPECT().Close()

	mocks.mc.EXPECT().
		RequestClaim(gomock.Any(), &external.RequestClaimRequest{
			OrderReference: "order-reference-a",
		}).
		Return(&external.RequestClaimResponse{
			Claim: "claim",
		}, nil)

	response, err := service.RetrieveClaimForOrder(ctx, &api.RetrieveClaimForOrderRequest{
		OrderReference:                "order-reference-a",
		OrderOrganizationSerialNumber: "organization-a",
	})

	assert.NoError(t, err)
	assert.Equal(t, response, &api.RetrieveClaimForOrderResponse{
		Claim: "claim",
	})
}
