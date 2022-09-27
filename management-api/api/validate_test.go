// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package api_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/api"
	common_testing "go.nlx.io/nlx/testing/testingutils"
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
				EndpointUrl: "",
			},
			err: "endpoint_url: cannot be blank; name: cannot be blank.",
		},
		"using_invalid_endpoint": {
			service: &api.CreateServiceRequest{
				Name:        "my-service",
				EndpointUrl: "invalid-endpoint",
			},
			err: "endpoint_url: must be a valid URL.",
		},
		"happy_flow": {
			service: &api.CreateServiceRequest{
				Name:        "my-service",
				EndpointUrl: "https://my-service.test",
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
				EndpointUrl: "",
			},
			err: "endpoint_url: cannot be blank; name: cannot be blank.",
		},
		"using_invalid_endpoint": {
			service: &api.UpdateServiceRequest{
				Name:        "my-service",
				EndpointUrl: "invalid-endpoint",
			},
			err: "endpoint_url: must be a valid URL.",
		},
		"happy_flow": {
			service: &api.UpdateServiceRequest{
				Name:        "my-service",
				EndpointUrl: "https://my-service.test",
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

func TestOutwayValidate(t *testing.T) {
	pkiDir := filepath.Join("..", "..", "testing", "pki")

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.NLXTestInternal)
	require.NoError(t, err)

	testPublicKeyPEM, err := certBundle.PublicKeyPEM()
	require.NoError(t, err)

	tests := map[string]struct {
		req *api.RegisterOutwayRequest
		err string
	}{
		"without_name": {
			req: &api.RegisterOutwayRequest{
				Name:           "",
				SelfAddressApi: "mock-address",
				PublicKeyPem:   testPublicKeyPEM,
				Version:        "unknown",
			},
			err: "name: cannot be blank.",
		},
		"without_public_key_pem": {
			req: &api.RegisterOutwayRequest{
				Name:           "outway42.test",
				SelfAddressApi: "mock-address",
				Version:        "unknown",
			},
			err: "public_key_pem: cannot be blank.",
		},
		"without_self_address_api": {
			req: &api.RegisterOutwayRequest{
				Name:         "outway42.test",
				Version:      "unknown",
				PublicKeyPem: testPublicKeyPEM,
			},
			err: "self_address_api: cannot be blank.",
		},
		"without_version": {
			req: &api.RegisterOutwayRequest{
				Name:           "outway42.test",
				SelfAddressApi: "mock-address",
				PublicKeyPem:   testPublicKeyPEM,
			},
			err: "version: cannot be blank.",
		},
		"happy_flow": {
			req: &api.RegisterOutwayRequest{
				Name:           "outway42.test",
				SelfAddressApi: "mock-address",
				PublicKeyPem:   testPublicKeyPEM,
				Version:        "unknown",
			},
			err: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			err := tt.req.Validate()
			if err != nil {
				assert.EqualError(t, err, tt.err)
				return
			}

			assert.Equal(t, nil, err)
		})
	}
}
