// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

// CertificateBundle bundles a certificate, private key and root certificate pool
type CertificateBundle struct {
	rootCAs *x509.CertPool
	keyPair *tls.Certificate

	// Base64 encoded of the Subject Public Key Information (SPKI) fingerprint
	publicKeyFingerprint string
}

func (c *CertificateBundle) RootCAs() *x509.CertPool {
	return c.rootCAs
}

func (c *CertificateBundle) Certificate() *x509.Certificate {
	return c.keyPair.Leaf
}

func (c *CertificateBundle) PublicKeyFingerprint() string {
	return c.publicKeyFingerprint
}

// TLSConfig returns a new tls.Config with the certifcate and root ca
func (c *CertificateBundle) TLSConfig(options ...ConfigOption) *tls.Config {
	t := &tls.Config{
		Certificates: []tls.Certificate{*c.keyPair},
		RootCAs:      c.rootCAs,
		MinVersion:   tls.VersionTLS12,
	}

	for _, option := range options {
		option(t)
	}

	return t
}

func NewBundleFromFiles(certFile, keyFile, rootCertFile string) (*CertificateBundle, error) {
	certPEM, err := ioutil.ReadFile(filepath.Clean(certFile))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read certificate file")
	}

	keyPEM, err := ioutil.ReadFile(filepath.Clean(keyFile))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read private key file")
	}

	caRootPEM, err := ioutil.ReadFile(filepath.Clean(rootCertFile))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read root certificate file")
	}

	return NewBundle(certPEM, keyPEM, caRootPEM)
}

func NewBundle(certPEM, keyPEM, rootCertPEM []byte) (*CertificateBundle, error) {
	keyPair, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse certificate/key pair")
	}

	// Ignore error, certificate is already parsed by X509KeyPair. We wouldn't have come this far.
	keyPair.Leaf, _ = x509.ParseCertificate(keyPair.Certificate[0])

	rootCAs, rootCertificate, err := NewCertPool(rootCertPEM)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse root CA certificate")
	}

	intermediates := newIntermediatePool(&keyPair)

	opts := x509.VerifyOptions{
		Roots:         rootCAs,
		Intermediates: intermediates,
	}

	if _, err := keyPair.Leaf.Verify(opts); err != nil {
		_, ok := err.(x509.UnknownAuthorityError)
		if ok {
			return nil, fmt.Errorf("failed to verify certificate: certificate is signed by '%s' and not by provided root CA of '%s'", keyPair.Leaf.Issuer.String(), rootCertificate.Subject.String())
		}

		return nil, errors.Wrap(err, "failed to verify certificate")
	}

	bundle := &CertificateBundle{
		rootCAs:              rootCAs,
		keyPair:              &keyPair,
		publicKeyFingerprint: PublicKeyFingerprint(keyPair.Leaf),
	}

	return bundle, nil
}

// PublicKeyFingerprint generates the base64 encoded fingerprint of the Subject Public Key Information (SPKI)
func PublicKeyFingerprint(certificate *x509.Certificate) string {
	sum := sha256.Sum256(certificate.RawSubjectPublicKeyInfo)

	return base64.StdEncoding.EncodeToString(sum[:])
}

func parseCertificate(certPEM []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, fmt.Errorf("unable to decode pem for certificate")
	}

	return x509.ParseCertificate(block.Bytes)
}

func newIntermediatePool(cert *tls.Certificate) *x509.CertPool {
	p := x509.NewCertPool()

	for _, pem := range cert.Certificate[1:] {
		c, err := x509.ParseCertificate(pem)
		if err == nil {
			p.AddCert(c)
		}
	}

	return p
}
