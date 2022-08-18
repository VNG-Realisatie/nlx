// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain"
)

func (s *TXLogService) CreateRecord(ctx context.Context, req *api.CreateRecordRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateRecord")

	sourceOrg, err := domain.NewOrganization(req.SourceOrganization)
	if err != nil {
		s.logger.Error("error creating source org", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	destinationOrg, err := domain.NewOrganization(req.DestOrganization)
	if err != nil {
		s.logger.Error("error creating destination org", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	service, err := domain.NewService(req.ServiceName)
	if err != nil {
		s.logger.Error("error creating service", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	order, err := domain.NewOrder(&domain.NewOrderArgs{
		Delegator: req.Delegator,
		Reference: req.OrderReference,
	})
	if err != nil {
		s.logger.Error("failed to create order model", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	direction := domain.OUT

	if req.Direction == api.CreateRecordRequest_IN {
		direction = domain.IN
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
		DataSubjects:  req.DataSubjects,
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
