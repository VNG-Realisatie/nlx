// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpc

import (
	"context"
	"encoding/json"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/app/command"
)

func (s *Server) CreateRecord(ctx context.Context, req *api.CreateRecordRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateRecord")

	direction := "out"

	if req.Direction == api.CreateRecordRequest_IN {
		direction = "in"
	}

	dataSubjects := map[string]string{}

	for _, dataSubject := range req.DataSubjects {
		dataSubjects[dataSubject.Key] = dataSubject.Value
	}

	err := s.app.Commands.CreateRecord.Handle(ctx, &command.NewRecordArgs{
		SourceOrganization:      req.SourceOrganization,
		DestinationOrganization: req.DestOrganization,
		Direction:               direction,
		ServiceName:             req.ServiceName,
		OrderReference:          req.OrderReference,
		Delegator:               req.Delegator,
		Data:                    json.RawMessage(req.Data),
		TransactionID:           req.LogrecordID,
		DataSubjects:            dataSubjects,
	})
	if err != nil {
		return nil, ResponseFromError(err)
	}

	return &emptypb.Empty{}, nil
}
