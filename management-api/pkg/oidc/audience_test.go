// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package oidc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAudienceUnmarshal(t *testing.T) {
	tests := map[string]struct {
		input string
		want  *audience
	}{
		"audience_as_string": {
			input: "\"audience\"",
			want:  &audience{"audience"},
		},
		"audience_as_array": {
			input: "[\"audience\"]",
			want:  &audience{"audience"},
		},
		"multiple_audiences": {
			input: "[\"audience-1\", \"audience-2\"]",
			want:  &audience{"audience-1", "audience-2"},
		},
	}

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			a := &audience{}

			err := json.Unmarshal([]byte(test.input), a)
			assert.NoError(t, err)

			assert.Equal(t, tc.want, a)
		})
	}
}
