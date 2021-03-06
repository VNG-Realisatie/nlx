// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"

	"go.nlx.io/nlx/common/nlxversion"
)

// RegisterOutwayVersion registers an outway
func (db PostgreSQLDirectoryDatabase) RegisterOutwayVersion(ctx context.Context, version nlxversion.Version) error {
	_, err := db.registerOutwayStatement.Exec(version)
	if err != nil {
		return fmt.Errorf("failed to log the outway version: %v", err)
	}

	return nil
}
