// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpc

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

func (s *Server) CreateRecord(ctx context.Context, req *api.CreateRecordRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateRecord")

	direction := domain.OUT

	if req.Direction == api.CreateRecordRequest_IN {
		direction = domain.IN
	}

	dataSubjects := map[string]string{}

	for _, dataSubject := range req.DataSubjects {
		dataSubjects[dataSubject.Key] = dataSubject.Value
	}

	record, err := domain.NewRecord(&domain.NewRecordArgs{
		SourceOrganization:      req.SourceOrganization,
		DestinationOrganization: req.DestOrganization,
		Direction:               direction,
		ServiceName:             req.ServiceName,
		OrderReference:          req.OrderReference,
		Delegator:               req.Delegator,
		Data:                    json.RawMessage(req.Data),
		TransactionID:           req.LogrecordID,
		CreatedAt:               s.clock.Now(),
		DataSubjects:            dataSubjects,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid record: %s", err))
	}

	err = s.app.Commands.CreateRecord.Handle(ctx, record)
	if err != nil {
		s.logger.Error("failed to create record model", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	return &emptypb.Empty{}, nil
}
