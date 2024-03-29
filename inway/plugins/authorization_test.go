// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/inway/plugins"
)

//nolint:funlen // this is a test
func TestAuthorizationPlugin(t *testing.T) {
	tests := map[string]struct {
		args                         *plugins.AuthRequest
		authServerEnabled            bool
		authServerResponse           interface{}
		authServerResponseStatusCode int
		wantErr                      *httperrors.NLXNetworkError
		wantHTTPStatusCode           int
	}{
		"when_auth_server_returns_non_OK_status": {
			args: &plugins.AuthRequest{
				Input: &plugins.AuthRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
					Service: &plugins.Service{
						Grants: []*plugins.Grant{
							{
								OrganizationSerialNumber: "00000000000000000001",
								PublicKeyFingerprint:     "mock-public-key-fingerprint",
							},
						},
					},
				},
			},
			authServerResponse: &plugins.AuthResponse{
				Result: true,
			},
			authServerEnabled:            true,
			authServerResponseStatusCode: http.StatusUnauthorized,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.IAS1,
				Code:     httperrors.ErrorWhileAuthorizingRequestErr,
				Message:  "error authorizing request",
			},
		},
		"when_auth_server_returns_invalid_response": {
			args: &plugins.AuthRequest{
				Input: &plugins.AuthRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
					Service: &plugins.Service{
						Grants: []*plugins.Grant{
							{
								OrganizationSerialNumber: "00000000000000000001",
								PublicKeyFingerprint:     "mock-public-key-fingerprint",
							},
						},
					},
				},
			},
			authServerResponse: struct {
				Invalid string `json:"invalid"`
			}{
				Invalid: "this is an invalid response",
			},
			authServerEnabled:            true,
			authServerResponseStatusCode: http.StatusOK,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.IAS1,
				Code:     httperrors.UnauthorizedErr,
				Message:  "authorization server denied request",
			},
		},
		"when_auth_server_fails": {
			args: &plugins.AuthRequest{
				Input: &plugins.AuthRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
					Service: &plugins.Service{
						Grants: []*plugins.Grant{
							{
								OrganizationSerialNumber: "00000000000000000001",
								PublicKeyFingerprint:     "mock-public-key-fingerprint",
							},
						},
					},
				},
			},
			authServerEnabled:            true,
			authServerResponseStatusCode: http.StatusInternalServerError,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.IAS1,
				Code:     httperrors.ErrorWhileAuthorizingRequestErr,
				Message:  "error authorizing request",
			},
		},
		"when_auth_server_returns_no_access": {
			args: &plugins.AuthRequest{
				Input: &plugins.AuthRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
					Service: &plugins.Service{
						Grants: []*plugins.Grant{
							{
								OrganizationSerialNumber: "00000000000000000001",
								PublicKeyFingerprint:     "mock-public-key-fingerprint",
							},
						},
					},
				},
			},
			authServerEnabled: true,
			authServerResponse: &plugins.AuthResponse{
				Result: false,
			},
			authServerResponseStatusCode: http.StatusOK,
			wantHTTPStatusCode:           httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.IAS1,
				Code:     httperrors.UnauthorizedErr,
				Message:  "authorization server denied request",
			},
		},
		"when_access_grant_not_found": {
			args: &plugins.AuthRequest{
				Input: &plugins.AuthRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
					Service: &plugins.Service{
						Name:   "mock-service",
						Grants: []*plugins.Grant{},
					},
				},
			},
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.AccessDeniedErr,
				Message:  "permission denied, organization \"00000000000000000001\" or public key fingerprint \"mock-public-key-fingerprint\" is not allowed access.",
			},
		},
		"happy_flow_with_auth_server": {
			args: &plugins.AuthRequest{
				Input: &plugins.AuthRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
					Service: &plugins.Service{
						Grants: []*plugins.Grant{
							{
								OrganizationSerialNumber: "00000000000000000001",
								PublicKeyFingerprint:     "mock-public-key-fingerprint",
							},
						},
					},
				},
			},
			authServerEnabled: true,
			authServerResponse: &plugins.AuthResponse{
				Result: true,
			},
			authServerResponseStatusCode: http.StatusOK,
			wantHTTPStatusCode:           http.StatusOK,
		},
		"happy_flow": {
			args: &plugins.AuthRequest{
				Input: &plugins.AuthRequestInput{
					Headers: http.Header{
						"Proxy-Authorization": []string{"Bearer abc"},
					},
					Service: &plugins.Service{
						Grants: []*plugins.Grant{
							{
								OrganizationSerialNumber: "00000000000000000001",
								PublicKeyFingerprint:     "mock-public-key-fingerprint",
							},
						},
					},
				},
			},
			authServerEnabled:  false,
			wantHTTPStatusCode: http.StatusOK,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&plugins.Destination{
				Organization: tt.args.Input.OrganizationSerialNumber,
				Service:      tt.args.Input.Service,
				Path:         tt.args.Input.Path,
			}, nil, &plugins.AuthInfo{
				OrganizationSerialNumber: "00000000000000000001",
				PublicKeyFingerprint:     "mock-public-key-fingerprint",
			})

			for k, values := range tt.args.Input.Headers {
				for _, v := range values {
					context.Request.Header.Add(k, v)
				}
			}

			var gotAuthorizationServerRequest []byte

			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					body := r.Body
					defer r.Body.Close()

					var err error
					gotAuthorizationServerRequest, err = io.ReadAll(body)
					assert.NoError(t, err)

					w.WriteHeader(tt.authServerResponseStatusCode)

					b, err := json.Marshal(tt.authServerResponse)
					assert.NoError(t, err)

					_, err = w.Write(b)
					assert.NoError(t, err)
				}),
			)

			plugin := plugins.NewAuthorizationPlugin(&plugins.NewAuthorizationPluginArgs{
				CA:                  nil,
				ServiceURL:          server.URL,
				AuthorizationClient: http.DefaultClient,
				AuthServerEnabled:   tt.authServerEnabled,
			})

			err := plugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()

			defer response.Body.Close()

			contents, err := io.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantHTTPStatusCode, response.StatusCode)

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			} else if tt.authServerEnabled {
				wantAuthorizationServiceRequest, err := json.Marshal(tt.args)
				assert.NoError(t, err)

				assert.Equal(t, wantAuthorizationServiceRequest, gotAuthorizationServerRequest)
			}
		})
	}
}
