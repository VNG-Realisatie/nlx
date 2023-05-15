// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import "net/http"

type serviceDetailPage struct {
	*BasePage

	Location            string
	Environment         string
	Introduction        serviceDetailPageIntroduction
	SearchResults       ServicesSearchResults
	ServiceDetailDrawer ServiceDetailDrawer
}

type serviceDetailPageIntroduction struct {
	Title       string
	Description string
}

type ServiceDetailDrawer struct {
	ServiceName              string
	OrganizationSerialNumber string
	OrganizationName         string
	APISpecificationType     string
	PublicSupportContact     string
	IsOnline                 bool
	SpecificationURL         string
	DocumentationURL         string
}

func (p *serviceDetailPage) render(w http.ResponseWriter) error {
	baseTemplate := p.TemplateWithHelpers()

	t, err := baseTemplate.
		ParseFS(
			tplFolder,
			"templates/base.html",
			"templates/overview.html",
			"templates/service-detail.html",
			"templates/partials/services-search-results.html",
			"templates/partials/domain-navigation.html",
		)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "base.html", p)
	if err != nil {
		return err
	}

	return nil
}
