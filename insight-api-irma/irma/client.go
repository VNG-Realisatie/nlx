package irma

import (
	"crypto/rsa"
	"net/url"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Client wraps the irma api server
type Client struct {
	logger *zap.Logger

	ServiceProviderName string

	IRMAEndpointURL *url.URL
	RSASigningKey   *rsa.PrivateKey
}

// NewClient prepares a new client for use
func NewClient(logger *zap.Logger, serviceProviderName string, irmaEndpointURL string, rsaSigningKey *rsa.PrivateKey) (*Client, error) {
	client := &Client{
		logger:              logger,
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
