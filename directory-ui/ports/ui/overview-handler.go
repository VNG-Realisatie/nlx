// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"context"
	"net/http"
)

func (s *Server) overviewHandler(w http.ResponseWriter, _ *http.Request) {
	services, err := s.app.Queries.ListServices.Handle(context.Background(), "", true)
	if err != nil {
		s.logger.Error("overview handler, could not execute list services query", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	searchResults := make([]*ServicesSearchResult, len(services))

	for i, srv := range services {
		searchResults[i] = &ServicesSearchResult{
			ServiceName:              srv.Name,
			OrganizationSerialNumber: srv.OrganizationSerialNumber,
			OrganizationName:         srv.OrganizationName,
			IsOnline:                 srv.IsOnline,
			APISpecificationType:     srv.APISpecificationType,
		}
	}

	page := overviewPage{
		BasePage:    s.basePage,
		Location:    "/",
		Environment: s.environment,
		Introduction: overviewPageIntroduction{
			Title:       "Directory",
			Description: "In deze NLX directory vindt u een overzicht van alle beschikbare API’s per NLX omgeving (demo, pre-productie en productie).",
		},
		SearchResults: searchResults,
	}

	err = page.render(w)
	if err != nil {
		s.logger.Error("could not render overview page", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
