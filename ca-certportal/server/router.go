// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudflare/cfssl/info"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.uber.org/zap"

	certportal "go.nlx.io/nlx/ca-certportal"
)

type CertPortal struct {
	logger        *zap.Logger
	router        chi.Router
	listenAddress string
	httpServer    *http.Server
}

func setSecurityHeadersHandler(next http.Handler) http.Handler {
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

func NewCertPortal(l *zap.Logger, createSigner certportal.CreateSignerFunc, serialNumberGenerator certportal.SerialNumberGeneratorFunc, listenAddress string) *CertPortal {
	i := &CertPortal{
		logger:        l,
		listenAddress: listenAddress,
	}

	r := chi.NewRouter()
	r.Use(setSecurityHeadersHandler)
	r.Route("/api", func(r chi.Router) {
		r.Post("/request_certificate", requestCertificateHandler(i.logger, createSigner, serialNumberGenerator))
	})

	r.Get("/root.crt", rootCertHandler(i.logger, createSigner))

	workDir, err := os.Getwd()
	if err != nil {
		l.Fatal("failed to get working directory")
	}

	filesDir := filepath.Join(workDir, "public")
	r.Get("/*", http.FileServer(http.Dir(filesDir)).ServeHTTP)

	i.router = r

	i.httpServer = &http.Server{
		Addr:    listenAddress,
		Handler: r,
	}

	return i
}

func (c *CertPortal) Run() error {
	c.logger.Info("starting certportal")

	err := c.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (c *CertPortal) Shutdown(ctx context.Context) error {
	c.logger.Debug("shutting down")

	return c.httpServer.Shutdown(ctx)
}

func (c *CertPortal) GetRouter() chi.Router {
	return c.router
}

func requestCertificateHandler(logger *zap.Logger, createSigner certportal.CreateSignerFunc, serialNumberGenerator certportal.SerialNumberGeneratorFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &certificateRequest{}
		if err := render.DecodeJSON(r.Body, data); err != nil {
			http.Error(w, "could not decode request body", http.StatusBadRequest)
			return
		}

		certificate, err := certportal.RequestCertificate(data.Csr, createSigner, serialNumberGenerator)
		if err != nil {
			switch err {
			case certportal.ErrFailedToSignCSR:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				logger.Error("failed to sign csr", zap.Error(err))

				return

			case certportal.ErrFailedToParseCSR:
				w.WriteHeader(http.StatusBadRequest)
				http.Error(w, "failed to parse csr", http.StatusBadRequest)

				return

			case certportal.ErrFailedToCreateSigner:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				logger.Error("failed to create signer", zap.Error(err))

				return

			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				logger.Error("failed to request certificate", zap.Error(err))

				return
			}
		}

		render.Status(r, http.StatusCreated)
		render.SetContentType(render.ContentTypeJSON)

		err = render.Render(w, r, &certificateResponse{
			Certificate: string(certificate),
		})
		if err != nil {
			logger.Error("error rendering response", zap.Error(err))
		}
	}
}

type certificateRequest struct {
	Csr string `json:"csr"`
}

type certificateResponse struct {
	Certificate string `json:"certificate"`
}

func (rd *certificateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func rootCertHandler(logger *zap.Logger, createSigner certportal.CreateSignerFunc) func(http.ResponseWriter, *http.Request) {
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
