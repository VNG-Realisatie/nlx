package verwerkingenlogging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VerwerkingenLoggingAPI interface {
	GetToken(request *GetTokenRequest) (*GetTokenResponse, error)
	WriteLog(request *WriteLogRequest) error
}

type VerwerkingenLoggingAPIHTTPClient struct {
	baseURL string
}

func NewVerwerkingenLoggingHTTPClient(baseURL string) (*VerwerkingenLoggingAPIHTTPClient, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("baseURL is empty")
	}

	return &VerwerkingenLoggingAPIHTTPClient{
		baseURL: baseURL,
	}, nil
}

func (v *VerwerkingenLoggingAPIHTTPClient) GetToken(request *GetTokenRequest) (*GetTokenResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/token", v.baseURL), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	getTokenResponse := &GetTokenResponse{}

	err = json.Unmarshal(body, getTokenResponse)
	if err != nil {
		return nil, err
	}

	return getTokenResponse, nil
}

func (v *VerwerkingenLoggingAPIHTTPClient) WriteLog(request *WriteLogRequest) error {
	responseGetToken, err := v.GetToken(&GetTokenRequest{
		Scopes: []string{"read:confidential", "update:confidential", "create:confidential", "read:normal", "update:normal", "create:normal"},
	})
	if err != nil {
		return fmt.Errorf("error getting token: %s", err)
	}

	if request.Vertrouwelijkheid == "" {
		return fmt.Errorf("vertrouwelijkheid cannot be empty")
	}

	if request.Tijdstip.IsZero() {
		return fmt.Errorf("tijdstip is invalid")
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/verwerkingsacties", v.baseURL), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", responseGetToken.Access))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	println(string(content))

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("request failed. status code: %d", res.StatusCode)
	}

	return nil
}
