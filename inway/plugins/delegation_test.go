// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/delegation"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/plugins"
)

//nolint:funlen // this is a test
func TestDelegationPlugin(t *testing.T) {
	delegationPlugin := plugins.NewDelegationPlugin()

	var pkiDir = filepath.Join("..", "..", "testing", "pki")

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	cert.PublicKey()

	certPEM, _ := cert.PublicKeyPEM()
	certFingerprint := cert.PublicKeyFingerprint()

	validClaim, err := getJWTAsSignedString(cert, "delegatee-org", "mock-service")
	assert.Nil(t, err)

	validClaimOtherService, err := getJWTAsSignedString(cert, "delegatee-org", "mock-service-other")
	assert.Nil(t, err)

	validClaimOtherDelegatee, err := getJWTAsSignedString(cert, "nlx-hackerman", "mock-service")
	assert.Nil(t, err)

	validClaimOtherDelegateeAndService, err := getJWTAsSignedString(cert, "nlx-hackerman", "mock-service-without-valid-grant")
	assert.Nil(t, err)

	tests := map[string]struct {
		service            *plugins.Service
		claim              string
		expectedStatusCode int
		expectedMessage    string
		delegationSuccess  bool
	}{
		"invalid_claim_format": {
			claim:              "invalid-claim",
			expectedStatusCode: http.StatusInternalServerError,
			expectedMessage:    "nlx-inway: unable to verify claim\n",
			delegationSuccess:  false,
		},
		"delegatee_is_not_requesting_organization": {
			claim: validClaimOtherDelegatee,
			service: &plugins.Service{
				Name: "mock-service",
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedMessage:    "nlx-inway: no access\n",
			delegationSuccess:  false,
		},
		"delegatee_does_not_have_access_to_service": {
			claim: validClaimOtherDelegateeAndService,
			service: &plugins.Service{
				Name:   "mock-service-without-valid-grant",
				Grants: []*plugins.Grant{},
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedMessage:    "nlx-inway: no access\n",
			delegationSuccess:  false,
		},
		"delegatee_does_not_have_service_in_claims": {
			claim: validClaimOtherService,
			service: &plugins.Service{
				Name: "mock-service",
				Grants: []*plugins.Grant{
					{
						OrganizationName:     "issuer-org",
						PublicKeyPEM:         certPEM,
						PublicKeyFingerprint: certFingerprint,
					},
				},
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedMessage:    "nlx-inway: no access\n",
			delegationSuccess:  false,
		},
		"happy_flow": {
			claim: validClaim,
			service: &plugins.Service{
				Name: "mock-service",
				Grants: []*plugins.Grant{
					{
						OrganizationName:     "issuer-org",
						PublicKeyPEM:         certPEM,
						PublicKeyFingerprint: certFingerprint,
					},
				},
			},
			expectedStatusCode: http.StatusOK,
			delegationSuccess:  true,
		},
		"happy_flow_without_delegation": {
			claim:              "",
			expectedStatusCode: http.StatusOK,
			delegationSuccess:  false,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&plugins.Destination{
				Service:      tt.service,
				Organization: "nlx-test",
			}, nil, &plugins.AuthInfo{
				OrganizationName:     "delegatee-org",
				PublicKeyFingerprint: certFingerprint,
			})

			context.Request.Header.Add("X-NLX-Request-Claim", tt.claim)

			err := delegationPlugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()
			defer response.Body.Close()

			contents, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedMessage, string(contents))
			assert.Equal(t, tt.expectedStatusCode, response.StatusCode)

			if tt.delegationSuccess {
				assert.Equal(t, "issuer-org", context.LogData["delegator"])
				assert.Equal(t, "order-reference", context.LogData["orderReference"])

				assert.Equal(t, "issuer-org", context.AuthInfo.OrganizationName)
				assert.Equal(t, certFingerprint, context.AuthInfo.PublicKeyFingerprint)
			}
		})
	}
}

func getJWTAsSignedString(orgCert *common_tls.CertificateBundle, delegatee, service string) (string, error) {
	claims := delegation.JWTClaims{
		Delegatee:      delegatee,
		OrderReference: "order-reference",
		Services: []delegation.Service{
			{
				Service:      service,
				Organization: "nlx-test",
			},
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "issuer-org",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedString, err := token.SignedString(orgCert.PrivateKey())
	if err != nil {
		return "", err
	}

	return signedString, nil
}
