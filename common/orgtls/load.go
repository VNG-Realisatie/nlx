// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package orgtls

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Load loads the root certs and own cert/key
func Load(options TLSOptions) (*x509.CertPool, *x509.Certificate, error) {
	roots, err := LoadRootCert(options.NLXRootCert)
	if err != nil {
		return nil, nil, err
	}

	certPEM, err := ioutil.ReadFile(options.OrgCertFile)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open and read organization certificate file `%s`", options.OrgCertFile)
	}
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, nil, errors.Errorf("failed to parse PEM for organization certificate `%s`", options.OrgCertFile)
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to parse x509 for organization certificate `%s`", options.OrgCertFile)
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}
	if _, err := cert.Verify(opts); err != nil {
		return nil, nil, errors.Wrap(err, "failed to verify certificate: not signed by root CA")
	}

	return roots, cert, nil
}

// LoadRootCert loads the certificate from file and adds it to a new x509.CertPool which is returned.
func LoadRootCert(rootCertFile string) (*x509.CertPool, error) {
	roots := x509.NewCertPool()
	rootPEM, err := ioutil.ReadFile(rootCertFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read root CA certificate file `%s`", rootCertFile)
	}
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		return nil, errors.Errorf("failed to parse PEM for root certificate `%s`", rootCertFile)
	}
	return roots, nil
}
