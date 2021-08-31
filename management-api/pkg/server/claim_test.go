// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/delegation"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// nolint:funlen // this is a test
func TestRequestClaim(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		request        *external.RequestClaimRequest
		ctx            context.Context
		setup          func(*common_tls.CertificateBundle, serviceMocks)
		wantCode       codes.Code
		wantErrMessage string
		wantValidUntil time.Time
	}{
		"when_the_proxy_metadata_is_missing": {
			request:        &external.RequestClaimRequest{},
			ctx:            context.Background(),
			wantCode:       codes.Internal,
			wantErrMessage: "missing metadata from the management proxy",
		},
		"when_providing_an_empty_order_reference": {
			request: &external.RequestClaimRequest{
				OrderReference: "",
			},
			ctx:            setProxyMetadata(context.Background()),
			wantCode:       codes.InvalidArgument,
			wantErrMessage: "an order reference must be provided",
		},
		"when_public_key_is_invalid": {
			request: &external.RequestClaimRequest{
				OrderReference: "arbitrary-order-reference",
			},
			ctx: setProxyMetadata(context.Background()),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "arbitrary-order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    "organization-a",
						PublicKeyPEM: "arbitrary-public-key-in-pem-format",
					}, nil)
			},
			wantCode:       codes.Internal,
			wantErrMessage: "invalid public key format",
		},
		"when_public_key_fingerprint_does_not_equal_metadata_fingerprint": {
			request: &external.RequestClaimRequest{
				OrderReference: "arbitrary-order-reference",
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs( // data from setProxyMetadata()
				"nlx-organization", "organization-a",
				"nlx-public-key-der", "ZHVtbXktcHVibGljLWtleQo=",
				"nlx-public-key-fingerprint", "invalid=",
			)),
			setup: func(orgCerts *common_tls.CertificateBundle, mocks serviceMocks) {
				publicKeyPEM, _ := orgCerts.PublicKeyPEM()

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "arbitrary-order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    "organization-a",
						PublicKeyPEM: publicKeyPEM,
						ValidUntil:   now.Add(4 * time.Hour),
					}, nil)
			},
			wantCode:       codes.Unauthenticated,
			wantErrMessage: "invalid public key for order",
		},
		"when_order_is revoked": {
			request: &external.RequestClaimRequest{
				OrderReference: "arbitrary-order-reference",
			},
			ctx: setProxyMetadata(context.Background()),
			setup: func(orgCerts *common_tls.CertificateBundle, mocks serviceMocks) {
				publicKeyPEM, _ := orgCerts.PublicKeyPEM()

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "arbitrary-order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    "organization-a",
						PublicKeyPEM: publicKeyPEM,
						RevokedAt: sql.NullTime{
							Time:  now.Add(-1 * time.Hour),
							Valid: true,
						},
					}, nil)
			},
			wantCode:       codes.Unauthenticated,
			wantErrMessage: "order is revoked",
		},
		"when_order_is_no_longer_valid": {
			request: &external.RequestClaimRequest{
				OrderReference: "arbitrary-order-reference",
			},
			ctx: setProxyMetadata(context.Background()),
			setup: func(orgCerts *common_tls.CertificateBundle, mocks serviceMocks) {
				publicKeyPEM, _ := orgCerts.PublicKeyPEM()

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "arbitrary-order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    "organization-a",
						PublicKeyPEM: publicKeyPEM,
						ValidUntil:   now.Add(-1 * time.Hour),
					}, nil)
			},
			wantCode:       codes.Unauthenticated,
			wantErrMessage: "order is no longer valid",
		},
		"happy_flow_with_short_valid_until": {
			request: &external.RequestClaimRequest{
				OrderReference: "arbitrary-order-reference",
			},
			ctx: setProxyMetadata(context.Background()),
			setup: func(orgCerts *common_tls.CertificateBundle, mocks serviceMocks) {
				publicKeyPEM, _ := orgCerts.PublicKeyPEM()

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "arbitrary-order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    "organization-a",
						PublicKeyPEM: publicKeyPEM,
						ValidUntil:   now.Add(2 * time.Hour),
					}, nil)
			},
			wantValidUntil: now.Add(2 * time.Hour),
		},
		"happy_flow": {
			request: &external.RequestClaimRequest{
				OrderReference: "arbitrary-order-reference",
			},
			ctx: setProxyMetadata(context.Background()),
			setup: func(orgCerts *common_tls.CertificateBundle, mocks serviceMocks) {
				publicKeyPEM, _ := orgCerts.PublicKeyPEM()

				mocks.db.
					EXPECT().
					GetOutgoingOrderByReference(gomock.Any(), "arbitrary-order-reference").
					Return(&database.OutgoingOrder{
						Delegatee:    "organization-a",
						PublicKeyPEM: publicKeyPEM,
						ValidUntil:   now.Add(4 * time.Hour),
					}, nil)
			},
			wantValidUntil: now.Add(4 * time.Hour),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, bundle, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(bundle, mocks)
			}

			response, err := service.RequestClaim(tt.ctx, tt.request)

			if tt.wantCode != codes.OK {
				assert.Error(t, err)

				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
				assert.Equal(t, tt.wantErrMessage, st.Message())
			} else {
				assert.NoError(t, err)

				token, err := jwt.ParseWithClaims(response.Claim, &delegation.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
					return bundle.PublicKey(), nil
				})

				assert.NoError(t, err)

				claims := token.Claims.(*delegation.JWTClaims)
				assert.Equal(t, claims.Delegatee, "organization-a")
				assert.Equal(t, claims.OrderReference, "arbitrary-order-reference")
				assert.Equal(t, claims.Issuer, "nlx-test")
				assert.Equal(t, claims.ExpiresAt, tt.wantValidUntil.Unix())
			}
		})
	}
}
