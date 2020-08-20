// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		organisationName string
		expectedReturn   bool
	}{
		{
			"gemeente-turfbrug",
			true,
		},
		{
			"Gemeente Turfbrug",
			true,
		}, {
			"VNG Realisatie B.V.",
			true,
		},
		{
			"VNG Réalisatie B.V.",
			false,
		},
		{
			"gemeente/turfburg",
			false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, validateName(test.organisationName), test.expectedReturn)
	}
}
