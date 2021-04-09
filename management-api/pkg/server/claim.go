// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"time"

	"github.com/form3tech-oss/jwt-go"
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

	order, err := s.configDatabase.GetOrderByReference(ctx, req.OrderReference)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "order with reference %s not found", req.OrderReference)
		}

		return nil, status.Error(codes.Internal, "failed to find order")
	}

	claims := delegation.JWTClaims{
		Services:       []string{},
		Delegatee:      md.OrganizationName,
		OrderReference: req.OrderReference,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresInHours * time.Hour).Unix(),
			Issuer:    s.orgCert.Certificate().Subject.Organization[0],
		},
	}

	for _, service := range order.Services {
		claims.Services = append(claims.Services, service.ServiceName)
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
