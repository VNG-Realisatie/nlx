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
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

//nolint:funlen // its a unittest
func TestRequestClaim(t *testing.T) {
	tests := map[string]struct {
		request  *external.RequestClaimRequest
		setup    func(*mock_database.MockConfigDatabase) context.Context
		want     func(*testing.T, *common_tls.CertificateBundle, *external.RequestClaimResponse)
		wantCode codes.Code
	}{
		"when_the_proxy_metadata_is_missing": {
			request: &external.RequestClaimRequest{},
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
			wantCode: codes.Internal,
		},
		"when_providing_an_empty_order_reference": {
			request: &external.RequestClaimRequest{
				OrderReference: "",
			},
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return setProxyMetadata(context.Background())
			},
			wantCode: codes.InvalidArgument,
		},
		"happy_flow": {
			request: &external.RequestClaimRequest{
				OrderReference: "arbitrary-order-reference",
			},
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return setProxyMetadata(context.Background())
			},
			wantCode: codes.OK,
			want: func(t *testing.T, bundle *common_tls.CertificateBundle, resp *external.RequestClaimResponse) {
				token, err := jwt.ParseWithClaims(resp.Claim, &server.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
					return bundle.PublicKey(), nil
				})

				assert.NoError(t, err)

				claims := token.Claims.(*server.JWTClaims)
				assert.Equal(t, claims.Organization, "organization-a")
				assert.Equal(t, claims.OrderReference, "arbitrary-order-reference")
				assert.Equal(t, claims.Issuer, "nlx-test")
				assert.Equal(t, claims.ExpiresAt, time.Now().Add(time.Hour).Unix())
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db, _, bundle := newService(t)
			ctx := tt.setup(db)

			response, err := service.RequestClaim(ctx, tt.request)

			if tt.wantCode > 0 {
				assert.Error(t, err)

				st, ok := status.FromError(err)

				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
			} else {
				assert.NoError(t, err)
				if assert.NotNil(t, response) {
					tt.want(t, bundle, response)
				}
			}
		})
	}
}
