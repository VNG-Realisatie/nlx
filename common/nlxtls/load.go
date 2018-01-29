package nlxtls

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Load loads the root certs and own cert/key
func Load(options TLSOptions) (*x509.CertPool, *x509.Certificate, error) {
	roots := x509.NewCertPool()
	rootPEM, err := ioutil.ReadFile(options.NLXRootCert)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to read NLX CA root PEM file `%s`", options.NLXRootCert)
	}
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		return nil, nil, errors.Errorf("failed to parse NLX CA root PEM from file `%s`", options.NLXRootCert)
	}

	certPEM, err := ioutil.ReadFile(options.OrgCertFile)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to read certificate PEM file `%s`", options.OrgCertFile)
	}
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		return nil, nil, errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse certificate")
	}

	opts := x509.VerifyOptions{
		// DNSName: "mail.google.com",
		Roots: roots,
	}
	if _, err := cert.Verify(opts); err != nil {
		return nil, nil, errors.Wrap(err, "failed to verify certificate")
	}

	return roots, cert, nil
}
