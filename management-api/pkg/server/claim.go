// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"time"

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

	// TODO sign request via outway
	// Delegatee:      md.OrganizationSerialNumber,
	// OrderReference: req.OrderReference,
	// AccessProof:   convertOutgoingOrderAccessProofsToDelegationAccessProofs(order.OutgoingOrderAccessProofs),
	// ExpiresAt: expiresAt,

	return &external.RequestClaimResponse{
		Claim: "", // TODO get this string via outway SignOrderClaim request
	}, nil
}

func convertOutgoingOrderAccessProofsToDelegationAccessProofs(outgoingOrderAccessProofs []*database.OutgoingOrderAccessProof) []*delegation.AccessProof {
	claimAccessProofs := make([]*delegation.AccessProof, len(outgoingOrderAccessProofs))

	for i, outgoingOrderAccessProof := range outgoingOrderAccessProofs {
		claimAccessProofs[i] = &delegation.AccessProof{
			ServiceName:              outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.ServiceName,
			OrganizationSerialNumber: outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.Organization.SerialNumber,
			PublicKeyFingerprint:     outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.PublicKeyFingerprint,
		}
	}

	return claimAccessProofs
}
