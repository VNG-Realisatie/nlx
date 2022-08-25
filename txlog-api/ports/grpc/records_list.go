// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpc

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

func (s *Server) ListRecords(ctx context.Context, _ *emptypb.Empty) (*api.ListRecordsResponse, error) {
	s.logger.Info("rpc request ListRecords")

	records, err := s.app.Queries.ListRecords.Handle(ctx, maxAmountOfRecords)
	if err != nil {
		s.logger.Error("error getting record list from storage", zap.Error(err))
		return nil, status.Error(codes.Internal, "storage error")
	}

	response := dataModelToResponse(records)

	return response, nil
}

func dataModelToResponse(records []*domain.Record) *api.ListRecordsResponse {
	response := &api.ListRecordsResponse{}
	response.Records = make([]*api.ListRecordsResponse_Record, len(records))

	for i, r := range records {
		record := &api.ListRecordsResponse_Record{
			Source: &api.ListRecordsResponse_Record_Organization{
				SerialNumber: r.SourceOrganization(),
			},
			Destination: &api.ListRecordsResponse_Record_Organization{
				SerialNumber: r.DestinationOrganization(),
			},
			Direction: api.ListRecordsResponse_Record_Direction(api.ListRecordsResponse_Record_Direction_value[strings.ToUpper(string(r.Direction()))]),
			Service: &api.ListRecordsResponse_Record_Service{
				Name: r.ServiceName(),
			},
			Order: &api.ListRecordsResponse_Record_Order{
				Delegator: r.Delegator(),
				Reference: r.OrderReference(),
			},
			Data:          string(r.Data()),
			TransactionID: r.TransactionID(),
			CreatedAt:     timestamppb.New(r.CreatedAt()),
		}
		response.Records[i] = record
	}

	return response
}
