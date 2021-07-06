// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

const expiresInHours = 4

func (s *ManagementService) RequestClaim(ctx context.Context, req *external.RequestClaimRequest) (*external.RequestClaimResponse, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		s.logger.Error("failed to parse proxy metadata", zap.Error(err))
		return nil, err
	}

	if len(req.OrderReference) < 1 {
		return nil, status.Error(codes.InvalidArgument, "an order reference must be provided")
	}

	order, err := s.configDatabase.GetOutgoingOrderByReference(ctx, req.OrderReference)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "order with reference %s not found", req.OrderReference)
		}

		return nil, status.Error(codes.Internal, "failed to find order")
	}

	if order.Delegatee != md.OrganizationName {
		return nil, status.Errorf(codes.NotFound, "order with reference %s and organization %s not found", req.OrderReference, md.OrganizationName)
	}

	block, _ := pem.Decode([]byte(order.PublicKeyPEM))
	if block == nil {
		s.logger.Error("invalid public key format", zap.Error(err))

		return nil, status.Error(codes.Internal, "invalid public key format")
	}

	sum := sha256.Sum256(block.Bytes)
	fingerprint := base64.StdEncoding.EncodeToString(sum[:])

	if fingerprint != md.PublicKeyFingerprint {
		return nil, status.Errorf(codes.Unauthenticated, "invalid public key for order")
	}

	if time.Now().After(order.ValidUntil) {
		return nil, status.Errorf(codes.Unauthenticated, "order is no longer valid")
	}

	expiresAt := time.Now().Add(expiresInHours * time.Hour)

	if expiresAt.After(order.ValidUntil) {
		expiresAt = order.ValidUntil
	}

	claims := delegation.JWTClaims{
		Services:       make([]delegation.Service, len(order.Services)),
		Delegatee:      md.OrganizationName,
		OrderReference: req.OrderReference,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    s.orgCert.Certificate().Subject.Organization[0],
		},
	}

	for i, service := range order.Services {
		claims.Services[i] = delegation.Service{
			Organization: service.Organization,
			Service:      service.Service,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedString, err := token.SignedString(s.orgCert.PrivateKey())
	if err != nil {
		s.logger.Error("unable to create signed string from private key", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to sign claim")
	}

	return &external.RequestClaimResponse{
		Claim: signedString,
	}, nil
}
