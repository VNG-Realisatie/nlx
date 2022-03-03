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
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/outway"
	outwayapi "go.nlx.io/nlx/outway/api"
)

const expiresInHours = 4

var errMessageOrderRevoked = "order is revoked"
var errMessageOutwayUnableToSignClaim = "could not sign order claim via outway"

func (s *ManagementService) RequestClaim(ctx context.Context, req *external.RequestClaimRequest) (*external.RequestClaimResponse, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		s.logger.Error("failed to parse proxy metadata", zap.Error(err))
		return nil, err
	}

	if len(req.OrderReference) < 1 {
		return nil, status.Error(codes.InvalidArgument, "an order reference must be provided")
	}

	if req.ServiceName == "" {
		return nil, status.Error(codes.InvalidArgument, "a service name must be provided")
	}

	if req.ServiceOrganizationSerialNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "an organization serial number must be provided")
	}

	order, err := s.configDatabase.GetOutgoingOrderByReference(ctx, req.OrderReference)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "order with reference '%s' not found", req.OrderReference)
		}

		return nil, status.Error(codes.Internal, "failed to find order")
	}

	if order.Delegatee != md.OrganizationSerialNumber {
		return nil, status.Errorf(codes.NotFound, "order with reference '%s' and organization serial number '%s' not found", req.OrderReference, md.OrganizationSerialNumber)
	}

	outgoingAccessRequest := filterOutgoingAccessRequestFromOrder(order, req.ServiceOrganizationSerialNumber, req.ServiceName)
	if outgoingAccessRequest == nil {
		return nil, status.Errorf(codes.NotFound, "order with reference '%s' and organization serial number '%s' and service name '%s' not found", req.OrderReference, md.OrganizationSerialNumber, req.ServiceName)
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

	outways, err := s.configDatabase.GetOutwaysByPublicKeyFingerprint(ctx, outgoingAccessRequest.PublicKeyFingerprint)
	if err != nil {
		s.logger.Error("could not find outway", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not find outway")
	}

	outwayClient := createOutwayClient(ctx, s, outways)
	if outwayClient == nil {
		s.logger.Error("could not connect to outway", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not connect to outway")
	}

	signedOrderClaimResp, err := outwayClient.SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
		Delegatee:      md.OrganizationSerialNumber,
		OrderReference: req.OrderReference,
		AccessProof: &outwayapi.AccessProof{
			ServiceName:              outgoingAccessRequest.ServiceName,
			OrganizationSerialNumber: s.orgCert.GetOrganizationInfo().SerialNumber,
			PublicKeyFingerprint:     outgoingAccessRequest.PublicKeyFingerprint,
		},
		ExpiresAt: timestamppb.New(expiresAt),
	})
	if err != nil {
		s.logger.Error("could not sign order claim via outway", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not sign order claim via outway")
	}

	return &external.RequestClaimResponse{
		Claim: signedOrderClaimResp.SignedOrderclaim,
	}, nil
}

func createOutwayClient(ctx context.Context, s *ManagementService, outways []*database.Outway) outway.Client {
	for _, o := range outways {
		c, err := s.createOutwayClientFunc(ctx, o.SelfAddressAPI, s.internalCert)

		if err != nil {
			s.logger.Error("error creating outway client", zap.Error(err))
			continue
		}

		return c
	}

	return nil
}

func filterOutgoingAccessRequestFromOrder(order *database.OutgoingOrder, providerSerialNumber, serviceName string) *database.OutgoingAccessRequest {
	for _, a := range order.OutgoingOrderAccessProofs {
		if a.AccessProof.OutgoingAccessRequest.ServiceName == serviceName && a.AccessProof.OutgoingAccessRequest.Organization.SerialNumber == providerSerialNumber {
			return a.AccessProof.OutgoingAccessRequest
		}
	}

	return nil
}
