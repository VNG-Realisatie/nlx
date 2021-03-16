// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeInwayProxyAddress(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
		err   string
	}{
		"empty_address": {
			"",
			"",
			"invalid format for inway address: missing port in address",
		},
		"without_port": {
			"localhost",
			"",
			"invalid format for inway address: address localhost: missing port in address",
		},
		"happy_flow": {
			"localhost:8000",
			"localhost:8001",
			"invalid format for inway address: missing port in address",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			actual, err := computeInwayProxyAddress(tt.input)
			if err != nil {
				assert.EqualError(t, err, tt.err)
			}

			assert.Equal(t, tt.want, actual)
		})
	}
}
