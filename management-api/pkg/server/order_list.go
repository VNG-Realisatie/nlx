// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint:dupl // code is not actually duplicated, the linter has lost it's mind
package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/util/convert"
)

func (s *ManagementService) ListOutgoingOrders(ctx context.Context, _ *emptypb.Empty) (*api.ListOutgoingOrdersResponse, error) {
	s.logger.Info("rpc request ListOutgoingOrders")

	orders, err := s.configDatabase.ListOutgoingOrders(ctx)
	if err != nil {
		s.logger.Error("error getting outgoing orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve outgoing orders")
	}

	outgoingOrders := make([]*api.OutgoingOrder, len(orders))

	for i, order := range orders {
		outgoingOrders[i] = &api.OutgoingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegatee:   order.Delegatee,
			RevokedAt:   convert.SQLToProtoTimestamp(order.RevokedAt),
			ValidFrom:   timestamppb.New(order.ValidFrom),
			ValidUntil:  timestamppb.New(order.ValidUntil),
			Services:    convertOutgoingOrderServices(order.Services),
		}
	}

	return &api.ListOutgoingOrdersResponse{Orders: outgoingOrders}, nil
}

func (s *ManagementService) ListIncomingOrders(ctx context.Context, _ *emptypb.Empty) (*api.ListIncomingOrdersResponse, error) {
	s.logger.Info("rpc request ListIncomingOrders")

	orders, err := s.configDatabase.ListIncomingOrders(ctx)
	if err != nil {
		s.logger.Error("error getting incoming orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve received orders")
	}

	incomingOrders := make([]*api.IncomingOrder, len(orders))

	for i, order := range orders {
		incomingOrders[i] = &api.IncomingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegator:   order.Delegator,
			RevokedAt:   convert.SQLToProtoTimestamp(order.RevokedAt),
			ValidFrom:   timestamppb.New(order.ValidFrom),
			ValidUntil:  timestamppb.New(order.ValidUntil),
			Services:    convertIncomingOrderServices(order.Services),
		}
	}

	return &api.ListIncomingOrdersResponse{Orders: incomingOrders}, nil
}

func (s *ManagementService) ListOrders(ctx context.Context, _ *emptypb.Empty) (*external.ListOrdersResponse, error) {
	metadata, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to parse proxy metadata")
	}

	orders, err := s.configDatabase.ListOutgoingOrdersByOrganization(ctx, metadata.OrganizationName)
	if err != nil {
		s.logger.Error("error getting issued orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve external orders")
	}

	incomingOrders := make([]*api.IncomingOrder, len(orders))

	for i, order := range orders {
		incomingOrders[i] = &api.IncomingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegator:   s.orgCert.Certificate().Subject.Organization[0],
			RevokedAt:   convert.SQLToProtoTimestamp(order.RevokedAt),
			ValidFrom:   timestamppb.New(order.ValidFrom),
			ValidUntil:  timestamppb.New(order.ValidUntil),
			Services:    convertOutgoingOrderServices(order.Services),
		}
	}

	return &external.ListOrdersResponse{Orders: incomingOrders}, nil
}

func convertOutgoingOrderServices(services []database.OutgoingOrderService) []*api.OrderService {
	protoServices := make([]*api.OrderService, len(services))

	for i, service := range services {
		protoServices[i] = &api.OrderService{
			Organization: service.Organization,
			Service:      service.Service,
		}
	}

	return protoServices
}

func convertIncomingOrderServices(services []database.IncomingOrderService) []*api.OrderService {
	protoServices := make([]*api.OrderService, len(services))

	for i, service := range services {
		protoServices[i] = &api.OrderService{
			Organization: service.Organization,
			Service:      service.Service,
		}
	}

	return protoServices
}
