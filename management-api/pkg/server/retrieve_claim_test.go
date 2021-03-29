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

func TestRetrieveClaim(t *testing.T) {
	tests := map[string]struct {
		request *api.RetrieveClaimForOrderRequest
		setup   func(*server.ManagementService, serviceMocks) context.Context
		want    func(*testing.T, *common_tls.CertificateBundle, *external.RequestClaimResponse)
		wantErr error
	}{
		"when_the_proxy_metadata_is_missing": {
			request: &api.RetrieveClaimForOrderRequest{},
			setup: func(*server.ManagementService, serviceMocks) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_providing_an_empty_order_reference": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference: "",
			},
			setup: func(*server.ManagementService, serviceMocks) context.Context {
				return setProxyMetadata(context.Background())
			},
			wantErr: status.Error(codes.InvalidArgument, "an order reference must be provided"),
		},
		"when_providing_an_empty_order_organization_name": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference:        "arbitrary-order-reference",
				OrderOrganizationName: "",
			},
			setup: func(*server.ManagementService, serviceMocks) context.Context {
				return setProxyMetadata(context.Background())
			},
			wantErr: status.Error(codes.InvalidArgument, "an organization name of the order must be provided"),
		},
		"when_getting_the_organization_inway_proxy_address_fails": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference:        "order-reference-a",
				OrderOrganizationName: "organization-a",
			},
			setup: func(service *server.ManagementService, mocks serviceMocks) context.Context {
				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), gomock.Any()).
					Return("", errors.New("arbitrary error"))

				return setProxyMetadata(context.Background())
			},
			wantErr: status.Error(codes.Internal, "unable to retrieve claim"),
		},
		"when_creating_the_management_client_fails": {
			request: &api.RetrieveClaimForOrderRequest{
				OrderReference:        "order-reference-a",
				OrderOrganizationName: "organization-a",
			},
			setup: func(service *server.ManagementService, mocks serviceMocks) context.Context {
				mocks.dc.EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), gomock.Any()).
					Return("inway-address", nil)

				mocks.mc.EXPECT().Close()

				mocks.mc.EXPECT().
					RequestClaim(gomock.Any(), &external.RequestClaimRequest{
						OrderReference: "order-reference-a",
					}).
					Return(nil, errors.New("arbitrary error"))

				return setProxyMetadata(context.Background())
			},
			wantErr: status.Error(codes.Internal, "unable to retrieve claim"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			ctx := tt.setup(service, mocks)

			_, err := service.RetrieveClaimForOrder(ctx, tt.request)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestRetrieveClaimHappyFlow(t *testing.T) {
	service, _, mocks := newService(t)
	ctx := setProxyMetadata(context.Background())

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
		OrderReference:        "order-reference-a",
		OrderOrganizationName: "organization-a",
	})

	assert.NoError(t, err)
	assert.Equal(t, response, &api.RetrieveClaimForOrderResponse{
		Claim: "claim",
	})
}
