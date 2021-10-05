// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"

	"go.nlx.io/nlx/common/nlxversion"
)

// DirectoryDatabase is the interface for a directory database
type DirectoryDatabase interface {
	ListServices(ctx context.Context, organizationSerialNumber string) ([]*Service, error)
	RegisterOutwayVersion(ctx context.Context, version nlxversion.Version) error
	ListOrganizations(ctx context.Context) ([]*Organization, error)
	GetOrganizationInwayAddress(ctx context.Context, organizationSerialNumber string) (string, error)
	ListVersionStatistics(ctx context.Context) ([]*VersionStatistics, error)

	Shutdown() error
}
