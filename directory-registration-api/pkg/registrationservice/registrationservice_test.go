// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValidOrganizationName(t *testing.T) {
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
		assert.Equal(t, isValidOrganizationName(test.organisationName), test.expectedReturn)
	}
}

func Test_IsValidServiceName(t *testing.T) {
	tests := []struct {
		serviceName    string
		expectedReturn bool
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
		assert.Equal(t, isValidServiceName(test.serviceName), test.expectedReturn)
	}
}
