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
const nlxPathParts = 3

func isNLXUrl(destinationURL *url.URL) bool {
	return strings.HasSuffix(destinationURL.Hostname(), nlxDomain)
}

func parseURLPath(urlPath string) (*plugins.Destination, error) {
	pathParts := strings.SplitN(strings.TrimPrefix(urlPath, "/"), "/", nlxPathParts)

	if len(pathParts) != nlxPathParts {
		return nil, fmt.Errorf("invalid path in url expecting: /organization/service/path")
	}

	return &plugins.Destination{
		Organization: pathParts[0],
		Service:      pathParts[1],
		Path:         pathParts[2],
	}, nil
}

func parseLocalNLXURL(destinationURL *url.URL) (*plugins.Destination, error) {
	hostNameParts := strings.Split(destinationURL.Hostname(), ".")
	if len(hostNameParts) != nlxHostNameParts {
		return nil, fmt.Errorf("invalid hostname expecting: service.organization.services.nlx.local")
	}

	return &plugins.Destination{
		Service:      hostNameParts[0],
		Organization: hostNameParts[1],
		Path:         destinationURL.Path,
	}, nil
}
