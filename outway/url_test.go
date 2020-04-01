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
		url      string
		isNLXUrl bool
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
		assert.Equal(t, test.isNLXUrl, result)
	}
}

func TestParseURLPath(t *testing.T) {
	destination, err := parseURLPath("/organization/service/path")
	assert.Nil(t, err)
	assert.Equal(t, "organization", destination.Organization)
	assert.Equal(t, "service", destination.Service)
	assert.Equal(t, "path", destination.Path)

	_, err = parseURLPath("/organization/service")
	assert.EqualError(t, err, "invalid path in url expecting: /organization/service/path")
}

func TestParseLocalNLXURL(t *testing.T) {
	destinationURL, err := url.Parse("http://service.organization.service.nlx.local/path")
	assert.Nil(t, err)
	destination, err := parseLocalNLXURL(destinationURL)
	assert.Nil(t, err)
	assert.Equal(t, "organization", destination.Organization)
	assert.Equal(t, "service", destination.Service)
	assert.Equal(t, "/path", destination.Path)

	destinationURL, err = url.Parse("http://service.organization.invalid.service.nlx.local/path")
	assert.Nil(t, err)
	_, err = parseLocalNLXURL(destinationURL)
	assert.EqualError(t, err, "invalid hostname expecting: service.organization.services.nlx.local")
}
