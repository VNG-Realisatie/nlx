// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import "context"

// DirectoryDatabase is the interface for a directory database
type DirectoryDatabase interface {
	ListServices(ctx context.Context) ([]*Service, error)
}
