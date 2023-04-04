// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"net/http"
)

type participantsPage struct {
	Location      string
	Environment   string
	Introduction  participantsPageIntroduction
	SearchResults ParticipantsSearchResults
}

type participantsPageIntroduction struct {
	Title       string
	Description string
}

func (p *participantsPage) render(w http.ResponseWriter) error {
	baseTemplate := templateWithSVGHelper()

	t, err := baseTemplate.
		ParseFS(
			tplFolder,
			"templates/base.html",
			"templates/participants.html",
			"templates/partials/participants-search-results.html",
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
