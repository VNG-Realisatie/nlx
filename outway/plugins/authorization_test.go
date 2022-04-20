// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/httperrors"
)

//nolint:funlen // this is a test
func TestAuthorizationPlugin(t *testing.T) {
	tests := map[string]struct {
		args                         *authRequest
		authServerResponse           interface{}
		authServerResponseStatusCode int
		wantError                    string
		wantHTTPStatusCode           int
	}{
		"when_auth_server_returns_non_OK_status": {
			args: &authRequest{
				Input: &authRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
				},
			},
			authServerResponse: &authResponse{
				Result: true,
			},
			authServerResponseStatusCode: http.StatusUnauthorized,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantError:                    "nlx-outway: error authorizing request\n",
		},
		"when_auth_server_returns_invalid_response": {
			args: &authRequest{
				Input: &authRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
				},
			},
			authServerResponse: struct {
				Invalid string `json:"invalid"`
			}{
				Invalid: "this is an invalid response",
			},
			authServerResponseStatusCode: http.StatusOK,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantError:                    "nlx-outway: authorization server denied request\n",
		},
		"when_auth_server_fails": {
			args: &authRequest{
				Input: &authRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
				},
			},
			authServerResponseStatusCode: http.StatusInternalServerError,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantError:                    "nlx-outway: error authorizing request\n",
		},
		"when_auth_server_returns_no_access": {
			args: &authRequest{
				Input: &authRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
				},
			},
			authServerResponse: &authResponse{
				Result: false,
			},
			authServerResponseStatusCode: http.StatusOK,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantError:                    "nlx-outway: authorization server denied request\n",
		},
		"happy_flow": {
			args: &authRequest{
				Input: &authRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
				},
			},
			authServerResponse: &authResponse{
				Result: true,
			},
			authServerResponseStatusCode: http.StatusOK,
			wantHTTPStatusCode:           http.StatusOK,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&Destination{
				OrganizationSerialNumber: tt.args.Input.OrganizationSerialNumber,
				Service:                  tt.args.Input.Service,
				Path:                     tt.args.Input.Path,
			})

			for k, values := range tt.args.Input.Headers {
				for _, v := range values {
					context.Request.Header.Add(k, v)
				}
			}

			var gotAuthorizationServiceRequest []byte

			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					body := r.Body

					var err error
					gotAuthorizationServiceRequest, err = io.ReadAll(body)
					assert.NoError(t, err)

					w.WriteHeader(tt.authServerResponseStatusCode)

					b, err := json.Marshal(tt.authServerResponse)
					assert.NoError(t, err)

					_, err = w.Write(b)
					assert.NoError(t, err)
				}),
			)

			plugin := NewAuthorizationPlugin(&NewAuthorizationPluginArgs{
				CA:                  nil,
				ServiceURL:          server.URL,
				AuthorizationClient: *http.DefaultClient,
			})

			err := plugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()

			defer response.Body.Close()

			contents, err := io.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantError, string(contents))
			assert.Equal(t, tt.wantHTTPStatusCode, response.StatusCode)

			if tt.wantError == "" {
				wantAuthorizationServiceRequest, err := json.Marshal(tt.args)
				assert.NoError(t, err)

				assert.Equal(t, wantAuthorizationServiceRequest, gotAuthorizationServiceRequest)
			}
		})
	}
}
