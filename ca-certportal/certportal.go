// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package certportal

import (
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"

	"github.com/cloudflare/cfssl/signer"
)

type CreateSignerFunc func() (signer.Signer, error)

var (
	ErrFailedToParseCSR     = errors.New("failed to parse csr")
	ErrFailedToCreateSigner = errors.New("unable to create signer")
	ErrFailedToSignCSR      = errors.New("failed to sign csr")
)

type Certificate []byte

func RequestCertificate(certificateSigningRequest string, createSigner CreateSignerFunc) (Certificate, error) {
	csr, err := parseCertificateSigningRequest(certificateSigningRequest)
	if err != nil {
		return nil, ErrFailedToParseCSR
	}

	signReq := signer.SignRequest{
		Request: certificateSigningRequest,
	}

	if !hasSAN(csr) {
		signReq.Hosts = []string{csr.Subject.CommonName}
	}

	s, err := createSigner()
	if err != nil {
		return nil, ErrFailedToCreateSigner
	}

	cert, err := s.Sign(signReq)
	if err != nil {
		return nil, ErrFailedToSignCSR
	}

	return cert, nil
}

func parseCertificateSigningRequest(certificateSigningRequest string) (*x509.CertificateRequest, error) {
	block, _ := pem.Decode([]byte(certificateSigningRequest))

	if block == nil {
		return nil, errors.New("failed to decode certificate request as PEM")
	}

	return x509.ParseCertificateRequest(block.Bytes)
}

var sanOID = asn1.ObjectIdentifier{2, 5, 29, 17} // subjectAltName

func hasSAN(csr *x509.CertificateRequest) bool {
	if len(csr.DNSNames) > 0 {
		return true
	}

	for _, extension := range csr.Extensions {
		if extension.Id.Equal(sanOID) {
			return true
		}
	}

	return false
}
