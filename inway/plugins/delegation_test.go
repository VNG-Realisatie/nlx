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

	validClaim, err := getJWTAsSignedString(&jwtArgs{
		delegatorCert:               delegatorCertBundle,
		delegatorSerialNumber:       delegatorCertBundle.Certificate().Subject.SerialNumber,
		delegateeSerialNumber:       delegateeCertBundle.Certificate().Subject.SerialNumber,
		serviceProviderSerialNumber: serviceProviderSerialNumber,
		serviceName:                 "mock-service",
	})

	assert.Nil(t, err)

	validClaimOtherService, err := getJWTAsSignedString(&jwtArgs{
		delegatorCert:               delegatorCertBundle,
		delegatorSerialNumber:       delegatorCertBundle.Certificate().Subject.SerialNumber,
		delegateeSerialNumber:       delegateeCertBundle.Certificate().Subject.SerialNumber,
		serviceProviderSerialNumber: serviceProviderSerialNumber,
		serviceName:                 "mock-service-other",
	})
	assert.Nil(t, err)

	validClaimOtherDelegatee, err := getJWTAsSignedString(&jwtArgs{
		delegatorCert:               delegatorCertBundle,
		delegatorSerialNumber:       delegatorCertBundle.Certificate().Subject.SerialNumber,
		delegateeSerialNumber:       "nlx-hackerman",
		serviceProviderSerialNumber: serviceProviderSerialNumber,
		serviceName:                 "mock-service",
	})
	assert.Nil(t, err)

	validClaimOtherDelegateeAndService, err := getJWTAsSignedString(&jwtArgs{
		delegatorCert:               delegatorCertBundle,
		delegatorSerialNumber:       delegatorCertBundle.Certificate().Subject.SerialNumber,
		delegateeSerialNumber:       "nlx-hackerman",
		serviceProviderSerialNumber: serviceProviderSerialNumber,
		serviceName:                 "mock-service-without-grant",
	})
	assert.Nil(t, err)

	type args struct {
		service *plugins.Service
		claim   string
	}

	tests := map[string]struct {
		args                  *args
		wantStatusCode        int
		wantMessage           string
		wantDelegationSuccess bool
	}{
		"invalid_claim_format": {
			args: &args{
				claim: "invalid-claim",
			},
			wantStatusCode:        http.StatusInternalServerError,
			wantMessage:           "nlx-inway: unable to verify claim\n",
			wantDelegationSuccess: false,
		},
		"delegatee_is_not_requesting_organization": {
			args: &args{
				claim: validClaimOtherDelegatee,
				service: &plugins.Service{
					Name: "mock-service",
				},
			},
			wantStatusCode:        http.StatusUnauthorized,
			wantMessage:           "nlx-inway: no access\n",
			wantDelegationSuccess: false,
		},
		"delegatee_does_not_have_access_to_service": {
			args: &args{
				claim: validClaimOtherDelegateeAndService,
				service: &plugins.Service{
					Name:   "mock-service-without-valid-grant",
					Grants: []*plugins.Grant{},
				},
			},
			wantStatusCode:        http.StatusUnauthorized,
			wantMessage:           "nlx-inway: no access\n",
			wantDelegationSuccess: false,
		},
		"delegatee_does_not_have_service_in_claims": {
			args: &args{
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
			},
			wantStatusCode:        http.StatusUnauthorized,
			wantMessage:           "nlx-inway: no access\n",
			wantDelegationSuccess: false,
		},
		"happy_flow": {
			args: &args{
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
			},
			wantStatusCode:        http.StatusOK,
			wantDelegationSuccess: true,
		},
		"happy_flow_without_delegation": {
			args: &args{
				claim: "",
			},
			wantStatusCode:        http.StatusOK,
			wantDelegationSuccess: false,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&plugins.Destination{
				Service:      tt.args.service,
				Organization: serviceProviderSerialNumber,
			}, nil, &plugins.AuthInfo{
				OrganizationSerialNumber: delegateeCertBundle.Certificate().Subject.SerialNumber,
				PublicKeyFingerprint:     delegatorCertBundle.PublicKeyFingerprint(),
			})

			context.Request.Header.Add("X-NLX-Request-Claim", tt.args.claim)

			err := delegationPlugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()
			defer response.Body.Close()

			contents, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantMessage, string(contents))
			assert.Equal(t, tt.wantStatusCode, response.StatusCode)

			if tt.wantDelegationSuccess {
				assert.Equal(t, delegatorCertBundle.Certificate().Subject.SerialNumber, context.LogData["delegator"])
				assert.Equal(t, "order-reference", context.LogData["orderReference"])

				assert.Equal(t, delegatorCertBundle.Certificate().Subject.SerialNumber, context.AuthInfo.OrganizationSerialNumber)
				assert.Equal(t, delegatorCertBundle.PublicKeyFingerprint(), context.AuthInfo.PublicKeyFingerprint)
			}
		})
	}
}

type jwtArgs struct {
	delegatorCert               *common_tls.CertificateBundle
	delegatorSerialNumber       string
	delegateeSerialNumber       string
	serviceProviderSerialNumber string
	serviceName                 string
}

func getJWTAsSignedString(args *jwtArgs) (string, error) {
	claims := delegation.JWTClaims{
		Delegatee:      args.delegateeSerialNumber,
		OrderReference: "order-reference",
		AccessProofs: []*delegation.AccessProof{
			{
				OrganizationSerialNumber: args.delegatorSerialNumber,
				ServiceName:              args.serviceName,
				PublicKeyFingerprint:     args.delegatorCert.PublicKeyFingerprint(),
			},
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			Issuer:    args.delegatorSerialNumber,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedString, err := token.SignedString(args.delegatorCert.PrivateKey())
	if err != nil {
		return "", err
	}

	println(signedString)

	return signedString, nil
}
