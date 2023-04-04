// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package query

import (
	"context"

	"go.nlx.io/nlx/directory-ui/domain"
)

type ListParticipantsHandler struct {
	participantRepository domain.Repository
}

func NewListParticipantsHandler(repository domain.Repository) *ListParticipantsHandler {
	return &ListParticipantsHandler{
		participantRepository: repository,
	}
}

func (l *ListParticipantsHandler) Handle(ctx context.Context, organizationNameFilter string) ([]*Participant, error) {
	participants, err := l.participantRepository.ListParticipants(ctx, organizationNameFilter)
	if err != nil {
		return nil, err
	}

	result := make([]*Participant, len(participants))

	for i, s := range participants {
		result[i] = &Participant{
			Name:          s.Name,
			Since:         s.Since,
			ServicesCount: s.ServicesCount,
			InwaysCount:   s.InwaysCount,
			OutwaysCount:  s.OutwaysCount,
		}
	}

	return result, nil
}
