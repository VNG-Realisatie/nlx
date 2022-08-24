// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/txlog-api/domain"
)

//nolint:funlen // this is a test
func Test_NewRecord(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		args        *domain.NewRecordArgs
		expectedErr string
	}{
		"empty_source_organization": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "SourceOrganization: cannot be blank.",
		},
		"invalid_source_organization": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "000000000000000000001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "SourceOrganization: too long, max 20 bytes.",
		},
		"empty_destination_organization": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "DestinationOrganization: cannot be blank.",
		},
		"invalid_destination": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "000000000000000000002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "DestinationOrganization: too long, max 20 bytes.",
		},
		"empty_direction": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               "",
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "Direction: cannot be blank.",
		},
		"empty_service_name": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "ServiceName: cannot be blank.",
		},
		"empty_transaction_id": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "",
				CreatedAt:               now,
			},
			expectedErr: "TransactionID: cannot be blank.",
		},
		"empty_created_at": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
			},
			expectedErr: "CreatedAt: cannot be blank.",
		},
		"order_reference_without_delegator": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "empty delegator, both the delegator and order reference should be provided",
		},
		"delegator_without_order_reference": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "",
				Delegator:               "0003",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "empty order reference, both the delegator and order reference should be provided",
		},
		"invalid_delegator": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "000000000000000000001",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "Delegator: too long, max 20 bytes.",
		},
		"happy_flow": {
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service-2",
				OrderReference:          "test-reference",
				Delegator:               "0004",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               now,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewRecord(tt.args)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.args.SourceOrganization, result.SourceOrganization())
				assert.Equal(t, tt.args.DestinationOrganization, result.DestinationOrganization())
				assert.Equal(t, tt.args.Direction, result.Direction())
				assert.Equal(t, tt.args.ServiceName, result.ServiceName())
				assert.Equal(t, tt.args.Delegator, result.Delegator())
				assert.Equal(t, tt.args.Data, result.Data())
				assert.Equal(t, tt.args.TransactionID, result.TransactionID())
				assert.Equal(t, tt.args.CreatedAt, result.CreatedAt())
			}
		})
	}
}
