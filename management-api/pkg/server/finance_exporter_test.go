package server

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinanceCSVExporter(t *testing.T) {
	tests := map[string]struct {
		rows     []FinanceRow
		csvLines []string
	}{
		"exporting_an_empty_csv_still_writes_headers": {
			csvLines: []string{
				"Organization,Service name,Month,Price per request,Number of requests,Setup costs,Monthly costs,Transaction costs",
			},
		},

		"exporting_a_single_row_works": {
			rows: []FinanceRow{
				{
					Organization:     "Test",
					ServiceName:      "service-test",
					Month:            "januari 2021",
					PricePerRequest:  125,
					NumberOfRequests: 194,
					SetupCosts:       1500,
					MonthlyCosts:     1000,
					TransactionCosts: 125 * 194,
				},
			},
			csvLines: []string{
				"Organization,Service name,Month,Price per request,Number of requests,Setup costs,Monthly costs,Transaction costs",
				`Test,service-test,januari 2021,"€1,25",194,"€15,00","€10,00","€242,50"`,
			},
		},

		"exporting_multiple_rows_works": {
			rows: []FinanceRow{
				{
					Organization:     "Test",
					ServiceName:      "service-test 1",
					Month:            "januari 2021",
					PricePerRequest:  125,
					NumberOfRequests: 194,
					SetupCosts:       1500,
					MonthlyCosts:     1000,
					TransactionCosts: 125 * 194,
				},
				{
					Organization:     "Test",
					ServiceName:      "service-test 2",
					Month:            "januari 2021",
					PricePerRequest:  25,
					NumberOfRequests: 75,
					SetupCosts:       1400,
					MonthlyCosts:     900,
					TransactionCosts: 125 * 75,
				},
			},
			csvLines: []string{
				"Organization,Service name,Month,Price per request,Number of requests,Setup costs,Monthly costs,Transaction costs",
				`Test,service-test 1,januari 2021,"€1,25",194,"€15,00","€10,00","€242,50"`,
				`Test,service-test 2,januari 2021,"€0,25",75,"€14,00","€9,00","€93,75"`,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			buff := &bytes.Buffer{}
			exporter := NewFinanceCSVExporter(buff)

			for i := range tt.rows {
				err := exporter.Export(&tt.rows[i])
				assert.NoError(t, err)
			}

			err := exporter.Close()
			assert.NoError(t, err)

			assert.Equal(t, tt.csvLines, strings.Split(strings.TrimSpace(buff.String()), "\n"))
		})
	}
}
