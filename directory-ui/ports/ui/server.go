// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"go.nlx.io/nlx/directory-ui/adapters/logger"
	"go.nlx.io/nlx/directory-ui/app"
)

//go:embed templates/**
var tplFolder embed.FS

type Server struct {
	environment string
	staticPath  string
	app         *app.Application
	logger      logger.Logger
	httpServer  *http.Server
	basePage    *BasePage
}

var validEnvironments = []string{"demo", "preprod", "prod"}

var environmentNameToUrls = map[string]string{
	"demo":    "https://directory.demo.nlx.io/",
	"preprod": "https://directory.preprod.nlx.io/",
	"prod":    "https://directory.prod.nlx.io/",
}

func New(environment, staticPath string, lgr logger.Logger, a *app.Application) (*Server, error) {
	if !slices.Contains(validEnvironments, environment) {
		return nil, fmt.Errorf("invalid environment. options: %s", strings.Join(validEnvironments, ", "))
	}

	if lgr == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}

	if a == nil {
		return nil, fmt.Errorf("app cannot be nil")
	}

	basePage, err := NewBasePage(staticPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create base page")
	}

	server := &Server{
		staticPath:  staticPath,
		environment: environment,
		logger:      lgr,
		app:         a,
		basePage:    basePage,
	}

	return server, nil
}

func (s *Server) ListenAndServe(address string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", s.overviewHandler)
	r.Post("/search-services", s.searchServicesHandler)
	r.Get("/participants", s.participantsHandler)
	r.Post("/search-participants", s.searchParticipantsHandler)
	r.Get("/{organizationSerialNumber}/{serviceName}/", s.serviceDetailHandler)

	filesDir := http.Dir(s.staticPath)
	r.Handle("/*", http.FileServer(filesDir))

	const readHeaderTimeout = 5 * time.Second

	s.httpServer = &http.Server{
		Addr:              address,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	err := s.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}
