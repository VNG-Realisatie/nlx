// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

type searchParticipantsPage struct {
	SearchResults ParticipantsSearchResults
}

func (p *searchParticipantsPage) render(w http.ResponseWriter) error {
	baseTemplate := templateWithSVGHelper()

	t, err := baseTemplate.ParseFS(
		tplFolder,
		"templates/partials/participants-search-results.html",
	)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "participants-search-results.html", p)
	if err != nil {
		return err
	}

	return nil
}
