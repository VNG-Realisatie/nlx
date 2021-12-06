package directoryapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func TestSetOrganizationContactDetailsValidate(t *testing.T) {
	tests := map[string]struct {
		request *directoryapi.SetOrganizationContactDetailsRequest
		err     string
	}{
		"invalid_email_address": {
			request: &directoryapi.SetOrganizationContactDetailsRequest{
				EmailAddress: "invalid-email",
			},
			err: "email_address: must be a valid email address.",
		},
		"happy_flow": {
			request: &directoryapi.SetOrganizationContactDetailsRequest{
				EmailAddress: "mock@email.com",
			},
			err: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			err := tt.request.Validate()
			if err != nil {
				assert.EqualError(t, err, tt.err)
				return
			}

			assert.Equal(t, nil, err)
		})
	}
}
