// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpc

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/app/query"
)

const maxAmountOfRecords = 100

func (s *Server) ListRecords(ctx context.Context, _ *api.ListRecordsRequest) (*api.ListRecordsResponse, error) {
	s.logger.Info("rpc request ListRecords")

	records, err := s.app.Queries.ListRecords.Handle(ctx, maxAmountOfRecords)
	if err != nil {
		s.logger.Error("error getting record list from storage", err)
		return nil, status.Error(codes.Internal, "storage error")
	}

	response := dataModelToResponse(records)

	return response, nil
}

func dataModelToResponse(records []*query.Record) *api.ListRecordsResponse {
	response := &api.ListRecordsResponse{}
	response.Records = make([]*api.ListRecordsResponse_Record, len(records))

	for i, r := range records {
		recordResponse := &api.ListRecordsResponse_Record{
			Source: &api.ListRecordsResponse_Record_Organization{
				SerialNumber: r.SourceOrganization,
			},
			Destination: &api.ListRecordsResponse_Record_Organization{
				SerialNumber: r.DestinationOrganization,
			},
			Direction: directionToProto(r.Direction),
			Service: &api.ListRecordsResponse_Record_Service{
				Name: r.ServiceName,
			},
			Order: &api.ListRecordsResponse_Record_Order{
				Delegator: r.Delegator,
				Reference: r.OrderReference,
			},
			Data:          string(r.Data),
			TransactionId: r.TransactionID,
			CreatedAt:     timestamppb.New(r.CreatedAt),
		}
		response.Records[i] = recordResponse
	}

	return response
}

func directionToProto(direction string) api.ListRecordsResponse_Record_Direction {
	val, ok := api.ListRecordsResponse_Record_Direction_value[fmt.Sprintf("DIRECTION_%s", strings.ToUpper(direction))]
	if !ok {
		return api.ListRecordsResponse_Record_DIRECTION_UNSPECIFIED
	}

	return api.ListRecordsResponse_Record_Direction(val)
}
