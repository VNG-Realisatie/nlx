// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

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
)

func (s *ManagementService) ListIssuedOrders(ctx context.Context, _ *emptypb.Empty) (*api.ListIssuedOrdersResponse, error) {
	s.logger.Info("rpc request ListIssuedOrders")

	orders, err := s.configDatabase.ListIssuedOrders(ctx)
	if err != nil {
		s.logger.Error("error getting issued orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve issued orders")
	}

	ordersResponse := convertOrderToProto(orders)

	return &api.ListIssuedOrdersResponse{Orders: ordersResponse}, nil
}

func (s *ManagementService) ListOrders(ctx context.Context, _ *emptypb.Empty) (*external.ListOrdersResponse, error) {
	metadata, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to parse proxy metadata")
	}

	orders, err := s.configDatabase.ListOrdersByOrganization(ctx, metadata.OrganizationName)
	if err != nil {
		s.logger.Error("error getting issued orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve issued orders")
	}

	ordersResponse := convertOrderToProto(orders)

	return &external.ListOrdersResponse{Orders: ordersResponse}, nil
}

func convertOrderToProto(orders []*database.Order) []*api.Order {
	result := make([]*api.Order, len(orders))

	for i, order := range orders {
		services := make([]*api.Order_Service, len(order.Services))

		for j, service := range order.Services {
			services[j] = &api.Order_Service{
				Service:      service.Service,
				Organization: service.Organization,
			}
		}

		result[i] = &api.Order{
			Reference:   order.Reference,
			Description: order.Description,
			Delegatee:   order.Delegatee,
			Services:    services,
			ValidFrom:   timestamppb.New(order.ValidFrom),
			ValidUntil:  timestamppb.New(order.ValidUntil),
		}
	}

	return result
}
