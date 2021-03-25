// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	mock_transactionlog "go.nlx.io/nlx/common/transactionlog/mock"
	"go.nlx.io/nlx/inway/config"
	mock_api "go.nlx.io/nlx/management-api/api/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestInwayProxyRequest(t *testing.T) {
	cert := createCertBundle()

	client := setupClient(cert)
	//nolint:dupl // this is a test
	tests := []struct {
		requestPath  string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{
			"/mock-service-public/dummy",
			"dummy-ID",
			http.StatusOK,
			"",
		},
		{
			"/mock-service-whitelist/dummy",
			"dummy-ID",
			http.StatusOK,
			"",
		},
		{
			"/mock-service-whitelist-unauthorized/dummy",
			"dummy-ID",
			http.StatusForbidden,
			"nlx-inway: permission denied, organization \"nlx-test\" or public key \"60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=\" is not allowed access.\n",
		},
		{
			"/mock-service-unspecified-unauthorized/dummy",
			"dummy-ID",
			http.StatusForbidden,
			"nlx-inway: permission denied, organization \"nlx-test\" or public key \"60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=\" is not allowed access.\n",
		},
		{
			"/mock-service",
			"dummy-ID",
			http.StatusBadRequest,
			"nlx-inway: invalid path in url\n"},
		{
			"/mock-service/fictive",
			"dummy-ID",
			http.StatusBadRequest,
			"nlx-inway: no endpoint for service\n",
		},
		{
			"/mock-service-public/dummy",
			"",
			http.StatusBadRequest,
			"nlx-inway: missing logrecord id\n",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.requestPath, func(t *testing.T) {
			proxyRequestMockServer, mockEndPoint, mocks := newTestEnv(t, cert)
			proxyRequestMockServer.StartTLS()

			defer proxyRequestMockServer.Close()
			defer mockEndPoint.Close()

			mocks.tl.EXPECT().
				AddRecord(gomock.Any()).
				AnyTimes().
				Return(nil)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", proxyRequestMockServer.URL, test.requestPath), nil)
			assert.Nil(t, err)
			req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
			resp, err := client.Do(req)
			assert.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, test.statusCode, resp.StatusCode)

			defer resp.Body.Close()
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("error parsing result.body", err)
			}
			assert.Equal(t, test.errorMessage, string(bytes))
		})
	}
}

func TestDoelbinding(t *testing.T) {
	cert := createCertBundle()

	tests := map[string]struct {
		setup        func(*inwayMocks)
		statusCode   int
		errorMessage string
	}{
		"happy_flow": {
			func(m *inwayMocks) {
				m.tl.EXPECT().
					AddRecord(&transactionlog.Record{
						SrcOrganization:  "nlx-test",
						DestOrganization: "nlx-test",
						ServiceName:      "mock-service-public",
						LogrecordID:      "dummyID",
						Data: map[string]interface{}{
							"request-path":              "/dummy",
							"doelbinding-process-id":    "123456",
							"doelbinding-data-elements": "mock-element",
						},
					}).
					Return(nil)
			},
			http.StatusOK,
			"",
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			proxyRequestMockServer, mockEndPoint, mocks := newTestEnv(t, cert)
			proxyRequestMockServer.StartTLS()

			defer proxyRequestMockServer.Close()
			defer mockEndPoint.Close()

			client := setupClient(cert)

			test.setup(mocks)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), nil)
			assert.Nil(t, err)
			req.Header.Add("X-NLX-Logrecord-Id", "dummyID")
			req.Header.Add("X-NLX-Request-Process-Id", "123456")
			req.Header.Add("X-NLX-Request-Data-Elements", "mock-element")

			resp, err := client.Do(req)
			assert.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, test.statusCode, resp.StatusCode)

			defer resp.Body.Close()
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("error parsing result.body", err)
			}
			assert.Equal(t, test.errorMessage, string(bytes))
		})
	}
}

//nolint:funlen // this is a test
func TestInwayProxyDelegatedRequest(t *testing.T) {
	cert := createCertBundle()
	validClaim, err := getJWTAsSignedString(cert)
	assert.Nil(t, err)

	tests := map[string]struct {
		setup        func(*inwayMocks)
		path         string
		claim        string
		statusCode   int
		errorMessage string
	}{
		"invalid_claim_format": {
			func(m *inwayMocks) {},
			"mock-service-whitelist/dummy",
			"foo-bar-baz",
			http.StatusInternalServerError,
			"nlx-inway: unable to verify claim\n",
		},
		"delegator_does_not_have_access_to_service": {
			func(m *inwayMocks) {},
			"mock-service-whitelist-unauthorized/dummy",
			validClaim,
			http.StatusUnauthorized,
			"nlx-inway: no access\n",
		},
		"error_failed_to_write_transaction_log": {
			func(m *inwayMocks) {
				m.tl.EXPECT().
					AddRecord(gomock.Any()).
					Return(errors.New("arbitrary error"))
			},
			"mock-service-whitelist/dummy",
			validClaim,
			http.StatusInternalServerError,
			"nlx-inway: server error\n",
		},
		"happy_flow": {
			func(m *inwayMocks) {
				m.tl.EXPECT().
					AddRecord(&transactionlog.Record{
						SrcOrganization:  "nlx-test",
						DestOrganization: "nlx-test",
						ServiceName:      "mock-service-whitelist",
						LogrecordID:      "dummyID",
						Data: map[string]interface{}{
							"request-path": "/dummy",
						},
						Delegator:      "nlx-test",
						OrderReference: "order-reference",
					}).
					Return(nil)
			},
			"mock-service-whitelist/dummy",
			validClaim,
			http.StatusOK,
			"",
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			proxyRequestMockServer, mockEndPoint, mocks := newTestEnv(t, cert)
			proxyRequestMockServer.StartTLS()

			defer proxyRequestMockServer.Close()
			defer mockEndPoint.Close()

			client := setupClient(cert)

			test.setup(mocks)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", proxyRequestMockServer.URL, test.path), nil)
			assert.Nil(t, err)
			req.Header.Add("X-NLX-Logrecord-Id", "dummyID")
			req.Header.Add("X-NLX-Request-Claim", test.claim)

			resp, err := client.Do(req)
			assert.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, test.statusCode, resp.StatusCode)

			defer resp.Body.Close()
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("error parsing result.body", err)
			}
			assert.Equal(t, test.errorMessage, string(bytes))
		})
	}
}

// Clients with no organization specified in the certificate
// should not be allowed on the nlx network.
func TestInwayNoOrganizationNameInClientCertificate(t *testing.T) {
	cert := createCertBundle()

	certWithoutOrganizationName, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-without-name-chain.pem"),
		filepath.Join(pkiDir, "org-without-name-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	paths := []string{
		"/mock-service-public/dummy/",
		"/mock-service-whitelist/dummy",
		"/mock-service-whitelist-unauthorized/dummy",
		"/mock-service-unspecified-unauthorized/dummy",
		"/mock-service",
		"/mock-service-whitelist/fictive",
		"/mock-service-public/dummy",
	}

	for _, path := range paths {
		path := path

		t.Run(path, func(t *testing.T) {
			proxyRequestMockServer, mockEndPoint, _ := newTestEnv(t, cert)
			proxyRequestMockServer.StartTLS()

			defer proxyRequestMockServer.Close()
			defer mockEndPoint.Close()

			url := fmt.Sprintf("%s%s", proxyRequestMockServer.URL, path)
			req, err := http.NewRequest("GET", url, nil)
			assert.Nil(t, err)

			req.Header.Add("X-NLX-Logrecord-Id", "arbitrary-logrecord-id")

			noOrgClient := setupClient(certWithoutOrganizationName)

			resp, err := noOrgClient.Do(req)
			assert.Nil(t, err)

			defer resp.Body.Close()

			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("error parsing result.body", err)
			}

			assert.Equal(t, "nlx-inway: invalid certificate provided: missing organizations attribute in subject\n", string(bytes))
		})
	}
}

func createCertBundle() *common_tls.CertificateBundle {
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	return cert
}

func getJWTAsSignedString(orgCert *common_tls.CertificateBundle) (string, error) {
	claims := server.JWTClaims{
		Organization:   "delegatee-organization-name",
		OrderReference: "order-reference",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "nlx-test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedString, err := token.SignedString(orgCert.PrivateKey())
	if err != nil {
		return "", err
	}

	return signedString, nil
}

type inwayMocks struct {
	dc *mock_api.MockDelegationClient
	tl *mock_transactionlog.MockTransactionLogger
}

func newTestEnv(t *testing.T, cert *common_tls.CertificateBundle) (proxy, mock *httptest.Server, mocks *inwayMocks) {
	mockEndPoint := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		mockEndPoint.Close()
		t.Helper()
		ctrl.Finish()
	})

	mocks = &inwayMocks{
		dc: mock_api.NewMockDelegationClient(ctrl),
		tl: mock_transactionlog.NewMockTransactionLogger(ctrl),
	}

	pem, err := cert.PublicKeyPEM()
	assert.Nil(t, err)

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-whitelist"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "whitelist",
		},
		AuthorizationWhitelist: []config.AuthorizationWhitelistItem{
			{
				OrganizationName: "nlx-test",
				PublicKeyPEM:     pem,
			},
		},
	}
	serviceConfig.Services["mock-service-whitelist-unauthorized"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "whitelist",
		},
		AuthorizationWhitelist: []config.AuthorizationWhitelistItem{{OrganizationName: "nlx-forbidden"}},
	}
	serviceConfig.Services["mock-service-unspecified-unauthorized"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "",
		},
	}
	serviceConfig.Services["mock-service-public"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "none",
		},
	}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	iw, err := NewInway(
		logger,
		mocks.tl,
		testProcess,
		"",
		"localhost:1812",
		"localhost:1813",
		cert,
		"localhost:1815",
	)
	assert.Nil(t, err)

	iw.delegationClient = mocks.dc

	endPoints := []ServiceEndpoint{}

	for serviceName := range serviceConfig.Services {
		serviceDetails := serviceConfig.Services[serviceName]
		endpoint, endpointErr := iw.NewHTTPServiceEndpoint(serviceName, &serviceDetails, nil)

		if endpointErr != nil {
			t.Fatal("failed to create service endpoint", err)
		}

		endPoints = append(endPoints, endpoint)
	}

	err = iw.SetServiceEndpoints(endPoints)
	assert.Nil(t, err)

	proxyRequestMockServer := httptest.NewUnstartedServer(http.HandlerFunc(iw.handleProxyRequest))
	proxyRequestMockServer.TLS = cert.TLSConfig(cert.WithTLSClientAuth())

	return proxyRequestMockServer, mockEndPoint, mocks
}
