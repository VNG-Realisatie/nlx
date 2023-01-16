// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package grpcdirectory

import (
	"context"
	"strings"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-ui-htmx/domain"
)

func (l *Directory) ListParticipants(ctx context.Context, organizationNameFilter string) ([]*domain.Participant, error) {
	response, err := l.client.ListParticipants(ctx, &directoryapi.ListParticipantsRequest{})
	if err != nil {
		return nil, err
	}

	result := []*domain.Participant{}

	for _, s := range response.Participants {
		organizationContainsQuery := strings.Contains(strings.ToLower(s.Organization.Name), strings.ToLower(organizationNameFilter))

		if !organizationContainsQuery {
			continue
		}

		result = append(result, &domain.Participant{
			Name:          s.Organization.Name,
			Since:         s.CreatedAt.AsTime(),
			ServicesCount: uint(s.Statistics.Services),
			InwaysCount:   uint(s.Statistics.Inways),
			OutwaysCount:  uint(s.Statistics.Outways),
		})
	}

	return result, nil
}
