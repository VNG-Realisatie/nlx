// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi_test

import (
	"bytes"
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

	"go.nlx.io/nlx/common/derrsa"
	insightapi "go.nlx.io/nlx/insight-api"
	insightconfig "go.nlx.io/nlx/insight-api/config"
	"go.nlx.io/nlx/insight-api/irma"
	mock "go.nlx.io/nlx/insight-api/mock"
)

const publicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyJJQqF7nvhfWHLhnVkCfM0VPq8YEoyjQsDnm2RMJXSJe2pkUWRbFyqdR6bgmv1ZEtduWMhbKVQHcWE8ZfaYSQmVe92obSwOMYcXL0VS+ouTD0vjAWc0oucZ0U1adndMykgugi9PaWf0I7nCEgYRlsf62sXoF+YqwPiWLPPdzsoBDsmtWOYGfSxtJEucXXbbjiBvRsgipHqCwYsuXEc5TDXqE2iyrgDwVrme9vXxAHGRQgaISUTkOZUcVg3zkGx4SreGz5Hh1iL0VPOA0y4SCtl0a/6ccWxnapfZm03ygoxpsMoYhkFKRP32BYYorPMglck3IeW3fp24gZ7TnqeF29QIDAQAB"
const privateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC05Ac22sFpvXDYaQA/B/+WcXv8EPiPyYs1ZAcQ8ZG7eMp9kOv8+JQEXW2CbesVWw0QYQICDATrdCs6niC2qydOZYwWUQjt13JVekT9PvDi8dP5EpMyWaiDe2rDj3CGgZCXIwiMTQ6I8OUxUCPt20M7I2vQce+Alh+2/XvYmflop8n6BuEX+pRJi9Vw4X1i+rplGzdfDq6c5kdRTYCbcL2aWjgG85mafkRqdrPb7NuZDGB9p6CLjzkHrYiAKsazGH1A1t7Y4bhngKqQrMN4k2MNHbIY9D8cawMWkPMtI1vBxZLmIRV3NpWasgvzKMMpnok7xivmVgU8FPUQkZgA8BBZAgMBAAECggEAK7hxhfCZjtUa0TOPu6xTOilzrhjr+tTbsKvciVvZvVYUmwTMBPiPzf5G2Z93klHPdoX71kLRbZdGW4Sco4n6lhg1I6+yWMoZ+E71HcB4uGF6ulii+yhwclcCFwI0UE5AhEcTadW2DaMrwh98j6DPPxvwkxD2sj2WrMPXiyKsBX8f4mx0asgkQVfFC8tMzgfLj+tDiovOSqhzxkpYDHxp8ACdC+a4+81HBfAgkmR2x+DIOD6o1mW0vnlf2BFo6ZcoUXwqtFXK8xGBKPFpOfPB+hg85gYVStD0vMbTgFmeH5HKTmpZmhJqY4BTSEPGf3JP0/tddkT81IoUwNFS6JQmAQKBgQDuNg7xwxy4f4SL/cNKeKbd/DtIUQPdDuKWyHaM2pctrL5U5Kf26vWV13oqYCJdC3/odZVTnAzoOi8TAay8g1C31D/BCgUt1IW7077Sgb5Jsazmpco6DnZj+ufpKnp/ZHzSNP6dgI/dQEIfG9DBT/xqGSO1bemwm1goRJz7//zMTQKBgQDCZik218RGLFr3KWpXvgj5gVSABxIsuyhvIfPwzGoCBb7YEjPEQtOgTMtM7wn1TY5HA05GY6M8DFxaUbyaCor5AG3N5ZvoCfysoTgezw8WjCXzTNhh308iP2q3JZ2B+C2zPf0aSK/1BntzryYewnoNcyHsGb28KrXtM88N4UnqPQKBgCOjiNbY1xovUdhT7fzdUjHSA9iM7mQLTxE6CqqGJaoatxsiXpLNklKJu2hNm7aJ+uf/d4jbxv6Tfel9DafiiZgHNEagRigWLK/uPRVnfd2urGyRj1DiSwooRrwWs98NXLNiZFmSG3QBoiLfWXsiiWQiQLprKFRY2Xak1UvKf7rdAoGAB+0CYSoK5pGIY+tcWpd05jdPqqifJRO8YkuQFpE/ATYawdR8J9RRrId1An38efPfiSWpW1VUom4eldAfUGh9oglScMKbyKofkyo/j4IBq3mrUnAfol3obA0J3M27zkDAHD66wweTpPnOrrjFZRuovkOjbmzeP32+TR1/o6E70kECgYA2cBAFXM0eBIavlK96Pz1Yl90aFDai43sauZG9KUjjjF1Fcpn7+DaHK8H6l5DiqqUe6Ll2coRD9VZoLmQEIH5ON3I1OfioiKUXgqba1JXQO7TOkC8lFO2znpzxrdbyaJG2nw6SnjqEIf/4lyWi5oHZBvTwMVhnucBQ4c3cCwptOA=="

func TestNewInsightAPI(t *testing.T) {
	logger := zap.NewNop()
	ctrl := gomock.NewController(t)
	mockInsightLogFetcher := mock.NewMockInsightLogFetcher(ctrl)
	mockIrmaHandler := mock.NewMockJWTHandler(ctrl)
	config, err := insightconfig.LoadInsightConfig(logger, "../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	insightAPI, err := insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, "invalidKey", publicKey)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error decoding private key:")
	assert.Nil(t, insightAPI)

	insightAPI, err = insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, privateKey, "invalidKey")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error decoding public key:")
	assert.Nil(t, insightAPI)

	insightAPI, err = insightapi.NewInsightAPI(logger, config, mockIrmaHandler, mockInsightLogFetcher, privateKey, publicKey)
	assert.Nil(t, err)
	assert.NotNil(t, insightAPI)

}

func TestListDataSubjectHandler(t *testing.T) {
	logger := zap.NewNop()
	ctrl := gomock.NewController(t)
	mockInsightLogFetcher := mock.NewMockInsightLogFetcher(ctrl)

	mockIrmaHandler := mock.NewMockJWTHandler(ctrl)
	mockIrmaHandler.EXPECT().GenerateAndSignJWT(gomock.Any(), gomock.Any(), gomock.Any()).Return("generatedjwt", nil)

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

	expectedPrivateKey, err := derrsa.DecodeDEREncodedRSAPrivateKey(bytes.NewBufferString(privateKey))
	assert.Nil(t, err)
	tests := []struct {
		request               *insightapi.GenerateJWTRequest
		statuscode            int
		response              string
		jwtGenerationResponse func()
	}{
		{nil,
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
		{
			&insightapi.GenerateJWTRequest{
				DataSubjects: []string{"kenteken", "burgerservicenummer"},
			},
			500,
			"failed to generate JWT\n",
			func() {
				mockIrmaHandler.EXPECT().GenerateAndSignJWT(gomock.Any(), "insight", expectedPrivateKey).Return("", fmt.Errorf("JWT generation failed"))
			},
		},
		{
			&insightapi.GenerateJWTRequest{
				DataSubjects: []string{"kenteken", "burgerservicenummer"},
			},
			200,
			"generatedjwt",
			func() {
				mockIrmaHandler.EXPECT().GenerateAndSignJWT(gomock.Any(), "insight", expectedPrivateKey).Return("generatedjwt", nil)
			},
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
	expectedPublicKey, err := derrsa.DecodeDEREncodedRSAPublicKey(bytes.NewBufferString(publicKey))
	assert.Nil(t, err)

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
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, expectedPublicKey).Return(nil, &irma.VerificationResultClaims{}, fmt.Errorf("JWT verification failed"))
			},
			nil,
		},
		{
			"page=1&rowsPerPage=10",
			400,
			"failed to retrieve log records\n",
			func() {
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, expectedPublicKey).Return(nil, &irma.VerificationResultClaims{}, nil)
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
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, expectedPublicKey).Return(nil, &irma.VerificationResultClaims{}, nil)
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
				mockIrmaHandler.EXPECT().VerifyIRMAVerificationResult(expectedJWTBytes, expectedPublicKey).Return(nil, &irma.VerificationResultClaims{}, nil)
			},
			func() {
				mockInsightLogFetcher.EXPECT().GetLogRecords(10, 1, expectedDataSubjects, gomock.Any()).Return(nil, fmt.Errorf("error getting log records"))
			},
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
