// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"

	"go.nlx.io/nlx/management-ui-fsc/adapters/logger"
	"go.nlx.io/nlx/management-ui-fsc/app"
	"go.nlx.io/nlx/management-ui-fsc/ports/ui/i18n"
	jsoni18n "go.nlx.io/nlx/management-ui-fsc/ports/ui/i18n/json"
)

//go:embed templates/**
var tplFolder embed.FS

type Server struct {
	locale     string
	staticPath string
	app        *app.Application
	i18n       i18n.I18n
	logger     logger.Logger
	httpServer *http.Server
	basePage   *BasePage
}

func New(locale, staticPath string, lgr logger.Logger, a *app.Application) (*Server, error) {
	if locale == "" {
		return nil, fmt.Errorf("locale must be set (nl/en)")
	}

	if lgr == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}

	if a == nil {
		return nil, fmt.Errorf("app cannot be nil")
	}

	translations, err := jsoni18n.New(locale)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create new json i18n instance")
	}

	basePage, err := NewBasePage(staticPath, translations)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create base page")
	}

	server := &Server{
		locale:     locale,
		staticPath: staticPath,
		logger:     lgr,
		i18n:       translations,
		app:        a,
		basePage:   basePage,
	}

	return server, nil
}

const compressionLevel = 5

func (s *Server) ListenAndServe(address string) error {
	r := chi.NewRouter()
	r.Use(middleware.Compress(compressionLevel))
	r.Use(middleware.Logger)
	r.Get("/", s.loginHandler)
	r.Get("/inways-and-outways", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inways-and-outways/inways", http.StatusMovedPermanently)
	})
	r.Get("/inways-and-outways/inways", s.inwaysHandler)
	r.Get("/services", s.servicesHandler)

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
