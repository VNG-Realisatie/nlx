// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package http

import (
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	directoryapi "go.nlx.io/nlx/directory-api/api"
)

func createRouter(s *Server) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		_, err := io.Copy(w, strings.NewReader(directoryapi.SwaggerJSONDirectory))
		if err != nil {
			s.logger.Error("failed writing response")
		}
	})

	apiSpecHandler, err := newAPISpecHandler(s.httpClient, s.db, s.logger)
	if err != nil {
		s.logger.Fatal("Error creating API spec handler", zap.Error(err))
	}

	r.Route("/api/organizations/{organization_serial_number}/services/{service_name}", func(r chi.Router) {
		r.Use(s.ServiceCtx)
		r.Get("/api-spec", apiSpecHandler.handleFunc())
	})

	return r
}
