package domain_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/domain"
)

//nolint funlen: this is a test function
func Test_NewIncomingOrder(t *testing.T) {
	t.Parallel()

	type orderParams struct {
		reference   string
		description string
		delegator   string
		revokedAt   *time.Time
		validFrom   time.Time
		validUntil  time.Time
		services    []domain.IncomingOrderService
	}

	validFrom := time.Now().Add(-24 * time.Hour)
	validUntil := time.Now().Add(24 * time.Hour)

	tests := map[string]struct {
		order       orderParams
		expectedErr error
	}{
		"empty_reference": {
			order: orderParams{
				reference:   "",
				description: "my-description",
				delegator:   "my-delegator",
				revokedAt:   nil,
				validFrom:   validFrom,
				validUntil:  validUntil,
				services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "my-organization"),
				},
			},
			expectedErr: errors.New("reference: cannot be blank"),
		},
		"empty_description": {
			order: orderParams{
				reference:   "my-reference",
				description: "",
				delegator:   "my-delegator",
				revokedAt:   nil,
				validFrom:   validFrom,
				validUntil:  validUntil,
				services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "my-organization"),
				},
			},
			expectedErr: errors.New("description: cannot be blank"),
		},
		"empty_delegator": {
			order: orderParams{
				reference:   "my-reference",
				description: "my-description",
				delegator:   "",
				revokedAt:   nil,
				validFrom:   validFrom,
				validUntil:  validUntil,
				services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "my-organization"),
				},
			},
			expectedErr: errors.New("delegator: cannot be blank"),
		},
		"valid_from_is_after_valid_until": {
			order: orderParams{
				reference:   "my-reference",
				description: "my-description",
				delegator:   "my-delegator",
				revokedAt:   nil,
				validFrom:   time.Now(),
				validUntil:  time.Now().Add(-1 * time.Hour),
				services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "my-organization"),
				},
			},
			expectedErr: errors.New("valid from: order can not expire before the start date"),
		},
		"happy_flow": {
			order: orderParams{
				reference:   "my-reference",
				description: "my-description",
				delegator:   "my-delegator",
				revokedAt:   nil,
				validFrom:   validFrom,
				validUntil:  validUntil,
				services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "my-organization"),
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewIncomingOrder(
				tt.order.reference,
				tt.order.description,
				tt.order.delegator,
				tt.order.revokedAt,
				tt.order.validFrom,
				tt.order.validUntil,
				tt.order.services,
			)

			require.Equal(t, tt.expectedErr, err)
			if tt.expectedErr == nil {
				require.Equal(t, tt.order.reference, result.Reference())
				require.Equal(t, tt.order.description, result.Description())
				require.Equal(t, tt.order.delegator, result.Delegator())
				require.Equal(t, tt.order.revokedAt, result.RevokedAt())
				require.Equal(t, tt.order.validFrom, result.ValidFrom())
				require.Equal(t, tt.order.validUntil, result.ValidUntil())
				require.Equal(t, tt.order.services, result.Services())
			}
		})
	}
}
