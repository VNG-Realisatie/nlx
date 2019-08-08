// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type key int

const (
	serviceKey key = iota
)

type Service struct {
	ID   int
	Name string
}

func (s *Server) ServiceCtx(next http.Handler) http.Handler {
	selectServiceStatement, err := s.db.Preparex(`
		SELECT
			s.id AS id,
			s.name AS name
		FROM directory.services s
		INNER JOIN directory.organizations o ON o.id = s.organization_id
        WHERE o.name = $1 AND s.name = $2
	`)

	if err != nil {
		s.logger.Fatal("Error preparing selectServiceStatement")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		organizationName := chi.URLParam(r, "organization_name")
		serviceName := chi.URLParam(r, "service_name")

		var service Service
		err := selectServiceStatement.QueryRowx(organizationName, serviceName).StructScan(&service)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), serviceKey, service)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
