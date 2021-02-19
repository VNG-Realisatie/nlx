// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Row map[string]interface{}

type Exporter interface {
	Add(row Row) error
	Close() error
}

type CSVExporter struct {
	csv     *csv.Writer
	headers []string
}

func NewCSVExporter(writer io.Writer) *CSVExporter {
	return &CSVExporter{
		csv: csv.NewWriter(writer),
	}
}

func (exporter *CSVExporter) writeHeaders(row Row) error {
	exporter.headers = []string{}

	for name := range row {
		exporter.headers = append(exporter.headers, name)
	}

	if err := exporter.csv.Write(exporter.headers); err != nil {
		return err
	}

	return nil
}

func (exporter *CSVExporter) Export(row Row) error {
	if exporter.headers == nil {
		if err := exporter.writeHeaders(row); err != nil {
			return err
		}
	}

	csvRow := make([]string, len(exporter.headers))

	for i, name := range exporter.headers {
		value, ok := row[name]

		if !ok || value == nil {
			continue
		}

		switch val := value.(type) {
		case string:
			csvRow[i] = val
		default:
			csvRow[i] = fmt.Sprintf("%v", value)
		}
	}

	return exporter.csv.Write(csvRow)
}

func (exporter *CSVExporter) Close() error {
	exporter.csv.Flush()

	if err := exporter.csv.Error(); err != nil {
		return err
	}

	return nil
}
