// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package domain

import (
	"context"
)

type Repository interface {
	ListServices(ctx context.Context, filters ListServicesFilters) ([]*Service, error)
	ListParticipants(ctx context.Context, organizationNameFilter string) ([]*Participant, error)

	Shutdown() error
}

type ListServicesFilters struct {
	Query                  string
	IncludeOfflineServices bool
}
