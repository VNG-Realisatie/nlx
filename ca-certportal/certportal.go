// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package certportal

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// CertPortal struct
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

// NewCertPortal creates a new CertPortal and sets it up to handle requests.
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
	r.Get("/*", http.HandlerFunc(http.FileServer(http.Dir(filesDir)).ServeHTTP))

	i.router = r

	return i
}

// GetRouter returns the router of the CertPortal
func (c *CertPortal) GetRouter() chi.Router {
	return c.router
}
