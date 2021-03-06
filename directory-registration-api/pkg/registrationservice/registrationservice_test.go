// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/pkg/registrationservice"
)

const testOrganizationName = "Test Organization Name"
const testInvalidOrganizationName = ""

func testGetOrganizationNameFromRequest(ctx context.Context) (string, error) {
	return testOrganizationName, nil
}

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
		assert.Equal(t, registrationservice.IsValidOrganizationName(test.organisationName), test.expectedReturn)
	}
}
