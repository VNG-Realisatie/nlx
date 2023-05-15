// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import "net/http"

type searchData struct {
	*BasePage

	SearchResults ServicesSearchResults
}

func (p *searchData) render(w http.ResponseWriter) error {
	baseTemplate := p.TemplateWithHelpers()

	t, err := baseTemplate.
		ParseFS(
			tplFolder,
			"templates/partials/services-search-results.html",
		)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "services-search-results.html", p)
	if err != nil {
		return err
	}

	return nil
}
