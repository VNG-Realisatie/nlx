// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestVerifyClaim(t *testing.T) {
	tests := map[string]struct {
		request *api.VerifyClaimRequest
		setup   func(*server.ManagementService, *common_tls.CertificateBundle, serviceMocks) context.Context
		want    func(*testing.T, *common_tls.CertificateBundle, *external.RequestClaimResponse)
		wantErr error
	}{
		"when_the_proxy_metadata_is_missing": {
			request: &api.VerifyClaimRequest{},
			setup: func(*server.ManagementService, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_providing_an_empty_claim": {
			request: &api.VerifyClaimRequest{
				Claim: "",
			},
			setup: func(*server.ManagementService, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return setProxyMetadata(context.Background())
			},
			wantErr: status.Error(codes.InvalidArgument, "a claim must be provided"),
		},
		"when_providing_an_empty_service_name": {
			request: &api.VerifyClaimRequest{
				Claim:       "arbitrary-claim",
				ServiceName: "",
			},
			setup: func(*server.ManagementService, *common_tls.CertificateBundle, serviceMocks) context.Context {
				return setProxyMetadata(context.Background())
			},
			wantErr: status.Error(codes.InvalidArgument, "a service name must be provided"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, orgCert, mocks := newService(t)
			ctx := tt.setup(service, orgCert, mocks)

			_, err := service.VerifyClaim(ctx, tt.request)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestVerifyClaimWhenAccessGrantIsRevoked(t *testing.T) {
	service, orgCert, mocks := newService(t)
	ctx := setProxyMetadata(context.Background())

	mocks.db.EXPECT().
		GetLatestAccessGrantForService(ctx, "delegator-organization-name", "service-name").
		Return(&database.AccessGrant{
			RevokedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}, nil)

	claim, err := getJWTAsSignedString(orgCert)
	assert.Nil(t, err)

	response, err := service.VerifyClaim(ctx, &api.VerifyClaimRequest{
		ServiceName: "service-name",
		Claim:       claim,
	})

	assert.Nil(t, response)
	assert.ErrorIs(t, err, status.Error(codes.Unauthenticated, "access grant for service has been revoked"))
}

func TestVerifyClaimFailedToRetrieveAccessGrant(t *testing.T) {
	service, orgCert, mocks := newService(t)
	ctx := setProxyMetadata(context.Background())

	mocks.db.EXPECT().
		GetLatestAccessGrantForService(ctx, "delegator-organization-name", "service-name").
		Return(nil, errors.New("arbitrary error"))

	claim, err := getJWTAsSignedString(orgCert)
	assert.Nil(t, err)

	response, err := service.VerifyClaim(ctx, &api.VerifyClaimRequest{
		ServiceName: "service-name",
		Claim:       claim,
	})

	assert.Nil(t, response)
	assert.ErrorIs(t, err, status.Error(codes.Internal, "unable to verify claim"))
}

func TestVerifyClaimWithoutAccessGrant(t *testing.T) {
	service, orgCert, mocks := newService(t)
	ctx := setProxyMetadata(context.Background())

	mocks.db.EXPECT().
		GetLatestAccessGrantForService(ctx, "delegator-organization-name", "service-name").
		Return(nil, database.ErrNotFound)

	claim, err := getJWTAsSignedString(orgCert)
	assert.Nil(t, err)

	response, err := service.VerifyClaim(ctx, &api.VerifyClaimRequest{
		ServiceName: "service-name",
		Claim:       claim,
	})

	assert.Nil(t, response)
	assert.ErrorIs(t, err, status.Error(codes.Unauthenticated, "no access grant available for service"))
}

func TestVerifyClaimHappyFlow(t *testing.T) {
	service, orgCert, mocks := newService(t)
	ctx := setProxyMetadata(context.Background())

	publicKeyPEM, err := orgCert.PublicKeyPEM()
	assert.Nil(t, err)

	mocks.db.EXPECT().
		GetLatestAccessGrantForService(ctx, "delegator-organization-name", "service-name").
		Return(&database.AccessGrant{
			IncomingAccessRequest: &database.IncomingAccessRequest{
				PublicKeyPEM:         publicKeyPEM,
				PublicKeyFingerprint: orgCert.PublicKeyFingerprint(),
			},
		}, nil)

	claim, err := getJWTAsSignedString(orgCert)
	assert.Nil(t, err)

	response, err := service.VerifyClaim(ctx, &api.VerifyClaimRequest{
		ServiceName: "service-name",
		Claim:       claim,
	})

	assert.NoError(t, err)
	assert.Equal(t, &api.VerifyClaimResponse{
		OrderOrganizationName: "delegator-organization-name",
		OrderReference:        "order-reference",
	}, response)
}

func getJWTAsSignedString(orgCert *common_tls.CertificateBundle) (string, error) {
	claims := server.JWTClaims{
		Organization:   "delegatee-organization-name",
		OrderReference: "order-reference",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "delegator-organization-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedString, err := token.SignedString(orgCert.PrivateKey())
	if err != nil {
		return "", err
	}

	return signedString, nil
}
