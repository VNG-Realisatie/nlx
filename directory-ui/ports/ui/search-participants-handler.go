// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"context"
	"net/http"
)

func (s *Server) searchParticipantsHandler(w http.ResponseWriter, r *http.Request) {
	environment := r.PostFormValue("environment")

	if environment != s.environment {
		w.Header().Set(HxRedirectHeader, environmentNameToUrls[environment])
		return
	}

	participants, err := s.app.Queries.ListParticipants.Handle(context.Background(), r.PostFormValue("search"))
	if err != nil {
		s.logger.Error("could not execute list participants query", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	page := searchParticipantsPage{
		BasePage:      s.basePage,
		SearchResults: make([]*ParticipantsSearchResult, len(participants)),
	}

	for i, participant := range participants {
		page.SearchResults[i] = &ParticipantsSearchResult{
			OrganizationName: participant.Name,
			ParticipantSince: participant.Since.Format("06 January 2006"),
			ServicesCount:    uint32(participant.ServicesCount),
			InwaysCount:      uint32(participant.InwaysCount),
			OutwaysCount:     uint32(participant.OutwaysCount),
		}
	}

	err = page.render(w)
	if err != nil {
		s.logger.Error("could not render search participants template", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
