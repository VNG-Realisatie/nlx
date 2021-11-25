package server

import (
	"context"
	"errors"
	"fmt"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ManagementService) UpdateOutgoingOrder(ctx context.Context, request *api.OutgoingOrderRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request UpdateOutgoingOrder")

	order := convertOutgoingOrder(request)

	if err := validateOutgoingOrder(order); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid outgoing order: %s", err))
	}

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	services := []auditlog.RecordService{}

	for _, service := range order.Services {
		services = append(services, auditlog.RecordService{
			Organization: auditlog.RecordServiceOrganization{
				SerialNumber: service.Organization.SerialNumber,
				Name:         service.Organization.Name,
			},
			Service: service.Service,
		})
	}

	orderInDB, err := s.configDatabase.GetOutgoingOrderByReference(ctx, order.Reference)
	if err != nil {
		s.logger.Error("failed to fetch order in database", zap.Error(err))

		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "could not find outgoing order in management database")
		}
		return nil, status.Error(codes.Internal, "failed to fetch outgoing order in management database")
	}

	order.ID = orderInDB.ID

	err = s.auditLogger.OrderOutgoingUpdate(ctx, userInfo.username, userInfo.userAgent, order.Delegatee, order.Reference, services)
	if err != nil {
		s.logger.Error("failed to write auditlog", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to write to auditlog")
	}

	if err := s.configDatabase.UpdateOutgoingOrder(ctx, order); err != nil {
		s.logger.Error("failed to create outgoing order", zap.Error(err))

		return nil, status.Errorf(codes.Internal, "failed to update outgoing order")
	}

	return &emptypb.Empty{}, nil
}
