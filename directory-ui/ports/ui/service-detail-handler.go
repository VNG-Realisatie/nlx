// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func (s *Server) serviceDetailHandler(w http.ResponseWriter, r *http.Request) {
	services, err := s.app.Queries.ListServices.Handle(context.Background(), "", true)
	if err != nil {
		s.logger.Error("list services", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	searchResults := make([]*ServicesSearchResult, len(services))

	for i, service := range services {
		searchResults[i] = &ServicesSearchResult{
			ServiceName:              service.Name,
			OrganizationSerialNumber: service.OrganizationSerialNumber,
			OrganizationName:         service.OrganizationName,
			APISpecificationType:     service.APISpecificationType,
			IsOnline:                 service.IsOnline,
		}
	}

	page := serviceDetailPage{
		Location:    "/",
		Environment: s.environment,
		Introduction: serviceDetailPageIntroduction{
			Title:       "Directory",
			Description: "In deze NLX directory vindt u een overzicht van alle beschikbare API’s per NLX omgeving (demo, pre-productie en productie).",
		},
		SearchResults:       searchResults,
		ServiceDetailDrawer: ServiceDetailDrawer{},
	}

	for _, service := range services {
		if service.Name != chi.URLParam(r, "serviceName") ||
			service.OrganizationSerialNumber != chi.URLParam(r, "organizationSerialNumber") {
			continue
		}

		specificationURL := ""

		if service.APISpecificationType != "" {
			specificationURL = getSpecificationURL(s.environment, service.OrganizationSerialNumber, service.Name)
		}

		page.ServiceDetailDrawer.ServiceName = service.Name
		page.ServiceDetailDrawer.OrganizationSerialNumber = service.OrganizationSerialNumber
		page.ServiceDetailDrawer.OrganizationName = service.OrganizationName
		page.ServiceDetailDrawer.APISpecificationType = service.APISpecificationType
		page.ServiceDetailDrawer.PublicSupportContact = service.PublicSupportContact
		page.ServiceDetailDrawer.DocumentationURL = service.DocumentationURL
		page.ServiceDetailDrawer.SpecificationURL = specificationURL
		page.ServiceDetailDrawer.IsOnline = service.IsOnline
	}

	err = page.render(w)
	if err != nil {
		s.logger.Error("render service detail page", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func getSpecificationURL(environment, organizationSerialNumber, serviceName string) string {
	environmentURL := environmentNameToUrls[environment]

	qp := url.Values{
		"url": []string{fmt.Sprintf("%sapi/organizations/%s/services/%s/api-spec", environmentURL, organizationSerialNumber, serviceName)},
	}

	return fmt.Sprintf("https://redocly.github.io/redoc/?%s", qp.Encode())
}
