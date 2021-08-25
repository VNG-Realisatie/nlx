// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint dupl: we have both incoming and outgoing orders
func (s *ManagementService) RevokeOutgoingOrder(ctx context.Context, request *api.RevokeOutgoingOrderRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request RevokeOutgoingOrder")

	if request.Delegatee == "" {
		return nil, status.Error(codes.InvalidArgument, "delegatee is required")
	}

	if request.Reference == "" {
		return nil, status.Error(codes.InvalidArgument, "reference is required")
	}

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.OrderOutgoingRevoke(ctx, userInfo.username, userInfo.userAgent, request.Delegatee, request.Reference)
	if err != nil {
		s.logger.Error("failed to write auditlog", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to write to auditlog")
	}

	if err := s.configDatabase.RevokeOutgoingOrderByReference(ctx, request.Delegatee, request.Reference, time.Now()); err != nil {
		s.logger.Error("failed to revoke outgoing order", zap.Error(err))

		if err == database.ErrNotFound {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("outgoing order with delegatee %s and reference %s does not exist", request.Delegatee, request.Reference))
		}

		return nil, status.Errorf(codes.Internal, "failed to revoke outgoing order")
	}

	return &emptypb.Empty{}, nil
}
