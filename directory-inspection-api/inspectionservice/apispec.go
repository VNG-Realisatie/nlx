// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func getInwayAPISpecs(h *http.Client, inwayAddress, serviceName string) ([]byte, error) {
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
