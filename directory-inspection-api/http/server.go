// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package http

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	nlxhttp "go.nlx.io/nlx/common/http"
)

type Server struct {
	db         *sqlx.DB
	mux        *chi.Mux
	httpClient *http.Client
	logger     *zap.Logger
}

func NewServer(db *sqlx.DB, caCertPool *x509.CertPool, certKeyPair *tls.Certificate, logger *zap.Logger) *Server {
	h := &Server{
		db:         db,
		httpClient: nlxhttp.NewHTTPClient(caCertPool, certKeyPair),
		logger:     logger,
	}

	h.mux = createRouter(h)

	return h
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Mount(pattern string, h http.Handler) {
	s.mux.Mount(pattern, h)
}
