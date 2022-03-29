// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package outway

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/outway/plugins"
)

func TestIsNLXURL(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{
			"http://service.organization.services.nlx.local/path",
			true,
		}, {
			"http://google.nl",
			false,
		}, {
			"https://kubernetes.io/docs/tutorials/kubernetes-basics",
			false,
		},
	}

	for _, test := range tests {
		destinationURL, err := url.Parse(test.url)
		assert.Nil(t, err)

		result := isNLXUrl(destinationURL)
		assert.Equal(t, test.want, result)
	}
}

func TestParseURLPath(t *testing.T) {
	tests := map[string]struct {
		path            string
		wantDestination *plugins.Destination
		wantErr         error
	}{
		"invalid_format": {
			path:    "/serialNumber",
			wantErr: fmt.Errorf("invalid path in url expecting: /serialNumber/service"),
		},
		"happy_flow_with_path": {
			path: "/serialNumber/service/path",
			wantDestination: &plugins.Destination{
				OrganizationSerialNumber: "serialNumber",
				Service:                  "service",
				Path:                     "/path",
			},
		},
		"happy_flow_trailing_slash": {
			path: "/serialNumber/service/",
			wantDestination: &plugins.Destination{
				OrganizationSerialNumber: "serialNumber",
				Service:                  "service",
				Path:                     "/",
			},
		},
		"happy_flow_empty_path": {
			path: "/serialNumber/service",
			wantDestination: &plugins.Destination{
				OrganizationSerialNumber: "serialNumber",
				Service:                  "service",
				Path:                     "",
			}},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			destination, err := parseURLPath(tt.path)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantDestination, destination)
		})
	}

}

func TestParseLocalNLXURL(t *testing.T) {
	destinationURL, err := url.Parse("http://service.00000000000000000001.service.nlx.local/path")
	assert.Nil(t, err)
	destination, err := parseLocalNLXURL(destinationURL)
	assert.Nil(t, err)
	assert.Equal(t, "00000000000000000001", destination.OrganizationSerialNumber)
	assert.Equal(t, "service", destination.Service)
	assert.Equal(t, "/path", destination.Path)

	destinationURL, err = url.Parse("http://service.00000000000000000001.invalid.service.nlx.local/path")
	assert.Nil(t, err)
	_, err = parseLocalNLXURL(destinationURL)
	assert.EqualError(t, err, "invalid hostname expecting: service.serialNumber.services.nlx.local")
}
