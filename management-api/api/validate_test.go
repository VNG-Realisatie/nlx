package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/api"
)

// nolint:dupl // this is a test function
func TestCreateServiceValidate(t *testing.T) {
	tests := map[string]struct {
		service *api.CreateServiceRequest
		err     string
	}{
		"without_name_and_endpoint": {
			service: &api.CreateServiceRequest{
				Name:        "",
				EndpointURL: "",
			},
			err: "endpointURL: cannot be blank; name: cannot be blank.",
		},
		"using_invalid_endpoint": {
			service: &api.CreateServiceRequest{
				Name:        "my-service",
				EndpointURL: "invalid-endpoint",
			},
			err: "endpointURL: must be a valid URL.",
		},
		"happy_flow": {
			service: &api.CreateServiceRequest{
				Name:        "my-service",
				EndpointURL: "https://my-service.test",
			},
			err: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			err := tt.service.Validate()
			if err != nil {
				assert.EqualError(t, err, tt.err)
				return
			}

			assert.Equal(t, nil, err)
		})
	}
}

// nolint:dupl // this is a test function
func TestUpdateServiceValidate(t *testing.T) {
	tests := map[string]struct {
		service *api.UpdateServiceRequest
		err     string
	}{
		"without_name_and_endpoint": {
			service: &api.UpdateServiceRequest{
				Name:        "",
				EndpointURL: "",
			},
			err: "endpointURL: cannot be blank; name: cannot be blank.",
		},
		"using_invalid_endpoint": {
			service: &api.UpdateServiceRequest{
				Name:        "my-service",
				EndpointURL: "invalid-endpoint",
			},
			err: "endpointURL: must be a valid URL.",
		},
		"happy_flow": {
			service: &api.UpdateServiceRequest{
				Name:        "my-service",
				EndpointURL: "https://my-service.test",
			},
			err: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			err := tt.service.Validate()
			if err != nil {
				assert.EqualError(t, err, tt.err)
				return
			}

			assert.Equal(t, nil, err)
		})
	}
}

func TestInwayValidate(t *testing.T) {
	tests := map[string]struct {
		inway *api.Inway
		err   string
	}{
		"without_name": {
			inway: &api.Inway{
				Name: "",
			},
			err: "name: cannot be blank.",
		},
		"happy_flow": {
			inway: &api.Inway{
				Name: "inway42.test",
			},
			err: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			err := tt.inway.Validate()
			if err != nil {
				assert.EqualError(t, err, tt.err)
				return
			}

			assert.Equal(t, nil, err)
		})
	}
}
