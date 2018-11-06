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
		return nil, nil, errors.Wrapf(err, "failed to read certificate PEM file `%s`", options.OrgCertFile)
	}
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, nil, errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse certificate")
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}
	if _, err := cert.Verify(opts); err != nil {
		return nil, nil, errors.Wrap(err, "failed to verify certificate")
	}

	return roots, cert, nil
}

// LoadRootCert loads the certificate from file and adds it to a new x509.CertPool which is returned.
func LoadRootCert(rootCertFile string) (*x509.CertPool, error) {
	roots := x509.NewCertPool()
	rootPEM, err := ioutil.ReadFile(rootCertFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read CA root PEM file `%s`", rootCertFile)
	}
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		return nil, errors.Errorf("failed to parse CA root PEM from file `%s`", rootCertFile)
	}
	return roots, nil
}
