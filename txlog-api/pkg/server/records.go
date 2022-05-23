// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain"
)

const maxAmountOfRecords = 100

func (s *TXLogService) ListRecords(ctx context.Context, _ *emptypb.Empty) (*api.ListRecordsResponse, error) {
	s.logger.Info("rpc request ListRecords")

	records, err := s.storage.ListRecords(ctx, maxAmountOfRecords)
	if err != nil {
		s.logger.Error("error getting record list from storage", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	response := &api.ListRecordsResponse{}
	response.Records = make([]*api.Record, len(records))

	for i, r := range records {
		response.Records[i] = convertFromDatabaseRecord(r)
	}

	return response, nil
}

func convertFromDatabaseRecord(m *domain.Record) *api.Record {
	record := &api.Record{
		Source: &api.Organization{
			SerialNumber: m.Source().SerialNumber(),
		},
		Destination: &api.Organization{
			SerialNumber: m.Destination().SerialNumber(),
		},
		Direction: api.Direction(api.Direction_value[strings.ToUpper(string(m.Direction()))]),
		Service: &api.Service{
			Name: m.Service().Name(),
		},
		Order: &api.Order{
			Delegator: m.Order().Delegator(),
			Reference: m.Order().Reference(),
		},
		Data:          string(m.Data()),
		TransactionID: m.TransactionID(),
		CreatedAt:     timestamppb.New(m.CreatedAt()),
	}

	return record
}
