// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"gopkg.in/yaml.v3"
)

type openAPIVersion struct {
	OpenAPI string `json:"openapi" yaml:"openapi"`
	Swagger string `json:"swagger" yaml:"swagger"`
}

const (
	openAPI2 = "OpenAPI2"
	openAPI3 = "OpenAPI3"
)

func ParseAPISpectType(data []byte) (string, error) {
	versionCheck, err := parseVersion(data)
	if err != nil {
		return "", err
	}

	var versionValue string
	if versionCheck.OpenAPI != "" {
		versionValue = versionCheck.OpenAPI
	} else if versionCheck.Swagger != "" {
		versionValue = versionCheck.Swagger
	}

	switch versionValue {
	case "2.0":
		return openAPI2, nil
	case "3.0.0", "3.0.1", "3.0.2":
		return openAPI3, nil
	}

	return "", fmt.Errorf("documentation format is neither openAPI2 or openAPI3")
}

func parseVersion(data []byte) (*openAPIVersion, error) {
	if len(data) == 0 {
		return nil, errors.New("empty input")
	}

	version := &openAPIVersion{}

	// If it looks like JSON, try to parse it as
	if data[0] == '{' {
		err := json.Unmarshal(data, version)
		if err == nil {
			return version, nil
		}
	}

	// JSON failed, try it as YAML
	err := yaml.Unmarshal(data, version)
	if err == nil {
		return version, nil
	}

	return nil, errors.New("unable to parse version")
}

func getAPISpecsTypeViaInway(httpClient *http.Client, inwayAddress, serviceName string) (string, error) {
	data, err := getAPISpecsViaInway(httpClient, inwayAddress, serviceName)
	if err != nil {
		return "", err
	}

	return ParseAPISpectType(data)
}

func getAPISpecsViaInway(h *http.Client, inwayAddress, serviceName string) ([]byte, error) {
	inwayURL := url.URL{
		Scheme: "https",
		Host:   inwayAddress,
		Path:   path.Join("/.nlx/api-spec-doc/", serviceName),
	}

	res, err := h.Get(inwayURL.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
