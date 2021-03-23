// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

var ErrNoAccessGrantAvailable = errors.New("no access grant available for service")
var ErrNoAccessGrantRevoked = errors.New("access grant is revoked")

func (s *ManagementService) VerifyClaim(ctx context.Context, req *api.VerifyClaimRequest) (*api.VerifyClaimResponse, error) {
	_, parsePmErr := s.parseProxyMetadata(ctx)
	if parsePmErr != nil {
		s.logger.Error("failed to parse proxy metadata", zap.Error(parsePmErr))
		return nil, parsePmErr
	}

	if len(req.Claim) < 1 {
		return nil, status.Error(codes.InvalidArgument, "a claim must be provided")
	}

	if len(req.ServiceName) < 1 {
		return nil, status.Error(codes.InvalidArgument, "a service name must be provided")
	}

	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(req.Claim, claims, func(token *jwt.Token) (interface{}, error) {
		accessGrant, err := getAccessGrant(ctx, s.configDatabase, claims.Issuer, req.ServiceName)
		if err != nil {
			return nil, err
		}

		return getPublicKeyFromAccessGrant(accessGrant)
	})

	if err != nil {
		validationError, ok := err.(*jwt.ValidationError)
		if !ok {
			s.logger.Error("casting error to jwt validation error failed", zap.Error(err))
			return nil, status.Error(codes.Internal, "unable to verify claim")
		}

		if errors.Is(validationError.Inner, ErrNoAccessGrantAvailable) {
			return nil, status.Error(codes.Unauthenticated, "no access grant available for service")
		}

		if errors.Is(validationError.Inner, ErrNoAccessGrantRevoked) {
			return nil, status.Error(codes.Unauthenticated, "access grant for service has been revoked")
		}

		s.logger.Error("failed to parse jwt", zap.Error(err))

		return nil, status.Error(codes.Internal, "unable to verify claim")
	}

	return &api.VerifyClaimResponse{
		OrderOrganizationName: claims.Issuer,
		OrderReference:        claims.OrderReference,
	}, nil
}

func getPublicKeyFromAccessGrant(grant *database.AccessGrant) (crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(grant.IncomingAccessRequest.PublicKeyPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func getAccessGrant(ctx context.Context, cd database.ConfigDatabase, organizationName, serviceName string) (*database.AccessGrant, error) {
	accessGrant, err := cd.GetLatestAccessGrantForService(ctx, organizationName, serviceName)
	if errIsNotFound(err) {
		return nil, ErrNoAccessGrantAvailable
	}

	if err != nil {
		return nil, err
	}

	if accessGrant.RevokedAt.Valid && accessGrant.RevokedAt.Time.Before(time.Now()) {
		return nil, ErrNoAccessGrantRevoked
	}

	return accessGrant, nil
}
