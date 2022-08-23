// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain"
)

//nolint:gocyclo // complexity will be reduced once we simplify the domain
func (s *TXLogService) CreateRecord(ctx context.Context, req *api.CreateRecordRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateRecord")

	sourceOrg, err := domain.NewOrganization(req.SourceOrganization)
	if err != nil {
		s.logger.Error("error creating source org", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("source organization: %s", err))
	}

	destinationOrg, err := domain.NewOrganization(req.DestOrganization)
	if err != nil {
		s.logger.Error("error creating destination org", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("destination organization: %s", err))
	}

	service, err := domain.NewService(req.ServiceName)
	if err != nil {
		s.logger.Error("error creating service", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("service: %s", err))
	}

	var order *domain.Order

	if len(req.Delegator) > 0 && len(req.OrderReference) > 0 {
		newOrder, orderErr := domain.NewOrder(&domain.NewOrderArgs{
			Delegator: req.Delegator,
			Reference: req.OrderReference,
		})
		if orderErr != nil {
			s.logger.Error("failed to create order model", zap.Error(err))
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("order: %s", err))
		}

		order = newOrder
	}

	direction := domain.OUT

	if req.Direction == api.CreateRecordRequest_IN {
		direction = domain.IN
	}

	dataSubjects := map[string]string{}

	for _, dataSubject := range req.DataSubjects {
		dataSubjects[dataSubject.Key] = dataSubject.Value
	}

	record, err := domain.NewRecord(&domain.NewRecordArgs{
		Source:        sourceOrg,
		Destination:   destinationOrg,
		Direction:     direction,
		Service:       service,
		Order:         order,
		Data:          json.RawMessage(req.Data),
		TransactionID: req.LogrecordID,
		CreatedAt:     s.clock.Now(),
		DataSubjects:  dataSubjects,
	})
	if err != nil {
		s.logger.Error("failed to instantiate record model", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	err = s.storage.CreateRecord(ctx, record)
	if err != nil {
		s.logger.Error("failed to create record model", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	return &emptypb.Empty{}, nil
}
