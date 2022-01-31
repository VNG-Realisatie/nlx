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

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/common/delegation"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/plugins"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // this is a test
func TestDelegationPlugin(t *testing.T) {
	delegationPlugin := plugins.NewDelegationPlugin()

	const (
		serviceProviderSerialNumber = "00000000000000000099"
	)

	pkiDir := filepath.Join("..", "..", "testing", "pki")

	delegatorCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	delegateeCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
	require.NoError(t, err)

	delegatorCertBundle.PublicKey()

	certPEM, err := delegatorCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	validClaim, err := getJWTAsSignedString(delegatorCertBundle, delegateeCertBundle.Certificate().Subject.SerialNumber, serviceProviderSerialNumber, "mock-service")
	assert.Nil(t, err)

	validClaimOtherService, err := getJWTAsSignedString(delegatorCertBundle, delegateeCertBundle.Certificate().Subject.SerialNumber, serviceProviderSerialNumber, "mock-service-other")
	assert.Nil(t, err)

	validClaimOtherDelegatee, err := getJWTAsSignedString(delegatorCertBundle, "nlx-hackerman", serviceProviderSerialNumber, "mock-service")
	assert.Nil(t, err)

	validClaimOtherDelegateeAndService, err := getJWTAsSignedString(delegatorCertBundle, "nlx-hackerman", serviceProviderSerialNumber, "mock-service-without-valid-grant")
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
						OrganizationSerialNumber: delegatorCertBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM:             certPEM,
						PublicKeyFingerprint:     delegatorCertBundle.PublicKeyFingerprint(),
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
						OrganizationSerialNumber: delegatorCertBundle.Certificate().Subject.SerialNumber,
						PublicKeyPEM:             certPEM,
						PublicKeyFingerprint:     delegatorCertBundle.PublicKeyFingerprint(),
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
				Organization: serviceProviderSerialNumber,
			}, nil, &plugins.AuthInfo{
				OrganizationSerialNumber: delegateeCertBundle.Certificate().Subject.SerialNumber,
				PublicKeyFingerprint:     delegatorCertBundle.PublicKeyFingerprint(),
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
				assert.Equal(t, delegatorCertBundle.Certificate().Subject.SerialNumber, context.LogData["delegator"])
				assert.Equal(t, "order-reference", context.LogData["orderReference"])

				assert.Equal(t, delegatorCertBundle.Certificate().Subject.SerialNumber, context.AuthInfo.OrganizationSerialNumber)
				assert.Equal(t, delegatorCertBundle.PublicKeyFingerprint(), context.AuthInfo.PublicKeyFingerprint)
			}
		})
	}
}

// nolint:unparam // we want to keep the param name for readability
func getJWTAsSignedString(delegatorOrgCert *common_tls.CertificateBundle, delegateeSerialNumber, serviceProviderSerialNumber, service string) (string, error) {
	claims := delegation.JWTClaims{
		Delegatee:      delegateeSerialNumber,
		OrderReference: "order-reference",
		Services: []delegation.Service{
			{
				OrganizationSerialNumber: serviceProviderSerialNumber,
				Service:                  service,
			},
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    delegatorOrgCert.Certificate().Subject.SerialNumber,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedString, err := token.SignedString(delegatorOrgCert.PrivateKey())
	if err != nil {
		return "", err
	}

	return signedString, nil
}
