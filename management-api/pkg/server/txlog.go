// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/txlog"
	txlogapi "go.nlx.io/nlx/txlog-api/api"
)

type TXLogService struct {
	api.UnimplementedTXLogServer

	logger          *zap.Logger
	txlogClient     txlog.Client
	directoryClient directory.Client
}

func NewTXLogService(logger *zap.Logger, txlogClient txlog.Client, directoryClient directory.Client) *TXLogService {
	return &TXLogService{
		logger:          logger,
		txlogClient:     txlogClient,
		directoryClient: directoryClient,
	}
}

func (s *ManagementService) IsTXLogEnabled(context.Context, *api.IsTXLogEnabledRequest) (*api.IsTXLogEnabledResponse, error) {
	return &api.IsTXLogEnabledResponse{
		Enabled: s.txlogClient != nil,
	}, nil
}

func (s *TXLogService) ListRecords(ctx context.Context, _ *api.TXLogListRecordsRequest) (*api.TXLogListRecordsResponse, error) {
	s.logger.Info("rpc request ListRecords")

	organizations, err := s.directoryClient.ListOrganizations(ctx, &directoryapi.ListOrganizationsRequest{})
	if err != nil {
		s.logger.Error("failed to retrieve organizations from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "txlog error")
	}

	oinToOrgNameHash := convertOrganizationsToHash(organizations)

	resp, err := s.txlogClient.ListRecords(ctx, &txlogapi.ListRecordsRequest{})
	if err != nil {
		s.logger.Error("error getting records list from txlog", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "txlog error")
	}

	records := make([]*api.TXLogRecord, len(resp.Records))

	for i, r := range resp.Records {
		var order *api.TXLogOrder
		if r.Order != nil {
			order = &api.TXLogOrder{
				Delegator: &api.TXLogOrganization{
					SerialNumber: r.Order.Delegator,
					Name:         oinToOrgNameHash[r.Order.Delegator],
				},
				Reference: r.Order.Reference,
			}
		}

		records[i] = &api.TXLogRecord{
			Source: &api.TXLogOrganization{
				SerialNumber: r.Source.SerialNumber,
				Name:         oinToOrgNameHash[r.Source.SerialNumber],
			},
			Destination: &api.TXLogOrganization{
				SerialNumber: r.Destination.SerialNumber,
				Name:         oinToOrgNameHash[r.Destination.SerialNumber],
			},
			Direction: api.TXLogDirection(r.Direction),
			Service: &api.TXLogService{
				Name: r.Service.Name,
			},
			Order:         order,
			Data:          r.Data,
			TransactionId: r.TransactionId,
			CreatedAt:     r.CreatedAt,
		}
	}

	return &api.TXLogListRecordsResponse{Records: records}, nil
}
