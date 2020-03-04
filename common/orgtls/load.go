// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package orgtls

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

// Load loads the root certs and own cert/key
func Load(options TLSOptions) (*x509.CertPool, *tls.Certificate, error) {
	keyPair, err := tls.LoadX509KeyPair(options.OrgCertFile, options.OrgKeyFile)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to load organization certificate '%s", options.OrgCertFile)
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to parse organization certificate '%s", options.OrgCertFile)
	}

	rootCert, err := loadCertificate(options.NLXRootCert)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to load root CA certificate '%s", options.NLXRootCert)
	}

	roots := x509.NewCertPool()
	roots.AddCert(rootCert)

	intermediates := createIntermediatePool(&keyPair)

	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
	}

	if _, err := keyPair.Leaf.Verify(opts); err != nil {
		_, ok := err.(x509.UnknownAuthorityError)
		if ok {
			return nil, nil, fmt.Errorf("failed to verify certificate: certificate is signed by '%s' and not by provided root CA of '%s'", keyPair.Leaf.Issuer.String(), rootCert.Subject.String())
		}

		return nil, nil, errors.Wrap(err, "failed to verify certificate")
	}

	return roots, &keyPair, nil
}

// LoadRootCert loads the certificate from file and adds it to a new x509.CertPool which is returned.
func LoadRootCert(rootCertFile string) (*x509.CertPool, error) {
	rootPEM, err := ioutil.ReadFile(filepath.Clean(rootCertFile))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read root CA certificate file `%s`", rootCertFile)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		return nil, errors.Errorf("failed to parse PEM for root certificate `%s`", rootCertFile)
	}
	return roots, nil
}

func createIntermediatePool(cert *tls.Certificate) *x509.CertPool {
	pool := x509.NewCertPool()

	for _, pem := range cert.Certificate[1:] {
		c, err := x509.ParseCertificate(pem)
		if err == nil {
			pool.AddCert(c)
		}
	}

	return pool
}

func loadCertificate(filePath string) (*x509.Certificate, error) {
	certPEM, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open and read certificate file `%s`", filePath)
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, fmt.Errorf("unable to decode pem for certificate `%s`", filePath)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse x509 for certificate `%s`", filePath)
	}

	return cert, nil
}
