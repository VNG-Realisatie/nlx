// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/txlog-api/domain"
)

//nolint:funlen // this is a test
func Test_NewOrder(t *testing.T) {
	tests := map[string]struct {
		args        *domain.NewOrderArgs
		expectedErr string
	}{
		"when_no_delegator": {
			args: &domain.NewOrderArgs{
				Delegator: "",
				Reference: "test-reference",
			},
			expectedErr: "Delegator: cannot be blank.",
		},
		"when_no_reference": {
			args: &domain.NewOrderArgs{
				Delegator: "test-delegator",
				Reference: "",
			},
			expectedErr: "Reference: cannot be blank.",
		},
		"happy_flow": {
			args: &domain.NewOrderArgs{
				Delegator: "test-delegator",
				Reference: "test-reference",
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewOrder(
				tt.args,
			)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.args.Delegator, result.Delegator())
				assert.Equal(t, tt.args.Reference, result.Reference())
			}
		})
	}
}
