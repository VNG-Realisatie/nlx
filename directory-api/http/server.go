// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	nlxhttp "go.nlx.io/nlx/common/http"
	common_tls "go.nlx.io/nlx/common/tls"
)

type Server struct {
	db         *sqlx.DB
	mux        *chi.Mux
	httpClient *http.Client
	logger     *zap.Logger
}

func NewServer(db *sqlx.DB, certificate *common_tls.CertificateBundle, logger *zap.Logger) *Server {
	h := &Server{
		db:         db,
		httpClient: nlxhttp.NewHTTPClient(certificate),
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
