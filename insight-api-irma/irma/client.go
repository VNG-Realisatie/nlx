package irma

import (
	"crypto/rsa"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Client wraps the irma api server
type Client struct {
	logger *zap.Logger

	ServiceProviderName string

	IRMAEndpointURL    *url.URL
	RSASignPrivateKey  *rsa.PrivateKey
	RSAVerifyPublicKey *rsa.PublicKey
}

// NewClient prepares a new client for use
func NewClient(logger *zap.Logger, serviceProviderName string, irmaEndpointURL string, rsaSignPrivateKey *rsa.PrivateKey, rsaVerifyPublicKey *rsa.PublicKey) (*Client, error) {
	client := &Client{
		logger:              logger,
		ServiceProviderName: serviceProviderName,
		RSASignPrivateKey:   rsaSignPrivateKey,
		RSAVerifyPublicKey:  rsaVerifyPublicKey,
	}

	var err error
	client.IRMAEndpointURL, err = url.Parse(irmaEndpointURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse irmaEndpointURL")
	}

	return client, nil
}

func (c *Client) jwtVerifyKeyFunc(*jwt.Token) (interface{}, error) {
	return c.RSAVerifyPublicKey, nil
}
