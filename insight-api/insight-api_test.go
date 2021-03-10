// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	insightapi "go.nlx.io/nlx/insight-api"
	insightconfig "go.nlx.io/nlx/insight-api/config"
	"go.nlx.io/nlx/insight-api/irma"
	mock "go.nlx.io/nlx/insight-api/mock"
)

var (
	privateKey, _ = rsa.GenerateKey(rand.Reader, 4096)
	publicKey, _  = privateKey.Public().(*rsa.PublicKey)
)

func TestNewInsightAPI(t *testing.T) {
	logger := zap.NewNop()
	ctrl := gomock.NewController(t)
	mockInsightLogFetcher := mock.NewMockInsightLogFetcher(ctrl)
	mockIrmaHandler := mock.NewMockJWTHandler(ctrl)
	config, err := insightconfig.LoadInsightConfig(logger, "../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	insightAPI, err := insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, privateKey, publicKey)
	assert.Nil(t, err)
	assert.NotNil(t, insightAPI)
}

func TestListDataSubjectHandler(t *testing.T) {
	logger := zap.NewNop()
	ctrl := gomock.NewController(t)
	mockInsightLogFetcher := mock.NewMockInsightLogFetcher(ctrl)

	mockIrmaHandler := mock.NewMockJWTHandler(ctrl)

	config, err := insightconfig.LoadInsightConfig(logger, "../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	insightAPI, err := insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, privateKey, publicKey)
	assert.Nil(t, err)
	assert.NotNil(t, insightAPI)

	mockServer := httptest.NewServer(insightAPI)
	defer mockServer.Close()

	client := http.Client{}

	req, err := http.NewRequest("GET", mockServer.URL+"/getDataSubjects", nil)
	if err != nil {
		t.Fatal("error creating http request", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("error doing http request", err)
	}
	defer resp.Body.Close()

	expectedResponse := &insightapi.GetDataSubjectsResponse{}
	expectedResponse.DataSubjects = make(map[string]insightapi.DataSubject)
	expectedResponse.DataSubjects["burgerservicenummer"] = insightapi.DataSubject{
		Label: "Burgerservicenummer",
	}
	expectedResponse.DataSubjects["kenteken"] = insightapi.DataSubject{
		Label: "Kenteken",
	}
	getDataSubjectsResponse := &insightapi.GetDataSubjectsResponse{}
	err = json.NewDecoder(resp.Body).Decode(getDataSubjectsResponse)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedResponse, getDataSubjectsResponse)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("content-type"))
}

func TestGenerateJWTHandler(t *testing.T) {
	logger := zap.NewNop()

	ctrl := gomock.NewController(t)
	mockInsightLogFetcher := mock.NewMockInsightLogFetcher(ctrl)

	mockIrmaHandler := mock.NewMockJWTHandler(ctrl)

	config, err := insightconfig.LoadInsightConfig(logger, "../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	insightAPI, err := insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, privateKey, publicKey)
	assert.Nil(t, err)
	assert.NotNil(t, insightAPI)

	mockServer := httptest.NewServer(insightAPI)
	defer mockServer.Close()

	client := http.Client{}

	tests := []struct {
		request               *insightapi.GenerateJWTRequest
		statuscode            int
		response              string
		jwtGenerationResponse func()
	}{
		{
			nil,
			400,
			"incorrect request data\n",
			nil,
		},
		{
			&insightapi.GenerateJWTRequest{
				DataSubjects: []string{"non-existing"},
			},
			400,
			"incorrect dataSubject requested\n",
			nil,
		},
	}

	for _, test := range tests {
		var jsonBytes []byte
		if test.request != nil {
			jsonBytes, err = json.Marshal(test.request)
			assert.Nil(t, err)
		}

		if test.jwtGenerationResponse != nil {
			test.jwtGenerationResponse()
		}

		req, err := http.NewRequest("POST", mockServer.URL+"/generateJWT", bytes.NewReader(jsonBytes))
		if err != nil {
			t.Fatal("error creating http request", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("error doing http request", err)
		}

		assert.Equal(t, test.statuscode, resp.StatusCode)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		assert.Nil(t, err)

		assert.Equal(t, test.response, string(bodyBytes))

		if resp.StatusCode == http.StatusOK {
			assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("content-type"))
		}
	}
}

//nolint:funlen // this is a test
func TestLogFetcherHandler(t *testing.T) {
	logger := zap.NewNop()

	ctrlIRMA := gomock.NewController(t)
	ctrlLogFetcher := gomock.NewController(t)
	mockInsightLogFetcher := mock.NewMockInsightLogFetcher(ctrlLogFetcher)
	mockIrmaHandler := mock.NewMockJWTHandler(ctrlIRMA)

	config, err := insightconfig.LoadInsightConfig(logger, "../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	expectedDataSubjects := make(map[string][]string)

	for dataSubjectKey, dataSubjectProperties := range config.DataSubjects {
		for _, irmaAttribute := range dataSubjectProperties.IrmaAttributes {
			expectedDataSubjects[string(irmaAttribute)] = append(expectedDataSubjects[string(irmaAttribute)], dataSubjectKey)
		}
	}

	insightAPI, err := insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, privateKey, publicKey)
	assert.Nil(t, err)
	assert.NotNil(t, insightAPI)

	mockServer := httptest.NewServer(insightAPI)
	defer mockServer.Close()

	client := http.Client{}

	expectedGetLogsResponse := &insightapi.GetLogRecordsResponse{
		Page:        1,
		RowCount:    1,
		RowsPerPage: 1,
		Records:     []*insightapi.Record{},
	}

	jsonBytes, err := json.Marshal(expectedGetLogsResponse)
	assert.Nil(t, err)

	expectedGetLogsResponseString := fmt.Sprintf("%s\n", string(jsonBytes))

	jwt := "dummyjwt"
	expectedJWTBytes := []byte(jwt)

	tests := []struct {
		urlParameters                string
		statuscode                   int
		response                     string
		verifyIRMAVerificationResult func()
		logFetcherResults            func()
	}{
		{
			"page=1&rowsPerPage=10",
			400,
			"invalid irma jwt\n",
			func() {
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, publicKey).Return(nil, &irma.VerificationResultClaims{}, fmt.Errorf("JWT verification failed"))
			},
			nil,
		},
		{
			"page=1&rowsPerPage=10",
			400,
			"failed to retrieve log records\n",
			func() {
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, publicKey).Return(nil, &irma.VerificationResultClaims{}, nil)
			},
			func() {
				mockInsightLogFetcher.EXPECT().GetLogRecords(10, 1, expectedDataSubjects, gomock.Any()).Return(nil, fmt.Errorf("error getting log records"))
			},
		},
		{
			"page=1&rowsPerPage=10",
			200,
			expectedGetLogsResponseString,
			func() {
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, publicKey).Return(nil, &irma.VerificationResultClaims{}, nil)
			},
			func() {
				mockInsightLogFetcher.EXPECT().GetLogRecords(10, 1, expectedDataSubjects, gomock.Any()).Return(expectedGetLogsResponse, nil)
			},
		},
		{
			"page==1&rowsPerPage=10",
			400,
			"failed to parse URL values\n",
			func() {
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, publicKey).Return(nil, &irma.VerificationResultClaims{}, nil)
			},
			func() {},
		},
	}

	for _, test := range tests {
		if test.verifyIRMAVerificationResult != nil {
			test.verifyIRMAVerificationResult()
		}

		if test.logFetcherResults != nil {
			test.logFetcherResults()
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/fetch?%s", mockServer.URL, test.urlParameters), strings.NewReader(jwt))
		if err != nil {
			t.Fatal("error creating http request", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("error doing http request", err)
		}

		assert.Equal(t, test.statuscode, resp.StatusCode)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		assert.Nil(t, err)
		assert.Equal(t, test.response, string(bodyBytes))

		if resp.StatusCode == http.StatusOK {
			assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("content-type"))
		}
	}
}

func TestHappyOptionsHandler(t *testing.T) {
	logger := zap.NewNop()

	ctrlIRMA := gomock.NewController(t)
	ctrlLogFetcher := gomock.NewController(t)
	mockInsightLogFetcher := mock.NewMockInsightLogFetcher(ctrlLogFetcher)
	mockIrmaHandler := mock.NewMockJWTHandler(ctrlIRMA)

	config, err := insightconfig.LoadInsightConfig(logger, "../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	insightAPI, err := insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, privateKey, publicKey)
	assert.Nil(t, err)
	assert.NotNil(t, insightAPI)

	mockServer := httptest.NewServer(insightAPI)
	defer mockServer.Close()

	client := http.Client{}
	tests := []struct {
		path       string
		statuscode int
	}{
		{
			"/fetch",
			200,
		},
		{
			"/generateJWT",
			200,
		},
		{
			"/getDataSubjects",
			200,
		},
	}

	expectedHeaders := http.Header{}
	expectedHeaders.Set("Access-Control-Allow-Origin", "*")
	expectedHeaders.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	expectedHeaders.Set("X-Content-Type-Options", "nosniff")

	for _, test := range tests {
		req, err := http.NewRequest("OPTIONS", fmt.Sprintf("%s%s", mockServer.URL, test.path), nil)
		assert.Nil(t, err)
		resp, err := client.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, test.statuscode, resp.StatusCode)

		for key, expectedValue := range expectedHeaders {
			actualValue := resp.Header[key]
			assert.Equal(t, expectedValue, actualValue)
		}

		resp.Body.Close()
	}
}
