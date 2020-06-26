// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	userAgent = "nlx-management"
)

type Client interface {
	ListServices() ([]*InspectionAPIService, error)
}

type HTTPClient struct {
	client    *http.Client
	baseURL   *url.URL
	userAgent string
}

func NewClient(endpointURL string) (Client, error) {
	httpClient := &http.Client{}

	baseURL, err := url.Parse(endpointURL)
	if err != nil {
		return nil, err
	}

	baseURL.Path += "/directory/"

	c := &HTTPClient{
		client:    httpClient,
		baseURL:   baseURL,
		userAgent: userAgent,
	}

	return c, nil
}

func (d *HTTPClient) newRequest(method, pathStr string) (*http.Request, error) {
	u, err := d.baseURL.Parse(pathStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", d.userAgent)

	return req, nil
}

func (d *HTTPClient) sendRequest(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := d.client.Do(req)

	if err != nil {
		return nil, err
	}

	if c := resp.StatusCode; c < 200 || c > 299 {
		return nil, errors.New("request failed")
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}
