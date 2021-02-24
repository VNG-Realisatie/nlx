// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"encoding/csv"
	"io"
	"strconv"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const Cents = 100

var financeExportHeaders = []string{
	"Organization",
	"Service name",
	"Month",
	"Price per request",
	"Number of requests",
	"Setup costs",
	"Monthly costs",
	"Transaction costs",
}

type FinanceRow struct {
	Organization     string
	ServiceName      string
	Month            string
	PricePerRequest  int
	NumberOfRequests int
	SetupCosts       int
	MonthlyCosts     int
	TransactionCosts int
}

type FinanceExporter interface {
	Export(row *FinanceRow) error
	Close() error
}

type FinanceCSVExporter struct {
	scale   int
	printer *message.Printer
	headers bool
	csv     *csv.Writer
}

func NewFinanceCSVExporter(writer io.Writer) *FinanceCSVExporter {
	scale, _ := currency.Cash.Rounding(currency.EUR)

	return &FinanceCSVExporter{
		csv:     csv.NewWriter(writer),
		printer: message.NewPrinter(language.Dutch),
		scale:   scale,
	}
}

func (exporter *FinanceCSVExporter) writeHeaders() error {
	exporter.headers = true

	if err := exporter.csv.Write(financeExportHeaders); err != nil {
		return err
	}

	return nil
}

func (exporter *FinanceCSVExporter) formatCurrency(amountInCents int) string {
	amount := number.Decimal(float32(amountInCents)/Cents, number.Scale(exporter.scale))

	return exporter.printer.Sprintf("%v%v", currency.Symbol(currency.EUR), amount)
}

func (exporter *FinanceCSVExporter) Export(row *FinanceRow) error {
	if !exporter.headers {
		if err := exporter.writeHeaders(); err != nil {
			return err
		}
	}

	return exporter.csv.Write([]string{
		row.Organization,
		row.ServiceName,
		row.Month,
		exporter.formatCurrency(row.PricePerRequest),
		strconv.Itoa(row.NumberOfRequests),
		exporter.formatCurrency(row.SetupCosts),
		exporter.formatCurrency(row.MonthlyCosts),
		exporter.formatCurrency(row.TransactionCosts),
	})
}

func (exporter *FinanceCSVExporter) Close() error {
	if !exporter.headers {
		if err := exporter.writeHeaders(); err != nil {
			return err
		}
	}

	exporter.csv.Flush()

	if err := exporter.csv.Error(); err != nil {
		return err
	}

	return nil
}
