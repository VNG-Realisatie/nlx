// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type authRequest struct {
	Headers      http.Header `json:"headers"`
	Organization string      `json:"organization"`
	Service      string      `json:"service"`
}

type authResponse struct {
	Authorized bool   `json:"authorized"`
	Reason     string `json:"reason,omitempty"`
}

type authSettings struct {
	serviceURL string
	ca         *x509.CertPool
}

func (o *Outway) authorizeRequest(h http.Header, d *destination) (*authResponse, error) {
	req, err := http.NewRequest(http.MethodPost, o.authorizationSettings.serviceURL, nil)
	if err != nil {
		return nil, err
	}
	authRequest := &authRequest{
		Headers:      h,
		Organization: d.Organization,
		Service:      d.Service,
	}

	body, err := json.Marshal(authRequest)
	if err != nil {
		return nil, err
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	resp, err := o.authorizationClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authorization service return non 200 status code. status code: %d", resp.StatusCode)
	}

	authResponse := &authResponse{}
	err = json.NewDecoder(resp.Body).Decode(authResponse)
	if err != nil {
		return nil, err
	}

	return authResponse, nil
}

func (o *Outway) stripHeaders(r *http.Request, receiverOrganization string) {
	if o.organizationName != receiverOrganization {
		r.Header.Del("X-NLX-Requester-User")
		r.Header.Del("X-NLX-Requester-Claims")
		r.Header.Del("X-NLX-Request-Subject-Identifier")
		r.Header.Del("X-NLX-Request-Application-Id")
		r.Header.Del("X-NLX-Request-User-Id")
		r.Header.Del("X-NLX-Request-Data-Subject")
	}
	r.Header.Del("Proxy-Authorization")
}
