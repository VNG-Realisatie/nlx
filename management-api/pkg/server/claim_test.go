// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestRequestClaim(t *testing.T) {
	tests := map[string]struct {
		request  *external.RequestClaimRequest
		ctx      context.Context
		want     func(*testing.T, *common_tls.CertificateBundle, *external.RequestClaimResponse)
		wantCode codes.Code
	}{
		"when_the_proxy_metadata_is_missing": {
			request:  &external.RequestClaimRequest{},
			ctx:      context.Background(),
			wantCode: codes.Internal,
		},
		"when_providing_an_empty_order_reference": {
			request: &external.RequestClaimRequest{
				OrderReference: "",
			},
			ctx:      setProxyMetadata(context.Background()),
			wantCode: codes.InvalidArgument,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, _ := newService(t)

			_, err := service.RequestClaim(tt.ctx, tt.request)
			assert.Error(t, err)

			st, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, st.Code())
		})
	}
}

func TestRequestClaimHappyFlow(t *testing.T) {
	service, bundle, _ := newService(t)
	ctx := setProxyMetadata(context.Background())

	response, err := service.RequestClaim(ctx, &external.RequestClaimRequest{
		OrderReference: "arbitrary-order-reference",
	})
	assert.NoError(t, err)
	assert.NotNil(t, response)

	token, err := jwt.ParseWithClaims(response.Claim, &server.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return bundle.PublicKey(), nil
	})

	assert.NoError(t, err)

	claims := token.Claims.(*server.JWTClaims)
	assert.Equal(t, claims.Organization, "organization-a")
	assert.Equal(t, claims.OrderReference, "arbitrary-order-reference")
	assert.Equal(t, claims.Issuer, "nlx-test")
	assert.Equal(t, claims.ExpiresAt, time.Now().Add(time.Hour).Unix())
}
