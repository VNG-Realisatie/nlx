// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/management-api/api"
	mock "go.nlx.io/nlx/management-api/api/mock"
)

var testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
	"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
	"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func TestDelegationPlugin(t *testing.T) {
	tests := map[string]struct {
		wantErr            bool
		wantMessage        string
		wantHTTPStatusCode int
		setup              func(*mock.MockManagementClient, *DelegationPlugin)
		setHeaders         func(*http.Request)
	}{
		"missing_order_reference_returns_an_errors": {
			wantErr:            true,
			wantHTTPStatusCode: http.StatusInternalServerError,
			wantMessage:        "failed to parse delegation metadata\n",
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "00000000000000000001")
			},
		},

		"missing_delegator_returns_an_errors": {
			wantErr:            true,
			wantHTTPStatusCode: http.StatusInternalServerError,
			wantMessage:        "failed to parse delegation metadata\n",
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Order-Reference", "test-ref-123")
			},
		},

		"error_while_retrieving_claim_returns_an_error": {
			wantErr:            true,
			wantHTTPStatusCode: http.StatusInternalServerError,
			wantMessage:        "failed to retrieve claim\n",
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "00000000000000000001")
				r.Header.Add("X-NLX-Request-Order-Reference", "test-ref-123")
			},
			setup: func(client *mock.MockManagementClient, plugin *DelegationPlugin) {
				client.EXPECT().
					RetrieveClaimForOrder(gomock.Any(), &api.RetrieveClaimForOrderRequest{
						OrderOrganizationSerialNumber: "00000000000000000001",
						OrderReference:                "test-ref-123",
					}).
					Return(nil, errors.New("something went wrong"))
			},
		},

		"error_when_retrieving_invalid_jwt": {
			wantErr:            true,
			wantHTTPStatusCode: http.StatusInternalServerError,
			wantMessage:        "failed to parse JWT\n",
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "00000000000000000001")
				r.Header.Add("X-NLX-Request-Order-Reference", "test-ref-123")
			},
			setup: func(client *mock.MockManagementClient, plugin *DelegationPlugin) {
				client.EXPECT().
					RetrieveClaimForOrder(gomock.Any(), &api.RetrieveClaimForOrderRequest{
						OrderOrganizationSerialNumber: "00000000000000000001",
						OrderReference:                "test-ref-123",
					}).
					Return(&api.RetrieveClaimForOrderResponse{
						Claim: "invalid_jwt",
					}, nil)
			},
		},

		"missing_claim_results_in_requesting_a_claim": {
			wantHTTPStatusCode: http.StatusOK,
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "00000000000000000001")
				r.Header.Add("X-NLX-Request-Order-Reference", "test-ref-123")
			},
			setup: func(client *mock.MockManagementClient, plugin *DelegationPlugin) {
				client.EXPECT().
					RetrieveClaimForOrder(gomock.Any(), &api.RetrieveClaimForOrderRequest{
						OrderOrganizationSerialNumber: "00000000000000000001",
						OrderReference:                "test-ref-123",
					}).
					Return(&api.RetrieveClaimForOrderResponse{
						Claim: testToken,
					}, nil)
			},
		},

		"order_has_been_revoked": {
			wantErr:            true,
			wantHTTPStatusCode: http.StatusUnauthorized,
			wantMessage:        "order is revoked\n",
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "00000000000000000001")
				r.Header.Add("X-NLX-Request-Order-Reference", "test-ref-123")
			},
			setup: func(client *mock.MockManagementClient, plugin *DelegationPlugin) {
				client.EXPECT().
					RetrieveClaimForOrder(gomock.Any(), &api.RetrieveClaimForOrderRequest{
						OrderOrganizationSerialNumber: "00000000000000000001",
						OrderReference:                "test-ref-123",
					}).
					Return(nil, status.Error(codes.Unauthenticated, "order is revoked"))
			},
		},

		"invalid_claim_results_in_requesting_a_new_claim": {
			wantHTTPStatusCode: http.StatusOK,
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "00000000000000000001")
				r.Header.Add("X-NLX-Request-Order-Reference", "test-ref-123")
			},
			setup: func(client *mock.MockManagementClient, plugin *DelegationPlugin) {
				client.EXPECT().
					RetrieveClaimForOrder(gomock.Any(), &api.RetrieveClaimForOrderRequest{
						OrderOrganizationSerialNumber: "00000000000000000001",
						OrderReference:                "test-ref-123",
					}).
					Return(&api.RetrieveClaimForOrderResponse{
						Claim: testToken,
					}, nil)

				plugin.claims.Store("00000000000000000001/test-ref-123", &claimData{
					Raw: testToken,
					JWTClaims: delegation.JWTClaims{
						StandardClaims: jwt.StandardClaims{
							ExpiresAt: 1,
						},
						Delegatee:      "00000000000000000001",
						OrderReference: "test-ref-123",
					},
				})
			},
		},

		"required_headers_with_valid_claims_succeeds": {
			wantHTTPStatusCode: http.StatusOK,
			setHeaders: func(r *http.Request) {
				r.Header.Add("X-NLX-Request-Delegator", "00000000000000000001")
				r.Header.Add("X-NLX-Request-Order-Reference", "test-ref-123")
			},
			setup: func(client *mock.MockManagementClient, plugin *DelegationPlugin) {
				plugin.claims.Store("00000000000000000001/test-ref-123", &claimData{
					Raw: "claim",
					JWTClaims: delegation.JWTClaims{
						StandardClaims: jwt.StandardClaims{},
						Delegatee:      "00000000000000000001",
						OrderReference: "test-ref-123",
					},
				})
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			client := mock.NewMockManagementClient(ctrl)
			context := fakeContext(&Destination{})

			tt.setHeaders(context.Request)

			plugin := NewDelegationPlugin(client)

			if tt.setup != nil {
				tt.setup(client, plugin)
			}

			err := plugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()

			defer response.Body.Close()

			contents, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			if tt.wantErr {
				assert.Equal(t, tt.wantMessage, string(contents))
			} else {
				assert.Equal(t, "00000000000000000001", context.LogData["delegator"])
				assert.Equal(t, "test-ref-123", context.LogData["orderReference"])
			}

			assert.Equal(t, tt.wantHTTPStatusCode, response.StatusCode)
		})
	}
}
