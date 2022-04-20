// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
	"go.nlx.io/nlx/management-api/pkg/outway"
	outwayapi "go.nlx.io/nlx/outway/api"
)

const expiresInHours = 4

var errMessageOutwayUnableToSignClaim = "could not sign order claim via outway"

func (s *ManagementService) RequestClaim(ctx context.Context, req *external.RequestClaimRequest) (*external.RequestClaimResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, grpcerrors.NewFromValidationError(err)
	}

	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		s.logger.Error("failed to parse proxy metadata", zap.Error(err))
		return nil, err
	}

	order, err := s.configDatabase.GetOutgoingOrderByReference(ctx, req.OrderReference)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, grpcerrors.New(codes.NotFound, external.ErrorReason_ORDER_NOT_FOUND, "order not found", nil)
		}

		return nil, grpcerrors.NewInternal("failed to find order", nil)
	}

	if order.Delegatee != md.OrganizationSerialNumber {
		return nil, grpcerrors.New(codes.PermissionDenied, external.ErrorReason_ORDER_NOT_FOUND_FOR_ORG, "order does not exist for your organization", nil)
	}

	outgoingAccessRequest := filterOutgoingAccessRequestFromOrder(order, req.ServiceOrganizationSerialNumber, req.ServiceName)
	if outgoingAccessRequest == nil {
		return nil, grpcerrors.New(codes.NotFound, external.ErrorReason_ORDER_DOES_NOT_CONTAIN_SERVICE, "service not found in order", nil)
	}

	if order.RevokedAt.Valid {
		return nil, grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ORDER_REVOKED, "order is revoked", nil)
	}

	if time.Now().After(order.ValidUntil) {
		return nil, grpcerrors.New(codes.Unauthenticated, external.ErrorReason_ORDER_EXPIRED, "order is expired", nil)
	}

	expiresAt := time.Now().Add(expiresInHours * time.Hour)

	if expiresAt.After(order.ValidUntil) {
		expiresAt = order.ValidUntil
	}

	delegateeFingerprint, err := tls.PemPublicKeyFingerprint([]byte(order.PublicKeyPEM))
	if err != nil {
		s.logger.Error("invalid public key format", zap.Error(err))
		return nil, grpcerrors.NewInternal("invalid public key format", nil)
	}

	outways, err := s.configDatabase.GetOutwaysByPublicKeyFingerprint(ctx, outgoingAccessRequest.PublicKeyFingerprint)
	if err != nil {
		s.logger.Error("could not find outway", zap.Error(err))
		return nil, grpcerrors.NewInternal("could not find outway", nil)
	}

	outwayClient := createOutwayClient(ctx, s, outways)
	if outwayClient == nil {
		s.logger.Error("could not connect to outway", zap.Error(err))
		return nil, grpcerrors.NewInternal("could not connect to outway", nil)
	}

	signedOrderClaimResp, err := outwayClient.SignOrderClaim(ctx, &outwayapi.SignOrderClaimRequest{
		Delegatee:                     md.OrganizationSerialNumber,
		DelegateePublicKeyFingerprint: delegateeFingerprint,
		OrderReference:                req.OrderReference,
		AccessProof: &outwayapi.AccessProof{
			ServiceName:              outgoingAccessRequest.ServiceName,
			OrganizationSerialNumber: s.orgCert.GetOrganizationInfo().SerialNumber,
			PublicKeyFingerprint:     outgoingAccessRequest.PublicKeyFingerprint,
		},
		ExpiresAt: timestamppb.New(expiresAt),
	})
	if err != nil {
		s.logger.Error("could not sign order claim via outway", zap.Error(err))
		return nil, grpcerrors.NewInternal("could not sign order claim via outway", nil)
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
