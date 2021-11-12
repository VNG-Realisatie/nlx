// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-api/domain"
)

//nolint:funlen // this is a test
func Test_NewVersionStatistics(t *testing.T) {
	tests := map[string]struct {
		args        domain.NewVersionStatisticsArgs
		expectedErr error
	}{
		"when_no_gateway_type": {
			args: domain.NewVersionStatisticsArgs{
				GatewayType: "",
				Version:     "11.0.1",
				Amount:      10,
			},
			expectedErr: errors.New("GatewayType: cannot be blank."),
		},
		"when_no_version": {
			args: domain.NewVersionStatisticsArgs{
				GatewayType: domain.TypeInway,
				Version:     "",
				Amount:      10,
			},
			expectedErr: errors.New("Version: cannot be blank."),
		},
		"happy_flow": {
			args: domain.NewVersionStatisticsArgs{
				GatewayType: domain.TypeInway,
				Version:     "11.0.1",
				Amount:      10,
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewVersionStatistics(
				&tt.args,
			)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, tt.args.GatewayType, result.GatewayType())
			assert.Equal(t, tt.args.Version, result.Version())
			assert.Equal(t, tt.args.Amount, result.Amount())
		})
	}
}
