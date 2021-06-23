// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import "context"

// DirectoryDatabase is the interface for a directory database
type DirectoryDatabase interface {
	RegisterService(params *RegisterServiceParams) error
	SetOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error
	ClearOrganizationInway(ctx context.Context, organizationName string) error
}
