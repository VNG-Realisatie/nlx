// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"

	"go.nlx.io/nlx/common/nlxversion"
)

// DirectoryDatabase is the interface for a directory database
type DirectoryDatabase interface {
	ListServices(ctx context.Context, organizationName string) ([]*Service, error)
	RegisterOutwayVersion(ctx context.Context, version nlxversion.Version) error
	ListOrganizations(ctx context.Context) ([]*Organization, error)
	GetOrganizationInwayAddress(ctx context.Context, organizationName string) (string, error)
}
