// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"bytes"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
)

func (s *ManagementService) IsFinanceEnabled(ctx context.Context, request *emptypb.Empty) (*api.IsFinanceEnabledResponse, error) {
	return &api.IsFinanceEnabledResponse{
		Enabled: s.txlogDatabase != nil,
	}, nil
}

func (s *ManagementService) DownloadFinanceExport(ctx context.Context, request *emptypb.Empty) (*api.DownloadFinanceExportResponse, error) {
	services, err := s.configDatabase.ListServices(ctx)
	if err != nil {
		s.logger.Error("failed to list services", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	servicesMap := map[string]*database.Service{}

	for _, service := range services {
		servicesMap[service.Name] = service
	}

	records := []txlogdb.Record{}

	if s.txlogDatabase != nil {
		records, err = s.txlogDatabase.FilterRecords(
			ctx,
			&txlogdb.Filters{
				Destination: s.orgCert.Certificate().Subject.Organization[0],
				Direction:   transactionlog.DirectionIn,
			},
		)
		if err != nil {
			s.logger.Error("failed to filter records", zap.Error(err))

			return nil, status.Error(codes.Internal, "database error")
		}
	}

	buff := &bytes.Buffer{}
	exporter := NewFinanceCSVExporter(buff)

	for _, record := range records {
		svc := &database.Service{}

		if match, ok := servicesMap[record.ServiceName]; ok {
			svc = match
		}

		if err := exporter.Export(&FinanceRow{
			Organization:     record.Source,
			ServiceName:      record.ServiceName,
			Month:            record.CreatedAt.Format("01 2006"),
			PricePerRequest:  svc.RequestCosts,
			NumberOfRequests: record.RequestCount,
			SetupCosts:       svc.OneTimeCosts,
			MonthlyCosts:     svc.MonthlyCosts,
			TransactionCosts: svc.RequestCosts * record.RequestCount,
		}); err != nil {
			s.logger.Error("failed to export CSV", zap.Error(err))

			return nil, status.Error(codes.Internal, "write error")
		}
	}

	if err := exporter.Close(); err != nil {
		s.logger.Error("failed to export CSV", zap.Error(err))

		return nil, status.Error(codes.Internal, "write error")
	}

	return &api.DownloadFinanceExportResponse{
		Data: buff.Bytes(),
	}, nil
}
