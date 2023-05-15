// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

type overviewPage struct {
	*BasePage

	Location      string
	Introduction  overviewPageIntroduction
	Environment   string
	SearchResults ServicesSearchResults
}

type overviewPageIntroduction struct {
	Title       string
	Description string
}

func (p *overviewPage) render(w http.ResponseWriter) error {
	baseTemplate := p.TemplateWithHelpers()

	t, err := baseTemplate.
		ParseFS(
			tplFolder,
			"templates/base.html",
			"templates/overview.html",
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
