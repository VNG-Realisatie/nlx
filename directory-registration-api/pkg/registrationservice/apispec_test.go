// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package registrationservice_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/pkg/registrationservice"
)

const (
	openAPI2 = "OpenAPI2"
	openAPI3 = "OpenAPI3"
)

func TestParseAPISpectType(t *testing.T) {
	errUnknownVersion := errors.New("documentation format is neither openAPI2 or openAPI3")

	tests := []struct {
		name     string
		data     string
		expected string
		err      error
	}{
		{
			"json_2.0",
			`{"swagger":"2.0"}`,
			openAPI2,
			nil,
		},
		{
			"json_3.0.0",
			`{"openapi":"3.0.0"}`,
			openAPI3,
			nil,
		},
		{
			"json_3.0.1",
			`{"openapi":"3.0.1"}`,
			openAPI3,
			nil,
		},
		{
			"json_3.0.2",
			`{"openapi":"3.0.2"}`,
			openAPI3,
			nil,
		},
		{
			"json_1.0",
			`{"swagger":"1.0"}`,
			"",
			errUnknownVersion,
		},
		{
			"yaml_2.0",
			`swagger: "2.0"`,
			openAPI2,
			nil,
		},
		{
			"yaml_2.0_float",
			`swagger: 2.0`,
			openAPI2,
			nil,
		},
		{
			"yaml_2.0_float",
			`swagger: 2.0`,
			openAPI2,
			nil,
		},
		{
			"empty_input",
			``,
			"",
			errors.New("empty input"),
		},
		{
			"invalid_input",
			`this is invalid`,
			"",
			errors.New("unable to parse version"),
		},
		{
			"json_incomplete",
			`{}`,
			"",
			errUnknownVersion,
		},
		{
			"yaml_incomplete",
			`---`,
			"",
			errUnknownVersion,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			actual, err := registrationservice.ParseAPISpectType([]byte(test.data))

			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.err, err)
		})
	}
}
