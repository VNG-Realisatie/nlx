// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeInwayProxyAddress(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    string
		wantErr string
	}{
		"empty_address": {
			input:   "",
			want:    "",
			wantErr: "empty inway address provided",
		},
		"without_port": {
			input:   "localhost",
			want:    "",
			wantErr: "invalid format for inway address: address localhost: missing port in address",
		},
		"happy_flow": {
			input:   "localhost:8000",
			want:    "localhost:8001",
			wantErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			got, err := computeInwayProxyAddress(tt.input)

			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
