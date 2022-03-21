// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package outway

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
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
	destination, err := parseURLPath("/serialNumber/service/path")
	assert.Nil(t, err)
	assert.Equal(t, "serialNumber", destination.OrganizationSerialNumber)
	assert.Equal(t, "service", destination.Service)
	assert.Equal(t, "path", destination.Path)

	destination, err = parseURLPath("/serialNumber/service")
	assert.Nil(t, err)
	assert.Equal(t, "serialNumber", destination.OrganizationSerialNumber)
	assert.Equal(t, "service", destination.Service)
	assert.Equal(t, "", destination.Path)
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
