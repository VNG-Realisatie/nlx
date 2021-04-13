// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/inway/plugins"
)

func TestAuthorizationPlugin(t *testing.T) {
	authorizationPlugin := plugins.NewAuthorizationPlugin()

	tests := map[string]struct {
		service            *plugins.Service
		expectedStatusCode int
		expectedMessage    string
	}{
		"unauthorized": {
			service: &plugins.Service{
				Name:   "mock-service",
				Grants: []*plugins.Grant{},
			},
			expectedStatusCode: http.StatusForbidden,
			expectedMessage:    "nlx-inway: permission denied, organization \"mock-org\" or public key fingerprint \"mock-public-key-fingerprint\" is not allowed access.\n",
		},
		"happy_flow": {
			service: &plugins.Service{
				Name: "mock-service",
				Grants: []*plugins.Grant{
					{
						OrganizationName:     "mock-org",
						PublicKeyFingerprint: "mock-public-key-fingerprint",
					},
				},
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&plugins.Destination{
				Service: tt.service,
			}, nil, &plugins.AuthInfo{
				OrganizationName:     "mock-org",
				PublicKeyFingerprint: "mock-public-key-fingerprint",
			})

			err := authorizationPlugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()
			defer response.Body.Close()

			contents, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedMessage, string(contents))
			assert.Equal(t, tt.expectedStatusCode, response.StatusCode)
		})
	}
}
