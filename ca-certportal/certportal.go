// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package certportal

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/sign"
	"github.com/cloudflare/cfssl/signer"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// CertPortal struct
type CertPortal struct {
	logger *zap.Logger
	caHost string
	router chi.Router
}

// NewCertPortal creates a new CertPortal and sets it up to handle requests.
func NewCertPortal(l *zap.Logger, caHost string) *CertPortal {
	i := &CertPortal{
		logger: l,
		caHost: caHost,
	}
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/request_certificate", requestCertificateHandler(i.logger, i.createSigner))
	})

	r.Get("/root.crt", rootCertHandler(i.logger, i.createSigner))

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(r, "/", http.Dir(filesDir))
	i.router = r
	return i
}

// GetRouter returns the router of the CertPortal
func (c *CertPortal) GetRouter() chi.Router {
	return c.router
}

func (c *CertPortal) createSigner() (signer.Signer, error) {
	signer, err := sign.SignerFromConfig(cli.Config{
		Remote: c.caHost,
	})
	if err != nil {
		return nil, err
	}

	return signer, nil
}
