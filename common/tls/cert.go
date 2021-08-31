// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls

import (
	"crypto"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	// key file can't be:
	// - world readable (0o004)
	// - world writable (0o002)
	// - any executable (0o111)
	invalidPermissions = 0o117
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

func (c *CertificateBundle) PrivateKey() crypto.PrivateKey {
	return c.keyPair.PrivateKey
}

func (c *CertificateBundle) PublicKeyFingerprint() string {
	return c.publicKeyFingerprint
}

func (c *CertificateBundle) PublicKey() crypto.PublicKey {
	return c.Certificate().PublicKey
}

func (c *CertificateBundle) PublicKeyPEM() (string, error) {
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: c.Certificate().RawSubjectPublicKeyInfo,
	})

	if publicKeyPEM == nil {
		return "", errors.New("invalid public key")
	}

	return string(publicKeyPEM), nil
}

// TLSConfig returns a new tls.Config with the certifcate and root ca
func (c *CertificateBundle) TLSConfig(options ...ConfigOption) *tls.Config {
	config := NewConfig(options...)
	config.Certificates = []tls.Certificate{*c.keyPair}
	config.RootCAs = c.rootCAs

	return config
}

// VerifyPrivateKeyPermissions verifies if a file has its permissions configured in way we
// deem as safe.
func VerifyPrivateKeyPermissions(unsanitizedPath string) error {
	file, err := os.Open(filepath.Clean(unsanitizedPath))
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	perms := stat.Mode().Perm()

	if perms&invalidPermissions != 0 {
		return fmt.Errorf(
			"file permissions too open. the file should not allow execution or be readable and writeable for everybody",
		)
	}

	return nil
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
		return nil, fmt.Errorf("failed to parse certificate/key pair: %v", err)
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
		publicKeyFingerprint: X509PublicKeyFingerprint(keyPair.Leaf),
	}

	return bundle, nil
}

// createFingerprint creates a base64-ed sha256 fingerprint for a []byte
func createFingerprint(bytes []byte) string {
	sum := sha256.Sum256(bytes)
	return base64.StdEncoding.EncodeToString(sum[:])
}

// X509PublicKeyFingerprint generates the base64 encoded fingerprint of the Subject PemPublicKeyFingerprint Key Information (SPKI)
func X509PublicKeyFingerprint(certificate *x509.Certificate) string {
	return createFingerprint(certificate.RawSubjectPublicKeyInfo)
}

// PemPublicKeyFingerprint generates the base64 encoded fingerprint of the Subject PemPublicKeyFingerprint Key Information (SPKI)
func PemPublicKeyFingerprint(pemBytes []byte) (string, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return "", fmt.Errorf("unable to decode pem for certificate")
	}

	return createFingerprint(block.Bytes), nil
}

// NewConfig returns a new tls.Config with sane defaults
func NewConfig(options ...ConfigOption) *tls.Config {
	config := &tls.Config{
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
	}

	for _, option := range options {
		option(config)
	}

	return config
}

func ParseCertificate(certPEM []byte) (*x509.Certificate, error) {
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
