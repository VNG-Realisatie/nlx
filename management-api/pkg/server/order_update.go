// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"fmt"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ManagementService) UpdateOutgoingOrder(ctx context.Context, request *api.UpdateOutgoingOrderRequest) (*emptypb.Empty, error) {
	err := s.authorize(ctx, permissions.UpdateOutgoingOrder)
	if err != nil {
		return nil, err
	}

	s.logger.Info("rpc request UpdateOutgoingOrder")

	userInfo, userAgent, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	orderInDB, err := s.configDatabase.GetOutgoingOrderByReference(ctx, request.Reference)
	if err != nil {
		s.logger.Error("failed to fetch order in database", zap.Error(err))

		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "could not find outgoing order in management database")
		}

		return nil, status.Error(codes.Internal, "failed to fetch outgoing order in management database")
	}

	updateOutgoingOrder := &database.UpdateOutgoingOrder{
		ID:             orderInDB.ID,
		Reference:      request.Reference,
		Description:    request.Description,
		PublicKeyPEM:   request.PublicKeyPEM,
		ValidFrom:      request.ValidFrom.AsTime(),
		ValidUntil:     request.ValidUntil.AsTime(),
		AccessProofIds: request.AccessProofIds,
	}

	if errValidation := validateUpdateOutgoingOrder(updateOutgoingOrder); errValidation != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid update outgoing order: %s", errValidation))
	}

	accessProofs, err := s.configDatabase.GetAccessProofs(ctx, updateOutgoingOrder.AccessProofIds)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not retrieve access proofs")
	}

	err = s.auditLogger.OrderOutgoingUpdate(ctx, userInfo.Email, userAgent, orderInDB.Delegatee, orderInDB.Reference, accessProofsToAuditLogRecordServices(accessProofs))
	if err != nil {
		s.logger.Error("failed to write auditlog", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to write to auditlog")
	}

	if err := s.configDatabase.UpdateOutgoingOrder(ctx, updateOutgoingOrder); err != nil {
		s.logger.Error("failed to update outgoing order", zap.Error(err))

		return nil, status.Errorf(codes.Internal, "failed to update outgoing order")
	}

	return &emptypb.Empty{}, nil
}

func accessProofsToAuditLogRecordServices(accessProofs []*database.AccessProof) []auditlog.RecordService {
	services := make([]auditlog.RecordService, len(accessProofs))

	for i, a := range accessProofs {
		if a.OutgoingAccessRequest != nil {
			services[i] = auditlog.RecordService{
				Organization: auditlog.RecordServiceOrganization{
					SerialNumber: a.OutgoingAccessRequest.Organization.SerialNumber,
					Name:         a.OutgoingAccessRequest.Organization.Name,
				},
				Service: a.OutgoingAccessRequest.ServiceName,
			}
		}
	}

	return services
}
