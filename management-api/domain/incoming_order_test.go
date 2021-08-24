package domain_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/domain"
)

func Test_NewIncomingOrder(t *testing.T) {
	type orderParams struct {
		reference string
	}

	tests := map[string]struct {
		order       orderParams
		expectedErr error
	}{
		"empty_reference": {
			order: orderParams{
				reference: "",
			},
			expectedErr: errors.New("reference: cannot be blank"),
		},
		"happy_flow": {
			order: orderParams{
				reference: "my-reference",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewIncomingOrder(
				tt.order.reference,
			)

			require.Equal(t, tt.expectedErr, err)
			if tt.expectedErr == nil {
				assert.Equal(t, tt.order.reference, result.Reference())
			}
		})
	}
}
