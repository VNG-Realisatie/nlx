// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package certportal

import (
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"io"
	"net/http"

	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

var sanOID = asn1.ObjectIdentifier{2, 5, 29, 17} // subjectAltName

// function type to enable mocking of the signer
type createSignerFunc func() (signer.Signer, error)

func requestCertificateHandler(logger *zap.Logger, createSigner createSignerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &CertificateRequest{}
		if err := render.DecodeJSON(r.Body, data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error("error reading request body", zap.Error(err))

			return
		}

		csr, err := parseCertificateRequest(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error("parse certificate request", zap.Error(err))

			return
		}

		signReq := signer.SignRequest{
			Request: data.Csr,
		}

		if !hasSAN(csr) {
			signReq.Hosts = []string{csr.Subject.CommonName}
		}

		s, err := createSigner()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("error creating certificate signer", zap.Error(err))

			return
		}

		cert, err := s.Sign(signReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("error signing request", zap.Error(err))

			return
		}

		render.Status(r, http.StatusCreated)
		render.SetContentType(render.ContentTypeJSON)

		err = render.Render(w, r, &certificateResponse{
			Certificate: string(cert),
		})
		if err != nil {
			logger.Error("error rendering response", zap.Error(err))
		}
	}
}

func parseCertificateRequest(request *CertificateRequest) (*x509.CertificateRequest, error) {
	block, _ := pem.Decode([]byte(request.Csr))

	if block == nil {
		return nil, errors.New("decoding certificate request as PEM")
	}

	return x509.ParseCertificateRequest(block.Bytes)
}

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

// CertificateRequest contains the csr
type CertificateRequest struct {
	Csr string `json:"csr"`
}

type certificateResponse struct {
	Certificate string `json:"certificate"`
}

func (rd *certificateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func rootCertHandler(logger *zap.Logger, createSigner createSignerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		sign, err := createSigner()
		if err != nil {
			logger.Error("error obtaining root.crt from cfssl root CA", zap.Error(err))
			http.Error(w, "failed to create new cfssl signer", http.StatusInternalServerError)

			return
		}

		resp, err := sign.Info(info.Req{})
		if err != nil {
			logger.Error("error obtaining root.crt from cfssl root CA", zap.Error(err))
			http.Error(w, "failed to obtain root.crt from cfssl root CA", http.StatusInternalServerError)

			return
		}

		_, err = io.WriteString(w, resp.Certificate)
		if err != nil {
			logger.Error("error in sending root certificate as response", zap.Error(err))
			http.Error(w, "error in sending root certificate as response ", http.StatusInternalServerError)
		}
	}
}
