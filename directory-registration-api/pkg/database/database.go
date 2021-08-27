// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import "context"

type DirectoryDatabase interface {
	SetOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error
	ClearOrganizationInway(ctx context.Context, organizationName string) error
	Shutdown()
}
