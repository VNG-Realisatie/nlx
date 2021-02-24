// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"bytes"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
)

func (service *ManagementService) IsFinanceEnabled(ctx context.Context, request *types.Empty) (*api.IsFinanceEnabledResponse, error) {
	return &api.IsFinanceEnabledResponse{
		Enabled: service.txlogDatabase != nil,
	}, nil
}

func (service *ManagementService) DownloadFinanceExport(ctx context.Context, request *types.Empty) (*api.DownloadFinanceExportResponse, error) {
	services, err := service.configDatabase.ListServices(ctx)
	if err != nil {
		service.logger.Error("failed to list services", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	servicesMap := map[string]*database.Service{}

	for _, service := range services {
		servicesMap[service.Name] = service
	}

	records := []txlogdb.Record{}

	if service.txlogDatabase != nil {
		records, err = service.txlogDatabase.FilterRecords(
			ctx,
			&txlogdb.Filters{
				Destination: service.orgCert.Certificate().Issuer.CommonName,
				Direction:   transactionlog.DirectionIn,
			},
		)
		if err != nil {
			service.logger.Error("failed to filter records", zap.Error(err))

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
			Month:            record.CreatedAt.Format("Januari 2006"),
			PricePerRequest:  svc.RequestCosts,
			NumberOfRequests: record.RequestCount,
			SetupCosts:       svc.OneTimeCosts,
			MonthlyCosts:     svc.MonthlyCosts,
			TransactionCosts: svc.RequestCosts * record.RequestCount,
		}); err != nil {
			service.logger.Error("failed to export CSV", zap.Error(err))

			return nil, status.Error(codes.Internal, "write error")
		}
	}

	if err := exporter.Close(); err != nil {
		service.logger.Error("failed to export CSV", zap.Error(err))

		return nil, status.Error(codes.Internal, "write error")
	}

	return &api.DownloadFinanceExportResponse{
		Data: buff.Bytes(),
	}, nil
}
