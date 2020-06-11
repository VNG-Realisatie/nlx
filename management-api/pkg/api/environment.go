// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func environmentRoutes(a *API) chi.Router {
	r := chi.NewRouter()
	r.Get("/", a.environmentHandler)

	return r
}

func (a API) environmentHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, a.environment)
}
