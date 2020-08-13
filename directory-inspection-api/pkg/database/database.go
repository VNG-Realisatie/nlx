// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"

	"go.nlx.io/nlx/common/nlxversion"
)

// DirectoryDatabase is the interface for a directory database
type DirectoryDatabase interface {
	// TODO: rename to GetServicesForOrganization?
	ListServices(ctx context.Context, organizationName string) ([]*Service, error)
	// TODO: replace nlx version with separate version & component properties
	RegisterOutwayVersion(ctx context.Context, version nlxversion.NlxVersion) error
}
