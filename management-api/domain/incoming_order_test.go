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

	validFrom := time.Now().Add(-24 * time.Hour)
	validUntil := time.Now().Add(24 * time.Hour)

	tests := map[string]struct {
		order       *domain.NewIncomingOrderArgs
		expectedErr error
	}{
		"empty_reference": {
			order: &domain.NewIncomingOrderArgs{
				Reference:   "",
				Description: "my-description",
				Delegator:   "my-delegator",
				RevokedAt:   nil,
				ValidFrom:   validFrom,
				ValidUntil:  validUntil,
				Services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "00000000000000000001", "my-organization"),
				},
			},
			expectedErr: errors.New("reference: cannot be blank"),
		},
		"empty_description": {
			order: &domain.NewIncomingOrderArgs{
				Reference:   "my-reference",
				Description: "",
				Delegator:   "my-delegator",
				RevokedAt:   nil,
				ValidFrom:   validFrom,
				ValidUntil:  validUntil,
				Services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "00000000000000000001", "my-organization"),
				},
			},
			expectedErr: errors.New("description: cannot be blank"),
		},
		"empty_delegator": {
			order: &domain.NewIncomingOrderArgs{
				Reference:   "my-reference",
				Description: "my-description",
				Delegator:   "",
				RevokedAt:   nil,
				ValidFrom:   validFrom,
				ValidUntil:  validUntil,
				Services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "00000000000000000001", "my-organization"),
				},
			},
			expectedErr: errors.New("delegator: cannot be blank"),
		},
		"empty_services": {
			order: &domain.NewIncomingOrderArgs{
				Reference:   "my-reference",
				Description: "my-description",
				Delegator:   "my-delegator",
				RevokedAt:   nil,
				ValidFrom:   validFrom,
				ValidUntil:  validUntil,
				Services:    []domain.IncomingOrderService{},
			},
			expectedErr: nil,
		},
		"valid_from_is_after_valid_until": {
			order: &domain.NewIncomingOrderArgs{
				Reference:   "my-reference",
				Description: "my-description",
				Delegator:   "my-delegator",
				RevokedAt:   nil,
				ValidFrom:   time.Now(),
				ValidUntil:  time.Now().Add(-1 * time.Hour),
				Services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "00000000000000000001", "my-organization"),
				},
			},
			expectedErr: errors.New("valid from: order can not expire before the start date"),
		},
		"happy_flow": {
			order: &domain.NewIncomingOrderArgs{
				Reference:   "my-reference",
				Description: "my-description",
				Delegator:   "my-delegator",
				RevokedAt:   nil,
				ValidFrom:   validFrom,
				ValidUntil:  validUntil,
				Services: []domain.IncomingOrderService{
					domain.NewIncomingOrderService("my-service", "00000000000000000001", "my-organization"),
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewIncomingOrder(tt.order)

			require.Equal(t, tt.expectedErr, err)
			if tt.expectedErr == nil {
				require.Equal(t, tt.order.Reference, result.Reference())
				require.Equal(t, tt.order.Description, result.Description())
				require.Equal(t, tt.order.Delegator, result.Delegator())
				require.Equal(t, tt.order.RevokedAt, result.RevokedAt())
				require.Equal(t, tt.order.ValidFrom, result.ValidFrom())
				require.Equal(t, tt.order.ValidUntil, result.ValidUntil())
				require.Equal(t, tt.order.Services, result.Services())
			}
		})
	}
}
