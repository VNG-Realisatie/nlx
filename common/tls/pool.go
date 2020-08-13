// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls

import (
	"crypto/x509"
	"io/ioutil"

	"github.com/pkg/errors"
)

func NewCertPoolFromFile(file string) (*x509.CertPool, *x509.Certificate, error) {
	pem, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to read certificate file")
	}

	return NewCertPool(pem)
}

func NewCertPool(pem []byte) (*x509.CertPool, *x509.Certificate, error) {
	c, err := parseCertificate(pem)
	if err != nil {
		return nil, nil, err
	}

	p := x509.NewCertPool()
	p.AddCert(c)

	return p, c, nil
}
