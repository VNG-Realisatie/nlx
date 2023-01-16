// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package grpcdirectory

import (
	"context"
	"strings"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-ui-htmx/domain"
)

func (l *Directory) ListServices(ctx context.Context, filters domain.ListServicesFilters) ([]*domain.Service, error) {
	response, err := l.client.ListServices(ctx, &directoryapi.ListServicesRequest{})
	if err != nil {
		return nil, err
	}

	result := []*domain.Service{}

	for _, service := range response.Services {
		serviceContainsQuery := strings.Contains(strings.ToLower(service.Name), strings.ToLower(filters.Query))
		organizationContainsQuery := strings.Contains(strings.ToLower(service.Organization.Name), strings.ToLower(filters.Query))

		if !serviceContainsQuery && !organizationContainsQuery {
			continue
		}

		isOnline := len(service.Inways) > 0 && service.Inways[0].State == directoryapi.Inway_STATE_UP

		result = append(result, &domain.Service{
			Name:                     service.Name,
			OrganizationName:         service.Organization.Name,
			OrganizationSerialNumber: service.Organization.SerialNumber,
			IsOnline:                 isOnline,
			APISpecificationType:     service.ApiSpecificationType,
			PublicSupportContact:     service.PublicSupportContact,
			DocumentationURL:         service.DocumentationUrl,
		})
	}

	return result, nil
}
