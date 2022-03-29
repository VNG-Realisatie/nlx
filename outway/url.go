// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package outway

import (
	"fmt"
	"net/url"
	"strings"

	"go.nlx.io/nlx/outway/plugins"
)

const nlxDomain = "services.nlx.local"
const nlxHostNameParts = 5
const nlxPathParts = 2

func isNLXUrl(destinationURL *url.URL) bool {
	return strings.HasSuffix(destinationURL.Hostname(), nlxDomain)
}

func parseURLPath(urlPath string) (*plugins.Destination, error) {
	pathParts := strings.SplitN(strings.TrimPrefix(urlPath, "/"), "/", nlxPathParts)
	if len(pathParts) != nlxPathParts {
		return nil, fmt.Errorf("invalid path in url expecting: /serialNumber/service")
	}

	organizationSerialNumber := pathParts[0]

	indexSlash := strings.Index(pathParts[1], "/")

	var serviceName = pathParts[1]

	var path = ""

	if indexSlash > -1 {
		serviceName = pathParts[1][0:indexSlash]
		path = pathParts[1][indexSlash:]
	}

	return &plugins.Destination{
		OrganizationSerialNumber: organizationSerialNumber,
		Service:                  serviceName,
		Path:                     path,
	}, nil
}

func parseLocalNLXURL(destinationURL *url.URL) (*plugins.Destination, error) {
	hostNameParts := strings.Split(destinationURL.Hostname(), ".")
	if len(hostNameParts) != nlxHostNameParts {
		return nil, fmt.Errorf("invalid hostname expecting: service.serialNumber.services.nlx.local")
	}

	return &plugins.Destination{
		Service:                  hostNameParts[0],
		OrganizationSerialNumber: hostNameParts[1],
		Path:                     destinationURL.Path,
	}, nil
}
