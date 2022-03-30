// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-api/pkg/directory"
)

const (
	openAPI2 = "OpenAPI2"
	openAPI3 = "OpenAPI3"
)

func TestParseAPISpecificationType(t *testing.T) {
	errUnknownVersion := errors.New("documentation format is neither openAPI2 or openAPI3")

	tests := map[string]struct {
		input   string
		want    string
		wantErr error
	}{
		"json_2.0": {

			input:   `{"swagger":"2.0"}`,
			want:    openAPI2,
			wantErr: nil,
		},
		"json_3.0.0": {

			input:   `{"openapi":"3.0.0"}`,
			want:    openAPI3,
			wantErr: nil,
		},
		"json_3.0.1": {

			input:   `{"openapi":"3.0.1"}`,
			want:    openAPI3,
			wantErr: nil,
		},
		"json_3.0.2": {

			input:   `{"openapi":"3.0.2"}`,
			want:    openAPI3,
			wantErr: nil,
		},
		"json_1.0": {
			input:   `{"swagger":"1.0"}`,
			want:    "",
			wantErr: errUnknownVersion,
		},
		"yaml_2.0": {

			input:   `swagger: "2.0"`,
			want:    openAPI2,
			wantErr: nil,
		},
		"yaml_2.0_float": {
			input:   `swagger: 2.0`,
			want:    openAPI2,
			wantErr: nil,
		},
		"empty_input": {
			input:   ``,
			want:    "",
			wantErr: errors.New("unable to parse openapi specification version: empty input"),
		},
		"invalid_json": {
			input:   `{`,
			want:    "",
			wantErr: errors.New("unable to parse openapi specification version: unexpected end of JSON input"),
		},
		"invalid_input": {
			input:   `this is invalid`,
			want:    "",
			wantErr: errors.New("unable to parse openapi specification version: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `this is...` into directory.openAPIVersion"),
		},
		"json_incomplete": {
			input:   `{}`,
			want:    "",
			wantErr: errUnknownVersion,
		},
		"yaml_incomplete": {
			input:   `---`,
			want:    "",
			wantErr: errUnknownVersion,
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			actual, err := directory.ParseAPISpecificationType([]byte(test.input))

			assert.Equal(t, test.want, actual)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
