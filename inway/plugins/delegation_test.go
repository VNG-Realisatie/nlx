// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/common/httperrors"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/plugins"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // this is a test
func TestDelegationPlugin(t *testing.T) {
	delegationPlugin := plugins.NewDelegationPlugin()

	const serviceProviderSerialNumber = "00000000000000000099"

	pkiDir := filepath.Join("..", "..", "testing", "pki")

	delegatorCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	delegateeCertBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTestB)
	require.NoError(t, err)

	delegatorCertBundle.PublicKey()

	certPEM, err := delegatorCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	validClaim, err := getJWTAsSignedString(&jwtArgs{
		delegatorCert:                 delegatorCertBundle,
		delegatorSerialNumber:         delegatorCertBundle.Certificate().Subject.SerialNumber,
		delegateeSerialNumber:         delegateeCertBundle.Certificate().Subject.SerialNumber,
		delegateePublicKeyFingerprint: delegateeCertBundle.PublicKeyFingerprint(),
		serviceProviderSerialNumber:   serviceProviderSerialNumber,
		serviceName:                   "mock-service",
	})
	assert.Nil(t, err)

	type args struct {
		service                       *plugins.Service
		claim                         string
		delegateeSerialNumber         string
		delegateePublicKeyFingerprint string
	}

	tests := map[string]struct {
		args                  *args
		wantStatusCode        int
		wantErr               *httperrors.NLXNetworkError
		wantDelegationSuccess bool
	}{
		"invalid_claim_format": {
			args: &args{
				claim: "invalid-claim",
			},
			wantStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.UnableToVerifyClaim,
				Message:  "unable to verify claim",
			},
			wantDelegationSuccess: false,
		},
		"delegatee_is_not_requesting_organization": {
			args: &args{
				service: &plugins.Service{
					Name: "mock-service",
				},
				claim:                         validClaim,
				delegateePublicKeyFingerprint: "public-key-fingerprint",
				delegateeSerialNumber:         "00000000000000000099",
			},
			wantStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.RequestingOrganizationIsNotDelegatee,
				Message:  "no access. organization serialnumber does not match the delegatee organization serialnumber of the order",
			},
			wantDelegationSuccess: false,
		},
		"delegatee_pub_key_fingerprint_not_found_in_order": {
			args: &args{
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
				claim: func() string {
					claim, err := getJWTAsSignedString(&jwtArgs{
						delegatorCert:                 delegatorCertBundle,
						delegatorSerialNumber:         delegatorCertBundle.Certificate().Subject.SerialNumber,
						delegateeSerialNumber:         delegateeCertBundle.Certificate().Subject.SerialNumber,
						delegateePublicKeyFingerprint: delegatorCertBundle.PublicKeyFingerprint(),
						serviceProviderSerialNumber:   serviceProviderSerialNumber,
						serviceName:                   "mock-service",
					})
					require.NoError(t, err)

					return claim
				}(),
				delegateeSerialNumber:         delegateeCertBundle.Certificate().Subject.SerialNumber,
				delegateePublicKeyFingerprint: delegateeCertBundle.PublicKeyFingerprint(),
			},
			wantStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.RequestingOrganizationIsNotDelegatee,
				Message:  "no access. public key of the connection does not match the delegatee public key of the order",
			},
			wantDelegationSuccess: false,
		},
		"delegator_does_not_have_access_to_service": {
			args: &args{
				service: &plugins.Service{
					Name:   "mock-service-without-valid-grant",
					Grants: []*plugins.Grant{},
				},
				claim:                         validClaim,
				delegateeSerialNumber:         delegateeCertBundle.Certificate().Subject.SerialNumber,
				delegateePublicKeyFingerprint: delegateeCertBundle.PublicKeyFingerprint(),
			},
			wantStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.DelegatorDoesNotHaveAccessToService,
				Message:  "no access. delegator does not have access to the service for the public key in the claim",
			},
			wantDelegationSuccess: false,
		},
		"happy_flow": {
			args: &args{
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
				claim:                         validClaim,
				delegateeSerialNumber:         delegateeCertBundle.Certificate().Subject.SerialNumber,
				delegateePublicKeyFingerprint: delegateeCertBundle.PublicKeyFingerprint(),
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
				OrganizationSerialNumber: tt.args.delegateeSerialNumber,
				PublicKeyFingerprint:     tt.args.delegateePublicKeyFingerprint,
			})

			context.Request.Header.Add("X-NLX-Request-Claim", tt.args.claim)

			err := delegationPlugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()
			defer response.Body.Close()

			contents, err := io.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantStatusCode, response.StatusCode)

			if tt.wantDelegationSuccess {
				assert.Equal(t, delegatorCertBundle.Certificate().Subject.SerialNumber, context.LogData["delegator"])
				assert.Equal(t, "order-reference", context.LogData["orderReference"])

				assert.Equal(t, delegatorCertBundle.Certificate().Subject.SerialNumber, context.AuthInfo.OrganizationSerialNumber)
				assert.Equal(t, delegatorCertBundle.PublicKeyFingerprint(), context.AuthInfo.PublicKeyFingerprint)
			}

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			}
		})
	}
}

type jwtArgs struct {
	delegatorCert                 *common_tls.CertificateBundle
	delegatorSerialNumber         string
	delegateeSerialNumber         string
	delegateePublicKeyFingerprint string
	serviceProviderSerialNumber   string
	serviceName                   string
}

func getJWTAsSignedString(args *jwtArgs) (string, error) {
	claims := delegation.JWTClaims{
		Delegatee:                     args.delegateeSerialNumber,
		DelegateePublicKeyFingerprint: args.delegateePublicKeyFingerprint,
		OrderReference:                "order-reference",
		AccessProof: &delegation.AccessProof{
			OrganizationSerialNumber: args.delegatorSerialNumber,
			ServiceName:              args.serviceName,
			PublicKeyFingerprint:     "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
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

	return signedString, nil
}
