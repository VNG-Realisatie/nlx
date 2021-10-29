// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

const expiresInHours = 4

var errMessageOrderRevoked = "order is revoked"

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

	if order.Delegatee != md.OrganizationSerialNumber {
		return nil, status.Errorf(codes.NotFound, "order with reference %s and organization serialnumber %s not found", req.OrderReference, md.OrganizationSerialNumber)
	}

	fingerprint, err := tls.PemPublicKeyFingerprint([]byte(order.PublicKeyPEM))
	if err != nil {
		s.logger.Error("invalid public key format", zap.Error(err))
		return nil, status.Error(codes.Internal, "invalid public key format")
	}

	if fingerprint != md.PublicKeyFingerprint {
		return nil, status.Errorf(codes.Unauthenticated, "invalid public key for order")
	}

	if order.RevokedAt.Valid {
		return nil, status.Errorf(codes.Unauthenticated, errMessageOrderRevoked)
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
		Delegatee:      md.OrganizationSerialNumber,
		OrderReference: req.OrderReference,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    s.orgCert.Certificate().Subject.SerialNumber,
		},
	}

	for i, service := range order.Services {
		claims.Services[i] = delegation.Service{
			OrganizationSerialNumber: service.Organization.SerialNumber,
			Service:                  service.Service,
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
