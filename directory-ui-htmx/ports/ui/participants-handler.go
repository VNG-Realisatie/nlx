// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package uiport

import (
	"context"
	"net/http"
)

func (s *Server) participantsHandler(w http.ResponseWriter, _ *http.Request) {
	participants, err := s.app.Queries.ListParticipants.Handle(context.Background(), "")
	if err != nil {
		s.logger.Error("list participants", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	searchResults := make([]*ParticipantsSearchResult, len(participants))

	for i, participant := range participants {
		searchResults[i] = &ParticipantsSearchResult{
			OrganizationName: participant.Name,
			ParticipantSince: participant.Since.Format("06 January 2006"),
			ServicesCount:    uint32(participant.ServicesCount),
			InwaysCount:      uint32(participant.InwaysCount),
			OutwaysCount:     uint32(participant.OutwaysCount),
		}
	}

	page := participantsPage{
		Location:    "/participants",
		Environment: s.environment,
		Introduction: participantsPageIntroduction{
			Title:       "Directory Deelnemers",
			Description: "In dit overzicht vindt u alle deelnemers aan dit NLX ecosysteem (demo, pre-productie en productie).",
		},
		SearchResults: searchResults,
	}

	err = page.render(w)
	if err != nil {
		s.logger.Error("could not render participants page", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
