// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/config"
	"go.nlx.io/nlx/management-api/api"
	mock_api "go.nlx.io/nlx/management-api/api/mock"
)

type inwayMocks struct {
	dc *mock_api.MockDelegationClient
}

func newTestEnv(t *testing.T, cert *common_tls.CertificateBundle) (proxy, mock *httptest.Server, mocks *inwayMocks) {
	// Mock endpoint (service)
	mockEndPoint := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	// defer is missing do this in you test!!
	// defer mockEndPoint.Close()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks = &inwayMocks{
		dc: mock_api.NewMockDelegationClient(ctrl),
	}

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-whitelist"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "whitelist",
		},
		AuthorizationWhitelist: []config.AuthorizationWhitelistItem{{OrganizationName: "nlx-test"}},
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
	iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", "localhost:1813", cert, "localhost:1815")
	assert.Nil(t, err)

	iw.delegationClient = mocks.dc

	endPoints := []ServiceEndpoint{}
	// Add service endpoints
	for serviceName := range serviceConfig.Services {
		serviceDetails := serviceConfig.Services[serviceName]
		endpoint, errr := iw.NewHTTPServiceEndpoint(serviceName, &serviceDetails, nil)
		if errr != nil {
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

func TestInwayProxyRequest(t *testing.T) {
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	proxyRequestMockServer, mockEndPoint, _ := newTestEnv(t, cert)
	proxyRequestMockServer.StartTLS()

	defer proxyRequestMockServer.Close()
	defer mockEndPoint.Close()

	client := setupClient(cert)
	//nolint:dupl // this is a test
	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-service-whitelist/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-service-whitelist-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-inway: permission denied, organization \"nlx-test\" or public key \"60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service-unspecified-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-inway: permission denied, organization \"nlx-test\" or public key \"60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx-inway: invalid path in url\n"},
		{fmt.Sprintf("%s/mock-service/fictive", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx-inway: no endpoint for service\n"},
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "", http.StatusBadRequest, "nlx-inway: missing logrecord id\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.url, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.url, nil)
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

func TestInwayProxyDelegatedRequest(t *testing.T) {
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	tests := map[string]struct {
		setup        func(*inwayMocks)
		statusCode   int
		errorMessage string
	}{
		"error_verify_claim_response_unauthenticated": {
			func(m *inwayMocks) {
				m.dc.EXPECT().
					VerifyClaim(gomock.Any(), &api.VerifyClaimRequest{
						Claim:       "mock-claim",
						ServiceName: "mock-service-public",
					}).
					Return(nil, status.Error(codes.Unauthenticated, "arbitrary error"))
			},
			http.StatusUnauthorized,
			"nlx-inway: claim is invalid. error: arbitrary error\n",
		},
		"error_verify_claim_response_internal": {
			func(m *inwayMocks) {
				m.dc.EXPECT().
					VerifyClaim(gomock.Any(), &api.VerifyClaimRequest{
						Claim:       "mock-claim",
						ServiceName: "mock-service-public",
					}).
					Return(nil, status.Error(codes.Internal, "arbitrary error"))
			},
			http.StatusInternalServerError,
			"nlx-inway: failed to verify claim\n",
		},
		"happy_flow": {
			func(m *inwayMocks) {
				m.dc.EXPECT().
					VerifyClaim(gomock.Any(), &api.VerifyClaimRequest{
						Claim:       "mock-claim",
						ServiceName: "mock-service-public",
					}).
					Return(&api.VerifyClaimResponse{
						OrderOrganizationName: "test",
						OrderReference:        "test",
					}, nil)
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
			req.Header.Add("X-NLX-Request-Claim", "mock-claim")

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

func TestInwayNoOrgProxyRequest(t *testing.T) {
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	certNoOrg, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-without-name-chain.pem"),
		filepath.Join(pkiDir, "org-without-name-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	// Clients with no organization specified in the certificate
	// should not be allowed on the nlx network.
	proxyRequestMockServer, mockEndPoint, _ := newTestEnv(t, cert)
	proxyRequestMockServer.StartTLS()

	defer proxyRequestMockServer.Close()
	defer mockEndPoint.Close()

	//nolint:dupl // this is atest
	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, ""},
		{fmt.Sprintf("%s/mock-service-whitelist/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, ""},
		{fmt.Sprintf("%s/mock-service-whitelist-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-outway: could not handle your request, organization \"nlx-test\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service-unspecified-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-outway: could not handle your request, organization \"nlx-test\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: invalid path in url\n"},
		{fmt.Sprintf("%s/mock-service/fictive", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: no endpoint for service\n"},
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "", http.StatusBadRequest, "nlx-outway: missing logrecord id\n"},
	}

	noOrgClient := setupClient(certNoOrg)

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		assert.Nil(t, err)

		req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
		resp, err := noOrgClient.Do(req)
		assert.Nil(t, err)

		defer resp.Body.Close()

		if resp.StatusCode != 400 {
			t.Fatalf(
				`result: "%d" for url "%s", expected http status code : "%d"`,
				resp.StatusCode, test.url, 400)
		}
	}
}
