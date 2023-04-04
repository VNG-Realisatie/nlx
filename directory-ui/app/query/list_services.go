// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package query

import (
	"context"

	"go.nlx.io/nlx/directory-ui/domain"
)

type ListServicesHandler struct {
	repository domain.Repository
}

func NewListServicesHandler(repository domain.Repository) *ListServicesHandler {
	return &ListServicesHandler{
		repository: repository,
	}
}

func (l *ListServicesHandler) Handle(ctx context.Context, query string, showOffline bool) ([]*Service, error) {
	services, err := l.repository.ListServices(ctx, domain.ListServicesFilters{
		Query:                  query,
		IncludeOfflineServices: showOffline,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*Service, len(services))

	for i, s := range services {
		result[i] = &Service{
			Name:                     s.Name,
			OrganizationName:         s.OrganizationName,
			OrganizationSerialNumber: s.OrganizationSerialNumber,
			IsOnline:                 s.IsOnline,
			APISpecificationType:     s.APISpecificationType,
			PublicSupportContact:     s.PublicSupportContact,
			DocumentationURL:         s.DocumentationURL,
		}
	}

	return result, nil
}
