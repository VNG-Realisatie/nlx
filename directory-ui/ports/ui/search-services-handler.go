// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"context"
	"net/http"
)

func (s *Server) searchServicesHandler(w http.ResponseWriter, r *http.Request) {
	environment := r.PostFormValue("environment")

	if environment != s.environment {
		w.Header().Set(HxRedirectHeader, environmentNameToUrls[environment])
		return
	}

	services, err := s.app.Queries.ListServices.Handle(context.Background(), r.PostFormValue("search"), true)
	if err != nil {
		s.logger.Error("could not execute list services query", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	data := searchData{
		BasePage:      s.basePage,
		SearchResults: make([]*ServicesSearchResult, len(services)),
	}

	for i, service := range services {
		data.SearchResults[i] = &ServicesSearchResult{
			ServiceName:              service.Name,
			OrganizationSerialNumber: service.OrganizationSerialNumber,
			OrganizationName:         service.OrganizationName,
			IsOnline:                 service.IsOnline,
			APISpecificationType:     service.APISpecificationType,
		}
	}

	err = data.render(w)
	if err != nil {
		s.logger.Error("could not render search services page", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
