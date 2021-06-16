// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type CertPortal struct {
	logger *zap.Logger
	router chi.Router
}

func SetSecurityHeadersHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=0, no-cache, no-store, must-revalidate")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; frame-ancestors 'none'")
		w.Header().Set("Permissions-Policy", "accelerometer=(), ambient-light-sensor=(), autoplay=(), battery=(), camera=(), display-capture=(), document-domain=(), encrypted-media=(), execution-while-not-rendered=(), execution-while-out-of-viewport=(), fullscreen=(), geolocation=(), gyroscope=(), layout-animations=(), legacy-image-formats=(), magnetometer=(), microphone=(), midi=(), navigation-override=(), oversized-images=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), sync-xhr=(), usb=(), vr=(), wake-lock=(), screen-wake-lock=(), web-share=(), xr-spatial-tracking=()")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Referrer-Policy", "same-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	})
}

func NewCertPortal(l *zap.Logger, createSigner createSignerFunc) *CertPortal {
	i := &CertPortal{
		logger: l,
	}
	r := chi.NewRouter()
	r.Use(SetSecurityHeadersHandler)
	r.Route("/api", func(r chi.Router) {
		r.Post("/request_certificate", requestCertificateHandler(i.logger, createSigner))
	})

	r.Get("/root.crt", rootCertHandler(i.logger, createSigner))

	workDir, err := os.Getwd()
	if err != nil {
		l.Fatal("failed to get working directory")
	}

	filesDir := filepath.Join(workDir, "public")
	r.Get("/*", http.FileServer(http.Dir(filesDir)).ServeHTTP)

	i.router = r

	return i
}

func (c *CertPortal) GetRouter() chi.Router {
	return c.router
}

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
