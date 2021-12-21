// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/txlog-api/domain"
)

//nolint:funlen // this is a test
func Test_NewRecord(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		args        *domain.NewRecordArgs
		expectedErr string
	}{
		"when_no_source": {
			args: &domain.NewRecordArgs{
				Source:        nil,
				Destination:   createNewOrganization(t, "0002"),
				Direction:     domain.IN,
				Service:       createNewService(t, "test-service-2"),
				Order:         createNewOrder(t, "0003", "test-reference"),
				Data:          []byte(`{"test": "value"}`),
				TransactionID: "abcde",
				CreatedAt:     now,
			},
			expectedErr: "Source: cannot be blank.",
		},
		"when_no_dest": {
			args: &domain.NewRecordArgs{
				Source:        createNewOrganization(t, "0001"),
				Destination:   nil,
				Direction:     domain.IN,
				Service:       createNewService(t, "test-service"),
				Order:         createNewOrder(t, "0003", "test-reference"),
				Data:          []byte(`{"test": "value"}`),
				TransactionID: "abcde",
				CreatedAt:     now,
			},
			expectedErr: "Destination: cannot be blank.",
		},
		"when_no_direction": {
			args: &domain.NewRecordArgs{
				Source:        createNewOrganization(t, "0001"),
				Destination:   createNewOrganization(t, "0002"),
				Direction:     "",
				Service:       createNewService(t, "test-service"),
				Order:         createNewOrder(t, "0003", "test-reference"),
				Data:          []byte(`{"test": "value"}`),
				TransactionID: "abcde",
				CreatedAt:     now,
			},
			expectedErr: "Direction: cannot be blank.",
		},
		"when_no_service": {
			args: &domain.NewRecordArgs{
				Source:        createNewOrganization(t, "0001"),
				Destination:   createNewOrganization(t, "0002"),
				Direction:     domain.IN,
				Service:       nil,
				Order:         createNewOrder(t, "0003", "test-reference"),
				Data:          []byte(`{"test": "value"}`),
				TransactionID: "abcde",
				CreatedAt:     now,
			},
			expectedErr: "Service: cannot be blank.",
		},
		"when_no_transaction_id": {
			args: &domain.NewRecordArgs{
				Source:        createNewOrganization(t, "0001"),
				Destination:   createNewOrganization(t, "0002"),
				Direction:     domain.IN,
				Service:       createNewService(t, "test-service"),
				Order:         createNewOrder(t, "0003", "test-reference"),
				Data:          []byte(`{"test": "value"}`),
				TransactionID: "",
				CreatedAt:     now,
			},
			expectedErr: "TransactionID: cannot be blank.",
		},
		"when_no_created_at": {
			args: &domain.NewRecordArgs{
				Source:        createNewOrganization(t, "0001"),
				Destination:   createNewOrganization(t, "0002"),
				Direction:     domain.IN,
				Service:       createNewService(t, "test-service"),
				Order:         createNewOrder(t, "0003", "test-reference"),
				Data:          []byte(`{"test": "value"}`),
				TransactionID: "abcde",
				CreatedAt:     time.Time{},
			},
			expectedErr: "CreatedAt: cannot be blank.",
		},
		"happy_flow": {
			args: &domain.NewRecordArgs{
				Source:        createNewOrganization(t, "0001"),
				Destination:   createNewOrganization(t, "0002"),
				Direction:     domain.IN,
				Service:       createNewService(t, "test-service"),
				Order:         createNewOrder(t, "0004", "test-reference-2"),
				Data:          []byte(`{"test": "value"}`),
				TransactionID: "abcde",
				CreatedAt:     now,
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

				assert.Equal(t, tt.args.Source.SerialNumber(), result.Source().SerialNumber())
				assert.Equal(t, tt.args.Destination.SerialNumber(), result.Destination().SerialNumber())
				assert.Equal(t, tt.args.Direction, result.Direction())
				assert.Equal(t, tt.args.Service.Name(), result.Service().Name())
				assert.Equal(t, tt.args.Order.Delegator(), result.Order().Delegator())
				assert.Equal(t, tt.args.Data, result.Data())
				assert.Equal(t, tt.args.TransactionID, result.TransactionID())
				assert.Equal(t, tt.args.CreatedAt, result.CreatedAt())
			}
		})
	}
}

func createNewOrganization(t *testing.T, serialNumber string) *domain.Organization {
	m, err := domain.NewOrganization(serialNumber)
	require.NoError(t, err)

	return m
}

func createNewService(t *testing.T, name string) *domain.Service {
	m, err := domain.NewService(name)
	require.NoError(t, err)

	return m
}

func createNewOrder(t *testing.T, delegator, reference string) *domain.Order {
	m, err := domain.NewOrder(&domain.NewOrderArgs{
		Delegator: delegator,
		Reference: reference,
	})
	require.NoError(t, err)

	return m
}
