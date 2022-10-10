// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"crypto"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/delegation"
	outwayapi "go.nlx.io/nlx/outway/api"
	"go.nlx.io/nlx/outway/pkg/server"
)

func TestSignOrderClaim(t *testing.T) {
	now := time.Now()

	const mockSignedClaim = "signed-claim"

	tests := map[string]struct {
		signer  server.SignFunction
		request *outwayapi.SignOrderClaimRequest
		want    *outwayapi.SignOrderClaimResponse
		wantErr error
	}{
		"invalid_expires_at": {
			request: &outwayapi.SignOrderClaimRequest{
				Delegatee:                     "00000000000000000001",
				DelegateePublicKeyFingerprint: "public-key-fingerprint",
				OrderReference:                "order-reference",
				AccessProof: &outwayapi.AccessProof{
					ServiceName:              "mock-service",
					OrganizationSerialNumber: "00000000000000000002",
					PublicKeyFingerprint:     "public-key-fingerprint-2",
				},
				ExpiresAt: nil,
			},
			wantErr: status.Error(codes.Internal, "invalid expiry time provided"),
		},
		"signing_claim_fails": {
			signer: func(_ crypto.PrivateKey, _ *delegation.JWTClaims) (string, error) {
				return "", fmt.Errorf("arbitrary error")
			},
			request: &outwayapi.SignOrderClaimRequest{
				Delegatee:                     "00000000000000000001",
				DelegateePublicKeyFingerprint: "public-key-fingerprint",
				OrderReference:                "order-reference",
				AccessProof: &outwayapi.AccessProof{
					ServiceName:              "mock-service",
					OrganizationSerialNumber: "00000000000000000002",
					PublicKeyFingerprint:     "public-key-fingerprint-2",
				},
				ExpiresAt: timestamppb.New(now),
			},
			wantErr: status.Error(codes.Internal, "unable to sign claim"),
		},
		"happy_flow": {
			signer: func(_ crypto.PrivateKey, _ *delegation.JWTClaims) (string, error) {
				return mockSignedClaim, nil
			},
			request: &outwayapi.SignOrderClaimRequest{
				Delegatee:                     "00000000000000000001",
				DelegateePublicKeyFingerprint: "public-key-fingerprint",
				OrderReference:                "order-reference",
				AccessProof: &outwayapi.AccessProof{
					ServiceName:              "mock-service",
					OrganizationSerialNumber: "00000000000000000002",
					PublicKeyFingerprint:     "public-key-fingerprint-2",
				},
				ExpiresAt: timestamppb.New(now),
			},
			want: &outwayapi.SignOrderClaimResponse{
				SignedOrderClaim: mockSignedClaim,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service := newService(t, tt.signer)

			ctx := context.Background()

			actual, err := service.SignOrderClaim(ctx, tt.request)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}
