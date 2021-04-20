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
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) ListIssuedOrders(ctx context.Context, _ *emptypb.Empty) (*api.ListIssuedOrdersResponse, error) {
	s.logger.Info("rpc request ListIssuedOrders")

	orders, err := s.configDatabase.ListIssuedOrders(ctx)
	if err != nil {
		s.logger.Error("error getting issued orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve issued orders")
	}

	ordersResponse := convertFromDatabaseToView(orders)

	return &api.ListIssuedOrdersResponse{Orders: ordersResponse}, nil
}

func convertFromDatabaseToView(orders []*database.Order) []*api.ListIssuedOrdersResponse_Order {
	result := make([]*api.ListIssuedOrdersResponse_Order, len(orders))

	for i, order := range orders {
		orderResponse := &api.ListIssuedOrdersResponse_Order{
			Reference:   order.Reference,
			Description: order.Description,
			Delegatee:   order.Delegatee,
			ValidFrom:   timestamppb.New(order.ValidFrom),
			ValidUntil:  timestamppb.New(order.ValidUntil),
		}

		services := make([]*api.ListIssuedOrdersResponse_Order_Service, len(order.Services))

		for j, service := range order.Services {
			services[j] = &api.ListIssuedOrdersResponse_Order_Service{
				Service:      service.Service,
				Organization: service.Organization,
			}
		}

		orderResponse.Services = services

		result[i] = orderResponse
	}

	return result
}
