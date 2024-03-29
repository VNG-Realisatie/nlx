// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package server

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
	"go.nlx.io/nlx/management-api/pkg/util/convert"
)

func convertIncomingOrderToDatabase(incomingOrders chan *external.IncomingOrder) []*database.IncomingOrder {
	orders := []*database.IncomingOrder{}

	for order := range incomingOrders {
		services := make([]database.IncomingOrderService, len(order.Services))

		for i, service := range order.Services {
			services[i] = database.IncomingOrderService{
				Organization: database.IncomingOrderServiceOrganization{
					Name:         service.Organization.Name,
					SerialNumber: service.Organization.SerialNumber,
				},
				Service: service.Service,
			}
		}

		orders = append(orders, &database.IncomingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegator:   order.Delegator.SerialNumber,
			RevokedAt:   convert.ProtoToSQLTimestamp(order.RevokedAt),
			ValidFrom:   order.ValidFrom.AsTime(),
			ValidUntil:  order.ValidUntil.AsTime(),
			Services:    services,
		})
	}

	return orders
}

func (s *ManagementService) SynchronizeOrders(ctx context.Context, _ *api.SynchronizeOrdersRequest) (*api.SynchronizeOrdersResponse, error) {
	err := s.authorize(ctx, permissions.SynchronizeIncomingOrders)
	if err != nil {
		return nil, err
	}

	response, err := s.directoryClient.ListOrganizations(ctx, &directoryapi.ListOrganizationsRequest{})
	if err != nil {
		s.logger.Error("failed to list response", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to list response")
	}

	oinToOrgNameHash := convertOrganizationsToHash(response)

	ordersChan := make(chan *external.IncomingOrder)

	wc := &sync.WaitGroup{}
	wc.Add(len(response.Organizations))

	go func() {
		wc.Wait()
		close(ordersChan)
	}()

	for _, organization := range response.Organizations {
		go func(org *directoryapi.Organization) {
			defer wc.Done()

			orders, err := s.fetchOrganizationOrders(ctx, org)
			if err != nil {
				s.logger.Error(
					"unable to synchronize organization orders",
					zap.String("organization-serial-number", org.SerialNumber),
					zap.String("organization-name", org.Name),
					zap.Error(err),
				)

				return
			}

			for _, order := range orders {
				ordersChan <- order
			}
		}(organization)
	}

	orders := convertIncomingOrderToDatabase(ordersChan)

	if len(orders) == 0 {
		return &api.SynchronizeOrdersResponse{Orders: []*external.IncomingOrder{}}, nil
	}

	if err := s.configDatabase.SynchronizeOrders(ctx, orders); err != nil {
		s.logger.Error("failed to synchronize database orders", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to synchronize database orders")
	}

	incomingOrders := make([]*external.IncomingOrder, len(orders))

	for i, order := range orders {
		incomingOrders[i] = &external.IncomingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegator: &external.Organization{
				SerialNumber: order.Delegator,
				Name:         oinToOrgNameHash[order.Delegator],
			},
			RevokedAt:  convert.SQLToProtoTimestamp(order.RevokedAt),
			ValidFrom:  timestamppb.New(order.ValidFrom),
			ValidUntil: timestamppb.New(order.ValidUntil),
			Services:   convertIncomingOrderServices(order.Services, oinToOrgNameHash),
		}
	}

	return &api.SynchronizeOrdersResponse{Orders: incomingOrders}, nil
}

func (s *ManagementService) fetchOrganizationOrders(ctx context.Context, organization *directoryapi.Organization) ([]*external.IncomingOrder, error) {
	inwayProxyAddress, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, organization.SerialNumber)
	if err != nil {
		return nil, err
	}

	externalManagementClient, err := s.createManagementClientFunc(ctx, inwayProxyAddress, s.orgCert)
	if err != nil {
		return nil, err
	}

	defer externalManagementClient.Close()

	response, err := externalManagementClient.ListOrders(ctx, &external.ListOrdersRequest{})
	if err != nil {
		return nil, err
	}

	return response.Orders, nil
}
