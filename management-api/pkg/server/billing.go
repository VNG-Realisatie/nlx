// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"bytes"
	"strconv"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
)

const (
	Cents = 100
)

func (service *ManagementService) IsBillingEnabled(ctx context.Context, request *types.Empty) (*api.IsBillingEnabledResponse, error) {
	return &api.IsBillingEnabledResponse{
		Enabled: service.txlogDatabase != nil,
	}, nil
}

func (service *ManagementService) DownloadBillingExport(ctx context.Context, request *types.Empty) (*api.DownloadBillingExportResponse, error) {
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

	printer := message.NewPrinter(language.Dutch)
	scale, _ := currency.Cash.Rounding(currency.EUR)

	formatCurrency := func(amountInCents int) string {
		amount := number.Decimal(float32(amountInCents)/Cents, number.Scale(scale))

		return printer.Sprintf("%v%v", currency.Symbol(currency.EUR), amount)
	}

	buff := &bytes.Buffer{}
	exporter := NewCSVExporter(buff)

	for _, record := range records {
		svc := &database.Service{}

		if match, ok := servicesMap[record.ServiceName]; ok {
			svc = match
		}

		row := map[string]interface{}{
			"Organization":       record.Source,
			"svc":                record.ServiceName,
			"Month":              record.CreatedAt.Format("Januari 2006"),
			"Price per request":  formatCurrency(svc.RequestCosts),
			"Number of requests": strconv.Itoa(record.RequestCount),
			"Setup costs":        formatCurrency(svc.OneTimeCosts),
			"Monthly costs":      formatCurrency(svc.MonthlyCosts),
			"Transaction costs":  formatCurrency(svc.RequestCosts * record.RequestCount),
		}

		if err := exporter.Export(row); err != nil {
			service.logger.Error("failed to export CSV", zap.Error(err))

			return nil, status.Error(codes.Internal, "write error")
		}
	}

	if err := exporter.Close(); err != nil {
		service.logger.Error("failed to export CSV", zap.Error(err))

		return nil, status.Error(codes.Internal, "write error")
	}

	return &api.DownloadBillingExportResponse{
		Data: buff.Bytes(),
	}, nil
}
