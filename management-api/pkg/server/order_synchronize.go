package server

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) SynchronizeOrders(ctx context.Context, _ *emptypb.Empty) (*api.SynchronizeOrdersResponse, error) {
	response, err := s.directoryClient.ListOrganizations(ctx, &emptypb.Empty{})
	if err != nil {
		s.logger.Error("failed to list response", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to list response")
	}

	ordersChan := make(chan *api.IncomingOrder)

	wc := &sync.WaitGroup{}
	wc.Add(len(response.Organizations))

	go func() {
		wc.Wait()
		close(ordersChan)
	}()

	for _, organization := range response.Organizations {
		go func(org *inspectionapi.ListOrganizationsResponse_Organization) {
			defer wc.Done()

			orders, err := s.fetchOrganizationOrders(ctx, org)
			if err != nil {
				s.logger.Error(
					"unable to synchronize organization orders",
					zap.String("organization", org.Name),
					zap.Error(err),
				)

				return
			}

			for _, order := range orders {
				ordersChan <- order
			}
		}(organization)
	}

	orders := []*database.IncomingOrder{}

	for order := range ordersChan {
		services := make([]database.IncomingOrderService, len(order.Services))

		for i, service := range order.Services {
			services[i] = database.IncomingOrderService{
				Organization: service.Organization,
				Service:      service.Service,
			}
		}

		orders = append(orders, &database.IncomingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegator:   order.Delegator,
			ValidFrom:   order.ValidFrom.AsTime(),
			ValidUntil:  order.ValidUntil.AsTime(),
			Services:    services,
		})
	}

	if len(orders) == 0 {
		return &api.SynchronizeOrdersResponse{Orders: []*api.IncomingOrder{}}, nil
	}

	if err := s.configDatabase.SynchronizeOrders(ctx, orders); err != nil {
		s.logger.Error("failed to synchronize database orders", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to synchronize database orders")
	}

	incomingOrders := make([]*api.IncomingOrder, len(orders))

	for i, order := range orders {
		incomingOrders[i] = &api.IncomingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegator:   order.Delegator,
			ValidFrom:   timestamppb.New(order.ValidFrom),
			ValidUntil:  timestamppb.New(order.ValidUntil),
			Services:    convertIncomingOrderServices(order.Services),
		}
	}

	return &api.SynchronizeOrdersResponse{Orders: incomingOrders}, nil
}

func (s *ManagementService) fetchOrganizationOrders(ctx context.Context, organization *inspectionapi.ListOrganizationsResponse_Organization) ([]*api.IncomingOrder, error) {
	inwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, organization.Name)
	if err != nil {
		return nil, err
	}

	externalManagementClient, err := s.createManagementClientFunc(ctx, inwayProxyAddress, s.orgCert)
	if err != nil {
		return nil, err
	}

	defer externalManagementClient.Close()

	response, err := externalManagementClient.ListOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return response.Orders, nil
}
