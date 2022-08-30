// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/strings"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "FirstName",
			want:  "first_name",
		},
		{
			input: "LAST_NAME",
			want:  "last_name",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got := strings.ToSnakeCase(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
