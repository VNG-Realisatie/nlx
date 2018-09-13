package irma

import (
	"crypto/rsa"
	"net/url"

	"github.com/pkg/errors"
)

// Client wraps the irma api server
type Client struct {
	ServiceProviderName string

	IRMAEndpointURL *url.URL
	RSASigningKey   *rsa.PrivateKey
}

// NewClient prepares a new client for use
func NewClient(serviceProviderName string, irmaEndpointURL string, rsaSigningKey *rsa.PrivateKey) (*Client, error) {
	client := &Client{
		ServiceProviderName: serviceProviderName,
		RSASigningKey:       rsaSigningKey,
	}

	var err error
	client.IRMAEndpointURL, err = url.Parse(irmaEndpointURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse irmaEndpointURL")
	}

	return client, nil
}
