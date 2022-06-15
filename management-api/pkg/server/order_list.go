// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint:dupl // code is not actually duplicated, the linter has lost it's mind
package server

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/util/convert"
)

func (s *ManagementService) ListOutgoingOrders(ctx context.Context, _ *emptypb.Empty) (*api.ListOutgoingOrdersResponse, error) {
	s.logger.Info("rpc request ListOutgoingOrders")

	participants, err := s.directoryClient.ListParticipants(ctx, &emptypb.Empty{})
	if err != nil {
		s.logger.Error("error getting participants from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	oinToOrgNameHash := convertParticipantsToHash(participants)

	orders, err := s.configDatabase.ListOutgoingOrders(ctx)
	if err != nil {
		s.logger.Error("error getting outgoing orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve outgoing orders")
	}

	outgoingOrders := make([]*api.OutgoingOrder, len(orders))

	for i, order := range orders {
		accessProofs := make([]*api.AccessProof, len(order.OutgoingOrderAccessProofs))

		for j, outgoingOrderAccessProofs := range order.OutgoingOrderAccessProofs {
			outgoingOrderAccessProofs.AccessProof.OutgoingAccessRequest.Organization.Name = oinToOrgNameHash[outgoingOrderAccessProofs.AccessProof.OutgoingAccessRequest.Organization.SerialNumber]
			accessProofs[j] = convertAccessProof(outgoingOrderAccessProofs.AccessProof)
		}

		outgoingOrders[i] = &api.OutgoingOrder{
			Reference:    order.Reference,
			PublicKeyPem: order.PublicKeyPEM,
			Description:  order.Description,
			Delegatee: &api.Organization{
				SerialNumber: order.Delegatee,
				Name:         oinToOrgNameHash[order.Delegatee],
			},
			RevokedAt:    convert.SQLToProtoTimestamp(order.RevokedAt),
			ValidFrom:    timestamppb.New(order.ValidFrom),
			ValidUntil:   timestamppb.New(order.ValidUntil),
			AccessProofs: accessProofs,
		}
	}

	return &api.ListOutgoingOrdersResponse{Orders: outgoingOrders}, nil
}

func (s *ManagementService) ListIncomingOrders(ctx context.Context, _ *emptypb.Empty) (*api.ListIncomingOrdersResponse, error) {
	s.logger.Info("rpc request ListIncomingOrders")

	participants, err := s.directoryClient.ListParticipants(ctx, &emptypb.Empty{})
	if err != nil {
		s.logger.Error("error getting participants from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	oinToOrgNameHash := convertParticipantsToHash(participants)

	orders, err := s.configDatabase.ListIncomingOrders(ctx)
	if err != nil {
		s.logger.Error("error getting incoming orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve received orders")
	}

	incomingOrders := make([]*api.IncomingOrder, len(orders))

	for i, order := range orders {
		var revokedAt *timestamppb.Timestamp

		if order.RevokedAt() != nil {
			revokedAt = timestamppb.New(*order.RevokedAt())
		}

		incomingOrders[i] = &api.IncomingOrder{
			Reference:   order.Reference(),
			Description: order.Description(),
			Delegator: &api.Organization{
				SerialNumber: order.Delegator(),
				Name:         oinToOrgNameHash[order.Delegator()],
			},
			RevokedAt:  revokedAt,
			ValidFrom:  timestamppb.New(order.ValidFrom()),
			ValidUntil: timestamppb.New(order.ValidUntil()),
			Services:   convertDomainIncomingOrderServices(order.Services(), oinToOrgNameHash),
		}
	}

	return &api.ListIncomingOrdersResponse{Orders: incomingOrders}, nil
}

func (s *ManagementService) ListOrders(ctx context.Context, _ *emptypb.Empty) (*external.ListOrdersResponse, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to parse proxy metadata")
	}

	orders, err := s.configDatabase.ListOutgoingOrdersByOrganization(ctx, md.OrganizationSerialNumber)
	if err != nil {
		s.logger.Error("error getting issued orders from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to retrieve external orders")
	}

	incomingOrders := make([]*api.IncomingOrder, len(orders))

	for i, order := range orders {
		incomingOrders[i] = &api.IncomingOrder{
			Reference:   order.Reference,
			Description: order.Description,
			Delegator: &api.Organization{
				SerialNumber: s.orgCert.Certificate().Subject.SerialNumber,
			},
			RevokedAt:  convert.SQLToProtoTimestamp(order.RevokedAt),
			ValidFrom:  timestamppb.New(order.ValidFrom),
			ValidUntil: timestamppb.New(order.ValidUntil),
			Services:   convertOutgoingAccessProofsToOrderServices(order.OutgoingOrderAccessProofs),
		}
	}

	return &external.ListOrdersResponse{Orders: incomingOrders}, nil
}

func convertIncomingOrderServices(services []database.IncomingOrderService, oinToOrgNameHash map[string]string) []*api.OrderService {
	protoServices := make([]*api.OrderService, len(services))

	for i, service := range services {
		protoServices[i] = &api.OrderService{
			Organization: &api.Organization{
				SerialNumber: service.Organization.SerialNumber,
				Name:         oinToOrgNameHash[service.Organization.SerialNumber],
			},
			Service: service.Service,
		}
	}

	return protoServices
}

func convertOutgoingAccessProofsToOrderServices(outgoingOrderAccessProofs []*database.OutgoingOrderAccessProof) []*api.OrderService {
	orderServices := make(map[string]*api.OrderService)
	protoServices := make([]*api.OrderService, 0)

	for _, outgoingOrderAccessProof := range outgoingOrderAccessProofs {
		orderServiceKey := fmt.Sprintf("%s.%s", outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.Organization.SerialNumber, outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.ServiceName)

		_, ok := orderServices[orderServiceKey]
		if ok {
			continue
		}

		orderService := &api.OrderService{
			Organization: &api.Organization{
				SerialNumber: outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.Organization.SerialNumber,
				Name:         outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.Organization.Name,
			},
			Service: outgoingOrderAccessProof.AccessProof.OutgoingAccessRequest.ServiceName,
		}

		orderServices[orderServiceKey] = orderService
		protoServices = append(protoServices, orderService)
	}

	return protoServices
}
func convertDomainIncomingOrderServices(services []domain.IncomingOrderService, oinToOrgNameHash map[string]string) []*api.OrderService {
	protoServices := make([]*api.OrderService, len(services))

	for i, service := range services {
		protoServices[i] = &api.OrderService{
			Organization: &api.Organization{
				SerialNumber: service.OrganizationSerialNumber(),
				Name:         oinToOrgNameHash[service.OrganizationSerialNumber()],
			},
			Service: service.Service(),
		}
	}

	return protoServices
}
